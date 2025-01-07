## TASK

Implement a generic function, filterSlice, that filters elements from a slice based on a condition defined by a function passed as an argument. The goal is to return a new slice containing only the elements that meet the specified condition.

Expected Behavior
After completing the filterSlice function:

It should work with slices of any type.
It should dynamically filter elements based on the provided keepItem function.

Run the Tests: Run the tests using go test -v in your terminal. You will see that some tests fail. Analyze the output to understand why they are failing.

Expected output
oddNumbers: [1 3 5 7 9]
filtered fruits [cherry fig]
