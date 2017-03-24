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
			SourceType:       pbsights.SourceType(sight.Source.SourceType),
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
		Type:       pbsights.SightType(sight.Type),
		JsonLd:     sight.JsonLd,
		Properties: sight.Data,
		Meta:       sight.Meta,
	}

	return c.c.Write(ctx, ps)
}
