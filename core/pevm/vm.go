package pevm

import (
	"bytes"
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	corevm "github.com/PlatONnetwork/PlatON-Go/core/vm"
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
	signer                  coretypes.Signer
	statedb                 sdk.StateDB
	mvMemory                *MvMemory
	txs                     coretypes.Transactions
	beneficiaryLocationHash MemoryLocationHash
}

func NewVm(
	env *Env,
	cApp sdk.ContractsApp,
	signer coretypes.Signer,
	statedb sdk.StateDB,
	mvMemory *MvMemory,
	txs coretypes.Transactions) *Vm {
	return &Vm{
		env:                     env,
		cApp:                    cApp,
		signer:                  signer,
		statedb:                 statedb,
		mvMemory:                mvMemory,
		txs:                     txs,
		beneficiaryLocationHash: BasicLoc(env.Header.Coinbase),
	}
}

func (vm *Vm) Execute(txVersion *TxVersion) (result *VmExecutionResult, err error) {
	defer func() {
		if catchErr := recover(); catchErr != nil {
			if realErr, ok := catchErr.(error); ok {
				err = ToVmExecutionError(realErr)
			} else {
				panic(catchErr)
			}
		}
	}()
	var (
		tx       = vm.txs[txVersion.TxIdx]
		fromAddr = tx.FromAddr(vm.signer)
		fromHash = BasicLoc(fromAddr)
		toHash   MemoryLocationHash
	)
	if tx.To() != nil {
		toHash = BasicLoc(*tx.To())
	}

	db := NewVmDB(vm, txVersion.TxIdx, tx, fromAddr, fromHash, toHash)
	chainCtx := vm.env.ChainContext
	chainCfg := vm.env.ChainConfig
	vmCfg := vm.env.VMConfig
	header := vm.env.Header
	gasPool := new(core.GasPool).AddGas(header.GasLimit)
	usedGas := new(uint64)

	receipt, err := core.ApplyTransaction(chainCfg, chainCtx, gasPool, db, header, tx, usedGas, vmCfg, vm.cApp)
	switch err {
	case nil:
		writeSet := NewWriteSet()
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

		for addr, balance := range db.addBalances {
			lh := BasicLoc(addr)
			writeSet.Add(lh, NewLazyRecipient(addr, new(big.Int).Set(balance)))
		}
		for addr, balance := range db.subBalances {
			lh := BasicLoc(addr)
			writeSet.Add(lh, NewLazySender(addr, new(big.Int).Set(balance)))
		}

		if db.isLazy {
			db.vm.mvMemory.AddLazyAddresses([]common.Address{db.fromAddr, *tx.To()})
		}

		var flags FinishExecFlags
		if db.txIdx > 0 && !db.isLazy {
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

	case core.ErrNonceTooLow, core.ErrNonceTooHigh:
		if db.txIdx > 0 {
			return nil, ExecutionBlockingError{TxIdx: db.txIdx - 1}
		} else {
			return nil, ExecutionError{err}
		}
	case core.ErrGasLimitReached, corevm.ErrAbort:
		return nil, ErrFallbackToSequential
	}
	return nil, ExecutionError{err}
}
