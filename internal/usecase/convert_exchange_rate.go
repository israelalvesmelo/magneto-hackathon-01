package usecase

import "github.com/israelalvesmelo/magneto-hackathon-01/internal/entity"

type ConvertExchangeRateInput struct {
	FromCurrency string  `json:"from_currency"`
	ToCurrency   string  `json:"to_currency"`
	Amount       float64 `json:"amount"`
}

type ConvertExchangeRateUseCase struct {
	exchangeRateRepository entity.ExchangeRateRepository
}

func NewConvertExchangeRateUseCase(exchangeRateRepository entity.ExchangeRateRepository) *ConvertExchangeRateUseCase {
	return &ConvertExchangeRateUseCase{
		exchangeRateRepository: exchangeRateRepository,
	}
}

func (c *ConvertExchangeRateUseCase) Execute(input ConvertExchangeRateInput) (float64, error) {
	rate, err := c.exchangeRateRepository.GetExchangeRate(input.FromCurrency, input.ToCurrency)
	if err != nil {
		return 0, err
	}
	convertedAmount := input.Amount * rate
	return convertedAmount, nil
}
