package pevm

import (
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	coretypes "github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName           = "pevm"
	ModuleVersion uint64 = 0
)

var (
	_ module.Module               = (*Module)(nil)
	_ module.FillTransationModule = (*Module)(nil)
	_ module.TxExecutorModule     = (*Module)(nil)
)

type Module struct {
	logger log.Logger
}

func NewModule(cliCtx *cli.Context) (*Module, error) {
	return &Module{
		logger: log.New("module", ModuleName),
	}, nil
}

func (m *Module) Name() string    { return ModuleName }
func (m *Module) Version() uint64 { return ModuleVersion }

func (m *Module) FillTransactions(ctx sdk.WorkerContext, cb sdk.TxApplyCallbackApp) (coretypes.Transactions, coretypes.Receipts, error) {
	var (
		allReceipts = make(types.Receipts, 0)
		allTxs      = make(types.Transactions, 0)
		usedGas     uint64
		signer      = types.MakeSigner(ctx.ChainConfig(), ctx.Header().Number)
		begin       = time.Now()
		pevm        = NewPEVM(m.logger, ctx, cb)
	)

	sysTxs, err := cb.AddTxs(ctx)
	if err != nil {
		m.logger.Error("Failed to add system transactions",
			"blockNumber", ctx.Header().Number,
			"parentHash", ctx.Header().ParentHash,
			"err", err)
		return allTxs, allReceipts, err
	}
	if len(sysTxs) > 0 {
		m.logger.Debug("add system transactions", "sysTxs", len(sysTxs))
		result, _ := pevm.Run(sysTxs, true)
		allTxs = append(allTxs, result.Transactions...)
		allReceipts = append(allReceipts, result.Receipts...)
		usedGas += result.GasUsed
	}

	txpool := ctx.Backend().TxPool()
	pending := txpool.Pending(true, true)
	localTxs, remoteTxs := make(map[common.Address]types.Transactions), pending
	for _, account := range txpool.Locals() {
		if txs := remoteTxs[account]; len(txs) > 0 {
			delete(remoteTxs, account)
			localTxs[account] = txs
		}
	}

	sortedTxs, err := cb.SortTxs(ctx, localTxs, remoteTxs)
	if err != nil {
		return nil, nil, err
	}
	if len(sortedTxs) > 0 {
		result, _ := pevm.Run(sortedTxs, false)
		allTxs = append(allTxs, result.Transactions...)
		allReceipts = append(allReceipts, result.Receipts...)
		usedGas += result.GasUsed
	} else {
		localtimeout := false
		if len(localTxs) > 0 {
			txs := types.NewTransactionsByPriceAndNonce(signer, localTxs, ctx.Header().BaseFee)
			result, _ := pevm.Run(txs.PeekAll(), false)
			allTxs = append(allTxs, result.Transactions...)
			allReceipts = append(allReceipts, result.Receipts...)
			localtimeout = result.Timeout
			usedGas += result.GasUsed
		}
		if !localtimeout && len(remoteTxs) > 0 {
			txs := types.NewTransactionsByPriceAndNonce(signer, remoteTxs, ctx.Header().BaseFee)
			result, _ := pevm.Run(txs.PeekAll(), false)
			allTxs = append(allTxs, result.Transactions...)
			allReceipts = append(allReceipts, result.Receipts...)
			usedGas += result.GasUsed
		}
	}
	// NOTE: need set gas used to header
	ctx.Header().GasUsed = usedGas
	m.logger.Info("Fill transactions success",
		"blockNumber", ctx.Header().Number,
		"parentHash", ctx.Header().ParentHash,
		"count", len(allTxs),
		"elapsed", time.Since(begin))
	return allTxs, allReceipts, nil
}

func (m *Module) ExecuteTxs(ctx sdk.WorkerContext, cApp sdk.ContractsApp, txs types.Transactions) (types.Receipts, uint64, error) {
	var (
		header = ctx.Header()
		now    = time.Now()
		pevm   = NewPEVM(m.logger, ctx, cApp)
	)

	result, err := pevm.Run(txs, false)
	m.logger.Info("Execute transactions finished",
		"blockNumber", header.Number,
		"blockHash", header.Hash(),
		"count", len(txs),
		"elapsed", time.Since(now))
	return result.Receipts, result.GasUsed, err
}
