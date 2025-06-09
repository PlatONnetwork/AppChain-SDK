package pevm

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	coresdk "github.com/PlatONnetwork/PlatON-Go/core/sdk"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/miner"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/google/btree"
	"golang.org/x/sync/errgroup"
)

var AbortFallback = AbortFallbackToSequential{}

type AbortFallbackToSequential struct{}

func (e AbortFallbackToSequential) Error() string {
	return "abort to fallback to sequential"
}

func (e AbortFallbackToSequential) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

type AbortExecutionError struct {
	err error
}

func (e AbortExecutionError) Error() string {
	return fmt.Sprintf("Abort: %v", e.err)
}

func (e AbortExecutionError) Is(rhl error) bool {
	return reflect.TypeOf(e).Name() == reflect.TypeOf(rhl).Name()
}

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

type ExecutionResult struct {
	tx      *coretypes.Transaction
	receipt *coretypes.Receipt
	gasUsed uint64
}

type ExecutionResults struct {
	mu      sync.Mutex
	results []*ExecutionResult
}

func NewExecutionResults(blockSize int32) *ExecutionResults {
	return &ExecutionResults{
		results: make([]*ExecutionResult, blockSize),
	}
}

func (e *ExecutionResults) Set(index int32, result *ExecutionResult) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.results[index] = result
}

func (e *ExecutionResults) Range(f func(*ExecutionResult)) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for _, result := range e.results {
		f(result)
	}
}

type PEVMResult struct {
	Transactions coretypes.Transactions
	Receipts     coretypes.Receipts
	GasUsed      uint64
	Timeout      bool
}

type Env struct {
	Header        *coretypes.Header
	StateDB       sdk.StateDB
	ChainConfig   *params.ChainConfig
	VMConfig      vm.Config
	ChainContext  coresdk.ChainContext
	IsWorker      bool
	BlockDeadline time.Time
}

type PEVM struct {
	forceSequential  bool
	concurrencyLevel int
	txsBatch         int

	logger log.Logger
	env    *Env
	cApp   sdk.ContractsApp

	// Use in serial execution
	gp *core.GasPool

	cumulativeGasUsed uint64
	txCount           int

	// Use in paralle execution
	abortReason AbortReason
}

func NewPEVM(
	forceSequential bool,
	concurrencyLevel int,
	txsBatch int,
	logger log.Logger,
	env *Env,
	cApp sdk.ContractsApp) *PEVM {
	return &PEVM{
		forceSequential:  forceSequential,
		concurrencyLevel: concurrencyLevel,
		txsBatch:         txsBatch,
		logger:           logger,
		env:              env,
		cApp:             cApp,

		gp: new(core.GasPool).AddGas(env.Header.GasLimit),
	}
}

func (e *PEVM) Run(txs coretypes.Transactions, isSysTxs bool) (*PEVMResult, error) {
	if e.forceSequential || len(txs) < e.concurrencyLevel {
		return e.serialExecute(txs, isSysTxs)
	}
	return e.parallelExecute(txs, isSysTxs)
}

func (e *PEVM) serialExecute(txs coretypes.Transactions, isSysTxs bool) (*PEVMResult, error) {
	var (
		result *PEVMResult
		err    error
	)
	if e.env.IsWorker {
		result = e.commitTransactions(txs, isSysTxs)
	} else {
		result, err = e.applyTransactions(txs)
	}
	return result, err
}

func (e *PEVM) applyTransaction(tx *coretypes.Transaction) (*coretypes.Receipt, error) {
	var (
		chainCtx = e.env.ChainContext
		chainCfg = e.env.ChainConfig
		vmCfg    = e.env.VMConfig
		header   = e.env.Header
		statedb  = e.env.StateDB
	)
	receipt, err := core.ApplyTransaction(chainCfg, chainCtx, e.gp, statedb, header, tx, &e.cumulativeGasUsed, vmCfg, e.cApp)
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
		bd           = e.env.BlockDeadline
		header       = e.env.Header
		blockNumber  = header.Number
		parentHash   = header.ParentHash
		timestamp    = int64(header.Time)
		statedb      = e.env.StateDB
		signer       = coretypes.MakeSigner(e.env.ChainConfig, header.Number)
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
		GasUsed:      e.cumulativeGasUsed,
	}
}

