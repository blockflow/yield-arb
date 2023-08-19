package aavev3

import (
	"context"
	"fmt"
	"log"
	"math"
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

const AaveV3Name = "aavev3"

var aavev3AddressesProviders = map[string]string{
	"ethereum":  "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e",
	"polygon":   "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"avalanche": "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"arbitrum":  "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"optimism":  "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
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

var wethGatewayAddresses = map[string]string{
	// Testnets
	"ethereum_goerli": "0x2a498323acad2971a8b1936fd7540596dc9bbacd",
}

var nativeTokens = map[string]string{
	"ethereum":        "ETH",
	"ethereum_goerli": "ETH",
	"arbitrum":        "ETH",
	"arbitrum_goerli": "AETH",
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
		log.Printf("Failed to instantiate Lending Pool: %v", err)
		return err
	}

	// Instantiate UIPoolDataProvider
	uiPoolDataProviderAddress := common.HexToAddress(uiPoolDataProviders[chain])
	uiPoolDataProviderCaller, err := NewAaveV3UIPoolDataProviderCaller(uiPoolDataProviderAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate UIPoolDataProvider: %v", err)
		return err
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

// Returns the market.
// Assumes lending and borrowing tokens are the same.
func (a *AaveV3) GetMarkets() (*t.ProtocolChain, error) {
	log.Printf("Fetching market data for %v...", a.chain)
	startTime := time.Now()

	// Fetch reserve data for all tokens
	aggReserveData, _, err := a.uiPoolDataProviderCaller.GetReservesData(nil, a.addressesProviderAddress)
	if err != nil {
		log.Printf("Failed to fetch reserve data: %v", err)
	}

	// Filter out results for specified symbols
	var supplyMarkets []*t.MarketInfo
	var borrowMarkets []*t.MarketInfo
	for _, reserveData := range aggReserveData {
		if reserveData.IsPaused {
			continue
		}

		decimals := new(big.Int).Exp(big.NewInt(10), reserveData.Decimals, nil)
		ltv := new(big.Float).SetInt(reserveData.BaseLTVasCollateral)
		ltv.Quo(ltv, big.NewFloat(100))
		lendingAPY := utils.ConvertRayToPercentage(reserveData.LiquidityRate)
		borrowingAPY := utils.ConvertRayToPercentage(reserveData.VariableBorrowRate)

		amountSupplied := new(big.Int).Add(reserveData.TotalScaledVariableDebt, reserveData.AvailableLiquidity)
		amountSupplied.Quo(amountSupplied, decimals)
		supplyCap := new(big.Float).SetInt(new(big.Int).Sub(reserveData.SupplyCap, amountSupplied))
		if supplyCap.Cmp(big.NewFloat(0)) == 0 { // Infinite cap
			supplyCap = big.NewFloat(math.MaxFloat64)
		}

		availableLiquidity := new(big.Int).Quo(reserveData.AvailableLiquidity, decimals)
		borrowCap := new(big.Float).SetInt(availableLiquidity)
		if borrowCap.Cmp(big.NewFloat(0)) == 0 { // Cap is liquidity
			borrowCap = new(big.Float).SetInt(reserveData.AvailableLiquidity)
		}

		priceInUSD := new(big.Float).SetInt(reserveData.PriceInMarketReferenceCurrency)
		priceInUSD.Quo(priceInUSD, big.NewFloat(100000000))
		market := &t.MarketInfo{
			Protocol:   AaveV3Name,
			Chain:      a.chain,
			Token:      reserveData.Symbol,
			Decimals:   uint8(reserveData.Decimals.Int64()),
			LTV:        ltv,
			SupplyAPY:  lendingAPY,
			BorrowAPY:  borrowingAPY,
			SupplyCap:  supplyCap,
			BorrowCap:  borrowCap,
			PriceInUSD: priceInUSD,
		}
		supplyMarkets = append(supplyMarkets, market)
		borrowMarkets = append(borrowMarkets, market)
	}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return &t.ProtocolChain{
		Protocol:      AaveV3Name,
		Chain:         a.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}, nil
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

// Deposits the specified token into the protocol
func (a *AaveV3) Supply(wallet string, token string, amount *big.Int) (*types.Transaction, error) {
	walletAddress := common.HexToAddress(wallet)
	auth, err := accounts.GetAuth(a.cl, a.chainID, walletAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve auth: %v", err)
	}

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
	log.Printf("Supplied %v %v to %v on %v (%v)", amount, token, AaveV3Name, a.chain, tx.Hash())
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
	log.Printf("Withdrew %v %v from %v on %v (%v)", amount, token, AaveV3Name, a.chain, tx.Hash())
	return tx, nil
}

func (a *AaveV3) WithdrawAll(wallet string, token string) (*types.Transaction, error) {
	maxUint := new(big.Int).SetUint64(math.MaxUint64)
	return a.Withdraw(wallet, token, maxUint)
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
	log.Printf("Repayed %v %v to %v on %v (%v)", amount, token, AaveV3Name, a.chain, tx.Hash())
	return tx, nil
}

func (a *AaveV3) RepayAll(wallet string, token string) (*types.Transaction, error) {
	maxUint := new(big.Int).SetUint64(math.MaxUint64)
	return a.Repay(wallet, token, maxUint)
}
