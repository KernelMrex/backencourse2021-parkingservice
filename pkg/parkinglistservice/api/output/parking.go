package output

import "github.com/google/uuid"

type ParkingOutput struct {
	ID          uuid.UUID
	Name        string
	Description string
	Coordinates string
}
