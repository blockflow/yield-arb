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
)

type TokenMapping struct {
	Tokens map[string]string `json:"tokens"`
}

var RPCEndpoints = map[string]string{
	"ethereum":  "https://eth-mainnet.g.alchemy.com/v2/NiPLhDKdUp9f7e6BPsQeW4lRXAo2rtbZ",
	"polygon":   "https://polygon-mainnet.g.alchemy.com/v2/NiPLhDKdUp9f7e6BPsQeW4lRXAo2rtbZ",
	"avalanche": "https://rpc.ankr.com/avalanche",
	"arbitrum":  "https://arb1.arbitrum.io/rpc",
}

var TokenMappings = make(map[string]*TokenMapping)
var TokenAliases *TokenMapping

func init() {
	for chain := range RPCEndpoints {
		if err := loadMapping(chain); err != nil {
			log.Printf("Failed to load %v mapping: %v", chain, err)
		}
	}

	if err := loadAliases(); err != nil {
		log.Printf("Failed to load aliases: %v", err)
	}
}

// Loads token mappings into TokenMappings var
func loadMapping(chain string) error {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "mappings", chain+"_tokens.json")
	rawMapping, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open token mapping: %v", err)
		return err
	}
	defer rawMapping.Close()
	var parsedMapping TokenMapping
	readMapping, err := io.ReadAll(rawMapping)
	if err != nil {
		log.Printf("Failed to read token mapping: %v", err)
	}
	err = json.Unmarshal(readMapping, &parsedMapping)
	if err != nil {
		log.Printf("Failed to parse token mapping: %v", err)
		return err
	}

	TokenMappings[chain] = &parsedMapping
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
	var parsedMapping TokenMapping
	readMapping, err := io.ReadAll(rawMapping)
	if err != nil {
		log.Printf("Failed to read token aliases: %v", err)
	}
	err = json.Unmarshal(readMapping, &parsedMapping)
	if err != nil {
		log.Printf("Failed to parse token aliases: %v", err)
		return err
	}

	TokenAliases = &parsedMapping
	return nil
}

// Converts the token symbols to their respective addresses for the specified chain
// If not mapped, token will be excluded.
func ConvertSymbolsToAddresses(chain string, symbols []string) ([]string, error) {
	mapping, ok := TokenMappings[chain]
	if !ok {
		msg := fmt.Sprintf("Could not find %v mapping", chain)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	var result []string
	for _, symbol := range symbols {
		address, ok := mapping.Tokens[symbol]
		if ok {
			result = append(result, address)
		}
	}

	return result, nil
}

// Converts the token addresses to their respective symbols for the specified chain
// If not mapped, token will be excluded.
func ConvertAddressesToSymbols(chain string, addresses []string) ([]string, error) {
	mapping, ok := TokenMappings[chain]
	if !ok {
		msg := fmt.Sprintf("Could not find %v mapping", chain)
		log.Println(msg)
		return nil, errors.New(msg)
	}

	// Reverse mapping O(tokens)
	reversedMapping := make(map[string]string)
	for token, address := range mapping.Tokens {
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
