package main

import "fmt"

type Point struct {
	X int
	Y int
}

func (p *Point) MoveToOrigin() {
	p.X = 0
	p.Y = 0
}

func main() {

	p := Point{X: 1, Y: 2}
	ptr := &p

	fmt.Printf("prt (memory address): %p\n", ptr)
	fmt.Println("*ptr: ", *ptr)
	fmt.Println("*ptr: ", ptr)

	p.MoveToOrigin()
	fmt.Println(p)
}
