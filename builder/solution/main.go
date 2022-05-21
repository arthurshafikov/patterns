package main

import (
	"github.com/arthurshafikov/patterns/builder/solution/estates"
)

func main() {
	withoutDirectors()
}

// One of the way is to use new convenient methods to construct the House instance.
func withoutDirectors() {
	modernHouseWithGarder := estates.NewHouse().
		SelectStyle(estates.ModernStyle).
		BuildRooms(3).
		BuildGarden()
	modernHouseWithGarder.PrintInfo()
	// House is built in modern style, has 3 rooms and has garden.

	cottageWithGarageAndStatues := estates.NewHouse().
		SelectStyle(estates.CottageStyle).
		BuildRooms(10).
		BuildGarage().
		BuildFancyStatues()
	cottageWithGarageAndStatues.PrintInfo()
	// House is built in cottage style, has 10 rooms and also has garage, fancy statues.

	bungalowWithSwimmingPoolGardenAndGarage := estates.NewHouse().
		SelectStyle(estates.CottageStyle).
		BuildRooms(10).
		BuildSwimmingPool().
		BuildGarage().
		BuildGarden()
	bungalowWithSwimmingPoolGardenAndGarage.PrintInfo()
	// House is built in cottage style, has 10 rooms and also has swimming pool, garage, garden.

	simpleHouse := estates.NewHouse().
		SelectStyle(estates.ModernStyle).
		BuildRooms(2)
	simpleHouse.PrintInfo()
	// House is built in modern style, has 2 rooms.
}
