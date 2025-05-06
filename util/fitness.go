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

	// Add fitness calculation for all the sections
	overallFitness += calculateSectionUniqueness(individual.Chromosome.Sequence)

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

func calculateSectionUniqueness(wholeBoard [9][9]uint8) uint8 {
	fitness := uint8(0)

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[0][0],
		wholeBoard[0][1],
		wholeBoard[0][2],
		wholeBoard[1][0],
		wholeBoard[1][1],
		wholeBoard[1][2],
		wholeBoard[2][0],
		wholeBoard[2][1],
		wholeBoard[2][2],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[0][3],
		wholeBoard[0][4],
		wholeBoard[0][5],
		wholeBoard[1][3],
		wholeBoard[1][4],
		wholeBoard[1][5],
		wholeBoard[2][3],
		wholeBoard[2][4],
		wholeBoard[2][5],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[0][6],
		wholeBoard[0][7],
		wholeBoard[0][8],
		wholeBoard[1][6],
		wholeBoard[1][7],
		wholeBoard[1][8],
		wholeBoard[2][6],
		wholeBoard[2][7],
		wholeBoard[2][8],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[3][0],
		wholeBoard[3][1],
		wholeBoard[3][2],
		wholeBoard[4][0],
		wholeBoard[4][1],
		wholeBoard[4][2],
		wholeBoard[5][0],
		wholeBoard[5][1],
		wholeBoard[5][2],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[3][3],
		wholeBoard[3][4],
		wholeBoard[3][5],
		wholeBoard[4][3],
		wholeBoard[4][4],
		wholeBoard[4][5],
		wholeBoard[5][3],
		wholeBoard[5][4],
		wholeBoard[5][5],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[3][6],
		wholeBoard[3][7],
		wholeBoard[3][8],
		wholeBoard[4][6],
		wholeBoard[4][7],
		wholeBoard[4][8],
		wholeBoard[5][6],
		wholeBoard[5][7],
		wholeBoard[5][8],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[6][0],
		wholeBoard[6][1],
		wholeBoard[6][2],
		wholeBoard[7][0],
		wholeBoard[7][1],
		wholeBoard[7][2],
		wholeBoard[8][0],
		wholeBoard[8][1],
		wholeBoard[8][2],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[6][3],
		wholeBoard[6][4],
		wholeBoard[6][5],
		wholeBoard[7][3],
		wholeBoard[7][4],
		wholeBoard[7][5],
		wholeBoard[8][3],
		wholeBoard[8][4],
		wholeBoard[8][5],
	})

	fitness += calculateArrayUniqueness([]uint8{
		wholeBoard[6][6],
		wholeBoard[6][7],
		wholeBoard[6][8],
		wholeBoard[7][6],
		wholeBoard[7][7],
		wholeBoard[7][8],
		wholeBoard[8][6],
		wholeBoard[8][7],
		wholeBoard[8][8],
	})

	return fitness
}
