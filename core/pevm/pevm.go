package pevm

import (
	"errors"
	"fmt"
	"math/big"
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
	if result.receipt == nil {
		panic(fmt.Sprintf("empty receipt(txIdx: %d)", index))
	}
	e.results[index] = result
}

func (e *ExecutionResults) Range(f func(int, *ExecutionResult)) {
	e.mu.Lock()
	defer e.mu.Unlock()
	for i, result := range e.results {
		f(i, result)
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
	signer coretypes.Signer

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
		signer:           coretypes.NewLondonSigner(env.ChainConfig.ChainID),

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
		committedTxs = make(coretypes.Transactions, 0)
		receipts     = make(coretypes.Receipts, 0)
		timeout      bool
		begin        = time.Now()
		peeker       = miner.NewAppTxsPeeker(txs, e.signer)
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
		from, _ := coretypes.Sender(e.signer, tx)
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
			executionResults *ExecutionResults
			err              error
		)
		if isSysTxs || len(txs) <= batch {
			executionResults, err = e.parallelExecuteBatch(txs, isSysTxs)
			if err != nil {
				return &pevmResult, err
			}
			executionResults.Range(func(_ int, result *ExecutionResult) {
				receipt := result.receipt
				e.cumulativeGasUsed += receipt.GasUsed
				receipt.CumulativeGasUsed = e.cumulativeGasUsed
				receipt.TransactionIndex += uint(e.txCount)
				pevmResult.Receipts = append(pevmResult.Receipts, receipt)
			})
			pevmResult.Transactions = append(pevmResult.Transactions, txs...)
			pevmResult.GasUsed = e.cumulativeGasUsed
			e.txCount += len(txs)
		} else {
			peeker := NewTxsPeeker(txs, e.signer)
			execTxs := peeker.Peeks(batch)
			for len(execTxs) > 0 {
				executionResults, err = e.parallelExecuteBatch(execTxs, isSysTxs)
				if err != nil {
					return &pevmResult, err
				}
				executionResults.Range(func(i int, result *ExecutionResult) {
					if result == nil || result.receipt == nil {
						panic(fmt.Sprintf("empty result(index: %d, txCount: %d, count: %d)",
							i, e.txCount, len(execTxs)))
					}
					receipt := result.receipt
					e.cumulativeGasUsed += receipt.GasUsed
					receipt.CumulativeGasUsed = e.cumulativeGasUsed
					receipt.TransactionIndex += uint(e.txCount)
					pevmResult.Receipts = append(pevmResult.Receipts, receipt)
				})
				pevmResult.Transactions = append(pevmResult.Transactions, execTxs...)
				pevmResult.GasUsed = e.cumulativeGasUsed
				e.txCount += len(execTxs)

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
				execTxs = peeker.Peeks(batch)
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
		executionResults.Range(func(_ int, result *ExecutionResult) {
			receipt := result.receipt
			cumulativeGasUsed += receipt.GasUsed
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
	defer mvMemory.Release()
	vm := NewVm(e.env, e.cApp, e.signer, NewStateDBMut(e.env.StateDB), mvMemory, txs)

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

	scheduler.CheckStatus()

	statedb := e.env.StateDB
	committedLocations := make(map[MemoryLocationHash]bool)
	var abortErr error
	mvMemory.ConsumeLazyAddresses(func(addr common.Address) bool {
		locationHash := BasicLoc(addr)
		committedLocations[locationHash] = true
		if writeHistory := mvMemory.data.Get(locationHash); writeHistory != nil {
			var (
				balance = new(big.Int)
				nonce   uint64
			)
			writeHistory.Ascend(func(entry *item) bool {
				de := entry.Entry.(*DataEntry)
				if _, ok := de.Value.(*Basic); !ok {
					balance.Set(statedb.GetBalance(addr))
					nonce = statedb.GetNonce(addr)
				}
				return false
			})

			writeHistory.Ascend(func(entry *item) bool {
				tx := txs[entry.TxIdx]

				if entry, ok := entry.Entry.(*DataEntry); ok {
					if basic, vok := entry.Value.(*Basic); vok {
						account := basic.Account
						balance.Set(account.Balance)
						nonce = account.Nonce
					} else if lazy, vok := entry.Value.(*LazyRecipient); vok {
						balance.Add(balance, lazy.Balance)
					} else if lazy, vok := entry.Value.(*LazySender); vok {
						maxFee := tx.Gas()*tx.GasPrice().Uint64() + tx.Value().Uint64()
						if balance.Uint64() < maxFee {
							abortErr = fmt.Errorf("lack of fund for max fee(balance: %d, maxFee: %d)",
								balance.Uint64(), maxFee)
							return false
						}
						balance.Sub(balance, lazy.Balance)
						nonce += 1
					}
				}

				if tx.FromAddr(e.signer) == addr {
					var executeNonce uint64
					if nonce == 0 {
						abortErr = fmt.Errorf("unreachable error(addr: %s)", addr.Hex())
						return false
					}
					executeNonce = nonce - 1
					if executeNonce != tx.Nonce() {
						abortErr = fmt.Errorf("nonce mismatch(tx: %s, txIdx: %d, from: %s, nonce: %d, execute_nonce: %d)",
							tx.Hash().TerminalString(), entry.TxIdx, addr.Hex(), tx.Nonce(), executeNonce)
						return false
					}
				}

				statedb.SetBalance(addr, balance)
				if nonce > 0 {
					statedb.SetNonce(addr, nonce)
				}
				// End writeHistory.Ascend
				return true
			})
			if abortErr != nil {
				return false
			}
		}
		// End mvMomeory.ConsumeLazyAddresses()
		return true
	})
	if abortErr != nil {
		return nil, abortErr
	}

	for _, shard := range mvMemory.data.shards {
		for locationHash, writeHistory := range shard.histories {
			if committedLocations[locationHash] {
				continue
			}
			writeHistory.Ascend(func(d *item) bool {
				if entry, ok := d.Entry.(*DataEntry); ok {
					if basic, vok := entry.Value.(*Basic); vok {
						account := basic.Account
						if !account.Suicided {
							if account.Nonce > 0 {
								statedb.SetNonce(account.Addr, account.Nonce)
							}
							statedb.SetBalance(account.Addr, account.Balance)
						}
					} else if des, vok := entry.Value.(*SelfDestructed); vok {
						statedb.Suicide(des.Addr)
					} else if state, vok := entry.Value.(*State); vok {
						statedb.SetState(state.Addr, state.Key, state.Value)
					} else if codeHash, vok := entry.Value.(*CodeHash); vok {
						if code, ok := mvMemory.newByteCodes.Get(codeHash.CodeHash); ok {
							statedb.SetCode(codeHash.Addr, code)
						}
					}
				}
				return true
			})
		}
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
