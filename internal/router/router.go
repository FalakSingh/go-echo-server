package router

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo) {
	router := e.Group("/api")
	AuthRouter(router)
}
