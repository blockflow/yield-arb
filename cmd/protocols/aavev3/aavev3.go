package aavev3

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"
	"yield-arb/cmd/accounts"
	"yield-arb/cmd/transactions"
	"yield-arb/cmd/utils"

	"yield-arb/cmd/protocols/schema"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Account for new features such as isolation mode

type AaveV3 struct {
	chain                    string
	cl                       *ethclient.Client
	chainID                  *big.Int
	addressesProviderAddress common.Address
	poolAddress              common.Address
	poolContract             *AaveV3Pool
	uiPoolDataProviderCaller *AaveV3UIPoolDataProviderCaller
	wethGatewayTransactor    *WETHGatewayTransactor
}

// type AaveV3MarketParams struct {
// 	// Self calculated
// 	SupplyCapRemaining *big.Int
// 	TotalVariableDebt  *big.Int
// 	TotalStableDebt    *big.Int

// 	ReserveFactor          *big.Int
// 	AvailableLiquidity     *big.Int
// 	AverageStableRate      *big.Int
// 	StableRateSlope1       *big.Int
// 	StableRateSlope2       *big.Int
// 	VariableRateSlope1     *big.Int
// 	VariableRateSlope2     *big.Int
// 	BaseStableBorrowRate   *big.Int
// 	BaseVariableBorrowRate *big.Int
// 	OptimalRatio           *big.Int
// 	Unbacked               *big.Int

// 	// Constants not provided by UIPoolDataProvider
// 	OptimalStableToTotalDebtRatio *big.Int
// 	StableRateExcessOffset        *big.Int
// }

type ReserveData struct {
	Unbacked                *big.Int
	AccruedToTreasuryScaled *big.Int
	TotalAToken             *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}

const AaveV3Name = "aavev3"

var gasLimit = map[string]uint64{
	"ethereum": uint64(300000),
	"arbitrum": uint64(500000),
	"base":     uint64(300000),
}

var aavev3AddressesProviders = map[string]string{
	"ethereum":  "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e",
	"polygon":   "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"avalanche": "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"arbitrum":  "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"optimism":  "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"base":      "0xe20fCBdBfFC4Dd138cE8b2E6FBb6CB49777ad64D",
	// "fantom",
	// "harmony",
	// "metis",

	// Testnets
	"ethereum_goerli": "0xC911B590248d127aD18546B186cC6B324e99F02c",
	"avalanche_fuji":  "0x220c6A7D868FC38ECB47d5E69b99e9906300286A",
	"arbitrum_goerli": "0x4EEE0BB72C2717310318f27628B3c8a708E4951C",
	"optimism_goerli": "0x0b8FAe5f9Bf5a1a5867FB5b39fF4C028b1C2ebA9",
	"polygon_mumbai":  "0xeb7A892BB04A8f836bDEeBbf60897A7Af1Bf5d7F",
}
var uiPoolDataProviders = map[string]string{
	"ethereum":  "0x91c0eA31b49B69Ea18607702c5d9aC360bf3dE7d",
	"polygon":   "0xC69728f11E9E6127733751c8410432913123acf1",
	"avalanche": "0xF71DBe0FAEF1473ffC607d4c555dfF0aEaDb878d",
	"arbitrum":  "0x145dE30c929a065582da84Cf96F88460dB9745A7",
	"optimism":  "0xbd83DdBE37fc91923d59C8c1E0bDe0CccCa332d5",
	"base":      "0x174446a6741300cD2E7C1b1A636Fee99c8F83502",
	// "fantom",
	// "harmony",
	// "metis",

	// Testnets
	"ethereum_goerli": "0xb00A75686293Fea5DA122E8361f6815A0B0AF48E",
	"avalanche_fuji":  "0x08D07a855306400c8e499664f7f5247046274C77",
	"arbitrum_goerli": "0x583F04c0C4BDE3D7706e939F3Ea890Be9A20A5CF",
	"optimism_goerli": "0x9277eFbB991536a98a1aA8b735E9D26d887104C1",
	"polygon_mumbai":  "0x928d9A76705aA6e4a6650BFb7E7912e413Fe7341",
}
var poolDataProviders = map[string]string{
	"ethereum": "0x7B4EB56E7CD4b454BA8ff71E4518426369a138a3",
	"arbitrum": "0x69FA688f1Dc47d4B5d8029D5a35FB7a548310654",
	"base":     "0x2d8A3C5677189723C4cB8873CfC9C8976FDF38Ac",
}

