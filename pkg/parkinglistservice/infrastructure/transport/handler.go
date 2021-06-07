package transport

import (
	"context"
	grpc "parkinglistservice/api/parkinglistservice"
)

func (s *GRPCServer) GetParking(_ context.Context, _ *grpc.GetParkingRequest) (*grpc.GetParkingResponse, error) {
	resp := grpc.GetParkingResponse{
		Title:          "Test title",
		Description:    "Test description",
		Location:       "Test location",
		AvailableSlots: 10,
		ReservedSlots:  15,
	}

	return &resp, nil
}
