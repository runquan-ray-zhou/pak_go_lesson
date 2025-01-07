package main

import (
	"fmt"
)

func filterPositiveInts(nums []int) []int {

	result := []int{}

	for _, n := range nums {
		if n > 2 {
			result = append(result, n)
		}
	}

	return result
}

func filterLongStrings(strs []string) []string {
	result := []string{}

	for _, s := range strs {
		if len(s) > 5 {
			result = append(result, s)
		}
	}
	return result
}

func main() {

	numSlice := []int{0, 1, 2, 3}

	stringSlice := []string{"apple pie", "banana", "chocolate cake", "dates"}

	filteredNum := filterPositiveInts(numSlice)

	fileredString := filterLongStrings(stringSlice)

	fmt.Println("filteredNum:", filteredNum)
	fmt.Println("fileredString:", fileredString)

}
