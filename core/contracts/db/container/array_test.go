package container

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestArray(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	require.Nil(t, err)
	prefix := []byte("abc")
	n := NewArray[string](prefix, common.Address{}, statedb)
	require.Equal(t, uint32(0), n.Length())
	value := []string{"a", "b", "c"}
	for _, v := range value {
		n.Push(v)
	}
	n2 := NewArray[string](prefix, common.Address{}, statedb)
	require.Equal(t, uint32(3), n.Length())
	for i := uint32(0); i < n.Length(); i++ {
		v, err := n2.Index(i)
		require.Nil(t, err)
		require.Equal(t, value[i], v)
	}
}

func TestArrayContainer(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	require.Nil(t, err)
	prefix := []byte("abc")
	n := NewArray[*Map[string]](prefix, common.Address{}, statedb)
	length := uint32(3)
	for i := uint32(0); i < length; i++ {
		elem := n.MustIndex(n.Length())
		elem.MustSet("abc", fmt.Sprintf("%d", i))
		n.Push(elem)
	}

	require.Equal(t, length, n.Length())

	for i := uint32(0); i < length; i++ {
		x2, err := n.Index(i)
		require.Nil(t, err)
		actual, err := x2.Get("abc")
		require.Nil(t, err)
		require.Equal(t, fmt.Sprintf("%d", i), actual)
	}
	elem, err := n.Index(n.Length())
	require.Nil(t, err)
	require.NotNil(t, elem)
}
