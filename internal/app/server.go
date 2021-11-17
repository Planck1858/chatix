package app

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"githab.com/Planck1858/chatix/internal/user/rep"
	"githab.com/Planck1858/chatix/pkg/logging"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"net/http"
	"time"
)

type application struct {
	echo *echo.Echo
	log  logging.Logger
	cfg  *Config
}

func NewApplication(log logging.Logger, cfg *Config) App {
	e := echo.New()

	e.Server = &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Second * time.Duration(cfg.Server.TimeOut),
		WriteTimeout: time.Second * time.Duration(cfg.Server.TimeOut),
	}

	e.Use(middleware.Logger())
	e.Use(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
			echo.HeaderCookie, echo.HeaderXRequestID},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
	}))

	return App(
		&application{
			echo: e,
			log:  log,
			cfg:  cfg,
		})
}

func (a *application) Start(ctx context.Context) error {
	a.log.Info("application " + a.cfg.AppName + " initialization")

	// init postgres db
	pg, err := initPostgresDatabaseConnection(a.cfg)
	if err != nil {
		return err
	}

	// init repos
	userRep := rep.NewRepository(pg)

	// init application router

	// init services
	//NewUserController(userService).Mount(a.echo)

	// init http controllers

	// start http application
	a.log.Info("application server is starting")
	err = a.echo.StartServer(a.echo.Server)
	if err != nil {
		a.log.Fatal(err)

		shutErr := a.echo.Shutdown(context.Background())
		if shutErr != nil {
			a.log.Fatal("application shutdown")
		}
	}

	return err
}

func (a *application) Shutdown() error {
	err := a.echo.Shutdown(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func initPostgresDatabaseConnection(cfg *Config) (*sqlx.DB, error) {
	driverName, err := ocsql.Register("postgres", ocsql.WithAllTraceOptions())
	if err != nil {
		return nil, errors.Wrap(err, "failed init postgres tracer")
	}

	postgresConnect, err := sql.Open(driverName, cfg.Db.PostgresConnection)
	if err != nil {
		return nil, errors.Wrap(err, "failed open postgres connection")
	}

	postgresDb := sqlx.NewDb(postgresConnect, "postgres")
	err = postgresDb.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed ping postgres")
	}

	postgresDb.SetMaxIdleConns(cfg.Db.MaxIdleConnections)
	postgresDb.SetMaxOpenConns(cfg.Db.MaxOpenConnections)

	return postgresDb, nil
}
