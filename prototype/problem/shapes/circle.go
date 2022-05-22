package shapes

import (
	"math/rand"
)

type Circle struct {
	uniqueID int // We can't use struct copying because we want this field to stay unique among all shapes.
	X        int
	Y        int
	Radius   int

	opacity int // We can't copy opacity outside of this package.
}

func NewCircle(x, y, radius int) *Circle {
	return &Circle{
		uniqueID: rand.Intn(1000), // some ID generation...
		X:        x,
		Y:        y,
		Radius:   radius,
	}
}

func (c *Circle) SetOpacity(opacity int) {
	c.opacity = opacity
}
