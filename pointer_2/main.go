package main

import (
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Email *string
}

// validating if the email is equal to nil, if nil will return nil

func validateEmail(user *User) {

	// at := false

	if user.Email == nil {
		return
	}

	trimmedEmail := strings.TrimSpace(*user.Email)

	if !strings.Contains(trimmedEmail, "@") {
		user.Email = nil
	}

	*user.Email = trimmedEmail
}

func main() {
	email := "test@example.com"
	user := User{Name: "John Doe", Email: &email}

	validateEmail(&user)

	fmt.Println("User Name:", user.Name)
	if user.Email != nil {
		fmt.Println("User Email:", *user.Email)
	} else {
		fmt.Println("User Email: Invalid")
	}
}

/*

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

*/
