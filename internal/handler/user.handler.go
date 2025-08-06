package handler

import (
	services "echo-server/internal/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userSvc services.UserService
}

func NewUserHandler(userSvc services.UserService) *UserHandler {
	return &UserHandler{userSvc}
}

func (h *UserHandler) FetchUsers(c echo.Context) error {
	return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid"})
}
