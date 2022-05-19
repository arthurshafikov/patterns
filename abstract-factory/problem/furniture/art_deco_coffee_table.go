package furniture

import "fmt"

type ArtDecoCoffeeTable struct {
}

func NewArtDecoCoffeeTable() *ArtDecoCoffeeTable {
	return &ArtDecoCoffeeTable{}
}

func (c *ArtDecoCoffeeTable) DrinkCoffeeAt() {
	fmt.Println("Drinking coffee at Art Deco coffee table")
}
