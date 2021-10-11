package main

import (
	"context"
	"githab.com/Planck1858/chatix/user_service/internal/app"
	"githab.com/Planck1858/chatix/user_service/pkg/logging"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// logger
	zapLog := logging.NewLogger()
	defer zapLog.Sync()

	// main context
	ctx := context.Background()

	// application init
	a := app.NewServer(zapLog)
	err := a.Start(ctx)
	if err != nil {
		zapLog.Fatal(err)
	}

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
	if err != nil {
		zapLog.Fatal(err)
	}

	zapLog.With("signal", stop).Fatal("signal caught")
	zapLog.Info("stopping all jobs")
}
