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

	p, _ := protocols.GetProtocol("aavev3")

	// p.Connect("ethereum")

	// symbols, _ := p.GetLendingTokens()

	// specs, _ := p.GetLendingSpecs(symbols)
	// log.Println(specs)

	// lendingAPYs, _ := p.GetLendingAPYs(symbols)
	// log.Println(lendingAPYs)

	chains := []string{"ethereum_goerli", "avalanche_fuji", "polygon_mumbai"}
	chainPMs := make([]*protocols.ProtocolMarkets, len(chains))
	for i, chain := range chains {
		p.Connect(chain)
		pms, _ := p.GetMarkets()
		chainPMs[i] = &pms
	}

	strats, _ := arbitrage.CalculateStrategies(chainPMs)
	for _, strat := range strats[:5] {
		log.Println()
		for _, spec := range strat {
			log.Print(*spec)
		}
	}

}
