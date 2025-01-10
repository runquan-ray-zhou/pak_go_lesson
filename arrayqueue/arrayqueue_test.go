package main

import (
	"testing"
)

// Helper function to create a new ArrayQueue for testing.
func newArrayQueue[T any](capacity int) *ArrayQueue[T] {
	return NewArrayQueue[T](capacity)
}

func TestArrayQueue_NewArrayQueue(t *testing.T) {
	t.Run("New queue is empty", func(t *testing.T) {
		q := newArrayQueue[int](5)

		if q == nil {
			t.Fatal("NewArrayQueue returned nil")
		}

		if !q.IsEmpty() {
			t.Errorf("Expected new queue to be empty, but it's not")
		}

		if q.Size() != 0 {
			t.Errorf("Expected new queue to have size 0, but got %d", q.Size())
		}
	})

	t.Run("New queue with different types", func(t *testing.T) {
		intQueue := newArrayQueue[int](5)
		stringQueue := newArrayQueue[string](10)
		boolQueue := newArrayQueue[bool](3)

		if intQueue == nil || stringQueue == nil || boolQueue == nil {
			t.Fatal("NewArrayQueue returned nil for some type")
		}

		if !intQueue.IsEmpty() || !stringQueue.IsEmpty() || !boolQueue.IsEmpty() {
			t.Errorf("Expected new queues to be empty, but they're not")
		}
	})
}

func TestArrayQueue_Enqueue(t *testing.T) {
	t.Run("Enqueue into empty queue", func(t *testing.T) {
		q := newArrayQueue[int](5)
		q.Enqueue(10)

		if q.Size() != 1 {
			t.Errorf("Expected size 1, got %d", q.Size())
		}

		if q.Front() != 10 {
			t.Errorf("Expected front to be 10, got %d", q.Front())
		}
	})

	t.Run("Enqueue multiple elements", func(t *testing.T) {
		q := newArrayQueue[string](5)
		q.Enqueue("hello")
		q.Enqueue("world")

		if q.Size() != 2 {
			t.Errorf("Expected size 2, got %d", q.Size())
		}

		if q.Front() != "hello" {
			t.Errorf("Expected front to be \"hello\", got \"%s\"", q.Front())
		}
	})

	t.Run("Enqueue into full queue", func(t *testing.T) {
		q := newArrayQueue[int](3)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Enqueue(4) // Should not be added

		if q.Size() != 3 {
			t.Errorf("Expected size 3, got %d", q.Size())
		}

		if q.Front() != 1 {
			t.Errorf("Expected front to be 1, got %d", q.Front())
		}
	})
}

func TestArrayQueue_Dequeue(t *testing.T) {
	t.Run("Dequeue on empty queue", func(t *testing.T) {
		q := newArrayQueue[int](5)

		if !q.IsEmpty() {
			t.Fatal("Expected new queue to be empty")
		}

		q.Dequeue() // Should not panic

		if !q.IsEmpty() {
			t.Errorf("Expected queue to remain empty after Dequeue on empty queue")
		}
	})

	t.Run("Dequeue on non-empty queue", func(t *testing.T) {
		q := newArrayQueue[int](5)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		q.Dequeue()

		if q.Size() != 2 {
			t.Errorf("Expected size 2, got %d", q.Size())
		}

		if q.Front() != 2 {
			t.Errorf("Expected front to be 2, got %d", q.Front())
		}
	})

	t.Run("Dequeue until empty", func(t *testing.T) {
		q := newArrayQueue[string](5)
		q.Enqueue("one")
		q.Enqueue("two")
		q.Dequeue()
		q.Dequeue()

		if !q.IsEmpty() {
			t.Errorf("Expected queue to be empty after dequeuing all elements")
		}

		if q.Size() != 0 {
			t.Errorf("Expected size 0, got %d", q.Size())
		}
	})
}

