package model

import "github.com/google/uuid"

// Individual represents data about a single individual.
type Individual struct {
	ID          uuid.UUID // Unique identifier for the individual
	Name        string
	Chromosomes []uint16
}
