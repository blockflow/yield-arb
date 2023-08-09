package types

import "math/big"

type ProtocolMarkets struct {
	Protocol       string        `json:"protocol"`
	Chain          string        `json:"chain"`
	LendingSpecs   []*TokenSpecs `json:"lendingSpecs"`
	BorrowingSpecs []*TokenSpecs `json:"borrowingSpecs"`
}

type TokenSpecs struct {
	Protocol string     `json:"protocol"`
	Chain    string     `json:"chain"`
	Token    string     `json:"token"`
	LTV      *big.Float `json:"ltv"` // 0 if cannot be collateral
	APY      *big.Float `json:"apy"`
}

type AccountData struct {
	TotalCollateralBase         *big.Int `json:"totalCollateralBase"`
	TotalDebtBase               *big.Int `json:"totalDebtBase"`
	AvailableBorrowsBase        *big.Int `json:"availableBorrowsBase"`
	CurrentLiquidationThreshold *big.Int `json:"currentLiquidationThreshold"`
	LTV                         *big.Int `json:"ltv"`
	HealthFactor                *big.Int `json:"healthFactor"`
}
