package api

import (
	"encoding/json"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"yield-arb/cmd/arbitrage"
	"yield-arb/cmd/protocols"
	"yield-arb/cmd/protocols/schema"

	"github.com/ethereum/go-ethereum/core/types"
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
			Params:   map[string]interface{}{},
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
	ps := strings.Split(r.URL.Query().Get("protocols"), ",")
	chains := strings.Split(r.URL.Query().Get("chains"), ",")
	baseTokens := strings.Split(r.URL.Query().Get("baseTokens"), ",")
	maxLevels, _ := strconv.Atoi(r.URL.Query().Get("maxLevels"))
	initialAmountUSD, _ := new(big.Int).SetString(r.URL.Query().Get("initialAmountUSD"), 10)
	safetyFactor, _ := new(big.Int).SetString(r.URL.Query().Get("safetyFactor"), 10)

	var pcs []*schema.ProtocolChain
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
	collateralStrats := arbitrage.GetAllStrats(pcs, maxLevels)

	log.Println("Generating all steps...")
	strats := make([][]*schema.MarketInfo, 0)
	for _, baseToken := range baseTokens {
		baseStrats, ok := collateralStrats[baseToken]
		if ok {
			strats = append(strats, baseStrats...)
		}
	}
	strategies, err := arbitrage.CalcStrategies(psMap, strats, initialAmountUSD, safetyFactor)
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

func getTransactions(w http.ResponseWriter, r *http.Request) {
	var strats schema.Strategies
	if err := render.Bind(r, &strats); err != nil {
		log.Panicf("failed to bind input: %v", err)
	}

	txs := make([][]*types.Transaction, len(strats.Strategies))
	ps := make(map[string]protocols.Protocol)
	for i, strat := range strats.Strategies {
		for _, step := range strat.Steps {
			p, ok := ps[step.Market.Protocol]
			if !ok {
				var err error
				p, err = protocols.GetProtocol(step.Market.Protocol)
				if err != nil {
					log.Panicf("Failed to get protocol: %v", err)
				}
				ps[step.Market.Protocol] = p
			}
			p.Connect(step.Market.Chain)
			newTxs, err := p.GetTransactions("", step)
			if err != nil {
				log.Panicf("failed to get transactions: %v", err)
			}
			txs[i] = append(txs[i], newTxs...)
		}
	}

	res, err := json.Marshal(txs)
	if err != nil {
		log.Panicf("failed to marshal strategies: %v", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
