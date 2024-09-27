package operations

import (
	"slices"
	"testing"

	"github.com/markgoddard/reductionist/pkg/request"
)

func TestMax(t *testing.T) {
	var op Max
	data := [...]byte{1, 2, 3, 4, 5, 6, 7, 8}
	var request_data request.Data
	request_data.Dtype = "int64"
	result, err := op.Execute(data[:], request_data)
	if err != nil {
		t.Fatalf("got error %v", err)
	}
	if !slices.Equal(result, data[:]) {
		t.Fatalf("result not equal to input")
	}
}
