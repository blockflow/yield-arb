package protocols

import (
	"fmt"
	"math/big"

	"yield-arb/cmd/protocols/aavev3"
	"yield-arb/cmd/protocols/compoundv3"
	t "yield-arb/cmd/protocols/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Protocol interface {
	// Returns the chains supported by the protocol
	GetChains() ([]string, error)
	// Connects to the protocol on the specified chain
	Connect(chain string) error
	// // Returns the tokens supported by the protocol on the specified chain
	// // Tokens are represented as their symbols
	// GetLendingTokens() ([]string, error)
	// // Returns the tokens supported by the protocol on the specified chain
	// // Tokens are represented as their symbols
	// GetBorrowingTokens() ([]string, error)
	// // Returns the TokenSpecs for the specified tokens
	// // Tokens are represented as their symbols
	// GetLendingSpecs(symbols []string) ([]*t.TokenSpecs, error)
	// // Returns the TokenSpecs for the specified tokens
	// // Tokens are represented as their symbols
	// GetBorrowingSpecs(symbols []string) ([]*t.TokenSpecs, error)
	// Returns the markets for the protocol
	GetMarkets() (*t.ProtocolChain, error)

	// Lends the token to the protocol
	Supply(from common.Address, token string, amount *big.Int) (*types.Transaction, error)
	// // Withdraws the token from the protocol
	// Withdraw(user string, token string, amount *big.Int) error
	// Borrows the token from the protocol
	Borrow(from common.Address, token string, amount *big.Int) (*types.Transaction, error)
	// // Repays the token to the protocol
	// Repay(user string, token string, amount *big.Int) error

	// // Fetches the user's positions and leverage
	// GetAccountData(user string)
}

func GetProtocol(protocol string) (Protocol, error) {
	switch protocol {
	// case "aavev2":
	// 	return aavev2.NewAaveV2Protocol(), nil
	case "aavev3":
		return aavev3.NewAaveV3Protocol(), nil
	// case "compoundv2":
	// 	return compoundv2.NewCompoundV2Protocol(), nil
	case "compoundv3":
		return compoundv3.NewCompoundV3Protocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
