package sights

import (
	"github.com/intersights/intersights-go"
)

const (
	SourceWeb       intersights.SourceType = "web"
	SourceMobile    intersights.SourceType = "mobile"
	SourceAPI       intersights.SourceType = "api"
	SourcePhone     intersights.SourceType = "phone"
	SourceRealWorld intersights.SourceType = "realworld"
	SourceWebhook   intersights.SourceType = "webhook"
	SourceOther     intersights.SourceType = "other"

	//TypeEvent generic event
	TypeEvent intersights.SightType = "event"
	//TypeProperty setting a property on an entity
	TypeProperty intersights.SightType = "property"
	//TypeHeatmap heatmap data
	TypeHeatmap intersights.SightType = "heatmap"
	//TypeForm information about a form
	TypeForm intersights.SightType = "form"
)