var wethGatewayAddresses = map[string]string{
	// Testnets
	"ethereum_goerli": "0x2a498323acad2971a8b1936fd7540596dc9bbacd",
}

func (a *AaveV3) GetChains() ([]string, error) {
	return []string{
		"ethereum",
		"polygon",
		"avalanche",
		"arbitrum",
		"fantom",
		"harmony",
		"optimism",
		"metis",
	}, nil
}

func NewAaveV3Protocol() *AaveV3 {
	return &AaveV3{}
}

func (a *AaveV3) Connect(chain string) error {
	// Setup the client
	rpcEndpoint := utils.ChainConfigs[chain].RPCEndpoint
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Printf("Failed to connect to the %v client: %v", chain, err)
		return err
	}

	// Fetch chainid
	// TODO: Fetch this in multicall with pool address.
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to fetch chainid: %v", err)
		return err
	}

	// Instantiate AddressesProvider
	addressesProviderAddress := common.HexToAddress(aavev3AddressesProviders[chain])
	addressesProviderCaller, err := NewAaveV3AddressesProviderCaller(addressesProviderAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate AddressesProvider: %v", err)
		return err
	}

	// Fetch lending pool address
	lendingPoolAddress, err := addressesProviderCaller.GetPool(nil)
	if err != nil {
		log.Printf("Failed to get lending pool address: %v", err)
		return err
	}

	// Instantiate lending pool
	poolContract, err := NewAaveV3Pool(lendingPoolAddress, cl)
	if err != nil {
		return fmt.Errorf("failed to instantiate lending pool: %v", err)
	}

	// Instantiate UIPoolDataProvider
	uiPoolDataProviderAddress := common.HexToAddress(uiPoolDataProviders[chain])
	uiPoolDataProviderCaller, err := NewAaveV3UIPoolDataProviderCaller(uiPoolDataProviderAddress, cl)
	if err != nil {
		return fmt.Errorf("failed to instantiate UIPoolDataProvider: %v", err)
	}

	// Instantiate WETHGateway
	wethGatewayAddress := common.HexToAddress(wethGatewayAddresses[chain])
	wethGatewayTransactor, err := NewWETHGatewayTransactor(wethGatewayAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate WETHGateway: %v", err)
		return err
	}

	a.chain = chain
	a.cl = cl
	a.chainID = chainid
	a.addressesProviderAddress = addressesProviderAddress
	a.poolAddress = lendingPoolAddress
	a.poolContract = poolContract
	a.uiPoolDataProviderCaller = uiPoolDataProviderCaller
	a.wethGatewayTransactor = wethGatewayTransactor
	log.Printf("%v connected to %v (chainid: %v, pool: %v)", AaveV3Name, a.chain, a.chainID, lendingPoolAddress)
	return nil
}

