package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/laansdole/go-sightseeing/go-projects/subscription-service/internal/common/domain"
)

func TestNegativeBalanceStipulationValidation_Eligibility(t *testing.T) {
	tests := []struct {
		name        string
		stipulation domain.NegativeBalanceStipulationValidation
		want        domain.Validity
	}{
		{
			name: "GIVEN negative balance stipulation validation THEN invalid",
			stipulation: domain.NegativeBalanceStipulationValidation{
				CurrentBalance: -1,
			},
			want: domain.InvalidForCancellation,
		},
		{
			name: "GIVEN positive balance stipulation validation THEN valid",
			stipulation: domain.NegativeBalanceStipulationValidation{
				CurrentBalance: 1,
			},
			want: domain.ValidForCancellation,
		},
		{
			name: "GIVEN zero balance stipulation validation THEN valid",
			stipulation: domain.NegativeBalanceStipulationValidation{
				CurrentBalance: 0,
			},
			want: domain.ValidForCancellation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.stipulation.Validity()
			assert.Equal(t, tt.want, got)
		})
	}
}
