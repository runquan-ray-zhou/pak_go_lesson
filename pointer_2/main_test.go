package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		name     string
		email    *string
		expected *string
	}

	email1 := "test@example.com"
	email2 := "invalid-email"
	email3 := " another@example.com "
	email4 := ""

	tests := []testCase{
		{
			name:     "Valid Email",
			email:    &email1,
			expected: &email1,
		},
		{
			name:     "Invalid Email",
			email:    &email2,
			expected: nil,
		},
		{
			name:     "Valid Email with Spaces",
			email:    &email3,
			expected: stringPtr("another@example.com"),
		},
		{
			name:     "Empty String",
			email:    &email4,
			expected: nil,
		},
		{
			name:     "Nil Email",
			email:    nil,
			expected: nil,
		},
	}

	if withSubmit {
		email5 := "  valid@example.net  "
		tests = append(tests, []testCase{
			{
				name:     "Another valid email with spaces",
				email:    &email5,
				expected: stringPtr("valid@example.net"),
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		user := User{Name: "Test User", Email: test.email}
		validateEmail(&user)

		if !compareEmail(user.Email, test.expected) {
			failCount++
			t.Errorf(`Test Failed: %s
Input: %v
Expected: %v
Actual: %v
`,
				test.name,
				pointerToString(test.email),
				pointerToString(test.expected),
				pointerToString(user.Email),
			)
		} else {
			passCount++
			fmt.Printf(`Test Passed: %s
Input: %v
Expected: %v
Actual: %v
`,
				test.name,
				pointerToString(test.email),
				pointerToString(test.expected),
				pointerToString(user.Email),
			)
		}
		fmt.Println("------------------------------")
	}

	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func compareEmail(actual, expected *string) bool {
	if actual == nil && expected == nil {
		return true
	}
	if actual == nil || expected == nil {
		return false
	}
	return *actual == *expected
}

func pointerToString(ptr *string) string {
	if ptr == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", *ptr)
}

func stringPtr(val string) *string {
	return &val
}

var withSubmit = true

func TestValidateEmail_Empty(t *testing.T) {
	email := "test@example.com"
	originalEmail := email
	user := User{Name: "Test User", Email: &email}

	validateEmail(&user)

	fmt.Println("After validateEmail: user.Email:", user.Email, "value:", *user.Email)

	if user.Email == nil || *user.Email != originalEmail {
		t.Errorf("Expected email to be unchanged and not nil, but it was modified to %v", user.Email)
	}
}
