package main

import (
	"fmt"
	"net/http"

	"github.com/markgoddard/reductionist/pkg/config"
	"github.com/markgoddard/reductionist/pkg/handlers"
	"github.com/markgoddard/reductionist/pkg/operations"
	"github.com/markgoddard/reductionist/pkg/worker"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.Parse()
	pool := worker.NewPool(10)
	handleShutdown(pool)
	http.Handle("/v1/max", handlers.New(operations.Max{}, pool))
	http.Handle("/v1/min", handlers.New(operations.Min{}, pool))
	http.Handle("/v1/sum", handlers.New(operations.Sum{}, pool))
	addr := fmt.Sprintf(":%d", conf.Port)
	http.ListenAndServe(addr, nil)
}

func handleShutdown(pool *worker.Pool) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Caught Ctrl-C")
		pool.Close()
		fmt.Println("Waiting for workers to exit")
		pool.Join()
		os.Exit(1)
	}()
}
