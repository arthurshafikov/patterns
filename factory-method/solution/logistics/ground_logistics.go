package logistics

import (
	"github.com/arthurshafikov/patterns/factory-method/solution/transport"
)

type GroundLogistics struct {
	Logistics
}

func NewGroundLogistics() *GroundLogistics {
	return &GroundLogistics{}
}

func (gl *GroundLogistics) CreateTransport() transport.Transport {
	return transport.NewTruck()
}
