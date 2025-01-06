package main

import "fmt"

func incrementByOne(integer *int) {
	*integer++
}

func main() {

	x := 10

	fmt.Println(x)
	incrementByOne(&x)
	fmt.Println(x)

}

/*
Exercise: Increment a Value Using Pass-by-Reference

Objective:
Understand how to modify the value of a variable in a function using pass-by-reference with pointers.

Instructions:
    Write a function increment that:
    Takes a pointer to an integer as its argument.
    Increments the value of the integer by 1.
In the main function:
    Declare an integer variable x and set it to 10.
    Pass the address of x to the increment function.
    Print the value of x before and after calling the increment function.

expected output:
Before increment: 10
After increment: 11
*/
