package main

import (
	"log"
	"yield-arb/cmd/arbitrage"
	"yield-arb/cmd/protocols"
)

func main() {
	log.Println("Starting bot...")
	// p, err := protocols.GetProtocol("aavev2")
	// if err != nil {
	// 	panic(err)
	// }

	// p, _ := protocols.GetProtocol("aavev2")

	// p.Connect("ethereum")

	// symbols, _ := p.GetLendingTokens()

	// specs, _ := p.GetLendingSpecs(symbols)
	// log.Println(specs)

	// lendingAPYs, _ := p.GetLendingAPYs(symbols)
	// log.Println(lendingAPYs)

	chains := []string{"ethereum", "polygon", "avalanche"}
	// chains := []string{"ethereum_goerli", "avalanche_fuji", "polygon_mumbai"}
	chainPMs := make([]*protocols.ProtocolMarkets, len(chains))
	ps := []string{"aavev3"}
	for _, protocol := range ps {
		p, _ := protocols.GetProtocol(protocol)
		for i, chain := range chains {
			p.Connect(chain)
			pms, _ := p.GetMarkets()
			chainPMs[i] = &pms
		}
	}

	strats, _ := arbitrage.CalculateStrategies(chainPMs)
	var topStrats = make([][]*protocols.TokenSpecs, 5)
	copy(topStrats, strats)
	for _, strat := range topStrats {
		log.Println()
		for _, spec := range strat {
			log.Print(*spec)
		}
	}

	stratsV2, _ := arbitrage.CalculateStrategiesV2(chainPMs)
	for collateral, specs := range stratsV2 {
		log.Print(collateral)
		for _, spec := range specs {
			log.Print(*spec)
		}
	}

	// Test Deposit()
	// p.Connect("ethereum_goerli")
	// p.Deposit("0x18dC22D776aEFefD2538079409176086fcB6C741", "WETH", big.NewInt(1))
}
