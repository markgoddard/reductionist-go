package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RequestData struct {
	Path  string `json:"path"`
	Dtype string `json:"dtype"`
}

func sum(w http.ResponseWriter, req *http.Request) {
	var request_data RequestData
	if err := json.NewDecoder(req.Body).Decode(&request_data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(request_data.Path) == 0 {
		http.Error(w, "Path not provided", http.StatusBadRequest)
		return
	}
	if request_data.Dtype != "int64" {
		http.Error(w, "Unexpected dtype", http.StatusBadRequest)
		return
	}
	data, err := os.ReadFile(request_data.Path)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusBadRequest)
		return
	}
	reader := bytes.NewReader(data)
	var sum uint64
	for {
		var value int64
		err := binary.Read(reader, binary.LittleEndian, &value)
		if err != nil {
			if err != io.EOF {
				http.Error(w, "Failed to read int", http.StatusBadRequest)
			}
			break
		}
		fmt.Println(value)
		sum += uint64(value)
	}
	result := make([]byte, 8)
	binary.LittleEndian.PutUint64(result, sum)
	fmt.Println(sum, result)
	w.Write(result)
}

func main() {
	http.HandleFunc("/v1/sum", sum)
	http.ListenAndServe(":8080", nil)
}
