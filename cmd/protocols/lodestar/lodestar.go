package lodestar

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"
	"yield-arb/cmd/accounts"
	"yield-arb/cmd/transactions"
	"yield-arb/cmd/utils"

	t "yield-arb/cmd/protocols/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Lodestar struct {
	chain               string
	cl                  *ethclient.Client
	chainID             *big.Int
	comptrollerAddress  common.Address
	comptrollerContract *Comptroller
	lensAddress         common.Address
	lensContract        *Lens
}

type LodestarParams struct {
	SupplyCapRemaining    *big.Int
	TotalBorrows          *big.Int
	TotalReserves         *big.Int
	TotalCash             *big.Int
	ReserveFactorMantissa *big.Int
	// Constants
	Kink                   *big.Int
	BaseRatePerBlock       *big.Int
	MultiplierPerBlock     *big.Int
	JumpMultiplierPerBlock *big.Int
	BlocksPerYear          *big.Int
}

type InterestRateConstants struct {
	Kink                   *big.Int
	BaseRatePerBlock       *big.Int
	MultiplierPerBlock     *big.Int
	JumpMultiplierPerBlock *big.Int
}

const LodestarName = "lodestar"

var comptrollerAddresses = map[string]string{
	"arbitrum": "0xa86DD95c210dd186Fa7639F93E4177E97d057576",
}
var lensAddresses = map[string]string{
	"arbitrum": "0x24C25910aF4068B5F6C3b75252a36c4810849135",
}
var lusdcAddress = common.HexToAddress("0x1ca530f02DD0487cef4943c674342c5aEa08922F")
var blocksPerYear = big.NewInt(2628000)

func (l *Lodestar) GetChains() ([]string, error) {
	return []string{"arbitrum"}, nil
}

func NewLodestarProtocol() *Lodestar {
	return &Lodestar{}
}

func (l *Lodestar) Connect(chain string) error {
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

	// Instantiate comptroller
	comptrollerAddress := common.HexToAddress(comptrollerAddresses[chain])
	comptrollerContract, err := NewComptroller(comptrollerAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate comptroller: %v", err)
		return err
	}

	// Instantiate lens
	lensAddress := common.HexToAddress(lensAddresses[chain])
	lensContract, err := NewLens(lensAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate lens: %v", err)
		return err
	}

	l.chain = chain
	l.cl = cl
	l.chainID = chainid
	l.comptrollerAddress = comptrollerAddress
	l.comptrollerContract = comptrollerContract
	l.lensAddress = lensAddress
	l.lensContract = lensContract
	log.Printf("%v connected to %v (chainid: %v, lens: %v)", LodestarName, l.chain, l.chainID, lensAddress)
	return nil
}

