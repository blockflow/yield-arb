package dforce

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"math"
// 	"math/big"
// 	"time"
// 	"yield-arb/cmd/accounts"
// 	t "yield-arb/cmd/protocols/schema"
// 	"yield-arb/cmd/transactions"
// 	txs "yield-arb/cmd/transactions"
// 	"yield-arb/cmd/utils"

// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/ethclient"
// 	"golang.org/x/exp/slices"
// )

// type DForce struct {
// 	chain              string
// 	cl                 *ethclient.Client
// 	chainID            *big.Int
// 	controllerAddress  common.Address
// 	controllerContract *Controller
// 	oracleAddress      common.Address
// }

// type DForceParams struct {
// 	SupplyCapRemaining *big.Int
// 	TotalSupply        *big.Int
// 	TotalBorrows       *big.Int
// 	TotalReserves      *big.Int
// 	TotalCash          *big.Int
// 	ReserveRatio       *big.Int

// 	Optimal       *big.Int
// 	Base          *big.Int
// 	Slope1        *big.Int
// 	Slope2        *big.Int
// 	BlocksPerYear *big.Int
// }

// type InterestRateConstants struct {
// 	Optimal *big.Int
// 	Base    *big.Int
// 	Slope1  *big.Int
// 	Slope2  *big.Int
// }

// const DForceName = "dforce"

// var controllerAddresses = map[string]string{
// 	"arbitrum": "0x8E7e9eA9023B81457Ae7E6D2a51b003D421E5408",
// }

// // TODO: Change to mapping to support other chains
// var nativeToken = "iETH"

// var ignoreTokens = []common.Address{
// 	common.HexToAddress("0xe8c85B60Cb3bA32369c699015621813fb2fEA56c"), // TODO: iMUSX, no supply rate?
// 	common.HexToAddress("0x5BE49B2e04aC55A17c72aC37E3a85D9602322021"), // TODO: iMEUX, no supply rate?
// }
// var blocksPerYear = map[string]*big.Int{
// 	"arbitrum": big.NewInt(2425846),
// }

// // Returns the chains supported by the protocol
// func (d *DForce) GetChains() ([]string, error) {
// 	return []string{
// 		"arbitrum",
// 	}, nil
// }

// func NewDForceProtocol() *DForce {
// 	return &DForce{}
// }

// func (d *DForce) Connect(chain string) error {
// 	// Setup the client
// 	rpcEndpoint := utils.ChainConfigs[chain].RPCEndpoint
// 	cl, err := ethclient.Dial(rpcEndpoint)
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to the %v client: %v", chain, err)
// 	}

// 	// Fetch chainid
// 	// TODO: Fetch this in multicall with pool address.
// 	chainid, err := cl.ChainID(context.Background())
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch chainid: %v", err)
// 	}

// 	// Instantiate controller
// 	controllerAddress := common.HexToAddress(controllerAddresses[chain])
// 	controllerContract, err := NewController(controllerAddress, cl)
// 	if err != nil {
// 		return fmt.Errorf("failed to instantiate AddressesProvider: %v", err)
// 	}

// 	// Fetch price oracle
// 	oracleAddress, err := controllerContract.PriceOracle(nil)
// 	if err != nil {
// 		return fmt.Errorf("failed to fetch oracle address: %v", err)
// 	}

// 	d.chain = chain
// 	d.cl = cl
// 	d.chainID = chainid
// 	d.controllerAddress = controllerAddress
// 	d.controllerContract = controllerContract
// 	d.oracleAddress = oracleAddress
// 	log.Printf("%v connected to %v (chainid: %v, controller: %v)", DForceName, d.chain, d.chainID, controllerAddress)
// 	return nil
// }

