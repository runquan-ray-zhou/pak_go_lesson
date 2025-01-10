package main

import (
	"fmt"
	"testing"
)

func TestUpdateLibraryStats(t *testing.T) {
	type testCase struct {
		initialStats   LibraryStats
		newTransaction BookTransaction
		expected       LibraryStats
	}

	tests := []testCase{
		{
			initialStats:   LibraryStats{BooksBorrowed: 0, BooksReturned: 0, BooksNotReturned: 0},
			newTransaction: BookTransaction{BookTitle: "The Great Gatsby", IsReturned: true},
			expected:       LibraryStats{BooksBorrowed: 1, BooksReturned: 1, BooksNotReturned: 0},
		},
		{
			initialStats:   LibraryStats{BooksBorrowed: 1, BooksReturned: 1, BooksNotReturned: 0},
			newTransaction: BookTransaction{BookTitle: "1984", IsReturned: false},
			expected:       LibraryStats{BooksBorrowed: 2, BooksReturned: 1, BooksNotReturned: 1},
		},
	}

	if withSubmit {
		tests = append(tests, []testCase{
			{
				initialStats:   LibraryStats{BooksBorrowed: 2, BooksReturned: 1, BooksNotReturned: 1},
				newTransaction: BookTransaction{BookTitle: "To Kill a Mockingbird", IsReturned: false},
				expected:       LibraryStats{BooksBorrowed: 3, BooksReturned: 1, BooksNotReturned: 2},
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		stats := test.initialStats
		updateLibraryStats(&stats, test.newTransaction)
		if stats != test.expected {
			failCount++
			t.Errorf(`
---------------------------------
Test Failed:
  Initial Stats:
    BooksBorrowed=%d, BooksReturned=%d, BooksNotReturned=%d
  New Transaction:
    BookTitle=%s, IsReturned=%v
  Expected:
    BooksBorrowed=%d, BooksReturned=%d, BooksNotReturned=%d
  Actual:
    BooksBorrowed=%d, BooksReturned=%d, BooksNotReturned=%d
`, test.initialStats.BooksBorrowed, test.initialStats.BooksReturned, test.initialStats.BooksNotReturned,
				test.newTransaction.BookTitle, test.newTransaction.IsReturned,
				test.expected.BooksBorrowed, test.expected.BooksReturned, test.expected.BooksNotReturned,
				stats.BooksBorrowed, stats.BooksReturned, stats.BooksNotReturned)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  Initial Stats:
    BooksBorrowed=%d, BooksReturned=%d, BooksNotReturned=%d
  New Transaction:
    BookTitle=%s, IsReturned=%v
  Expected:
    BooksBorrowed=%d, BooksReturned=%d, BooksNotReturned=%d
  Actual:
    BooksBorrowed=%d, BooksReturned=%d, BooksNotReturned=%d
`, test.initialStats.BooksBorrowed, test.initialStats.BooksReturned, test.initialStats.BooksNotReturned,
				test.newTransaction.BookTitle, test.newTransaction.IsReturned,
				test.expected.BooksBorrowed, test.expected.BooksReturned, test.expected.BooksNotReturned,
				stats.BooksBorrowed, stats.BooksReturned, stats.BooksNotReturned)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
