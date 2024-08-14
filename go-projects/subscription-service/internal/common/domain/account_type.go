package domain

type AccountStatus string

const (
	TrialAccount     AccountStatus = "trial"
	CancelledAccount AccountStatus = "cancelled"
	LegacyAccount    AccountStatus = "legacy"
	ActiveAccount    AccountStatus = "active"
	UnknownAccount   AccountStatus = "unknown"
)

func AccountStatusOf(code string) AccountStatus {
	switch code {
	case "trial", "t":
		return TrialAccount
	case "cancelled", "c":
		return CancelledAccount
	case "legacy", "l":
		return LegacyAccount
	case "active", "a":
		return ActiveAccount
	default:
		return UnknownAccount
	}
}

func (status AccountStatus) String() string {
	switch status {
	case TrialAccount:
		return "trial"
	case CancelledAccount:
		return "cancelled"
	case LegacyAccount:
		return "legacy"
	case ActiveAccount:
		return "active"
	default:
		return "unknown"
	}
}
