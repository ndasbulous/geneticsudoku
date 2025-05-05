package util

import (
	"ndasbulous/geneticsudoku/model"
)

func CalculateFitness(individual model.Individual) uint8 {
	// Calculate the overallFitness of the individual based on the Sudoku rules
	// This is a placeholder implementation; you should replace it with actual logic
	overallFitness := uint8(0)

	// Example: Count the number of unique numbers in the individual's chromosome
	i := 0
	for i < 9 {
		// rowFitness := calculateRowOrColumnUniqueness(individual.Chromosome[i*9 : (i+1)*8])

		rowToCalculate := individual.Chromosome.Sequence[i][0:9]
		fitness := calculateArrayUniqueness(rowToCalculate)
		overallFitness += fitness

		// j := 0
		// for j < 9{
		// columnFitness := calculateRowOrColumnUniqueness(individual.Chromosome.Sequence[j])
		// }

		columnToCalculate := []uint8{
			individual.Chromosome.Sequence[0][i],
			individual.Chromosome.Sequence[1][i],
			individual.Chromosome.Sequence[2][i],
			individual.Chromosome.Sequence[3][i],
			individual.Chromosome.Sequence[4][i],
			individual.Chromosome.Sequence[5][i],
			individual.Chromosome.Sequence[6][i],
			individual.Chromosome.Sequence[7][i],
			individual.Chromosome.Sequence[8][i],
		}

		fitness = calculateArrayUniqueness(columnToCalculate)
		overallFitness += fitness

		i++
	}

	i = 0

	return overallFitness
}

func calculateArrayUniqueness(row []uint8) uint8 {
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
