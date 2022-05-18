package transport

type Ship struct {
}

func NewShip() *Ship {
	return &Ship{}
}

func (t *Ship) Deliver() string {
	return "Order has been delivered by sea!"
}

func (t *Ship) EstimateTimeMultiplier() int {
	return 25
}
