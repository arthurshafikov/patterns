package main

import (
	"github.com/arthurshafikov/patterns/builder/solution/builders"
	"github.com/arthurshafikov/patterns/builder/solution/estates"
)

func main() {
	withoutDirectors()
	withDirectors()
}

// One of the way is to use new convenient methods to construct the House instance.
func withoutDirectors() {
	modernHouseWithGarder := builders.NewHouseBuilder().
		SelectStyle(estates.ModernStyle).
		BuildRooms(3).
		BuildGarden().
		GetResult()
	modernHouseWithGarder.PrintInfo()
	// Estate is built in modern style, has 3 rooms and also has garden.

	cottageWithGarageAndStatues := builders.NewHouseBuilder().
		SelectStyle(estates.CottageStyle).
		BuildRooms(10).
		BuildGarage().
		BuildFancyStatues().
		GetResult()
	cottageWithGarageAndStatues.PrintInfo()
	// Estate is built in cottage style, has 10 rooms and also has garage, fancy statues.

	bungalowWithSwimmingPoolGardenAndGarage := builders.NewHouseBuilder().
		SelectStyle(estates.CottageStyle).
		BuildRooms(10).
		BuildSwimmingPool().
		BuildGarage().
		BuildGarden().
		GetResult()
	bungalowWithSwimmingPoolGardenAndGarage.PrintInfo()
	// Estate is built in cottage style, has 10 rooms and also has swimming pool, garage, garden.

	simpleHouse := builders.NewHouseBuilder().
		SelectStyle(estates.ModernStyle).
		BuildRooms(2).
		GetResult()
	simpleHouse.PrintInfo()
	// Estate is built in modern style, has 2 rooms.
}

// In most cases you won't need a director, but if there is a lot of repeated code sometimes it's good to make directors
// and allow them to incapsulate logic of creating objects.
func withDirectors() {
	houseBuilder := builders.NewHouseBuilder()
	hostelBuilder := builders.NewHostelBuilder()

	director := NewDirector()
	director.
		SetBuilder(houseBuilder).
		Create(CheapCost).
		PrintInfo()
	// Estate is built in bungalow style, has 2 rooms.

	director.
		SetBuilder(hostelBuilder).
		Create(ExpensiveCost).
		PrintInfo()
	// Estate is built in cottage style, has 15 rooms and also has swimming pool, garage, garden, fancy statues.
}
