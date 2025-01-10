package main

import (
	"errors"
	"fmt"
	"strings"
)

type customer struct {
	id      int
	balance float64
}

type transaction struct {
	customerID      int
	amount          float64
	transactionType string
}

// const ErrInsufficientFunds string = "Withdrawal with insufficient funds"
// const ErrUnknownTransactionType string = "Unknown transaction type"

var (
	ErrInsufficientFunds      = errors.New("Withdrawal with insufficient funds")
	ErrUnknownTransactionType = errors.New("Unknown transaction type")
)

func updateBalance(customer *customer, transaction transaction) error {

	if strings.Contains(transaction.transactionType, "deposit") {
		customer.balance = customer.balance + transaction.amount
		return nil
	}
	if strings.Contains(transaction.transactionType, "withdrawal") {
		customer.balance = customer.balance - transaction.amount
		return nil
	}
	if transaction.transactionType == "invalid" {
		return errors.New("haha")
	}
	return nil
}

const transactionDeposit = "Valid deposit"
const transactionWithdrawal string = "Valid withdrawal"

func main() {
	alice := &customer{id: 1, balance: 100.0}
	deposit := &transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}
	withdrawal := &transaction{customerID: 1, amount: 75, transactionType: transactionWithdrawal}
	invalid := &transaction{customerID: 1, amount: 25, transactionType: "invalid"}

	err := updateBalance(alice, deposit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Alice's balance after deposit: %.2f\n", alice.balance)
	}

	err = updateBalance(alice, withdrawal)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Alice's balance after withdrawal: %.2f\n", alice.balance)
	}

	err = updateBalance(alice, invalid)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Alice's balance after invalid transaction: %.2f\n", alice.balance)
	}
}

/*

package main

import (
	"errors"
	"fmt"
)

type customer struct {
	id      int
	balance float64
}

type transaction struct {
	customerID      int
	amount          float64
	transactionType string
}

var (
	ErrInsufficientFunds      = errors.New("insufficient funds")
	ErrUnknownTransactionType = errors.New("unknown transaction type")
)

const (
	transactionDeposit    = "deposit"
	transactionWithdrawal = "withdrawal"
)

func updateBalance(c *customer, t transaction) error {

	if t.transactionType == transactionDeposit {
		c.balance += t.amount
	} else if t.transactionType == transactionWithdrawal {
		if c.balance < t.amount {
			return ErrInsufficientFunds
		}
		c.balance -= t.amount
	} else {
		return ErrUnknownTransactionType
	}
	return nil
}

func main() {
	alice := customer{id: 1, balance: 100.0}
	deposit := transaction{customerID: 1, amount: 50, transactionType: transactionDeposit}
	withdrawal := transaction{customerID: 1, amount: 75, transactionType: transactionWithdrawal}
	invalid := transaction{customerID: 1, amount: 25, transactionType: "invalid"}

	err := updateBalance(&alice, deposit)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Alice's balance after deposit: %.2f\n", alice.balance)
	}

	err = updateBalance(&alice, withdrawal)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Alice's balance after withdrawal: %.2f\n", alice.balance)
	}

	err = updateBalance(&alice, invalid)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Alice's balance after invalid transaction: %.2f\n", alice.balance)
	}
}

*/
