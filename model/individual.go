package model

// Individual represents data about a single individual.
type Individual struct {
	ID          string
	Name        string
	Chromosomes []uint16
}
