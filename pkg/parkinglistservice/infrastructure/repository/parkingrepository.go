package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"parkinglistservice/pkg/parkinglistservice/domain/model"
)

type parkingRepository struct {
	db *sql.DB
}

func NewParkingRepository(db *sql.DB) model.ParkingRepository {
	return &parkingRepository{db: db}
}

func (pr *parkingRepository) Get(id uuid.UUID) (*model.Parking, error) {
	return nil, nil
}
