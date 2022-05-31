package tree

type Forest struct {
	trees []*Tree
}

func NewForest() *Forest {
	return &Forest{}
}

func (f *Forest) PlantTree(x, y int, name, color string, texture []byte) {
	tree := NewTree(x, y, name, color, texture)

	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}
