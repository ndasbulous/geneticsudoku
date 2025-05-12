package util

import "ndasbulous/geneticsudoku/constant"

func crossoverTwoChromosomes(individualA, individualB [9][9]uint8, crossoverMethod string, crossoverPoint uint8) ([9][9]uint8, [9][9]uint8) {
	childX := individualA
	childY := individualB

	if crossoverMethod == constant.CROSSOVER_METHOD_ROW {
		childX, childY = crossoverByRow(individualA, individualB, crossoverPoint)
	} else if crossoverMethod == constant.CROSSOVER_METHOD_COLUMN {
		childX, childY = crossoverByColumn(individualA, individualB, crossoverPoint)
		// } else if crossoverMethod == constant.CROSSOVER_METHOD_SECTION {
		// 	childX, childY = crossoverBySection(individualA, individualB, crossoverPoint)
	} else {
		panic("Invalid crossover method")
	}

	return childX, childY
}

func crossoverByRow(individualA, individualB [9][9]uint8, crossoverPoint uint8) ([9][9]uint8, [9][9]uint8) {
	childX := individualA
	childY := individualB

	// Do crossover on Row on crossover point
	if crossoverPoint == 1 {
		childX[0] = individualB[0]
		childY[0] = individualA[0]
	} else if crossoverPoint == 2 {
		childX[0] = individualB[0]
		childX[1] = individualB[1]
		childY[0] = individualA[0]
		childY[1] = individualA[1]
	} else if crossoverPoint == 3 {
		childX[0] = individualB[0]
		childX[1] = individualB[1]
		childX[2] = individualB[2]
		childY[0] = individualA[0]
		childY[1] = individualA[1]
		childY[2] = individualA[2]
	} else if crossoverPoint == 4 {
		childX[0] = individualB[0]
		childX[1] = individualB[1]
		childX[2] = individualB[2]
		childX[3] = individualB[3]
		childY[0] = individualA[0]
		childY[1] = individualA[1]
		childY[2] = individualA[2]
		childY[3] = individualA[3]
	} else if crossoverPoint == 5 {
		childX[0] = individualB[0]
		childX[1] = individualB[1]
		childX[2] = individualB[2]
		childX[3] = individualB[3]
		childX[4] = individualB[4]
		childY[0] = individualA[0]
		childY[1] = individualA[1]
		childY[2] = individualA[2]
		childY[3] = individualA[3]
		childY[4] = individualA[4]
	} else if crossoverPoint == 6 {
		childX[0] = individualB[0]
		childX[1] = individualB[1]
		childX[2] = individualB[2]
		childX[3] = individualB[3]
		childX[4] = individualB[4]
		childX[5] = individualB[5]
		childY[0] = individualA[0]
		childY[1] = individualA[1]
		childY[2] = individualA[2]
		childY[3] = individualA[3]
		childY[4] = individualA[4]
		childY[5] = individualA[5]
	} else if crossoverPoint == 7 {
		childX[0] = individualB[0]
		childX[1] = individualB[1]
		childX[2] = individualB[2]
		childX[3] = individualB[3]
		childX[4] = individualB[4]
		childX[5] = individualB[5]
		childX[6] = individualB[6]
		childY[0] = individualA[0]
		childY[1] = individualA[1]
		childY[2] = individualA[2]
		childY[3] = individualA[3]
		childY[4] = individualA[4]
		childY[5] = individualA[5]
		childY[6] = individualA[6]
	} else if crossoverPoint == 8 {
		childX[0] = individualB[0]
		childX[1] = individualB[1]
		childX[2] = individualB[2]
		childX[3] = individualB[3]
		childX[4] = individualB[4]
		childX[5] = individualB[5]
		childX[6] = individualB[6]
		childX[7] = individualB[7]
		childY[0] = individualA[0]
		childY[1] = individualA[1]
		childY[2] = individualA[2]
		childY[3] = individualA[3]
		childY[4] = individualA[4]
		childY[5] = individualA[5]
		childY[6] = individualA[6]
		childY[7] = individualA[7]
	}

	return childX, childY
}

func crossoverByColumn(individualA, individualB [9][9]uint8, crossoverPoint uint8) ([9][9]uint8, [9][9]uint8) {
	childX := individualA
	childY := individualB

	// Do crossover on Row on crossover point
	if crossoverPoint == 1 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
	} else if crossoverPoint == 2 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
		childX, childY = switchGeneByColumn(1, childX, childY)
	} else if crossoverPoint == 3 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
		childX, childY = switchGeneByColumn(1, childX, childY)
		childX, childY = switchGeneByColumn(2, childX, childY)
	} else if crossoverPoint == 4 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
		childX, childY = switchGeneByColumn(1, childX, childY)
		childX, childY = switchGeneByColumn(2, childX, childY)
		childX, childY = switchGeneByColumn(3, childX, childY)
	} else if crossoverPoint == 5 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
		childX, childY = switchGeneByColumn(1, childX, childY)
		childX, childY = switchGeneByColumn(2, childX, childY)
		childX, childY = switchGeneByColumn(3, childX, childY)
		childX, childY = switchGeneByColumn(4, childX, childY)
	} else if crossoverPoint == 6 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
		childX, childY = switchGeneByColumn(1, childX, childY)
		childX, childY = switchGeneByColumn(2, childX, childY)
		childX, childY = switchGeneByColumn(3, childX, childY)
		childX, childY = switchGeneByColumn(4, childX, childY)
		childX, childY = switchGeneByColumn(5, childX, childY)
	} else if crossoverPoint == 7 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
		childX, childY = switchGeneByColumn(1, childX, childY)
		childX, childY = switchGeneByColumn(2, childX, childY)
		childX, childY = switchGeneByColumn(3, childX, childY)
		childX, childY = switchGeneByColumn(4, childX, childY)
		childX, childY = switchGeneByColumn(5, childX, childY)
		childX, childY = switchGeneByColumn(6, childX, childY)
	} else if crossoverPoint == 8 {
		childX, childY = switchGeneByColumn(0, individualA, individualB)
		childX, childY = switchGeneByColumn(1, childX, childY)
		childX, childY = switchGeneByColumn(2, childX, childY)
		childX, childY = switchGeneByColumn(3, childX, childY)
		childX, childY = switchGeneByColumn(4, childX, childY)
		childX, childY = switchGeneByColumn(5, childX, childY)
		childX, childY = switchGeneByColumn(6, childX, childY)
		childX, childY = switchGeneByColumn(7, childX, childY)
	}

	return childX, childY
}

func switchGeneByColumn(column uint8, individualA, individualB [9][9]uint8) ([9][9]uint8, [9][9]uint8) {
	childX := individualA
	childY := individualB

	for i := 0; i < 9; i++ {
		childX[i][column] = individualB[i][column]
		childY[i][column] = individualA[i][column]
	}

	return childX, childY
}
