package domain

type StipulationName string

// Account Status Stipulations
const (
	StipulationNegativeBalancePrediction StipulationName = "Account Status/Negative Balance Prediction"
	StipulationActiveSubscription        StipulationName = "Account Status/Active Subscription"
	StipulationDeceasedAccountHolder     StipulationName = "Account Status/Deceased Account Holder"
)

// Payment Arrangement Stipulations
const (
	StipulationPendingCardChange         StipulationName = "Payment Arrangement/Pending Card Change"
	StipulationPendingPaymentArrangement StipulationName = "Payment Arrangement/Pending Payment Arrangement"
)
