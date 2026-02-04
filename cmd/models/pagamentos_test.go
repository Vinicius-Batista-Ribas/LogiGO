package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidatePayment(t *testing.T) {
	InitValidators()

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
