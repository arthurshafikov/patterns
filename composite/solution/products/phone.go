package products

type Phone struct {
}

func NewPhone() *Phone {
	return &Phone{}
}

func (p *Phone) GetPrice() int {
	return 15
}
