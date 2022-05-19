package furniture

import "fmt"

type ModernChair struct {
}

func NewModernChair() *ModernChair {
	return &ModernChair{}
}

func (c *ModernChair) SitOn() {
	fmt.Println("What a beautiful Modern chair")
}

func (c *ModernChair) HasLegs() bool {
	return false
}
