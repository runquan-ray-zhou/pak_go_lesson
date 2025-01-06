package main

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Resize(scaleFactor float64) {
	r.Width = r.Width * scaleFactor
	r.Height = r.Height * scaleFactor
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func main() {

	rect := &Rectangle{
		Width:  10.0,
		Height: 5.0,
	}
	fmt.Printf("Initial dimensions: Width = %.2f, Height = %.2f\n", rect.Width, rect.Height)
	fmt.Printf("Initial area: %.2f\n", rect.Area())

	rect.Resize(2.0)
	fmt.Printf("After resize: Width = %.2f, Height = %.2f\n", rect.Width, rect.Height)
	fmt.Printf("Updated area: %.2f\n", rect.Area())
}

/*

Exercise: Updating Struct Fields with Methods
Objective:
Understand how to use methods with structs to modify their fields.

Instructions:
Define a Rectangle struct with the following fields:

    Width (float64)
    Height (float64)

Write a method Resize for the Rectangle struct:

    The method should take a pointer receiver (*Rectangle).
    It should update the Width and Height by multiplying them with a scale factor passed as an argument.

Write another method Area for the Rectangle struct:

    It should return the area of the rectangle (Width * Height) as a float64.

In the main function:

    Create a Rectangle instance with initial values for Width and Height.
    Call the Resize method to scale the rectangle dimensions.
    Call the Area method to compute and print the updated area.

Expected output:
Initial dimensions: Width = 4, Height = 5, Area = 20
After resizing: Width = 8, Height = 10, Area = 80


*/
