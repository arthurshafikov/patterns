package tree

type Tree struct {
	x       int
	y       int
	name    string
	color   string
	texture []byte
}

func NewTree(x, y int, name, color string, texture []byte) *Tree {
	return &Tree{
		x:       x,
		y:       y,
		name:    name,
		color:   color,
		texture: texture,
	}
}

func (t *Tree) Draw() {
	// Draw the tree in canvas or somewhere else...
}
