# Go Interfaces Exercise: Payment Gateway

This exercise focuses on using interfaces in Go to model different payment methods and process payments using polymorphism.

## Scenario

You are developing a payment gateway that can process payments from different payment methods:

- Credit Card
- PayPal
- Bank Transfer

Each payment method has its own way of processing payments and generating receipts.

Expected output
Processing credit card payment of $10.00 with card ending in 3456
Receipt:
Credit Card Payment
Amount: $10.00
Card Number: 1234567890123456
Expiry Date: 12/24

Processing PayPal payment of $20.00 for email example@example.com
Receipt:
PayPal Payment
Amount: $20.00
Email: example@example.com

Processing bank transfer of $5.00 for account 1234567890
Receipt:
Bank Transfer Payment
Amount: $5.00
Account Number: 1234567890
Routing Number: 123456789

Verify: Run the tests again using go test -v. They should now all pass.
