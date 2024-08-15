package domain_test

import (
	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain/stub"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain"
)

func TestExtractAccountNumberFromResourceName(t *testing.T) {
	want := domain.AccountNumber("1234")
	got := domain.ExtractAccountNumberFrom("accounts/1234")
	assert.Equal(t, want, got, "ResourceName()")
}

func TestResourceNameFromAccountNumber(t *testing.T) {
	want := "accounts/1234"
	got := domain.AccountNumber("1234").ResourceName()
	assert.Equal(t, want, got, "ResourceName()")
}

func TestIsEqualAccountNumber(t *testing.T) {
	tests := []struct {
		name string
		acc  domain.AccountNumber
		comp string
		want bool
	}{
		{
			name: "GIVEN equal account numbers THEN returns true",
			acc:  domain.AccountNumber("1234"),
			comp: "1234",
			want: true,
		},
		{
			name: "GIVEN non-equal account numbers THEN returns false",
			acc:  domain.AccountNumber("1234"),
			comp: "5678",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.acc.IsEqual(tt.comp)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFilterLegacyAccount(t *testing.T) {
	tests := []struct {
		name     string
		accounts domain.Accounts
		want     domain.Accounts
	}{
		{
			name: "GIVEN a mixed list of accounts THEN returns only legacy accounts",
			accounts: domain.Accounts{
				stub.NewActiveAccountStub("active123"),
				stub.NewLegacyAccountStub("legacy456"),
				stub.NewCancelledAccountStub("cancelled789"),
				stub.NewLegacyAccountStub("legacy012"),
				stub.NewTrialAccountStub("trial345"),
			},
			want: domain.Accounts{
				stub.NewLegacyAccountStub("legacy456"),
				stub.NewLegacyAccountStub("legacy012"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.accounts.FilterLegacyAccount()
			assert.Equal(t, tt.want, got)
		})
	}
}
