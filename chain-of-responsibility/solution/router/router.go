package router

import (
	"github.com/arthurshafikov/patterns/chain-of-responsibility/solution/middlewares"
	"github.com/arthurshafikov/patterns/chain-of-responsibility/solution/models"
	"github.com/arthurshafikov/patterns/chain-of-responsibility/solution/pages"
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
	auth := middlewares.NewCheckAuth()
	auth.SetNext(middlewares.NewCheckBanned().SetNext(middlewares.NewCheckAdmin())) // we can wrap as many as we want
	if response := auth.Handle(request); response != nil {
		return response
	}

	return models.NewResponse(200, r.deps.AdminPage.Show())
}

func (r *Router) ShowIndexPage(request *models.Request) *models.Response {
	auth := middlewares.NewCheckAuth()
	auth.SetNext(middlewares.NewCheckBanned()) // we can wrap as many as we want
	if response := auth.Handle(request); response != nil {
		return response
	}

	return models.NewResponse(200, r.deps.IndexPage.Show())
}
