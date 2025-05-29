package pevm

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/x/pevm/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/miner"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/google/btree"
	"golang.org/x/sync/errgroup"
)

const ParallelExecuteTxnBatch = 64

type AbortReason struct {
	mu     sync.RWMutex
	reason error
}

func (r *AbortReason) Get() error {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.reason
}

func (r *AbortReason) Set(err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reason = err
}

func (r *AbortReason) InsertIfNotSet(err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.reason == nil {
		r.reason = err
	}
}

type PEVMResult struct {
	Transactions coretypes.Transactions
	Receipts     coretypes.Receipts
	GasUsed      uint64
	Timeout      bool
}

type PEVM struct {
	logger log.Logger
	ctx    sdk.WorkerContext
	cApp   sdk.ContractsApp

	// Use in serial execution
	gp      *core.GasPool
	usedGas uint64
	txCount int

	// Use in paralle execution
	mu               sync.Mutex
	executionResults []*PEVMResult
	abortReason      AbortReason
}

func NewPEVM(
	logger log.Logger,
	ctx sdk.WorkerContext,
	cApp sdk.ContractsApp) *PEVM {
	return &PEVM{
		logger: logger,
		ctx:    ctx,
		cApp:   cApp,

		gp: new(core.GasPool).AddGas(ctx.Header().GasLimit),
	}
}

func (e *PEVM) Run(txs coretypes.Transactions, isSysTxs bool) (*PEVMResult, error) {
	if len(txs) < runtime.NumCPU() {
		return e.serialExecute(txs, isSysTxs)
	}
	return e.parallelExecute(txs, isSysTxs)
}

func (e *PEVM) serialExecute(txs coretypes.Transactions, isSysTxs bool) (*PEVMResult, error) {
	var (
		result *PEVMResult
		err    error
	)
	if e.ctx.IsWorker() {
		result = e.commitTransactions(txs, isSysTxs)
	} else {
		result, err = e.applyTransactions(txs)
	}
	return result, err
}

func (e *PEVM) applyTransaction(tx *coretypes.Transaction) (*coretypes.Receipt, error) {
	var (
		chainCtx = e.ctx.Backend().ChainContext()
		chainCfg = e.ctx.ChainConfig()
		vmCfg    = *e.ctx.VMConfig()
		header   = e.ctx.Header()
		statedb  = e.ctx.StateDB()
	)
	receipt, err := core.ApplyTransaction(chainCfg, chainCtx, e.gp, statedb, header, tx, &e.usedGas, vmCfg, e.cApp)
	if err != nil {
		e.logger.Error("Failed to apply transaction",
			"blockNumber", header.Number,
			"parentHash", header.ParentHash,
			"txHash", tx.Hash(),
			"err", err)
	}
	return receipt, err
}

func (e *PEVM) commitTransactions(txs coretypes.Transactions, isSysTxs bool) *PEVMResult {
	var (
		bd           = e.ctx.BlockDeadline()
		header       = e.ctx.Header()
		blockNumber  = header.Number
		parentHash   = header.ParentHash
		timestamp    = int64(header.Time)
		statedb      = e.ctx.StateDB()
		signer       = coretypes.MakeSigner(e.ctx.ChainConfig(), header.Number)
		committedTxs = make(coretypes.Transactions, 0)
		receipts     = make(coretypes.Receipts, 0)
		timeout      bool
		begin        = time.Now()
		peeker       = miner.NewAppTxsPeeker(txs, signer)
	)

	for {
		now := time.Now()
		if bd.Before(now) && !isSysTxs {
			e.logger.Warn("interrupt current ex-executing",
				"blockNumber", blockNumber,
				"parentHash", parentHash,
				"now", now.UnixMilli(),
				"timestamp", timestamp,
				"deadlineDuration", bd.Sub(time.UnixMilli(timestamp)))
			timeout = true
			break
		}

		if e.gp.Gas() < params.TxGas {
			e.logger.Trace("Not enough gas for further transactions",
				"blockNumber", blockNumber,
				"parentHash", parentHash,
				"have", e.gp,
				"want", params.TxGas)
			break
		}

		tx := peeker.Peek()
		if tx == nil {
			break
		}
		from, _ := coretypes.Sender(signer, tx)
		statedb.Prepare(tx.Hash(), e.txCount)

		receipt, err := e.applyTransaction(tx)
		switch err {
		case core.ErrGasLimitReached:
			e.logger.Warn("Gas limit exceeded for current block",
				"blockNumber", blockNumber,
				"parentHash", parentHash,
				"parentHash", header.ParentHash,
				"txHash", tx.Hash(),
				"sender", from)
			peeker.Pop()

		case core.ErrNonceTooLow:
			peeker.Shift()

		case core.ErrNonceTooHigh:
			e.logger.Warn("Skipping account with high nonce",
				"blockNumber", blockNumber,
				"parentHash", parentHash,
				"txHash", tx.Hash(),
				"sender", from,
				"senderCurNonce", statedb.GetNonce(from),
				"tx.nonce", tx.Nonce())
			peeker.Pop()

		case vm.ErrAbort:
			e.logger.Warn("Skipping account with exec timeout tx",
				"blockNumber", blockNumber,
				"parentHash", parentHash,
				"txHash", tx.Hash(),
				"err", err)
			peeker.Shift()

		case nil:
			e.txCount++
			committedTxs = append(committedTxs, tx)
			receipts = append(receipts, receipt)
			peeker.Shift()
		}
	}
	e.logger.Info("Commit transactions finished",
		"txs", len(txs),
		"committedTxs", len(committedTxs),
		"timeout", timeout,
		"elapsed", time.Since(begin))
	return &PEVMResult{
		Transactions: committedTxs,
		Receipts:     receipts,
		Timeout:      timeout,
		GasUsed:      e.usedGas,
	}
}

