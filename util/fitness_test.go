package util

import (
	"testing"
)

type testScenario struct {
	testData       [9]uint8
	expectedResult uint8
}

// Test CalculateFitness
func TestCalculateRowOrColumnUniqueness(t *testing.T) {
	scenarioList := []testScenario{
		{
			testData:       [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedResult: 0,
		},
		{
			testData:       [9]uint8{2, 1, 3, 4, 5, 6, 7, 8, 9},
			expectedResult: 0,
		},
		{
			testData:       [9]uint8{9, 2, 3, 4, 5, 6, 7, 8, 1},
			expectedResult: 0,
		},
		{
			testData:       [9]uint8{1, 2, 3, 4, 5, 1, 7, 8, 9},
			expectedResult: 1,
		},
		{
			testData:       [9]uint8{1, 2, 3, 4, 5, 6, 7, 8, 2},
			expectedResult: 1,
		},
		{
			testData:       [9]uint8{5, 2, 3, 4, 5, 6, 7, 8, 9},
			expectedResult: 1,
		},
		{
			testData:       [9]uint8{1, 2, 3, 4, 5, 2, 7, 2, 9},
			expectedResult: 2,
		},
		{
			testData:       [9]uint8{3, 2, 3, 4, 5, 6, 7, 8, 3},
			expectedResult: 2,
		},
		{
			testData:       [9]uint8{1, 8, 3, 3, 5, 8, 7, 8, 9},
			expectedResult: 3,
		},
		{
			testData:       [9]uint8{1, 2, 3, 4, 3, 3, 7, 3, 3},
			expectedResult: 4,
		},
		{
			testData:       [9]uint8{7, 7, 7, 7, 7, 7, 7, 7, 7},
			expectedResult: 8,
		},
	}

	for _, scenario := range scenarioList {
		actual := calculateArrayUniqueness(scenario.testData[0:len(scenario.testData)])
		if scenario.expectedResult != actual {
			t.Errorf(`Expected fitness value doesn't match: Expected %v, Actual %v, testData %v`,
				scenario.expectedResult, actual, scenario.testData)
		}
	}
}
