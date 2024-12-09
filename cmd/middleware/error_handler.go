package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/israelalvesmelo/magneto-hackathon-01/internal/entity"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	lastErr := c.Errors.Last()

	if lastErr == nil {
		return
	}
	statusCode := http.StatusInternalServerError
	msgError := lastErr.Error()

	switch err := lastErr.Err.(type) {
	case *entity.ExchangeError:
		statusCode = err.StatusCode
		msgError = err.Error()
		fmt.Print("entrou aqui")

		fmt.Print(msgError)
	}
	fmt.Print(msgError)

	c.JSON(statusCode, msgError)

}
