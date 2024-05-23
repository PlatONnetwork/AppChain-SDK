package test

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

func NewEVM(statedb *state.StateDB, vmconfig vm.Config, txContext vm.TxContext, blockContext vm.BlockContext, app sdk.ContractsApp) *vm.EVM {
	return vm.NewEVM(blockContext, txContext, statedb, params.MainnetChainConfig, vmconfig, app)
}

func NewBlockContext() vm.BlockContext {
	return vm.BlockContext{
		CanTransfer: func(db vm.StateDB, address common.Address, b *big.Int) bool {
			return core.CanTransfer(db, address, b)

		},
		Transfer: func(db vm.StateDB, address common.Address, address2 common.Address, b *big.Int) {
			core.Transfer(db, address, address2, b)
		},
		GetHash:     vmTestBlockHash,
		BlockNumber: new(big.Int).SetUint64(10),
		Time:        new(big.Int).SetUint64(1),
		GasLimit:    100000000,
	}
}

func vmTestBlockHash(n uint64) common.Hash {
	return common.BytesToHash(crypto.Keccak256([]byte(big.NewInt(int64(n)).String())))
}
