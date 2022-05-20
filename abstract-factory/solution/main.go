package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/abstract-factory/solution/factories"
)

func main() {
	orders := []string{
		"modern",
		"victorian",
		"art_deco",
	}

	// Normal cases.
	for _, family := range orders {
		var factory factories.Factory

		switch family {
		case "modern":
			factory = factories.NewModernFactory()
		case "victorian":
			factory = factories.NewVictorianFactory()
		case "art_deco":
			factory = factories.NewArtDecoFactory()
		}

		chair := factory.CreateChair()
		fmt.Printf("Chair has legs: %v\n", chair.HasLegs())
		chair.SitOn()

		sofa := factory.CreateSofa()
		sofa.LayDown()

		coffeeTable := factory.CreateCoffeeTable()
		coffeeTable.DrinkCoffeeAt()

		fmt.Printf("Order was sent...\n\n")
	}

	NewFurniture()
}

// New function that was written by new programmer.
func NewFurniture() {
	// probably we can even disallow the ability to create furniture structures anywhere outside factories
	factory := factories.NewModernFactory()

	chair := factory.CreateChair()
	fmt.Printf("Chair has legs: %v\n", chair.HasLegs())
	chair.SitOn()

	sofa := factory.CreateSofa()
	sofa.LayDown()

	coffeeTable := factory.CreateCoffeeTable()
	coffeeTable.DrinkCoffeeAt()
	// all instances are from the same furniture family thanks to the factory
}
