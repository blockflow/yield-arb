package arbitrage

import (
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/types"
	"yield-arb/cmd/utils"

	"golang.org/x/exp/slices"
)

// // Calculates the net APY for the tokenspec triplets
// func calculateNetAPY(specs []*t.MarketInfo) *big.Float {
// 	loanAPY := new(big.Float).Sub(specs[2].SupplyAPY, specs[1].BorrowAPY)
// 	loanAPY.Mul(loanAPY, specs[0].LTV)
// 	loanAPY.Quo(loanAPY, big.NewFloat(100))
// 	netAPY := new(big.Float).Add(specs[0].SupplyAPY, loanAPY)
// 	return netAPY
// }

// // Calculates the net APY for any odd number of TokenSpecs.
// // TokenSpecs should be in alternating order of lend/borrow starting with lend.
// func CalculateNetAPYV2(specs []*t.MarketInfo) *big.Float {
// 	if len(specs) == 0 {
// 		// Base case
// 		return big.NewFloat(0)
// 	} else if len(specs)%2 == 0 {
// 		// Even, intermediate case
// 		assetNetAPY := new(big.Float).Sub(specs[1].SupplyAPY, specs[0].BorrowAPY)
// 		// If borrowing again
// 		nextLevelAPY := CalculateNetAPYV2(specs[2:])
// 		ltv := new(big.Float).Quo(specs[0].LTV, big.NewFloat(100))
// 		nextLevelAPY.Mul(nextLevelAPY, ltv)
// 		return assetNetAPY.Add(assetNetAPY, nextLevelAPY)
// 	} else {
// 		// Odd, start case
// 		nextLevelAPY := CalculateNetAPYV2(specs[1:])
// 		ltv := new(big.Float).Quo(specs[0].LTV, big.NewFloat(100))
// 		nextLevelAPY.Mul(nextLevelAPY, ltv)
// 		return new(big.Float).Add(specs[0].SupplyAPY, nextLevelAPY)
// 	}
// }

// // Compares the two tokenspec triplets' net APYs.
// // Returns true if a is larger than b, false otherwise.
// func moreNetAPY(a, b []*t.MarketInfo) bool {
// 	return calculateNetAPY(a).Cmp(calculateNetAPY((b))) == 1
// }

// // Calculates the markets' min cap in USD.
// func calcMarketsCap(markets []*t.MarketInfo) *big.Float {
// 	currentMarket := markets[0]
// 	marketsLen := len(markets)
// 	if marketsLen == 1 { // Base case, lend
// 		return new(big.Float).Mul(currentMarket.SupplyLiquidity, currentMarket.PriceInUSD)
// 	} else if marketsLen%2 == 0 { // Borrow
// 		// Return the min
// 		subCapUSD := calcMarketsCap(markets[1:])
// 		currentCapUSD := new(big.Float).Mul(currentMarket.BorrowLiquidity, currentMarket.PriceInUSD)
// 		if currentCapUSD.Cmp(subCapUSD) == -1 {
// 			return currentCapUSD
// 		} else {
// 			return subCapUSD
// 		}
// 	} else { // Lend
// 		ltv := new(big.Float).Quo(currentMarket.LTV, big.NewFloat(100))
// 		subCap := new(big.Float).Mul(calcMarketsCap(markets[1:]), ltv)
// 		currentCapUSD := new(big.Float).Mul(currentMarket.SupplyLiquidity, currentMarket.PriceInUSD)
// 		if currentMarket.SupplyLiquidity.Cmp(subCap) == -1 {
// 			return currentCapUSD
// 		} else {
// 			return subCap
// 		}
// 	}
// }

// func CalculateStratV2CapsUSD(strat map[string][]*t.MarketInfo) map[string]*big.Float {
// 	result := make(map[string]*big.Float)
// 	for collateral, markets := range strat {
// 		result[collateral] = calcMarketsCap(markets)
// 	}
// 	return result
// }

