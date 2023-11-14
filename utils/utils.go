package utils

import "encoding/binary"

// EncodeUint64ToBytes encodes provided uint64 to big endian byte slice
func EncodeUint64ToBytes(value uint64) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint64(result, value)

	return result
}

func EncodeUint32ToBytes(value uint32) []byte {
	result := make([]byte, 8)
	binary.BigEndian.PutUint32(result, value)

	return result
}

// EncodeBytesToUint64 big endian byte slice to uint64
func EncodeBytesToUint64(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func EncodeBytesToUint32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}
