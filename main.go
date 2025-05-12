package main

import (
	"fmt"
	"ndasbulous/geneticsudoku/model"
	"ndasbulous/geneticsudoku/util"
	"sort"

	"github.com/google/uuid"
)

func main() {
	runGenerationsWithPredefinedValues()
}

func runGenerationsWithPredefinedValues() {
	predefinedValues := [9][9]uint8{
		{0, 3, 5, 2, 6, 0, 0, 8, 1},
		{6, 8, 0, 5, 7, 1, 0, 9, 0},
		{1, 0, 0, 8, 0, 4, 5, 0, 2},
		{0, 2, 0, 1, 9, 0, 0, 4, 7},
		{3, 0, 0, 6, 8, 2, 9, 1, 5},
		{9, 5, 0, 7, 0, 3, 0, 2, 0},
		{0, 0, 9, 0, 0, 6, 0, 7, 4},
		{0, 4, 8, 9, 0, 7, 1, 3, 0},
		{7, 0, 0, 4, 0, 8, 0, 5, 9},
	}

	const populationSize = 1000
	population := [populationSize]model.Individual{}
	for index := 0; index < populationSize; index++ {
		individual := model.Individual{
			ID:         uuid.New(), // This should be generated dynamically
			Name:       "John Doe",
			Chromosome: model.Chromosome{util.GenerateRandomChromosomeWithPredefinedValues(predefinedValues)},
			Fitness:    0,
		}
		individual.Fitness = util.CalculateFitness(individual)

		population[index] = individual
	}

	sort.Slice(population[:], func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})
	for index := 0; index < populationSize; index++ {
		// if population[index].Fitness < 10 {
		fmt.Printf("Individual %d, Fitness: %d\n", index+1, population[index].Fitness)
		// }
	}
}
