package operations

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/markgoddard/reductionist/pkg/request"
	"io"
)

type Min struct {
}

func (min Min) Execute(data []byte, request_data request.Data) (result []byte, err error) {
	if request_data.Dtype != "int64" {
		return nil, errors.New("Unexpected dtype")
	}
	if len(data) == 0 {
		return nil, errors.New("No elements")
	}
	reader := bytes.NewReader(data)
	var min_res int64
	first := true
	for {
		var value int64
		err := binary.Read(reader, binary.LittleEndian, &value)
		if err != nil {
			if err != io.EOF {
				return nil, err
			}
			break
		}
		fmt.Println(value)
		if first {
			first = false
			min_res = value
		} else if value < min_res {
			min_res = value
		}
	}
	result = uint64ToBytes(uint64(min_res))
	fmt.Println(min_res, result)
	return
}
