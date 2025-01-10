package main

import (
	"fmt"
	"testing"
)

func TestUpdateBalance(t *testing.T) {
	type testCase struct {
		name          string
		customer      customer
		transaction   transaction
		expected      float64
		expectedError error
	}

	tests := []testCase{
		{
			name:          "Valid deposit",
			customer:      customer{id: 1, balance: 100.0},
			transaction:   transaction{customerID: 1, amount: 50.0, transactionType: transactionDeposit},
			expected:      150.0,
			expectedError: nil,
		},
		{
			name:          "Valid withdrawal",
			customer:      customer{id: 2, balance: 200.0},
			transaction:   transaction{customerID: 2, amount: 50.0, transactionType: transactionWithdrawal},
			expected:      150.0,
			expectedError: nil,
		},
		{
			name:          "Withdrawal with insufficient funds",
			customer:      customer{id: 3, balance: 50.0},
			transaction:   transaction{customerID: 3, amount: 100.0, transactionType: transactionWithdrawal},
			expected:      50.0,
			expectedError: ErrInsufficientFunds,
		},
		{
			name:          "Unknown transaction type",
			customer:      customer{id: 4, balance: 100.0},
			transaction:   transaction{customerID: 4, amount: 50.0, transactionType: "unknown"},
			expected:      100.0,
			expectedError: ErrUnknownTransactionType,
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{
				name:          "Multiple deposits and withdrawals",
				customer:      customer{id: 5, balance: 50.0},
				transaction:   transaction{customerID: 5, amount: 10.0, transactionType: transactionDeposit},
				expected:      60.0,
				expectedError: nil,
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			customerCopy := test.customer
			err := updateBalance(&customerCopy, test.transaction)

			if err != test.expectedError {
				failCount++
				t.Errorf("Expected error: %v, Actual error: %v", test.expectedError, err)
			}
			if customerCopy.balance != test.expected {
				failCount++
				t.Errorf("Expected balance: %.2f, Actual balance: %.2f", test.expected, customerCopy.balance)
			}

			if err == nil && customerCopy.balance == test.expected {
				passCount++
				fmt.Println("Test Passed:", test.name)
			} else {
				failCount++
				fmt.Println("Test Failed:", test.name)
			}
			fmt.Println("------------------------------")
		})
	}
	fmt.Printf("Tests Passed: %d, Tests Failed: %d\n", passCount, failCount)
}

var withSubmit = true