func (a *AaveV3) getRatios(aggReserveData []IUiPoolDataProviderV3AggregatedReserveData, aggReserveDataLength int) ([]*big.Int, []*big.Int, error) {
	// Fetch OptimalStableToTotalDebtRatio and StableRateExcessOffset
	optimalStableToTotalDebtRatios := make([]*big.Int, aggReserveDataLength)
	stableRateExcessOffsets := make([]*big.Int, aggReserveDataLength)
	interestRateABI, err := InterestRateStrategyMetaData.GetAbi()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch interest rate strategy abi: %v", err)
	}
	// Pack call data
	calls := make([]transactions.Multicall3Call3, 2*aggReserveDataLength)
	debtRatioData, err := interestRateABI.Pack("OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO calldata: %v", err)
	}
	maxExcessData, err := interestRateABI.Pack("MAX_EXCESS_USAGE_RATIO")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack MAX_EXCESS_USAGE_RATIO calldata: %v", err)
	}
	for i, reserveData := range aggReserveData {
		calls[i] = transactions.Multicall3Call3{
			Target:   reserveData.InterestRateStrategyAddress,
			CallData: debtRatioData,
		}
		calls[i+aggReserveDataLength] = transactions.Multicall3Call3{
			Target:   reserveData.InterestRateStrategyAddress,
			CallData: maxExcessData,
		}
	}
	// Make multicall
	responses, err := transactions.HandleMulticall(a.cl, &calls)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to handle multicall: %v", err)
	}
	// Unpack results
	type ReturnData struct {
		Data *big.Int
	}
	for i, response := range *responses {
		returnData := new(ReturnData)
		if i < aggReserveDataLength {
			err = interestRateABI.UnpackIntoInterface(returnData, "OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO", response.ReturnData)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to unpack OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO: %v", err)
			}
			optimalStableToTotalDebtRatios[i] = returnData.Data
		} else {
			err = interestRateABI.UnpackIntoInterface(returnData, "MAX_EXCESS_USAGE_RATIO", response.ReturnData)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to unpack MAX_EXCESS_USAGE_RATIO: %v", err)
			}
			stableRateExcessOffsets[i-aggReserveDataLength] = returnData.Data
		}
	}

	return optimalStableToTotalDebtRatios, stableRateExcessOffsets, nil
}

func (a *AaveV3) getReserveDatas(aggReserveData []IUiPoolDataProviderV3AggregatedReserveData, aggReserveDataLength int) ([]*ReserveData, error) {
	// Pack data
	calls := make([]transactions.Multicall3Call3, aggReserveDataLength)
	poolDataProviderABI, err := PoolDataProviderMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pool data provider abi: %v", err)
	}
	poolDataProviderAddress := common.HexToAddress(poolDataProviders[a.chain])
	for i, reserveData := range aggReserveData {
		data, err := poolDataProviderABI.Pack("getReserveData", reserveData.UnderlyingAsset)
		if err != nil {
			return nil, fmt.Errorf("failed to pack getReserveData calldata: %v", err)
		}
		calls[i] = transactions.Multicall3Call3{
			Target:   poolDataProviderAddress,
			CallData: data,
		}
	}
	// Make multicall
	responses, err := transactions.HandleMulticall(a.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to handle multicall: %v", err)
	}
	// Unpack results
	reserveDatas := make([]*ReserveData, aggReserveDataLength)
	for i, response := range *responses {
		reserveData := make([]*big.Int, 12)
		err = poolDataProviderABI.UnpackIntoInterface(&reserveData, "getReserveData", response.ReturnData)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack getReserveData: %v", err)
		}
		reserveDatas[i] = &ReserveData{
			TotalStableDebt:     reserveData[3],
			TotalVariableDebt:   reserveData[4],
			VariableBorrowIndex: reserveData[10],
		}
	}

	return reserveDatas, nil
}

