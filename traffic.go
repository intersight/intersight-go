package intersight

type TrafficChannel string

type TrafficSource struct {
	//Source - utm_source - vendor / advertiser ID / sub channel
	Source string
	//Channel - utm_medium
	Channel    TrafficChannel
	SubChannel string
	//Campaign - utm_name
	Campaign string
	//Term - utm_term
	Term        string
	TrackingIDs map[int32]string
}
