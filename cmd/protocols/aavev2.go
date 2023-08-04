package protocols

import (
	"context"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"yield-arb/cmd/tokens"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TODO: Map tokens to their addresses on different chains
// TODO: Add notifications for when config changes

var addressesProviderEthereum common.Address = common.HexToAddress("0xB53C1a33016B2DC2fF3653530bfF1848a515c8c5")
var addressesProviderPolygon common.Address = common.HexToAddress("0xd05e3E715d945B59290df0ae8eF85c1BdB684744")

type AaveV2 struct {
	chain               string
	cl                  *ethclient.Client
	chainid             big.Int
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

type ReserveDataOutput struct {
	Output AaveV2ReserveData
}

func NewAaveV2Protocol() *AaveV2 {
	return &AaveV2{}
}

func (a *AaveV2) GetChains() ([]string, error) {
	return []string{"ethereum", "polygon", "avalanche"}, nil
}

func (a *AaveV2) Connect(chain string) error {
	// Setup the client
	var cl *ethclient.Client
	switch chain {
	case "ethereum":
		_cl, clErr := ethclient.Dial("https://eth-mainnet.g.alchemy.com/v2/NiPLhDKdUp9f7e6BPsQeW4lRXAo2rtbZ")
		if clErr != nil {
			log.Printf("Failed to connect to the %v client: %v", chain, clErr)
			return clErr
		}
		cl = _cl
	case "polygon":
		_cl, clErr := ethclient.Dial("https://polygon-mainnet.g.alchemy.com/v2/NiPLhDKdUp9f7e6BPsQeW4lRXAo2rtbZ")
		if clErr != nil {
			log.Printf("Failed to connect to the %v client: %v", chain, clErr)
			return clErr
		}
		cl = _cl
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
	path := filepath.Join(dir, "abis", "aavev2", "addresses_provider.json")
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
	var addressesProviderContract *bind.BoundContract
	switch chain {
	case "ethereum":
		addressesProviderContract = bind.NewBoundContract(addressesProviderEthereum, parsedABI, cl, cl, cl)
	case "polygon":
		addressesProviderContract = bind.NewBoundContract(addressesProviderPolygon, parsedABI, cl, cl, cl)
	}

	// Fetch lending pool address
	results := []interface{}{new(common.Address)}
	callOpts := &bind.CallOpts{}
	err = addressesProviderContract.Call(callOpts, &results, "getLendingPool")
	if err != nil {
		log.Printf("Failed to fetch lending pool: %v", err)
		return err
	}

	// Instantiate lending pool
	rawLendingPoolABI, err := os.Open(filepath.Join(dir, "abis", "aavev2", "lending_pool.json"))
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
	a.chainid = *chainid
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
	symbols, err := tokens.ConvertAddressesToSymbols(a.chain, addressesString)
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

func (a *AaveV2) getReserveData(token string) (*AaveV2ReserveData, error) {
	result := []interface{}{new(ReserveDataOutput)}
	callOpts := &bind.CallOpts{}
	err := a.lendingPoolContract.Call(callOpts, &result, "getReserveData", common.HexToAddress(token))
	if err != nil {
		log.Printf("Failed to fetch reserve data: %v", err)
		return nil, err
	}
	return &result[0].(*ReserveDataOutput).Output, nil
}

// Convert a big.Int ray to a percentage
func convertRayToPercentage(ray *big.Int) *big.Float {
	rayAsFloat := new(big.Float).SetInt(ray)
	// Convert to 27 decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(27))
	rayAsFloat.Quo(rayAsFloat, divisor)
	// Convert to percentage
	rayAsFloat.Mul(rayAsFloat, big.NewFloat(100))
	return rayAsFloat
}

// TODO: Filter out paused tokens
// GetLendingAPYs returns the lending APYs for the given tokens
func (a *AaveV2) GetLendingAPYs(symbols []string) (map[string]*big.Float, error) {
	addresses, err := tokens.ConvertSymbolsToAddresses(a.chain, symbols)
	if err != nil {
		log.Printf("Failed to convert symbols to addresses: %v", err)
		return nil, err
	}
	if len(addresses) != len(symbols) {
		log.Printf("%v tokens lost during conversion", len(symbols)-len(addresses))
		return nil, err
	}

	lendingAPYs := make(map[string]*big.Float)
	for i, address := range addresses {
		reserveData, err := a.getReserveData(address)
		if err != nil {
			log.Printf("Failed to fetch reserve data: %v", err)
			return nil, err
		}
		liquidtyRate := convertRayToPercentage(reserveData.CurrentLiquidityRate)
		lendingAPYs[symbols[i]] = liquidtyRate
	}
	return lendingAPYs, nil
}

// GetBorrowingAPYs returns the variable borrowing APYs for the given tokens
func (a *AaveV2) GetBorrowingAPYs(symbols []string) (map[string]*big.Float, error) {
	addresses, err := tokens.ConvertSymbolsToAddresses(a.chain, symbols)
	if err != nil {
		log.Printf("Failed to convert symbols to addresses: %v", err)
		return nil, err
	}
	if len(addresses) != len(symbols) {
		log.Printf("%v tokens lost during conversion", len(symbols)-len(addresses))
		return nil, err
	}

	borrowingAPYs := make(map[string]*big.Float)
	for i, address := range addresses {
		reserveData, err := a.getReserveData(address)
		if err != nil {
			log.Printf("Failed to fetch reserve data: %v", err)
			return nil, err
		}
		variableBorrowRate := convertRayToPercentage(reserveData.CurrentVariableBorrowRate)
		borrowingAPYs[symbols[i]] = variableBorrowRate
	}
	return borrowingAPYs, nil
}
