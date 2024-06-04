package test

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
)

func NewMemoryStateDB() *state.StateDB {
	db := rawdb.NewMemoryDatabase()
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db), nil)
	return statedb
}
