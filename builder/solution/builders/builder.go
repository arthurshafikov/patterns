package builders

import "github.com/arthurshafikov/patterns/builder/solution/estates"

type EstateBuilder interface {
	SelectStyle(style string) EstateBuilder
	BuildRooms(numOfRooms int) EstateBuilder
	BuildSwimmingPool() EstateBuilder
	BuildGarage() EstateBuilder
	BuildGarden() EstateBuilder
	BuildFancyStatues() EstateBuilder
	GetResult() estates.Estate
}