// func (l *DForce) getInterestRateConstants(interestRateModels []common.Address) ([]*InterestRateConstants, error) {
// 	// Fetch interest rate constants
// 	// Pack calldata
// 	calls := make([]transactions.Multicall3Call3, len(interestRateModels)*4)
// 	interestRateModelABI, err := InterestRateModelMetaData.GetAbi()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get interest rate model abi: %v", err)
// 	}
// 	for i, interestRateModel := range interestRateModels {
// 		kinkData, err := interestRateModelABI.Pack("optimal")
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to pack optimal: %v", err)
// 		}
// 		baseRatePerBlockData, err := interestRateModelABI.Pack("base")
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to pack base: %v", err)
// 		}
// 		multiplierPerBlockData, err := interestRateModelABI.Pack("slope_1")
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to pack slope1: %v", err)
// 		}
// 		jumpMultiplierPerBlockData, err := interestRateModelABI.Pack("slope_2")
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to pack slope2: %v", err)
// 		}
// 		calls[i*4] = transactions.Multicall3Call3{
// 			Target:   interestRateModel,
// 			CallData: kinkData,
// 		}
// 		calls[i*4+1] = transactions.Multicall3Call3{
// 			Target:   interestRateModel,
// 			CallData: baseRatePerBlockData,
// 		}
// 		calls[i*4+2] = transactions.Multicall3Call3{
// 			Target:   interestRateModel,
// 			CallData: multiplierPerBlockData,
// 		}
// 		calls[i*4+3] = transactions.Multicall3Call3{
// 			Target:   interestRateModel,
// 			CallData: jumpMultiplierPerBlockData,
// 		}
// 	}
// 	// Perform multicall
// 	responses, err := transactions.HandleMulticall(l.cl, &calls)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to multicall: %v", err)
// 	}
// 	// Unpack responses
// 	type ResponseOutput struct {
// 		Data *big.Int
// 	}
// 	interestRateConstants := make([]*InterestRateConstants, len(interestRateModels))
// 	for i := range interestRateConstants {
// 		rateConstants := new(InterestRateConstants)
// 		output := new(ResponseOutput)
// 		if err := interestRateModelABI.UnpackIntoInterface(output, "optimal", (*responses)[i*4].ReturnData); err != nil {
// 			return nil, fmt.Errorf("failed to unpack %v: %v", "optimal", err)
// 		}
// 		rateConstants.Optimal = output.Data
// 		output = new(ResponseOutput)
// 		if err := interestRateModelABI.UnpackIntoInterface(output, "base", (*responses)[i*4+1].ReturnData); err != nil {
// 			return nil, fmt.Errorf("failed to unpack %v: %v", "base", err)
// 		}
// 		rateConstants.Base = output.Data
// 		output = new(ResponseOutput)
// 		if err := interestRateModelABI.UnpackIntoInterface(output, "slope_1", (*responses)[i*4+2].ReturnData); err != nil {
// 			return nil, fmt.Errorf("failed to unpack %v: %v", "slope1", err)
// 		}
// 		rateConstants.Slope1 = output.Data
// 		output = new(ResponseOutput)
// 		if err := interestRateModelABI.UnpackIntoInterface(output, "slope_2", (*responses)[i*4+3].ReturnData); err != nil {
// 			return nil, fmt.Errorf("failed to unpack %v: %v", "slope2", err)
// 		}
// 		rateConstants.Slope2 = output.Data
// 		interestRateConstants[i] = rateConstants
// 	}

// 	return interestRateConstants, nil
// }

// // Returns the markets for the protocol
// func (d *DForce) GetMarkets() ([]*t.ProtocolChain, error) {
// 	log.Printf("Fetching markets for %v on %v", DForceName, d.chain)
// 	startTime := time.Now()
// 	// Get all tokens
// 	allITokens, err := d.controllerContract.GetAlliTokens(nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get all iTokens: %v", err)
// 	}

