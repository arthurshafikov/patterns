package transport

type Plane struct {
}

func NewPlane() *Plane {
	return &Plane{}
}

func (t *Plane) Deliver() string {
	return "Order has been delivered by sky!"
}

func (t *Plane) EstimateTimeMultiplier() int {
	return 2
}
