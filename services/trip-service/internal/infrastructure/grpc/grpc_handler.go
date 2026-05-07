package grpc

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	pb "ride-sharing/shared/proto/trip"

	"google.golang.org/grpc"
)

type gRPCHandler struct {
	pb.UnimplementedTripServiceServer
	service domain.TripService
}

func NewgRPCHandler(server *grpc.Server, service domain.TripService) *gRPCHandler {
	handler := &gRPCHandler{
		service: service,
	}

	pb.RegisterTripServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) PreviewTrip(ctx context.Context, req *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {
	pickup := req.GetPickupLocation()
	destination := req.GetDropoffLocation()

	t, err := h.service.GetRoute(ctx, pickup, destination)
	if err != nil {
		return nil, err
	}
	return &pb.PreviewTripResponse{
		Route: route,
	}, nil
}