// 	// Aggregate calldata
// 	iTokenABI, err := ITokenMetaData.GetAbi()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get iToken abi: %v", err)
// 	}
// 	controllerABI, err := ControllerMetaData.GetAbi()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get controller abi: %v", err)
// 	}
// 	oracleABI, err := OracleMetaData.GetAbi()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get oracle abi: %v", err)
// 	}
// 	var calls []txs.Multicall3Call3
// 	iTokenMethods := []string{"symbol", "decimals", "supplyRatePerBlock", "borrowRatePerBlock", "totalSupply", "getCash", "interestRateModel", "totalBorrows", "totalReserves", "reserveRatio"}
// 	for _, itoken := range allITokens {
// 		if slices.Contains[common.Address](ignoreTokens, itoken) {
// 			continue
// 		}
// 		// iToken calls
// 		for _, method := range iTokenMethods {
// 			data, err := iTokenABI.Pack(method)
// 			if err != nil {
// 				return nil, fmt.Errorf("failed to pack %v calldata: %v", method, err)
// 			}
// 			calls = append(calls, txs.Multicall3Call3{
// 				Target:   itoken,
// 				CallData: data,
// 			})
// 		}
// 		// Controller calls. Fetch markets.
// 		marketsData, err := controllerABI.Pack("markets", itoken)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to pack markets calldata: %v", err)
// 		}
// 		calls = append(calls, txs.Multicall3Call3{
// 			Target:   d.controllerAddress,
// 			CallData: marketsData,
// 		})
// 		// Price Oracle calls. Get underlying price.
// 		priceData, err := oracleABI.Pack("getUnderlyingPrice", itoken)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to pack getUnderlyingPrice calldata: %v", err)
// 		}
// 		calls = append(calls, txs.Multicall3Call3{
// 			Target:   d.oracleAddress,
// 			CallData: priceData,
// 		})
// 	}

// 	// Perform multicall
// 	responses, err := txs.HandleMulticall(d.cl, &calls)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to multicall: %v", err)
// 	}

// 	// Unpack
// 	numTokens := len(allITokens) - len(ignoreTokens) // Exclude ignore token
// 	supplyMarkets := make([]*t.MarketInfo, numTokens)
// 	borrowMarkets := make([]*t.MarketInfo, numTokens)
// 	interestRateModels := make([]common.Address, numTokens)
// 	factor := len(iTokenMethods) + 2
// 	for i := 0; i < numTokens; i++ {
// 		j := i * factor
// 		// Unpack itoken calls
// 		var symbol string
// 		var decimals uint8
// 		var supplyRatePerBlock *big.Int
// 		var borrowRatePerBlock *big.Int
// 		var totalSupply *big.Int
// 		var getCash *big.Int
// 		var interestRateModel common.Address
// 		var totalBorrows *big.Int
// 		var totalReserves *big.Int
// 		var reserveRatio *big.Int
// 		results := []interface{}{&symbol, &decimals, &supplyRatePerBlock, &borrowRatePerBlock, &totalSupply, &getCash, &interestRateModel, &totalBorrows, &totalReserves, &reserveRatio}
// 		for k, method := range iTokenMethods {
// 			if err := iTokenABI.UnpackIntoInterface(results[k], method, (*responses)[j+k].ReturnData); err != nil {
// 				return nil, fmt.Errorf("failed to unpack %v: %v", method, err)
// 			}
// 		}
// 		interestRateModels[i] = interestRateModel
// 		// Unpack controller calls
// 		var controllerMarket struct {
// 			CollateralFactorMantissa *big.Int
// 			BorrowFactorMantissa     *big.Int
// 			BorrowCapacity           *big.Int
// 			SupplyCapacity           *big.Int
// 			MintPaused               bool
// 			RedeemPaused             bool
// 			BorrowPaused             bool
// 		}
// 		if err := controllerABI.UnpackIntoInterface(&controllerMarket, "markets", (*responses)[j+factor-2].ReturnData); err != nil {
// 			return nil, fmt.Errorf("failed to unpack markets: %v", err)
// 		}
// 		// Unpack oracle calls
// 		var underlyingPrice *big.Int
// 		if err := oracleABI.UnpackIntoInterface(&underlyingPrice, "getUnderlyingPrice", (*responses)[j+factor-1].ReturnData); err != nil {
// 			return nil, fmt.Errorf("failed to unpack getUnderlyingPrice: %v", err)
// 		}

