package main

import (
	"fmt"
)

func ReturnBooks(userID string, books []string, borrowed int) (int, error) {

	if len(books) > borrowed {
		return 0, fmt.Errorf("Invalid return - more than borrowed")
	}

	newBorrowedAmount := borrowed - len(books)

	return newBorrowedAmount, nil

}

func main() {

	users := map[string]struct {
		books           []string
		currentBorrowed int
	}{
		"user1": {books: []string{"Book A", "Book B"}, currentBorrowed: 3},
		"user2": {books: []string{"Book X", "Book Y", "Book Z"}, currentBorrowed: 2},
	}

	// Process each user
	for userID, user := range users {
		fmt.Printf("Processing return request for %s...\n", userID)
		newBorrowedCount, err := ReturnBooks(userID, user.books, user.currentBorrowed)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("User %s successfully returned books. New borrowed count: %d\n", userID, newBorrowedCount)
		}
		fmt.Println()
	}
}

/*

## Library Book Return System Exercise

## Overview

In this exercise, you'll be working on a library management system, specifically focusing on the book return process. Your task is to implement the ReturnBooks function, which handles the return of books by library users.

## Requirements

The function should process the return of books for a given user.
It should update the number of books currently borrowed by the user.
The function should handle error cases, such as when a user tries to return more books than they have borrowed.
Proper logging of returned books is required.

Verify: Run the tests again using go test -v. They should now all pass.

*/
