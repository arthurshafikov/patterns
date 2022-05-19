package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/factory-method/problem/order"
	"github.com/arthurshafikov/patterns/factory-method/problem/transport"
)

func main() {
	groundOrder := order.Order{
		TypeOfDelivery: "ground",
		Distance:       10,
	}
	seaOrder := order.Order{
		TypeOfDelivery: "sea",
		Distance:       1000,
	}
	skyOrder := order.Order{
		TypeOfDelivery: "sky",
		Distance:       500,
	}

	for _, ord := range []order.Order{groundOrder, seaOrder, skyOrder} {
		estimateTime(ord)
		deliver(ord)
	}
}

func deliver(order order.Order) {
	var text string

	// in this particular case we actually can move this switch into the cycle inside of main() method
	// but in real programm these switches would be everywhere in different classes etc.
	switch order.TypeOfDelivery {
	case "ground":
		text = transport.NewTruck().DeliverByGround()
	case "sea":
		text = transport.NewShip().DeliverBySea()
	case "sky":
		text = transport.NewPlane().DeliverBySky()
	}

	fmt.Println(text)
}

func estimateTime(order order.Order) {
	var estimateTimeMultiplier int

	switch order.TypeOfDelivery {
	case "ground":
		estimateTimeMultiplier = transport.NewTruck().GetEstimateTimeMultiplier()
	case "sea":
		estimateTimeMultiplier = transport.NewShip().GetEstimateTimeMultiplier()
	case "sky":
		estimateTimeMultiplier = transport.NewPlane().GetEstimateTimeMultiplier()
	}

	fmt.Printf(
		"Estimate time delivery for %v order: %v hours\n",
		order.TypeOfDelivery,
		estimateTimeMultiplier*order.Distance,
	)
}
