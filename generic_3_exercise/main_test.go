package main

import (
	"testing"
)

func TestUpdateValue(t *testing.T) {
	t.Run("Update int", func(t *testing.T) {
		ptr := new(int)
		*ptr = 5
		newValue := 10
		expected := 10

		updateValue(ptr, newValue)

		if *ptr != expected {
			t.Errorf("Expected: %d, Got: %d", expected, *ptr)
		}
	})

	t.Run("Update string", func(t *testing.T) {
		ptr := new(string)
		*ptr = "hello"
		newValue := "world"
		expected := "world"

		updateValue(ptr, newValue)

		if *ptr != expected {
			t.Errorf("Expected: %s, Got: %s", expected, *ptr)
		}
	})

	t.Run("Update struct", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		ptr := &Point{X: 1, Y: 2}
		newValue := Point{X: 3, Y: 4}
		expected := Point{X: 3, Y: 4}

		updateValue(ptr, newValue)

		if *ptr != expected {
			t.Errorf("Expected: %+v, Got: %+v", expected, *ptr)
		}
	})

	t.Run("Update with nil pointer", func(t *testing.T) {
		var intPtr *int = nil
		newValue := 10
		updateValue(intPtr, newValue)

		// Expecting no panic and no change
	})
}
