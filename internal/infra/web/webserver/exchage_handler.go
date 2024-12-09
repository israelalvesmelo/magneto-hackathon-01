package webserver

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/israelalvesmelo/magneto-hackathon-01/internal/entity"
	"github.com/israelalvesmelo/magneto-hackathon-01/internal/usecase"
)

type ExchangeHandler struct {
	CreateExchangeRateUseCase  usecase.CreateExchangeRateUseCase
	FindExchangeRateUseCase    usecase.FindExchangeRateUseCase
	ConvertExchangeRateUseCase usecase.ConvertExchangeRateUseCase
}

func NewExchangeHandler(createExchangeRateUseCase usecase.CreateExchangeRateUseCase,
	findExchangeRateUseCase usecase.FindExchangeRateUseCase,
	convertExchangeRateUseCase usecase.ConvertExchangeRateUseCase,
) *ExchangeHandler {
	return &ExchangeHandler{
		CreateExchangeRateUseCase:  createExchangeRateUseCase,
		FindExchangeRateUseCase:    findExchangeRateUseCase,
		ConvertExchangeRateUseCase: convertExchangeRateUseCase,
	}
}

func (h *ExchangeHandler) AddExchangeRate(c *gin.Context) {
	var dto usecase.CreateExchangeRateInput
	if err := c.ShouldBindJSON(&dto); err != nil {
		msgErr := "Erro ao converter JSON"
		c.Error(entity.NewExchangeError(http.StatusBadRequest, msgErr, err))
		return
	}
	err := h.CreateExchangeRateUseCase.Execute(dto)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Taxa de câmbio adicionada com sucesso!"})
}

func (h *ExchangeHandler) GetExchangeRate(c *gin.Context) {
	var dto usecase.FindExchangeRateInput
	dto.FromCurrency = c.Query("from_currency")
	dto.ToCurrency = c.Query("to_currency")
	if dto.FromCurrency == "" || dto.ToCurrency == "" {
		msgErr := "Parâmetros inválidos"
		c.Error(entity.NewExchangeError(http.StatusBadRequest, msgErr, nil))
		return
	}

	rate, err := h.FindExchangeRateUseCase.Execute(dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"rate": rate})
}

func (h *ExchangeHandler) ConvertAmount(c *gin.Context) {
	var dto usecase.ConvertExchangeRateInput
	dto.FromCurrency = c.Query("from_currency")
	dto.ToCurrency = c.Query("to_currency")
	amount := c.Query("amount")
	if dto.FromCurrency == "" || dto.ToCurrency == "" || amount == "" {
		msgErr := "Parâmetros inválidos"
		c.Error(entity.NewExchangeError(http.StatusBadRequest, msgErr, nil))
		return
	}
	dto.Amount, _ = strconv.ParseFloat(amount, 64)

	result, err := h.ConvertExchangeRateUseCase.Execute(dto)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"converted_amount": result})
}
