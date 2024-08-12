package stub

import "github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain"

// Create stubs for testing

func NewActiveAccountStub() domain.Account {
	return domain.ActiveAccount{
		BaseAccount: domain.BaseAccount{
			AccountNumber: "active123",
			Balance:       100.0,
		},
	}
}

func NewTrialAccountStub() domain.Account {
	return domain.TrialAccount{
		BaseAccount: domain.BaseAccount{
			AccountNumber: "trial123",
			Balance:       0.0,
		},
	}
}

func NewSuspendedAccountStub() domain.Account {
	return domain.SuspendedAccount{
		BaseAccount: domain.BaseAccount{
			AccountNumber: "suspended123",
			Balance:       50.0,
		},
	}
}

func NewCancelledAccountStub() domain.Account {
	return domain.CancelledAccount{
		BaseAccount: domain.BaseAccount{
			AccountNumber: "cancelled123",
			Balance:       0.0,
		},
		GracePeriodEnd: "2023-12-31",
	}
}

func NewLegacyAccountStub() domain.Account {
	return domain.LegacyAccount{
		BaseAccount: domain.BaseAccount{
			AccountNumber: "legacy123",
			Balance:       200.0,
		},
	}
}

func NewFamilyOrEnterpriseAccountStub() domain.Account {
	return domain.FamilyOrEnterpriseAccount{
		BaseAccount: domain.BaseAccount{
			AccountNumber: "family123",
			Balance:       300.0,
		},
		Members: 5,
	}
}
