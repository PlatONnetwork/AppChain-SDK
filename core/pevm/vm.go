package pevm

import (
	"bytes"

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
	flags           FinishExecFlags
}

type Vm struct {
	env                     *Env
	cApp                    sdk.ContractsApp
	statedb                 sdk.StateDB
	mvMemory                *MvMemory
	txs                     coretypes.Transactions
	beneficiaryLocationHash MemoryLocationHash
}

func NewVm(
	env *Env,
	cApp sdk.ContractsApp,
	statedb sdk.StateDB,
	mvMemory *MvMemory,
	txs coretypes.Transactions) *Vm {
	return &Vm{
		env:                     env,
		cApp:                    cApp,
		statedb:                 statedb,
		mvMemory:                mvMemory,
		txs:                     txs,
		beneficiaryLocationHash: BasicLoc(env.Header.Coinbase),
	}
}

func (vm *Vm) Execute(txVersion *TxVersion) (*VmExecutionResult, error) {
	var (
		tx       = vm.txs[txVersion.TxIdx]
		fromHash = BasicLoc(tx.FromAddr(coretypes.NewEIP155Signer(vm.env.ChainConfig.ChainID)))
		toHash   MemoryLocationHash
	)
	if tx.To() != nil {
		toHash = BasicLoc(*tx.To())
	}

	db := NewVmDB(vm, txVersion.TxIdx, tx, fromHash, toHash)
	chainCtx := vm.env.ChainContext
	chainCfg := vm.env.ChainConfig
	vmCfg := vm.env.VMConfig
	header := vm.env.Header
	gasPool := new(core.GasPool).AddGas(header.GasLimit)
	usedGas := new(uint64)

	receipt, err := core.ApplyTransaction(chainCfg, chainCtx, gasPool, db, header, tx, usedGas, vmCfg, vm.cApp)
	if err != nil {
		if db.txIdx > 0 {
			return nil, NewErrBlocking(db.txIdx, err)
		}
		return nil, NewErrExecution(db.txIdx, err)
	}

	writeSet := NewWriteSet(3)
	for addr, _ := range db.dirties {
		locationHash := BasicLoc(addr)
		account := db.readAccounts[locationHash]
		if account.Suicided {
			writeSet.Add(CodeHashLoc(addr), NewSelfDestructed(addr))
			continue
		}

		writeSet.Add(locationHash, NewBasic(addr, account))
		if account.NewCode {
			writeSet.Add(CodeHashLoc(addr), NewCodeHash(addr, account.CodeHash))
			db.vm.mvMemory.newByteCodes.Set(account.CodeHash, account.Code)
		}
	}

	for addr, states := range db.states {
		for key, val := range states {
			cloneKey := bytes.Clone([]byte(key))
			cloneVal := bytes.Clone(val)
			writeSet.Add(StateLoc(addr, cloneKey), NewState(addr, cloneKey, cloneVal))
		}
	}

	//if db.isLazy {
	//		db.vm.mvMemory.AddLazyAddresses([]common.Address{tx.FromAddr(coretypes.NewEIP155Signer(db.vm.ctx.ChainConfig().ChainID)), *tx.To()})
	//}

	var flags FinishExecFlags
	if db.txIdx > 0 {
		flags.Set(NeedValidation)
	}

	if db.vm.mvMemory.Record(txVersion, db.readSet, writeSet) {
		flags.Set(WroteNewLocation)
	}

	return &VmExecutionResult{
		executionResult: &PevmTxExecutionResult{
			receipt: receipt,
			gasUsed: *usedGas,
		},
		flags: flags,
	}, nil
}
