package middleware

import (
	"context"
	"strings"

	"backend/internal/config"
	"backend/pkg/errors"
	"backend/pkg/jwt"
	"github.com/cloudwego/hertz/pkg/app"
)

const (
	UserIDKey   = "user_id"
	UserPhoneKey = "user_phone"
)

func JWT() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(401, errors.ErrUnauthorized)
			c.Abort()
			return
		}

		jwtMgr := jwt.NewManager(config.AppConfig.JWTSecret)
		claims, err := jwtMgr.ValidateToken(tokenString)
		if err != nil {
			c.JSON(401, errors.ErrTokenExpired)
			c.Abort()
			return
		}

		c.Set(UserIDKey, claims.UserID)
		c.Set(UserPhoneKey, claims.Phone)
		c.Next(ctx)
	}
}
