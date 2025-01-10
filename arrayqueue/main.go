package main

import (
	"fmt"
)

func main() {
	queue := NewArrayQueue[int](5) // Create queue with capacity 5
	fmt.Print(queue.arr)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println("Queue:", queue.String())     // Output: 1 -> 2 -> 3
	fmt.Println("Front:", queue.Front())      // Output: 1
	fmt.Println("Size:", queue.Size())        // Output: 3
	fmt.Println("Is Empty:", queue.IsEmpty()) // Output: false

	queue.Dequeue()
	fmt.Println("Queue after dequeue:", queue.String()) // Output: 2 -> 3

	queue.Enqueue(4)
	queue.Enqueue(5)
	fmt.Println("Queue:", queue.String()) // Output: 2 -> 3 -> 4 -> 5

	queue.Enqueue(6)                      // Try to enqueue when full (currently does nothing)
	fmt.Println("Queue:", queue.String()) // Output: 2 -> 3 -> 4 -> 5 -> 6
}
