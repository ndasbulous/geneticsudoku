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
		Chromosome: model.Chromosome{Sequence: GenerateRandomChromosome()},
		Fitness:    0,
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

func GenerateRandomChromosome() [9][9]uint8 {
	// Generate a random chromosome for an individual
	var chromosome [9][9]uint8

	for i := 0; i < len(chromosome); i++ {
		for j := 0; j < len(chromosome[i]); j++ {
			chromosome[i][j] = (uint8(rand.Uint32()) % 9) + 1 // Random number between 1 and 9
		}
	}

	return chromosome
}

func GenerateRandomChromosomeWithPredefinedValues(predefinedValues [9][9]uint8) [9][9]uint8 {
	// Generate a random chromosome for an individual
	var chromosome [9][9]uint8

	for i := 0; i < len(chromosome); i++ {
		for j := 0; j < len(chromosome[i]); j++ {
			if predefinedValues[i][j] == 0 {
				// Generate random number if cell is not defined
				chromosome[i][j] = (uint8(rand.Uint32()) % 9) + 1 // Random number between 1 and 9
			}
		}
	}

	return chromosome
}
