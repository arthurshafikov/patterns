package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/abstract-factory/problem/furniture"
)

func main() {
	orders := []string{
		"modern",
		"victorian",
		"art_deco",
	}

	for _, order := range orders {
		switch order {
		case "modern":
			chair := furniture.NewModernChair()
			fmt.Printf("Chair has legs: %v\n", chair.HasLegs())
			chair.SitOn()

			sofa := furniture.NewModernSofa()
			sofa.LayDown()

			coffeeTable := furniture.NewModernCoffeeTable()
			coffeeTable.DrinkCoffeeAt()
		case "victorian":
			chair := furniture.NewVictorianChair()
			fmt.Printf("Chair has legs: %v\n", chair.HasLegs())
			chair.SitOn()

			sofa := furniture.NewVictorianSofa()
			sofa.LayDown()

			coffeeTable := furniture.NewVictorianCoffeeTable()
			coffeeTable.DrinkCoffeeAt()
		case "art_deco":
			chair := furniture.NewArtDecoChair()
			fmt.Printf("Chair has legs: %v\n", chair.HasLegs())
			chair.SitOn()

			sofa := furniture.NewArtDecoSofa()
			sofa.LayDown()

			coffeeTable := furniture.NewArtDecoCoffeeTable()
			coffeeTable.DrinkCoffeeAt()
		}

		fmt.Printf("Order was sent...\n\n")
	}
}
