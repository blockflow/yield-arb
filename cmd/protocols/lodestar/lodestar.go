package lodestar

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"
	"time"
	"yield-arb/cmd/accounts"
	"yield-arb/cmd/utils"

	t "yield-arb/cmd/protocols/types"
	txs "yield-arb/cmd/transactions"

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

const LodestarName = "lodestar"

var comptrollerAddresses = map[string]string{
	"arbitrum": "0xa86DD95c210dd186Fa7639F93E4177E97d057576",
}
var lensAddresses = map[string]string{
	"arbitrum": "0x24C25910aF4068B5F6C3b75252a36c4810849135",
}
var lusdcAddress = common.HexToAddress("0x1ca530f02DD0487cef4943c674342c5aEa08922F")

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
	prices := make([]*big.Float, numAssets)
	usdPrice := big.NewFloat(0)
	for i, metadata := range metadataAll {
		decimals := uint8(metadata.UnderlyingDecimals.Uint64())
		price := new(big.Float).SetInt(oraclePrices[i].UnderlyingPrice)
		price.Quo(price, big.NewFloat(math.Pow(10, 36-float64(decimals))))
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
	calls := make([]txs.Multicall3Call3, numAssets)
	comptrollerABI, err := ComptrollerMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comptroller abi: %v", err)
	}
	for i, market := range markets {
		data, err := comptrollerABI.Pack("supplyCaps", market)
		if err != nil {
			return nil, fmt.Errorf("failed to pack supply caps: %v", err)
		}
		calls[i] = txs.Multicall3Call3{
			Target:   l.comptrollerAddress,
			CallData: data,
		}
	}
	// Perform multicall
	responses, err := txs.HandleMulticall(l.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall: %v", err)
	}
	// Unpack responses
	supplyCaps := make([]*big.Float, numAssets)
	for i, response := range *responses {
		var supplyCap *big.Int
		if err := comptrollerABI.UnpackIntoInterface(&supplyCap, "supplyCaps", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", "supplyCaps", err)
		}
		supplyCaps[i] = new(big.Float).SetInt(supplyCap)
	}

	supplyMarkets := make([]*t.MarketInfo, numAssets)
	borrowMarkets := make([]*t.MarketInfo, numAssets)
	for i, metadata := range metadataAll {
		decimals := uint8(metadata.UnderlyingDecimals.Uint64())
		ltv := new(big.Float).Quo(new(big.Float).SetInt(metadata.CollateralFactorMantissa), utils.ETHMantissa)
		ltv.Mul(ltv, big.NewFloat(100))
		supplyAPY := utils.ConvertRatePerBlockToAPY(metadata.SupplyRatePerBlock)
		borrowAPY := utils.ConvertRatePerBlockToAPY(metadata.BorrowRatePerBlock)
		decimalsFactor := big.NewFloat(math.Pow(10, float64(decimals)))
		supplyCap := new(big.Float).Quo(supplyCaps[i], decimalsFactor)
		borrowCap := new(big.Float).Quo(new(big.Float).SetInt(metadata.TotalCash), decimalsFactor)
		market := &t.MarketInfo{
			Protocol:   LodestarName,
			Chain:      l.chain,
			Token:      symbols[i],
			Decimals:   decimals,
			LTV:        ltv,
			SupplyAPY:  supplyAPY,
			BorrowAPY:  borrowAPY,
			SupplyCap:  supplyCap,
			BorrowCap:  borrowCap,
			PriceInUSD: prices[i],
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

// Instantiate lToken contract.
func (l *Lodestar) InstantiateLToken(token string) (*LToken, error) {
	if token[0] != 'l' {
		return nil, fmt.Errorf("token must be an lToken")
	}

	iTokenAddress, err := utils.ConvertSymbolToAddress(l.chain, token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}

	iTokenContract, err := NewLToken(common.HexToAddress(iTokenAddress), l.cl)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	return iTokenContract, nil
}

// Deposits the specified token into the protocol
// TODO: Need to call EnterMarkets() to collateralize the supplied asset(s).
func (l *Lodestar) Supply(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	lTokenContract, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(l.cl, l.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	tx, err := lTokenContract.Mint(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to mint: %v", err)
	}

	log.Printf("Supplied %v %v to %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Withdraws the specified token from the protocol.
func (l *Lodestar) Withdraw(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	lTokenContract, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(l.cl, l.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	tx, err := lTokenContract.Redeem(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to redeem: %v", err)
	}

	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Withdraws all of the specified token from the protocol.
func (l *Lodestar) WithdrawAll(wallet string, token string) (*types.Transaction, error) {
	lTokenContract, err := l.InstantiateLToken(token)
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
	lTokenContract, err := l.InstantiateLToken(token)
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

	log.Printf("Borrowed %v %v from %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Repays the specified token to the protocol.
func (l *Lodestar) Repay(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	lTokenContract, err := l.InstantiateLToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate iToken: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(l.cl, l.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

	tx, err := lTokenContract.RepayBorrow(auth, amount)
	if err != nil {
		return nil, fmt.Errorf("failed to repay: %v", err)
	}

	log.Printf("Repaid %v %v to %v on %v (%v)", amount, token, LodestarName, l.chain, tx.Hash())
	return tx, nil
}

// Repays all of the specified token to the protocol.
func (l *Lodestar) RepayAll(wallet string, token string) (*types.Transaction, error) {
	lTokenContract, err := l.InstantiateLToken(token)
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