// 		// Fill out market info partially
// 		decimalsInt := big.NewInt(int64(decimals))
// 		// LTV
// 		ltv := new(big.Int).Mul(controllerMarket.CollateralFactorMantissa, big.NewInt(1e4))
// 		ltv.Quo(ltv, utils.ETHMantissaInt)
// 		// Caps
// 		decimalsFactor := big.NewFloat(math.Pow(10, float64(decimals)))
// 		supplyCap := new(big.Int).Sub(controllerMarket.SupplyCapacity, totalSupply)
// 		borrowCap := new(big.Float).SetInt(getCash)
// 		borrowCap.Quo(borrowCap, decimalsFactor)
// 		// Convert compv2 prices to 8 decimals
// 		inverseDecimals := new(big.Int).Sub(big.NewInt(36), decimalsInt)
// 		price := new(big.Int).Mul(underlyingPrice, big.NewInt(1e8))
// 		price.Quo(price, new(big.Int).Exp(big.NewInt(10), inverseDecimals, nil))
// 		market := &t.MarketInfo{
// 			Protocol:   DForceName,
// 			Chain:      d.chain,
// 			Token:      symbol,
// 			Decimals:   decimalsInt,
// 			LTV:        ltv,
// 			PriceInUSD: price,
// 			Params: DForceParams{
// 				SupplyCapRemaining: supplyCap,
// 				TotalBorrows:       totalBorrows,
// 				TotalReserves:      totalReserves,
// 				TotalCash:          getCash,
// 				ReserveRatio:       reserveRatio,

// 				TotalSupply:   totalSupply,
// 				BlocksPerYear: blocksPerYear[d.chain],
// 			},
// 		}
// 		supplyMarkets[i] = market
// 		borrowMarkets[i] = market
// 	}

// 	// Fetch interest rate model params
// 	interestRateConstants, err := d.getInterestRateConstants(interestRateModels)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get interest rate constants: %v", err)
// 	}

// 	// Fill out rest of market info
// 	for i, market := range supplyMarkets {
// 		params := market.Params.(DForceParams)
// 		params.Optimal = interestRateConstants[i].Optimal
// 		params.Base = interestRateConstants[i].Base
// 		params.Slope1 = interestRateConstants[i].Slope1
// 		params.Slope2 = interestRateConstants[i].Slope2
// 		market.Params = params
// 	}

// 	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
// 	log.Printf("Time elapsed: %v", time.Since(startTime))
// 	return []*t.ProtocolChain{{
// 		Protocol:      DForceName,
// 		Chain:         d.chain,
// 		SupplyMarkets: supplyMarkets,
// 		BorrowMarkets: borrowMarkets,
// 	}}, nil
// }

// // A fork of CompoundV2 with different base (instead of Mantissa)
// func (*DForce) CalcAPY(market *t.MarketInfo, amount *big.Int, isSupply bool) (*big.Int, *big.Int, error) {
// 	params, ok := market.Params.(DForceParams)
// 	if !ok {
// 		return nil, nil, fmt.Errorf("failed to cast params to AaveV3MarketParams")
// 	}

// 	// Check for caps
// 	actualAmount := amount
// 	availableLiquidity := new(big.Int).Sub(params.TotalCash, params.TotalReserves)
// 	if isSupply && params.SupplyCapRemaining.Cmp(amount) == -1 {
// 		actualAmount = params.SupplyCapRemaining
// 	} else if !isSupply && availableLiquidity.Cmp(amount) == -1 {
// 		actualAmount = availableLiquidity
// 	}

// 	var rate *big.Int
// 	if isSupply {
// 		params.TotalCash.Add(params.TotalCash, actualAmount)
// 		rate = getSupplyRate(&params)
// 	} else {
// 		params.TotalBorrows.Add(params.TotalBorrows, actualAmount)
// 		params.TotalCash.Sub(params.TotalCash, actualAmount)
// 		rate = getBorrowRate(&params)
// 	}
// 	return utils.ConvertRatePerBlockToAPY(rate, params.BlocksPerYear), actualAmount, nil
// }

// // TODO: Add checks for 100% utilization rate
// // (https://github.com/dforce-network/LendingContractsV2/blob/master/contracts/InterestRateModel/InterestRateModel.sol#L87)
// func utilizationRate(params *DForceParams) *big.Int {
// 	if params.TotalBorrows.Cmp(big.NewInt(0)) == 0 {
// 		return big.NewInt(0)
// 	}
// 	borrowsMantissa := new(big.Int).Mul(params.TotalBorrows, utils.ETHMantissaInt)
// 	supply := new(big.Int).Add(params.TotalCash, params.TotalBorrows)
// 	supply.Sub(supply, params.TotalReserves)
// 	return new(big.Int).Div(borrowsMantissa, supply)
// }

// func getBorrowRate(params *DForceParams) *big.Int {
// 	util := utilizationRate(params)

