package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

type ErrorPaymentMethod struct{}

func (e *ErrorPaymentMethod) ProcessPayment(amount float64) error {
	return errors.New("payment failed")
}

func (e *ErrorPaymentMethod) GenerateReceipt(amount float64) string {
	return ""
}

func TestProcessPayment(t *testing.T) {
	tests := []struct {
		name            string
		pm              PaymentMethod
		amount          float64
		expectError     bool
		expectedReceipt string
	}{
		{
			name:            "Credit Card Payment",
			pm:              &CreditCard{CardNumber: "1234567890123456", ExpiryDate: "12/24", CVV: "123"},
			amount:          10.0,
			expectError:     false,
			expectedReceipt: "Credit Card Payment\nAmount: $10.00\nCard Number: 1234567890123456\nExpiry Date: 12/24",
		},
		{
			name:            "PayPal Payment",
			pm:              &PayPal{Email: "test@example.com"},
			amount:          20.0,
			expectError:     false,
			expectedReceipt: "PayPal Payment\nAmount: $20.00\nEmail: test@example.com",
		},
		{
			name:            "Bank Transfer Payment",
			pm:              &BankTransfer{AccountNumber: "1234567890", RoutingNumber: "987654321"},
			amount:          5.0,
			expectError:     false,
			expectedReceipt: "Bank Transfer Payment\nAmount: $5.00\nAccount Number: 1234567890\nRouting Number: 987654321",
		},
		{
			name:            "Error Payment",
			pm:              &ErrorPaymentMethod{},
			amount:          50.0,
			expectError:     true,
			expectedReceipt: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			processPayment(tt.pm, tt.amount)

			// Restore stdout
			w.Close()
			os.Stdout = oldStdout

			// Read the captured output
			var buf bytes.Buffer
			io.Copy(&buf, r)
			output := buf.String()

			if tt.expectError {
				if !strings.Contains(output, "Payment failed:") {
					t.Errorf("Expected an error message, but none was printed")
				}
			} else {
				if !strings.Contains(output, tt.expectedReceipt) {
					t.Errorf("Expected receipt: %s, got: %s", tt.expectedReceipt, output)
				}
			}
		})
	}
}
