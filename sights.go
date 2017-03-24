package intersights

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
	RequestID string `json:"r"`
	//ProfileVendor - The application responsible for generating this sight
	ProfileVendor    string `json:"pv"`
	RequestStartTime Time `json:"rt"`
	SourceType       SourceType `json:"st"`
	Hostname         string `json:"h"`
	IP               string `json:"ip"`
	UserAgent        string `json:"ua"`
	Interactions     map[string]int32  `json:"in"`
	TrafficSource    TrafficSource `json:"ts"`
}

type Sight struct {
	ID        string  `json:"id"`
	SightTime Time  `json:"t"`
	Source    SightSource `json:"src"`
	Type      SightType `json:"type"`

	//ProfileNamespace - The namespace this profile exists in
	ProfileNamespace string `json:"pn"`
	//ProfileID - Which profile to track the sight against
	ProfileID string `json:"p"`
	//Data - Used for profile property sights
	Data map[string]string `json:"d"`
	//Meta - Used for any meta not defined
	Meta map[string]string `json:"m"`
	//JsonLd - Used to describe the sight
	JsonLd string `json:"ld"`
}
