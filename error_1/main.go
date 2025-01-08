package main

import "fmt"

func borrowBooks(userID string, books []string, accountLimit int) (int, error) {

	totalBorrowed := len(books)

	if totalBorrowed > accountLimit {
		return 0, fmt.Errorf("user %s exceeded borrowing limit: allowed %d, attempted %d", userID, accountLimit, totalBorrowed)
	}

	for _, book := range books {
		fmt.Printf("User %s borrowed books: %s\n", userID, book)
	}

	return totalBorrowed, nil
}

func main() {

	users := map[string]struct {
		books        []string
		accountLimit int
	}{
		"user1": {books: []string{"Book A", "Book B"}, accountLimit: 3},
		"user2": {books: []string{"Book X", "Book Y", "Book Z", "Book W"}, accountLimit: 2},
	}

	for userID, user := range users {
		fmt.Printf("Processing borrow request for %s...\n", userID)

		totalBorrowed, err := borrowBooks(userID, user.books, user.accountLimit)

		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("User %s successfully borrowed %d books\n", userID, totalBorrowed)
		}

		fmt.Println()
	}
}
