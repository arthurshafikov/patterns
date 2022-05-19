package furniture

import "fmt"

type VictorianSofa struct {
}

func NewVictorianSofa() *VictorianSofa {
	return &VictorianSofa{}
}

func (c *VictorianSofa) LayDown() {
	fmt.Println("Sleeping on Victorian sofa!")
}
