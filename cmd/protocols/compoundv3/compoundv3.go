package compoundv3

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
	t "yield-arb/cmd/protocols/types"
	txs "yield-arb/cmd/transactions"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CompoundV3 struct {
	chain         string
	cl            *ethclient.Client
	chainID       *big.Int
	cometAddress  common.Address
	cometContract *Comet
	baseToken     string // symbol
}

const CompoundV3Name = "CompoundV3"

var compv3CometAddresses = map[string]string{
	"arbitrum":        "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA",
	"arbitrum_goerli": "0x1d573274E19174260c5aCE3f2251598959d24456",
	"polygon":         "0xF25212E676D1F7F89Cd72fFEe66158f541246445",
	"polygon_mumbai":  "0xF09F0369aB0a875254fB565E52226c88f10Bc839",
}

func NewCompoundV3Protocol() *CompoundV3 {
	return &CompoundV3{}
}

// Returns the chains supported by the protocol
func (c *CompoundV3) GetChains() ([]string, error) {
	return []string{"arbitrum", "arbitrum_goerli", "polygon", "polygon_mumbai"}, nil
}

// Connects to the protocol on the specified chain
func (c *CompoundV3) Connect(chain string) error {
	// Setup the client
	rpcEndpoint := utils.ChainConfigs[chain].RPCEndpoint
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Printf("Failed to connect to the %v client: %v", chain, err)
		return err
	}

	// Fetch chainId
	chainId, err := cl.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to fetch chainid: %v", err)
		return err
	}

	// Instantiate comet
	cometAddress := common.HexToAddress(compv3CometAddresses[chain])
	cometContract, err := NewComet(cometAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate comet: %v", err)
		return err
	}

	// Get base token symbol
	baseTokenAddress, err := cometContract.BaseToken(nil)
	if err != nil {
		log.Printf("Failed to fetch base token: %v", err)
		return err
	}

	symbols, err := utils.ConvertAddressesToSymbols(chain, []string{baseTokenAddress.Hex()})
	if err != nil {
		log.Printf("Failed to convert base token: %v", err)
		return err
	}

	c.chain = chain
	c.cl = cl
	c.chainID = chainId
	c.cometAddress = cometAddress
	c.cometContract = cometContract
	c.baseToken = symbols[0]

	log.Printf("%v connected to %v (chainid: %v)", CompoundV3Name, c.chain, c.chainID)
	return nil
}

// Uses multicall to reduce RPC calls
func (c *CompoundV3) getAllAssetInfos() ([]*CometCoreAssetInfo, error) {
	// Fetch number of assets first
	numAssets, err := c.cometContract.NumAssets(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch num assets: %v", err)
	}

	// Aggregate calldata
	cometABI, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comet abi: %v", err)
	}
	calls := make([]txs.Multicall3Call3, numAssets)
	for i := uint8(0); i < numAssets; i++ {
		data, err := cometABI.Pack("getAssetInfo", i)
		if err != nil {
			return nil, fmt.Errorf("failed to pack get asset info calldata: %v", err)
		}
		calls[i] = txs.Multicall3Call3{
			Target:   c.cometAddress,
			CallData: data,
		}
	}

	// Perform multicall
	responses, err := txs.HandleMulticall(c.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall asset info: %v", err)
	}

	// Unpack into CometCoreAssetInfo
	type ReturnData struct {
		Data CometCoreAssetInfo
	}
	result := make([]*CometCoreAssetInfo, numAssets)
	for i, response := range *responses {
		returnData := new(ReturnData)
		if err := cometABI.UnpackIntoInterface(returnData, "getAssetInfo", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack asset info: %v", err)
		}
		result[i] = &returnData.Data
	}

	return result, nil
}

// Uses multicall to reduce RPC calls
func (c *CompoundV3) getPrices(pfs []common.Address) ([]*big.Float, error) {
	pfLength := len(pfs)
	// Aggregate calldata
	cometABI, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comet abi: %v", err)
	}
	calls := make([]txs.Multicall3Call3, pfLength)
	for i, pf := range pfs {
		data, err := cometABI.Pack("getPrice", pf)
		if err != nil {
			return nil, fmt.Errorf("failed to pack get asset info calldata: %v", err)
		}
		calls[i] = txs.Multicall3Call3{
			Target:   c.cometAddress,
			CallData: data,
		}
	}

	// Perform multicall
	responses, err := txs.HandleMulticall(c.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall asset info: %v", err)
	}

	// Unpack into CometCoreAssetInfo
	type ReturnData struct {
		Data *big.Int
	}
	result := make([]*big.Float, pfLength)
	for i, response := range *responses {
		returnData := new(ReturnData)
		if err := cometABI.UnpackIntoInterface(returnData, "getPrice", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack asset info: %v", err)
		}
		result[i] = new(big.Float).SetInt(returnData.Data)
	}

	return result, nil
}

