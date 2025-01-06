package main

import (
	"fmt"
	"testing"
)

func TestUpdateTitle(t *testing.T) {
	type testCase struct {
		book     Book
		newTitle string
		expected string
	}

	tests := []testCase{
		{
			book: Book{
				Title:  "The Hitchhiker's Guide to the Galaxy",
				Author: "Douglas Adams",
				ISBN:   "978-0345391803",
			},
			newTitle: "The Restaurant at the End of the Universe",
			expected: "The Restaurant at the End of the Universe",
		},
		{
			book: Book{
				Title:  "Pride and Prejudice",
				Author: "Jane Austen",
				ISBN:   "978-0141439518",
			},
			newTitle: "Sense and Sensibility",
			expected: "Sense and Sensibility",
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{
				book: Book{
					Title:  "The Go Programming Language",
					Author: "Alan A. A. Donovan, Brian W. Kernighan",
					ISBN:   "978-0134190440",
				},
				newTitle: "Go in Practice",
				expected: "Go in Practice",
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		originalTitle := test.book.Title
		test.book.updateTitle(test.newTitle)
		if test.book.Title != test.expected {
			failCount++
			t.Errorf(`Test Failed:
Input:
Title: %v
Author: %v
ISBN: %v
New Title: %v
Expected: %v
Actual: %v
`, originalTitle, test.book.Author, test.book.ISBN, test.newTitle, test.expected, test.book.Title)
		} else {
			passCount++
			fmt.Printf(`Test Passed:
Input:
Title: %v
Author: %v
ISBN: %v
New Title: %v
Expected: %v
Actual: %v
`, originalTitle, test.book.Author, test.book.ISBN, test.newTitle, test.expected, test.book.Title)
		}
		fmt.Println("------------------------------")
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
