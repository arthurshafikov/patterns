package transport

type Transport interface {
	Deliver() string
	EstimateTimeMultiplier() int
}
