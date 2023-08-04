package protocols

import "math/big"

type Protocol interface {
	// Returns the chains supported by the protocol
	GetChains() ([]string, error)
	// Connects to the protocol on the specified chain
	Connect(chain string) error
	// Returns the tokens supported by the protocol on the specified chain
	// Tokens are represented as their symbols
	GetLendingTokens() ([]string, error)
	// Returns the tokens supported by the protocol on the specified chain
	// Tokens are represented as their symbols
	GetBorrowingTokens() ([]string, error)
	// Returns the APYs for the specified tokens
	// Tokens are represented as their symbols
	GetLendingAPYs(tokens []string) (map[string]*big.Float, error)
	// Returns the APYs for the specified tokens
	// Tokens are represented as their symbols
	GetBorrowingAPYs(tokens []string) (map[string]*big.Float, error)
}
