package domain

import (
	"github.com/intersights/intersights-go"
)

const (
	StatusDeactivated intersights.DomainStatus = "deactivated"
	StatusActive      intersights.DomainStatus = "active"
	StatusApproval    intersights.DomainStatus = "approval"

	VerifyWeb intersights.DomainVerifyMode = "web"
	VerifyDNS intersights.DomainVerifyMode = "dns"
)
