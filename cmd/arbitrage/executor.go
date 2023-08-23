package arbitrage

import (
	"fmt"
	"log"
	"math"
	"math/big"
	p "yield-arb/cmd/protocols"
	t "yield-arb/cmd/protocols/types"
)

// EnterStrategy executes the arbitrage strategy.
//
// Parameters:
//   - wallet: The user's wallet address.
//   - ps: The protocol chains to execute the strategy on. Keys are defined as protocol:chain.
//   - strategy: A slice containing market information to determine the strategy to execute.
//   - initialCollateralAmount: The initial amount of collateral to start the strategy with.
//   - ltvFactor: The factor to multiply the Loan-to-Value (LTV) by. It serves as a safety factor and is represented in percent.
//
// Returns:
//   - error: Returns an error if the strategy execution fails.
func EnterStrategy(wallet string, ps map[string]*p.Protocol, strategy []*t.MarketInfo, initialCollateralAmount *big.Float, ltvFactor *big.Float) error {
	if len(strategy) == 0 {
		return fmt.Errorf("strategy is empty")
	}
	currentAmountUSD := new(big.Float).Mul(initialCollateralAmount, strategy[0].PriceInUSD)
	currentAmountUSD.Quo(currentAmountUSD, big.NewFloat(math.Pow(10, float64(strategy[0].Decimals))))
	for i, market := range strategy {
		p, ok := ps[market.Protocol+":"+market.Chain]
		if !ok {
			return fmt.Errorf("protocol chain not found: %s:%s", market.Protocol, market.Chain)
		}
		if i%2 == 0 { // Lend
			currentAmount := new(big.Float).Quo(currentAmountUSD, market.PriceInUSD)
			currentAmount.Mul(currentAmount, big.NewFloat(math.Pow(10, float64(market.Decimals))))
			amount, _ := currentAmount.Int(nil)
			tx, err := (*p).Supply(wallet, market.Token, amount)
			if err != nil {
				return fmt.Errorf("failed to supply: %v", err)
			}
			log.Printf("Lent %v %v to %v (%v)", amount, market.Token, market.Protocol, tx.Hash())

			// Calculate next amount
			ltvAdjusted := new(big.Float).Mul(market.LTV, ltvFactor)
			ltvAdjusted.Quo(ltvAdjusted, big.NewFloat(10000))
			currentAmountUSD.Mul(currentAmountUSD, ltvAdjusted)
		} else { // Borrow
			// Calculate the amount to lend
			currentAmount := new(big.Float).Quo(currentAmountUSD, market.PriceInUSD)
			currentAmount.Mul(currentAmount, big.NewFloat(math.Pow(10, float64(market.Decimals))))
			amount, _ := currentAmount.Int(nil)
			tx, err := (*p).Borrow(wallet, market.Token, amount)
			if err != nil {
				return fmt.Errorf("failed to borrow %v of %v: %v", amount, market.Token, err)
			}
			log.Printf("Borrowed %v %v from %v (%v)", amount, market.Token, market.Protocol, tx.Hash())
		}
	}
	return nil
}

// Exit the strategy and liquidate all positions.
// Repays slightly more than debt to account for interest.
func ExitStrategy(wallet string, ps map[string]*p.Protocol, strategy []*t.MarketInfo) error {
	// Starting with the end, withdraw all collateral and repay all debt
	for i := len(strategy) - 1; i >= 0; i-- {
		market := strategy[i]
		p, ok := ps[market.Protocol+":"+market.Chain]
		if !ok {
			return fmt.Errorf("protocol chain not found: %s:%s", market.Protocol, market.Chain)
		}
		if i%2 == 0 { // Withdraw
			tx, err := (*p).WithdrawAll(wallet, market.Token)
			if err != nil {
				return fmt.Errorf("failed to withdraw: %v", err)
			}
			log.Printf("Withdrew all %v from %v (%v)", market.Token, market.Protocol, tx.Hash())
		} else { // Repay
			tx, err := (*p).RepayAll(wallet, market.Token)
			if err != nil {
				return fmt.Errorf("failed to repay: %v", err)
			}
			log.Printf("Repayed all %v to %v (%v)", market.Token, market.Protocol, tx.Hash())
		}
	}
	return nil
}
