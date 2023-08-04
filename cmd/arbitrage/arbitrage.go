package arbitrage

import (
	"log"
	"math/big"
	p "yield-arb/cmd/protocols"

	"golang.org/x/exp/slices"
)

var ApprovedCollateralTokens = []string{"USDC", "USDT", "WETH", "stETH"}

// Calculates the net APY for the tokenspec triplets
func calculateNetAPY(specs []*p.TokenSpecs) *big.Float {
	loanAPY := new(big.Float).Sub(specs[2].APY, specs[1].APY)
	loanAPY.Mul(loanAPY, specs[0].LTV)
	loanAPY.Quo(loanAPY, big.NewFloat(100))
	netAPY := new(big.Float).Add(specs[0].APY, loanAPY)
	return netAPY
}

// Compares the two tokenspec triplets' net APYs.
// Returns true if a is larger than b, false otherwise.
func moreNetAPY(a, b []*p.TokenSpecs) bool {
	return calculateNetAPY(a).Cmp(calculateNetAPY((b))) == 1
}

// Calculates the strategies and ranks them
func CalculateStrategies(pms []*p.ProtocolMarkets) ([][]*p.TokenSpecs, error) {
	log.Println("Calculating the best yield arb strategies...")
	// Find best lend rates per token
	maxTokenLendSpecs := make(map[string]*p.TokenSpecs)
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
	collateralSpecs := make(map[string][]*p.TokenSpecs)
	for _, pm := range pms {
		market := pm.Protocol + ":" + pm.Chain
		for _, spec := range pm.LendingSpecs {
			if slices.Contains(ApprovedCollateralTokens, spec.Token) {
				collateralSpecs[market] = append(collateralSpecs[market], spec)
			}
		}
	}

	// Pair borrow rates with best collateral rates within markets
	var tokenPairs [][]*p.TokenSpecs
	for _, pm := range pms {
		market := pm.Protocol + ":" + pm.Chain
		// Iterate over BorrowingSpecs
		for _, borrowSpec := range pm.BorrowingSpecs {
			// Iterate over collateral specs
			for _, collateralSpec := range collateralSpecs[market] {
				tokenPairs = append(tokenPairs, []*p.TokenSpecs{
					collateralSpec,
					borrowSpec,
				})
			}
		}
	}

	// Find best strat per pair
	bestStrats := make([][]*p.TokenSpecs, len(tokenPairs))
	for i, pair := range tokenPairs {
		bestStrats[i] = []*p.TokenSpecs{
			pair[0],
			pair[1],
			maxTokenLendSpecs[pair[1].Token],
		}
	}

	// Sort strats in descending order
	slices.SortFunc(bestStrats, moreNetAPY)

	return bestStrats, nil
}
