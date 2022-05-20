package factories

import "github.com/arthurshafikov/patterns/abstract-factory/solution/furniture"

type ArtDecoFactory struct {
}

func NewArtDecoFactory() *ArtDecoFactory {
	return &ArtDecoFactory{}
}

func (a *ArtDecoFactory) CreateChair() Chair {
	return furniture.NewArtDecoChair()
}

func (a *ArtDecoFactory) CreateSofa() Sofa {
	return furniture.NewArtDecoSofa()
}

func (a *ArtDecoFactory) CreateCoffeeTable() CoffeeTable {
	return furniture.NewArtDecoCoffeeTable()
}
