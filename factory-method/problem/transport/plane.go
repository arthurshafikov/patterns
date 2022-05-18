package transport

type Plane struct {
}

func NewPlane() *Plane {
	return &Plane{}
}

func (t *Plane) DeliverBySky() string { // classes don't follow the same interface
	return "Order has been delivered by sky!"
}

func (t *Plane) GetEstimateTimeMultiplier() int {
	return 2
}
