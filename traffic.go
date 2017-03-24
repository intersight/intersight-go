package intersights

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
	Term string
	//Creative - ID for creative used e.g. Banner ID
	Creative string
	//Offer - ID for any specific offer specified
	Offer string
	//TrackingIDs - any traffic source defined tracking IDs eg click ID
	TrackingIDs map[int32]string
}
