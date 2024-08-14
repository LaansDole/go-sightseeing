package docs

/*
Business Rule: Active Subscription Cancellation

Stipulation Name: StipulationActiveSubscription = "Account Status/Active Subscription"

Scenario:
The user has an active subscription that is either paid monthly or yearly.

Business Rule:
- Monthly Subscription:
  - The user is allowed to cancel the subscription at any time during the month.
  - After cancellation, the service remains active until the end of the current billing month.
  - The user continues to have access to the service until the end of the month.
  - No further billing occurs after the current billing month ends.

- Yearly Subscription:
  - The user is allowed to cancel the subscription at any time during the year.
  - After cancellation, the service remains active until the end of the current billing year.
  - The user continues to have access to the service until the end of the year.
  - No further billing occurs after the current billing year ends.

Implementation:
- The system should check if the subscription is currently active and allow cancellation mid-period.
- Upon cancellation, the service should continue until the end of the current billing cycle.
*/
