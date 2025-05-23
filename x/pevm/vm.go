package pevm

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/pevm/types"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type VM struct {
	ctx                     sdk.WorkerContext
	statedb                 sdk.StateDB
	mvMemory                *MvMemory
	txs                     coretypes.Transactions
	beneficiaryLocationHash types.MemoryLocationHash
}

func NewVM(
	ctx sdk.WorkerContext,
	statedb sdk.StateDB,
	mvMemory *MvMemory,
	txs coretypes.Transactions) *VM {
	return &VM{
		ctx:                     ctx,
		statedb:                 statedb,
		mvMemory:                mvMemory,
		txs:                     txs,
		beneficiaryLocationHash: types.NewBasicLoc(ctx.Header().Coinbase).Hash(),
	}
}

func (vm *VM) Execute(txVersion *types.TxVersion) error {
	return nil
}
