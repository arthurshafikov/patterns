package furniture

import "fmt"

type VictorianCoffeeTable struct {
}

func NewVictorianCoffeeTable() *VictorianCoffeeTable {
	return &VictorianCoffeeTable{}
}

func (c *VictorianCoffeeTable) DrinkCoffeeAt() {
	fmt.Println("Drinking coffee at Victorian coffee table")
}
