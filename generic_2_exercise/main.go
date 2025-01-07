package main

import "fmt"

func filterSlice[T any](input []T, keepItem func(n T) bool) []T {

	newSlice := []T{}

	for _, val := range input {
		if keepItem(val) {
			newSlice = append(newSlice, val)
		}
	}
	return newSlice
}

func keepOddNumber(n int) bool {
	if n%2 != 0 {
		return true
	}
	return false
}

func keepCherryAndFig(n string) bool {
	if n == "cherry" {
		return true
	}
	if n == "fig" {
		return true
	}
	return false
}

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	strings := []string{"apple", "banana", "cherry", "date", "fig", "grape"}

	oddNumbers := filterSlice(numbers, keepOddNumber)
	fmt.Println(oddNumbers)
	cheeryAndFig := filterSlice(strings, keepCherryAndFig)
	fmt.Println(cheeryAndFig)
}

/*

Implement a generic function, filterSlice, that filters elements from a slice based on a condition defined by a function passed as an argument. The goal is to return a new slice containing only the elements that meet the specified condition.

Expected Behavior
After completing the filterSlice function:

It should work with slices of any type.
It should dynamically filter elements based on the provided keepItem function.

Run the Tests: Run the tests using go test -v in your terminal. You will see that some tests fail. Analyze the output to understand why they are failing.

Expected output
oddNumbers: [1 3 5 7 9]
filtered fruits [cherry fig]

*/
