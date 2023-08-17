package compoundv3

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
	"yield-arb/cmd/accounts"
	t "yield-arb/cmd/protocols/types"
	txs "yield-arb/cmd/transactions"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CompoundV3 struct {
	chain          string
	cl             *ethclient.Client
	chainID        *big.Int
	configAddress  common.Address
	configContract *Configurator
	cometAddress   common.Address
	cometContract  *Comet
}

const CompoundV3Name = "CompoundV3"

var compv3ConfigAddresses = map[string]string{
	"arbitrum":        "0xb21b06D71c75973babdE35b49fFDAc3F82Ad3775",
	"arbitrum_goerli": "0x1Ead344570F0f0a0cD86d95d8adDC7855C8723Fb",
}

var compv3CometAddresses = map[string]string{
	"arbitrum":        "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA",
	"arbitrum_goerli": "0x1d573274E19174260c5aCE3f2251598959d24456",
	"polygon":         "0xF25212E676D1F7F89Cd72fFEe66158f541246445",
	"polygon_mumbai":  "0xF09F0369aB0a875254fB565E52226c88f10Bc839",
}

var decimals = map[string]uint8{
	"ETH":    18,
	"USDC.e": 8,
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

	// Instantiate configurator
	configAddress := common.HexToAddress(compv3ConfigAddresses[chain])
	configContract, err := NewConfigurator(configAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate configurator: %v", err)
		return err
	}

	// Instantiate comet
	cometAddress := common.HexToAddress(compv3CometAddresses[chain])
	cometContract, err := NewComet(cometAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate comet: %v", err)
		return err
	}

	c.chain = chain
	c.cl = cl
	c.chainID = chainId
	c.configAddress = configAddress
	c.configContract = configContract
	c.cometAddress = cometAddress
	c.cometContract = cometContract

	log.Printf("%v connected to %v (chainid: %v)", CompoundV3Name, c.chain, c.chainID)
	return nil
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
		// Prices have 8 decimals for compv3
		price := new(big.Int).Quo(returnData.Data, big.NewInt(100000000))
		result[i] = new(big.Float).SetInt(price)
	}

	return result, nil
}

// Returns the amount supplied for each token.
func (c *CompoundV3) getTotalsCollateral(assets []CometConfigurationAssetConfig) ([]*big.Int, error) {
	// Aggregate calldata
	cometABI, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comet abi: %v", err)
	}
	numAssets := len(assets)
	calls := make([]txs.Multicall3Call3, numAssets)
	for i, asset := range assets {
		data, err := cometABI.Pack("totalsCollateral", asset.Asset)
		if err != nil {
			return nil, fmt.Errorf("failed to pack totals collateral calldata: %v", err)
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
		TotalSupplyAsset *big.Int
		Reserved         *big.Int
	}
	result := make([]*big.Int, numAssets)
	for i, response := range *responses {
		returnData := new(ReturnData)
		if err := cometABI.UnpackIntoInterface(returnData, "totalsCollateral", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack totalsCollateral: %v", err)
		}
		result[i] = returnData.TotalSupplyAsset
	}

	return result, nil
}

/*
Seconds Per Year = 60 * 60 * 24 * 365
Utilization = getUtilization()
Supply Rate = getSupplyRate(Utilization)
Supply APR = Supply Rate / (10 ^ 18) * Seconds Per Year * 100

SupplyRate, BorrowRate, BorrowCap (available base)
*/
func (c *CompoundV3) getBaseStats() (*big.Float, *big.Float, *big.Float, error) {
	// Get utilization
	utilization, err := c.cometContract.GetUtilization(nil)
	if err != nil {
		log.Printf("Failed to get utilization: %v", err)
		return nil, nil, nil, err
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

	// Get total supply
	totalSupply, err := c.cometContract.TotalSupply(nil)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to fetch total supply: %v", err)
	}
	utilFloat := new(big.Float).SetInt(utilization)
	utilizationPerc := utilFloat.Quo(utilFloat, utils.ETHMantissa)
	totalSupplyAdjustedInt := totalSupply.Quo(totalSupply, big.NewInt(1000000)) // Base factor for compv3 decimals
	totalSupplyAdjusted := new(big.Float).SetInt(totalSupplyAdjustedInt)
	amountBorrowed := utilizationPerc.Mul(utilizationPerc, totalSupplyAdjusted)
	availableLiquidity := new(big.Float).Sub(totalSupplyAdjusted, amountBorrowed)

	return supplyRate, borrowRate, availableLiquidity, nil
}

