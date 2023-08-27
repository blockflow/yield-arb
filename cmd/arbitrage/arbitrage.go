package arbitrage

import (
	"fmt"
	"math/big"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/types"
	"yield-arb/cmd/utils"
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
func CalcStratAPY(ps map[string]*protocols.Protocol, strat []*types.MarketInfo, initial, safety *big.Int) (*big.Int, *big.Int, error) {
	if len(strat) == 0 { // Base case
		return big.NewInt(0), big.NewInt(0)
	}

	apy := big.NewInt(0)
	var amount *big.Int
	*amount = *initial     // Copy value to avoid mutating
	if len(strat)%2 == 0 { // Lend
		p := ps[strat[0].Protocol]
		supplyAPY, supplyAmount, err := (*p).CalcAPY(strat[0], amount, true)
		if err != nil {
			return big.NewInt(0), big.NewInt(0), fmt.Errorf("failed to calc supply apy: %v", err)
		}
		borrowAPY, borrowAmount, err := (*p).CalcAPY(strat[1], amount, true)
		if err != nil {
			return big.NewInt(0), big.NewInt(0), fmt.Errorf("failed to calc borrow apy: %v", err)
		}

		assetNetAPY := new(big.Float).Sub(strat[1].SupplyAPY, strat[0].BorrowAPY)
		// If borrowing again
		nextLevelAPY := CalculateNetAPYV2(strat[2:])
		ltv := new(big.Float).Quo(strat[0].LTV, big.NewFloat(100))
		nextLevelAPY.Mul(nextLevelAPY, ltv)
		return assetNetAPY.Add(assetNetAPY, nextLevelAPY)
	} else { // Borrow
		apy.Sub(apy, market.BorrowAPY)
	}
}
