package main

type StackInterface[T any] interface {
	Top() T
	Push(val T)
	Pop()
	Empty() bool
	Size() int
	String() string
}

type ArrayStack[T any] struct {
	arr []T
}

func NewArrayStack[T any]() *ArrayStack[T] {
	return &ArrayStack[T]{arr: []T{}}
}

func (s *ArrayStack[T]) Top() T {

}

func (s *ArrayStack[T]) Push(val T) {

}

func (s *ArrayStack[T]) Pop() {

}

func (s *ArrayStack[T]) Empty() bool {

}

func (s *ArrayStack[T]) Size() int {

}

func (s *ArrayStack[T]) String() string {

}
