package models

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	InitValidators()
	os.Exit(m.Run())
}

func TestValidatePayment(t *testing.T) {

	testCases := []struct {
		name        string
		payment     Payment
		expectError bool
	}{
		{
			name: "valid payment",
			payment: Payment{
				Date:        time.Now(),
				Amount:      100.0,
				Description: "Test payment",
				Method:      Credit,
				Type:        Expense,
			},
			expectError: false,
		}, {
			name: "future date",
			payment: Payment{
				Date: time.Now().Add(24 * time.Hour),
			},
			expectError: true,
		}, {
			name: "negative amount",
			payment: Payment{
				Amount: -50.0,
			},
			expectError: true,
		}, {
			name: "Missing description",
			payment: Payment{
				Date:        time.Now(),
				Amount:      100.0,
				Description: "",
				Method:      Credit,
				Type:        Expense,
			},
			expectError: true,
		}, {
			name: "invalid method",
			payment: Payment{
				Date:        time.Now(),
				Amount:      100.0,
				Description: "Test payment",
				Method:      "InvalidMethod",
				Type:        Expense,
			},
			expectError: true,
		}, {
			name: "invalid type",
			payment: Payment{
				Date:        time.Now(),
				Amount:      100.0,
				Description: "Test payment",
				Method:      Credit,
				Type:        "InvalidType",
			},
			expectError: true,
		}, {
			name: "required fields missing",
			payment: Payment{
				Date:        time.Now(),
				Description: "Test payment",
				Method:      Credit,
				Type:        Expense,
			},
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePayment(&tc.payment)

			if tc.expectError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestPayment_BeforeCreate_Invalid(t *testing.T) {

	p := Payment{
		Amount: -100.0,
	}

	err := p.BeforeCreate(nil)

	if err == nil {
		t.Errorf("esperava erro no BeforeCreate para pagamento inválido")
	}
}

func TestPaymentMethodIsValid(t *testing.T) {
	assert.True(t, Pix.IsValid())
	assert.True(t, Credit.IsValid())
	assert.False(t, PaymentMethod("INVALID").IsValid())
}

func TestPaymentTypeIsValid(t *testing.T) {
	assert.True(t, Income.IsValid())
	assert.True(t, Expense.IsValid())
	assert.False(t, PaymentType("INVALID").IsValid())
}
