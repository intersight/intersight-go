package sights

import (
	"github.com/intersights/intersights-go"
)

const (
	SourceWeb       intersight.SourceType = "web"
	SourceMobile    intersight.SourceType = "mobile"
	SourceAPI       intersight.SourceType = "api"
	SourcePhone     intersight.SourceType = "phone"
	SourceRealWorld intersight.SourceType = "realworld"
	SourceWebhook   intersight.SourceType = "webhook"
	SourceOther     intersight.SourceType = "other"

	//TypeEvent generic event
	TypeEvent intersight.SightType = "event"
	//TypeProperty setting a property on an entity
	TypeProperty intersight.SightType = "property"
	//TypeHeatmap heatmap data
	TypeHeatmap intersight.SightType = "heatmap"
	//TypeForm information about a form
	TypeForm intersight.SightType = "form"
)
