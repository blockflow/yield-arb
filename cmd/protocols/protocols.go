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
	// Returns the TokenSpecs for the specified tokens
	// Tokens are represented as their symbols
	GetLendingSpecs(symbols []string) ([]*TokenSpecs, error)
	// Returns the TokenSpecs for the specified tokens
	// Tokens are represented as their symbols
	GetBorrowingSpecs(symbols []string) ([]*TokenSpecs, error)
	// Returns the markets for the protocol
	GetMarkets() (ProtocolMarkets, error)
}

type ProtocolMarkets struct {
	Protocol       string        `json:"protocol"`
	Chain          string        `json:"chain"`
	LendingSpecs   []*TokenSpecs `json:"lendingSpecs"`
	BorrowingSpecs []*TokenSpecs `json:"borrowingSpecs"`
}

type TokenSpecs struct {
	Protocol string     `json:"protocol"`
	Chain    string     `json:"chain"`
	Token    string     `json:"token"`
	LTV      *big.Float `json:"ltv"` // 0 if cannot be collateral
	APY      *big.Float `json:"apy"`
}

func GetProtocol(protocol string) (Protocol, error) {
	switch protocol {
	case "aavev2":
		return NewAaveV2Protocol(), nil
	case "aavev3":
		return NewAaveV3Protocol(), nil
	// case "compound":
	// 	return NewCompoundProtocol(), nil
	// case "dydx":
	// 	return NewDydxProtocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
