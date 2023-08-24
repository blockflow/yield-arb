package arbitrage

var ApprovedCollateralTokens = []string{"USDC", "USDT", "ETH", "stETH"}

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

// Finds all the strategies that can be made with the given tokens using BFS.
// Returns a map of token symbol to list of strategies.
//
// Params:
//   - pcs: List of protocol chains
//   - baseTokens: List of base tokens to start with
// func GetAllStrats(pcs []*t.ProtocolChain, baseTokens []string) map[string][]*t.MarketInfo {

// }