func TestArrayQueue_Front(t *testing.T) {
	t.Run("Front on empty queue", func(t *testing.T) {
		q := newArrayQueue[int](5)

		if !q.IsEmpty() {
			t.Fatal("Expected new queue to be empty")
		}

		if q.Front() != 0 {
			t.Errorf("Expected front of empty queue to return zero value, got %v", q.Front())
		}
	})

	t.Run("Front on non-empty queue", func(t *testing.T) {
		q := newArrayQueue[string](5)
		q.Enqueue("hello")
		q.Enqueue("world")

		if q.Front() != "hello" {
			t.Errorf("Expected front to be \"hello\", got \"%s\"", q.Front())
		}

		if q.Size() != 2 {
			t.Errorf("Expected size to remain 2 after calling Front(), got %d", q.Size())
		}
	})
}

func TestArrayQueue_IsEmpty(t *testing.T) {
	t.Run("Empty queue", func(t *testing.T) {
		q := newArrayQueue[int](5)

		if !q.IsEmpty() {
			t.Errorf("Expected new queue to be empty, but it's not")
		}
	})

	t.Run("Non-empty queue", func(t *testing.T) {
		q := newArrayQueue[string](5)
		q.Enqueue("hello")

		if q.IsEmpty() {
			t.Errorf("Expected queue to be non-empty after enqueuing an element, but it's empty")
		}
	})

	t.Run("Empty after Enqueue and Dequeue", func(t *testing.T) {
		q := newArrayQueue[float64](5)
		q.Enqueue(1.23)
		q.Dequeue()

		if !q.IsEmpty() {
			t.Errorf("Expected queue to be empty after enqueuing and then dequeuing an element, but it's not")
		}
	})
}

func TestArrayQueue_Size(t *testing.T) {
	t.Run("Size of empty queue", func(t *testing.T) {
		q := newArrayQueue[int](5)

		if q.Size() != 0 {
			t.Errorf("Expected size of new queue to be 0, got %d", q.Size())
		}
	})

	t.Run("Size after Enqueue", func(t *testing.T) {
		q := newArrayQueue[string](5)
		q.Enqueue("one")
		q.Enqueue("two")

		if q.Size() != 2 {
			t.Errorf("Expected size to be 2 after enqueuing two elements, got %d", q.Size())
		}
	})

	t.Run("Size after Dequeue", func(t *testing.T) {
		q := newArrayQueue[float64](5)
		q.Enqueue(1.1)
		q.Enqueue(2.2)
		q.Dequeue()

		if q.Size() != 1 {
			t.Errorf("Expected size to be 1 after enqueuing two elements and dequeuing one, got %d", q.Size())
		}
	})

	t.Run("Size after multiple Enqueue and Dequeue", func(t *testing.T) {
		q := newArrayQueue[int](5)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Dequeue()
		q.Enqueue(3)
		q.Dequeue()
		q.Enqueue(4)

		if q.Size() != 2 {
			t.Errorf("Expected size to be 2 after multiple enqueues and dequeues, got %d", q.Size())
		}
	})
}

func TestArrayQueue_String(t *testing.T) {
	t.Run("String on empty queue", func(t *testing.T) {
		q := newArrayQueue[int](5)

		if q.String() != "" {
			t.Errorf("Expected String() of an empty queue to be \"\", got \"%s\"", q.String())
		}
	})

	t.Run("String on queue with one element", func(t *testing.T) {
		q := newArrayQueue[string](5)
		q.Enqueue("hello")

		expected := "hello"
		if q.String() != expected {
			t.Errorf("Expected String() to be \"%s\", got \"%s\"", expected, q.String())
		}
	})

	t.Run("String on queue with multiple elements", func(t *testing.T) {
		q := newArrayQueue[int](5)
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		expected := "1 -> 2 -> 3"
		if q.String() != expected {
			t.Errorf("Expected String() to be \"%s\", got \"%s\"", expected, q.String())
		}
	})

	t.Run("String after Enqueue and Dequeue", func(t *testing.T) {
		q := newArrayQueue[float64](5)
		q.Enqueue(1.1)
		q.Enqueue(2.2)
		q.Dequeue()
		q.Enqueue(3.3)

		expected := "2.2 -> 3.3"
		if q.String() != expected {
			t.Errorf("Expected String() to be \"%s\", got \"%s\"", expected, q.String())
		}
	})
}