// 	var rate *big.Int
// 	if util.Cmp(params.Optimal) != 1 {
// 		rate = new(big.Int).Mul(util, params.Slope1)
// 		rate.Div(rate, params.Optimal)
// 		rate.Add(rate, params.Base)
// 	} else {
// 		baseRates := new(big.Int).Add(params.Base, params.Slope1)

// 		excessUtil := new(big.Int).Sub(util, params.Optimal)
// 		maxExcellUtil := new(big.Int).Sub(utils.ETHMantissaInt, params.Optimal)
// 		excessRate := new(big.Int).Mul(excessUtil, params.Slope2)

// 		rate = new(big.Int).Add(baseRates, excessRate)
// 		rate.Div(rate, maxExcellUtil)
// 	}

// 	// Divide by blocks per year
// 	return rate.Div(rate, params.BlocksPerYear)
// }

// // supplyRate = borrowRate × (1-reserveFactor) × borrowsPer
// func getSupplyRate(params *DForceParams) *big.Int {
// 	borrowRate := getBorrowRate(params)
// 	oneMinusReserveRatio := new(big.Int).Sub(utils.ETHMantissaInt, params.ReserveRatio)
// 	util := utilizationRate(params)
// 	rate := new(big.Int).Mul(borrowRate, oneMinusReserveRatio)
// 	rate.Mul(rate, util)
// 	return rate.Div(rate, new(big.Int).Exp(utils.ETHMantissaInt, big.NewInt(2), nil))
// }

// // Instantiate iToken contract.
// func (d *DForce) InstantiateIToken(token string) (*IToken, common.Address, error) {
// 	if token[0] != 'i' {
// 		return nil, common.Address{}, fmt.Errorf("token must be an iToken")
// 	}

// 	itoken, err := utils.ConvertSymbolToAddress(d.chain, token)
// 	if err != nil {
// 		return nil, common.Address{}, fmt.Errorf("failed to convert symbol to address: %v", err)
// 	}
// 	iTokenAddress := common.HexToAddress(itoken)

// 	iTokenContract, err := NewIToken(iTokenAddress, d.cl)
// 	if err != nil {
// 		return nil, common.Address{}, fmt.Errorf("failed to instantiate iToken: %v", err)
// 	}

// 	return iTokenContract, iTokenAddress, nil
// }

// // Lends the token to the protocol
// func (d *DForce) Supply(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
// 	iTokenContract, iTokenAddress, err := d.InstantiateIToken(token)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
// 	}

// 	walletAddress := common.HexToAddress(wallet)
// 	auth, err := accounts.GetAuth(d.cl, d.chainID, walletAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
// 	}

// 	// Check if entered market
// 	enteredMarket, err := d.controllerContract.HasEnteredMarket(nil, walletAddress, iTokenAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get entered markets: %v", err)
// 	}
// 	if !enteredMarket {
// 		// Enter market
// 		log.Printf("Entering %v market on %v", token, DForceName)
// 		tx, err := d.controllerContract.EnterMarkets(auth, []common.Address{iTokenAddress})
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to enter market: %v", err)
// 		}
// 		// Wait for tx to be mined
// 		_, err = txs.WaitForConfirmations(d.cl, tx, 0)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
// 		}
// 		log.Printf("Entered %v market on %v (%v)", token, DForceName, tx.Hash())
// 	}

// 	var tx *types.Transaction
// 	if token == nativeToken {
// 		// Send native token
// 		// Manually added iETH mint() function to abi before generating bindings
// 		auth.Value = amount
// 		tx, err = iTokenContract.Mint0(auth, walletAddress)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to mint: %v", err)
// 		}
// 	} else {
// 		// Handle approvals
// 		tokenAddress, err := iTokenContract.Underlying(nil)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to get underlying token address: %v", err)
// 		}
// 		_, err = transactions.ApproveERC20IfNeeded(d.cl, auth, tokenAddress, walletAddress, iTokenAddress, amount)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to approve: %v", err)
// 		}

// 		tx, err = iTokenContract.Mint(auth, walletAddress, amount)
// 		if err != nil {
// 			return nil, fmt.Errorf("failed to mint: %v", err)
// 		}
// 	}

