package transport

type Truck struct {
}

func NewTruck() *Truck {
	return &Truck{}
}

func (t *Truck) DeliverByGround() string { // classes don't follow the same interface
	return "Order has been delivered by ground!"
}

func (t *Truck) GetEstimateTimeMultiplier() int {
	return 5
}