// Returns the market.
// Assumes lending and borrowing tokens are the same.
func (a *AaveV3) GetMarkets() ([]*schema.ProtocolChain, error) {
	log.Printf("Fetching market data for %v...", a.chain)
	startTime := time.Now()

	// Fetch reserve data for all tokens
	aggReserveData, _, err := a.uiPoolDataProviderCaller.GetReservesData(nil, a.addressesProviderAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reserve data: %v", err)
	}
	aggReserveDataLength := len(aggReserveData)

	// Fetch OptimalStableToTotalDebtRatio and StableRateExcessOffset
	optimalStableToTotalDebtRatios, stableRateExcessOffsets, err := a.getRatios(aggReserveData, aggReserveDataLength)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ratios: %v", err)
	}

	// Fetch reserve data
	reserveDatas, err := a.getReserveDatas(aggReserveData, aggReserveDataLength)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reserve datas: %v", err)
	}

	// Filter out results for specified symbols
	var supplyMarkets []*schema.MarketInfo
	var borrowMarkets []*schema.MarketInfo
	for i, reserveData := range aggReserveData {
		if reserveData.IsPaused {
			continue
		}

		var supplyCap *big.Int
		if reserveData.SupplyCap.Cmp(big.NewInt(0)) == 0 { // Infinite cap
			supplyCap = utils.MaxUint256
		} else {
			decimals := new(big.Int).Exp(big.NewInt(10), reserveData.Decimals, nil)
			amountSupplied := new(big.Int).Add(reserveData.TotalScaledVariableDebt, reserveData.AvailableLiquidity)
			supplyCap = new(big.Int).Mul(reserveData.SupplyCap, decimals)
			supplyCap.Sub(supplyCap, amountSupplied)
		}

		// Check mapping for USDC.e
		symbol, err := utils.ConvertAddressToSymbol(a.chain, reserveData.UnderlyingAsset.Hex())
		if err != nil {
			return nil, fmt.Errorf("failed to convert address to symbol: %v", err)
		}

		market := &schema.MarketInfo{
			Protocol:   AaveV3Name,
			Chain:      a.chain,
			Token:      symbol,
			Decimals:   reserveData.Decimals,
			LTV:        reserveData.BaseLTVasCollateral,
			PriceInUSD: reserveData.PriceInMarketReferenceCurrency,
			Params: map[string]interface{}{
				"supplyCapRemaining": supplyCap,
				"totalVariableDebt":  reserveDatas[i].TotalVariableDebt,
				"totalStableDebt":    reserveDatas[i].TotalStableDebt,

				"reserveFactor":          reserveData.ReserveFactor,
				"availableLiquidity":     reserveData.AvailableLiquidity,
				"averageStableRate":      reserveData.AverageStableRate,
				"stableRateSlope1":       reserveData.StableRateSlope1,
				"stableRateSlope2":       reserveData.StableRateSlope2,
				"variableRateSlope1":     reserveData.VariableRateSlope1,
				"variableRateSlope2":     reserveData.VariableRateSlope2,
				"baseStableBorrowRate":   reserveData.BaseStableBorrowRate,
				"baseVariableBorrowRate": reserveData.BaseVariableBorrowRate,
				"optimalRatio":           reserveData.OptimalUsageRatio,
				"unbacked":               reserveData.Unbacked,

				"optimalStableToTotalDebtRatio": optimalStableToTotalDebtRatios[i],
				"stableRateExcessOffset":        stableRateExcessOffsets[i],
			},
		}
		supplyMarkets = append(supplyMarkets, market)
		borrowMarkets = append(borrowMarkets, market)
	}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return []*schema.ProtocolChain{{
		Protocol:      AaveV3Name,
		Chain:         a.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}}, nil
}

func (*AaveV3) CalcAPY(market *schema.MarketInfo, amount *big.Int, isSupply bool) (*big.Int, *big.Int, error) {
	supplyCapRemaining := market.Params["supplyCapRemaining"].(*big.Int)
	availableLiquidity := market.Params["availableLiquidity"].(*big.Int)

	// Check for caps
	actualAmount := amount
	if isSupply && supplyCapRemaining.Cmp(amount) == -1 {
		actualAmount = supplyCapRemaining
	} else if !isSupply && availableLiquidity.Cmp(amount) == -1 {
		actualAmount = availableLiquidity
	}

	if isSupply {
		currentLiquidityRate, _, _ := calculateInterestRates(market.Params, actualAmount, big.NewInt(0))
		return currentLiquidityRate, actualAmount, nil
	}
	_, _, currentVariableRate := calculateInterestRates(market.Params, big.NewInt(0), actualAmount)
	return currentVariableRate, actualAmount, nil
}

