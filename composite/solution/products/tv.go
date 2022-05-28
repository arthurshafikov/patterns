package products

type TV struct {
}

func NewTV() *TV {
	return &TV{}
}

func (p *TV) GetPrice() int {
	return 50
}
