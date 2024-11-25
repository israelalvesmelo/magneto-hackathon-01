package usecase

import "github.com/israelalvesmelo/magneto-hackathon-01/internal/entity"

type FindExchangeRateInput struct {
	FromCurrency string `json:"from_currency"`
	ToCurrency   string `json:"to_currency"`
}
type FindExchangeRateUseCase struct {
	exchangeRateRepository entity.ExchangeRateRepository
}

func NewFindExchangeRateUseCase(exchangeRateRepository entity.ExchangeRateRepository) *FindExchangeRateUseCase {
	return &FindExchangeRateUseCase{
		exchangeRateRepository: exchangeRateRepository,
	}
}
func (c *FindExchangeRateUseCase) Execute(input FindExchangeRateInput) (float64, error) {
	rate, err := c.exchangeRateRepository.GetExchangeRate(input.FromCurrency, input.ToCurrency)
	if err != nil {
		return 0, err
	}
	return rate, nil
}
