package models

import (
	"time"

	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	Date        time.Time `validate:"nonzero"`
	Amount      int64     `validate:"min=1"`
	Description string    `validate:"nonzero"`

	Method PaymentMethod `validate:"nonzero"`
	Type   PaymentType   `validate:"nonzero"`
}
type PaymentMethod string

const (
	Pix    PaymentMethod = "PIX"
	Credit PaymentMethod = "CREDIT"
	Debit  PaymentMethod = "DEBIT"
	Cash   PaymentMethod = "CASH"
)

type PaymentType string

const (
	Income  PaymentType = "INCOME"  // dinheiro recebido
	Expense PaymentType = "EXPENSE" // dinheiro gasto
)

func ValidatePayment(payment *Payment) error {
	if err := validator.Validate(payment); err != nil {
		return err
	}
	return nil
}
