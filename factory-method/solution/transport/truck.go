package transport

type Truck struct {
}

func NewTruck() *Truck {
	return &Truck{}
}

func (t *Truck) Deliver() string {
	return "Order has been delivered by ground!"
}

func (t *Truck) EstimateTimeMultiplier() int {
	return 5
}
