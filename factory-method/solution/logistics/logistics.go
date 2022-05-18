package logistics

import (
	"github.com/arthurshafikov/patterns/factory-method/solution/transport"
)

type Logistics interface {
	CreateTransport() transport.Transport
}
