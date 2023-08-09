package aavev3

import (
	"context"
	"log"
	"math/big"
	"time"
	"yield-arb/cmd/accounts"
	txs "yield-arb/cmd/transactions"
	"yield-arb/cmd/utils"

	t "yield-arb/cmd/protocols/types"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/exp/slices"
)

// TODO: Account for new features such as isolation mode

type AaveV3 struct {
	chain                    string
	cl                       *ethclient.Client
	chainid                  *big.Int
	addressesProviderAddress common.Address
	poolAddress              string
	poolABI                  abi.ABI
	poolContract             *AaveV3Pool
	uiPoolDataProviderCaller *AaveV3UIPoolDataProviderCaller
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

	// Instantiate UIPoolDataProvider
	uiPoolDataProviderAddress := common.HexToAddress(uiPoolDataProviders[chain])
	uiPoolDataProviderCaller, err := NewAaveV3UIPoolDataProviderCaller(uiPoolDataProviderAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate UIPoolDataProvider: %v", err)
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

	a.chain = chain
	a.cl = cl
	a.chainid = chainid
	a.addressesProviderAddress = addressesProviderAddress
	a.poolAddress = lendingPoolAddress.Hex()
	a.poolContract = poolContract
	a.uiPoolDataProviderCaller = uiPoolDataProviderCaller
	log.Printf("%v connected to %v (chainid: %v, pool: %v)", AaveV3Name, a.chain, a.chainid, lendingPoolAddress)
	return nil
}

// Returns the symbols of the supported tokens
func (a *AaveV3) getReservesList() ([]string, error) {
	addresses, err := a.poolContract.AaveV3PoolCaller.GetReservesList(nil)
	if err != nil {
		log.Printf("Failed to fetch reserves list: %v", err)
		return nil, err
	}

	// Convert to string
	addressesString := make([]string, len(addresses))
	for i, address := range addresses {
		addressesString[i] = address.Hex()
	}

	// Convert to symbols
	symbols, err := utils.ConvertAddressesToSymbols(a.chain, addressesString)
	if err != nil {
		log.Printf("Failed to convert addresses to symbols: %v", err)
		return nil, err
	}

	return symbols, nil
}

// GetLendingTokens returns the tokens that can be lent on the given chain
func (a *AaveV3) GetLendingTokens() ([]string, error) {
	return a.getReservesList()
}

// GetBorrowingTokens returns the tokens that can be borrowed on the given chain
func (a *AaveV3) GetBorrowingTokens() ([]string, error) {
	return a.getReservesList()
}

// Get the lending TokenSpecs for the specified tokens. Filters out paused tokens.
func (a *AaveV3) getTokenSpecs(symbols []string, isLending bool) ([]*t.TokenSpecs, error) {
	addresses, err := utils.ConvertSymbolsToAddresses(a.chain, symbols)
	if err != nil {
		log.Printf("Failed to convert symbols to addresses: %v", err)
		return nil, err
	}
	if len(addresses) == 0 {
		return []*t.TokenSpecs{}, nil
	}

	// Fetch reserve data for all tokens
	aggReserveData, _, err := a.uiPoolDataProviderCaller.GetReservesData(nil, a.addressesProviderAddress)
	if err != nil {
		log.Printf("Failed to fetch reserve data: %v", err)
	}

	// Filter out results for specified symbols
	var tokenSpecs []*t.TokenSpecs
	for _, reserveData := range aggReserveData {
		if reserveData.IsPaused {
			continue
		}

		if slices.Contains[string](symbols, reserveData.Symbol) {
			ltv := new(big.Float).SetInt(reserveData.BaseLTVasCollateral)
			ltv.Quo(ltv, big.NewFloat(100))
			var apy *big.Float
			if isLending {
				apy = utils.ConvertRayToPercentage(reserveData.LiquidityRate)
			} else {
				apy = utils.ConvertRayToPercentage(reserveData.VariableBorrowRate)
			}
			tokenSpecs = append(tokenSpecs, &t.TokenSpecs{
				Protocol: AaveV3Name,
				Chain:    a.chain,
				Token:    reserveData.Symbol,
				LTV:      ltv,
				APY:      apy,
			})
		}
	}
	return tokenSpecs, nil
}

func (a *AaveV3) GetLendingSpecs(symbols []string) ([]*t.TokenSpecs, error) {
	return a.getTokenSpecs(symbols, true)
}

func (a *AaveV3) GetBorrowingSpecs(symbols []string) ([]*t.TokenSpecs, error) {
	return a.getTokenSpecs(symbols, false)
}

// Returns the market.
// Assumes lending and borrowing tokens are the same.
// TODO: Check for supply capped markets
func (a *AaveV3) GetMarkets() (*t.ProtocolMarkets, error) {
	log.Printf("Fetching market data for %v...", a.chain)
	startTime := time.Now()

	// Fetch reserve data for all tokens
	aggReserveData, _, err := a.uiPoolDataProviderCaller.GetReservesData(nil, a.addressesProviderAddress)
	if err != nil {
		log.Printf("Failed to fetch reserve data: %v", err)
	}

	// Filter out results for specified symbols
	var lendingSpecs []*t.TokenSpecs
	var borrowingSpecs []*t.TokenSpecs
	for _, reserveData := range aggReserveData {
		if reserveData.IsPaused {
			continue
		}

		ltv := new(big.Float).SetInt(reserveData.BaseLTVasCollateral)
		ltv.Quo(ltv, big.NewFloat(100))
		lendingAPY := utils.ConvertRayToPercentage(reserveData.LiquidityRate)
		borrowingAPY := utils.ConvertRayToPercentage(reserveData.VariableBorrowRate)
		lendingSpecs = append(lendingSpecs, &t.TokenSpecs{
			Protocol: AaveV3Name,
			Chain:    a.chain,
			Token:    reserveData.Symbol,
			LTV:      ltv,
			APY:      lendingAPY,
		})
		borrowingSpecs = append(borrowingSpecs, &t.TokenSpecs{
			Protocol: AaveV3Name,
			Chain:    a.chain,
			Token:    reserveData.Symbol,
			LTV:      ltv,
			APY:      borrowingAPY,
		})
	}

	log.Printf("Time elapsed: %v", time.Since(startTime))

	return &t.ProtocolMarkets{
		Protocol:       AaveV3Name,
		Chain:          a.chain,
		LendingSpecs:   lendingSpecs,
		BorrowingSpecs: borrowingSpecs,
	}, nil
}

// Deposits the specified token into the protocol
func (a *AaveV3) Supply(from string, token string, amount *big.Int) (*common.Hash, error) {
	methodName := "supply"
	// TODO: Check allowance
	assets, err := utils.ConvertSymbolsToAddresses(a.chain, []string{token})
	if err != nil {
		log.Printf("Failed to convert token symbol to address: %v", err)
		return nil, err
	}
	data, err := a.poolABI.Pack(
		methodName,
		common.HexToAddress(assets[0]),
		amount,
		common.HexToAddress(from), // onBehalfOf, can update to delegate borrowing to another address
		uint16(0),
	)
	if err != nil {
		log.Printf("Failed to pack method args: %v", err)
		return nil, err
	}

	// Build transaction
	log.Println("Building deposit tx...")
	tx, err := txs.BuildTransaction(a.cl, from, a.poolAddress, big.NewInt(0), data)
	if err != nil {
		log.Printf("Failed to build deposit transaction: %v", err)
		return nil, err
	}

	// Sign transaction
	signedTx, err := accounts.SignTransaction(from, *a.chainid, tx)
	if err != nil {
		log.Printf("Failed to sign tx: %v", err)
		return nil, err
	}

	// Send transaction
	hash, err := txs.SendTransaction(a.cl, signedTx)
	if err != nil {
		log.Printf("Failed to send deposit tx: %v", err)
		return nil, err
	}
	return hash, nil
}

// Borrows the specified token from the protocol
// func (a *AaveV3) Borrow(from string, token string, amount *big.Int) (*common.Hash, error) {

// }
