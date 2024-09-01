package request

import (
	"encoding/json"
	"errors"
	"io"
)

type Data struct {
	Path  string `json:"path"`
	Dtype string `json:"dtype"`
}

func FromJson(r io.Reader) (data Data, err error) {
	err = json.NewDecoder(r).Decode(&data)
	if err != nil {
		return
	}
	if len(data.Path) == 0 {
		err = errors.New("Path not provided")
		return
	}
	if len(data.Dtype) == 0 {
		err = errors.New("Dtype not provided")
		return
	}
	return
}
