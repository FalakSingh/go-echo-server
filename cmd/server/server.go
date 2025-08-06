package server

import (
	"echo-server/config"
	appMiddleware "echo-server/internal/middleware"
	"echo-server/internal/router"
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo *echo.Echo
}

func NewServer() *Server {
	config.Init()
	config.InitMongoDB()
	e := echo.New()
	e.HTTPErrorHandler = appMiddleware.CustomErrorHandler
	log.Println("Error handler registered")
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())
	router.Init(e)
	log.Println("Router initialized")
	return &Server{
		echo: e,
	}
}

func (s *Server) Start() {
	port := config.Env.PORT
	s.echo.Logger.Fatal(s.echo.Start(":" + port))
}
