package main

import (
	"github.com/markgoddard/reductionist/pkg/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/sum", handlers.Sum)
	http.ListenAndServe(":8080", nil)
}
