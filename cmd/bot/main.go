package main

import (
	"log"
	"math/big"
	"time"
	p "yield-arb/cmd/protocols"
)

func main() {
	log.Println("Starting bot...")
	startTime := time.Now()

	// chains := []string{"arbitrum"}
	// var chainPMs []*t.ProtocolChain
	// ps := []string{"aavev3", "compoundv3", "dforce", "lodestar"}
	// for _, protocol := range ps {
	// 	p, err := p.GetProtocol(protocol)
	// 	if err != nil {
	// 		log.Panicf("Failed to get protocol: %v", err)
	// 	}
	// 	for _, chain := range chains {
	// 		p.Connect(chain)
	// 		pms, err := p.GetMarkets()
	// 		if err != nil {
	// 			log.Panicf("failed to get markets: %v", err)
	// 		}
	// 		chainPMs = append(chainPMs, pms)
	// 	}
	// }

	// log.Println("Calculating strategy v2...")
	// stratsV2, _ := arbitrage.CalculateStrategiesV2(chainPMs)
	// caps := arbitrage.CalculateStratV2CapsUSD(stratsV2)
	// for collateral, specs := range stratsV2 {
	// 	log.Printf("%v: %v", collateral, arbitrage.CalculateNetAPYV2(specs))
	// 	log.Printf("Cap in USD: $%v", caps[collateral])
	// 	for _, spec := range specs {
	// 		log.Print(spec.Protocol, " ", spec.Token, " ", spec.SupplyAPY, " ", spec.BorrowAPY)
	// 		// prettySpec, _ := json.MarshalIndent(spec, "", "  ")
	// 		// log.Print(string(prettySpec))
	// 	}
	// }

	// Test Deposit()
	p, _ := p.GetProtocol("aavev3")
	p.Connect("ethereum_goerli")
	_, err := p.Supply("0x18dC22D776aEFefD2538079409176086fcB6C741", "WETH", big.NewInt(100000000000000))
	if err != nil {
		log.Printf("failed to supply: %v", err)
	}

	// Test Borrow
	// p, _ := p.GetProtocol("aavev3")
	// p.Connect("ethereum_goerli")
	// _, err := p.Borrow(common.HexToAddress("0x18dC22D776aEFefD2538079409176086fcB6C741"), "USDC", big.NewInt(31))
	// if err != nil {
	// 	log.Printf("failed to borrow: %v", err)
	// }

	log.Printf("Total time elapsed: %v", time.Since(startTime))
}
