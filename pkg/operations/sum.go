package operations

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/markgoddard/reductionist/pkg/request"
	"io"
)

type Sum struct {
}

func (sum Sum) Execute(data []byte, request_data request.Data) (result []byte, err error) {
	if request_data.Dtype != "int64" {
		return nil, errors.New("Unexpected dtype")
	}
	reader := bytes.NewReader(data)
	var sum_res uint64
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
		sum_res += uint64(value)
	}
	result = uint64ToBytes(sum_res)
	fmt.Println(sum_res, result)
	return
}

func uint64ToBytes(n uint64) (result []byte) {
	result = make([]byte, 8)
	binary.LittleEndian.PutUint64(result, n)
	return
}
