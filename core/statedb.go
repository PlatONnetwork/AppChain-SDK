package core

import (
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func CloneStateDB(statedb sdk.StateDB) sdk.StateDB {
	if s, ok := statedb.(*state.StateDB); ok {
		return s.NewStateDB()
	}
	return nil
}
