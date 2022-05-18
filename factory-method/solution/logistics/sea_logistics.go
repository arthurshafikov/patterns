package logistics

import (
	"github.com/arthurshafikov/patterns/factory-method/solution/transport"
)

type SeaLogistics struct {
	Logistics
}

func NewSeaLogistics() *SeaLogistics {
	return &SeaLogistics{}
}

func (gl *SeaLogistics) CreateTransport() transport.Transport {
	return transport.NewShip()
}
