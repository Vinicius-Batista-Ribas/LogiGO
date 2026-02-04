package erros

import "errors"

var (
	ErrInvalidAmount   = errors.New("amount must be greater than zero")
	ErrPaymentNotFound = errors.New("payment not found")
	ErrInvalidStatus   = errors.New("invalid payment status")
	ErrInvalidType     = errors.New("invalid payment type")
)
