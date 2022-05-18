package transport

type Ship struct {
}

func NewShip() *Ship {
	return &Ship{}
}

func (t *Ship) DeliverBySea() string { // classes don't follow the same interface
	return "Order has been delivered by sea!"
}

func (t *Ship) GetEstimateTimeMultiplier() int {
	return 25
}
