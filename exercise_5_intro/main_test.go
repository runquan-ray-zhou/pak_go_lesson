package main

import "testing"

func TestRectangle(t *testing.T) {
	tests := []struct {
		initial      Rectangle
		scale        float64
		expected     Rectangle
		expectedArea float64
	}{
		{Rectangle{Width: 4, Height: 5}, 2, Rectangle{Width: 8, Height: 10}, 80},
		{Rectangle{Width: 6, Height: 3}, 1.5, Rectangle{Width: 9, Height: 4.5}, 40.5},
	}

	for _, test := range tests {
		r := test.initial
		r.Resize(test.scale)

		// Check dimensions after resizing
		if r.Width != test.expected.Width || r.Height != test.expected.Height {
			t.Errorf("Resize failed: expected dimensions %+v, got %+v", test.expected, r)
		}

		// Check area after resizing
		area := r.Area()
		if area != test.expectedArea {
			t.Errorf("Area failed: expected %.2f, got %.2f", test.expectedArea, area)
		}
	}
}
