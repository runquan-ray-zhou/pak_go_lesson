package main

import (
	"reflect"
	"testing"
)

func TestUpdatePerson(t *testing.T) {
	tests := []struct {
		input    Person
		expected Person
	}{
		{Person{Name: "John", Age: 25}, Person{Name: "Updated", Age: 26}},
		{Person{Name: "Alice", Age: 30}, Person{Name: "Updated", Age: 31}},
	}

	for _, test := range tests {
		p := test.input
		updatePerson(&p) // Pass by reference

		if !reflect.DeepEqual(p, test.expected) {
			t.Errorf("Expected %+v, got %+v", test.expected, p)
		}
	}
}
