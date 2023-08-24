package main

import (
	"log"
	"math/big"
	"time"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/types"
)

func main() {
	log.Println("Starting bot...")
	startTime := time.Now()

	const protocol = "aavev3"
	const chain = "arbitrum"
	const wallet = "0x18dC22D776aEFefD2538079409176086fcB6C741"

	chains := []string{"arbitrum"}
	var chainPMs []*types.ProtocolChain
	ps := []string{"aavev3"}
	// ps := []string{"aavev3", "compoundv3", "dforce"}
	psMap := make(map[string]*protocols.Protocol)
	for _, protocol := range ps {
		p, err := protocols.GetProtocol(protocol)
		if err != nil {
			log.Panicf("Failed to get protocol: %v", err)
		}
		psMap[protocol+":"+"arbitrum"] = &p
		for _, chain := range chains {
			p.Connect(chain)
			pms, err := p.GetMarkets()
			if err != nil {
				log.Panicf("failed to get markets: %v", err)
			}
			chainPMs = append(chainPMs, pms)
		}

		apy, amount, err := p.CalcAPY(chainPMs[0].SupplyMarkets[0], big.NewInt(1), false)
		if err != nil {
			log.Panicf("failed to calc apy: %v", err)
		}
		log.Printf("Token: %v, APY: %v, Amount: %v", chainPMs[0].SupplyMarkets[0].Token, apy, amount)
	}

	// log.Println("Calculating strategy v2...")
	// stratsV2, _ := arbitrage.CalculateStrategiesV2(chainPMs)
	// caps := arbitrage.CalculateStratV2CapsUSD(stratsV2)
	// for collateral, specs := range stratsV2 {
	// 	log.Println("----------------------------------------")
	// 	log.Printf("%v: %v", collateral, arbitrage.CalculateNetAPYV2(specs))
	// 	log.Printf("Cap in USD: $%v", caps[collateral])
	// 	for _, spec := range specs {
	// 		// log.Print(spec.Protocol, " ", spec.Token, " ", spec.SupplyAPY, " ", spec.BorrowAPY)
	// 		prettySpec, _ := json.MarshalIndent(spec, "", "  ")
	// 		log.Print(string(prettySpec))
	// 	}
	// }

	// Enter strat
	// err := arbitrage.EnterStrategy(wallet, psMap, stratsV2["ETH"], big.NewFloat(1000000000000000), big.NewFloat(50))
	// if err != nil {
	// 	log.Printf("failed to enter strategy: %v", err)
	// }

	// Exit strat
	// err := arbitrage.ExitStrategy(wallet, psMap, stratsV2["ETH"])
	// if err != nil {
	// 	log.Printf("failed to exit strategy: %v", err)
	// }

	// Test Deposit()
	// p, _ := protocols.GetProtocol(protocol)
	// p.Connect(chain)
	// _, err := p.Supply(wallet, "USDC.e", big.NewInt(1000000))
	// if err != nil {
	// 	log.Printf("failed to supply: %v", err)
	// }

	// Test Borrow
	// p, _ := protocols.GetProtocol(protocol)
	// p.Connect(chain)
	// _, err := p.Borrow(wallet, "AAVE", big.NewInt(441039744165))
	// if err != nil {
	// 	log.Printf("failed to borrow: %v", err)
	// }

	// Test Repay
	// p, _ := protocols.GetProtocol(protocol)
	// p.Connect(chain)
	// _, err := p.RepayAll(wallet, "lwstETH")
	// if err != nil {
	// 	log.Printf("failed to repay: %v", err)
	// }

	// Test Withdraw
	// p, _ := protocols.GetProtocol(protocol)
	// p.Connect(chain)
	// _, err := p.WithdrawAll(wallet, "lUSDC.e")
	// if err != nil {
	// 	log.Printf("failed to withdraw: %v", err)
	// }

	log.Printf("Total time elapsed: %v", time.Since(startTime))
}
