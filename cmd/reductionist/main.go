package main

import (
	"fmt"
	"net/http"

	"github.com/markgoddard/reductionist/pkg/config"
	"github.com/markgoddard/reductionist/pkg/handlers"
	"github.com/markgoddard/reductionist/pkg/operations"
	"github.com/markgoddard/reductionist/pkg/worker"
)

func main() {
	conf := config.Parse()
	pool := worker.NewPool(10)
	http.Handle("/v1/max", handlers.New(operations.Max{}, &pool))
	http.Handle("/v1/min", handlers.New(operations.Min{}, &pool))
	http.Handle("/v1/sum", handlers.New(operations.Sum{}, &pool))
	addr := fmt.Sprintf(":%d", conf.Port)
	http.ListenAndServe(addr, nil)
}
