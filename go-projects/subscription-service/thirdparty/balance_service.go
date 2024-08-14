package thirdparty

import (
	"errors"
	"time"
)

// Simulated response delay to mimic real-world API calls
const responseDelay = 500 * time.Millisecond

// Simulated data: Map of account numbers to their corresponding balances
var accountBalances = map[string]float64{
	"123456789": 100.00, // Example account with a positive balance
	"987654321": -10.00, // Example account with a negative balance
	"555555555": 0.00,   // Example account with a zero balance
}

// GetAccountBalance simulates a third-party service that returns the balance for an account number
func GetAccountBalance(accountNumber string) (float64, error) {
	// Simulate response time
	time.Sleep(responseDelay)

	// Check if the account number exists in the simulated data
	if balance, exists := accountBalances[accountNumber]; exists {
		return balance, nil
	}

	// Return an error if the account number is not found
	return 0.00, errors.New("unable to retrieve balance: account number not found")
}
