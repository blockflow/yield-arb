package api

import (
	"log"
	"math/big"
	"net/http"
	"strings"
	"yield-arb/cmd/arbitrage"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/types"
	"yield-arb/cmd/utils"

	"github.com/go-chi/chi"
)

func getStrats(w http.ResponseWriter, r *http.Request) {
	approvedCollateralTokens := strings.Split(chi.URLParam(r, "approvedCollateralTokens"), ",")
	chains := strings.Split(chi.URLParam(r, "chains"), ",")
	ps := strings.Split(chi.URLParam(r, "protocols"), ",")
	initialAmountUSD, ok := new(big.Int).SetString(chi.URLParam(r, "initialAmountUSD"), 10)
	if !ok {
		log.Panicf("failed to parse initialAmountUSD: %v", chi.URLParam(r, "initialAmountUSD"))
	}
	safety, ok := new(big.Int).SetString(chi.URLParam(r, "safetyFactor"), 10)
	if !ok {
		log.Panicf("failed to parse safety: %v", chi.URLParam(r, "safetyFactor"))
	}

	var pcs []*types.ProtocolChain
	psMap := make(map[string]*protocols.Protocol)
	for _, protocol := range ps {
		p, err := protocols.GetProtocol(protocol)
		if err != nil {
			log.Panicf("Failed to get protocol: %v", err)
		}
		psMap[protocol] = &p
		for _, chain := range chains {
			p.Connect(chain)
			pms, err := p.GetMarkets()
			if err != nil {
				log.Panicf("failed to get markets: %v", err)
			}
			pcs = append(pcs, pms...)
		}
	}

	log.Println("Calculating all strats...")
	collateralStrats := arbitrage.GetAllStrats(pcs, 2)
	for _, collateral := range approvedCollateralTokens {
		log.Println("----------------------------------------")
		log.Println(collateral)
		for _, strats := range collateralStrats[collateral] {
			log.Println("Strat:")
			for _, market := range strats {
				log.Println(market.Protocol, " ", market.Token, " ", market.Decimals, " ", market.LTV, " ", market.PriceInUSD)
			}
		}
	}

	log.Println("Generating all steps...")
	strats := make([][]*types.MarketInfo, 1)
	for _, collateral := range approvedCollateralTokens {
		collStrats, ok := collateralStrats[collateral]
		if ok {
			strats = append(strats, collStrats...)
		}
	}
	strategies, err := arbitrage.CalcStrategies(psMap, strats, initialAmountUSD, safety)
	if err != nil {
		log.Panicf("failed to calc strategies: %v", err)
	}
	arbitrage.SortStrategies(strategies)
	for _, strat := range strategies[:5] {
		log.Println("----------------------------------------")
		log.Printf("APY: %v", utils.ConvertRayToPercentage(strat.APY))
		for _, step := range strat.Steps {
			log.Println(step.Market.Protocol, step.Market.Token, step.IsSupply, utils.ConvertRayToPercentage(step.APY), step.Amount)
		}
	}
}
