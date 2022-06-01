package tree

type TreeFactory struct {
	treeTypes []*TreeType
}

func NewTreeFactory() *TreeFactory {
	return &TreeFactory{}
}

func (tf *TreeFactory) GetTreeType(name, color string, texture []byte) *TreeType {
	for _, treeType := range tf.treeTypes {
		if treeType.name == name && treeType.color == color && string(treeType.texture) == string(texture) {
			return treeType
		}
	}

	newTreeType := NewTreeType(name, color, texture)
	tf.treeTypes = append(tf.treeTypes, newTreeType)

	return newTreeType
}
