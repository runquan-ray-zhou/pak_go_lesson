package main

import (
	"fmt"
)

type CreditCard struct {
	CardNumber, ExpiryDate, CVV string
}

type PayPay struct {
	Email string
}

type BankTransfer struct {
	AccountNumber, RoutingNumber string
}

type PaymentMethod interface {
	ProcessPayment(amount float64) error
	GenerateReceipt(amount float64) string
}

func (c *CreditCard) ProcessPayment(amount float64) error {
	fmt.Printf("Credit Card Payment\n Amount: %v\n Card Number: %v\n Expiry Date: %v\n", amount, c.CardNumber, c.ExpiryDate)

	return nil
}

func (c *CreditCard) GenerateReceipt(amount float64) string {
	return fmt.Sprintf("Credit Card Payment\n Amount: %v\n Card Number: %v\n Expiry Date: %v\n", amount, c.CardNumber, c.ExpiryDate)
}

func processPayment(p PaymentMethod, amount float64) {
	err := p.ProcessPayment(amount)

	if err != nil {
		fmt.Println("Payment failure", err)
		return
	}

	receipt := p.GenerateReceipt(amount)
	fmt.Println(receipt)
}

func main() {
	creditCard := &CreditCard{CardNumber: "1234567890123456", ExpiryDate: "12/24", CVV: "123"}
	// payPal := &PayPal{Email: "example@example.com"}
	// bankTransfer := &BankTransfer{AccountNumber: "1234567890", RoutingNumber: "123456789"}

	processPayment(creditCard, 10.0)
	fmt.Println()
	// processPayment(payPal, 20.0)
	// fmt.Println()
	// processPayment(bankTransfer, 5.0)
}

/*
package main

import (
	"fmt"
)

type PaymentMethod interface {
	ProcessPayment(amount float64) error
	GenerateReceipt(amount float64) string
}

type CreditCard struct {
	CardNumber string
	ExpiryDate string
	CVV        string
}

func (c *CreditCard) ProcessPayment(amount float64) error {

	//this is where you make API call
	fmt.Printf("Processing credit card payment of $%.2f with card ending in %s\n", amount, c.CardNumber[len(c.CardNumber)-4:])

	return nil
}

func (c *CreditCard) GenerateReceipt(amount float64) string {
	return fmt.Sprintf("Credit Card Payment\nAmount: $%.2f\nCard Number: %s\nExpiry Date: %s", amount, c.CardNumber, c.ExpiryDate)
}

type PayPal struct {
	Email string
}

func (p *PayPal) ProcessPayment(amount float64) error {

	//this is where you make API call
	fmt.Printf("Processing PayPal payment of $%.2f for email %s\n", amount, p.Email)

	return nil
}

func (p *PayPal) GenerateReceipt(amount float64) string {
	return fmt.Sprintf("PayPal Payment\nAmount: $%.2f\nEmail: %s", amount, p.Email)
}

type BankTransfer struct {
	AccountNumber string
	RoutingNumber string
}

func (b *BankTransfer) ProcessPayment(amount float64) error {
	// In a real system, you would integrate with a banking API here.
	fmt.Printf("Processing bank transfer of $%.2f for account %s\n", amount, b.AccountNumber)
	return nil
}

func (b *BankTransfer) GenerateReceipt(amount float64) string {
	return fmt.Sprintf("Bank Transfer Payment\nAmount: $%.2f\nAccount Number: %s\nRouting Number: %s", amount, b.AccountNumber, b.RoutingNumber)
}

func processPayment(pm PaymentMethod, amount float64) {
	err := pm.ProcessPayment(amount)

	if err != nil {
		fmt.Println("Payment failed:", err)
		return
	}

	receipt := pm.GenerateReceipt(amount)
	fmt.Println("Receipt:\n", receipt)

}

func main() {
	creditCard := &CreditCard{CardNumber: "1234567890123456", ExpiryDate: "12/24", CVV: "123"}
	PayPal := &PayPal{Email: "example@example.com"}
	bankTransfer := &BankTransfer{AccountNumber: "1234567890", RoutingNumber: "123456789"}

	processPayment(creditCard, 10.0)
	fmt.Println()
	processPayment(PayPal, 20.0)
	fmt.Println()
	processPayment(bankTransfer, 5.0)
}

*/
