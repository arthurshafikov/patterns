package pages

type AdminPage struct {
}

func NewAdminPage() *AdminPage {
	return &AdminPage{}
}

func (a *AdminPage) Show() string {
	return "Some HTML code of an admin page that should not be showed to users..."
}
