Exercise: Update a Struct Field Using Pass-by-Reference

Objective:
Modify the fields of a struct using pass-by-reference with pointers.

Instructions:
Define a Person struct with the following fields:

    Name (string)
    Age (int)
Write a function updatePerson that:

    Takes a pointer to a Person struct as an argument.
    Updates the Name to "Updated" and increments the Age by 1.
    In the main function:

Create a Person instance with initial values for Name and Age.
    Pass the address of the Person instance to the updatePerson function.
    Print the Person struct before and after calling the function.

Expected output:
Before update: {Name: John, Age: 25}
After update: {Name: Updated, Age: 26}
