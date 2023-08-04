package main

import (
	"log"
	"yield-arb/cmd/arbitrage"
	"yield-arb/cmd/protocols"
)

func main() {
	log.Println("Starting bot...")
	p, err := protocols.GetProtocol("aavev2")
	if err != nil {
		panic(err)
	}

	p.Connect("ethereum")

	// symbols, _ := p.GetLendingTokens()
	// log.Println(symbols)

	// lendingAPYs, _ := p.GetLendingAPYs(symbols)
	// log.Println(lendingAPYs)

	ethereumPMs, _ := p.GetMarkets()
	// p.Connect("polygon")
	// polygonPMs, _ := p.GetMarkets()
	p.Connect("avalanche")
	avalanchePMs, _ := p.GetMarkets()

	strats, _ := arbitrage.CalculateStrategies([]*protocols.ProtocolMarkets{
		&ethereumPMs, &avalanchePMs,
	})
	for _, strat := range strats[:5] {
		log.Println()
		for _, spec := range strat {
			log.Print(*spec)
		}
	}

}
