package arbitrage

import (
	"log"
	"math/big"
	t "yield-arb/cmd/protocols/types"
	"yield-arb/cmd/utils"

	"golang.org/x/exp/slices"
)

type ArbPathLvl struct {
	TotalAmount *big.Float // In USD
	Specs       []*t.TokenSpecs
}

var ApprovedCollateralTokens = []string{"USDC", "USDT", "ETH", "stETH"}

// Calculates the net APY for the tokenspec triplets
func calculateNetAPY(specs []*t.TokenSpecs) *big.Float {
	loanAPY := new(big.Float).Sub(specs[2].SupplyAPY, specs[1].BorrowAPY)
	loanAPY.Mul(loanAPY, specs[0].LTV)
	loanAPY.Quo(loanAPY, big.NewFloat(100))
	netAPY := new(big.Float).Add(specs[0].SupplyAPY, loanAPY)
	return netAPY
}

// Calculates the net APY for any odd number of TokenSpecs.
// TokenSpecs should be in alternating order of lend/borrow starting with lend.
func CalculateNetAPYV2(specs []*t.TokenSpecs) *big.Float {
	if len(specs) == 0 {
		// Base case
		return big.NewFloat(0)
	} else if len(specs)%2 == 0 {
		// Even, intermediate case
		assetNetAPY := new(big.Float).Sub(specs[1].SupplyAPY, specs[0].BorrowAPY)
		// If borrowing again
		nextLevelAPY := CalculateNetAPYV2(specs[2:])
		ltv := new(big.Float).Quo(specs[0].LTV, big.NewFloat(100))
		nextLevelAPY.Mul(nextLevelAPY, ltv)
		return assetNetAPY.Add(assetNetAPY, nextLevelAPY)
	} else {
		// Odd, start case
		nextLevelAPY := CalculateNetAPYV2(specs[1:])
		ltv := new(big.Float).Quo(specs[0].LTV, big.NewFloat(100))
		nextLevelAPY.Mul(nextLevelAPY, ltv)
		return new(big.Float).Add(specs[0].SupplyAPY, nextLevelAPY)
	}
}

// Compares the two tokenspec triplets' net APYs.
// Returns true if a is larger than b, false otherwise.
func moreNetAPY(a, b []*t.TokenSpecs) bool {
	return calculateNetAPY(a).Cmp(calculateNetAPY((b))) == 1
}

// Returns true if a's Supply APY is larger than b's, false otherwise.
func moreSupplyAPY(a, b *t.TokenSpecs) bool {
	return a.SupplyAPY.Cmp(b.SupplyAPY) == 1
}

// Calculates the weighted avg apy for a list of TokenSpecs.
// Alternates between supply/borrow.
func weightedAvgAPY(specs []*t.TokenSpecs) *big.Float {
	var result *big.Float
	var totalWeight *big.Float
	for i, spec := range specs {
		if i%2 == 1 { // If odd, then supply
			result.Add(result, new(big.Float).Mul(spec.SupplyCap, spec.SupplyAPY))
			totalWeight.Add(totalWeight, spec.SupplyCap)
		} else { // Even, borrow
			result.Add(result, new(big.Float).Mul(spec.BorrowCap, spec.BorrowAPY))
			totalWeight.Add(totalWeight, spec.BorrowCap)
		}
	}
	return result.Quo(result, totalWeight)
}

