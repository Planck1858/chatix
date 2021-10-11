package app

import (
	"context"
	"githab.com/Planck1858/chatix/user_service/pkg/logging"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"time"
)

const (
	appName = "user-service"
	timeOut = time.Second * 15
)

type server struct {
	echo *echo.Echo
	log  logging.Logger
}

func NewServer(log logging.Logger) App {
	e := echo.New()

	e.Server = &http.Server{
		Addr:         ":8000",
		ReadTimeout:  timeOut,
		WriteTimeout: timeOut,
	}

	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	}))

	return App(
		&server{
			echo: e,
			log:  log,
		})
}

func (s *server) Start(ctx context.Context) error {
	s.log.Info(appName)
	s.log.Info("Application initialized")
	//userService := user.New()

	//NewUserController(userService).Mount(s.echo)

	err := s.echo.StartServer(s.echo.Server)
	if err != nil {
		s.log.Fatal(err)

		shutErr := s.echo.Shutdown(context.Background())
		if shutErr != nil {
			s.log.Fatal("server shutdown")
		}
	}

	return err
}

func initDatabaseConnection(cfg *Config) (*sqlx.DB, error) {

	// open database connection pool
	db, err := sqlx.Open("pgx", cfg.Connection)
	if err != nil {
		return nil, errors.Wrap(err, "can't connect to Postgres")
	}

	if cfg.MaxOpenConnections != 0 && cfg.MaxIdleConnections != 0 {
		db.SetMaxIdleConns(cfg.MaxIdleConnections)
		db.SetMaxOpenConns(cfg.MaxOpenConnections)
	}

	return db, nil
}

func (s *server) Shutdown() error {
	err := s.echo.Shutdown(context.Background())
	if err != nil {
		return err
	}

	return nil
}
