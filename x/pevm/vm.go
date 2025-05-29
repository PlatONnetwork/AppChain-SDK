package pevm

import (
	"bytes"

	"github.com/PlatONnetwork/AppChain-SDK/x/pevm/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type EvmStateTransactions = map[common.Address]map[string][]byte

type PevmTxExecutionResult struct {
	receipt *coretypes.Receipt
	gasUsed uint64
}

type VmExecutionResult struct {
	executionResult *PevmTxExecutionResult
	flags           types.FinishExecFlags
}

type Vm struct {
	ctx                     sdk.WorkerContext
	cApp                    sdk.ContractsApp
	statedb                 sdk.StateDB
	mvMemory                *MvMemory
	txs                     coretypes.Transactions
	beneficiaryLocationHash types.MemoryLocationHash
}

func NewVm(
	ctx sdk.WorkerContext,
	cApp sdk.ContractsApp,
	statedb sdk.StateDB,
	mvMemory *MvMemory,
	txs coretypes.Transactions) *Vm {
	return &Vm{
		ctx:                     ctx,
		cApp:                    cApp,
		statedb:                 statedb,
		mvMemory:                mvMemory,
		txs:                     txs,
		beneficiaryLocationHash: types.BasicLoc(ctx.Header().Coinbase),
	}
}

func (vm *Vm) Execute(txVersion *types.TxVersion) (*VmExecutionResult, error) {
	var (
		tx       = vm.txs[txVersion.TxIdx]
		fromHash = types.BasicLoc(tx.FromAddr(coretypes.NewEIP155Signer(vm.ctx.ChainConfig().ChainID)))
		toHash   types.MemoryLocationHash
	)
	if tx.To() != nil {
		toHash = types.BasicLoc(*tx.To())
	}

	db := NewVmDB(vm, txVersion.TxIdx, tx, fromHash, toHash)
	chainCtx := vm.ctx.Backend().ChainContext()
	chainCfg := vm.ctx.ChainConfig()
	vmCfg := *vm.ctx.VMConfig()
	header := vm.ctx.Header()
	gasPool := new(core.GasPool).AddGas(header.GasLimit)
	usedGas := new(uint64)

	receipt, err := core.ApplyTransaction(chainCfg, chainCtx, gasPool, db, header, tx, usedGas, vmCfg, vm.cApp)
	if err != nil {
		if db.txIdx > 0 {
			return nil, NewErrBlocking(db.txIdx, err)
		}
		return nil, NewErrExecution(db.txIdx, err)
	}

	writeSet := types.NewWriteSet(3)
	for addr, _ := range db.dirties {
		locationHash := types.BasicLoc(addr)
		account := db.readAccounts[locationHash]
		if account.Suicided {
			writeSet.Add(types.CodeHashLoc(addr), types.NewSelfDestructed(addr))
			continue
		}

		writeSet.Add(locationHash, types.NewBasic(addr, account))
		if account.NewCode {
			writeSet.Add(types.CodeHashLoc(addr), types.NewCodeHash(addr, account.CodeHash))
			db.vm.mvMemory.newByteCodes.Set(account.CodeHash, account.Code)
		}
	}

	for addr, states := range db.states {
		for key, val := range states {
			cloneKey := bytes.Clone([]byte(key))
			cloneVal := bytes.Clone(val)
			writeSet.Add(types.StateLoc(addr, cloneKey), types.NewState(addr, cloneKey, cloneVal))
		}
	}

	//if db.isLazy {
	//		db.vm.mvMemory.AddLazyAddresses([]common.Address{tx.FromAddr(coretypes.NewEIP155Signer(db.vm.ctx.ChainConfig().ChainID)), *tx.To()})
	//}

	var flags types.FinishExecFlags
	if db.txIdx > 0 {
		flags.Set(types.NeedValidation)
	}

	if db.vm.mvMemory.Record(txVersion, db.readSet, writeSet) {
		flags.Set(types.WroteNewLocation)
	}

	return &VmExecutionResult{
		executionResult: &PevmTxExecutionResult{
			receipt: receipt,
			gasUsed: *usedGas,
		},
		flags: flags,
	}, nil
}
