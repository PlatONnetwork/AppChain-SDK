package db

import "github.com/PlatONnetwork/PlatON-Go/rlp"

type StoreValueDecoder interface {
	DecodeValue(b []byte, val any) error
}

type DefaultValueDecoder struct {
}

func (s *DefaultValueDecoder) DecodeValue(b []byte, val any) error {
	return rlp.DecodeBytes(b, val)
}
