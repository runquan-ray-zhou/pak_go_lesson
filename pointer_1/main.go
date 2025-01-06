package main

import "fmt"

type Product struct {
	Name          string
	OriginalPrice float64
	DiscountPrice *float64
}

func updateDiscountPrice(product *Product, discountPercent float64) {
	if product.OriginalPrice == 0 {
		discountPrice := new(float64)
		*discountPrice = 0
		product.DiscountPrice = discountPrice
	}
	if discountPercent < 0 || discountPercent > product.OriginalPrice {
		return
	} else {
		discountPrice := new(float64)
		*discountPrice = product.OriginalPrice - product.OriginalPrice*discountPercent/100
		product.DiscountPrice = discountPrice
	}
}

func main() {

	newProduct := &Product{Name: "Example Product", OriginalPrice: 100}

	var amount float64 = 10
	discountPercent := amount

	updateDiscountPrice(newProduct, discountPercent)

	fmt.Println(newProduct)
}

/*

# Introduction to Pointers - Exercise

This exercise focuses on understanding pointers in Go, particularly how they interact with struct fields and function calls.

## Scenario

We are building a simple inventory management system. We have a `Product` struct that stores information about a product:

### Your Task

Understand the Code: Carefully examine the code in main.go and main_test.go. Understand how the Product struct is defined and how updateDiscountPrice is intended to work.

Run the Tests: Run the tests using go test -v in your terminal. You will see that some tests fail. Analyze the output to understand why they are failing.

Expected output
Product Name: Example Product
Original Price: 100
Discounted Price: 90


package main

import "fmt"

type Product struct {
	Name          string
	OriginalPrice float64
	DiscountPrice *float64
}

func updateDiscountPrice(product *Product, discountPercent float64) {
	// Handle invalid discount percentages
	if discountPercent < 0 || discountPercent > 100 {
		product.DiscountPrice = nil
		return
	}

	// Handle zero original price
	if product.OriginalPrice == 0 {
		product.DiscountPrice = floatPtr(0.0)
		return
	}

	// Calculate and update the discount price
	discountPrice := product.OriginalPrice * (1 - discountPercent/100)
	product.DiscountPrice = floatPtr(discountPrice)
}

func floatPtr(value float64) *float64 {
	return &value
}

func main() {
	newProduct := &Product{Name: "Example Product", OriginalPrice: 100}

	var amount float64 = 10
	discountPercent := amount

	updateDiscountPrice(newProduct, discountPercent)

	fmt.Println(newProduct)
}



*/
