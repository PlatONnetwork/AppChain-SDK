package pevm

import (
	"fmt"
	"runtime"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/miner"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const ParallelExecuteTxnBatch = 64

type PEVM struct {
	logger log.Logger
	ctx    sdk.WorkerContext
	cApp   sdk.ContractsApp

	// Use in serial execution
	gp      *core.GasPool
	usedGas uint64
	txCount int
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

func (e *PEVM) Run(txs types.Transactions, isSysTxs bool) (types.Transactions, types.Receipts, uint64, bool, error) {
	if len(txs) < runtime.NumCPU() {
		return e.serialExecute(txs, isSysTxs)
	}
	return e.parallelExecute(txs, isSysTxs)
}

func (e *PEVM) serialExecute(txs types.Transactions, isSysTxs bool) (types.Transactions, types.Receipts, uint64, bool, error) {
	var (
		executedTxs types.Transactions
		receipts    types.Receipts
		timeout     bool
		err         error
	)
	if e.ctx.IsWorker() {
		executedTxs, receipts, timeout = e.commitTransactions(txs, isSysTxs)
	} else {
		receipts, err = e.applyTransactions(txs)
	}
	return executedTxs, receipts, e.usedGas, timeout, err
}

func (e *PEVM) parallelExecute(txs types.Transactions, isSysTxs bool) (types.Transactions, types.Receipts, uint64, bool, error) {
	return nil, nil, 0, false, nil
}

func (e *PEVM) applyTransaction(tx *types.Transaction) (*types.Receipt, error) {
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

func (e *PEVM) commitTransactions(txs types.Transactions, isSysTxs bool) (types.Transactions, types.Receipts, bool) {
	var (
		bd           = e.ctx.BlockDeadline()
		header       = e.ctx.Header()
		blockNumber  = header.Number
		parentHash   = header.ParentHash
		timestamp    = int64(header.Time)
		statedb      = e.ctx.StateDB()
		signer       = coretypes.MakeSigner(e.ctx.ChainConfig(), header.Number)
		committedTxs = make(types.Transactions, 0)
		receipts     = make(types.Receipts, 0)
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
	return committedTxs, receipts, timeout
}

func (e *PEVM) applyTransactions(txs types.Transactions) (types.Receipts, error) {
	var (
		receipts = make(types.Receipts, 0)
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
	return receipts, nil
}
