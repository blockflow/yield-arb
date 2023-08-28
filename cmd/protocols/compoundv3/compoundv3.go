package compoundv3

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"
	"yield-arb/cmd/accounts"
	t "yield-arb/cmd/protocols/types"
	"yield-arb/cmd/transactions"
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

type CompoundV3Params struct {
	SupplyCapRemaining *big.Int
	TotalSupply        *big.Int
	TotalBorrows       *big.Int
	// Constants
	Base      *big.Int
	SlopeLow  *big.Int
	Kink      *big.Int
	SlopeHigh *big.Int
}

type CompoundV3Stats struct {
	TotalSupply     *big.Int
	TotalBorrows    *big.Int
	SupplyBase      *big.Int
	BorrowBase      *big.Int
	SupplySlopeLow  *big.Int
	BorrowSlopeLow  *big.Int
	SupplySlodeHigh *big.Int
	BorrowSlopeHigh *big.Int
}

const CompoundV3Name = "CompoundV3"

var compv3ConfigAddresses = map[string]string{
	"arbitrum":        "0xb21b06D71c75973babdE35b49fFDAc3F82Ad3775",
	"arbitrum_goerli": "0x1Ead344570F0f0a0cD86d95d8adDC7855C8723Fb",
	"ethereum_goerli": "0xB28495db3eC65A0e3558F040BC4f98A0d588Ae60",
}

