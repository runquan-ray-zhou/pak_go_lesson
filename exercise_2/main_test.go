package main

import "testing"

func TestSwap(t *testing.T) {
	tests := []struct {
		a, b      int
		expectedA int
		expectedB int
	}{
		{5, 10, 10, 5},       // Positive integers
		{0, -5, -5, 0},       // Zero and negative integer
		{-10, -20, -20, -10}, // Negative integers
		{42, 42, 42, 42},     // Identical integers
	}

	for _, tt := range tests {
		a, b := tt.a, tt.b
		swap(&a, &b)

		if a != tt.expectedA || b != tt.expectedB {
			t.Errorf("swap failed: got (%d, %d), expected (%d, %d)", a, b, tt.expectedA, tt.expectedB)
		}
	}
}
