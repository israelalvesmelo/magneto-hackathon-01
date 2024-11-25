package usecase

import "github.com/israelalvesmelo/magneto-hackathon-01/internal/entity"

type CreateExchangeRateInput struct {
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
	Rate         float64 `json:"rate"`
}

type CreateExchangeRateUseCase struct {
	exchangeRateRepository entity.ExchangeRateRepository
}

func NewCreateExchangeRateUseCase(exchangeRateRepository entity.ExchangeRateRepository) *CreateExchangeRateUseCase {
	return &CreateExchangeRateUseCase{
		exchangeRateRepository: exchangeRateRepository,
	}
}

func (c *CreateExchangeRateUseCase) Execute(input CreateExchangeRateInput) error {
	exchangeRate := entity.NewExchangeRate(input.FromCurrency, input.ToCurrency, input.Rate)
	err := c.exchangeRateRepository.AddExchangeRate(exchangeRate.FromCurrency, exchangeRate.ToCurrency, exchangeRate.Rate)
	if err != nil {
		return err
	}
	return nil
}