// Returns the markets for the protocol
func (c *CompoundV3) GetMarkets() (*t.ProtocolChain, error) {
	log.Printf("Fetching market data for %v...", c.chain)
	startTime := time.Now()

	// Get config
	config, err := c.configContract.GetConfiguration(nil, c.cometAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch configuration: %v", err)
	}

	// Get all token info
	numAssets := len(config.AssetConfigs)

	// Get prices
	pfs := make([]common.Address, numAssets+1)
	for i, assetInfo := range config.AssetConfigs {
		pfs[i] = assetInfo.PriceFeed
	}
	pfs[numAssets] = config.BaseTokenPriceFeed
	prices, err := c.getPrices(pfs)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all prices: %v", err)
	}

	// Get amounts supplied
	amountsSupplied, err := c.getTotalsCollateral(config.AssetConfigs)
	if err != nil {
		return nil, fmt.Errorf("failed to get amounts supplied: %v", err)
	}

	// Fill in LTV and APY for collateral tokens
	supplyMarkets := make([]*t.MarketInfo, numAssets+1)
	for i, assetInfo := range config.AssetConfigs {
		symbols, err := utils.ConvertAddressesToSymbols(c.chain, []string{assetInfo.Asset.Hex()})
		if err != nil {
			return nil, fmt.Errorf("failed to convert symbol: %v", err)
		} else if len(symbols) == 0 {
			return nil, fmt.Errorf("token address mapping missing")
		}
		symbol := symbols[0]
		decimals := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(assetInfo.Decimals)), nil)
		// Has LTV, no APY
		ltv := big.NewFloat(float64(assetInfo.BorrowCollateralFactor))
		ltv.Quo(ltv, big.NewFloat(10000000000000000))
		supplyCapInt := new(big.Int).Sub(assetInfo.SupplyCap, amountsSupplied[i])
		supplyCapInt.Quo(supplyCapInt, decimals)
		supplyCap := new(big.Float).SetInt(supplyCapInt)
		supplyMarkets[i] = &t.MarketInfo{
			Protocol:   CompoundV3Name,
			Chain:      c.chain,
			Token:      symbol,
			Decimals:   assetInfo.Decimals,
			LTV:        ltv,
			SupplyAPY:  big.NewFloat(0),
			BorrowAPY:  big.NewFloat(0),
			SupplyCap:  supplyCap,
			BorrowCap:  big.NewFloat(0),
			PriceInUSD: prices[i],
		}
	}

	// Base token, has APY, no LTV
	symbols, err := utils.ConvertAddressesToSymbols(c.chain, []string{config.BaseToken.Hex()})
	if err != nil {
		return nil, fmt.Errorf("failed to convert base address to token: %v", err)
	}
	supplyAPY, borrowAPY, borrowCap, err := c.getBaseStats()
	if err != nil {
		return nil, fmt.Errorf("failed to get base aprs: %v", err)
	}
	decimals := decimals[symbols[0]]
	market := &t.MarketInfo{
		Protocol:   CompoundV3Name,
		Chain:      c.chain,
		Token:      symbols[0],
		Decimals:   decimals,
		LTV:        big.NewFloat(0),
		SupplyAPY:  supplyAPY,
		BorrowAPY:  borrowAPY,
		SupplyCap:  big.NewFloat(math.MaxFloat64),
		BorrowCap:  borrowCap,
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
func (c *CompoundV3) Supply(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(c.cl, c.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	address, err := utils.ConvertSymbolToAddress(c.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	tx, err := c.cometContract.Supply(auth, common.HexToAddress(address), amount)
	if err != nil {
		return nil, fmt.Errorf("failed to supply: %v", err)
	}

	log.Printf("Supplied %v %v to %v on %v (%v)", amount, token, CompoundV3Name, c.chain, tx.Hash())
	return tx, nil
}

// Borrows the token from the protocol.
// Note: CompoundV3 uses the withdraw function to borrow base asset.
func (c *CompoundV3) Borrow(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(c.cl, c.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	address, err := utils.ConvertSymbolToAddress(c.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	tx, err := c.cometContract.Withdraw(auth, common.HexToAddress(address), amount)
	if err != nil {
		return nil, fmt.Errorf("failed to borrow: %v", err)
	}

	log.Printf("Borrowed %v %v from %v on %v (%v)", amount, token, CompoundV3Name, c.chain, tx.Hash())
	return tx, nil
}

// Withdraws the token from the protocol.
func (c *CompoundV3) Withdraw(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(c.cl, c.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	address, err := utils.ConvertSymbolToAddress(c.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	tx, err := c.cometContract.Withdraw(auth, common.HexToAddress(address), amount)
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw: %v", err)
	}

	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, CompoundV3Name, c.chain, tx.Hash())
	return tx, nil
}

// Repays the token to the protocol.
// Note: CompoundV3 uses the supply function to repay base asset.
func (c *CompoundV3) Repay(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(c.cl, c.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	address, err := utils.ConvertSymbolToAddress(c.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	tx, err := c.cometContract.Supply(auth, common.HexToAddress(address), amount)
	if err != nil {
		return nil, fmt.Errorf("failed to repay: %v", err)
	}

	log.Printf("Repaid %v %v to %v on %v (%v)", amount, token, CompoundV3Name, c.chain, tx.Hash())
	return tx, nil
}
