package intersights

import "time"

type AccountType string
type AccountStatus string
type AccountLifecycle string
type AccountOrientation string

type Account struct {
	ID            string
	BrandID       string
	AccountNumber string
	Name          string
	PrimaryEmail  string
	PrimaryPhone  string
	//Language IETF language tag e.g. en / en-US / en-GB
	Language    string
	Timezone    string
	Country     string
	Type        AccountType
	SubType     string
	Lifecycle   AccountLifecycle
	Orientation AccountOrientation

	TrafficSource TrafficSource

	ParentAccountID string
	//AccountManager [type]accountId e.g. [technical]john, [sales]sarah
	AccountManager map[string]string

	OpenDate   time.Time
	CloseDate  time.Time
	TrialStart time.Time
	TrialEnd   time.Time

	//ImportantDates significant dates relating to this account e.g. reopen
	ImportantDates map[string]time.Time

	//Useful Flags

	//Loyal Loyal Customer to brand
	Loyal bool
	//HPA High Priority Account
	HPA bool
	//Fraud Fraudulent Account
	Fraud bool
	//Disputes Account Had Disputes
	Disputes bool
	//Discount Account Interested in Discounts
	Discount bool
	//Impulse Will buy impulsively
	Impulse bool
	//Test TestAccount Account
	TestAccount bool
	//Purchased Account purchased anything from brand
	Purchased bool
	//Renewing Account is renewing
	Renewing bool
	//AtRisk Usually unpaid bills, at risk of cancel
	AtRisk bool

	//Billing Information
	Balance          int64
	IsDelinquent     bool
	TotalPaid        int64
	TotalRefunded    int64
	TotalDisputed    int64
	TotalDisputeLost int64
	DelinquentSince  time.Time

	Products []AccountProduct
}
type AccountProduct struct {
	ProductID string
	FeatureID string
	Status    ProductStatus
}
