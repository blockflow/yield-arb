package main

import (
	"log"
	"yield-arb/cmd/protocols"
)

func main() {
	log.Println("Starting bot...")
	p, err := protocols.GetProtocol("aavev2")
	if err != nil {
		panic(err)
	}
	p.Connect("ethereum")

	symbols, _ := p.GetLendingTokens()
	log.Println(symbols)

	lendingAPYs, _ := p.GetLendingAPYs(symbols)
	log.Println(lendingAPYs)
}
