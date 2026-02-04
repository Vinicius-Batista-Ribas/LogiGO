package dto

import (
	"strings"
	"time"

	"LogiGO/cmd/models"
)

type CreatePaymentDTO struct {
	Date        time.Time `json:"date"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`

	Method models.PaymentMethod `json:"method"`
	Type   models.PaymentType   `json:"type"`
}

func (dto *CreatePaymentDTO) ToModel() *models.Payment {
	return &models.Payment{
		Date:        dto.Date,
		Amount:      dto.Amount,
		Description: dto.Description,
		Method:      dto.Method,
		Type:        dto.Type,
	}
}

type PaymentSummaryDTO struct {
	Income  float64 `json:"income"`
	Expense float64 `json:"expense"`
	Balance float64 `json:"balance"`
}

func (dto *CreatePaymentDTO) Normalize() {
	dto.Method = models.PaymentMethod(
		strings.ToUpper(string(dto.Method)),
	)
	dto.Type = models.PaymentType(
		strings.ToUpper(string(dto.Type)),
	)
}