/*
Calculates the best strategies with dynamic path lengths and ranks them.
Limit to max of 3 levels (lends) to reduce interest rate risk.
Seeks to maximize: xa + ra(-ya + xb + rb(xc - yb)).
Takes into account supply/borrow caps and will suggest multiple paths
to support specified amount of base token.

find max Xc paths
  - group by token
  - calculate max amount for each group
  - sort by apy

find max Xb paths
  - find all singles
  - find all pairs (with max Xc)
  - calculate max amount for each group
  - sort by apy

find max Xa paths
  - pair all Yas with Xb lists
  - calculate max amount for each Ya path
  - combine Yas until collateral amount is met
    -
*/
func CalculateStrategiesV3(pms []*t.ProtocolMarkets, collateralToken string, collateralAmount *big.Float) (map[string][]*ArbPathLvl, error) {
	// Solve 3rd level, max of lend for each asset
	// Group by token
	thirdLvls := make(map[string]*ArbPathLvl)
	for _, pm := range pms {
		for _, xc := range pm.LendingSpecs {
			xcSymbol := utils.CommonSymbol(xc.Token)
			lvl := thirdLvls[xcSymbol]
			lvl.Specs = append(lvl.Specs, xc)
			// Keep track of total amount
			amount := new(big.Float).Mul(xc.SupplyCap, xc.PriceInUSD)
			lvl.TotalAmount.Add(lvl.TotalAmount, amount)
		}
	}
	// Sort each group by APY
	for _, lvl := range thirdLvls {
		slices.SortFunc[*t.TokenSpecs](lvl.Specs, moreSupplyAPY)
	}

	// Solve 2nd level, max of 2 lends taking into account LTV.
	// 1st lend and borrow assets must be from same protocol.
	// Borrow and 2nd lend assets have to match.
	// Can make block recursive to support more levels.
	var secondLvls = make(map[string][]*ArbPathLvl)
	for _, pm := range pms {
		for _, xb := range pm.LendingSpecs {
			xbSymbol := utils.CommonSymbol(xb.Token)
			lvls := secondLvls[xbSymbol]

			// Add all singles
			lvls = append(lvls, &ArbPathLvl{
				Specs:       []*t.TokenSpecs{xb},
				TotalAmount: xb.SupplyCap,
			})

			// Add all pairs, including total amounts
			for _, yb := range pm.BorrowingSpecs {
				xcGroup := thirdLvls[yb.Token]
				// Total amount is min of (xb,yb,and xc)
				minAmountUSD := new(big.Float).Mul(xb.SupplyCap, xb.PriceInUSD)
				// Take xb's LTV into account
				ybAmountAdjustedUSD := new(big.Float).Quo(yb.BorrowCap, xb.LTV)
				ybAmountAdjustedUSD.Mul(ybAmountAdjustedUSD, yb.PriceInUSD)
				if ybAmountAdjustedUSD.Cmp(minAmountUSD) == 1 {
					minAmountUSD = ybAmountAdjustedUSD
				}
				// Add on xc group
				xcAmountAdjustedUSD := new(big.Float).Quo(xcGroup.TotalAmount, xb.LTV)
				xcAmountAdjustedUSD.Mul(xcAmountAdjustedUSD, xcGroup.Specs[0].PriceInUSD)
				if xcAmountAdjustedUSD.Cmp(minAmountUSD) == 1 {
					minAmountUSD = xcAmountAdjustedUSD
				}
				lvls = append(lvls, &ArbPathLvl{
					TotalAmount: minAmountUSD,
					Specs:       []*t.TokenSpecs{xb, yb, xcGroup.Specs},
				})
			}
		}
	}

	// TODO: Sort second lvls

	// Solve 3rd level, max of 3 lends taking into account LTV.
	// - pair all Yas with Xb lists
	// - calculate max amount for each Ya path
	// - combine Yas until collateral amount is met
	firstLvls := make(map[string][]*ArbPathLvl)
	for _, pm := range pms {
		for _, xa := range pm.LendingSpecs {
			xaSymbol := utils.CommonSymbol(xa.Token)
			// Check if approved collateral
			if !slices.Contains(ApprovedCollateralTokens, xaSymbol) {
				continue
			}

			lvls := firstLvls[xaSymbol]
			// Add singular
			lvls = append(lvls, &ArbPathLvl{
				TotalAmount: new(big.Float).Mul(xa.SupplyCap, xa.PriceInUSD),
				Specs:       []*t.TokenSpecs{xa},
			})
			// Add second lvls
			for _, ya := range pm.BorrowingSpecs {
				// Calculate xa + ra(maxXbPath - ya)
				yaSymbol := utils.CommonSymbol(ya.Token)
				maxXbPath, ok := secondLvls[yaSymbol]
				if !ok {
					// No additional path
					continue
				}
				log.Print(maxXbPath)
			}
		}
	}

	return firstLvls, nil
}

