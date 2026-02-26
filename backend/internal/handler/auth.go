package handler

import (
	"context"

	"backend/internal/dto"
	"backend/internal/service"
	"backend/pkg/errors"
	"backend/pkg/response"
	"github.com/cloudwego/hertz/pkg/app"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(ctx context.Context, c *app.RequestContext) {
	var req dto.LoginRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(400, response.BadRequest(err.Error()))
		return
	}

	result, err := h.authService.Login(ctx, req.Phone, req.Password)
	if err != nil {
		if err == errors.ErrInvalidCredentials {
			c.JSON(401, response.Unauthorized(err.Error()))
			return
		}
		c.JSON(500, response.InternalServerError(err.Error()))
		return
	}

	c.JSON(200, response.Success(result))
}

func (h *AuthHandler) RefreshToken(ctx context.Context, c *app.RequestContext) {
	var req dto.RefreshRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(400, response.BadRequest(err.Error()))
		return
	}

	result, err := h.authService.RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		if err == errors.ErrTokenExpired {
			c.JSON(401, response.Unauthorized(err.Error()))
			return
		}
		c.JSON(500, response.InternalServerError(err.Error()))
		return
	}

	c.JSON(200, response.Success(result))
}
