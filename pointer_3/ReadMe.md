# Go Pointers Exercise: Update Library Statistics

This exercise focuses on using pointers in Go to modify struct fields within a function.

## Scenario

You are developing a system for tracking library statistics. You have two structs:

- **`LibraryStats`:** Stores the overall statistics for the library, including the number of books borrowed, returned, and not returned.
- **`BookTransaction`:** Represents a single transaction involving a book, indicating the book's title and whether it was returned or not.

Your task is to implement the `updateLibraryStats` function.

Run the Tests: Run the tests using go test -v in your terminal. You will see that some tests fail. Analyze the output to understand why they are failing.

Expect output
Initial Stats: {10 5 5}
Stats after transaction 1 (returned): {11 6 5}
Stats after transaction 2 (not returned): {12 6 6}
