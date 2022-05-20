package furniture

import "fmt"

type ArtDecoChair struct {
}

func NewArtDecoChair() *ArtDecoChair {
	return &ArtDecoChair{}
}

func (c *ArtDecoChair) SitOn() {
	fmt.Println("What a beautiful Art Deco chair")
}

func (c *ArtDecoChair) HasLegs() bool {
	return true
}
