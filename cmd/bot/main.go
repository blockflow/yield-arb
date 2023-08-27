package main

import (
	"log"
	"time"
	"yield-arb/cmd/arbitrage"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/types"
)

func main() {
	log.Println("Starting bot...")
	startTime := time.Now()

	const protocol = "aavev3"
	const chain = "arbitrum"
	const wallet = "0x18dC22D776aEFefD2538079409176086fcB6C741"
	var ApprovedCollateralTokens = []string{"USDC"}
	// var ApprovedCollateralTokens = []string{"USDC", "USDT", "ETH", "stETH"}

	chains := []string{"arbitrum"}
	var pcs []*types.ProtocolChain
	// ps := []string{"compoundv3"}
	ps := []string{"aavev3", "dforce"}
	// ps := []string{"aavev3", "compoundv3", "dforce", "lodestar"}
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
			pcs = append(pcs, pms)
		}

		// apy, amount, err := p.CalcAPY(pcs[0].SupplyMarkets[4], new(big.Int).Exp(big.NewInt(0), big.NewInt(24), nil), true)
		// if err != nil {
		// 	log.Panicf("failed to calc apy: %v", err)
		// }
		// log.Printf("Token: %v, APY: %v, Amount: %v", pcs[0].SupplyMarkets[4].Token, apy, amount)
	}

	log.Println("Calculating all strats...")
	collateralStrats := arbitrage.GetAllStrats(pcs, 2)
	// caps := arbitrage.CalculateStratV2CapsUSD(stratsV2)
	for _, collateral := range ApprovedCollateralTokens {
		log.Println("----------------------------------------")
		log.Println(collateral)
		for _, strats := range collateralStrats[collateral] {
			log.Println("Strat:")
			for _, market := range strats {
				log.Println(market.Protocol, " ", market.Token, " ", market.Decimals, " ", market.LTV, " ", market.PriceInUSD)
			}
		}
	}

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
