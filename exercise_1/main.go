package main

import "fmt"

func main() {

	x := 10
	p := &x

	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(*p)

	*p = 20

	fmt.Println(x)

}

/*
Objective:
Understand how to use pointers to access and modify variable values.

Instructions:
    Create a variable x with the value 10.
    Create a pointer p that stores the address of x.
Print:
    The value of x.
    The address of x.
    The value of p.
    The value stored at the address p points to (dereference the pointer).

Modify the value of x via the pointer p by setting it to 20.
Print the updated value of x.

Sample output:
Value of x: 10
Address of x: 0xc000018030
Value of pointer p: 0xc000018030
Value at pointer p: 10
Updated value of x: 20

*/
