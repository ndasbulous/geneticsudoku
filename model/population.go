package model

import "github.com/google/uuid"

// Population represents a collection of individuals.
type Population struct {
	ID             uuid.UUID // Unique identifier for the population
	Name           string
	Individuals    []Individual
	Generation     uint64     // The generation number of the population
	BestFitness    float64    // The best fitness score of the population
	AverageFitness float64    // The average fitness score of the population
	TotalFitness   float64    // The total fitness score of the population
	BestIndividual Individual // The best individual in the population
}
