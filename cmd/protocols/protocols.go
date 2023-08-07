package protocols

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
	GetMarkets() (*ProtocolMarkets, error)

	// Lends the token to the protocol
	Deposit(from string, token string, amount *big.Int) (*common.Hash, error)
	// // Withdraws the token from the protocol
	// Withdraw(user string, token string, amount *big.Int) error
	// // Borrows the token from the protocol
	// Borrow(user string, token string, amount *big.Int) error
	// // Repays the token to the protocol
	// Repay(user string, token string, amount *big.Int) error

	// // Fetches the user's positions and leverage
	// GetAccountData(user string)
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

type AccountData struct {
	TotalCollateralBase         *big.Int `json:"totalCollateralBase"`
	TotalDebtBase               *big.Int `json:"totalDebtBase"`
	AvailableBorrowsBase        *big.Int `json:"availableBorrowsBase"`
	CurrentLiquidationThreshold *big.Int `json:"currentLiquidationThreshold"`
	LTV                         *big.Int `json:"ltv"`
	HealthFactor                *big.Int `json:"healthFactor"`
}

func GetProtocol(protocol string) (Protocol, error) {
	switch protocol {
	case "aavev2":
		return NewAaveV2Protocol(), nil
	case "aavev3":
		return NewAaveV3Protocol(), nil
	case "compoundv2":
		return NewCompoundV2Protocol(), nil
	// case "dydx":
	// 	return NewDydxProtocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
