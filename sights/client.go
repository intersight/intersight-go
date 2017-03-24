package sights

import (
	"context"

	"github.com/golang/protobuf/ptypes"
	"github.com/intersights/intersights-go"
	"github.com/intersights/proto-go/pbsights"
	"google.golang.org/grpc"
)

const (
	appID = "sights"
)

//SightsClient Sights Client
type Client struct {
	projectID string
	c         pbsights.SightsClient
}

//NewClient - create a new client to the sights server
func NewClient(ctx context.Context, projectID string, opts ...grpc.DialOption) (*Client, error) {

	client, err := grpc.DialContext(ctx, intersights.GetGrpcLocation(intersights.Vendor, appID), opts...)
	if err != nil {
		return nil, err
	}
	c := &Client{
		projectID: projectID,
		c:         pbsights.NewSightsClient(client),
	}
	return c, nil
}

func (c Client) Write(ctx context.Context, sight intersights.Sight) (*pbsights.SightResponse, error) {
	ctx = intersights.SetCallContext(ctx, c.projectID)

	rst, _ := ptypes.TimestampProto(sight.Source.RequestStartTime.Time)
	st, _ := ptypes.TimestampProto(sight.SightTime.Time)

	var sourceType pbsights.SourceType
	var sightType pbsights.SightType

	switch sight.Source.SourceType {
	case SourceWeb:
		sourceType = pbsights.SourceType_Web
		break
	case SourceMobile:
		sourceType = pbsights.SourceType_Mobile
		break
	case SourceAPI:
		sourceType = pbsights.SourceType_API
		break
	case SourcePhone:
		sourceType = pbsights.SourceType_Phone
		break
	case SourceRealWorld:
		sourceType = pbsights.SourceType_Real_World
		break
	case SourceWebhook:
		sourceType = pbsights.SourceType_Webhook
		break
	case SourceOther:
		sourceType = pbsights.SourceType_Other
		break
	}

	switch sight.Type {
	case TypeProperty:
		sightType = pbsights.SightType_Property
		break
	case TypeEvent:
		sightType = pbsights.SightType_Event
		break
	case TypeHeatmap:
		sightType = pbsights.SightType_Heatmap
		break
	case TypeForm:
		sightType = pbsights.SightType_Form
		break
	}

	ps := &pbsights.Sight{
		ProfileNamespace: sight.ProfileNamespace,
		ProfileId:        sight.ProfileID,
		Source: &pbsights.SightSource{
			DeviceId:         sight.Source.DeviceID,
			SessionId:        sight.Source.SessionID,
			IdentityId:       sight.Source.IdentityID,
			TrackerId:        sight.Source.TrackerID,
			RequestId:        sight.Source.RequestID,
			ProfileVendor:    sight.Source.ProfileVendor,
			SourceType:       sourceType,
			Hostname:         sight.Source.Hostname,
			Ip:               sight.Source.IP,
			UserAgent:        sight.Source.UserAgent,
			Interactions:     sight.Source.Interactions,
			RequestStartTime: rst,
			TrafficSource: &pbsights.TrafficSource{
				Source:      sight.Source.TrafficSource.Source,
				Channel:     string(sight.Source.TrafficSource.Channel),
				SubChannel:  sight.Source.TrafficSource.SubChannel,
				Campaign:    sight.Source.TrafficSource.Campaign,
				Term:        sight.Source.TrafficSource.Term,
				Creative:    sight.Source.TrafficSource.Creative,
				Offer:       sight.Source.TrafficSource.Offer,
				TrackingIds: sight.Source.TrafficSource.TrackingIDs,
			},
		},
		Time:       st,
		Type:       sightType,
		JsonLd:     sight.JsonLd,
		Properties: sight.Data,
		Meta:       sight.Meta,
	}

	return c.c.Write(ctx, ps)
}
