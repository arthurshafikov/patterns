package shapes

import "math/rand"

type Shape interface {
	Clone() Shape
	SetOpacity(opacity int)
}

func generateUniqueID() int {
	return rand.Intn(1000) // some ID generation...
}