/*
method: getSupplyRate
Seconds Per Year = 60 * 60 * 24 * 365
Utilization = getUtilization()
Supply Rate = getSupplyRate(Utilization)
Supply APR = Supply Rate / (10 ^ 18) * Seconds Per Year * 100

method: getBorrowRate
Seconds Per Year = 60 * 60 * 24 * 365
Utilization = getUtilization()
Borrow Rate = getBorrowRate(Utilization)
Borrow APR = Borrow Rate / (10 ^ 18) * Seconds Per Year * 100
*/
func (c *CompoundV3) getBaseAPRs() (*big.Float, *big.Float, error) {
	// Get utilization
	utilization, err := c.cometContract.GetUtilization(nil)
	if err != nil {
		log.Printf("Failed to get utilization: %v", err)
		return nil, nil, err
	}

	// Get supply rate
	supplyRateInt, err := c.cometContract.GetSupplyRate(nil, utilization)
	if err != nil {
		log.Printf("Failed to get supply rate: %v", err)
	}
	borrowRateInt, err := c.cometContract.GetBorrowRate(nil, utilization)
	if err != nil {
		log.Printf("Failed to get borrow rate: %v", err)
	}
	supplyRate := big.NewFloat(float64(supplyRateInt))
	borrowRate := big.NewFloat(float64(borrowRateInt))

	// Calculate APRs
	supplyRate.Mul(supplyRate, utils.SecPerYear)
	borrowRate.Mul(borrowRate, utils.SecPerYear)
	supplyRate.Mul(supplyRate, big.NewFloat(100))
	borrowRate.Mul(borrowRate, big.NewFloat(100))
	supplyRate.Quo(supplyRate, utils.ETHMantissa)
	borrowRate.Quo(borrowRate, utils.ETHMantissa)

	return supplyRate, borrowRate, nil
}

// Returns the markets for the protocol
func (c *CompoundV3) GetMarkets() (*t.ProtocolChain, error) {
	startTime := time.Now()
	log.Println("Fetching markets...")

	// Get all token info
	allAssetInfos, err := c.getAllAssetInfos()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all asset infos: %v", err)
	}
	numAssets := len(allAssetInfos)

	// Get prices
	pfs := make([]common.Address, numAssets+1)
	for i, assetInfo := range allAssetInfos {
		pfs[i] = assetInfo.PriceFeed
	}
	pfs[numAssets], err = c.cometContract.BaseTokenPriceFeed(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get base token price feed: %v", err)
	}
	prices, err := c.getPrices(pfs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all prices: %v", err)
	}

	// Fill in LTV and APY for collateral tokens
	supplyMarkets := make([]*t.MarketInfo, numAssets+1)
	for i, assetInfo := range allAssetInfos {
		symbols, err := utils.ConvertAddressesToSymbols(c.chain, []string{assetInfo.Asset.Hex()})
		if err != nil {
			return nil, fmt.Errorf("failed to convert symbol: %v", err)
		} else if len(symbols) == 0 {
			return nil, fmt.Errorf("token address mapping missing")
		}
		symbol := symbols[0]
		// Has LTV, no APY
		ltv := big.NewFloat(float64(assetInfo.BorrowCollateralFactor))
		ltv.Quo(ltv, big.NewFloat(10000000000000000))
		supplyMarkets[i] = &t.MarketInfo{
			Protocol:   CompoundV3Name,
			Chain:      c.chain,
			Token:      symbol,
			LTV:        ltv,
			SupplyAPY:  big.NewFloat(0),
			BorrowAPY:  big.NewFloat(0),
			SupplyCap:  new(big.Float).SetInt(assetInfo.SupplyCap),
			BorrowCap:  big.NewFloat(0),
			PriceInUSD: prices[i],
		}
	}

	// Base token, has APY, no LTV
	// TODO: use multicall to bundle these
	supplyAPY, borrowAPY, err := c.getBaseAPRs()
	if err != nil {
		return nil, fmt.Errorf("failed to get base aprs: %v", err)
	}
	market := &t.MarketInfo{
		Protocol:   CompoundV3Name,
		Chain:      c.chain,
		Token:      c.baseToken,
		LTV:        big.NewFloat(0),
		SupplyAPY:  supplyAPY,
		BorrowAPY:  borrowAPY,
		SupplyCap:  big.NewFloat(math.MaxFloat64),
		BorrowCap:  big.NewFloat(0),
		PriceInUSD: prices[numAssets],
	}
	supplyMarkets[numAssets] = market
	borrowMarkets := []*t.MarketInfo{market}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return &t.ProtocolChain{
		Protocol:      CompoundV3Name,
		Chain:         c.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}, nil
}

// Lends the token to the protocol
func (c *CompoundV3) Supply(from common.Address, token string, amount *big.Int) (*types.Transaction, error) {
	return nil, nil
}

func (c *CompoundV3) Borrow(from common.Address, token string, amount *big.Int) (*types.Transaction, error) {
	return nil, nil
}

// TODO: Handle ETH transactions differently!!!
