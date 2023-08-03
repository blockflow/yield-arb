package protocols

type Protocol interface {
	GetChains() ([]string, error)
	Connect(chain string) error
	GetLendingTokens() ([]string, error)
	// GetBorrowingTokens() ([]string, error)
	// GetLendingAPYs(tokens []string) (map[string]float64, error)
	// GetBorrowingAPYs(tokens []string) (map[string]float64, error)
}
