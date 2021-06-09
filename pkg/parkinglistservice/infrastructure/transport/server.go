package transport

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(codes.InvalidArgument, "cannot parse uuid")
	}

	parkingModel, err := s.parkingRepository.Get(id)
	if err != nil {
		log.Errorf("cannot get from repository: %v", errors.WithStack(err))
		return nil, status.Error(codes.NotFound, "cannot get parking from repository")
	}

	return &parking.GetParkingResponse{
		Title:          parkingModel.Name,
		Description:    parkingModel.Description,
		Location:       parkingModel.Coordinates,
		AvailableSlots: 0,
		ReservedSlots:  0,
	}, nil
}

func (s *GRPCServer) AddParking(_ context.Context, r *parking.AddParkingRequest) (*parking.AddParkingResponse, error) {
	id := uuid.New()

	if err := s.parkingRepository.Add(&model.Parking{
		ID:          id,
		Name:        r.GetTitle(),
		Description: r.GetDescription(),
		Coordinates: r.GetLocation(),
	}); err != nil {
		log.Errorf("could not add new parking: %v", errors.WithStack(err))
		return nil, status.Error(codes.Aborted, "could not add new parking")
	}

	return &parking.AddParkingResponse{Id: id.String()}, nil
}
