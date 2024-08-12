package domain

type StipulationName string

// Account Status Stipulations
const (
	StipulationNegativeBalance       StipulationName = "Account Status/Negative Balance"
	StipulationActiveSubscription    StipulationName = "Account Status/Active Subscription"
	StipulationNoActiveServices      StipulationName = "Account Status/No Active Services"
	StipulationRefundProcessing      StipulationName = "Account Status/Refund Processing"
	StipulationDeceasedAccountHolder StipulationName = "Account Status/Deceased Account Holder"
	StipulationUnpaidService         StipulationName = "Account Status/Unpaid Service"
)

// Payment Arrangement Stipulations
const (
	StipulationPendingCardChange         StipulationName = "Payment Arrangement/Pending Card Change"
	StipulationPendingPaymentArrangement StipulationName = "Payment Arrangement/Pending Payment Arrangement"
)
