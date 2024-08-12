package domain

type NegativeBalanceStipulationValidation struct {
	CurrentBalance float64
}

var _ StipulationValidation = NegativeBalanceStipulationValidation{}

func (v NegativeBalanceStipulationValidation) StipulationName() StipulationName {
	return StipulationNegativeBalance
}

func (v NegativeBalanceStipulationValidation) Validity() Validity {
	if v.CurrentBalance < 0 {
		return InvalidForCancellation
	}
	return ValidForCancellation
}