func (e *PEVM) applyTransactions(txs coretypes.Transactions) (*PEVMResult, error) {
	var (
		receipts = make(coretypes.Receipts, 0)
		header   = e.ctx.Header()
		statedb  = e.ctx.StateDB()
		begin    = time.Now()
	)

	for i, tx := range txs {
		statedb.Prepare(tx.Hash(), i)
		receipt, err := e.applyTransaction(tx)
		if err != nil {
			return nil, fmt.Errorf("could not apply tx %d [%v]: %w", i, tx.Hash(), err)
		}
		receipts = append(receipts, receipt)
	}
	e.logger.Info("apply transactions finished",
		"blockNumber", header.Number,
		"blockHash", header.Hash(),
		"txs", len(txs),
		"elapsed", time.Since(begin))
	return &PEVMResult{Receipts: receipts, GasUsed: e.usedGas}, nil
}

func (e *PEVM) parallelExecute(txs coretypes.Transactions, isSysTxs bool) (*PEVMResult, error) {
	if len(txs) == 0 {
		return &PEVMResult{}, nil
	}

	// TODO: batch parallel in minging

	blockSize := uint32(len(txs))
	e.executionResults = make([]*PEVMResult, blockSize)
	scheduler := NewScheduler(blockSize)
	txIdxs := make([]uint32, blockSize)
	i := uint32(0)
	for ; i < blockSize; i++ {
		txIdxs[i] = i
	}
	mvMemory := NewMvMemory(int(blockSize), map[types.MemoryLocationHash][]uint32{
		types.BasicLoc(e.ctx.Header().Coinbase): txIdxs,
	}, []common.Address{e.ctx.Header().Coinbase})
	vm := NewVm(e.ctx, e.cApp, e.ctx.StateDB(), mvMemory, txs)

	g := new(errgroup.Group)
	for j := 0; j < runtime.NumCPU(); j++ {
		g.Go(func() error {
			task := scheduler.NextTask()
			for task != nil {
				switch task.(type) {
				case *types.Execution:
					task = e.tryExecute(vm, scheduler, task.Version())
				case *types.Validation:
					task = e.tryValidate(mvMemory, scheduler, task.Version())
				}

				if e.abortReason.Get() != nil {
					break
				}

				if task == nil {
					task = scheduler.NextTask()
				}
			}
			return nil
		})
	}
	g.Wait()

	if reason := e.abortReason.Get(); reason != nil {
		return nil, reason
	}

	var (
		cumulativeGasUsed uint64
		receipts          coretypes.Receipts
	)
	for i := 0; i < int(blockSize); i++ {
		receipt := e.executionResults[i].Receipts[0]
		cumulativeGasUsed += receipt.CumulativeGasUsed
		receipt.CumulativeGasUsed = cumulativeGasUsed
		receipt.TransactionIndex = uint(e.txCount) + receipt.TransactionIndex
		receipts = append(receipts, e.executionResults[i].Receipts...)
	}

	statedb := e.ctx.StateDB()
	for _, writeHistory := range mvMemory.data.Items() {
		writeHistory.Ascend(func(item btree.Item) bool {
			d := item.(*DataEntry)
			switch d.Entry.(type) {
			case *types.DataEntry:
				entry := d.Entry.(*types.DataEntry)
				switch entry.Value.(type) {
				case *types.Basic:
					basic := entry.Value.(*types.Basic)
					account := basic.Account
					if !account.Suicided {
						statedb.SetNonce(account.Addr, account.Nonce)
						statedb.SetBalance(account.Addr, account.Balance)
					}
				case *types.SelfDestructed:
					des := entry.Value.(*types.SelfDestructed)
					statedb.Suicide(des.Addr)
				case *types.State:
					state := entry.Value.(*types.State)
					statedb.SetState(state.Addr, state.Key, state.Value)
				case *types.CodeHash:
					codeHash := entry.Value.(*types.CodeHash)
					if code, ok := mvMemory.newByteCodes.Get(codeHash.CodeHash); ok {
						statedb.SetCode(codeHash.Addr, code)
					}
				}
			}
			return true
		})
	}

	e.txCount = len(txs)

	return &PEVMResult{
		Transactions: txs,
		Receipts:     receipts,
		GasUsed:      cumulativeGasUsed,
	}, nil
}

func (e *PEVM) tryExecute(vm *Vm, scheduler *Scheduler, txVersion types.TxVersion) types.Task {
	for {
		result, err := vm.Execute(&txVersion)
		if err != nil {
			if errors.Is(err, &ErrBlocking{}) {
				if !scheduler.AddDependency(txVersion.TxIdx, err.(ErrBlocking).TxIdx) &&
					e.abortReason.Get() == nil {
					continue
				}
			}
			scheduler.Abort()
			e.abortReason.InsertIfNotSet(err)
			return nil
		}
		e.mu.Lock()
		e.executionResults[txVersion.TxIdx] = &PEVMResult{
			Receipts: coretypes.Receipts{result.executionResult.receipt},
			GasUsed:  result.executionResult.gasUsed,
			Timeout:  false,
		}
		e.mu.Unlock()
		return scheduler.FinishExecution(txVersion, result.flags)
	}

}

func (e *PEVM) tryValidate(mvMemory *MvMemory, scheduler *Scheduler, txVersion types.TxVersion) types.Task {
	readSetValid := mvMemory.ValidateReadLocations(txVersion.TxIdx)
	aborted := !readSetValid && scheduler.TryValidationAbort(txVersion)
	if aborted {
		mvMemory.ConvertWritesToEstimate(txVersion.TxIdx)
	}
	return scheduler.FinishValidation(txVersion, aborted)
}
