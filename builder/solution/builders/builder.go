package builders

import "github.com/arthurshafikov/patterns/builder/solution/estates"

type Builder interface {
	SelectStyle(style string) Builder
	BuildRooms(numOfRooms int) Builder
	BuildSwimmingPool() Builder
	BuildGarage() Builder
	BuildGarden() Builder
	BuildFancyStatues() Builder
	GetResult() estates.Estate
}
