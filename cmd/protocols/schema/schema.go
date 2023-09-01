package schema

import (
	"encoding/json"
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

// Leaves out the Params field
func (m *MarketInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Protocol   string `json:"protocol"`
		Chain      string `json:"chain"`
		Token      string `json:"token"`
		Decimals   string `json:"decimals"`
		LTV        string `json:"ltv"`        // In basis points, 0 if cannot be collateral
		PriceInUSD string `json:"priceInUSD"` // How much USD is required to purchase 1 ether unit, with 8 decimals
	}{
		Protocol:   m.Protocol,
		Chain:      m.Chain,
		Token:      m.Token,
		Decimals:   m.Decimals.String(),
		LTV:        m.LTV.String(),
		PriceInUSD: m.PriceInUSD.String(),
	})
}

func (m *MarketInfo) UnmarshalJSON(data []byte) error {
	// Declare an anonymous struct with BigNum as string
	var temp struct {
		Protocol   string `json:"protocol"`
		Chain      string `json:"chain"`
		Token      string `json:"token"`
		Decimals   string `json:"decimals"`
		LTV        string `json:"ltv"`        // In basis points, 0 if cannot be collateral
		PriceInUSD string `json:"priceInUSD"` // How much USD is required to purchase 1 ether unit, with 8 decimals
	}

	// Unmarshal the json bytes into the temp struct
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Populate the fields of ComplexResponse
	m.Protocol = temp.Protocol
	m.Chain = temp.Chain
	m.Token = temp.Token

	// Convert string to big.Int
	m.Decimals = new(big.Int)
	m.Decimals.SetString(temp.Decimals, 10)
	m.LTV = new(big.Int)
	m.LTV.SetString(temp.LTV, 10)
	m.PriceInUSD = new(big.Int)
	m.PriceInUSD.SetString(temp.PriceInUSD, 10)

	return nil
}

// Stringifies the APY and Amount
func (s *StrategyStep) MarshalJSON() ([]byte, error) {
	type Alias StrategyStep
	return json.Marshal(&struct {
		APY    string `json:"apy"`
		Amount string `json:"amount"`
		*Alias
	}{
		APY:    s.APY.String(),
		Amount: s.Amount.String(),
		Alias:  (*Alias)(s),
	})
}

func (s *StrategyStep) UnmarshalJSON(data []byte) error {
	// Declare an anonymous struct with BigNum as string
	var temp struct {
		Market   *MarketInfo `json:"market"`
		IsSupply bool        `json:"isSupply"`
		APY      string      `json:"apy"`
		Amount   string      `json:"amount"`
	}

	// Unmarshal the json bytes into the temp struct
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Populate the fields of ComplexResponse
	s.Market = temp.Market
	s.IsSupply = temp.IsSupply

	// Convert string to big.Int
	s.APY = new(big.Int)
	s.APY.SetString(temp.APY, 10)
	s.Amount = new(big.Int)
	s.Amount.SetString(temp.Amount, 10)

	return nil
}

// Stringifies the APY
func (s *Strategy) MarshalJSON() ([]byte, error) {
	type Alias Strategy
	return json.Marshal(&struct {
		APY string `json:"apy"`
		*Alias
	}{
		APY:   s.APY.String(),
		Alias: (*Alias)(s),
	})
}

func (s *Strategy) UnmarshalJSON(data []byte) error {
	// Declare an anonymous struct with BigNum as string
	var temp struct {
		Steps []*StrategyStep `json:"steps"`
		APY   string          `json:"apy"`
	}

	// Unmarshal the json bytes into the temp struct
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	// Populate the fields of ComplexResponse
	s.Steps = temp.Steps

	// Convert string to big.Int
	s.APY = new(big.Int)
	s.APY.SetString(temp.APY, 10)

	return nil
}
