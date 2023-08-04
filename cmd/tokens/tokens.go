package tokens

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type TokenMapping struct {
	Tokens map[string]string `json:"tokens"`
}

func loadMapping(chain string) (*TokenMapping, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)
	path := filepath.Join(dir, "mappings", chain+"_tokens.json")
	rawMapping, err := os.Open(path)
	if err != nil {
		log.Printf("Failed to open token mapping: %v", err)
		return nil, err
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
		return nil, err
	}
	return &parsedMapping, nil
}

// Converts the token symbols to their respective addresses for the specified chain
// If not mapped, token will be excluded.
func ConvertSymbolsToAddresses(chain string, symbols []string) ([]string, error) {
	mapping, err := loadMapping(chain)
	if err != nil {
		log.Printf("Failed to load mapping: %v", err)
		return nil, err
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
	mapping, err := loadMapping(chain)
	if err != nil {
		log.Printf("Failed to load mapping: %v", err)
		return nil, err
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
