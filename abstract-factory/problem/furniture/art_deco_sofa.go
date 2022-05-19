package furniture

import "fmt"

type ArtDecoSofa struct {
}

func NewArtDecoSofa() *ArtDecoSofa {
	return &ArtDecoSofa{}
}

func (c *ArtDecoSofa) LayDown() {
	fmt.Println("Sleeping on Art Deco sofa!")
}
