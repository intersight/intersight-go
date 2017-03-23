package account

import (
	"github.com/intersights/intersights-go"
)

const (
	TypeUnknown     intersights.AccountType = "unknown"
	TypeResidential intersights.AccountType = "residential"
	TypeBusiness    intersights.AccountType = "business"
	TypeStudent     intersights.AccountType = "student"
	TypeOther       intersights.AccountType = "other"

	StatusInitialised  intersights.AccountStatus = "initialised"
	StatusProvisioning intersights.AccountStatus = "provisioning"
	StatusActive       intersights.AccountStatus = "active"
	StatusDisabled     intersights.AccountStatus = "disabled"
	StatusSuspended    intersights.AccountStatus = "suspended"
	StatusExpired      intersights.AccountStatus = "expired"
	StatusCancelled    intersights.AccountStatus = "cancelled"
	StatusClosed       intersights.AccountStatus = "closed"

	//LifecycleCold Provided details by a third party
	LifecycleCold intersights.AccountLifecycle = "cold"
	//LifecycleSubscriber Subscribed to a newsletter
	LifecycleSubscriber intersights.AccountLifecycle = "subscriber"
	//LifecycleSupport Contacted support
	LifecycleSupport intersights.AccountLifecycle = "support"
	//LifecycleLead Completed a product interest form e.g. more information about product
	LifecycleLead intersights.AccountLifecycle = "lead"
	// LifecycleMarketingQualifiedLead - Completed automated activities like viewing product demo
	LifecycleMarketingQualifiedLead intersights.AccountLifecycle = "marketingqualifiedlead"
	// LifecycleSalesQualifiedLead - Spoken to a sales rep for qualification
	LifecycleSalesQualifiedLead intersights.AccountLifecycle = "salesqualifiedlead"
	//LifecycleTrial Subscribed for a trial account
	LifecycleTrial intersights.AccountLifecycle = "trial"
	//LifecycleFree Subscribed for a free account
	LifecycleFree intersights.AccountLifecycle = "free"
	//LifecycleOpportunity Started Order Flow
	LifecycleOpportunity intersights.AccountLifecycle = "opportunity"
	//LifecycleCustomer Paying Customer
	LifecycleCustomer intersights.AccountLifecycle = "customer"
	//LifecycleCancelled Removed account (from paying)
	LifecycleCancelled intersights.AccountLifecycle = "cancelled"
	//LifecycleLapsed Auto Cancelled
	LifecycleLapsed intersights.AccountLifecycle = "lapsed"
	//LifecycleLost Removed account (From free)
	LifecycleLost intersights.AccountLifecycle = "lost"

	//OrientationCost focuses on least cost
	OrientationCost intersights.AccountOrientation = "cost"
	//OrientationValue focuses on value for money
	OrientationValue intersights.AccountOrientation = "value"
	//OrientationTechnology quality product, cost of less value
	OrientationTechnology intersights.AccountOrientation = "technology"
	//OrientationBrand anything brand does, will buy
	OrientationBrand intersights.AccountOrientation = "brand"
)
