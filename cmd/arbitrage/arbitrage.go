package arbitrage

import (
	"log"
	"math/big"
	"yield-arb/cmd/protocols"

	"golang.org/x/exp/slices"
)

type TokenSpecs struct {
	Protocol string     `json:"protocol"`
	Chain    string     `json:"chain"`
	Token    string     `json:"token"`
	APY      *big.Float `json:"apy"`
}

var ApprovedCollateralTokens = []string{"USDC", "USDT", "WETH", "stETH"}

// Calculates the net APY for the tokenspec triplets
func calculateNetAPY(tokens []*TokenSpecs) *big.Float {
	netAPY := new(big.Float).Sub(tokens[0].APY, tokens[1].APY)
	netAPY.Add(netAPY, tokens[2].APY)
	return netAPY
}

// Compares the two tokenspec triplets' net APYs.
// Returns true if a is larger than b, false otherwise.
func moreNetAPY(a, b []*TokenSpecs) bool {
	return calculateNetAPY(a).Cmp(calculateNetAPY((b))) == 1
}

// Calculates the strategies and ranks them
func CalculateStrategies(pms []*protocols.ProtocolMarkets) ([][]*TokenSpecs, error) {
	log.Println("Calculating the best yield arb strategies...")
	// Find best lend rates per token
	maxTokenLendAPYs := make(map[string]*TokenSpecs)
	for _, pm := range pms {
		for token, apy := range pm.LendingAPYs {
			// If higher APY found
			maxToken, ok := maxTokenLendAPYs[token]
			if !ok || apy.Cmp(maxToken.APY) == 1 {
				maxTokenLendAPYs[token] = &TokenSpecs{
					Protocol: pm.Protocol,
					Chain:    pm.Chain,
					Token:    token,
					APY:      apy,
				}
			}
		}
	}

	// Find best approved collateral rates per market
	maxMarketLendAPYs := make(map[string]*TokenSpecs)
	for _, pm := range pms {
		market := pm.Protocol + ":" + pm.Chain
		for token, apy := range pm.LendingAPYs {
			// Make sure approved collateral
			if !slices.Contains(ApprovedCollateralTokens, token) {
				continue
			}
			// If higher APY found
			maxMarket, ok := maxMarketLendAPYs[market]
			if !ok || apy.Cmp(maxMarket.APY) == 1 {
				maxMarketLendAPYs[market] = &TokenSpecs{
					Protocol: pm.Protocol,
					Chain:    pm.Chain,
					Token:    token,
					APY:      apy,
				}
			}
		}
	}

	// Pair borrow rates with best collateral rates within markets
	var tokenPairs [][]*TokenSpecs
	for _, pm := range pms {
		market := pm.Protocol + ":" + pm.Chain
		for token, apy := range pm.BorrowingAPYs {
			tokenPairs = append(tokenPairs, []*TokenSpecs{
				maxMarketLendAPYs[market],
				{
					Protocol: pm.Protocol,
					Chain:    pm.Chain,
					Token:    token,
					APY:      apy,
				},
			})
		}
	}

	// Find best strat per pair
	bestStrats := make([][]*TokenSpecs, len(tokenPairs))
	for i, pair := range tokenPairs {
		// bleh := maxTokenLendAPYs[pair[1].Token]
		bestStrats[i] = []*TokenSpecs{
			pair[0],
			pair[1],
			maxTokenLendAPYs[pair[1].Token],
		}
	}

	// Sort strats in descending order
	slices.SortFunc(bestStrats, moreNetAPY)

	return bestStrats, nil
}