func (e *PEVM) applyTransactions(txs coretypes.Transactions) (*PEVMResult, error) {
	var (
		receipts = make(coretypes.Receipts, 0)
		header   = e.env.Header
		statedb  = e.env.StateDB
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
	return &PEVMResult{Receipts: receipts, GasUsed: e.cumulativeGasUsed}, nil
}

func (e *PEVM) parallelExecute(txs coretypes.Transactions, isSysTxs bool) (*PEVMResult, error) {
	if len(txs) == 0 {
		return &PEVMResult{}, nil
	}

	begin := time.Now()

	var pevmResult PEVMResult
	if e.env.IsWorker {
		var (
			timestamp        int64 = int64(e.env.Header.Time)
			blockDeadline          = e.env.BlockDeadline
			batch                  = e.txsBatch
			startIndex       int
			endIndex         int
			executionResults *ExecutionResults
			err              error
		)
		if isSysTxs || len(txs) <= batch {
			executionResults, err = e.parallelExecuteBatch(txs, isSysTxs)
			if err != nil {
				return &pevmResult, err
			}
			executionResults.Range(func(result *ExecutionResult) {
				receipt := result.receipt
				e.cumulativeGasUsed += receipt.CumulativeGasUsed
				receipt.CumulativeGasUsed = e.cumulativeGasUsed
				receipt.TransactionIndex += uint(e.txCount)
				pevmResult.Receipts = append(pevmResult.Receipts, receipt)
			})
			pevmResult.Transactions = append(pevmResult.Transactions, txs...)
			pevmResult.GasUsed = e.cumulativeGasUsed
			e.txCount += len(txs)
		} else {
			count := len(txs)
			endIndex = startIndex + batch
			for endIndex <= count {
				execTxs := txs[startIndex:endIndex]
				executionResults, err = e.parallelExecuteBatch(execTxs, isSysTxs)
				if err != nil {
					return &pevmResult, err
				}
				executionResults.Range(func(result *ExecutionResult) {
					receipt := result.receipt
					e.cumulativeGasUsed += receipt.CumulativeGasUsed
					receipt.CumulativeGasUsed = e.cumulativeGasUsed
					receipt.TransactionIndex += uint(e.txCount)
					pevmResult.Receipts = append(pevmResult.Receipts, receipt)
				})
				pevmResult.Transactions = append(pevmResult.Transactions, execTxs...)
				pevmResult.GasUsed = e.cumulativeGasUsed
				e.txCount = len(execTxs)

				now := time.Now()
				if blockDeadline.Before(time.Now()) {
					e.logger.Warn("interrupt current ex-executing",
						"blockNumber", e.env.Header.Number,
						"parentHash", e.env.Header.ParentHash,
						"now", now.UnixMilli(),
						"timestamp", timestamp,
						"deadlineDuration", blockDeadline.Sub(time.UnixMilli(timestamp)))
					pevmResult.Timeout = true
					break
				}

				startIndex = endIndex
				if startIndex >= count-1 {
					break
				}
				endIndex = startIndex + batch
				if endIndex > count {
					endIndex = count
				}
			}
		}
	} else {
		executionResults, err := e.parallelExecuteBatch(txs, isSysTxs)
		if err != nil {
			e.logger.Error("parallel execute failed",
				"number", e.env.Header.Number,
				"hash", e.env.Header.Hash(),
				"err", err)
			return &pevmResult, err
		}
		var cumulativeGasUsed uint64
		executionResults.Range(func(result *ExecutionResult) {
			receipt := result.receipt
			cumulativeGasUsed += receipt.CumulativeGasUsed
			receipt.CumulativeGasUsed = cumulativeGasUsed
			pevmResult.Receipts = append(pevmResult.Receipts, receipt)
		})
		pevmResult.Transactions = append(pevmResult.Transactions, txs...)
		pevmResult.GasUsed = cumulativeGasUsed
		e.logger.Info("parallel execute success",
			"number", e.env.Header.Number,
			"hash", e.env.Header.Hash())
	}

	e.logger.Info("Parallel execute finish",
		"number", e.env.Header.Number,
		"parentHash", e.env.Header.ParentHash,
		"txs", len(txs),
		"committedTxs", len(pevmResult.Transactions),
		"timeout", pevmResult.Timeout,
		"elapsed", time.Since(begin))
	return &pevmResult, nil
}

func (e *PEVM) parallelExecuteBatch(txs coretypes.Transactions, isSysTxs bool) (*ExecutionResults, error) {
	if len(txs) == 0 {
		return &ExecutionResults{}, nil
	}

	blockSize := int32(len(txs))
	executionResults := NewExecutionResults(blockSize)
	scheduler := NewScheduler(blockSize)
	txIdxs := make([]int32, blockSize)
	i := int32(0)
	for ; i < blockSize; i++ {
		txIdxs[i] = i
	}
	mvMemory := NewMvMemory(int(blockSize), map[MemoryLocationHash][]int32{
		BasicLoc(e.env.Header.Coinbase): txIdxs,
	}, []common.Address{e.env.Header.Coinbase})
	vm := NewVm(e.env, e.cApp, e.env.StateDB, mvMemory, txs)

	g := new(errgroup.Group)
	for j := 0; j < e.concurrencyLevel; j++ {
		g.Go(func() error {
			task := scheduler.NextTask()
			for task != nil {
				switch task.(type) {
				case *Execution:
					task = e.tryExecute(vm, scheduler, executionResults, task.Version())
				case *Validation:
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
		// FIXME: all txs fallback to sequential
		if errors.Is(reason, AbortFallback) {
			pevmRes, err := e.serialExecute(txs, isSysTxs)
			if err != nil {
				return nil, err
			}
			for i, receipt := range pevmRes.Receipts {
				r := receipt
				executionResults.Set(int32(i), &ExecutionResult{
					receipt: r,
					gasUsed: r.GasUsed,
				})
			}
			return executionResults, nil
		}
		return nil, reason
	}

	statedb := e.env.StateDB
	for _, writeHistory := range mvMemory.data.Items() {
		writeHistory.Ascend(func(itm btree.Item) bool {
			d := itm.(*item)
			switch d.Entry.(type) {
			case *DataEntry:
				entry := d.Entry.(*DataEntry)
				switch entry.Value.(type) {
				case *Basic:
					basic := entry.Value.(*Basic)
					account := basic.Account
					if !account.Suicided {
						if account.Addr != e.env.Header.Coinbase && account.Nonce > 0 { // FIXME: coinbase maybe a sender
							statedb.SetNonce(account.Addr, account.Nonce)
						}
						statedb.SetBalance(account.Addr, account.Balance)
					}
				case *SelfDestructed:
					des := entry.Value.(*SelfDestructed)
					statedb.Suicide(des.Addr)
				case *State:
					state := entry.Value.(*State)
					statedb.SetState(state.Addr, state.Key, state.Value)
				case *CodeHash:
					codeHash := entry.Value.(*CodeHash)
					if code, ok := mvMemory.newByteCodes.Get(codeHash.CodeHash); ok {
						statedb.SetCode(codeHash.Addr, code)
					}
				}
			}
			return true
		})
	}
	return executionResults, nil
}

func (e *PEVM) tryExecute(vm *Vm, scheduler *Scheduler, executionResults *ExecutionResults, txVersion TxVersion) Task {
	for {
		result, err := vm.Execute(&txVersion)
		if err != nil {
			if errors.Is(err, ErrRetry) {
				if e.abortReason.Get() == nil {
					continue
				}
			} else if errors.Is(err, ErrFallbackToSequential) {
				scheduler.Abort()
				e.abortReason.InsertIfNotSet(AbortFallback)
			} else if errors.Is(err, ExecutionBlockingError{}) {
				if !scheduler.AddDependency(txVersion.TxIdx, err.(ExecutionBlockingError).TxIdx) &&
					e.abortReason.Get() == nil {
					continue
				}
			} else {
				scheduler.Abort()
				e.abortReason.InsertIfNotSet(AbortExecutionError{err})
			}
			return nil
		}
		executionResults.Set(txVersion.TxIdx, &ExecutionResult{
			receipt: result.executionResult.receipt,
			gasUsed: result.executionResult.gasUsed,
		})
		return scheduler.FinishExecution(txVersion, result.flags)
	}

}

func (e *PEVM) tryValidate(mvMemory *MvMemory, scheduler *Scheduler, txVersion TxVersion) Task {
	readSetValid := mvMemory.ValidateReadLocations(txVersion.TxIdx)
	aborted := !readSetValid && scheduler.TryValidationAbort(txVersion)
	if aborted {
		mvMemory.ConvertWritesToEstimate(txVersion.TxIdx)
	}
	return scheduler.FinishValidation(txVersion, aborted)
}
