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

// TODO: Add notifications for when config changes

type AaveV2 struct {
	chain               string
	cl                  *ethclient.Client
	chainid             *big.Int
	lendingPoolContract *bind.BoundContract
}

type AaveV2ReserveData struct {
	Configurationstruct struct {
		Data *big.Int `json:"data"`
	} `json:"configuration"`
	LiquidityIndex              *big.Int       `json:"liquidityIndex"`
	VariableBorrowIndex         *big.Int       `json:"variableBorrowIndex"`
	CurrentLiquidityRate        *big.Int       `json:"currentLiquidityRate"`
	CurrentVariableBorrowRate   *big.Int       `json:"currentVariableBorrowRate"`
	CurrentStableBorrowRate     *big.Int       `json:"currentStableBorrowRate"`
	LastUpdateTimestamp         *big.Int       `json:"lastUpdateTimestamp"`
	ATokenAddress               common.Address `json:"aTokenAddress"`
	StableDebtTokenAddress      common.Address `json:"stableDebtTokenAddress"`
	VariableDebtTokenAddress    common.Address `json:"variableDebtTokenAddress"`
	InterestRateStrategyAddress common.Address `json:"interestRateStrategyAddress"`
	ID                          uint8          `json:"id"`
}

type AaveV2ReserveDataOutput struct {
	Output AaveV2ReserveData
}

const AaveV2Name = "aavev2"

var aavev2AddressesProviders = map[string]string{
	"ethereum":  "0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5",
	"polygon":   "0xd05e3E715d945B59290df0ae8eF85c1BdB684744",
	"avalanche": "0xb6A86025F0FE1862B372cb0ca18CE3EDe02A318f",

	// Testnets
	"ethereum_goerli": "0x5E52dEc931FFb32f609681B8438A51c675cc232d",
	"polygon_mumbai":  "0x178113104fEcbcD7fF8669a0150721e231F0FD4B",
	"avalanche_fuji":  "0x7fdC1FdF79BE3309bf82f4abdAD9f111A6590C0f",
}

func NewAaveV2Protocol() *AaveV2 {
	return &AaveV2{}
}

func (a *AaveV2) GetChains() ([]string, error) {
	return []string{"ethereum", "polygon", "avalanche"}, nil
}

func (a *AaveV2) Connect(chain string) error {
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
	path := filepath.Join(dir, "abis", AaveV2Name, "addresses_provider.json")
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
	addressesProvider := common.HexToAddress(aavev2AddressesProviders[chain])
	addressesProviderContract := bind.NewBoundContract(addressesProvider, parsedABI, cl, cl, cl)

	// Fetch lending pool address
	results := []interface{}{new(common.Address)}
	callOpts := &bind.CallOpts{}
	err = addressesProviderContract.Call(callOpts, &results, "getLendingPool")
	if err != nil {
		log.Printf("Failed to fetch lending pool: %v", err)
		return err
	}

	// Instantiate lending pool
	rawLendingPoolABI, err := os.Open(filepath.Join(dir, "abis", AaveV2Name, "lending_pool.json"))
	if err != nil {
		log.Printf("Failed to open lending_pool.json: %v", err)
		return err
	}
	defer rawLendingPoolABI.Close()
	parsedLendingPoolABI, err := abi.JSON(rawLendingPoolABI)
	if err != nil {
		log.Printf("Failed to parse lending_pool.json: %v", err)
		return err
	}
	lendingPoolContract := bind.NewBoundContract(*results[0].(*common.Address), parsedLendingPoolABI, cl, cl, cl)

	a.chain = chain
	a.cl = cl
	a.chainid = chainid
	a.lendingPoolContract = lendingPoolContract
	log.Printf("Connected to %v (chainid: %v, lendingpool: %v)", a.chain, a.chainid, *results[0].(*common.Address))
	return nil
}

// Returns the symbols of the supported tokens
func (a *AaveV2) getReservesList() ([]string, error) {
	results := []interface{}{new([]common.Address)}
	callOpts := &bind.CallOpts{}
	err := a.lendingPoolContract.Call(callOpts, &results, "getReservesList")
	if err != nil {
		log.Printf("Failed to fetch tokens: %v", err)
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
func (a *AaveV2) GetLendingTokens() ([]string, error) {
	return a.getReservesList()
}

// GetBorrowingTokens returns the tokens that can be borrowed on the given chain
func (a *AaveV2) GetBorrowingTokens() ([]string, error) {
	return a.getReservesList()
}

// Get ReserveData directly from smart contract
func (a *AaveV2) getReserveData(token string) (*AaveV2ReserveData, error) {
	result := []interface{}{new(AaveV2ReserveDataOutput)}
	callOpts := &bind.CallOpts{}
	err := a.lendingPoolContract.Call(callOpts, &result, "getReserveData", common.HexToAddress(token))
	if err != nil {
		log.Printf("Failed to fetch reserve data: %v", err)
		return nil, err
	}
	return &result[0].(*AaveV2ReserveDataOutput).Output, nil
}

// Get the lending TokenSpecs for the specified tokens. Filters out paused tokens.
func (a *AaveV2) getTokenSpecs(symbols []string, getAPY func(*AaveV2ReserveData) *big.Int) ([]*TokenSpecs, error) {
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
			Protocol: AaveV2Name,
			Chain:    a.chain,
			Token:    symbols[i],
			LTV:      ltv,
			APY:      liquidtyRate,
		})
	}
	return tokenSpecs, nil
}

func getLendingAPYV2(rd *AaveV2ReserveData) *big.Int {
	return rd.CurrentLiquidityRate
}

func getBorrowingAPYV2(rd *AaveV2ReserveData) *big.Int {
	return rd.CurrentVariableBorrowRate
}

func (a *AaveV2) GetLendingSpecs(symbols []string) ([]*TokenSpecs, error) {
	return a.getTokenSpecs(symbols, getLendingAPYV2)
}

func (a *AaveV2) GetBorrowingSpecs(symbols []string) ([]*TokenSpecs, error) {
	return a.getTokenSpecs(symbols, getBorrowingAPYV2)
}

// Returns the market.
// Assumes lending and borrowing tokens are the same.
func (a *AaveV2) GetMarkets() (*ProtocolMarkets, error) {
	log.Printf("Fetching market data for %v...", a.chain)
	lendingTokens, _ := (*a).GetLendingTokens()
	lendingSpecs, _ := (*a).GetLendingSpecs(lendingTokens)
	borrowingSpecs, _ := (*a).GetBorrowingSpecs(lendingTokens)

	return &ProtocolMarkets{
		Protocol:       AaveV2Name,
		Chain:          a.chain,
		LendingSpecs:   lendingSpecs,
		BorrowingSpecs: borrowingSpecs,
	}, nil
}

// Deposits the specified token into the protocol
func (a *AaveV2) Deposit(from string, token string, amount *big.Int) (*common.Hash, error) {
	return nil, nil
}
