package main

import "fmt"

// func splitIntSlice(s []int) ([]int, []int) {

// 	mid := len(s) / 2

// 	return s[:mid], s[mid:]

// }

// func splitStringSlice(s []string) ([]string, []string) {
// 	mid := len(s) / 2

// 	return s[:mid], s[mid:]
// }

func splitArySlice[T any](s []T) ([]T, []T) {
	mid := len(s) / 2

	return s[:mid], s[mid:]
}

func main() {

	numSlice := []int{0, 1, 2, 3}

	stringSlice := []string{"a", "b", "c", "d"}

	numSliceResult1, numSliceResult2 := splitArySlice(numSlice)

	fmt.Println("numSliceResult1:", numSliceResult1, "numSliceResult2:", numSliceResult2)

	stringSliceResult1, stringSliceResult2 := splitArySlice(stringSlice)

	fmt.Println("stringSliceResult1:", stringSliceResult1, "stringSliceResult2:", stringSliceResult2)

}
