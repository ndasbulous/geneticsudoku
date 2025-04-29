package util

import (
	"ndasbulous/geneticsudoku/model"

	"github.com/google/uuid"
)

func GenerateRandomIndividual() model.Individual {
	// Generate a random individual with a unique ID and name
	individual := model.Individual{
		ID:         uuid.New(), // This should be generated dynamically
		Name:       "John Doe",
		Chromosome: make([]uint8, 81), // Assuming a 9x9 Sudoku grid flattened to a single slice
		Fitness:    0.0,
	}

	return individual
}

func GenerateInitialPopulation(size int) model.Population {
	// Generate a population of individuals
	population := model.Population{
		ID:             uuid.UUID{}, // This should be generated dynamically
		Name:           "Population 1",
		Individuals:    make([]model.Individual, size),
		Generation:     1,
		BestFitness:    0.0,
		AverageFitness: 0.0,
		TotalFitness:   0.0,
	}

	for i := 0; i < size; i++ {
		population.Individuals[i] = GenerateRandomIndividual()
	}

	return population
}
