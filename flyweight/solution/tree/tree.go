package tree

type Tree struct {
	x        int
	y        int
	treeType *TreeType
}

func NewTree(x, y int, treeType *TreeType) *Tree {
	return &Tree{
		x:        x,
		y:        y,
		treeType: treeType,
	}
}

func (t *Tree) Draw() {
	t.treeType.Draw(t.x, t.y)
}
