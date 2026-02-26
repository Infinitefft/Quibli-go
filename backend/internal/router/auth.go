package router

import (
	"backend/internal/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func RegisterAuthRoutes(r *server.Hertz, authHandler *handler.AuthHandler) {
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/login", authHandler.Login)
		authGroup.POST("/refresh", authHandler.RefreshToken)
	}
}
