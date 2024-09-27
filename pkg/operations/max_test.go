package operations

import (
	"bytes"
	"encoding/binary"
	"slices"
	"testing"

	"github.com/markgoddard/reductionist/pkg/request"
)

func TestMax(t *testing.T) {
	tests := []struct {
		dtype      string
		data, want []byte
		err        bool
	}{
		{
			"int64",
			tsToBytes([]int64{1, 2, 3, 4}),
			tsToBytes([]int64{4}),
			false,
		},
		{
			"int32",
			tsToBytes([]int32{1, 2, 3, 4}),
			tsToBytes([]int32{4}),
			false,
		},
		{
			"float64",
			tsToBytes([]float64{1, 2, 3, 4}),
			tsToBytes([]float64{4}),
			false,
		},
		// error: no input data
		{
			"int32",
			tsToBytes([]int32{}),
			nil,
			true,
		},
		// error: unexpected dtype
		{
			"unknown",
			nil,
			nil,
			true,
		},
	}
	for _, tt := range tests {
		var op Max
		var request_data request.Data
		request_data.Dtype = tt.dtype
		result, err := op.Execute(tt.data, request_data)
		if err != nil && !tt.err {
			t.Fatalf("got unexpected error %v", err)
		}
		if err == nil && tt.err {
			t.Fatalf("did not get expected error")
		}
		if !tt.err && !slices.Equal(result, tt.want) {
			t.Fatalf("result not equal to input")
		}
	}
}

func tsToBytes[T comparable](data []T) []byte {
	var result bytes.Buffer
	binary.Write(&result, binary.LittleEndian, data)
	return result.Bytes()
}
