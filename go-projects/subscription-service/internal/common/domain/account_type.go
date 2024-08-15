package domain

type AccountType string

const (
	AccountTypeLegacy     AccountType = "legacy"
	AccountTypeIndividual AccountType = "individual"
	AccountTypeBusiness   AccountType = "business"
	AccountTypeUnknown    AccountType = "unknown"
)

func (acc AccountType) String() string {
	switch acc {
	case AccountTypeLegacy:
		return "legacy"
	case AccountTypeIndividual:
		return "individual"
	case AccountTypeBusiness:
		return "business"
	default:
		return "unknown"
	}
}

func (acc Account) IsLegacy() bool {
	if acc.Type == AccountTypeLegacy || acc.Type.String() == "legacy" {
		return true
	}
	return false
}
