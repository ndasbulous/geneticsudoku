package util

import (
	"ndasbulous/geneticsudoku/model"
	"testing"

	"github.com/google/uuid"
)

type testRowOrColumnFitnessScenario struct {
	testData       [9]uint8
	expectedResult uint8
}

type testOverallFitnessScenario struct {
	testData       [9][9]uint8
	expectedResult uint8
}

// Test CalculateFitness
func TestCalculateRowOrColumnUniqueness(t *testing.T) {
	scenarioList := []testRowOrColumnFitnessScenario{
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

func TestCalculateFitness(t *testing.T) {
	scenarioList := []testOverallFitnessScenario{
		{
			testData: [9][9]uint8{
				{4, 3, 5, 2, 6, 9, 7, 8, 1},
				{6, 8, 2, 5, 7, 1, 4, 9, 3},
				{1, 9, 7, 8, 3, 4, 5, 6, 2},
				{8, 2, 6, 1, 9, 5, 3, 4, 7},
				{3, 7, 4, 6, 8, 2, 9, 1, 5},
				{9, 5, 1, 7, 4, 3, 6, 2, 8},
				{5, 1, 9, 3, 2, 6, 8, 7, 4},
				{2, 4, 8, 9, 5, 7, 1, 3, 6},
				{7, 6, 3, 4, 1, 8, 2, 5, 9},
			},
			expectedResult: 0,
		},
	}

	for _, scenario := range scenarioList {
		individualToTest := model.Individual{
			ID:         uuid.New(), // This should be generated dynamically
			Name:       "John Doe",
			Chromosome: model.Chromosome{scenario.testData},
			Fitness:    0,
		}

		actual := CalculateFitness(individualToTest)
		if scenario.expectedResult != actual {
			t.Errorf(`Expected fitness value doesn't match: Expected %v, Actual %v, testData %v`,
				scenario.expectedResult, actual, scenario.testData)
		}
	}
}
