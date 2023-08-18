package arbitrage

import (
	"fmt"
	"log"
	"math/big"
	p "yield-arb/cmd/protocols"
	t "yield-arb/cmd/protocols/types"
)

// Executes the arbitrage strategy.
// pcs: The protocol chains to execute the strategy on. Keys are protocol:chain.
// strategy: The strategy to execute.
// initialCollateralAmount: The amount of collateral to start with.
// ltvFactor: The factor to multiply the LTV by. Safety factor.
func EnterStrategy(wallet string, ps map[string]*p.Protocol, strategy []*t.MarketInfo, initialCollateralAmount *big.Float, ltvFactor *big.Float) error {
	currentAmount := initialCollateralAmount
	for i, market := range strategy {
		p, ok := ps[market.Protocol+":"+market.Chain]
		if !ok {
			return fmt.Errorf("protocol chain not found: %s:%s", market.Protocol, market.Chain)
		}
		if i%2 == 0 { // Lend
			amount, _ := currentAmount.Int(nil)
			tx, err := (*p).Supply(wallet, market.Token, amount)
			if err != nil {
				return fmt.Errorf("failed to supply: %v", err)
			}
			log.Printf("Lent %v %v to %v (%v)", amount, market.Token, market.Protocol, tx.Hash())
		} else { // Borrow
			// Calculate the amount to lend
			ltv := new(big.Float).Quo(market.LTV, big.NewFloat(100))
			ltv.Quo(ltv, ltvFactor)
			currentAmount.Mul(currentAmount, ltv)
			amount, _ := currentAmount.Int(nil)
			tx, err := (*p).Borrow(wallet, market.Token, amount)
			if err != nil {
				return fmt.Errorf("failed to borrow: %v", err)
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
