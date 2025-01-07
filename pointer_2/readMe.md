# Go Pointers Exercise: Email Validation

This exercise focuses on using pointers in Go to modify struct fields and handle optional values.

## Scenario

You are working on a user management system. You have a `User` struct that stores a user's name and email address. There's a possibility the email can be `nil` so make sure you safe guard it.

You need to implement the `validateEmail` function, which takes a pointer to a `User` struct.

Check the main_test.go file regarding the test cases

Run the Tests: Run the tests using go test -v in your terminal. You will see that some tests fail. Analyze the output to understand why they are failing.

Expect output
User Name: John Doe
User Email: test@example.com
