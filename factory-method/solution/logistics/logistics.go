package logistics

import (
	"github.com/arthurshafikov/patterns/factory-method/solution/transport"
)

type Logistics interface {
	CreateTransport() transport.Transport
}

func GetLogisticsByOrderType(orderType string) Logistics {
	switch orderType {
	case "ground":
		return NewGroundLogistics()
	case "sea":
		return NewSeaLogistics()
	case "sky":
		return NewSkyLogistics()
	default:
		panic("Not defined type of delivery")
	}
}
