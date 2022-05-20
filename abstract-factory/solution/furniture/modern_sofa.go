package furniture

import "fmt"

type ModernSofa struct {
}

func NewModernSofa() *ModernSofa {
	return &ModernSofa{}
}

func (c *ModernSofa) LayDown() {
	fmt.Println("Sleeping on Modern sofa!")
}
