package middlewares

import "github.com/arthurshafikov/patterns/chain-of-responsibility/solution/models"

type CheckAdmin struct {
	NextMiddleware MiddlewareInterface
}

func NewCheckAdmin() *CheckAdmin {
	return &CheckAdmin{}
}

func (m *CheckAdmin) Handle(req *models.Request) *models.Response {
	if !req.User.IsAdmin {
		return models.NewResponse(403, "Forbidden")
	}

	if m.NextMiddleware != nil {
		return m.NextMiddleware.Handle(req)
	}

	return nil
}

func (m *CheckAdmin) SetNext(middleware MiddlewareInterface) MiddlewareInterface {
	m.NextMiddleware = middleware

	return m
}
