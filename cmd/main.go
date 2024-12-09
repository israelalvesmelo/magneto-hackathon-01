package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/israelalvesmelo/magneto-hackathon-01/cmd/middleware"
	"github.com/israelalvesmelo/magneto-hackathon-01/configs"
	repositories "github.com/israelalvesmelo/magneto-hackathon-01/internal/infra/database"
	"github.com/israelalvesmelo/magneto-hackathon-01/internal/infra/web/webserver"

	"github.com/israelalvesmelo/magneto-hackathon-01/internal/usecase"
	"github.com/israelalvesmelo/magneto-hackathon-01/pkg/database"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := database.InitDB(config.DBLocation)
	if err != nil {
		fmt.Printf("Erro ao conectar ao banco de dados: %v\n", err)
		return
	}

	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	exchangeRateRepository := repositories.NewExchangeRateRepository(db)

	createExchangeRateUseCase := usecase.NewCreateExchangeRateUseCase(exchangeRateRepository)
	findExchangeRateUseCase := usecase.NewFindExchangeRateUseCase(exchangeRateRepository)
	convertExchangeRateUseCase := usecase.NewConvertExchangeRateUseCase(exchangeRateRepository)

	exchangeHandler := webserver.NewExchangeHandler(*createExchangeRateUseCase,
		*findExchangeRateUseCase, *convertExchangeRateUseCase)

	r.POST("/exchange-rate", exchangeHandler.AddExchangeRate)
	r.GET("/exchange-rate", exchangeHandler.GetExchangeRate)
	r.GET("/convert", exchangeHandler.ConvertAmount)

	r.Run()
}
