package db

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type S struct {
	Name string
}

func TestDB(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	require.Nil(t, err)
	store := NewStore([]byte{}, common.Address{}, statedb)
	key := []byte("hh")
	SetState(store, key, "hello,world")
	name, err := GetState[string](store, key)
	require.Nil(t, err)
	t.Log(name)
	SetState(store, key, &S{Name: "hello"})
	n, err := GetState[*S](store, key)
	t.Log(n.Name)
	type Info struct {
		Name  string
		Num   uint64
		Value *big.Int
	}
	require.Nil(t, SetState(store, key, &Info{Name: "alice", Num: 10, Value: big.NewInt(100)}))
	info, err := GetState[*Info](store, key)
	require.Equal(t, big.NewInt(100), info.Value)
	require.Equal(t, "alice", info.Name)
	require.Equal(t, uint64(10), info.Num)
}

func TestDBKey(t *testing.T) {
	statedb, err := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	require.Nil(t, err)
	store := NewStore([]byte{}, common.Address{}, statedb)
	require.Nil(t, SetState(store, nil, nil))
}
