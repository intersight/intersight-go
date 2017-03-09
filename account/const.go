package account

import "github.com/intersight/intersight-go"

const (
	TypeUnknown     intersight.AccountType = "unknown"
	TypeResidential intersight.AccountType = "residential"
	TypeBusiness    intersight.AccountType = "business"
	TypeStudent     intersight.AccountType = "student"
	TypeOther       intersight.AccountType = "other"

	StatusInitialised  intersight.AccountStatus = "initialised"
	StatusProvisioning intersight.AccountStatus = "provisioning"
	StatusActive       intersight.AccountStatus = "active"
	StatusDisabled     intersight.AccountStatus = "disabled"
	StatusSuspended    intersight.AccountStatus = "suspended"
	StatusExpired      intersight.AccountStatus = "expired"
	StatusCancelled    intersight.AccountStatus = "cancelled"
	StatusClosed       intersight.AccountStatus = "closed"

	//LifecycleCold Provided details by a third party
	LifecycleCold intersight.AccountLifecycle = "cold"
	//LifecycleSubscriber Subscribed to a newsletter
	LifecycleSubscriber intersight.AccountLifecycle = "subscriber"
	//LifecycleSupport Contacted support
	LifecycleSupport intersight.AccountLifecycle = "support"
	//LifecycleLead Completed a product interest form e.g. more information about product
	LifecycleLead intersight.AccountLifecycle = "lead"
	// LifecycleMarketingQualifiedLead - Completed automated activities like viewing product demo
	LifecycleMarketingQualifiedLead intersight.AccountLifecycle = "marketingqualifiedlead"
	// LifecycleSalesQualifiedLead - Spoken to a sales rep for qualification
	LifecycleSalesQualifiedLead intersight.AccountLifecycle = "salesqualifiedlead"
	//LifecycleTrial Subscribed for a trial account
	LifecycleTrial intersight.AccountLifecycle = "trial"
	//LifecycleFree Subscribed for a free account
	LifecycleFree intersight.AccountLifecycle = "free"
	//LifecycleOpportunity Started Order Flow
	LifecycleOpportunity intersight.AccountLifecycle = "opportunity"
	//LifecycleCustomer Paying Customer
	LifecycleCustomer intersight.AccountLifecycle = "customer"
	//LifecycleCancelled Removed account (from paying)
	LifecycleCancelled intersight.AccountLifecycle = "cancelled"
	//LifecycleLapsed Auto Cancelled
	LifecycleLapsed intersight.AccountLifecycle = "lapsed"
	//LifecycleLost Removed account (From free)
	LifecycleLost intersight.AccountLifecycle = "lost"

	//OrientationCost focuses on least cost
	OrientationCost intersight.AccountOrientation = "cost"
	//OrientationValue focuses on value for money
	OrientationValue intersight.AccountOrientation = "value"
	//OrientationTechnology quality product, cost of less value
	OrientationTechnology intersight.AccountOrientation = "technology"
	//OrientationBrand anything brand does, will buy
	OrientationBrand intersight.AccountOrientation = "brand"
)