// Finds all the strategies that can be made with the given tokens using BFS in reverse.
//
// Params:
//   - pcs: List of protocol chains
//   - maxLevels: Maximum number of levels to search
//
// Returns:
//   - map[string][][]*t.MarketInfo: Map of collateral to list of strategies
func GetAllStrats(pcs []*types.ProtocolChain, maxLevels int) map[string][][]*types.MarketInfo {
	// Mapping of base token to list of strategies
	result := make(map[string][][]*types.MarketInfo)
	// BFS
	for level := 0; level < maxLevels; level++ {
		for _, pc := range pcs {
			// If first level, only supply tokens
			if level == 0 {
				for _, supplyMarket := range pc.SupplyMarkets {
					symbol := utils.CommonSymbol(supplyMarket.Token)
					result[symbol] = append(result[symbol], []*types.MarketInfo{supplyMarket})
				}
			} else {
				// Add to existing strategies
				for _, supplyMarket := range pc.SupplyMarkets {
					supplySymbol := utils.CommonSymbol(supplyMarket.Token)
					for _, borrowMarket := range pc.BorrowMarkets {
						// Skip if same token as supply
						if supplyMarket.Token == borrowMarket.Token {
							continue
						}
						borrowSymbol := utils.CommonSymbol(borrowMarket.Token)
						for _, strategy := range result[borrowSymbol] {
							// Skip if strategy already has level amount
							if len(strategy) == 2*level+1 {
								continue
							}
							// Skip if same pc (borrowing and lending same token from same protocol)
							if supplyMarket.Protocol == strategy[0].Protocol && supplyMarket.Chain == strategy[0].Chain {
								continue
							}
							// Add to the front of existing strategy
							newStrat := make([]*types.MarketInfo, len(strategy)+2)
							newStrat[0] = supplyMarket
							newStrat[1] = borrowMarket
							copy(newStrat[2:], strategy)
							result[supplySymbol] = append(result[supplySymbol], newStrat)
						}
					}
				}
			}
		}
	}
	return result
}

