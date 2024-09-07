package operations

import (
	"errors"
	"fmt"
	"github.com/markgoddard/reductionist/pkg/request"
)

type Max struct {
}

func (max Max) Execute(data []byte, request_data request.Data) (result []byte, err error) {
	if request_data.Dtype != "int64" {
		return nil, errors.New("Unexpected dtype")
	}
	if len(data) == 0 {
		return nil, errors.New("No elements")
	}
	buf, err := bytesToInt64s(data)
	if err != nil {
		return nil, err
	}
	var max_res int64 = buf[0]
	for _, value := range buf[1:] {
		fmt.Println(value)
		if value > max_res {
			max_res = value
		}
	}
	result = uint64ToBytes(uint64(max_res))
	fmt.Println(max_res, result)
	return
}
