package middlewares

import "github.com/arthurshafikov/patterns/chain-of-responsibility/solution/models"

type CheckAuth struct {
	NextMiddleware MiddlewareInterface
}

func NewCheckAuth() *CheckAuth {
	return &CheckAuth{}
}

func (m *CheckAuth) Handle(req *models.Request) *models.Response {
	if req.User == nil {
		return models.NewResponse(401, "Unauthorized")
	}

	if m.NextMiddleware != nil {
		return m.NextMiddleware.Handle(req)
	}

	return nil
}

func (m *CheckAuth) SetNext(middleware MiddlewareInterface) MiddlewareInterface {
	m.NextMiddleware = middleware

	return m
}