var compv3CometAddresses = map[string]string{
	"arbitrum":        "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA",
	"arbitrum_goerli": "0x1d573274E19174260c5aCE3f2251598959d24456",
	"ethereum_goerli": "0x3EE77595A8459e93C2888b13aDB354017B198188",
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
func (c *CompoundV3) getPrices(pfs []common.Address) ([]*big.Int, error) {
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
	result := make([]*big.Int, pfLength)
	for i, response := range *responses {
		returnData := new(ReturnData)
		if err := cometABI.UnpackIntoInterface(returnData, "getPrice", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack asset info: %v", err)
		}
		// Prices have 8 decimals for compv3
		result[i] = new(big.Int).Quo(returnData.Data, big.NewInt(100000000))
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
*/
func (c *CompoundV3) getBaseStats() (*CompoundV3Stats, error) {
	// Get utilization
	utilization, err := c.cometContract.GetUtilization(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get utilization: %v", err)
	}

	// Get supply rate
	supplyRateInt, err := c.cometContract.GetSupplyRate(nil, utilization)
	if err != nil {
		return nil, fmt.Errorf("failed to get supply rate: %v", err)
	}
	borrowRateInt, err := c.cometContract.GetBorrowRate(nil, utilization)
	if err != nil {
		return nil, fmt.Errorf("failed to get borrow rate: %v", err)
	}
	log.Print("Util: ", utilization)
	log.Print("Rates: ", supplyRateInt, borrowRateInt)

	// Get total supply
	totalSupply, err := c.cometContract.TotalSupply(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch total supply: %v", err)
	}
	totalBorrows, err := c.cometContract.TotalBorrow(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch total borrows: %v", err)
	}

	// Get base rates
	supplyRateBase, err := c.cometContract.SupplyPerSecondInterestRateBase(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch supply base rate: %v", err)
	}
	borrowRateBase, err := c.cometContract.BorrowPerSecondInterestRateBase(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch borrow base rate: %v", err)
	}

	// Get slope rates
	supplySlopeLow, err := c.cometContract.SupplyPerSecondInterestRateSlopeLow(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch supply slope low: %v", err)
	}
	borrowSlopeLow, err := c.cometContract.BorrowPerSecondInterestRateSlopeLow(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch borrow slope low: %v", err)
	}
	supplySlopeHigh, err := c.cometContract.SupplyPerSecondInterestRateSlopeHigh(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch supply slope high: %v", err)
	}
	borrowSlopeHigh, err := c.cometContract.BorrowPerSecondInterestRateSlopeHigh(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch borrow slope high: %v", err)
	}

	return &CompoundV3Stats{
		TotalSupply:     totalSupply,
		TotalBorrows:    totalBorrows,
		SupplyBase:      supplyRateBase,
		BorrowBase:      borrowRateBase,
		SupplySlopeLow:  supplySlopeLow,
		BorrowSlopeLow:  borrowSlopeLow,
		SupplySlodeHigh: supplySlopeHigh,
		BorrowSlopeHigh: borrowSlopeHigh,
	}, nil
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
		symbol, err := utils.ConvertAddressToSymbol(c.chain, assetInfo.Asset.Hex())
		if err != nil {
			return nil, fmt.Errorf("failed to convert symbol: %v", err)
		}
		// Has LTV, no APY
		ltv := big.NewFloat(float64(assetInfo.BorrowCollateralFactor))
		ltv.Quo(ltv, big.NewFloat(10000000000000000))
		supplyCap := new(big.Int).Sub(assetInfo.SupplyCap, amountsSupplied[i])
		supplyMarkets[i] = &t.MarketInfo{
			Protocol:   CompoundV3Name,
			Chain:      c.chain,
			Token:      symbol,
			Decimals:   big.NewInt(int64(assetInfo.Decimals)),
			LTV:        big.NewInt(int64(assetInfo.BorrowCollateralFactor)),
			PriceInUSD: prices[i],
			Params: &CompoundV3Params{
				SupplyCapRemaining: supplyCap,
				TotalSupply:        amountsSupplied[i],
			},
		}
	}

	// Base token, has APY, no LTV
	symbol, err := utils.ConvertAddressToSymbol(c.chain, config.BaseToken.Hex())
	if err != nil {
		return nil, fmt.Errorf("failed to convert base address to token: %v", err)
	}
	baseStats, err := c.getBaseStats()
	if err != nil {
		return nil, fmt.Errorf("failed to get base aprs: %v", err)
	}
	decimals := decimals[symbol]
	market := &t.MarketInfo{
		Protocol:   CompoundV3Name,
		Chain:      c.chain,
		Token:      symbol,
		Decimals:   big.NewInt(int64(decimals)),
		LTV:        big.NewInt(0),
		PriceInUSD: prices[numAssets],
		Params: &CompoundV3Params{
			SupplyCapRemaining: utils.MaxUint256,
			TotalSupply:        baseStats.TotalSupply,
			TotalBorrows:       baseStats.TotalBorrows,

			Base:      baseStats.SupplyBase,
			SlopeLow:  baseStats.SupplySlopeLow,
			Kink:      big.NewInt(int64(config.SupplyKink)),
			SlopeHigh: baseStats.SupplySlodeHigh,
		},
	}
	supplyMarkets[numAssets] = market
	borrowMarkets := []*t.MarketInfo{{
		Protocol:   CompoundV3Name,
		Chain:      c.chain,
		Token:      symbol,
		Decimals:   big.NewInt(int64(decimals)),
		LTV:        big.NewInt(0),
		PriceInUSD: prices[numAssets],
		Params: &CompoundV3Params{
			SupplyCapRemaining: utils.MaxUint256,
			TotalSupply:        baseStats.TotalSupply,
			TotalBorrows:       baseStats.TotalBorrows,

			Base:      baseStats.BorrowBase,
			SlopeLow:  baseStats.BorrowSlopeLow,
			Kink:      big.NewInt(int64(config.BorrowKink)),
			SlopeHigh: baseStats.BorrowSlopeHigh,
		}}}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return &t.ProtocolChain{
		Protocol:      CompoundV3Name,
		Chain:         c.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}, nil
}

func (c *CompoundV3) CalcAPY(market *t.MarketInfo, amount *big.Int, isSupply bool) (*big.Int, *big.Int, error) {
	prettySpec, _ := json.MarshalIndent(market, "", "  ")
	log.Print(string(prettySpec))
	params, ok := market.Params.(*CompoundV3Params)
	if !ok {
		return nil, nil, fmt.Errorf("failed to cast params to CompoundV3Params")
	}

	actualAmount := amount
	availableLiquidity := new(big.Int).Sub(params.TotalSupply, params.TotalBorrows)
	if isSupply && amount.Cmp(params.SupplyCapRemaining) == 1 {
		actualAmount = params.SupplyCapRemaining
	} else if !isSupply && amount.Cmp(availableLiquidity) == 1 {
		actualAmount = availableLiquidity
	}

	// If not base market (totalBorrows is nil), 0 APY
	if params.TotalBorrows == nil {
		return big.NewInt(0), actualAmount, nil
	}

	// Calc utilization
	supply := params.TotalSupply
	borrows := params.TotalBorrows
	if isSupply {
		supply.Add(supply, actualAmount)
	} else {
		borrows.Add(borrows, actualAmount)
	}
	utilization := new(big.Int).Div(new(big.Int).Mul(borrows, utils.ETHMantissaInt), supply)
	log.Print("Util: ", utilization)

	// Calculate rate per second
	var ratePerSecond *big.Int
	if utilization.Cmp(params.Kink) < 1 {
		ratePerSecond = new(big.Int).Add(params.Base, utils.ManMul(params.SlopeLow, utilization))
	} else {
		ratePerSecond = new(big.Int).Add(params.Base, utils.ManMul(params.SlopeHigh, params.Kink))
		ratePerSecond.Add(ratePerSecond, utils.ManMul(params.SlopeHigh, new(big.Int).Sub(utilization, params.Kink)))
	}
	log.Print("RatePerSecond: ", ratePerSecond)

	// Calculate APY
	apy := new(big.Int).Mul(ratePerSecond, utils.SecPerYear)
	apy.Mul(apy, big.NewInt(1e9)) // Convert to ray

	return apy, actualAmount, nil
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
	tokenAddress := common.HexToAddress(address)
	// Handle approvals
	_, err = transactions.ApproveERC20IfNeeded(c.cl, auth, tokenAddress, walletAddress, c.cometAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to approve: %v", err)
	}
	tx, err := c.cometContract.Supply(auth, tokenAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to supply: %v", err)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(c.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Supplied %v %v to %v on %v (%v)", amount, token, CompoundV3Name, c.chain, tx.Hash())
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

	// Wait
	_, err = transactions.WaitForConfirmations(c.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, CompoundV3Name, c.chain, tx.Hash())
	return tx, nil
}

// Only base asset earns interest.
func (c *CompoundV3) WithdrawAll(wallet string, token string) (*types.Transaction, error) {
	address, err := utils.ConvertSymbolToAddress(c.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	tokenAddress := common.HexToAddress(address)
	walletAddress := common.HexToAddress(wallet)

	// Get balance
	collateral, err := c.cometContract.UserCollateral(nil, walletAddress, tokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch collateral: %v", err)
	}
	return c.Withdraw(wallet, token, collateral.Balance)
}

// Borrows the token from the protocol.
// Note: CompoundV3 uses the withdraw function to borrow base asset.
func (c *CompoundV3) Borrow(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	return c.Withdraw(wallet, token, amount)
}

// Repays the token to the protocol.
// Note: CompoundV3 uses the supply function to repay base asset.
func (c *CompoundV3) Repay(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	return c.Supply(wallet, token, amount)
}

func (c *CompoundV3) RepayAll(wallet string, token string) (*types.Transaction, error) {
	return c.Repay(wallet, token, utils.MaxUint256)
}
