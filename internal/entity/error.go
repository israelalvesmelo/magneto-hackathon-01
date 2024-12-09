package entity

import (
	"fmt"
)

type ExchangeError struct {
	StatusCode int
	Message    string
	Err        error
}

func NewExchangeError(statusCode int, message string, err error) *ExchangeError {
	return &ExchangeError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func (e *ExchangeError) Error() string {
	if e.Err == nil {
		return e.Message
	}
	return fmt.Sprintf("message: %s, error: %v", e.Message, e.Err)
}
