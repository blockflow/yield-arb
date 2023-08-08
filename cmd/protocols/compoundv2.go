package protocols

import (
	"context"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"yield-arb/cmd/utils"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CompoundV2 struct {
	chain               string
	cl                  *ethclient.Client
	chainId             *big.Int
	cetherAddress       common.Address
	comptrollerContract *bind.BoundContract
	cerc20ParsedABI     *abi.ABI
	cetherContract      *bind.BoundContract
}

const CompoundV2Name = "compoundv2"

// TODO: Update for testnet
const compv2ComptrollerAddress = "0x3d9819210A31b4961b30EF54bE2aeD79B9c9Cd3B"
const compv2CetherAddress = "0x4Ddc2D193948926D02f9B1fE9e1daa0718270ED5"
const compv2ComptrollerName = "comptroller"
const compv2Cerc20Name = "cerc20"
const compv2CetherName = "cether"

// Mapping of token symbols to cERC20 address.
// Might need to be refreshed.
var compv2Markets map[string]common.Address

func NewCompoundV2Protocol() *CompoundV2 {
	return &CompoundV2{}
}

// Returns the chains supported by the protocol
func (c *CompoundV2) GetChains() ([]string, error) {
	return []string{"ethereum", "ethereum_goerli"}, nil
}

// Connects to the protocol on the specified chain
func (c *CompoundV2) Connect(chain string) error {
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

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	c.cetherAddress = common.HexToAddress(compv2CetherAddress)

	// Instantiate comptroller
	comptrollerPath := filepath.Join(dir, "abis", CompoundV2Name, compv2ComptrollerName+".json")
	comptrollerRawABI, err := os.Open(comptrollerPath)
	if err != nil {
		log.Printf("Failed to open %v.json: %v", compv2ComptrollerName, err)
		return err
	}
	defer comptrollerRawABI.Close()
	comptrollerParsedABI, err := abi.JSON(comptrollerRawABI)
	if err != nil {
		log.Printf("Failed to parse %v.json: %v", compv2ComptrollerName, err)
		return err
	}
	comptrollerContract := bind.NewBoundContract(common.HexToAddress(compv2ComptrollerAddress), comptrollerParsedABI, cl, cl, cl)

	// Instantiate cerc20 abi
	cerc20Path := filepath.Join(dir, "abis", CompoundV2Name, compv2Cerc20Name+".json")
	cerc20RawABI, err := os.Open(cerc20Path)
	if err != nil {
		log.Printf("Failed to open %v.json: %v", compv2Cerc20Name, err)
		return err
	}
	defer cerc20RawABI.Close()
	cerc20ParsedABI, err := abi.JSON(cerc20RawABI)
	if err != nil {
		log.Printf("Failed to parse %v.json: %v", compv2Cerc20Name, err)
		return err
	}

	// Instantiate cether
	cetherPath := filepath.Join(dir, "abis", CompoundV2Name, "cether.json")
	cetherRawABI, err := os.Open(cetherPath)
	if err != nil {
		log.Printf("Failed to open %v.json: %v", compv2CetherName, err)
		return err
	}
	defer cetherRawABI.Close()
	cetherParsedABI, err := abi.JSON(cetherRawABI)
	if err != nil {
		log.Printf("Failed to parse %v.json: %v", compv2CetherName, err)
		return err
	}
	cetherContract := bind.NewBoundContract(c.cetherAddress, cetherParsedABI, cl, cl, cl)

	c.chain = chain
	c.cl = cl
	c.chainId = chainId
	c.comptrollerContract = comptrollerContract
	c.cerc20ParsedABI = &cerc20ParsedABI
	c.cetherContract = cetherContract

	// Load erc20 mapping
	if err := c.loadERC20Mapping(); err != nil {
		log.Printf("Failed to load erc20 mapping: %v", err)
		return err
	}

	log.Printf("Connected to %v (chainid: %v)", c.chain, c.chainId)
	return nil
}

// Returns the addresses of the listed cerc20 contracts.
// Requires c.comptrollerContract to be set.
func (c *CompoundV2) getAllMarkets() ([]common.Address, error) {
	results := []interface{}{new([]common.Address)}
	err := c.comptrollerContract.Call(&bind.CallOpts{}, &results, "getAllMarkets")
	if err != nil {
		log.Printf("Failed to fetch tokens: %v", err)
		return nil, err
	}
	addresses := *results[0].(*[]common.Address)

	return addresses, nil
}

// Loads the erc20 mapping.
func (c *CompoundV2) loadERC20Mapping() error {
	var newERC20Mapping = make(map[string]common.Address)
	markets, err := c.getAllMarkets()
	if err != nil {
		log.Printf("Failed to get all markets: %v", err)
		return err
	}
	for _, cerc20Address := range markets {
		cerc20Contract := c.loadCERC20Contract(cerc20Address)

		// Fetch erc20 symbol
		symbolResults := []interface{}{new(string)}
		if err = cerc20Contract.Call(&bind.CallOpts{}, &symbolResults, "symbol"); err != nil {
			log.Printf("Failed to fetch erc20 symbol: %v", err)
			return err
		}
		// Trim the "c" prefix
		erc20Symbol := strings.TrimPrefix(*symbolResults[0].(*string), "c")

		newERC20Mapping[erc20Symbol] = cerc20Address
	}

	compv2Markets = newERC20Mapping
	return nil
}

// Returns the symbols of the supported tokens.
// Fetches symbols on chain.
func (c *CompoundV2) getAllTokens() ([]string, error) {
	results := []interface{}{new([]common.Address)}
	err := c.comptrollerContract.Call(&bind.CallOpts{}, &results, "getAllMarkets")
	if err != nil {
		log.Printf("Failed to fetch tokens: %v", err)
		return nil, err
	}
	addresses := *results[0].(*[]common.Address)

	// Convert to symbols
	symbols := make([]string, len(addresses))
	for i, address := range addresses {
		cerc20Contract := c.loadCERC20Contract(address)
		symbolResults := []interface{}{new(string)}
		if err := cerc20Contract.Call(&bind.CallOpts{}, &symbolResults, "symbol"); err != nil {
			log.Printf("Failed to fetch token symbol: %v", err)
			return nil, err
		}
		// Trim the "c" prefix
		symbol := strings.TrimPrefix(*symbolResults[0].(*string), "c")

		symbols[i] = symbol
	}

	return symbols, nil
}

// Returns the tokens supported by the protocol on the specified chain
// Tokens are represented as their symbols
func (c *CompoundV2) GetLendingTokens() ([]string, error) {
	return c.getAllTokens()
}

// Returns the tokens supported by the protocol on the specified chain
// Tokens are represented as their symbols
func (c *CompoundV2) GetBorrowingTokens() ([]string, error) {
	return c.getAllTokens()
}

func (c *CompoundV2) loadCERC20Contract(cerc20Address common.Address) *bind.BoundContract {
	return bind.NewBoundContract(cerc20Address, *c.cerc20ParsedABI, c.cl, c.cl, c.cl)
}

// Base method for fetching token specs.
// Skips ltv if borrow.
func (c *CompoundV2) getTokenSpecs(symbols []string, rateMethod string) ([]*TokenSpecs, error) {
	specs := make([]*TokenSpecs, len(symbols))
	for i, symbol := range symbols {
		cerc20Address := compv2Markets[symbol]
		cerc20Contract := c.loadCERC20Contract(cerc20Address)
		// Fetch ltv only if lend
		ltv := new(big.Float)
		if rateMethod == "supplyRatePerBlock" {
			type MarketOutput struct {
				IsListed                 bool
				CollateralFactorMantissa *big.Int
				IsComped                 bool
			}
			marketResult := []interface{}{new(MarketOutput)}
			if err := c.comptrollerContract.Call(&bind.CallOpts{}, &marketResult, "markets", cerc20Address); err != nil {
				log.Printf("Failed to fetch %v ltv: %v", symbol, err)
				return nil, err
			}
			ltv.SetInt(marketResult[0].(*MarketOutput).CollateralFactorMantissa)
			ltv.Quo(ltv, new(big.Float).SetUint64(10000000000000000))
		} else {
			ltv.SetInt64(0)
		}

		// Calculate APY: APY = ((((Rate / ETH Mantissa * Blocks Per Day + 1) ^ Days Per Year)) - 1) * 100
		type SupplyRatePerBlockOutput struct {
			RatePerBlock *big.Int
		}
		supplyRatePerBlockResults := []interface{}{new(SupplyRatePerBlockOutput)}
		startTime := time.Now()
		if err := cerc20Contract.Call(&bind.CallOpts{}, &supplyRatePerBlockResults, rateMethod); err != nil {
			log.Printf("Failed to fetch %v rate: %v", symbol, err)
			return nil, err
		}
		endTime := time.Now()
		elapsedTime := endTime.Sub(startTime)
		log.Printf("Elapsed time: %v", elapsedTime)
		supplyRatePerBlock := new(big.Float).SetInt(supplyRatePerBlockResults[0].(*SupplyRatePerBlockOutput).RatePerBlock)
		dailyRate := new(big.Float).Quo(supplyRatePerBlock, utils.ETHMantissa)
		dailyRate.Mul(dailyRate, utils.ETHBlocksPerDay)
		dailyRate.Add(dailyRate, big.NewFloat(1))
		dailyRateFloat, _ := dailyRate.Float64()
		yearlyTotal := math.Pow(dailyRateFloat, 365) // Days per year
		apy := big.NewFloat(yearlyTotal)
		apy.Sub(apy, big.NewFloat(1))
		apy.Mul(apy, big.NewFloat(100))

		specs[i] = &TokenSpecs{
			Protocol: CompoundV2Name,
			Chain:    c.chain,
			Token:    symbol,
			LTV:      ltv,
			APY:      apy,
		}
	}
	return specs, nil
}

// Returns the TokenSpecs for the specified tokens
// Tokens are represented as their symbols
func (c *CompoundV2) GetLendingSpecs(symbols []string) ([]*TokenSpecs, error) {
	return c.getTokenSpecs(symbols, "supplyRatePerBlock")
}

// Returns the TokenSpecs for the specified tokens
// Tokens are represented as their symbols
func (c *CompoundV2) GetBorrowingSpecs(symbols []string) ([]*TokenSpecs, error) {
	return c.getTokenSpecs(symbols, "borrowRatePerBlock")
}

// Returns the markets for the protocol
func (c *CompoundV2) GetMarkets() (*ProtocolMarkets, error) {
	log.Println("Fetching markets...")
	symbols, err := c.GetLendingTokens()
	if err != nil {
		log.Printf("Failed to fetch tokens: %v", err)
	}
	lendingSpecs, err := c.GetLendingSpecs(symbols)
	if err != nil {
		log.Printf("Failed to fetch lending specs: %v", err)
	}
	borrowingSpecs, err := c.GetBorrowingSpecs(symbols)
	if err != nil {
		log.Printf("Failed to fetch borrowing specs: %v", err)
	}
	return &ProtocolMarkets{
		Protocol:       CompoundV2Name,
		Chain:          c.chain,
		LendingSpecs:   lendingSpecs,
		BorrowingSpecs: borrowingSpecs,
	}, nil
}

// // Lends the token to the protocol
func (c *CompoundV2) Deposit(from string, token string, amount *big.Int) (*common.Hash, error) {
	return nil, nil
}

// TODO: Handle ETH transactions differently!!!