// Default all borrows to variable rate
func calculateInterestRates(params map[string]interface{}, liquidityAdded, liquidityTaken *big.Int) (currentLiquidityRate, currentStableRate, currentVariableRate *big.Int) {
	totalVariableDebt := params["totalVariableDebt"].(*big.Int)
	totalStableDebt := params["totalStableDebt"].(*big.Int)
	reserveFactor := params["reserveFactor"].(*big.Int)
	availableLiquidity := params["availableLiquidity"].(*big.Int)
	averageStableRate := params["averageStableRate"].(*big.Int)
	stableRateSlope1 := params["stableRateSlope1"].(*big.Int)
	stableRateSlope2 := params["stableRateSlope2"].(*big.Int)
	variableRateSlope1 := params["variableRateSlope1"].(*big.Int)
	variableRateSlope2 := params["variableRateSlope2"].(*big.Int)
	baseStableBorrowRate := params["baseStableBorrowRate"].(*big.Int)
	baseVariableBorrowRate := params["baseVariableBorrowRate"].(*big.Int)
	optimalRatio := params["optimalRatio"].(*big.Int)
	unbacked := params["unbacked"].(*big.Int)
	optimalStableToTotalDebtRatio := params["optimalStableToTotalDebtRatio"].(*big.Int)
	stableRateExcessOffset := params["stableRateExcessOffset"].(*big.Int)

	MAX_EXCESS_USAGE_RATIO := new(big.Int).Sub(utils.Ray, optimalRatio)
	MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO := new(big.Int).Sub(utils.Ray, optimalStableToTotalDebtRatio)

	newTotalVariableDebt := new(big.Int).Add(totalVariableDebt, liquidityTaken)
	totalDebt := new(big.Int).Add(totalStableDebt, newTotalVariableDebt)

	currentStableRate = new(big.Int).Set(baseStableBorrowRate)
	currentVariableRate = new(big.Int).Set(baseVariableBorrowRate)

	stableToTotalDebtRatio := big.NewInt(0)
	borrowUsageRatio := big.NewInt(0)
	supplyUsageRatio := big.NewInt(0)

	if totalDebt.Cmp(big.NewInt(0)) != 0 {
		stableToTotalDebtRatio = utils.RayDiv(totalStableDebt, totalDebt)
		availableLiquidity := new(big.Int).Add(availableLiquidity, liquidityAdded)
		availableLiquidity.Sub(availableLiquidity, liquidityTaken)

		availableLiquidityPlusDebt := new(big.Int).Add(availableLiquidity, totalDebt)
		borrowUsageRatio = utils.RayDiv(totalDebt, availableLiquidityPlusDebt)
		supplyUsageRatio = utils.RayDiv(totalDebt, new(big.Int).Add(availableLiquidityPlusDebt, unbacked))
	}

	if borrowUsageRatio.Cmp(optimalRatio) == 1 {
		excessBorrowUsageRatio := new(big.Int).Sub(borrowUsageRatio, optimalRatio)
		excessBorrowUsageRatio = utils.RayDiv(excessBorrowUsageRatio, MAX_EXCESS_USAGE_RATIO)

		currentStableRate.Add(currentStableRate, stableRateSlope1)
		currentStableRate.Add(currentStableRate, utils.RayMul(stableRateSlope2, excessBorrowUsageRatio))

		currentVariableRate.Add(currentVariableRate, variableRateSlope1)
		currentVariableRate.Add(currentVariableRate, utils.RayMul(variableRateSlope2, excessBorrowUsageRatio))
	} else {
		stableRate := utils.RayMul(stableRateSlope1, borrowUsageRatio)
		stableRate = utils.RayDiv(stableRate, optimalRatio)
		currentStableRate.Add(currentStableRate, stableRate)

		variableRate := utils.RayMul(variableRateSlope1, borrowUsageRatio)
		variableRate = utils.RayDiv(variableRate, optimalRatio)
		currentVariableRate.Add(currentVariableRate, variableRate)
	}

	if stableToTotalDebtRatio.Cmp(optimalStableToTotalDebtRatio) == 1 {
		diff := new(big.Int).Sub(stableToTotalDebtRatio, optimalStableToTotalDebtRatio)

		excessStableDebtRatio := utils.RayDiv(diff, MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO)

		currentStableRate.Add(currentStableRate, utils.RayMul(stableRateExcessOffset, excessStableDebtRatio))
	}

	overallBorrowRate := getOverallBorrowRate(
		totalStableDebt,
		newTotalVariableDebt,
		currentVariableRate,
		averageStableRate)

	currentLiquidityRate = utils.RayMul(overallBorrowRate, supplyUsageRatio)
	currentLiquidityRate = utils.PercentMul(currentLiquidityRate,
		new(big.Int).Sub(utils.PercentageFactor, reserveFactor))

	return
}

