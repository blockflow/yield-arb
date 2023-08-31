package schema

import (
	"math/big"
	"net/http"
)

type ProtocolChain struct {
	Protocol      string        `json:"protocol"`
	Chain         string        `json:"chain"`
	SupplyMarkets []*MarketInfo `json:"lendingSpecs"`
	BorrowMarkets []*MarketInfo `json:"borrowingSpecs"`
}

type MarketInfo struct {
	Protocol   string      `json:"protocol"`
	Chain      string      `json:"chain"`
	Token      string      `json:"token"`
	Decimals   *big.Int    `json:"decimals"`
	LTV        *big.Int    `json:"ltv"`        // In basis points, 0 if cannot be collateral
	PriceInUSD *big.Int    `json:"priceInUSD"` // How much USD is required to purchase 1 ether unit, with 8 decimals
	Params     interface{} `json:"params"`     // State of the market, e.g. total supplied, total borrowed, etc. Cannot be a pointer.
}

type AccountData struct {
	TotalCollateralBase         *big.Int `json:"totalCollateralBase"`
	TotalDebtBase               *big.Int `json:"totalDebtBase"`
	AvailableBorrowsBase        *big.Int `json:"availableBorrowsBase"`
	CurrentLiquidationThreshold *big.Int `json:"currentLiquidationThreshold"`
	LTV                         *big.Int `json:"ltv"`
	HealthFactor                *big.Int `json:"healthFactor"`
}

type StrategyStep struct {
	Market   *MarketInfo `json:"market"`
	IsSupply bool        `json:"isSupply"`
	APY      *big.Int    `json:"apy"`
	Amount   *big.Int    `json:"amount"`
}

type Strategy struct {
	Steps []*StrategyStep `json:"steps"`
	APY   *big.Int        `json:"apy"`
}

func (s *Strategy) Bind(r *http.Request) error {
	return nil
}
