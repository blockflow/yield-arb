package main

import (
	"log"
	"time"
	"yield-arb/cmd/arbitrage"
	p "yield-arb/cmd/protocols"
	t "yield-arb/cmd/protocols/types"
)

func main() {
	log.Println("Starting bot...")
	startTime := time.Now()
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

	chains := []string{"arbitrum_goerli"}
	// chains := []string{"ethereum", "polygon", "avalanche"}
	// chains := []string{"ethereum_goerli", "avalanche_fuji", "polygon_mumbai"}
	var chainPMs []*t.ProtocolChain
	ps := []string{"aavev3", "compoundv3"}
	// ps := []string{"compoundv3", "aavev3"}
	for _, protocol := range ps {
		p, err := p.GetProtocol(protocol)
		if err != nil {
			log.Printf("Failed to get protocol: %v", err)
		}
		for _, chain := range chains {
			p.Connect(chain)
			pms, _ := p.GetMarkets()
			chainPMs = append(chainPMs, pms)
		}
	}

	// // strats, _ := arbitrage.CalculateStrategies(chainPMs)
	// // var topStrats = make([][]*t.TokenSpecs, 5)
	// // copy(topStrats, strats)
	// // for _, strat := range topStrats {
	// // 	log.Println()
	// // 	for _, spec := range strat {
	// // 		log.Print(*spec)
	// // 	}
	// // }

	stratsV2, _ := arbitrage.CalculateStrategiesV2(chainPMs)
	caps := arbitrage.CalculateStratV2CapsUSD(stratsV2)
	for collateral, specs := range stratsV2 {
		log.Printf("%v: %v", collateral, arbitrage.CalculateNetAPYV2(specs))
		log.Printf("Cap in USD: $%v", caps[collateral])
		for _, spec := range specs {
			log.Print(*spec)
		}
	}

	// Test Deposit()
	// p, _ := p.GetProtocol("aavev3")
	// p.Connect("ethereum_goerli")
	// _, err := p.Supply(common.HexToAddress("0x18dC22D776aEFefD2538079409176086fcB6C741"), "WETH", big.NewInt(1))
	// if err != nil {
	// 	log.Printf("failed to supply: %v", err)
	// }

	// Test Borrow
	// p, _ := p.GetProtocol("aavev3")
	// p.Connect("ethereum_goerli")
	// _, err := p.Borrow(common.HexToAddress("0x18dC22D776aEFefD2538079409176086fcB6C741"), "USDC", big.NewInt(31))
	// if err != nil {
	// 	log.Printf("failed to borrow: %v", err)
	// }

	log.Printf("Total time elapsed: %v", time.Since(startTime))
}
