package shapes

type Rectangle struct {
	X      int
	Y      int
	Width  int
	Height int

	opacity int // We can't copy opacity AT ALL from outside of this package.
}

func NewRectangle(x, y, width, height int) *Rectangle {
	return &Rectangle{
		X:      x,
		Y:      y,
		Width:  width,
		Height: height,
	}
}

func (r *Rectangle) SetOpacity(opacity int) {
	r.opacity = opacity
}
