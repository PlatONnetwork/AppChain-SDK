package statesender

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initAccountNonce(statedb sdk.StateDB, addr common.Address) {
	statedb.SetNonce(addr, 1)
}
