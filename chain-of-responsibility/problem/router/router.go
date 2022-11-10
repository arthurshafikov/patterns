package router

import (
	"github.com/arthurshafikov/patterns/chain-of-responsibility/problem/models"
	"github.com/arthurshafikov/patterns/chain-of-responsibility/problem/pages"
)

type Dependencies struct {
	AdminPage pages.Page
	IndexPage pages.Page
}

type Router struct {
	deps Dependencies
}

func NewRouter(deps Dependencies) *Router {
	return &Router{
		deps: deps,
	}
}

func (r *Router) InitRoutes() {
	// some router initialization
}

func (r *Router) ShowAdminPage(request *models.Request) *models.Response {
	user := request.User

	// check if user unauthentificated (duplicated)
	if user == nil {
		return models.NewResponse(401, "Unauthorized")
	}

	// check if user is banned (duplicated)
	if user.IsBanned {
		return models.NewResponse(401, "You are banned on this resource")
	}

	// check if user is not an admin
	if !user.IsAdmin {
		return models.NewResponse(403, "Forbidden")
	}

	return models.NewResponse(200, r.deps.AdminPage.Show())
}

func (r *Router) ShowIndexPage(request *models.Request) *models.Response {
	user := request.User

	// check if user unauthentificated (duplicated)
	if user == nil {
		return models.NewResponse(401, "Unauthorized")
	}

	// check if user is banned (duplicated)
	if user.IsBanned {
		return models.NewResponse(401, "You are banned on this resource")
	}

	return models.NewResponse(200, r.deps.IndexPage.Show())
}
