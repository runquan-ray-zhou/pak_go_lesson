# Go Error Handling Exercise: Updating Customer Balances

This exercise focuses on error handling in Go, specifically using custom error constants and handling different transaction types when updating customer balances.

## Scenario

You're working on a banking application. Complete the following variables and struct

- **`customer` struct:** Represents a customer with an ID and an account balance.
- **`transaction` struct:** Represents a transaction with a customer ID, amount, and transaction type (deposit or withdrawal).
  **Constants:**

You need to define the following constants to represent valid transaction types:

- **`transactionDeposit`:** Should have the string value `"deposit"`.
- **`transactionWithdrawal`:** Should have the string value `"withdrawal"`.

Expected output
Alice's balance after deposit: 150.00
Alice's balance after withdrawal: 75.00
Error: unknown transaction type
