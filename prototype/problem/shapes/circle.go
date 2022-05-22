package shapes

type Circle struct {
	X      int
	Y      int
	Radius int

	opacity int // We can't copy opacity AT ALL from outside of this package.
}

func NewCircle(x, y, radius int) *Circle {
	return &Circle{
		X:      x,
		Y:      y,
		Radius: radius,
	}
}

func (c *Circle) SetOpacity(opacity int) {
	c.opacity = opacity
}
