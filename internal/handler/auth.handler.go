package handler

import (
	"echo-server/internal/dto"
	"echo-server/internal/helper/response"
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
	registerPayload := &dto.RegisterPayload{}
	if err := c.Bind(registerPayload); err != nil {
		return err
	}
	user := &model.User{Email: registerPayload.Email, Name: registerPayload.Name, Password: registerPayload.Password}
	err := h.authSvc.CreateUser(user)
	if err != nil {
		return err
	}
	return response.Success(c, http.StatusCreated, "User created successfully", user)
}

func (h *AuthHandler) Login(c echo.Context) error {
	loginPayload := &dto.LoginPayload{}
	if err := c.Bind(loginPayload); err != nil {
		return err
	}
	user := &model.User{Email: loginPayload.Email, Password: loginPayload.Password}
	err := h.authSvc.LoginUser(user)
	if err != nil {
		return err
	}
	return response.Success(c, http.StatusOK, "Login successful", user)
}
