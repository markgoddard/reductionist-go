package operations

import (
	"bytes"
	"encoding/binary"
	"reflect"
)

func bytesToInt64s(data []byte) ([]int64, error) {
	reader := bytes.NewReader(data)
	result := make([]int64, len(data) / 8)
	err := binary.Read(reader, binary.LittleEndian, result)
	return result, err
}

func uint64ToBytes(n uint64) (result []byte) {
	result = make([]byte, 8)
	binary.LittleEndian.PutUint64(result, n)
	return
}

func bytesToTs[T comparable] (data []byte) ([]T, error) {
	var sizer T
	reader := bytes.NewReader(data)
	tSize := reflect.TypeOf(sizer).Size()
	result := make([]T, len(data) / int(tSize))
	err := binary.Read(reader, binary.LittleEndian, result)
	return result, err
}

func tToBytes[T comparable](n T) []byte {
	var result bytes.Buffer
	binary.Write(&result, binary.LittleEndian, n)
	return result.Bytes()
}
