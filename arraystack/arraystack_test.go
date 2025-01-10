package main

import (
	"testing"
)

// Helper function to create a new array-based stack for each test.
func newArrayStack[T any]() *ArrayStack[T] {
	return NewArrayStack[T]()
}

func TestArrayStack_NewArrayStack(t *testing.T) {
	t.Run("New stack is empty", func(t *testing.T) {
		stack := newArrayStack[int]()

		if stack == nil {
			t.Fatal("NewArrayStack returned nil")
		}

		if !stack.Empty() {
			t.Errorf("Expected new stack to be empty, but it's not")
		}

		if stack.Size() != 0 {
			t.Errorf("Expected new stack to have size 0, but got %d", stack.Size())
		}
	})

	t.Run("New stack with different types", func(t *testing.T) {
		intStack := newArrayStack[int]()
		stringStack := newArrayStack[string]()
		boolStack := newArrayStack[bool]()

		if intStack == nil || stringStack == nil || boolStack == nil {
			t.Fatal("NewArrayStack returned nil for some type")
		}

		if !intStack.Empty() || !stringStack.Empty() || !boolStack.Empty() {
			t.Errorf("Expected new stacks to be empty, but they're not")
		}
	})
}

func TestArrayStack_Top(t *testing.T) {
	t.Run("Top on empty stack returns zero value", func(t *testing.T) {
		stack := newArrayStack[int]()

		if !stack.Empty() {
			t.Fatal("Expected new stack to be empty")
		}

		// Test with different types to check for their zero values
		intVal := newArrayStack[int]().Top()
		if intVal != 0 {
			t.Errorf("Expected top of empty int stack to be 0, got %d", intVal)
		}

		stringVal := newArrayStack[string]().Top()
		if stringVal != "" {
			t.Errorf("Expected top of empty string stack to be \"\", got \"%s\"", stringVal)
		}

		boolVal := newArrayStack[bool]().Top()
		if boolVal != false {
			t.Errorf("Expected top of empty bool stack to be false, got %t", boolVal)
		}
	})

	t.Run("Top on non-empty stack returns correct element", func(t *testing.T) {
		stack := newArrayStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		if stack.Size() != 3 {
			t.Fatalf("Expected stack size to be 3, got %d", stack.Size())
		}

		topVal := stack.Top()
		if topVal != 3 {
			t.Errorf("Expected top element to be 3, got %d", topVal)
		}

		if stack.Size() != 3 {
			t.Errorf("Expected stack size to still be 3 after calling Top(), got %d", stack.Size())
		}
	})
}

func TestArrayStack_Push(t *testing.T) {
	t.Run("Push to empty stack", func(t *testing.T) {
		stack := newArrayStack[int]()

		if !stack.Empty() {
			t.Fatal("Expected new stack to be empty")
		}

		stack.Push(1)

		if stack.Empty() {
			t.Errorf("Expected stack to be non-empty after Push")
		}

		if stack.Size() != 1 {
			t.Errorf("Expected stack size to be 1, got %d", stack.Size())
		}

		if stack.Top() != 1 {
			t.Errorf("Expected top element to be 1, got %v", stack.Top())
		}
	})

	t.Run("Push multiple elements", func(t *testing.T) {
		stack := newArrayStack[string]()
		stack.Push("hello")
		stack.Push("world")
		stack.Push("!")

		if stack.Size() != 3 {
			t.Errorf("Expected stack size to be 3, got %d", stack.Size())
		}

		if stack.Top() != "!" {
			t.Errorf("Expected top element to be \"!\", got \"%v\"", stack.Top())
		}
	})

	t.Run("Push different types", func(t *testing.T) {
		intStack := newArrayStack[int]()
		intStack.Push(10)
		intStack.Push(20)

		stringStack := newArrayStack[string]()
		stringStack.Push("one")
		stringStack.Push("two")

		if intStack.Top() != 20 {
			t.Errorf("Expected top of int stack to be 20, got %v", intStack.Top())
		}

		if stringStack.Top() != "two" {
			t.Errorf("Expected top of string stack to be \"two\", got \"%v\"", stringStack.Top())
		}
	})
}

