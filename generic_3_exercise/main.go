package main

import "fmt"

func updateValue[T any](value *T, updateValue T) {
	if value == nil {
		return
	}
	*value = updateValue

}

func main() {

	intPtr := new(int)
	*intPtr = 5
	fmt.Println("Before update:", *intPtr)
	updateValue(intPtr, 10)
	fmt.Println("Updated int:", *intPtr)

	strPtr := new(string)
	*strPtr = "hello"
	fmt.Println("Before update:", *strPtr)
	updateValue(strPtr, "world")
	fmt.Println("Updated string:", *strPtr)

	type Point struct {
		X, Y int
	}
	pointPtr := &Point{X: 1, Y: 2}
	fmt.Println("Before update:", *pointPtr)
	updateValue(pointPtr, Point{X: 3, Y: 4})
	fmt.Println("Updated Point:", *pointPtr)
}

/*

Exercise: Generic UpdateValue Function with Pointers

Scenario:

You need to create a generic function updateValue that can update the value pointed to by a pointer. The function should work with different data types (e.g., int, string, custom structs)

expected output
Before update: 5
Updated int: 10
Before update: hello
Updated string: world
Before update: {1 2}
Updated Point: {3 4}

*/
