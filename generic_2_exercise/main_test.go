package main

import (
	"slices"
	"testing"
)

func TestFilterSlice(t *testing.T) {
	t.Run("filter positive integers", func(t *testing.T) {
		input := []int{-2, -1, 0, 1, 2, 3}
		expected := []int{1, 2, 3}
		result := filterSlice(input, func(n int) bool { return n > 0 })

		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})

	t.Run("filter strings longer than 3 chars", func(t *testing.T) {
		input := []string{"a", "ab", "abc", "abcd", "abcde"}
		expected := []string{"abcd", "abcde"}
		result := filterSlice(input, func(s string) bool { return len(s) > 3 })

		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})

	t.Run("filter empty slice", func(t *testing.T) {
		var input []int
		var expected []int
		result := filterSlice(input, func(n int) bool { return n > 0 })

		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})

	t.Run("filter where no elements match", func(t *testing.T) {
		input := []int{1, 2, 3}
		var expected []int
		result := filterSlice(input, func(n int) bool { return n > 10 })

		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})

	t.Run("filter where all elements match", func(t *testing.T) {
		input := []int{1, 2, 3}
		expected := []int{1, 2, 3}
		result := filterSlice(input, func(n int) bool { return n > 0 })

		if !slices.Equal(result, expected) {
			t.Errorf("Expected %v but got %v", expected, result)
		}
	})
}
