package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/factory-method/solution/logistics"
	"github.com/arthurshafikov/patterns/factory-method/solution/order"
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
		var logis logistics.Logistics

		switch ord.TypeOfDelivery {
		case "ground":
			logis = logistics.NewGroundLogistics()
		case "sea":
			logis = logistics.NewSeaLogistics()
		case "sky":
			logis = logistics.NewSkyLogistics()
		}

		transport := logis.CreateTransport()

		fmt.Println(transport.Deliver())
		fmt.Printf(
			"Estimate time delivery for %v order: %v hours\n",
			ord.TypeOfDelivery,
			transport.EstimateTimeMultiplier()*ord.Distance,
		)
	}
}
