package main

import "testing"

func TestBankAccount(t *testing.T) {
	t.Run("Test Deposit", func(t *testing.T) {
		account := BankAccount{AccountHolder: "TestUser", Balance: 0}

		if !account.Deposit(500) || account.Balance != 500 {
			t.Errorf("Deposit failed: expected balance 500, got %.2f", account.Balance)
		}

		if account.Deposit(-100) {
			t.Error("Deposit should fail for negative amount")
		}
	})

	t.Run("Test Withdraw", func(t *testing.T) {
		account := BankAccount{AccountHolder: "TestUser", Balance: 500}

		if !account.Withdraw(200) || account.Balance != 300 {
			t.Errorf("Withdraw failed: expected balance 300, got %.2f", account.Balance)
		}

		if account.Withdraw(1000) {
			t.Error("Withdraw should fail for insufficient funds")
		}
	})

	t.Run("Test Transfer", func(t *testing.T) {
		account1 := BankAccount{AccountHolder: "Alice", Balance: 1000}
		account2 := BankAccount{AccountHolder: "Bob", Balance: 500}

		if !account1.Transfer(&account2, 300) {
			t.Error("Transfer failed: expected success")
		}

		if account1.Balance != 700 || account2.Balance != 800 {
			t.Errorf("Transfer failed: expected balances 700 and 800, got %.2f and %.2f",
				account1.Balance, account2.Balance)
		}

		if account1.Transfer(&account2, 2000) {
			t.Error("Transfer should fail for insufficient funds")
		}
	})
}
