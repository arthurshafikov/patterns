package tree

type Forest struct {
	trees []*Tree

	treeFactory *TreeFactory
}

func NewForest(treeFactory *TreeFactory) *Forest {
	return &Forest{
		treeFactory: treeFactory,
	}
}

func (f *Forest) PlantTree(x, y int, name, color string, texture []byte) {
	treeType := f.treeFactory.GetTreeType(name, color, texture)
	tree := NewTree(x, y, treeType)

	f.trees = append(f.trees, tree)
}

func (f *Forest) Draw() {
	for _, tree := range f.trees {
		tree.Draw()
	}
}