// Calculates the total APY for the strategy, initial liquidity, and safety factor.
// Attempts to provide max liquidity at each level.
//
// Params:
//   - ps: Map of protocol name to protocol
//   - strat: Strategy to calculate
//   - initialUSD: Initial liquidity in USD, with 8 decimals
//   - safety: Safety factor in basis points
//
// Returns:
//   - *big.Int: Total APY in ray
//   - []*t.StrategyStep: List of steps to take
//   - error: Error if any
func CalcStratSteps(ps map[string]*protocols.Protocol, strat []*types.MarketInfo, initialUSD, safety *big.Int) (*big.Int, []*types.StrategyStep, error) {
	if len(strat) == 0 { // Base case
		return big.NewInt(0), nil, nil
	}
	p, ok := ps[strat[0].Protocol]
	if !ok {
		return big.NewInt(0), nil, fmt.Errorf("protocol not found: %v", strat[0].Protocol)
	}

	var totalAPY *big.Int
	var nextLevelAPY *big.Int
	var nextSteps []*types.StrategyStep

	isSupply := len(strat)%2 == 1
	market := strat[0]
	decimals := new(big.Int).Exp(big.NewInt(10), market.Decimals, nil)
	initialUSDScaled := new(big.Int).Mul(initialUSD, decimals) // Add decimals
	if market.PriceInUSD.Cmp(big.NewInt(0)) == 0 {
		b, _ := json.MarshalIndent(market, "", "  ")
		return big.NewInt(0), nil, fmt.Errorf("price is 0: %v", string(b)) // TODO: Remove this
	}
	initialAmount := new(big.Int).Div(initialUSDScaled, market.PriceInUSD) // Remove 8 decimals

	currentAPY, currentAmount, err := (*p).CalcAPY(market, initialAmount, isSupply)
	if err != nil {
		return big.NewInt(0), nil, fmt.Errorf("failed to calc apy: %v", err)
	}
	currentAmountUSD := new(big.Int).Mul(currentAmount, market.PriceInUSD) // Add 8 decimals
	currentAmountUSD.Div(currentAmountUSD, decimals)                       // Remove decimals

	if isSupply { // Lend
		// Adjust for LTV and safety factor
		ltv := utils.BasisMul(market.LTV, safety)
		currentAmountUSD = utils.BasisMul(currentAmountUSD, ltv) // Percentage

		// Next level
		nextLevelAPY, nextSteps, err = CalcStratSteps(ps, strat[1:], currentAmountUSD, safety)
		if err != nil {
			return big.NewInt(0), nil, fmt.Errorf("failed to calc next level apy: %v", err)
		}
		// Calculating manually instead of using ltv to account for lower amount in next level
		var nextLevelAPYAdjusted *big.Int
		if nextSteps != nil {
			nextLevelAPYAdjusted = new(big.Int).Mul(nextLevelAPY, nextSteps[0].Amount)
			nextLevelAPYAdjusted.Div(nextLevelAPYAdjusted, currentAmount)
		} else {
			nextLevelAPYAdjusted = nextLevelAPY // Should be 0
		}

		totalAPY = new(big.Int).Add(currentAPY, nextLevelAPYAdjusted)
	} else { // Borrow
		// Next level
		nextLevelAPY, nextSteps, err = CalcStratSteps(ps, strat[1:], currentAmountUSD, safety)
		if err != nil {
			return big.NewInt(0), nil, fmt.Errorf("failed to calc next level apy: %v", err)
		}

		// If next level has lower amount, recalculate current level
		// TODO: Add a tolerance threshold
		if nextSteps != nil {
			nextStep := nextSteps[0]
			nextStepDecimals := new(big.Int).Exp(big.NewInt(10), nextStep.Market.Decimals, nil)
			nextStepAmountUSD := new(big.Int).Mul(nextStep.Amount, nextStep.Market.PriceInUSD) // Add 8 decimals
			nextStepAmountUSD.Div(nextStepAmountUSD, nextStepDecimals)
			if nextStepAmountUSD.Cmp(currentAmountUSD) == -1 {
				log.Print("Lower amount, diff: ", new(big.Int).Sub(currentAmountUSD, nextStepAmountUSD))
				newBorrowUSDScaled := new(big.Int).Mul(nextStepAmountUSD, decimals)
				currentAmount = newBorrowUSDScaled.Div(newBorrowUSDScaled, market.PriceInUSD)
				currentAPY, _, err = (*p).CalcAPY(market, currentAmount, false)
				if err != nil {
					return big.NewInt(0), nil, fmt.Errorf("failed to calc borrow apy: %v", err)
				}
			}
		}

		totalAPY = new(big.Int).Sub(nextLevelAPY, currentAPY)
	}
	totalSteps := make([]*types.StrategyStep, len(nextSteps)+1)
	totalSteps[0] = &types.StrategyStep{
		Market:   market,
		IsSupply: isSupply,
		APY:      currentAPY,
		Amount:   currentAmount,
	}
	copy(totalSteps[1:], nextSteps)
	return totalAPY, totalSteps, nil
}

// Calculates the strategies' steps and total APYs.
//
// Params:
//   - ps: Map of protocol name to protocol
//   - strats: List of strategies to calculate
//   - initialUSD: Initial liquidity in USD, with 8 decimals
//   - safety: Safety factor in basis points
//
// Returns:
//   - []*t.Strategy: List of strategies with steps and total APYs
//   - error: Error if any
func CalcStrategies(ps map[string]*protocols.Protocol, strats [][]*types.MarketInfo, initialUSD, safety *big.Int) ([]*types.Strategy, error) {
	result := make([]*types.Strategy, len(strats))
	for i, strat := range strats {
		totalAPY, steps, err := CalcStratSteps(ps, strat, initialUSD, safety)
		if err != nil {
			return nil, fmt.Errorf("failed to calc strat steps: %v", err)
		}
		result[i] = &types.Strategy{
			Steps: steps,
			APY:   totalAPY,
		}
	}
	return result, nil
}

// Compares the two strategies' APYs.
// Returns true if a is larger than b, false otherwise.
func apyMore(a, b *types.Strategy) bool {
	return a.APY.Cmp(b.APY) == 1
}

// Sorts the strategies by APY in descending order.
func SortStrategies(strats []*types.Strategy) {
	slices.SortFunc[*types.Strategy](strats, apyMore)
}
