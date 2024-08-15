package domain

import (
	"fmt"
	"strings"
)

type AccountNumber string

// Account interface representing a generic account
type Account struct {
	AccountNumber AccountNumber
	Balances      AccountBalances
	Status        AccountStatus
	Type          AccountType
}

type AccountBalances struct {
	CurrentMoney float64
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

type Accounts []Account

func (accList Accounts) FilterLegacyAccount() Accounts {
	var legacyAccounts Accounts
	for _, acc := range accList {
		if acc.Status == AccountStatusActive && acc.IsLegacy() {
			legacyAccounts = append(legacyAccounts, acc)
		}
	}
	return legacyAccounts
}
