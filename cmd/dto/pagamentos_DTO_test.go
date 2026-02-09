package dto

import (
	"LogiGO/cmd/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreatePaymentDTO_ToModel(t *testing.T) {
	date := time.Now()
	dto := CreatePaymentDTO{
		Date:        date,
		Amount:      100.0,
		Description: "Test Payment",
		Method:      models.Credit,
		Type:        models.Expense,
	}
	payment := dto.ToModel()

	assert.Equal(t, dto.Date, payment.Date)
	assert.Equal(t, dto.Amount, payment.Amount)
	assert.Equal(t, dto.Description, payment.Description)
	assert.Equal(t, dto.Method, payment.Method)
	assert.Equal(t, dto.Type, payment.Type)
}

func TestCreatePaymentDTO_Normalize(t *testing.T) {
	dto := CreatePaymentDTO{
		Method: "credit",
		Type:   "expense",
	}
	dto.Normalize()

	assert.Equal(t, models.Credit, dto.Method)
	assert.Equal(t, models.Expense, dto.Type)
}
