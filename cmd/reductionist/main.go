package main

import (
	"github.com/markgoddard/reductionist/pkg/handlers"
	"github.com/markgoddard/reductionist/pkg/operations"
	"net/http"
)

func main() {
	http.Handle("/v1/min", handlers.New(operations.Min{}))
	http.Handle("/v1/sum", handlers.New(operations.Sum{}))
	http.ListenAndServe(":8080", nil)
}