func (l *Lodestar) getInterestRateConstants(markets []common.Address) ([]*InterestRateConstants, error) {
	// Fetch interest rate models
	// Pack calldata
	calls := make([]transactions.Multicall3Call3, len(markets))
	ltokenABI, err := LTokenMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get ltoken abi: %v", err)
	}
	for i, market := range markets {
		data, err := ltokenABI.Pack("interestRateModel")
		if err != nil {
			return nil, fmt.Errorf("failed to pack interest rate model: %v", err)
		}
		calls[i] = transactions.Multicall3Call3{
			Target:   market,
			CallData: data,
		}
	}
	// Perform multicall
	responses, err := transactions.HandleMulticall(l.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall: %v", err)
	}
	// Unpack responses
	interestRateModels := make([]*common.Address, len(markets))
	for i, response := range *responses {
		interestRateModels[i] = new(common.Address) // Allocate memory
		if err := ltokenABI.UnpackIntoInterface(interestRateModels[i], "interestRateModel", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", "interestRateModel", err)
		}
	}

	// Fetch interest rate constants
	// Pack calldata
	calls = make([]transactions.Multicall3Call3, len(interestRateModels)*4)
	interestRateModelABI, err := InterestRateModelMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get interest rate model abi: %v", err)
	}
	for i, interestRateModel := range interestRateModels {
		kinkData, err := interestRateModelABI.Pack("kink")
		if err != nil {
			return nil, fmt.Errorf("failed to pack kink: %v", err)
		}
		baseRatePerBlockData, err := interestRateModelABI.Pack("baseRatePerBlock")
		if err != nil {
			return nil, fmt.Errorf("failed to pack baseRatePerBlock: %v", err)
		}
		multiplierPerBlockData, err := interestRateModelABI.Pack("multiplierPerBlock")
		if err != nil {
			return nil, fmt.Errorf("failed to pack multiplierPerBlock: %v", err)
		}
		jumpMultiplierPerBlockData, err := interestRateModelABI.Pack("jumpMultiplierPerBlock")
		if err != nil {
			return nil, fmt.Errorf("failed to pack jumpMultiplierPerBlock: %v", err)
		}
		calls[i*4] = transactions.Multicall3Call3{
			Target:   *interestRateModel,
			CallData: kinkData,
		}
		calls[i*4+1] = transactions.Multicall3Call3{
			Target:   *interestRateModel,
			CallData: baseRatePerBlockData,
		}
		calls[i*4+2] = transactions.Multicall3Call3{
			Target:   *interestRateModel,
			CallData: multiplierPerBlockData,
		}
		calls[i*4+3] = transactions.Multicall3Call3{
			Target:   *interestRateModel,
			CallData: jumpMultiplierPerBlockData,
		}
	}
	// Perform multicall
	responses, err = transactions.HandleMulticall(l.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall: %v", err)
	}
	// Unpack responses
	type ResponseOutput struct {
		Data *big.Int
	}
	interestRateConstants := make([]*InterestRateConstants, len(interestRateModels))
	for i := range interestRateConstants {
		rateConstants := new(InterestRateConstants)
		output := new(ResponseOutput)
		if err := interestRateModelABI.UnpackIntoInterface(output, "kink", (*responses)[i*4].ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", "kink", err)
		}
		rateConstants.Kink = output.Data
		output = new(ResponseOutput)
		if err := interestRateModelABI.UnpackIntoInterface(output, "baseRatePerBlock", (*responses)[i*4+1].ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", "baseRatePerBlock", err)
		}
		rateConstants.BaseRatePerBlock = output.Data
		output = new(ResponseOutput)
		if err := interestRateModelABI.UnpackIntoInterface(output, "multiplierPerBlock", (*responses)[i*4+2].ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", "multiplierPerBlock", err)
		}
		rateConstants.MultiplierPerBlock = output.Data
		output = new(ResponseOutput)
		if err := interestRateModelABI.UnpackIntoInterface(output, "jumpMultiplierPerBlock", (*responses)[i*4+3].ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", "jumpMultiplierPerBlock", err)
		}
		rateConstants.JumpMultiplierPerBlock = output.Data
		interestRateConstants[i] = rateConstants
	}

	return interestRateConstants, nil
}

