package middlewares

import "github.com/arthurshafikov/patterns/chain-of-responsibility/solution/models"

type CheckBanned struct {
	NextMiddleware MiddlewareInterface
}

func NewCheckBanned() *CheckBanned {
	return &CheckBanned{}
}

func (m *CheckBanned) Handle(req *models.Request) *models.Response {
	if req.User.IsBanned {
		return models.NewResponse(401, "You are banned on this resource")
	}

	if m.NextMiddleware != nil {
		return m.NextMiddleware.Handle(req)
	}

	return nil
}

func (m *CheckBanned) SetNext(middleware MiddlewareInterface) MiddlewareInterface {
	m.NextMiddleware = middleware

	return m
}
