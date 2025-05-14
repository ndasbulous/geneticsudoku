package util

import (
	"math/rand"
)

// Mutate individuals by flipping a random bit except for the predefined values
func Mutate(individual, predefinedValues [9][9]uint8) [9][9]uint8 {
	for i := 0; i < len(individual); i++ {
		for j := 0; j < len(individual[i]); j++ {
			if predefinedValues[i][j] == 0 {
				individual[i][j] = uint8(rand.Intn(9) + 1) // Random number between 1 and 9
			}
		}
	}
	return individual
}
