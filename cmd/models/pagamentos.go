package models

import (
	"reflect"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

/*
|--------------------------------------------------------------------------
| Model
|--------------------------------------------------------------------------
*/

type Payment struct {
	gorm.Model

	Date        time.Time `json:"date" validate:"required,notfuture"`
	Amount      float64   `json:"amount" validate:"required,gt=0"`
	Description string    `json:"description" validate:"required,notblank"`

	Method PaymentMethod `json:"method" validate:"required,payment_method"`
	Type   PaymentType   `json:"type" validate:"required,payment_type"`
}

/*
|--------------------------------------------------------------------------
| Enums
|--------------------------------------------------------------------------
*/

type PaymentMethod string

const (
	Pix    PaymentMethod = "PIX"
	Credit PaymentMethod = "CREDIT"
	Debit  PaymentMethod = "DEBIT"
	Cash   PaymentMethod = "CASH"
)

func (pm PaymentMethod) IsValid() bool {
	switch pm {
	case Pix, Credit, Debit, Cash:
		return true
	default:
		return false
	}
}

type PaymentType string

const (
	Income  PaymentType = "INCOME"
	Expense PaymentType = "EXPENSE"
)

func (pt PaymentType) IsValid() bool {
	switch pt {
	case Income, Expense:
		return true
	default:
		return false
	}
}

/*
|--------------------------------------------------------------------------
| Validator
|--------------------------------------------------------------------------
*/

var validate = validator.New()

func InitValidators() {
	validate.RegisterValidation("notblank", notBlank)
	validate.RegisterValidation("notfuture", notFuture)
	validate.RegisterValidation("payment_method", paymentMethodValidator)
	validate.RegisterValidation("payment_type", paymentTypeValidator)

	// usa o nome do campo do JSON nos erros
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("json")
	})
}

/*
|--------------------------------------------------------------------------
| Custom rules
|--------------------------------------------------------------------------
*/

func notBlank(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return strings.TrimSpace(value) != ""
}

func notFuture(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if !ok {
		return false
	}
	return !date.After(time.Now())
}

func paymentMethodValidator(fl validator.FieldLevel) bool {
	pm, ok := fl.Field().Interface().(PaymentMethod)
	if !ok {
		return false
	}
	return pm.IsValid()
}

func paymentTypeValidator(fl validator.FieldLevel) bool {
	pt, ok := fl.Field().Interface().(PaymentType)
	if !ok {
		return false
	}
	return pt.IsValid()
}

/*
|--------------------------------------------------------------------------
| Public validation
|--------------------------------------------------------------------------
*/

func ValidatePayment(payment *Payment) error {
	return validate.Struct(payment)
}

/*
|--------------------------------------------------------------------------
| GORM hooks
|--------------------------------------------------------------------------
*/

func (p *Payment) BeforeCreate(tx *gorm.DB) error {
	return ValidatePayment(p)
}

func (p *Payment) BeforeUpdate(tx *gorm.DB) error {
	return ValidatePayment(p)
}
