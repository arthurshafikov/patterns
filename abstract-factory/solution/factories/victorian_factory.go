package factories

import "github.com/arthurshafikov/patterns/abstract-factory/solution/furniture"

type VictorianFactory struct {
}

func NewVictorianFactory() *VictorianFactory {
	return &VictorianFactory{}
}

func (a *VictorianFactory) CreateChair() Chair {
	return furniture.NewVictorianChair()
}

func (a *VictorianFactory) CreateSofa() Sofa {
	return furniture.NewVictorianSofa()
}

func (a *VictorianFactory) CreateCoffeeTable() CoffeeTable {
	return furniture.NewVictorianCoffeeTable()
}
