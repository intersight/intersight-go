package event

import "github.com/intersight/intersight-go"

const (
	SourceWeb       intersight.SourceType = "web"
	SourceMobile    intersight.SourceType = "mobile"
	SourceAPI       intersight.SourceType = "api"
	SourcePhone     intersight.SourceType = "phone"
	SourceRealWorld intersight.SourceType = "realworld"
	SourceWebhook   intersight.SourceType = "webhook"
	SourceOther     intersight.SourceType = "other"

	//TypeEvent generic event
	TypeEvent intersight.EventType = "event"
	//TypeProperty setting a property on an entity
	TypeProperty intersight.EventType = "property"
	//TypeHeatmap heatmap data
	TypeHeatmap intersight.EventType = "heatmap"
	//TypeForm information about a form
	TypeForm intersight.EventType = "form"
)
