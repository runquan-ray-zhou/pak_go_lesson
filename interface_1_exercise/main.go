// main.go
package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius int
}

type Rectangle struct {
	Width  int
	Height int
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

func (c Circle) Perimeter() float64 {
	return math.Pi * float64(2*c.Radius)
}

func (r Rectangle) Area() float64 {
	return float64(r.Width * r.Height)
}

func (r Rectangle) Perimeter() float64 {
	return float64(2*r.Width + 2*r.Height)
}

func main() {
	// Create instances of Circle and Rectangle
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Print their areas and perimeters
	shapes := []Shape{circle, rectangle}
	for _, shape := range shapes {
		fmt.Printf("Area: %.2f\n", shape.Area())
		fmt.Printf("Perimeter: %.2f\n", shape.Perimeter())
	}
}

/*

Interface Exercise
This exercise is designed to help you understand and practice implementing interfaces in Go. You will define an interface and implement it in multiple struct types, then write test cases to verify your implementations.

Objective
Define a Shape Interface:

Create an interface with methods to calculate the area and perimeter of a shape.

Implement the Interface:

Implement the interface in two struct types: Circle and Rectangle.

Expected output
Area: 78.54
Perimeter: 31.42
Area: 24.00
Perimeter: 20.00

Verify: Run the tests again using go test -v. They should now all pass.

*/
