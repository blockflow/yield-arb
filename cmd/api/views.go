package api

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"yield-arb/cmd/arbitrage"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/types"

	"github.com/go-chi/render"
)

type GetStratsInput struct {
	BaseTokens       []string `json:"baseTokens"`
	Chains           []string `json:"chains"`
	Protocols        []string `json:"protocols"`
	InitialAmountUSD *big.Int `json:"initialAmountUSD"`
	SafetyFactor     *big.Int `json:"safetyFactor"`
}

func (i *GetStratsInput) Bind(r *http.Request) error {
	return nil
}

func getStrats(w http.ResponseWriter, r *http.Request) {
	var input GetStratsInput
	if err := render.Bind(r, &input); err != nil {
		log.Panicf("failed to bind input: %v", err)
	}

	var pcs []*types.ProtocolChain
	psMap := make(map[string]*protocols.Protocol)
	for _, protocol := range input.Protocols {
		p, err := protocols.GetProtocol(protocol)
		if err != nil {
			log.Panicf("Failed to get protocol: %v", err)
		}
		psMap[protocol] = &p
		for _, chain := range input.Chains {
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

	log.Println("Generating all steps...")
	strats := make([][]*types.MarketInfo, 1)
	for _, baseToken := range input.BaseTokens {
		baseStrats, ok := collateralStrats[baseToken]
		if ok {
			strats = append(strats, baseStrats...)
		}
	}
	strategies, err := arbitrage.CalcStrategies(psMap, strats, input.InitialAmountUSD, input.SafetyFactor)
	if err != nil {
		log.Panicf("failed to calc strategies: %v", err)
	}
	arbitrage.SortStrategies(strategies)

	res, err := json.Marshal(strategies)
	if err != nil {
		log.Panicf("failed to marshal strategies: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
