package main

import (
	"github.com/Planck1858/chatix/api"
	"github.com/Planck1858/chatix/pkg/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"net/http"
)

type server struct {
	echo *echo.Echo
	log *zap.Logger
}

func NewServer(log *zap.Logger) *server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT,  echo.DELETE},
	}))

	e.Validator = NewValidator()

	return &server{
		echo:       e,
		log:        log,
	}
}

func (s *server) Start() error {
	userService := user.New()

	api.NewUserController(userService).Mount(s.echo)

	go func() {
		if err := s.echo.Start(":8000"); err != nil && err != http.ErrServerClosed {
			s.log.Fatal("shutting down the server")
		}
	}()

	return nil
}