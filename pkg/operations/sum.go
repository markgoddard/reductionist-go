package operations

import (
	"errors"
	"fmt"

	"github.com/markgoddard/reductionist/pkg/request"
)

type Sum struct {
}

func (sum Sum) Execute(data []byte, request_data request.Data) (result []byte, err error) {
	if request_data.Dtype != "int64" {
		return nil, errors.New("unexpected dtype")
	}
	buf, err := bytesToInt64s(data)
	if err != nil {
		return nil, err
	}
	var sum_res uint64
	for _, value := range buf {
		fmt.Println(value)
		sum_res += uint64(value)
	}
	result = uint64ToBytes(sum_res)
	fmt.Println(sum_res, result)
	return
}
