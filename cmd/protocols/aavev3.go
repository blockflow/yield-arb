package protocols

import (
	"context"
	"log"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Account for new features such as isolation mode

type AaveV3 struct {
	chain        string
	cl           *ethclient.Client
	chainid      big.Int
	poolContract *bind.BoundContract
}

type AaveV3ReserveData struct {
	Configurationstruct struct {
		Data *big.Int `json:"data"`
	} `json:"configuration"`
	LiquidityIndex              *big.Int       `json:"liquidityIndex"`
	CurrentLiquidityRate        *big.Int       `json:"currentLiquidityRate"`
	VariableBorrowIndex         *big.Int       `json:"variableBorrowIndex"`
	CurrentVariableBorrowRate   *big.Int       `json:"currentVariableBorrowRate"`
	CurrentStableBorrowRate     *big.Int       `json:"currentStableBorrowRate"`
	LastUpdateTimestamp         *big.Int       `json:"lastUpdateTimestamp"`
	ID                          uint16         `json:"id"`
	ATokenAddress               common.Address `json:"aTokenAddress"`
	StableDebtTokenAddress      common.Address `json:"stableDebtTokenAddress"`
	VariableDebtTokenAddress    common.Address `json:"variableDebtTokenAddress"`
	InterestRateStrategyAddress common.Address `json:"interestRateStrategyAddress"`
	AccruedToTreasury           *big.Int       `json:"accruedToTreasury"`
	Unbacked                    *big.Int       `json:"unbacked"`
	IsolationModeTotalDebt      *big.Int       `json:"isolationModeTotalDebt"`
}

type AaveV3ReserveDataOutput struct {
	Output AaveV3ReserveData
}

const AaveV3Name = "aavev3"

var aavev3AddressesProviders = map[string]string{
	"ethereum":  "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e",
	"polygon":   "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"avalanche": "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"arbitrum":  "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	// "fantom",
	// "harmony",
	// "optimism",
	// "metis",

	// Testnets
	"ethereum_goerli": "0xC911B590248d127aD18546B186cC6B324e99F02c",
	"avalanche_fuji":  "0x220c6A7D868FC38ECB47d5E69b99e9906300286A",
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

	// Load addresses provider abi
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "abis", AaveV3Name, "addresses_provider.json")
	rawABI, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open addresses_provider.json: %v", err)
		return err
	}
	defer rawABI.Close()
	parsedABI, err := abi.JSON(rawABI)
	if err != nil {
		log.Printf("Failed to parse addresses_provider.json: %v", err)
		return err
	}

	// Instantiate addresses provider
	addressesProvider := common.HexToAddress(aavev3AddressesProviders[chain])
	addressesProviderContract := bind.NewBoundContract(addressesProvider, parsedABI, cl, cl, cl)

	// Fetch lending pool address
	results := []interface{}{new(common.Address)}
	callOpts := &bind.CallOpts{}
	err = addressesProviderContract.Call(callOpts, &results, "getPool")
	if err != nil {
		log.Printf("Failed to fetch lending pool: %v", err)
		return err
	}

	// Instantiate lending pool
	rawPoolABI, err := os.Open(filepath.Join(dir, "abis", AaveV3Name, "pool.json"))
	if err != nil {
		log.Printf("Failed to open pool.json: %v", err)
		return err
	}
	defer rawPoolABI.Close()
	parsedPoolABI, err := abi.JSON(rawPoolABI)
	if err != nil {
		log.Printf("Failed to parse pool.json: %v", err)
		return err
	}
	poolContract := bind.NewBoundContract(*results[0].(*common.Address), parsedPoolABI, cl, cl, cl)

	a.chain = chain
	a.cl = cl
	a.chainid = *chainid
	a.poolContract = poolContract
	log.Printf("Connected to %v (chainid: %v, pool: %v)", a.chain, a.chainid, *results[0].(*common.Address))
	return nil
}

