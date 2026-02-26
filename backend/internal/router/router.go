package router

import (
	"backend/internal/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

type Router struct {
	authHandler *handler.AuthHandler
}

func NewRouter(authHandler *handler.AuthHandler) *Router {
	return &Router{
		authHandler: authHandler,
	}
}

func (r *Router) Register(h *server.Hertz) {
	RegisterAuthRoutes(h, r.authHandler)
}
