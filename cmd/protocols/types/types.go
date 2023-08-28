package types

import "math/big"

type ProtocolChain struct {
	Protocol      string        `json:"protocol"`
	Chain         string        `json:"chain"`
	SupplyMarkets []*MarketInfo `json:"lendingSpecs"`
	BorrowMarkets []*MarketInfo `json:"borrowingSpecs"`
}

type MarketInfo struct {
	Protocol   string `json:"protocol"`
	Chain      string `json:"chain"`
	Token      string `json:"token"`
	Decimals   *big.Int
	LTV        *big.Int    // In basis points, 0 if cannot be collateral
	PriceInUSD *big.Int    // How much USD is required to purchase 1 ether unit, with 8 decimals
	Params     interface{} // State of the market, e.g. total supplied, total borrowed, etc. Cannot be a pointer.
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
	Market   *MarketInfo
	IsSupply bool
	APY      *big.Int
	Amount   *big.Int
}
