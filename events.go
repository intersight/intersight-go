package intersight

import "time"

type SourceType string
type EventType string

type EventSource struct {
	DeviceID   string
	SessionID  string
	IdentityID string
	AppID      string
	//TrackerID allows an app to narrow down the source of an event e.g. specific js tracker with differing configs
	TrackerID string
	//RequestID is a micro session the event is attached to, e.g. A single page load would have a single request ID, but many events
	RequestID        string
	RequestStartTime time.Time
	SourceType       SourceType
	Hostname         string
	IP               string
	UserAgent        string
	Interactions     map[string]int32
}

type Event struct {
	ID            string
	EventTime     time.Time
	Source        EventSource
	TrafficSource TrafficSource
	Key           string //event name | heatmap send type | ID to store properties against
	Type          EventType
	Data          map[string]string
	Meta          map[string]string
}
