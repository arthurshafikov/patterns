package products

type Hammer struct {
}

func NewHammer() *Hammer {
	return &Hammer{}
}

func (h *Hammer) GetPrice() int {
	return 5
}
