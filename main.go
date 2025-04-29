package main

import (
	"fmt"
	"ndasbulous/geneticsudoku/model"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello, World!")

	var individualA = model.Individual{
		ID:          uuid.New(), // This should be generated dynamically
		Name:        "John Doe",
		Chromosomes: []uint16{1, 2, 3},
	}

	fmt.Println("Individual A:", individualA)
}
