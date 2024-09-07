package operations

import (
	"errors"
	"fmt"
	"github.com/markgoddard/reductionist/pkg/request"
	"golang.org/x/exp/constraints"
)

type Max struct {
}

func (max Max) Execute(data []byte, request_data request.Data) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("No elements")
	}
	return dispatch(data, request_data)
}

func dispatch(data []byte, request_data request.Data) ([]byte, error) {
	switch request_data.Dtype {
	case "int64":
		return doMax[int64](data, request_data)
	case "int32":
		return doMax[int32](data, request_data)
	case "uint64":
		return doMax[uint64](data, request_data)
	case "uint32":
		return doMax[uint64](data, request_data)
	case "float64":
		return doMax[float64](data, request_data)
	case "float32":
		return doMax[float32](data, request_data)
	default:
		return nil, errors.New("Unexpected dtype")
	}
}

func doMax[T constraints.Ordered](data []byte, request_data request.Data) ([]byte, error) {
	buf, err := bytesToTs[T](data)
	if err != nil {
		return nil, err
	}
	var max_res T = buf[0]
	for _, value := range buf[1:] {
		fmt.Println(value)
		if value > max_res {
			max_res = value
		}
	}
	result := tToBytes(max_res)
	fmt.Println(max_res, result)
	return result, nil
}