// Returns the symbols of the supported tokens
func (a *AaveV3) getReservesList() ([]string, error) {
	results := []interface{}{new([]common.Address)}
	callOpts := &bind.CallOpts{}
	err := a.poolContract.Call(callOpts, &results, "getReservesList")
	if err != nil {
		log.Printf("Failed to fetch lending tokens: %v", err)
		return nil, err
	}
	addresses := *results[0].(*[]common.Address)

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

// Get ReserveData directly from smart contract
func (a *AaveV3) getReserveData(token string) (*AaveV3ReserveData, error) {
	result := []interface{}{new(AaveV3ReserveDataOutput)}
	callOpts := &bind.CallOpts{}
	err := a.poolContract.Call(callOpts, &result, "getReserveData", common.HexToAddress(token))
	if err != nil {
		log.Printf("Failed to fetch reserve data: %v", err)
		return nil, err
	}
	return &result[0].(*AaveV3ReserveDataOutput).Output, nil
}

// Get the lending TokenSpecs for the specified tokens. Filters out paused tokens.
func (a *AaveV3) getTokenSpecs(symbols []string, getAPY func(*AaveV3ReserveData) *big.Int) ([]*TokenSpecs, error) {
	addresses, err := utils.ConvertSymbolsToAddresses(a.chain, symbols)
	if err != nil {
		log.Printf("Failed to convert symbols to addresses: %v", err)
		return nil, err
	}
	if len(addresses) != len(symbols) {
		log.Printf("%v tokens lost during conversion", len(symbols)-len(addresses))
		return nil, err
	}

	var tokenSpecs []*TokenSpecs
	for i, address := range addresses {
		reserveData, err := a.getReserveData(address)
		if err != nil {
			log.Printf("Failed to fetch reserve data: %v", err)
			return nil, err
		}

		// Check if paused
		config := new(big.Int).Rsh(reserveData.Configurationstruct.Data, 58)
		config = config.And(config, big.NewInt(1))
		borrowingEnabled := config.Uint64() == 1
		if !borrowingEnabled {
			continue
		}

		// Calculate LTV in percentage
		mask := big.NewInt(0xFFFF)                                             // 16-bit mask
		ltvInt := new(big.Int).And(reserveData.Configurationstruct.Data, mask) // mask
		ltv := new(big.Float).SetInt(ltvInt)                                   // convert to float
		ltv.Quo(ltv, big.NewFloat(100))                                        // convert to percentage

		// Get the lending/borrowing rate
		liquidtyRate := utils.ConvertRayToPercentage(getAPY(reserveData))
		tokenSpecs = append(tokenSpecs, &TokenSpecs{
			Protocol: AaveV3Name,
			Chain:    a.chain,
			Token:    symbols[i],
			LTV:      ltv,
			APY:      liquidtyRate,
		})
	}
	return tokenSpecs, nil
}

func getLendingAPYV3(rd *AaveV3ReserveData) *big.Int {
	return rd.CurrentLiquidityRate
}

func getBorrowingAPYV3(rd *AaveV3ReserveData) *big.Int {
	return rd.CurrentVariableBorrowRate
}

func (a *AaveV3) GetLendingSpecs(symbols []string) ([]*TokenSpecs, error) {
	return a.getTokenSpecs(symbols, getLendingAPYV3)
}

func (a *AaveV3) GetBorrowingSpecs(symbols []string) ([]*TokenSpecs, error) {
	return a.getTokenSpecs(symbols, getBorrowingAPYV3)
}

// Returns the market.
// Assumes lending and borrowing tokens are the same.
// TODO: Check for supply capped markets
func (a *AaveV3) GetMarkets() (ProtocolMarkets, error) {
	log.Printf("Fetching market data for %v...", a.chain)
	lendingTokens, _ := (*a).GetLendingTokens()
	lendingSpecs, _ := (*a).GetLendingSpecs(lendingTokens)
	borrowingSpecs, _ := (*a).GetBorrowingSpecs(lendingTokens)

	return ProtocolMarkets{
		Protocol:       AaveV2Name,
		Chain:          a.chain,
		LendingSpecs:   lendingSpecs,
		BorrowingSpecs: borrowingSpecs,
	}, nil
}
