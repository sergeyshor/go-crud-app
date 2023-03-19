package middleware

import "go-crud-app/internal/usecase"

type Middlewares struct {
	userMiddlewares
}

func New(u usecase.User, t usecase.Tweet) *Middlewares {
	return &Middlewares{
		userMiddlewares{u},
	}
}
