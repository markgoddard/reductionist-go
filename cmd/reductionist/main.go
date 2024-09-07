package main

import (
	"fmt"
	"github.com/markgoddard/reductionist/pkg/config"
	"github.com/markgoddard/reductionist/pkg/handlers"
	"github.com/markgoddard/reductionist/pkg/operations"
	"net/http"
)

func main() {
	conf := config.Parse()
	http.Handle("/v1/max", handlers.New(operations.Max{}))
	http.Handle("/v1/min", handlers.New(operations.Min{}))
	http.Handle("/v1/sum", handlers.New(operations.Sum{}))
	addr := fmt.Sprintf(":%d", conf.Port)
	http.ListenAndServe(addr, nil)
}
