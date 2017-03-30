package intersights

type SourceType string
type SightType string

type SightSource struct {
	ProjectID string `json:"p"`
	//TrackerID allows an app to narrow down the source of an event e.g. specific js tracker with differing configs
	TrackerID string `json:"t"`
	//RequestID is a micro session the event is attached to, e.g. A single page load would have a single request ID, but many events.  Can be split with / to describe a sub request.
	RequestID string `json:"r"`
	//ProfileVendor - The application responsible for generating this sight
	ProfileVendor    string `json:"pv"`
	RequestStartTime Time `json:"rt"`
	SourceType       SourceType `json:"st"`
	//SourceEntry entry point for the source e.g. URL, Phone Number
	SourceEntry  string `json:"se"`
	Hostname     string `json:"h"`
	IP           string `json:"ip"`
	UserAgent    string `json:"ua"`
	Interactions map[string]int32  `json:"in"`
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
	//Lookup properties to reference back to the profile, e.g. email, phone, ip
	Lookup map[string]string `json:"l"`
	//JsonLd - Used to describe the sight
	JsonLd string `json:"ld"`
}
