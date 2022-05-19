package furniture

import "fmt"

type VictorianChair struct {
}

func NewVictorianChair() *VictorianChair {
	return &VictorianChair{}
}

func (c *VictorianChair) SitOn() {
	fmt.Println("What a beautiful Victorian chair")
}

func (c *VictorianChair) HasLegs() bool {
	return true
}
