package stub

import "github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain"

// NewActiveAccountStub creates a stub for an active account.
func NewActiveAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 100.0, // Positive balance
		},
		Status: domain.ActiveAccount,
	}
}

// NewTrialAccountStub creates a stub for a trial account.
func NewTrialAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 0.0, // Zero balance
		},
		Status: domain.TrialAccount,
	}
}

// NewCancelledAccountStub creates a stub for a cancelled account.
func NewCancelledAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 0.0, // Zero balance
		},
		Status: domain.CancelledAccount,
	}
}

// NewLegacyAccountStub creates a stub for a legacy account.
func NewLegacyAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 200.0, // Positive balance
		},
		Status: domain.LegacyAccount,
	}
}

// NewUnknownAccountStub creates a stub for an unknown account type.
func NewUnknownAccountStub(number string) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: 50.0, // Neutral balance
		},
		Status: domain.UnknownAccount,
	}
}

// NewCustomAccountStub allows creating a custom account stub with specified properties.
func NewCustomAccountStub(number string, balance float64, status domain.AccountStatus) domain.Account {
	return domain.Account{
		AccountNumber: domain.AccountNumber(number),
		Balances: domain.AccountBalances{
			CurrentMoney: balance,
		},
		Status: status,
	}
}
