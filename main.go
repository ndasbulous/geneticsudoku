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
		fmt.Println(individual.Chromosome.Sequence[0])
		fmt.Println(individual.Chromosome.Sequence[1])
		fmt.Println(individual.Chromosome.Sequence[2])
		fmt.Println(individual.Chromosome.Sequence[3])
		fmt.Println(individual.Chromosome.Sequence[4])
		fmt.Println(individual.Chromosome.Sequence[5])
		fmt.Println(individual.Chromosome.Sequence[6])
		fmt.Println(individual.Chromosome.Sequence[7])
		fmt.Println(individual.Chromosome.Sequence[8])
		fmt.Println("fitness: ", individual.Fitness)
		fmt.Println("")
	}
}
