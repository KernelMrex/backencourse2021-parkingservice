package transport

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	parking "parkinglistservice/api/parkinglistservice"
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

func (s *GRPCServer) GetParking(_ context.Context, r *parking.GetParkingRequest) (*parking.GetParkingResponse, error) {
	id, err := uuid.Parse(r.GetId())
	if err != nil {
		log.Errorf("cannot parse uuid: %v", errors.WithStack(err))
		return nil, errors.New("cannot parse uuid")
	}

	parkingModel, err := s.parkingRepository.Get(id)
	if err != nil {
		log.Errorf("cannot get from repository: %v", errors.WithStack(err))
		return nil, errors.New("cannot get from repository")
	}

	return &parking.GetParkingResponse{
		Title:          parkingModel.Name,
		Description:    parkingModel.Description,
		Location:       parkingModel.Coordinates,
		AvailableSlots: 0,
		ReservedSlots:  0,
	}, nil
}
