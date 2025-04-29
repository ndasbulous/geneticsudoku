package util

import (
	"math/rand"
	"ndasbulous/geneticsudoku/model"

	"github.com/google/uuid"
)

func GenerateRandomIndividual() model.Individual {
	// Generate a random individual with a unique ID and name
	individual := model.Individual{
		ID:         uuid.New(), // This should be generated dynamically
		Name:       "John Doe",
		Chromosome: GenerateRandomChromosome(), // Assuming a 9x9 Sudoku grid flattened to a single slice
		Fitness:    0.0,
	}

	individual.Fitness = CalculateFitness(individual)

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

func GenerateRandomChromosome() []uint8 {
	// Generate a random chromosome for an individual
	chromosome := make([]uint8, 81) // Assuming a 9x9 Sudoku grid flattened to a single slice

	for i := 0; i < len(chromosome); i++ {
		chromosome[i] = uint8(rand.Uint32()) % 10 // Random number between 0 and 9
		// Note: This is a simple random number generation for demonstration purposes
		// You can also use a more sophisticated method to generate valid Sudoku numbers
	}

	return chromosome
}
