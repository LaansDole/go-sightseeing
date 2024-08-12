package domain

import (
	"fmt"
	"strings"
)

type AccountNumber string

type AccountBalances struct {
	CurrentBalance float64
}

// ExtractAccountNumberFrom
//
// TODO: This is a refactor from:
// NewAccountNumberFromResourceName
func ExtractAccountNumberFrom(resourceName string) AccountNumber {
	accountPrefix := "accounts/"
	return AccountNumber(strings.TrimPrefix(resourceName, accountPrefix))
}

func (acc AccountNumber) ResourceName() string {
	accountPrefix := "accounts/"
	return fmt.Sprintf("%s%s", accountPrefix, acc)
}
