package main

import (
	"fmt"
	"ndasbulous/geneticsudoku/util"
)

func main() {
	fmt.Println("Hello, World!")

	var individualA = util.GenerateRandomIndividual()

	fmt.Println("Individual A:", individualA)

	var populationA = util.GenerateInitialPopulation(10)
	for _, individual := range populationA.Individuals {
		fmt.Println(individual.ID)
		fmt.Println(individual.Name)
		fmt.Println(individual.Chromosome)
		fmt.Println(individual.Fitness)
	}
}
