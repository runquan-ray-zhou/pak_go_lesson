package main

import "fmt"

type Book struct {
	Title  string
	Author string
	ISBN   string
}

func (b *Book) updateTitle(newTitle string) {
	b.Title = newTitle
}

func main() {

	book := &Book{
		Title:  "The Lord of the Rings",
		Author: "J.R.R. Tolkien",
		ISBN:   "978-0-618-05398-2",
	}

	fmt.Println("Original Title:", book.Title)

	book.updateTitle("The Hobbit")

	fmt.Println("Updated Title:", book.Title)
}

/*

# Go Pointers Exercise: Update Book Title

This exercise focuses on using pointer receivers in Go methods to modify struct fields.

## Scenario

You're working on a library management system. You have a `Book` struct that represents a book in the library:

## Task

you need to implement the updateTitle method, which is associated with the Book struct. This method should take a new title as a string argument and update the Title field of the Book struct.

Verify: Run the tests again using go test -v. They should now all pass.


*/
