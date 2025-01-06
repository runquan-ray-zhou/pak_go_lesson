package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name          string
		originalPrice float64
		discount      float64
		expected      *float64
	}
	tests := []testCase{
		{
			name:          "10% discount",
			originalPrice: 100.0,
			discount:      10.0,
			expected:      floatPtr(90.0),
		},
		{
			name:          "50% discount",
			originalPrice: 50.0,
			discount:      50.0,
			expected:      floatPtr(25.0),
		},
		{
			name:          "0% discount",
			originalPrice: 100.0,
			discount:      0.0,
			expected:      floatPtr(100.0),
		},
		{
			name:          "Negative discount",
			originalPrice: 100.0,
			discount:      -10.0,
			expected:      nil,
		},
		{
			name:          "Discount > 100%",
			originalPrice: 100.0,
			discount:      110.0,
			expected:      nil,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				name:          "25% discount",
				originalPrice: 40.0,
				discount:      25.0,
				expected:      floatPtr(30.0),
			},
			{
				name:          "No discount on 0 price",
				originalPrice: 0.0,
				discount:      10.0,
				expected:      floatPtr(0.0),
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		product := Product{Name: "Test Product", OriginalPrice: test.originalPrice}
		updateDiscountPrice(&product, test.discount)

		if !compareDiscountPrice(product.DiscountPrice, test.expected) {
			failCount++
			t.Errorf(`Test Failed: %s
input:
* Original Price: %v
* Discount: %v
expected:
%v
actual:
%v
`,
				test.name,
				test.originalPrice,
				test.discount,
				pointerToString(test.expected),
				pointerToString(product.DiscountPrice),
			)
		} else {
			passCount++
			fmt.Printf(`Test Passed: %s
input:
* Original Price: %v
* Discount: %v
expected:
%v
actual:
%v
`,
				test.name,
				test.originalPrice,
				test.discount,
				pointerToString(test.expected),
				pointerToString(product.DiscountPrice),
			)
		}
		fmt.Println("------------------------------")
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func compareDiscountPrice(actual, expected *float64) bool {
	if actual == nil && expected == nil {
		return true
	}
	if actual == nil || expected == nil {
		return false
	}
	return *actual == *expected
}

func pointerToString(ptr *float64) string {
	if ptr == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", *ptr)
}

func floatPtr(val float64) *float64 {
	return &val
}

var withSubmit = true
