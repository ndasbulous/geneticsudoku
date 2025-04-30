package model

import "github.com/google/uuid"

// Individual represents data about a single individual.
type Individual struct {
	ID         uuid.UUID // Unique identifier for the individual
	Name       string
	Chromosome Chromosome // The chromosome of the individual, representing the Sudoku grid
	Fitness    float64    // The fitness score of the individual
}
