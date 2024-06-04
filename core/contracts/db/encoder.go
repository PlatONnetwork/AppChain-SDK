package db

import (
	utils "github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/big"
	"reflect"
)

var (
	bigInt = reflect.TypeOf(big.Int{})
)

type StoreKeyEncoder interface {
	EncodeKey(prefix []byte, key any) ([]byte, error)
}
type DefaultKeyEncoder struct {
}

func (d *DefaultKeyEncoder) EncodeKey(prefix []byte, key any) ([]byte, error) {
	keyBuf, err := encodeKey(key)
	if err != nil {
		return nil, err
	}
	return append(prefix, keyBuf...), nil
}

func encodeKey(key any) ([]byte, error) {
	switch v := key.(type) {
	case *big.Int:
		return math.PaddedBigBytes(v, 32), nil
	case big.Int:
		return math.PaddedBigBytes(&v, 32), nil
	case uint64:
		return utils.EncodeUint64ToBytes(v), nil
	case uint32:
		return utils.EncodeUint32ToBytes(v), nil
	case uint16:
		return utils.EncodeUint16ToBytes(v), nil
	case uint8:
		return []byte{v}, nil
	case bool:
		val := byte(1)
		if !key.(bool) {
			val = 0
		}
		return []byte{val}, nil
	case string:
		return []byte(v), nil
	case common.Address:
		return v.Bytes(), nil
	case common.Hash:
		return v.Bytes(), nil
	case []byte:
		return v, nil
	}
	return rlp.EncodeToBytes(key)
}

type StoreValueEncoder interface {
	EncodeValue(value any) ([]byte, error)
}

type DefaultValueEncoder struct {
}

func (s *DefaultValueEncoder) EncodeValue(value any) ([]byte, error) {
	return rlp.EncodeToBytes(value)
}
