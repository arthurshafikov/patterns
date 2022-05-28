package main

import "github.com/arthurshafikov/patterns/composite/solution/products"

type Box struct {
	Elements []products.HasPrice
}

func NewBox(elements []products.HasPrice) *Box {
	return &Box{
		Elements: elements,
	}
}

func (b *Box) GetPrice() int {
	var total int

	for _, elem := range b.Elements {
		total += elem.GetPrice()
	}

	return total
}
