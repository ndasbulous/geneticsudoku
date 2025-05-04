package util

import (
	"ndasbulous/geneticsudoku/model"
)

func CalculateFitness(individual model.Individual) uint8 {
	// Calculate the fitness of the individual based on the Sudoku rules
	// This is a placeholder implementation; you should replace it with actual logic
	fitness := uint8(0)

	// Example: Count the number of unique numbers in the individual's chromosome
	i := 0
	for i < 9 {
		// rowFitness := calculateRowOrColumnUniqueness(individual.Chromosome[i*9 : (i+1)*8])
		rowFitness := calculateRowOrColumnUniqueness(individual.Chromosome.Sequence[i])
		fitness += rowFitness

		// j := 0
		// for j < 9{
		// columnFitness := calculateRowOrColumnUniqueness(individual.Chromosome.Sequence[j])
		// }

		i++
	}

	i = 0
	// for i < 9 {
	// 	column := []uint8{
	// 		individual.Chromosome.Sequence[i], individual.Chromosome[i+9], individual.Chromosome[i+18],
	// 		individual.Chromosome[i+27], individual.Chromosome[i+36], individual.Chromosome[i+45],
	// 		individual.Chromosome[i+54], individual.Chromosome[i+63], individual.Chromosome[i+72],
	// 	}
	// 	columnFitness := calculateRowOrColumnUniqueness(column)
	// 	fitness += columnFitness
	// 	i++
	// }

	return fitness
}

func calculateRowOrColumnUniqueness(row [9]uint8) uint8 {
	// Create a map to count occurrences of each number
	counts := make(map[uint8]int)

	// Count occurrences of each non-zero number
	for _, num := range row {
		if num != 0 {
			counts[num]++
		}
	}

	// Calculate number of duplicate entries
	duplicateCount := uint8(0)
	for _, count := range counts {
		if count > 1 {
			duplicateCount += uint8(count - 1)
		}
	}

	return duplicateCount
}
