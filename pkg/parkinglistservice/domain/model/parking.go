package model

import "github.com/google/uuid"

type Parking struct {
	ID          uuid.UUID
	Name        string
	Description string
	Coordinates string
}

type ParkingRepository interface {
	Get(id uuid.UUID) (*Parking, error)
}
