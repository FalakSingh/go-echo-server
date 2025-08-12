package router

import (
	"echo-server/internal/wiring"

	"github.com/labstack/echo/v4"
)

func AuthRouter(router *echo.Group) {
	authHandler := wiring.AuthWiring()

	authRouter := router.Group("/auth")

	authRouter.POST("/register", authHandler.Register)
	authRouter.POST("/login", authHandler.Login)

}
