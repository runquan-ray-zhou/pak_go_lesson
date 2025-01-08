// main_test.go
package main

import (
	"math"
	"testing"
)

// Test cases to ensure Circle and Rectangle implement the Shape interface correctly
func TestShape(t *testing.T) {
	shapes := []Shape{
		Circle{Radius: 5},
		Rectangle{Width: 4, Height: 6},
	}

	// Expected results
	expectedAreas := []float64{
		math.Pi * 5 * 5,
		4 * 6,
	}
	expectedPerimeters := []float64{
		2 * math.Pi * 5,
		2 * (4 + 6),
	}

	for i, shape := range shapes {
		area := shape.Area()
		if area != expectedAreas[i] {
			t.Errorf("expected area %f, but got %f", expectedAreas[i], area)
		}
		perimeter := shape.Perimeter()
		if perimeter != expectedPerimeters[i] {
			t.Errorf("expected perimeter %f, but got %f", expectedPerimeters[i], perimeter)
		}
	}
}
