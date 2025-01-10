package main

import "fmt"

type LibraryStats struct {
	BooksBorrowed, BooksReturned, BooksNotReturned int
}

type BookTransaction struct {
	BookTitle  string
	IsReturned bool
}

func updateLibraryStats(stats *LibraryStats, transaction BookTransaction) {

	if transaction.IsReturned {
		stats.BooksBorrowed++
		stats.BooksReturned++
	} else {
		stats.BooksBorrowed++
		stats.BooksNotReturned++
	}
}

func main() {
	stats := LibraryStats{
		BooksBorrowed:    10,
		BooksReturned:    5,
		BooksNotReturned: 5,
	}

	transaction1 := BookTransaction{
		BookTitle:  "The Hitchhiker's Guide to the Galaxy",
		IsReturned: true,
	}

	transaction2 := BookTransaction{
		BookTitle:  "Pride and Prejudice",
		IsReturned: false,
	}

	fmt.Println("Initial Stats:", stats)

	updateLibraryStats(&stats, transaction1)
	fmt.Println("Stats after transaction 1 (returned):", stats)

	updateLibraryStats(&stats, transaction2)
	fmt.Println("Stats after transaction 2 (not returned):", stats)
}
