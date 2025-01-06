package main

import (
	"fmt"
)

func main() {

	x := 42
	p := &x

	fmt.Println("x: ", x)   // value of x
	fmt.Println("&x: ", &x) // address
	fmt.Println("p: ", p)   // address

	fmt.Println("*p: ", *p) // value of x after dereferencing p
	// fmt.Println("*p: ", x*) // not referencing value

	var check bool = x == *p
	fmt.Println(check)

	*p = 99

	fmt.Println("x: ", x)
	fmt.Println("====")
	fmt.Println("&x: ", &x)
	fmt.Println("p: ", p)

	fmt.Println("&p: ", &p)
	fmt.Println()

}
