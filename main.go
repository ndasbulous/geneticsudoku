package main

import (
	"fmt"
	"math"
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
		{1, 0, 0, 0, 0, 7, 0, 9, 0},
		{0, 3, 0, 0, 2, 0, 0, 0, 8},
		{0, 0, 9, 6, 0, 0, 5, 0, 0},
		{0, 0, 5, 3, 0, 0, 9, 0, 0},
		{0, 1, 0, 0, 8, 0, 0, 0, 2},
		{6, 0, 0, 0, 0, 4, 0, 0, 0},
		{3, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 4, 0, 0, 0, 0, 0, 0, 7},
		{0, 0, 7, 0, 0, 0, 3, 0, 0},
	}

	const populationSize = 2000
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
		// Calculate average fitness
		var totalFitness float64 = 0
		for _, individual := range population {
			totalFitness += float64(individual.Fitness)
		}
		averageFitness := totalFitness / float64(populationSize)
		fmt.Printf("Current Generation: %d, Best fitness: %d, Average fitness: %f\n", currentGeneration, population[0].Fitness, averageFitness)

		// Do crossover if best fitness is not equal to average fitness
		if population[0].Fitness != uint8(math.Floor(averageFitness)) {
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

		} else {
			// Do mutation on all individuals
			for i := 0; i < populationSize; i++ {
				// Create two mutated individuals
				mutatedIndividual := util.Mutate(population[i].Chromosome.Sequence, predefinedValues)
				mutatedIndividual2 := util.Mutate(population[i].Chromosome.Sequence, predefinedValues)

				// Create two new individuals with the mutated chromosome
				newIndividual := model.Individual{
					ID:         uuid.New(), // This should be generated dynamically
					Name:       "Mutated Individual",
					Chromosome: model.Chromosome{Sequence: mutatedIndividual},
					Fitness:    0,
				}

				newIndividual2 := model.Individual{
					ID:         uuid.New(), // This should be generated dynamically
					Name:       "Mutated Individual 2",
					Chromosome: model.Chromosome{Sequence: mutatedIndividual2},
					Fitness:    0,
				}

				// Calculate fitness for those two individuals
				newIndividual.Fitness = util.CalculateFitness(newIndividual)
				newIndividual2.Fitness = util.CalculateFitness(newIndividual2)

				// Add the new individual to the population
				newIndividuals = append(newIndividuals, newIndividual)
				newIndividuals = append(newIndividuals, newIndividual2)

				// Replace random individuals in the population with the new individuals
				population[rand.Intn(populationSize)] = newIndividuals[rand.Intn(len(newIndividuals))]
				population[rand.Intn(populationSize)] = newIndividuals[rand.Intn(len(newIndividuals))]
			}
		}

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
