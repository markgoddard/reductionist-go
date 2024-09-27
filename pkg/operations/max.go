package operations

import (
	"errors"
	"fmt"

	"github.com/markgoddard/reductionist/pkg/request"
	"golang.org/x/exp/constraints"
)

type Max struct{}

func (max Max) Execute(data []byte, request_data request.Data) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("no elements")
	}
	op, err := makeMaxT(request_data)
	if err != nil {
		return nil, err
	}
	return op.Execute(data, request_data)
}

type maxT[T constraints.Ordered] struct {
}

func makeMaxT(request_data request.Data) (Operation, error) {
	// Dispatch to a specific type.
	switch request_data.Dtype {
	case "int64":
		return maxT[int64]{}, nil
	case "int32":
		return maxT[int32]{}, nil
	case "uint64":
		return maxT[uint64]{}, nil
	case "uint32":
		return maxT[uint64]{}, nil
	case "float64":
		return maxT[float64]{}, nil
	case "float32":
		return maxT[float32]{}, nil
	default:
		return nil, errors.New("unexpected dtype")
	}
}

func (max maxT[T]) Execute(data []byte, request_data request.Data) ([]byte, error) {
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
