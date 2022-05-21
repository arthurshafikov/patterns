package estates

type House struct {
	AbstractEstate
}

func NewHouse() *House {
	return &House{}
}
