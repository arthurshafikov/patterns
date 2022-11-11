package pages

type IndexPage struct {
}

func NewIndexPage() *IndexPage {
	return &IndexPage{}
}

func (i *IndexPage) Show() string {
	return "Some casual HTML code of a simple page that can be shown to all authorized users"
}
