Exercise: Bank Account Management
Objective:
Design a system with structs and methods to manage bank accounts and perform transactions between them.

Instructions:
Create a BankAccount struct with the following fields:
    AccountHolder (string)
    Balance (float64)

Implement the following methods for the BankAccount struct:
    Deposit(amount float64):
        Add amount to Balance.
        Return true if successful, false if amount is negative or zero.

    Withdraw(amount float64):
        Subtract amount from Balance if sufficient funds exist.
        Return true if successful, false if insufficient funds or amount is negative.

Create a Transfer method:
    It should take another BankAccount pointer and an amount as arguments.
    Withdraw the amount from the current account and deposit it into the other account if the transaction is valid.
    Return true if the transfer is successful, false otherwise.

Expected output:
Before Transfer: Alice has $1000.00, Bob has $500.00
Transfer successful!
After Transfer: Alice has $800.00, Bob has $700.00
