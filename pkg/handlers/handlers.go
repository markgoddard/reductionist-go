package handlers

import (
	"errors"
	"net/http"
	"os"

	"github.com/markgoddard/reductionist/pkg/operations"
	"github.com/markgoddard/reductionist/pkg/request"
	"github.com/markgoddard/reductionist/pkg/worker"
)

type Operation struct {
	operation operations.Operation
	pool      *worker.Pool
}

func (operation Operation) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	request_data, err := request.FromJson(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data, err := readFile(request_data.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	job := worker.NewJob(operation.operation, data, request_data)
	operation.pool.Execute(&job)
	result, err := job.Wait()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(result)
}

func New(operation operations.Operation, pool *worker.Pool) Operation {
	return Operation{operation: operation, pool: pool}
}

func readFile(path string) (data []byte, err error) {
	data, err = os.ReadFile(path)
	if err != nil {
		return nil, errors.New("failed to open file")
	}
	return
}
