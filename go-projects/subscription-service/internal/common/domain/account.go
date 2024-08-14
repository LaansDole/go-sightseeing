package domain

import (
	"fmt"
	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/thirdparty"
	"strings"
)

type AccountNumber string

// Account interface representing a generic account
type Account interface {
	GetAccountNumber() AccountNumber
	GetBalance() float64
	GetType() string
	IsValidForCancellation() bool
}

// BaseAccount struct to be embedded in other account types
type BaseAccount struct {
	AccountNumber AccountNumber
	Balance       float64
}

// GetAccountNumber Implement common methods for BaseAccount
func (acc BaseAccount) GetAccountNumber() AccountNumber {
	return acc.AccountNumber
}

func (acc BaseAccount) GetBalance() float64 {
	// Fetch balance from the third-party service
	balance, err := thirdparty.GetAccountBalance(string(acc.AccountNumber))
	if err != nil {
		return 0.00 // Handle the error by returning a default balance
	}
	return balance
}

func (acc BaseAccount) IsValidForCancellation() bool {
	balance := acc.GetBalance()
	// For this example, let's say the account is only valid for cancellation if the balance is non-negative
	if balance >= 0 {
		return true
	}
	return false
}

// ActiveAccount struct representing an active account
type ActiveAccount struct {
	BaseAccount
}

func (acc ActiveAccount) GetType() string {
	return "active"
}

func (acc ActiveAccount) IsValidForCancellation() bool {
	// Add logic specific to ActiveAccount cancellation
	return true
}

// TrialAccount struct representing a trial account
type TrialAccount struct {
	BaseAccount
}

func (acc TrialAccount) GetType() string {
	return "trial"
}

func (acc TrialAccount) IsValidForCancellation() bool {
	// Add logic specific to TrialAccount cancellation
	return true
}

// SuspendedAccount struct representing a suspended account
type SuspendedAccount struct {
	BaseAccount
}

func (acc SuspendedAccount) GetType() string {
	return "suspended"
}

func (acc SuspendedAccount) IsValidForCancellation() bool {
	// Add logic specific to SuspendedAccount cancellation
	return false
}

// CancelledAccount struct representing an account that has been cancelled but is still active within a grace period
type CancelledAccount struct {
	BaseAccount
	GracePeriodEnd string // Example field for when the service actually stops
}

func (acc CancelledAccount) GetType() string {
	return "cancelled"
}

func (acc CancelledAccount) IsValidForCancellation() bool {
	// Generally, a cancelled account is already in the process of being canceled
	// But you might have additional logic here
	return false
}

// LegacyAccount struct representing a legacy account
type LegacyAccount struct {
	BaseAccount
}

func (acc LegacyAccount) GetType() string {
	return "legacy"
}

func (acc LegacyAccount) IsValidForCancellation() bool {
	// Add logic specific to LegacyAccount cancellation
	return true
}

// FamilyOrEnterpriseAccount struct representing a family or enterprise account
type FamilyOrEnterpriseAccount struct {
	BaseAccount
	Members int // Example field for number of members/users
}

func (acc FamilyOrEnterpriseAccount) GetType() string {
	return "family_or_enterprise"
}

func (acc FamilyOrEnterpriseAccount) IsValidForCancellation() bool {
	// Add logic specific to FamilyOrEnterpriseAccount cancellation
	return true
}

// ExtractAccountNumberFrom
//
// TODO: This is a refactor from:
// NewAccountNumberFromResourceName()
func ExtractAccountNumberFrom(resourceName string) AccountNumber {
	accountPrefix := "accounts/"
	return AccountNumber(strings.TrimPrefix(resourceName, accountPrefix))
}

func (acc AccountNumber) ResourceName() string {
	accountPrefix := "accounts/"
	return fmt.Sprintf("%s%s", accountPrefix, acc)
}

// IsEqual
//
// TODO: This is a refactor from:
// Is
func (acc AccountNumber) IsEqual(accountNumber string) bool {
	return string(acc) == accountNumber
}
