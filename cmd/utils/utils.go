package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"math/big"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type ChainConfig struct {
	RPCEndpoint string            `json:"rpcEndpoint"`
	Tokens      map[string]string `json:"tokens"`
}

var ChainConfigs = make(map[string]*ChainConfig)
var TokenAliases map[string]string

// Constants
var SecPerYear = big.NewFloat(60 * 60 * 24 * 365)
var ETHMantissa = new(big.Float).SetUint64(1000000000000000000) // 10**18
var ETHBlocksPerDay = big.NewFloat(7200)
var BlocksPerDay = map[string]*big.Float{
	"ethereum": big.NewFloat(7200),
	"arbitrum": big.NewFloat(105120000), // Avg block time of 0.3s
}

func init() {
	// Parse all config json files
	_, filename, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(filename)
	configPath := filepath.Join(currentDir, "configs")
	configDir, _ := os.ReadDir(configPath)
	for _, config := range configDir {
		log.Printf("Loading %v...", config.Name())
		if err := loadConfig(config.Name()); err != nil {
			log.Printf("Failed to load %v: %v", config.Name(), err)
		}
	}
	log.Println("Loading token aliases...")
	if err := loadAliases(); err != nil {
		log.Printf("Failed to load aliases: %v", err)
	}
}

// Loads the chain config including rpc endpoint and token mappings.
func loadConfig(configName string) error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "configs", configName)
	rawConfig, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open config: %v", err)
		return err
	}
	defer rawConfig.Close()
	var parsedConfig ChainConfig
	readConfig, err := io.ReadAll(rawConfig)
	if err != nil {
		log.Printf("Failed to read config: %v", err)
	}
	err = json.Unmarshal(readConfig, &parsedConfig)
	if err != nil {
		log.Printf("Failed to parse config: %v", err)
		return err
	}

	ChainConfigs[strings.TrimSuffix(configName, ".json")] = &parsedConfig
	return nil
}

// Loads the alias mapping into TokenAliases var
func loadAliases() error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "mappings", "token_aliases.json")
	rawMapping, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open token aliases: %v", err)
		return err
	}
	defer rawMapping.Close()
	var parsedMapping = make(map[string]string)
	readMapping, err := io.ReadAll(rawMapping)
	if err != nil {
		log.Printf("Failed to read token aliases: %v", err)
	}
	err = json.Unmarshal(readMapping, &parsedMapping)
	if err != nil {
		log.Printf("Failed to parse token aliases: %v", err)
		return err
	}

	TokenAliases = parsedMapping
	return nil
}

// Converts the token symbol to its respective address for the specified chain.
func ConvertSymbolToAddress(chain string, symbol string) (string, error) {
	config, ok := ChainConfigs[chain]
	if !ok {
		return "", fmt.Errorf("could not find %v config", chain)
	}

	address, ok := config.Tokens[symbol]
	if ok {
		return address, nil
	}
	return "", fmt.Errorf("could not find address for %v", symbol)
}

// Converts the token symbols to their respective addresses for the specified chain
// If not mapped, token will be excluded.
func ConvertSymbolsToAddresses(chain string, symbols []string) ([]string, error) {
	config, ok := ChainConfigs[chain]
	if !ok {
		msg := fmt.Sprintf("Could not find %v config", chain)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	var result []string
	for _, symbol := range symbols {
		address, ok := config.Tokens[symbol]
		if ok {
			result = append(result, address)
		}
	}

	return result, nil
}

func ConvertAddressToSymbol(chain string, address string) (string, error) {
	config, ok := ChainConfigs[chain]
	if !ok {
		return "", fmt.Errorf("could not find %v config", chain)
	}

	// Reverse mapping O(tokens)
	reversedMapping := make(map[string]string)
	for token, address := range config.Tokens {
		reversedMapping[address] = token
	}

	// Convert address to symbols O(addresses)
	symbol, ok := reversedMapping[address]
	if ok {
		return symbol, nil
	}
	return "", fmt.Errorf("could not find symbol for %v", address)
}

// Converts the token addresses to their respective symbols for the specified chain
// If not mapped, token will be excluded.
// TODO: cache reverse mapping
func ConvertAddressesToSymbols(chain string, addresses []string) ([]string, error) {
	config, ok := ChainConfigs[chain]
	if !ok {
		msg := fmt.Sprintf("Could not find %v config", chain)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	// Reverse mapping O(tokens)
	reversedMapping := make(map[string]string)
	for token, address := range config.Tokens {
		reversedMapping[address] = token
	}

	// Convert addresses to symbols O(addresses)
	var result []string
	for _, address := range addresses {
		symbol, ok := reversedMapping[address]
		if ok {
			result = append(result, symbol)
		}
	}

	return result, nil
}

// Returns the common symbol if an alias, otherwise return the arg.
func CommonSymbol(symbol string) string {
	if commonSymbol, ok := TokenAliases[symbol]; ok {
		return commonSymbol
	}
	return symbol
}

// Convert a big.Int ray to a percentage
func ConvertRayToPercentage(ray *big.Int) *big.Float {
	rayAsFloat := new(big.Float).SetInt(ray)
	// Convert to 27 decimals
	divisor := new(big.Float).SetFloat64(math.Pow10(27))
	rayAsFloat.Quo(rayAsFloat, divisor)
	// Convert to percentage
	rayAsFloat.Mul(rayAsFloat, big.NewFloat(100))
	return rayAsFloat
}

// Converts a per block rate into APY based on ETH block times.
func ConvertRatePerBlockToAPY(ratePerBlock *big.Int) *big.Float {
	ratePerBlockFloat := new(big.Float).SetInt(ratePerBlock)
	dailyRate := new(big.Float).Quo(ratePerBlockFloat, ETHMantissa)
	dailyRate.Mul(dailyRate, ETHBlocksPerDay)
	dailyRate.Add(dailyRate, big.NewFloat(1))
	dailyRateFloat, _ := dailyRate.Float64()
	yearlyTotal := math.Pow(dailyRateFloat, 365) // Days per year
	apy := big.NewFloat(yearlyTotal)
	apy.Sub(apy, big.NewFloat(1))
	apy.Mul(apy, big.NewFloat(100))
	return apy
}
