package main

import (
	"log"
	"yield-arb/cmd/protocols"
)

func main() {
	log.Println("Starting bot...")
	// p, err := protocols.GetProtocol("aavev2")
	// if err != nil {
	// 	panic(err)
	// }

	p, _ := protocols.GetProtocol("aavev3")

	p.Connect("ethereum")

	symbols, _ := p.GetLendingTokens()

	specs, _ := p.GetLendingSpecs(symbols)
	log.Println(specs)

	// lendingAPYs, _ := p.GetLendingAPYs(symbols)
	// log.Println(lendingAPYs)

	// ethereumPMs, _ := p.GetMarkets()
	// p.Connect("polygon")
	// polygonPMs, _ := p.GetMarkets()
	// p.Connect("avalanche")
	// avalanchePMs, _ := p.GetMarkets()

	// strats, _ := arbitrage.CalculateStrategies([]*protocols.ProtocolMarkets{
	// 	&ethereumPMs, &polygonPMs, &avalanchePMs,
	// })
	// for _, strat := range strats[:5] {
	// 	log.Println()
	// 	for _, spec := range strat {
	// 		log.Print(*spec)
	// 	}
	// }

}
