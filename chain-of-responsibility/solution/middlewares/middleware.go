package middlewares

import "github.com/arthurshafikov/patterns/chain-of-responsibility/solution/models"

type Middleware struct {
	NextMiddleware *Middleware
}

type MiddlewareInterface interface {
	Handle(req *models.Request) *models.Response
	SetNext(middleware MiddlewareInterface) MiddlewareInterface
}
