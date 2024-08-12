package domain

import (
	"github.com/laansdole/go-sightseeing/internal/common/errorf"
)

type Validity string

const (
	InvalidForCancellation Validity = "InvalidForCancellation"
	ValidForCancellation   Validity = "ValidForCancellation"
)

type StipulationValidation interface {
	StipulationName() StipulationName
	Validity() Validity
}

type StipulationValidations struct {
	// Group: Account status
	NegativeBalance NegativeBalanceStipulationValidation

	// Group: Payment Arrangements
}

type ValidateAccountForCancellationRequest struct {
	AccountNumber AccountNumber
	Balances      AccountBalances
}

type AccountBalances struct {
	CurrentBalance float64
}

type ValidateAccountForCancellationResponse struct {
	// AccountNumber of the account being validated
	AccountNumber AccountNumber

	// Validity determines if the account can be cancelled
	Validity Validity

	// Concrete types is more preferred than interface for StipulationValidation
	// because using concrete type will make marshalling/unmarshalling this ValidateAccountForCancellationResponse easier
	Stipulations StipulationValidations
}

func NewValidateAccountForCancellationResponse(
	accountNumber AccountNumber,
	stipulations []StipulationValidation,
) (ValidateAccountForCancellationResponse, error) {
	response := ValidateAccountForCancellationResponse{
		AccountNumber: accountNumber,
		Validity:      ValidForCancellation,
	}
	for _, stp := range stipulations {
		response.Validity = checkAccountValidity(response.Validity, stp.Validity())
		switch stp := stp.(type) {

		// Group: Account status
		case NegativeBalanceStipulationValidation:
			response.Stipulations.NegativeBalance = stp

		// Group: Payment Arrangements

		default:
			return ValidateAccountForCancellationResponse{}, errorf.IsInternalServerf("invalid stipulation type %T", stp)
		}
	}
	return response, nil
}

func checkAccountValidity(
	accountValidity Validity,
	stipulationValidity Validity,
) Validity {
	if accountValidity == InvalidForCancellation {
		return InvalidForCancellation
	}
	if stipulationValidity == InvalidForCancellation {
		return InvalidForCancellation
	}
	return ValidForCancellation
}
