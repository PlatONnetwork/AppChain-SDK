package db

import (
	utils "github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type Case struct {
	Value  interface{}
	Expect []byte
}

func TestEncode(t *testing.T) {
	type ID struct {
		Name string
		Num  uint64
	}
	mustrlp := func(i interface{}) []byte {
		v, err := rlp.EncodeToBytes(i)
		require.Nil(t, err)
		return v
	}
	testcases := []Case{
		{
			big.NewInt(100),
			math.PaddedBigBytes(big.NewInt(100), 32),
		},
		{
			*big.NewInt(100),
			math.PaddedBigBytes(big.NewInt(100), 32),
		},
		{
			uint64(32),
			utils.EncodeUint64ToBytes(32),
		},
		{
			uint32(32),
			utils.EncodeUint32ToBytes(32),
		},
		{
			uint16(32),
			utils.EncodeUint16ToBytes(32),
		},
		{
			uint8(32),
			[]byte{byte(32)},
		},
		{
			true,
			[]byte{1},
		},
		{
			"hello",
			[]byte("hello"),
		},
		{
			&ID{Name: "he", Num: 1},
			mustrlp(&ID{Name: "he", Num: 1}),
		},
	}
	mustencode := func(key any) []byte {
		v, err := encodeKey(key)
		require.Nil(t, err)
		return v
	}
	for _, c := range testcases {
		require.Equal(t, c.Expect, mustencode(c.Value))
	}
}
