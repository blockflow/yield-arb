package protocols

import (
	"fmt"
	"math/big"
)

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
	// Returns the markets for the protocol
	GetMarkets() (ProtocolMarkets, error)
}

type ProtocolMarkets struct {
	Protocol      string                `json:"protocol"`
	Chain         string                `json:"chain"`
	LendingAPYs   map[string]*big.Float `json:"lendingAPYs"`
	BorrowingAPYs map[string]*big.Float `json:"borrowingAPYs"`
}

func GetProtocol(protocol string) (Protocol, error) {
	switch protocol {
	case "aavev2":
		return NewAaveV2Protocol(), nil
	// case "compound":
	// 	return NewCompoundProtocol(), nil
	// case "dydx":
	// 	return NewDydxProtocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