// 	// Wait
// 	_, err = txs.WaitForConfirmations(d.cl, tx, 0)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
// 	}

// 	log.Printf("Supplied %v %v to %v on %v (%v)", amount, token, DForceName, d.chain, tx.Hash())
// 	return tx, nil
// }

// // Withdraws the token from the protocol.
// func (d *DForce) Withdraw(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
// 	iTokenContract, iTokenAddress, err := d.InstantiateIToken(token)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
// 	}

// 	walletAddress := common.HexToAddress(wallet)
// 	auth, err := accounts.GetAuth(d.cl, d.chainID, walletAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
// 	}
// 	// Handle approvals
// 	_, err = transactions.ApproveERC20IfNeeded(d.cl, auth, iTokenAddress, walletAddress, iTokenAddress, amount)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to approve: %v", err)
// 	}

// 	tx, err := iTokenContract.RedeemUnderlying(auth, walletAddress, amount)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to withdraw: %v", err)
// 	}

// 	// Wait
// 	_, err = txs.WaitForConfirmations(d.cl, tx, 0)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
// 	}

// 	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, DForceName, d.chain, tx.Hash())
// 	return tx, nil
// }

// // Withdraws all of the token from the protocol.
// func (d *DForce) WithdrawAll(wallet string, token string) (*types.Transaction, error) {
// 	iTokenContract, iTokenAddress, err := d.InstantiateIToken(token)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
// 	}
// 	walletAddress := common.HexToAddress(wallet)

// 	// Get iToken balance
// 	amount, err := iTokenContract.BalanceOf(nil, walletAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get balance: %v", err)
// 	}

// 	auth, err := accounts.GetAuth(d.cl, d.chainID, walletAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
// 	}

// 	// Handle approvals
// 	_, err = transactions.ApproveERC20IfNeeded(d.cl, auth, iTokenAddress, walletAddress, iTokenAddress, amount)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to approve: %v", err)
// 	}

// 	tx, err := iTokenContract.Redeem(auth, walletAddress, amount)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to withdraw: %v", err)
// 	}

// 	// Wait
// 	_, err = txs.WaitForConfirmations(d.cl, tx, 0)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
// 	}

// 	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, DForceName, d.chain, tx.Hash())
// 	return tx, nil
// }

// // Borrows the token from the protocol
// func (d *DForce) Borrow(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
// 	iTokenContract, _, err := d.InstantiateIToken(token)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
// 	}

// 	walletAddress := common.HexToAddress(wallet)
// 	auth, err := accounts.GetAuth(d.cl, d.chainID, walletAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
// 	}

// 	tx, err := iTokenContract.Borrow(auth, amount)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to borrow: %v", err)
// 	}

// 	// Wait
// 	_, err = txs.WaitForConfirmations(d.cl, tx, 0)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
// 	}

// 	log.Printf("Borrowed %v %v from %v on %v (%v)", amount, token, DForceName, d.chain, tx.Hash())
// 	return tx, nil
// }

// // Repays the token to the protocol.
// func (d *DForce) Repay(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
// 	iTokenContract, iTokenAddress, err := d.InstantiateIToken(token)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
// 	}

// 	walletAddress := common.HexToAddress(wallet)
// 	auth, err := accounts.GetAuth(d.cl, d.chainID, walletAddress)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
// 	}

// 	tokenAddress, err := iTokenContract.Underlying(nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to get underlying token address: %v", err)
// 	}
// 	// Handle approvals
// 	_, err = transactions.ApproveERC20IfNeeded(d.cl, auth, tokenAddress, walletAddress, iTokenAddress, amount)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to approve: %v", err)
// 	}

// 	tx, err := iTokenContract.RepayBorrow(auth, amount)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to repay: %v", err)
// 	}

// 	// Wait
// 	_, err = txs.WaitForConfirmations(d.cl, tx, 0)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
// 	}

// 	log.Printf("Repaid %v %v to %v on %v (%v)", amount, token, DForceName, d.chain, tx.Hash())
// 	return tx, nil
// }

// // Repays all of the token to the protocol.
// func (d *DForce) RepayAll(wallet string, token string) (*types.Transaction, error) {
// 	return d.Repay(wallet, token, utils.MaxUint256)
// }
