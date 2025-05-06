package util

import "ndasbulous/geneticsudoku/constant"

func crossoverTwoChromosomes(individualA, individualB [9][9]uint8, crossoverMethod string, crossoverPoint uint8) ([9][9]uint8, [9][9]uint8) {
	childX := individualA
	childY := individualB

	if crossoverMethod == constant.CROSSOVER_METHOD_ROW {
		// Do crossover on Row on crossover point
		if crossoverPoint == 1 {
			childX[0] = individualB[0]
			childY[0] = individualA[0]
		} else if crossoverPoint == 2 {
			childX[0] = individualB[0]
			childX[1] = individualB[1]
			childY[0] = individualA[0]
			childY[1] = individualA[1]
		}
	}

	return childX, childY
}
