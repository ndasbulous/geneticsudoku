package main

import (
	"fmt"
	"math/rand"
	"ndasbulous/geneticsudoku/constant"
	"ndasbulous/geneticsudoku/model"
	"ndasbulous/geneticsudoku/util"
	"sort"
	"time"

	"github.com/google/uuid"
)

func init() {
	// Initialize random seed
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func main() {
	runGenerationsWithPredefinedValues()
}

func runGenerationsWithPredefinedValues() {
	// Replace the crossover method line with this
	crossoverMethods := []string{
		constant.CROSSOVER_METHOD_ROW,
		constant.CROSSOVER_METHOD_COLUMN,
		constant.CROSSOVER_METHOD_SECTION,
	}

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
			Chromosome: model.Chromosome{Sequence: util.GenerateRandomChromosomeWithPredefinedValues(predefinedValues)},
			Fitness:    0,
		}
		individual.Fitness = util.CalculateFitness(individual)

		population[index] = individual
	}

	sort.Slice(population[:], func(i, j int) bool {
		return population[i].Fitness < population[j].Fitness
	})

	newIndividuals := []model.Individual{}

	var currentGeneration int = 1
	for population[0].Fitness != 0 {
		fmt.Printf("Current Generation: %d, Best fitness: %d\n", currentGeneration, population[0].Fitness)

		// Do crossover
		for i := 0; i < populationSize; i++ {
			// Crossover between two parents
			parent1 := population[i]
			// Choose one random individual from the population, not the same as parent1
			// Ensure parent2 is not the same as parent1
			parent2 := population[(uint(rand.Uint32()) % populationSize)]

			childChromosome1, childChromosome2 := util.CrossoverTwoChromosomes(
				parent1.Chromosome.Sequence,
				parent2.Chromosome.Sequence,
				crossoverMethods[rand.Intn(len(crossoverMethods))],
				uint8(rand.Intn(8)+1))

			// Create new individuals for the children
			child1 := model.Individual{
				ID:         uuid.New(), // This should be generated dynamically
				Name:       "Child 1",
				Chromosome: model.Chromosome{Sequence: childChromosome1},
				Fitness:    0,
			}
			child2 := model.Individual{
				ID:         uuid.New(), // This should be generated dynamically
				Name:       "Child 2",
				Chromosome: model.Chromosome{Sequence: childChromosome2},
				Fitness:    0,
			}

			// Calculate fitness for the children
			child1.Fitness = util.CalculateFitness(child1)
			child2.Fitness = util.CalculateFitness(child2)

			// Add the children to the population
			newIndividuals = append(newIndividuals, child1)
			newIndividuals = append(newIndividuals, child2)
		}

		// Sort new individuals by fitness
		sort.Slice(newIndividuals[:], func(i, j int) bool {
			return newIndividuals[i].Fitness < newIndividuals[j].Fitness
		})

		// Replace the worst individuals in the population with the new individuals
		for i := 0; i < populationSize/2; i++ {
			population[populationSize-1-i] = newIndividuals[i]
		}

		// Sort the population again
		sort.Slice(population[:], func(i, j int) bool {
			return population[i].Fitness < population[j].Fitness
		})

		// Clear the new individuals slice for the next generation
		newIndividuals = []model.Individual{}

		// Increment the generation count
		currentGeneration++
	}

	fmt.Printf("Best individual:\n")
	// Print individual in matrix format
	for i := 0; i < len(population[0].Chromosome.Sequence); i++ {
		for j := 0; j < len(population[0].Chromosome.Sequence[i]); j++ {
			fmt.Printf("%d ", population[0].Chromosome.Sequence[i][j])
		}
		fmt.Println()
	}
}
