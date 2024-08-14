package domain_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain"
	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain/stub"
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
			name: "GIVEN an active account with a positive balance THEN it should be valid for cancellation",
			args: args{
				account: stub.NewActiveAccountStub("active123"),
			},
			want: want{
				expectedResult: true,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a trial account with a zero balance THEN it should be valid for cancellation",
			args: args{
				account: stub.NewTrialAccountStub("trial456"),
			},
			want: want{
				expectedResult: true,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a cancelled account THEN it should not be valid for cancellation",
			args: args{
				account: stub.NewCancelledAccountStub("cancelled012"),
			},
			want: want{
				expectedResult: false,
			},
			wantErr: false,
		},
		{
			name: "GIVEN a legacy account with a positive balance THEN it should be valid for cancellation",
			args: args{
				account: stub.NewLegacyAccountStub("legacy345"),
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
			result := tt.args.account.Status

			// Assertions
			assert.Equal(t, tt.want.expectedResult, result)
		})
	}
}
