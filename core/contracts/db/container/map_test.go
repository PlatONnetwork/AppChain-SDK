package container

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMap(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	require.Nil(t, err)
	n := NewMap[string]([]byte("abc"), common.Address{}, statedb)
	value := map[string]string{"a": "a", "b": "b", "c": "c"}
	for k, v := range value {
		n.Set(k, v)
	}
	for k, v := range value {
		actual, err := n.Get(k)
		require.Nil(t, err)
		require.Equal(t, v, actual)
	}
	actual, err := n.Get("xx")
	require.Nil(t, err)
	require.Nil(t, "", actual)
}

func TestMapContainer(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	require.Nil(t, err)
	prefix := []byte("abc")
	n := NewMap[*Array[string]](prefix, common.Address{}, statedb)

	value := []string{"a", "b", "c"}
	for i, _ := range value {
		a := NewArray[string](n.CreatePrefix(uint32(i)), common.Address{}, statedb)
		for _, vv := range value {
			a.Push(vv)
		}
		if uint32(len(value)) != a.Length() {
			require.Equal(t, uint32(len(value)), a.Length())
		}

		n.Set(uint32(i), a)
	}
	n2 := NewMap[*Array[string]](prefix, common.Address{}, statedb)

	for i, _ := range value {
		a, err := n2.Get(uint32(i))
		require.Nil(t, err)
		require.Equal(t, uint32(len(value)), a.Length())
		for i, _ := range value {
			v, err := a.Index(uint32(i))
			require.Nil(t, err)
			require.Equal(t, value[i], v)
		}
	}
}
