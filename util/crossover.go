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
	}

	return childX, childY
}