func getOverallBorrowRate(
	totalStableDebt *big.Int,
	totalVariableDebt *big.Int,
	variableRate *big.Int,
	averageStableRate *big.Int,
) *big.Int {
	totalDebt := new(big.Int).Add(totalStableDebt, totalVariableDebt)
	if totalDebt.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}

	weightedVariableRate := utils.RayMul(utils.WadToRay(totalVariableDebt), variableRate)

	weightedStableRate := utils.RayMul(utils.WadToRay(totalStableDebt), averageStableRate)

	overallBorrowRate := new(big.Int).Add(weightedVariableRate, weightedStableRate)
	overallBorrowRate = utils.RayDiv(overallBorrowRate, utils.WadToRay(totalDebt))

	return overallBorrowRate
}

// // Returns the token balances for the specified wallet.
// func (a *AaveV3) GetBalances(wallet string) (map[string]*big.Int, error) {
// 	provider := common.HexToAddress(aavev3AddressesProviders[a.chain])
// 	walletAddress := common.HexToAddress(wallet)
// 	balances := make(map[string]*big.Int)

// 	// Second result is the user emode category id
// 	userReserves, _, err := a.uiPoolDataProviderCaller.GetUserReservesData(nil, provider, walletAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to fetch user reserves: %v", err)
// 	}

// 	for _, userReserve := range userReserves {
// 		asset := userReserve.UnderlyingAsset
// 		symbol, err := utils.ConvertAddressToSymbol(a.chain, asset.Hex())
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to convert address to symbol: %v", err)
// 		}
// 		// balances[symbol] = userReserve.
// 	}
// }

func (a *AaveV3) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)

	var txs []*types.Transaction
	var txErr error
	address, err := utils.ConvertSymbolToAddress(step.Market.Chain, step.Market.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol: %v", err)
	}
	tokenAddress := common.HexToAddress(address)

	poolABI, err := AaveV3PoolMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pool abi: %v", err)
	}

	var data []byte
	if step.IsSupply {
		approvalTx, err := transactions.GetApprovalTxIfNeeded(a.cl, step.Market.Chain, tokenAddress, walletAddress, a.poolAddress, step.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to get approval tx: %v", err)
		}
		txs = append(txs, approvalTx)

		data, err = poolABI.Pack("supply", tokenAddress, step.Amount, walletAddress, uint16(0))
		if err != nil {
			return nil, fmt.Errorf("failed to pack deposit calldata: %v", err)
		}
	} else {
		data, err = poolABI.Pack("borrow", tokenAddress, step.Amount, big.NewInt(2), uint16(0), walletAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to pack borrow calldata: %v", err)
		}
	}
	gas := gasLimit[step.Market.Chain]
	if !step.IsSupply {
		gas *= 2
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		To:   &a.poolAddress,
		Data: data,
		Gas:  gas,
	})
	txs = append(txs, tx)

	if txErr != nil {
		return nil, fmt.Errorf("failed to send supply tx: %v", txErr)
	}

	return txs, nil
}

// Deposits the specified token into the protocol
func (a *AaveV3) Supply(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth := &bind.TransactOpts{From: walletAddress}

	var tx *types.Transaction
	var txErr error
	address, err := utils.ConvertSymbolToAddress(a.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol: %v", err)
	}
	tokenAddress := common.HexToAddress(address)

	if token == "WETH" {
		// Does not support EIP2612 Permit. Make sure to approve beforehand.
		_, err := transactions.ApproveERC20IfNeeded(a.cl, auth, tokenAddress, walletAddress, a.poolAddress, amount)
		if err != nil {
			return nil, fmt.Errorf("failed to approve: %v", err)
		}

		tx, txErr = a.poolContract.Supply(auth, tokenAddress, amount, walletAddress, uint16(0))
	} else {
		// Sign EIP 2612 permit to use SupplyPermit
		signedPermit, err := accounts.SignEIP2612Permit(a.cl, a.chainID, a.chain, token, walletAddress, a.poolAddress, amount)
		if err != nil {
			return nil, fmt.Errorf("failed to sign permit: %v", err)
		}

		tx, txErr = a.poolContract.SupplyWithPermit(auth, tokenAddress, amount, walletAddress, uint16(0), signedPermit.Deadline, signedPermit.V, [32]byte(signedPermit.R), [32]byte(signedPermit.S))
	}

	if txErr != nil {
		return nil, fmt.Errorf("failed to send supply tx: %v", txErr)
	}

	return tx, nil
}

