package tree

type TreeType struct {
	name    string
	color   string
	texture []byte
}

func NewTreeType(name, color string, texture []byte) *TreeType {
	return &TreeType{
		name:    name,
		color:   color,
		texture: texture,
	}
}

func (tt *TreeType) Draw(x, y int) {
	// Draw the tree in canvas or somewhere else...
}
