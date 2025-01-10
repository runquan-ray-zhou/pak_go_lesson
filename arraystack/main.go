package main

import (
	"fmt"
)

func main() {
	arrayStack := NewArrayStack[float64]()

	arrayStack.Push(1.1)
	arrayStack.Push(2.2)
	arrayStack.Push(3.3)

	fmt.Println("ArrayStack:", arrayStack.String())

	fmt.Println("Top:", arrayStack.Top())
	fmt.Println("Size:", arrayStack.Size())
	fmt.Println("Empty:", arrayStack.Empty())

	arrayStack.Pop()

	fmt.Println("ArrayStack after Pop:", arrayStack.String())
}
