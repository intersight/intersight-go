package domain

import "github.com/intersight/intersight-go"

const (
	StatusDeactivated intersight.DomainStatus = "deactivated"
	StatusActive      intersight.DomainStatus = "active"
	StatusApproval    intersight.DomainStatus = "approval"

	VerifyWeb intersight.DomainVerifyMode = "web"
	VerifyDNS intersight.DomainVerifyMode = "dns"
)
