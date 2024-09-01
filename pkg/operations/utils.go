package operations

import (
	"encoding/binary"
	"bytes"
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
