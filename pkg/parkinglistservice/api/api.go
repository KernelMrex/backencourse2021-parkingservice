package api

import (
	"github.com/google/uuid"
	"parkinglistservice/pkg/parkinglistservice/api/output"
	"parkinglistservice/pkg/parkinglistservice/domain/model"
)

type Api interface {
	GetParkingByID(id uuid.UUID) *output.ParkingOutput
}

type api struct {
	parkingRepository model.ParkingRepository
}

func NewApi(parkingRepository model.ParkingRepository) Api {
	return &api{
		parkingRepository: parkingRepository,
	}
}

func (api *api) GetParkingByID(id uuid.UUID) *output.ParkingOutput {
	return nil
}
