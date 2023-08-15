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
	Decimals   uint8
	LTV        *big.Float `json:"ltv"` // 0 if cannot be collateral
	SupplyAPY  *big.Float `json:"supplyApy"`
	BorrowAPY  *big.Float `json:"borrowApy"`
	SupplyCap  *big.Float // In ether units, availability remaining
	BorrowCap  *big.Float // In ether units, availability remaining
	PriceInUSD *big.Float // How much USD is required to purchase 1 ether unit
}

type AccountData struct {
	TotalCollateralBase         *big.Int `json:"totalCollateralBase"`
	TotalDebtBase               *big.Int `json:"totalDebtBase"`
	AvailableBorrowsBase        *big.Int `json:"availableBorrowsBase"`
	CurrentLiquidationThreshold *big.Int `json:"currentLiquidationThreshold"`
	LTV                         *big.Int `json:"ltv"`
	HealthFactor                *big.Int `json:"healthFactor"`
}
