package main

import (
	"context"
	"flag"
	"fmt"
	"githab.com/Planck1858/chatix/internal/app"
	"githab.com/Planck1858/chatix/pkg/logging"
	"github.com/pkg/errors"
	"os"
	"os/signal"
	"syscall"
)

const configPath = "./cmd/config.yaml"

func main() {
	// logger
	zapLog := logging.NewLogger()
	defer zapLog.Sync()

	// main context
	ctx := context.Background()

	// init config
	configPath := flag.String("c", configPath, "set config path")
	flag.Parse()

	config, err := loadConfigFromYaml(*configPath)
	errorHandler(zapLog, err)

	// application init
	a := app.NewApplication(zapLog, config)

	err = a.Start(ctx)
	errorHandler(zapLog, err)

	// graceful shutdown
	shutdownSignal := make(chan os.Signal, 1)
	defer close(shutdownSignal)

	signal.Notify(shutdownSignal,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	stop := <-shutdownSignal

	// application shutdown
	err = a.Shutdown()
	errorHandler(zapLog, err)

	zapLog.With("signal", stop).Fatal("signal caught")
	zapLog.Info("stopping all jobs")
}

func loadConfigFromYaml(path string) (*app.Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't read config file: "+path)
	}
	defer f.Close()

	var conf app.Config
	err = yaml.NewDecoder(f).Decode(&conf)
	if err != nil {
		return nil, errors.Wrap(err, "can't parse yaml-file")
	}

	return &conf, nil
}

func errorHandler(log logging.Logger, err error) {
	if err == nil {
		return
	}
	log.Info("application is shutting down with error")
	log.Fatal(fmt.Errorf("%+v", err))
}
