package main

import "fmt"

type BankAccount struct {
	AccountHolder string
	Balance       float64
}

func (b *BankAccount) Deposit(amount float64) bool {
	if amount <= 0 {
		return false
	}
	b.Balance = b.Balance + amount
	if b.Balance <= 0 {
		return false
	}
	return true
}

func (b *BankAccount) Withdraw(amount float64) bool {
	if amount <= 0 {
		return false
	}
	if b.Balance > amount {
		b.Balance = b.Balance - amount
		return true
	}
	return false
}

func (b *BankAccount) Transfer(anotherAccount *BankAccount, amount float64) bool {
	if amount <= 0 || amount > b.Balance {
		return false
	}

	if b.Withdraw(amount) == false {
		return false
	}
	if anotherAccount.Deposit(amount) == false {
		return false
	}

	return true
}

func main() {

	account1 := &BankAccount{AccountHolder: "Alice", Balance: 1000}
	account2 := &BankAccount{AccountHolder: "Bob", Balance: 500}

	fmt.Printf("Before Transfer: %s has $%.2f, %s has $%.2f\n",
		account1.AccountHolder, account1.Balance, account2.AccountHolder, account2.Balance)

	success := account1.Transfer(account2, 200)
	if success {
		fmt.Println("Transfer successful!")
	} else {
		fmt.Println("Transfer failed!")
	}

	fmt.Printf("After Transfer: %s has $%.2f, %s has $%.2f\n",
		account1.AccountHolder, account1.Balance, account2.AccountHolder, account2.Balance)
}

/*

Exercise: Bank Account Management
Objective:
Design a system with structs and methods to manage bank accounts and perform transactions between them.

Instructions:
Create a BankAccount struct with the following fields:
    AccountHolder (string)
    Balance (float64)

Implement the following methods for the BankAccount struct:
    Deposit(amount float64):
        Add amount to Balance.
        Return true if successful, false if amount is negative or zero.

    Withdraw(amount float64):
        Subtract amount from Balance if sufficient funds exist.
        Return true if successful, false if insufficient funds or amount is negative.

Create a Transfer method:
    It should take another BankAccount pointer and an amount as arguments.
    Withdraw the amount from the current account and deposit it into the other account if the transaction is valid.
    Return true if the transfer is successful, false otherwise.

Expected output:
Before Transfer: Alice has $1000.00, Bob has $500.00
Transfer successful!
After Transfer: Alice has $800.00, Bob has $700.00


*/
