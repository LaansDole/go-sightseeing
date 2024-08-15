package domain

type AccountStatus string

const (
	AccountStatusTrial     AccountStatus = "trial"
	AccountStatusCancelled AccountStatus = "cancelled"
	AccountStatusActive    AccountStatus = "active"
	AccountStatusUnknown   AccountStatus = "unknown"
)

func AccountStatusOf(code string) AccountStatus {
	switch code {
	case "trial", "t":
		return AccountStatusTrial
	case "cancelled", "c":
		return AccountStatusCancelled
	case "active", "a":
		return AccountStatusActive
	default:
		return AccountStatusUnknown
	}
}

func (status AccountStatus) String() string {
	switch status {
	case AccountStatusTrial:
		return "trial"
	case AccountStatusCancelled:
		return "cancelled"
	case AccountStatusActive:
		return "active"
	default:
		return "unknown"
	}
}
