package transport

import (
	"parkinglistservice/pkg/parkinglistservice/domain/model"
)

type GRPCServer struct {
	parkingRepository model.ParkingRepository
}

func NewGRPCServer(parkingRepository model.ParkingRepository) *GRPCServer {
	return &GRPCServer{
		parkingRepository: parkingRepository,
	}
}
