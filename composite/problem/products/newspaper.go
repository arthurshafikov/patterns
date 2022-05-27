package products

type Newspaper struct {
}

func NewNewspaper() *Newspaper {
	return &Newspaper{}
}

func (p *Newspaper) GetPrice() int {
	return 1
}
