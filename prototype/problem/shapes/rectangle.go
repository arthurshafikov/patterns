package shapes

import (
	"math/rand"
)

type Rectangle struct {
	uniqueID int // We can't use struct copying because we want this field to stay unique among all shapes.
	X        int
	Y        int
	Width    int
	Height   int

	opacity int // We can't copy opacity outside of this package.

}

func NewRectangle(x, y, width, height int) *Rectangle {
	return &Rectangle{
		uniqueID: rand.Intn(1000), // some ID generation...
		X:        x,
		Y:        y,
		Width:    width,
		Height:   height,
	}
}

func (r *Rectangle) SetOpacity(opacity int) {
	r.opacity = opacity
}
