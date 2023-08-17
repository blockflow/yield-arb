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
		if i%2 == 1 { // Lend
			p, ok := ps[market.Protocol+":"+market.Chain]
			if !ok {
				return fmt.Errorf("protocol chain not found: %s:%s", market.Protocol, market.Chain)
			}
			// Lend
			amount, _ := currentAmount.Int(nil)
			tx, err := (*p).Supply(wallet, market.Token, amount)
			if err != nil {
				return fmt.Errorf("failed to supply: %v", err)
			}
			log.Printf("Lent %v %v to %v (%v)", amount, market.Token, market.Protocol, tx.Hash())
		} else { // Borrow
			p, ok := ps[market.Protocol+":"+market.Chain]
			if !ok {
				return fmt.Errorf("protocol chain not found: %s:%s", market.Protocol, market.Chain)
			}
			// Calculate the amount to lend
			ltv := new(big.Float).Quo(market.LTV, big.NewFloat(100))
			ltv.Quo(ltv, ltvFactor)
			currentAmount.Mul(currentAmount, ltv)
			amount, _ := currentAmount.Int(nil)
			// Borrow
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
func ExitStrategy(wallet string, ps map[string]*p.Protocol, strategy []*t.MarketInfo) error {
	// Starting with the end, withdraw all collateral and repay all debt

	return nil
}
