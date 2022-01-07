package main

import (
	"context"
	"contrib.go.opencensus.io/integrations/ocsql"
	"database/sql"
	"fmt"
	controllers "githab.com/Planck1858/chatix/internal/api/http"
	"githab.com/Planck1858/chatix/internal/config"
	user_rep "githab.com/Planck1858/chatix/internal/storage/user"
	user_service "githab.com/Planck1858/chatix/internal/user"
	"githab.com/Planck1858/chatix/pkg/logging"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/pkg/errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// logger
	logging.Init()
	log := logging.GetLogger()
	defer log.Sync()

	// init config
	cfg := config.GetConfig()

	// init application
	log.Info("application " + cfg.AppName + " initializing")

	e := echo.New()

	log.Info("http server middlewares initializing...")
	e.Use(
		// http request id
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{Generator: func() string { return uuid.New().String() }}),
		// removes a trailing slash from the request URI
		middleware.RemoveTrailingSlash(),
		// CORS config
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization,
				echo.HeaderCookie, echo.HeaderXRequestID},
			AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		}),
		// logging requests
		middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: "id=${id}:time=${time_unix} - method=${method} - uri=${uri} - status=${status}\n",
		}),
	)

	log.Info("database initializing...")
	pgDb, err := initPostgresDatabaseConnection(&cfg.Db)
	errCheck(log, err)

	log.Info("repositories initializing...")
	userRep := user_rep.NewRepository(pgDb)

	log.Info("services initializing...")
	//jwtHelper := jwt.NewHelper(log)
	userServ := user_service.NewUserService(log, userRep)

	log.Info("server controllers initializing...")
	controllers.NewUserController(log, userServ).Register(e)

	log.Info("http server initializing...")

	httpServer := &http.Server{
		Addr:         cfg.Server.Port,
		ReadTimeout:  time.Second * time.Duration(cfg.Server.TimeOut),
		WriteTimeout: time.Second * time.Duration(cfg.Server.TimeOut),
	}

	log.Info("application initializing is done")

	// application http server start
	go func() {
		log.Info("application http server is starting")
		err = e.StartServer(httpServer)
		errCheck(log, err)
	}()

	// graceful shutdown
	shutdownSignal := make(chan os.Signal, 1)
	defer close(shutdownSignal)

	signal.Notify(shutdownSignal,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	stop := <-shutdownSignal
	log.With("signal", stop).Fatal("signal caught")
	log.Info("stopping all jobs")

	// application shutdown
	log.Info("application http server is shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = e.Shutdown(ctx)
	errCheck(log, err)

	log.Info("database is shutting down")
	err = pgDb.Close()
	errCheck(log, err)

	log.Info("all jobs are done")
}

func initPostgresDatabaseConnection(cfg *config.DbCfg) (*sqlx.DB, error) {
	driverName, err := ocsql.Register("postgres", ocsql.WithAllTraceOptions())
	if err != nil {
		return nil, errors.Wrap(err, "failed init postgres tracer")
	}

	postgresConnect, err := sql.Open(driverName, cfg.PostgresConnection)
	if err != nil {
		return nil, errors.Wrap(err, "failed open postgres connection")
	}

	postgresDb := sqlx.NewDb(postgresConnect, "postgres")
	err = postgresDb.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "failed ping postgres")
	}

	postgresDb.SetMaxIdleConns(cfg.MaxIdleConnections)
	postgresDb.SetMaxOpenConns(cfg.MaxOpenConnections)

	return postgresDb, nil
}

func errCheck(log logging.Logger, err error) {
	if err == nil {
		return
	}
	log.With(fmt.Errorf("%+v", err)).
		Fatal("application is shutting down with error")
}