func (a *AaveV3) Withdraw(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(a.cl, a.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	address, err := utils.ConvertSymbolToAddress(a.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol: %v", err)
	}
	tokenAddress := common.HexToAddress(address)

	// Fetch atoken address
	reserveData, err := a.poolContract.GetReserveData(nil, tokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reserve data: %v", err)
	}
	// Handle approvals
	_, err = transactions.ApproveERC20IfNeeded(a.cl, auth, reserveData.ATokenAddress, walletAddress, a.poolAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to approve: %v", err)
	}

	tx, txErr := a.poolContract.Withdraw(auth, tokenAddress, amount, walletAddress)

	if txErr != nil {
		return nil, fmt.Errorf("failed to send withdraw tx: %v", txErr)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(a.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, AaveV3Name, a.chain, tx.Hash())
	return tx, nil
}

func (a *AaveV3) WithdrawAll(wallet string, token string) (*types.Transaction, error) {
	return a.Withdraw(wallet, token, utils.MaxUint256)
}

// Borrows the specified token from the protocol.
// Defaults to variable interest rates.
func (a *AaveV3) Borrow(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(a.cl, a.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	address, err := utils.ConvertSymbolToAddress(a.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol: %v", err)
	}
	tx, txErr := a.poolContract.Borrow(auth, common.HexToAddress(address), amount, big.NewInt(2), uint16(0), walletAddress)

	if txErr != nil {
		return nil, fmt.Errorf("failed to send borrow tx: %v", txErr)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(a.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Borrowed %v %v from %v on %v (%v)", amount, token, AaveV3Name, a.chain, tx.Hash())
	return tx, nil
}

// Borrows the specified token from the protocol.
// Defaults to variable interest rates.
func (a *AaveV3) Repay(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(a.cl, a.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	var tx *types.Transaction
	var txErr error
	address, err := utils.ConvertSymbolToAddress(a.chain, token)
	tokenAddress := common.HexToAddress(address)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol: %v", err)
	}

	if token == "WETH" {
		// Handle approvals
		_, err := transactions.ApproveERC20IfNeeded(a.cl, auth, tokenAddress, walletAddress, a.poolAddress, amount)
		if err != nil {
			return nil, fmt.Errorf("failed to approve: %v", err)
		}
		tx, txErr = a.poolContract.Repay(auth, tokenAddress, amount, big.NewInt(2), walletAddress)
	} else {
		// Sign EIP 2612 permit to use SupplyPermit
		signedPermit, err := accounts.SignEIP2612Permit(a.cl, a.chainID, a.chain, token, walletAddress, a.poolAddress, amount)
		if err != nil {
			return nil, fmt.Errorf("failed to sign permit: %v", err)
		}

		tx, txErr = a.poolContract.RepayWithPermit(auth, tokenAddress, amount, big.NewInt(2), walletAddress, signedPermit.Deadline, signedPermit.V, [32]byte(signedPermit.R), [32]byte(signedPermit.S))
	}

	if txErr != nil {
		return nil, fmt.Errorf("failed to send repay tx: %v", txErr)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(a.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Repaid %v %v to %v on %v (%v)", amount, token, AaveV3Name, a.chain, tx.Hash())
	return tx, nil
}

func (a *AaveV3) RepayAll(wallet string, token string) (*types.Transaction, error) {
	return a.Repay(wallet, token, utils.MaxUint256)
}
