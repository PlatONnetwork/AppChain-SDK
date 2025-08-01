package core

import (
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func CloneStateDB(statedb sdk.StateDB) sdk.StateDB {
	if s, ok := statedb.(*state.StateDB); ok {
		return s.NewStateDB()
	}
	return nil
}

func CommitDB(statedb sdk.StateDB) (common.Hash, error) {
	if s, ok := statedb.(*state.StateDB); ok {
		return s.Commit(true)
	}
	return common.Hash{}, errors.New("invalid statedb")
}

func CleanStateDB(statedb sdk.StateDB) {
	if s, ok := statedb.(*state.StateDB); ok {
		s.ClearParentReference()
	}
}
