package main

import (
	"fmt"
	"ndasbulous/geneticsudoku/model"
)

func main() {
	fmt.Println("Hello, World!")

	var individualA = model.Individual{
		ID:          "1",
		Name:        "John Doe",
		Chromosomes: []uint16{1, 2, 3},
	}

	fmt.Println("Individual A:", individualA)
}
