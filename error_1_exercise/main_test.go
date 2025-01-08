package main

import (
	"testing"
)

func TestReturnBooks(t *testing.T) {
	testCases := []struct {
		name            string
		userID          string
		books           []string
		currentBorrowed int
		expectedCount   int
		expectError     bool
	}{
		{
			name:            "Valid return - some books",
			userID:          "user1",
			books:           []string{"Book A", "Book B"},
			currentBorrowed: 3,
			expectedCount:   1,
			expectError:     false,
		},
		{
			name:            "Valid return - all books",
			userID:          "user2",
			books:           []string{"Book X", "Book Y", "Book Z"},
			currentBorrowed: 3,
			expectedCount:   0,
			expectError:     false,
		},
		{
			name:            "Invalid return - more than borrowed",
			userID:          "user3",
			books:           []string{"Book A", "Book B", "Book C"},
			currentBorrowed: 2,
			expectedCount:   0,
			expectError:     true,
		},
		{
			name:            "Valid return - no books",
			userID:          "user4",
			books:           []string{},
			currentBorrowed: 5,
			expectedCount:   5,
			expectError:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			newCount, err := ReturnBooks(tc.userID, tc.books, tc.currentBorrowed)

			if tc.expectError && err == nil {
				t.Errorf("Expected an error but got none")
			}

			if !tc.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if newCount != tc.expectedCount {
				t.Errorf("Expected new borrowed count to be %d, but got %d", tc.expectedCount, newCount)
			}
		})
	}
}
