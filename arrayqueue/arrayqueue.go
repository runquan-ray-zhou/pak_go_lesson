package main

import "fmt"

type QueueInterface[T any] interface {
	Enqueue(val T)  // Adds val to the rear of the queue.
	Dequeue()       // Removes the item from the front of the queue.
	Front() T       // Returns the front item of the queue without removing it.
	IsEmpty() bool  // Returns whether the queue is empty.
	Size() int      // Returns the number of elements in the queue.
	String() string // Returns a string representation of the queue.
}

type ArrayQueue[T any] struct {
	arr   []T
	front int
	rear  int
	size  int
}

func NewArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return &ArrayQueue[T]{
		arr:   make([]T, capacity),
		front: 0,
		rear:  -1,
		size:  0,
	}
}

func (q *ArrayQueue[T]) Enqueue(val T) {
	if q.size == len(q.arr) {
		return
	}

	q.rear = (q.rear + 1) % len(q.arr)
	q.arr[q.rear] = val
	q.size++
}

func (q *ArrayQueue[T]) Dequeue() {
	if q.IsEmpty() {
		return
	}

	q.front = (q.front + 1) % len(q.arr)
	q.size--
}

func (q *ArrayQueue[T]) Front() T {
	if q.IsEmpty() {
		var zero T
		return zero
	}
	return q.arr[q.front]
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *ArrayQueue[T]) Size() int {
	return q.size
}

func (q *ArrayQueue[T]) String() string {
	result := ""
	count := 0

	for i := q.front; count < q.size; i = (i + 1) % len(q.arr) {
		result += fmt.Sprintf("%v", q.arr[i])

		if count < q.size-1 {
			result += " -> "
		}
		count++
	}
	return result
}
