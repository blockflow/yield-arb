package main

import (
	"log"
	"time"
	p "yield-arb/cmd/protocols"
)

func main() {
	log.Println("Starting bot...")
	startTime := time.Now()

	// chains := []string{"arbitrum"}
	// var chainPMs []*types.ProtocolChain
	// ps := []string{"dforce"}
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

	const protocol = "dforce"
	const chain = "arbitrum"
	const wallet = "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	// Test Deposit()
	// p, _ := p.GetProtocol(protocol)
	// p.Connect(chain)
	// _, err := p.Supply(wallet, "iDAI", big.NewInt(1000000))
	// if err != nil {
	// 	log.Printf("failed to supply: %v", err)
	// }

	// Test Borrow
	// p, _ := p.GetProtocol(protocol)
	// p.Connect(chain)
	// _, err := p.Borrow(wallet, "iUSDC", big.NewInt(1000000))
	// if err != nil {
	// 	log.Printf("failed to borrow: %v", err)
	// }

	// Test Repay
	// p, _ := p.GetProtocol(protocol)
	// p.Connect(chain)
	// _, err := p.RepayAll(wallet, "iUSDC")
	// if err != nil {
	// 	log.Printf("failed to repay: %v", err)
	// }

	// Test Withdraw
	p, _ := p.GetProtocol(protocol)
	p.Connect(chain)
	_, err := p.WithdrawAll(wallet, "iDAI")
	if err != nil {
		log.Printf("failed to withdraw: %v", err)
	}

	log.Printf("Total time elapsed: %v", time.Since(startTime))
}
