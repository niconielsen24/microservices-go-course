package main

import (
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"
)

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (p *previewTripRequest) ToProto() *pb.PreviewTripRequest {
	return &pb.PreviewTripRequest{
		UserId: p.UserID,
		PickupLocation: &pb.Coordinate{
			Latitude:  p.Pickup.Latitude,
			Longitude: p.Pickup.Longitude,
		},
		DropoffLocation: &pb.Coordinate{
			Latitude:  p.Destination.Latitude,
			Longitude: p.Destination.Longitude,
		},
	}
}
