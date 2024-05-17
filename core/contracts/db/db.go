package db

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

type Store struct {
	Prefix       []byte
	Address      common.Address
	StateDB      vm.StateDB
	KeyEncoder   StoreKeyEncoder
	ValueEncoder StoreValueEncoder
	ValueDecoder StoreValueDecoder
}

func NewStore(prefix []byte, addr common.Address, statedb vm.StateDB) *Store {
	return &Store{
		Prefix:       prefix,
		Address:      addr,
		StateDB:      statedb,
		KeyEncoder:   &DefaultKeyEncoder{},
		ValueEncoder: &DefaultValueEncoder{},
		ValueDecoder: &DefaultValueDecoder{},
	}
}

func (s *Store) SetKeyEncoder(encoder StoreKeyEncoder) {
	s.KeyEncoder = encoder
}

func (s *Store) SetValueEncoder(encoder StoreValueEncoder) {
	s.ValueEncoder = encoder
}

func (s *Store) SetValueDecoder(decoder StoreValueDecoder) {
	s.ValueDecoder = decoder
}

func SetState(store *Store, key any, value any) error {
	keyBuf, err := store.KeyEncoder.EncodeKey(store.Prefix, key)
	if err != nil {
		return nil
	}
	valBuf, err := store.ValueEncoder.EncodeValue(value)
	if err != nil {
		return err
	}
	store.StateDB.SetState(store.Address, keyBuf, valBuf)
	return nil
}

func GetState[T any](store *Store, key any) (T, error) {
	var v T
	keyBuf, err := store.KeyEncoder.EncodeKey(store.Prefix, key)
	if err != nil {
		return v, err
	}

	val := store.StateDB.GetState(store.Address, keyBuf)
	if len(val) == 0 {
		return v, nil
	}
	err = store.ValueDecoder.DecodeValue(val, &v)
	if err != nil {
		return v, err
	}
	return v, nil
}
