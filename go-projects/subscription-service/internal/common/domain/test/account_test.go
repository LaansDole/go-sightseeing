package test

import (
	"testing"

	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain"
	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain/stub"
	"github.com/stretchr/testify/assert"
)

func TestIsValidForCancellation(t *testing.T) {
	type args struct {
		account domain.Account
	}
	type want struct {
		expectedResult bool
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name: "GIVEN an active account THEN it should be valid for cancellation",
			args: args{
				account: stub.NewActiveAccountStub(),
			},
			want: want{
				expectedResult: true,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a trial account THEN it should be valid for cancellation",
			args: args{
				account: stub.NewTrialAccountStub(),
			},
			want: want{
				expectedResult: true,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a suspended account THEN it should not be valid for cancellation",
			args: args{
				account: stub.NewSuspendedAccountStub(),
			},
			want: want{
				expectedResult: false,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a cancelled account THEN it should not be valid for cancellation",
			args: args{
				account: stub.NewCancelledAccountStub(),
			},
			want: want{
				expectedResult: false,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a legacy account THEN it should be valid for cancellation",
			args: args{
				account: stub.NewLegacyAccountStub(),
			},
			want: want{
				expectedResult: true,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a family or enterprise account THEN it should be valid for cancellation",
			args: args{
				account: stub.NewFamilyOrEnterpriseAccountStub(),
			},
			want: want{
				expectedResult: true,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Perform the domain function
			result := tt.args.account.IsValidForCancellation()

			// Assertions
			assert.Equal(t, tt.want.expectedResult, result)
		})
	}
}
