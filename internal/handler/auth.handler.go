package handler

import (
	"echo-server/internal/model"
	"echo-server/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authSvc service.AuthService
}

func NewAuthHandler(authSvc service.AuthService) *AuthHandler {
	return &AuthHandler{authSvc: authSvc}
}

func (h *AuthHandler) Register(c echo.Context) error {
	user := &model.User{}
	if err := c.Bind(user); err != nil {
		return err
	}
	err := h.authSvc.CreateUser(user)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, echo.Map{"error": "invalid"})
}
