package intersights

import "time"

type DomainStatus string
type DomainVerifyMode string

type Domain struct {
	Domain           string
	TargetSegmentIDs []string
	BrandID          string
	IsPrimary        bool
	AliasOf          string
	Status           DomainStatus
	VerifyCode       string
	Verified         bool
	VerifyDate       time.Time
	VerifyMode       DomainVerifyMode
}
