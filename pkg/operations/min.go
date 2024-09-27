package operations

import (
	"errors"
	"fmt"
	"github.com/markgoddard/reductionist/pkg/request"
)

type Min struct {
}

func (min Min) Execute(data []byte, request_data request.Data) (result []byte, err error) {
	if request_data.Dtype != "int64" {
		return nil, errors.New("unexpected dtype")
	}
	if len(data) == 0 {
		return nil, errors.New("no elements")
	}
	buf, err := bytesToInt64s(data)
	if err != nil {
		return nil, err
	}
	var min_res int64 = buf[0]
	for _, value := range buf[1:] {
		fmt.Println(value)
		if value < min_res {
			min_res = value
		}
	}
	result = uint64ToBytes(uint64(min_res))
	fmt.Println(min_res, result)
	return
}
