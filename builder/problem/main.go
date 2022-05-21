package main

func main() {
	modernHouseWithGarder := NewHouse(ModernStyle, 3, false, false, true, false)
	modernHouseWithGarder.PrintInfo()
	// House is built in modern style, has 3 rooms and has garden.

	cottageWithGarageAndStatues := NewHouse(CottageStyle, 10, false, true, false, true)
	cottageWithGarageAndStatues.PrintInfo()
	// House is built in cottage style, has 10 rooms and also has garage, fancy statues.

	bungalowWithSwimmingPoolGardenAndGarage := NewHouse(CottageStyle, 10, true, true, true, false)
	bungalowWithSwimmingPoolGardenAndGarage.PrintInfo()
	// House is built in cottage style, has 10 rooms and also has swimming pool, garage, garden.

	simpleHouse := NewHouse(ModernStyle, 2, false, false, false, false)
	simpleHouse.PrintInfo()
	// House is built in modern style, has 2 rooms.
}
