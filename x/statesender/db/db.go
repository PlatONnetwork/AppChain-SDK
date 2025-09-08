package db

import (
	"errors"
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	ErrStoreFailed  = errors.New("store failed")
	ErrRlpEncode    = errors.New("rlp encode failed")
	ErrRlpDecode    = errors.New("rlp decode failed")
	ErrNotFound     = errors.New("not found")
	ErrExist        = errors.New("already exist")
	ErrMisMatching  = errors.New("mismatching")
	ErrInvalidValue = errors.New("invalid value")
)

var (
	l2StateSenderCounterKey = []byte("l2StateSenderCounter") // "l2StateSenderCounter" => counter (It is a self increasing counter of l2 stateSender)
)

func IncrementCounter(db sdk.StateDB, addr common.Address) *big.Int {
	counter := GetCounter(db, addr)
	counter = new(big.Int).Add(counter, common.Big1)
	db.SetState(addr, l2StateSenderCounterKey, counter.Bytes())
	return counter
}

func GetCounter(db sdk.StateDBReader, addr common.Address) *big.Int {
	value := db.GetState(addr, l2StateSenderCounterKey)
	if len(value) == 0 {
		return big.NewInt(0)
	}
	return new(big.Int).SetBytes(value)
}
