package stub

import "github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain"

// NewActiveAccountStub creates a stub for an active individual account.
func NewActiveAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 100.0, // Positive balance
		},
		Status: domain.AccountStatusActive,   // Updated status naming
		Type:   domain.AccountTypeIndividual, // Assuming active accounts are individual by default
	}
}

// NewTrialAccountStub creates a stub for a trial account.
func NewTrialAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 0.0, // Zero balance
		},
		Status: domain.AccountStatusTrial,    // Updated status naming
		Type:   domain.AccountTypeIndividual, // Assuming trial accounts are individual by default
	}
}

// NewCancelledAccountStub creates a stub for a cancelled account.
func NewCancelledAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 0.0, // Zero balance
		},
		Status: domain.AccountStatusCancelled, // Updated status naming
		Type:   domain.AccountTypeIndividual,  // Assuming cancelled accounts are individual by default
	}
}

// NewLegacyAccountStub creates a stub for a legacy account.
func NewLegacyAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 200.0, // Positive balance
		},
		Status: domain.AccountStatusActive, // Legacy accounts are typically active
		Type:   domain.AccountTypeLegacy,   // Account type set to legacy
	}
}

// NewUnknownAccountStub creates a stub for an unknown account type.
func NewUnknownAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 50.0, // Neutral balance
		},
		Status: domain.AccountStatusUnknown, // Updated status naming
		Type:   domain.AccountTypeUnknown,   // Type is unknown
	}
}

// NewCustomAccountStub allows creating a custom account stub with specified properties.
func NewCustomAccountStub(number string, balance float64, status domain.AccountStatus, accountType domain.AccountType) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: balance,
		},
		Status: status,
		Type:   accountType,
	}
}
