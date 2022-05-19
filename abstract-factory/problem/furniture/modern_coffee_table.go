package furniture

import "fmt"

type ModernCoffeeTable struct {
}

func NewModernCoffeeTable() *ModernCoffeeTable {
	return &ModernCoffeeTable{}
}

func (c *ModernCoffeeTable) DrinkCoffeeAt() {
	fmt.Println("Drinking coffee at Modern coffee table")
}