// Returns the market.
// Assumes lending and borrowing tokens are the same.
func (l *Lodestar) GetMarkets() (*t.ProtocolChain, error) {
	log.Printf("Fetching market data for %v...", l.chain)
	startTime := time.Now()

	// Fetch all markets
	markets, err := l.comptrollerContract.GetAllMarkets(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get all markets: %v", err)
	}
	numAssets := len(markets)
	addresses := make([]string, numAssets)
	for i, market := range markets {
		addresses[i] = market.Hex()
	}
	symbols, err := utils.ConvertAddressesToSymbols(l.chain, addresses)
	if err != nil {
		return nil, fmt.Errorf("failed to convert addresses: %v", err)
	}

	// Fetch all market metadata, cTokenMetadataAll was manually changed to view in json
	metadataAll, err := l.lensContract.CTokenMetadataAll(nil, markets)
	if err != nil {
		return nil, fmt.Errorf("failed to get all token metadata: %v", err)
	}
	oraclePrices, err := l.lensContract.CTokenUnderlyingPriceAll(nil, markets)
	if err != nil {
		return nil, fmt.Errorf("failed to get all prices: %v", err)
	}

	// Convert compv2 prices to USD
	prices := make([]*big.Int, numAssets)
	usdPrice := big.NewInt(0)
	for i, metadata := range metadataAll {
		inverseDecimals := new(big.Int).Sub(big.NewInt(36), metadata.UnderlyingDecimals)
		price := new(big.Int).Mul(oraclePrices[i].UnderlyingPrice, big.NewInt(10000))
		price.Quo(price, new(big.Int).Exp(big.NewInt(10), inverseDecimals, nil))
		if metadata.CToken == lusdcAddress {
			*usdPrice = *price
		}
		prices[i] = price
	}
	for _, price := range prices {
		price.Quo(price, usdPrice)
	}

	// Fetch all supply caps
	// Aggregate calldata
	calls := make([]transactions.Multicall3Call3, numAssets)
	comptrollerABI, err := ComptrollerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comptroller abi: %v", err)
	}
	for i, market := range markets {
		data, err := comptrollerABI.Pack("supplyCaps", market)
		if err != nil {
			return nil, fmt.Errorf("failed to pack supply caps: %v", err)
		}
		calls[i] = transactions.Multicall3Call3{
			Target:   l.comptrollerAddress,
			CallData: data,
		}
	}
	// Perform multicall
	responses, err := transactions.HandleMulticall(l.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall: %v", err)
	}
	// Unpack responses
	// TODO: reformat like interestrateconstants
	supplyCaps := make([]*big.Int, numAssets)
	for i, response := range *responses {
		var supplyCap *big.Int
		if err := comptrollerABI.UnpackIntoInterface(&supplyCap, "supplyCaps", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", "supplyCaps", err)
		}
		supplyCaps[i] = supplyCap
	}

	// Fetch interest rate constants
	interestRateConstants, err := l.getInterestRateConstants(markets)
	if err != nil {
		return nil, fmt.Errorf("failed to get interest rate constants: %v", err)
	}

	supplyMarkets := make([]*t.MarketInfo, numAssets)
	borrowMarkets := make([]*t.MarketInfo, numAssets)
	for i, metadata := range metadataAll {
		ltv := new(big.Int).Mul(metadata.CollateralFactorMantissa, big.NewInt(10000))
		ltv.Quo(ltv, utils.ETHMantissaInt)
		supplyCap := new(big.Int).Sub(supplyCaps[i], metadata.TotalCash)
		supplyCap.Sub(supplyCap, metadata.TotalBorrows)
		supplyCap.Add(supplyCap, metadata.TotalReserves)
		market := &t.MarketInfo{
			Protocol:   LodestarName,
			Chain:      l.chain,
			Token:      symbols[i],
			Decimals:   metadata.UnderlyingDecimals,
			LTV:        ltv,
			PriceInUSD: prices[i],
			Params: LodestarParams{
				SupplyCapRemaining:    supplyCap,
				TotalBorrows:          metadata.TotalBorrows,
				TotalReserves:         metadata.TotalReserves,
				TotalCash:             metadata.TotalCash,
				ReserveFactorMantissa: metadata.ReserveFactorMantissa,

				Kink:                   interestRateConstants[i].Kink,
				BaseRatePerBlock:       interestRateConstants[i].BaseRatePerBlock,
				MultiplierPerBlock:     interestRateConstants[i].MultiplierPerBlock,
				JumpMultiplierPerBlock: interestRateConstants[i].JumpMultiplierPerBlock,
				BlocksPerYear:          blocksPerYear,
			},
		}
		supplyMarkets[i] = market
		borrowMarkets[i] = market
	}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return &t.ProtocolChain{
		Protocol:      LodestarName,
		Chain:         l.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}, nil
}

func (*Lodestar) CalcAPY(market *t.MarketInfo, amount *big.Int, isSupply bool) (*big.Int, *big.Int, error) {
	params, ok := market.Params.(LodestarParams)
	if !ok {
		return nil, nil, fmt.Errorf("failed to cast params to AaveV3MarketParams")
	}

	// Check for caps
	actualAmount := amount
	availableLiquidity := new(big.Int).Sub(params.TotalCash, params.TotalReserves)
	if isSupply && params.SupplyCapRemaining.Cmp(amount) == -1 {
		actualAmount = params.SupplyCapRemaining
	} else if !isSupply && availableLiquidity.Cmp(amount) == -1 {
		actualAmount = availableLiquidity
	}

	var rate *big.Int
	if isSupply {
		params.TotalCash.Add(params.TotalCash, actualAmount)
		rate = getSupplyRate(&params)
	} else {
		params.TotalBorrows.Add(params.TotalBorrows, actualAmount)
		params.TotalCash.Sub(params.TotalCash, actualAmount)
		rate = getBorrowRate(&params)
	}
	return utils.ConvertRatePerBlockToAPY(rate, params.BlocksPerYear), actualAmount, nil
}

func utilizationRate(params *LodestarParams) *big.Int {
	if params.TotalBorrows.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	util := new(big.Int).Mul(params.TotalBorrows, utils.ETHMantissaInt)
	denom := new(big.Int).Add(params.TotalCash, params.TotalBorrows)
	denom.Sub(denom, params.TotalReserves)
	return util.Div(util, denom)
}

func getBorrowRate(params *LodestarParams) *big.Int {
	util := utilizationRate(params)

	if util.Cmp(params.Kink) != 1 {
		rate := new(big.Int).Mul(util, params.MultiplierPerBlock)
		rate.Div(rate, utils.ETHMantissaInt)
		return rate.Add(rate, params.BaseRatePerBlock)
	} else {
		normalRate := new(big.Int).Mul(params.Kink, params.MultiplierPerBlock)
		normalRate.Div(normalRate, utils.ETHMantissaInt)
		normalRate.Add(normalRate, params.BaseRatePerBlock)

		excessUtil := new(big.Int).Sub(util, params.Kink)
		excessUtil.Mul(excessUtil, params.JumpMultiplierPerBlock)
		excessUtil.Div(excessUtil, utils.ETHMantissaInt)
		return normalRate.Add(normalRate, excessUtil)
	}
}

