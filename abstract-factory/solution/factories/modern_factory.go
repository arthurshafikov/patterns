package factories

import "github.com/arthurshafikov/patterns/abstract-factory/solution/furniture"

type ModernFactory struct {
}

func NewModernFactory() *ModernFactory {
	return &ModernFactory{}
}

func (a *ModernFactory) CreateChair() Chair {
	return furniture.NewModernChair()
}

func (a *ModernFactory) CreateSofa() Sofa {
	return furniture.NewModernSofa()
}

func (a *ModernFactory) CreateCoffeeTable() CoffeeTable {
	return furniture.NewModernCoffeeTable()
}