func TestArrayStack_Pop(t *testing.T) {
	t.Run("Pop on empty stack", func(t *testing.T) {
		stack := newArrayStack[int]()

		if !stack.Empty() {
			t.Fatal("Expected new stack to be empty")
		}

		stack.Pop() // Should not panic

		if !stack.Empty() {
			t.Errorf("Expected stack to remain empty after Pop on empty stack")
		}
	})

	t.Run("Pop on non-empty stack", func(t *testing.T) {
		stack := newArrayStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		if stack.Size() != 3 {
			t.Fatalf("Expected stack size to be 3, got %d", stack.Size())
		}

		stack.Pop()

		if stack.Size() != 2 {
			t.Errorf("Expected stack size to be 2 after Pop, got %d", stack.Size())
		}

		if stack.Top() != 2 {
			t.Errorf("Expected top element to be 2 after Pop, got %d", stack.Top())
		}
	})

	t.Run("Pop until empty", func(t *testing.T) {
		stack := newArrayStack[string]()
		stack.Push("one")
		stack.Push("two")
		stack.Push("three")

		stack.Pop()
		stack.Pop()
		stack.Pop()

		if !stack.Empty() {
			t.Errorf("Expected stack to be empty after popping all elements")
		}

		if stack.Size() != 0 {
			t.Errorf("Expected stack size to be 0 after popping all elements, got %d", stack.Size())
		}
	})
}

func TestArrayStack_Empty(t *testing.T) {
	t.Run("Empty stack", func(t *testing.T) {
		stack := newArrayStack[int]()

		if !stack.Empty() {
			t.Errorf("Expected new stack to be empty, but it's not")
		}
	})

	t.Run("Non-empty stack", func(t *testing.T) {
		stack := newArrayStack[string]()
		stack.Push("hello")

		if stack.Empty() {
			t.Errorf("Expected stack to be non-empty after pushing an element, but it's empty")
		}
	})

	t.Run("Empty after Push and Pop", func(t *testing.T) {
		stack := newArrayStack[float64]()
		stack.Push(1.23)
		stack.Pop()

		if !stack.Empty() {
			t.Errorf("Expected stack to be empty after pushing and then popping an element, but it's not")
		}
	})
}

func TestArrayStack_Size(t *testing.T) {
	t.Run("Size of empty stack", func(t *testing.T) {
		stack := newArrayStack[int]()

		if stack.Size() != 0 {
			t.Errorf("Expected size of new stack to be 0, got %d", stack.Size())
		}
	})

	t.Run("Size after Push", func(t *testing.T) {
		stack := newArrayStack[string]()
		stack.Push("one")
		stack.Push("two")

		if stack.Size() != 2 {
			t.Errorf("Expected stack size to be 2 after pushing two elements, got %d", stack.Size())
		}
	})

	t.Run("Size after Pop", func(t *testing.T) {
		stack := newArrayStack[float64]()
		stack.Push(1.1)
		stack.Push(2.2)
		stack.Pop()

		if stack.Size() != 1 {
			t.Errorf("Expected stack size to be 1 after pushing two elements and popping one, got %d", stack.Size())
		}
	})

	t.Run("Size after multiple Push and Pop", func(t *testing.T) {
		stack := newArrayStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Pop()
		stack.Push(3)
		stack.Pop()
		stack.Push(4)

		if stack.Size() != 2 {
			t.Errorf("Expected stack size to be 2 after multiple pushes and pops, got %d", stack.Size())
		}
	})
}

func TestArrayStack_String(t *testing.T) {
	t.Run("String on empty stack", func(t *testing.T) {
		stack := newArrayStack[int]()

		if stack.String() != "" {
			t.Errorf("Expected String() of an empty stack to be \"\", got \"%s\"", stack.String())
		}
	})

	t.Run("String on stack with one element", func(t *testing.T) {
		stack := newArrayStack[string]()
		stack.Push("hello")

		expected := "hello"
		if stack.String() != expected {
			t.Errorf("Expected String() to be \"%s\", got \"%s\"", expected, stack.String())
		}
	})

	t.Run("String on stack with multiple elements", func(t *testing.T) {
		stack := newArrayStack[int]()
		stack.Push(1)
		stack.Push(2)
		stack.Push(3)

		expected := "3 -> 2 -> 1"
		if stack.String() != expected {
			t.Errorf("Expected String() to be \"%s\", got \"%s\"", expected, stack.String())
		}
	})

	t.Run("String after Push and Pop", func(t *testing.T) {
		stack := newArrayStack[float64]()
		stack.Push(1.1)
		stack.Push(2.2)
		stack.Pop()
		stack.Push(3.3)

		expected := "3.3 -> 1.1"
		if stack.String() != expected {
			t.Errorf("Expected String() to be \"%s\", got \"%s\"", expected, stack.String())
		}
	})
}
