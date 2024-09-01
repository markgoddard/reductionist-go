package handlers

import (
	"errors"
	"github.com/markgoddard/reductionist/pkg/operations"
	"github.com/markgoddard/reductionist/pkg/request"
	"net/http"
	"os"
)

func Sum(w http.ResponseWriter, req *http.Request) {
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
	operation := operations.Sum{}
	result, err := operation.Execute(data, request_data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(result)
}

func readFile(path string) (data []byte, err error) {
	data, err = os.ReadFile(path)
	if err != nil {
		return nil, errors.New("Failed to open file")
	}
	return
}
