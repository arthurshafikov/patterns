package builders

import (
	"github.com/arthurshafikov/patterns/builder/solution/estates"
)

type HouseBuilder struct {
	house *estates.House
}

func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{
		house: estates.NewHouse(),
	}
}

func (h *HouseBuilder) SelectStyle(style string) EstateBuilder {
	h.house.Style = style

	return h
}

func (h *HouseBuilder) BuildRooms(numOfRooms int) EstateBuilder {
	h.house.NumOfRooms = numOfRooms

	return h
}

func (h *HouseBuilder) BuildSwimmingPool() EstateBuilder {
	h.house.HasSwimmingPool = true

	return h
}

func (h *HouseBuilder) BuildGarage() EstateBuilder {
	h.house.HasGarage = true

	return h
}

func (h *HouseBuilder) BuildGarden() EstateBuilder {
	h.house.HasGarden = true

	return h
}

func (h *HouseBuilder) BuildFancyStatues() EstateBuilder {
	h.house.HasFancyStatues = true

	return h
}

func (h *HouseBuilder) GetResult() estates.Estate {
	return h.house
}
