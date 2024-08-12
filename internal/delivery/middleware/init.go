package middleware

import (
	"golang_boilerplate/domain/delivery"
)

type (
	Middleware struct{}
)

func New() delivery.IMiddleware {
	return &Middleware{}
}
