package util

import (
	"testing"
)

type testScenario struct {
	testData       [9]uint8
	expectedResult float64
}

// Test CalculateFitness
func TestCalculateRowOrColumnUniqueness(t *testing.T) {
	scenarioList := []testScenario{
		{
			testData:       [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedResult: 1,
		},
		{
			testData:       [9]uint8{2, 1, 3, 4, 5, 6, 7, 8, 9},
			expectedResult: 1,
		},
		{
			testData:       [9]uint8{9, 2, 3, 4, 5, 6, 7, 8, 1},
			expectedResult: 1,
		},
	}

	for _, scenario := range scenarioList {
		actual := calculateRowOrColumnUniqueness(scenario.testData)
		if scenario.expectedResult != actual {
			t.Errorf(`Expected fitness value doesn't match: Expected %f, Actual %f, testData %v`,
				scenario.expectedResult, actual, scenario.testData)
		}
	}
}
