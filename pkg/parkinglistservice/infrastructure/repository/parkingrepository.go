package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"parkinglistservice/pkg/parkinglistservice/domain/model"
)

type parkingRepository struct {
	db *sql.DB
}

func NewParkingRepository(db *sql.DB) *parkingRepository {
	return &parkingRepository{db: db}
}

type parkingData struct {
	Name        string
	Description string
	Coordinates string
}

func (pr *parkingRepository) Get(uuid uuid.UUID) (*model.Parking, error) {
	row := pr.db.QueryRow("SELECT `name`, `description`, `coordinates` FROM `parking` WHERE `id`=?", uuid.String())
	parking := &parkingData{}

	if err := row.Scan(&parking.Name, &parking.Description, &parking.Coordinates); err != nil {
		return nil, errors.WithStack(err)
	}

	return &model.Parking{
		ID:          uuid,
		Name:        parking.Name,
		Description: parking.Description,
		Coordinates: parking.Coordinates,
	}, nil
}
