package main

import "fmt"

func swap(a, b *int) {
	// temp := *a
	*a = *b
	// *b = temp
}

func main() {
	x, y := 5, 10

	fmt.Printf("Before swap: x = %d, y = %d\n", x, y)

	swap(&x, &y)

	fmt.Printf("After swap: x = %d, y = %d\n", x, y)
}

/*
Instruction

Function to swap two integers using pointers

Expected output
Before swap: x = 5, y = 10
After swap: x = 10, y = 5

*/
