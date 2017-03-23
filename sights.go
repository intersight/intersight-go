package intersight

import "time"

type SourceType string
type SightType string

type SightSource struct {
	DeviceID   string `json:"d"`
	SessionID  string `json:"s"`
	IdentityID string `json:"i"`
	ProjectID  string `json:"p"`
	//TrackerID allows an app to narrow down the source of an event e.g. specific js tracker with differing configs
	TrackerID string `json:"t"`
	//RequestID is a micro session the event is attached to, e.g. A single page load would have a single request ID, but many events.  Can be split with / to describe a sub request.
	RequestID        string `json:"r"`
	RequestStartTime time.Time `json:"rt"`
	SourceType       SourceType `json:"st"`
	Hostname         string `json:"h"`
	IP               string `json:"ip"`
	UserAgent        string `json:"ua"`
	Interactions     map[string]int32  `json:"in"`
}

type Sight struct {
	ID            string  `json:"id"`
	SightTime     time.Time  `json:"st"`
	Source        SightSource `json:"s"`
	TrafficSource TrafficSource `json:"tr"`
	Key           string //event name | heatmap send type | ID to store properties against  `json:"k"`
	Type          SightType `json:"t"`
	Data          map[string]string `json:"d"`
	Meta          map[string]string `json:"m"`
}