func getSupplyRate(params *LodestarParams) *big.Int {
	oneMinusReserveFactor := new(big.Int).Sub(utils.ETHMantissaInt, params.ReserveFactorMantissa)
	borrowRate := getBorrowRate(params)
	rateToPool := new(big.Int).Mul(borrowRate, oneMinusReserveFactor)
	rateToPool.Div(rateToPool, utils.ETHMantissaInt)
	rate := new(big.Int).Mul(utilizationRate(params), rateToPool)
	return rate.Div(rate, utils.ETHMantissaInt)
}

// Instantiate lToken contract.
func (l *Lodestar) InstantiateLToken(token string) (*LToken, common.Address, error) {
	if token[0] != 'l' {
		return nil, common.Address{}, fmt.Errorf("token must be an lToken")
	}

	ltoken, err := utils.ConvertSymbolToAddress(l.chain, token)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	lTokenAddress := common.HexToAddress(ltoken)

	iTokenContract, err := NewLToken(lTokenAddress, l.cl)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	return iTokenContract, lTokenAddress, nil
}

// Deposits the specified token into the protocol
// TODO: Need to call EnterMarkets() to collateralize the supplied asset(s).
func (l *Lodestar) Supply(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	lTokenContract, lTokenAddress, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(l.cl, l.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	// Check if entered market
	enteredMarket, err := l.comptrollerContract.CheckMembership(nil, walletAddress, lTokenAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to check membership: %v", err)
	}
	if !enteredMarket {
		tx, err := l.comptrollerContract.EnterMarkets(auth, []common.Address{lTokenAddress})
		if err != nil {
			return nil, fmt.Errorf("failed to enter markets: %v", err)
		}
		log.Printf("Entered %v market on %v (%v)", token, LodestarName, tx.Hash())
	}

	tokenAddress, err := lTokenContract.Underlying(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying token address: %v", err)
	}
	// Handle approvals
	_, err = transactions.ApproveERC20IfNeeded(l.cl, auth, tokenAddress, walletAddress, lTokenAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to approve: %v", err)
	}

	tx, err := lTokenContract.Mint(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to mint: %v", err)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(l.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Supplied %v %v to %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Withdraws the specified token from the protocol.
// Withdraws in units of ltoken, not underlying.
func (l *Lodestar) Withdraw(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	lTokenContract, lTokenAddress, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(l.cl, l.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}
	// Handle approvals
	_, err = transactions.ApproveERC20IfNeeded(l.cl, auth, lTokenAddress, walletAddress, lTokenAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to approve: %v", err)
	}

	tx, err := lTokenContract.Redeem(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to redeem: %v", err)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(l.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Withdraws all of the specified token from the protocol.
func (l *Lodestar) WithdrawAll(wallet string, token string) (*types.Transaction, error) {
	lTokenContract, _, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}
	walletAddress := common.HexToAddress(wallet)

	// Get ltoken balance
	amount, err := lTokenContract.BalanceOf(nil, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return l.Withdraw(wallet, token, amount)
}

// Borrows the specified token from the protocol.
// Defaults to variable interest rates.
func (l *Lodestar) Borrow(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	lTokenContract, _, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(l.cl, l.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	tx, err := lTokenContract.Borrow(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to borrow: %v", err)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(l.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Borrowed %v %v from %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Repays the specified token to the protocol.
func (l *Lodestar) Repay(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	lTokenContract, lTokenAddress, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(l.cl, l.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	tokenAddress, err := lTokenContract.Underlying(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying token address: %v", err)
	}
	// Handle approvals
	_, err = transactions.ApproveERC20IfNeeded(l.cl, auth, tokenAddress, walletAddress, lTokenAddress, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to approve: %v", err)
	}

	tx, err := lTokenContract.RepayBorrow(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to repay: %v", err)
	}

	// Wait
	_, err = transactions.WaitForConfirmations(l.cl, tx, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to wait for tx to be mined: %v", err)
	}

	log.Printf("Repaid %v %v to %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Repays all of the specified token to the protocol.
func (l *Lodestar) RepayAll(wallet string, token string) (*types.Transaction, error) {
	lTokenContract, _, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}
	walletAddress := common.HexToAddress(wallet)

	// Get ltoken balance
	amount, err := lTokenContract.BorrowBalanceStored(nil, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to get balance: %v", err)
	}

	return l.Repay(wallet, token, amount)
}
