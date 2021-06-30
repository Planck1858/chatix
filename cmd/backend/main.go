package main

import "sync"

//import (
//	"context"
//	"go.uber.org/zap"
//	"log"
//	"os"
//	"os/signal"
//	"syscall"
//	"time"
//)
//
//func main() {
//	zapLog, err := zap.NewDevelopment()
//	if err != nil {
//		log.Fatal("error on zap logger init: " + err.Error())
//	}
//	defer zapLog.Sync()
//	log := zapLog.Sugar()
//
//	server := NewServer(zapLog)
//
//	err = server.Start()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals,
//		syscall.SIGHUP,
//		syscall.SIGINT,
//		syscall.SIGTERM,
//		syscall.SIGQUIT)
//	stop := <-signals
//	println(stop)
//
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	if err := server.echo.Shutdown(ctx); err != nil {
//		log.Fatal(err)
//	}
//}
const N = 10

func main() {
	m := make(map[int]int)
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			m[i] = i
			mu.Unlock()
		}()
	}

	wg.Wait()
	println(len(m))
}
