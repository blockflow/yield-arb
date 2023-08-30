package api

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"yield-arb/cmd/arbitrage"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/compoundv3"
	"yield-arb/cmd/protocols/schema"

	"github.com/go-chi/render"
)

type GetStratsInput struct {
	BaseTokens       []string `json:"baseTokens"`
	Chains           []string `json:"chains"`
	Protocols        []string `json:"protocols"`
	MaxLevels        int      `json:"maxLevels"`
	InitialAmountUSD *big.Int `json:"initialAmountUSD"`
	SafetyFactor     *big.Int `json:"safetyFactor"`
}

func (i *GetStratsInput) Bind(r *http.Request) error {
	return nil
}

func test(w http.ResponseWriter, r *http.Request) {
	const protocol = "compoundv3"
	const chain = "arbitrum"
	const wallet = "0x18dC22D776aEFefD2538079409176086fcB6C741"
	const token = "USDC"
	// Test Supply()
	p, _ := protocols.GetProtocol(protocol)
	p.Connect(chain)
	txs, err := p.GetTransactions(wallet, &schema.StrategyStep{
		Market: &schema.MarketInfo{
			Protocol: protocol,
			Chain:    chain,
			Token:    token,
			Params:   &compoundv3.CompoundV3Params{},
		},
		IsSupply: true,
		Amount:   big.NewInt(1e6),
	})
	if err != nil {
		log.Printf("failed to supply: %v", err)
	}
	res, err := json.Marshal(txs)
	if err != nil {
		log.Panicf("failed to marshal strategies: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}

func getStrats(w http.ResponseWriter, r *http.Request) {
	var input GetStratsInput
	if err := render.Bind(r, &input); err != nil {
		log.Panicf("failed to bind input: %v", err)
	}

	var pcs []*schema.ProtocolChain
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
	collateralStrats := arbitrage.GetAllStrats(pcs, input.MaxLevels)

	log.Println("Generating all steps...")
	strats := make([][]*schema.MarketInfo, 1)
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
