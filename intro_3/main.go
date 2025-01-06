package main

import "fmt"

func increment(x *int) {
	*x++
	fmt.Println("Inside increment: ", x)
}

func main() {
	x := 5

	increment(&x) // passing in the address

	fmt.Println("In Main: ", x)

}