// Calculates the best strategies with dynamic path lengths and ranks them.
// Limit to max of 3 levels (lends) to reduce interest rate risk.
// Seeks to maximize: xa + ra(-ya + xb + rb(xc - yb))
func CalculateStrategiesV2(pms []*t.ProtocolMarkets) (map[string][]*t.TokenSpecs, error) {
	// Solve 3rd level, max of lend for each asset
	maxXcs := make(map[string]*t.TokenSpecs)
	for _, pm := range pms {
		for _, xc := range pm.LendingSpecs {
			xcSymbol := utils.CommonSymbol(xc.Token)
			maxXc, ok := maxXcs[xcSymbol]
			// If higher lend APY found
			if !ok || xc.APY.Cmp(maxXc.APY) == 1 {
				maxXcs[xcSymbol] = xc
			}
		}
	}

	// Solve 2nd level, max of 2 lends taking into account LTV.
	// 1st lend and borrow assets must be from same protocol.
	// Borrow and 2nd lend assets have to match.
	// TODO: cache net apys
	// Can make block recursive to support more levels.
	maxXbPaths := make(map[string][]*t.TokenSpecs)
	for _, pm := range pms {
		for _, xb := range pm.LendingSpecs {
			xbSymbol := utils.CommonSymbol(xb.Token)
			maxXbPath, ok := maxXbPaths[xbSymbol]
			// If first or singular lend is better
			if !ok || xb.APY.Cmp(CalculateNetAPYV2(maxXbPath)) == 1 {
				maxXbPaths[xbSymbol] = []*t.TokenSpecs{xb}
			}
			// Check 2 level APYs
			for _, yb := range pm.BorrowingSpecs {
				// Calculate xb + rb(xc - yb)
				ybSymbol := utils.CommonSymbol(yb.Token)
				maxXcPath := maxXcs[ybSymbol]
				nextLevelAPY := new(big.Float).Sub(maxXcPath.APY, yb.APY)
				rb := new(big.Float).Quo(xb.LTV, big.NewFloat(100))
				nextLevelAPY.Mul(nextLevelAPY, rb)
				xbAPY := new(big.Float).Add(xb.APY, nextLevelAPY)
				// If two levels is better
				if xbAPY.Cmp(CalculateNetAPYV2(maxXbPaths[xbSymbol])) == 1 {
					maxXbPaths[xbSymbol] = []*t.TokenSpecs{xb, yb, maxXcPath}
				}
			}
		}
	}

	// Solve 3rd level, max of 3 lends taking into account LTV.
	maxXaPaths := make(map[string][]*t.TokenSpecs)
	for _, pm := range pms {
		for _, xa := range pm.LendingSpecs {
			xaSymbol := utils.CommonSymbol(xa.Token)
			// Check if approved collateral
			if !slices.Contains(ApprovedCollateralTokens, xaSymbol) {
				continue
			}

			maxXaPath, ok := maxXaPaths[xaSymbol]
			// If first or singular lend is better
			if !ok || xa.APY.Cmp(CalculateNetAPYV2(maxXaPath)) == 1 {
				maxXaPaths[xaSymbol] = []*t.TokenSpecs{xa}
			}
			// Check 2 and 3 level APYs
			for _, ya := range pm.BorrowingSpecs {
				// Calculate xa + ra(maxXbPath - ya)
				yaSymbol := utils.CommonSymbol(ya.Token)
				maxXbPath, ok := maxXbPaths[yaSymbol]
				if !ok {
					// No additional path
					continue
				}
				maxXbPathAPY := CalculateNetAPYV2(maxXbPath)
				nextLevelAPY := new(big.Float).Sub(maxXbPathAPY, ya.APY)
				ra := new(big.Float).Quo(xa.LTV, big.NewFloat(100))
				nextLevelAPY.Mul(nextLevelAPY, ra)
				xaAPY := new(big.Float).Add(xa.APY, nextLevelAPY)
				// If better
				if xaAPY.Cmp(CalculateNetAPYV2(maxXaPaths[xaSymbol])) == 1 {
					maxXaPaths[xaSymbol] = append([]*t.TokenSpecs{xa, ya}, maxXbPath...)
				}
			}
		}
	}

	return maxXaPaths, nil
}

// Calculates the strategies and ranks them
func CalculateStrategies(pms []*t.ProtocolMarkets) ([][]*t.TokenSpecs, error) {
	log.Println("Calculating the best yield arb strategies...")
	// Find best lend rates per token
	maxTokenLendSpecs := make(map[string]*t.TokenSpecs)
	for _, pm := range pms {
		for _, spec := range pm.LendingSpecs {
			// If higher APY found
			maxToken, ok := maxTokenLendSpecs[spec.Token]
			if !ok || spec.APY.Cmp(maxToken.APY) == 1 {
				maxTokenLendSpecs[spec.Token] = spec
			}
		}
	}

	// Get all approved collateral specs for each market
	collateralSpecs := make(map[string][]*t.TokenSpecs)
	for _, pm := range pms {
		market := pm.Protocol + ":" + pm.Chain
		for _, spec := range pm.LendingSpecs {
			if slices.Contains(ApprovedCollateralTokens, spec.Token) {
				collateralSpecs[market] = append(collateralSpecs[market], spec)
			}
		}
	}

	// Pair borrow rates with best collateral rates within markets
	var tokenPairs [][]*t.TokenSpecs
	for _, pm := range pms {
		market := pm.Protocol + ":" + pm.Chain
		// Iterate over BorrowingSpecs
		for _, borrowSpec := range pm.BorrowingSpecs {
			// Iterate over collateral specs
			for _, collateralSpec := range collateralSpecs[market] {
				tokenPairs = append(tokenPairs, []*t.TokenSpecs{
					collateralSpec,
					borrowSpec,
				})
			}
		}
	}

	// Find best strat per pair
	bestStrats := make([][]*t.TokenSpecs, len(tokenPairs))
	for i, pair := range tokenPairs {
		bestStrats[i] = []*t.TokenSpecs{
			pair[0],
			pair[1],
			maxTokenLendSpecs[pair[1].Token],
		}
	}

	// Sort strats in descending order
	slices.SortFunc(bestStrats, moreNetAPY)

	return bestStrats, nil
}
