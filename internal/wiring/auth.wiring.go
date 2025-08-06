package wiring

import (
	"echo-server/internal/handler"
	"echo-server/internal/repository"
	"echo-server/internal/service"
)

func AuthWiring() *handler.AuthHandler {
	userRepo := repository.NewUserRepository()
	authService := service.NewAuthService(userRepo)
	return handler.NewAuthHandler(authService)
}
