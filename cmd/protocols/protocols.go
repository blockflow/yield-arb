package protocols

import (
	"fmt"
	"math/big"

	"yield-arb/cmd/protocols/aavev3"
	"yield-arb/cmd/protocols/compoundv3"
	"yield-arb/cmd/protocols/schema"

	"github.com/ethereum/go-ethereum/core/types"
)

type Protocol interface {
	// Returns the chains supported by the protocol
	GetChains() ([]string, error)
	// Connects to the protocol on the specified chain
	Connect(chain string) error
	// Returns the markets for the protocol
	GetMarkets() ([]*schema.ProtocolChain, error)
	// Returns the APY and actual amount for the given token.
	// Actual amount is the amount that can be supplied/borrowed.
	// APY in ray.
	CalcAPY(market *schema.MarketInfo, amount *big.Int, isSupply bool) (*big.Int, *big.Int, error)

	// // Returns the supply/borrow token balances for the wallet.
	// // Positive balances are supplied, negative balances are borrowed.
	// GetBalances(wallet string) (map[string]*big.Int, error)

	// Returns the transactions required to execute the strategy step.
	GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error)
	// // Lends the token to the protocol
	// Supply(wallet string, token string, amount *big.Int) ([]*types.Transaction, error)
	// // Withdraws the token from the protocol
	// Withdraw(wallet string, token string, amount *big.Int) (*types.Transaction, error)
	// // Withdraws all of the token from the protocol
	// WithdrawAll(wallet string, token string) (*types.Transaction, error)
	// // Borrows the token from the protocol
	// Borrow(wallet string, token string, amount *big.Int) (*types.Transaction, error)
	// // Repays the token to the protocol
	// Repay(wallet string, token string, amount *big.Int) (*types.Transaction, error)
	// // Repays all of the token to the protocol
	// RepayAll(wallet string, token string) (*types.Transaction, error)
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
	// case "dforce":
	// 	return dforce.NewDForceProtocol(), nil
	// case "lodestar":
	// 	return lodestar.NewLodestarProtocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
