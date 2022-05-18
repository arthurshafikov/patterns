package logistics

import (
	"github.com/arthurshafikov/patterns/factory-method/solution/transport"
)

type SkyLogistics struct {
	Logistics
}

func NewSkyLogistics() *SkyLogistics {
	return &SkyLogistics{}
}

func (gl *SkyLogistics) CreateTransport() transport.Transport {
	return transport.NewPlane()
}
