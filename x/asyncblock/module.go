package asyncblock

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/core"
	"github.com/PlatONnetwork/AppChain-SDK/core/pevm"
	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName           = "asyncblock"
	ModuleVersion uint64 = 0
)

type Module struct {
	sync.Mutex
	FinalizeFeed
	backend             sdk.Backend
	chain               sdk.BlockchainContext
	logger              log.Logger
	p2p                 AsyncBlockP2P
	cs                  ConsensusState
	bsc                 *BlockStateCache
	computerSender      *ComputeSender
	stateComputeSender  *ComputeSender
	signer              types.Signer
	stateChangeCh       chan struct{}
	entrySizeLimit      int
	splitEntryThreshold int
	concurrencyLevel    int
	forceSequential     bool
	txsBatch            int
	blsKey              *bls.SecretKey
}

func NewModule(ctx *cli.Context) *Module {
	m := &Module{
		logger:             log.New("module", ModuleName),
		computerSender:     NewComputeSender(ctx.GlobalInt(ComputeSenderThreadFlag.Name), "p2p"),
		stateComputeSender: NewComputeSender(ctx.GlobalInt(ComputeSenderThreadFlag.Name), "exe"),

		stateChangeCh:       make(chan struct{}, 10),
		entrySizeLimit:      ctx.GlobalInt(EntrySizeFlag.Name),
		splitEntryThreshold: ctx.GlobalInt(SplitThresholdFlag.Name),
		concurrencyLevel:    ctx.GlobalInt(ConcurrencyLevelFlag.Name),
		forceSequential:     ctx.GlobalBool(ForceSequentialFlag.Name),
		txsBatch:            ctx.GlobalInt(TxsBatchFlag.Name),
	}
	m.p2p = NewAsyncBlockP2P(m.HandleMsg)
	m.bsc = NewBlockStateCache(m.stateComputeSender, m.genEnv, m.Finalize)
	m.SetBlsKey(ctx.GlobalString(BLsKeyFlagName))
	return m
}

func (m *Module) Name() string    { return ModuleName }
func (m *Module) Version() uint64 { return ModuleVersion }

func (m *Module) Init(ctx sdk.InitContext) error {
	m.signer = types.NewLondonSigner(ctx.Backend().ChainConfig().ChainID)
	m.backend = ctx.Backend()
	m.computerSender.Run(m.signer)
	m.stateComputeSender.Run(m.signer)
	go m.executeLoop()
	return nil
}

func (m *Module) SetBlsKey(path string) {
	if len(path) == 0 {
		return
	}
	priKey, err := bls.LoadBLS(path)
	if err != nil {
		m.logger.Crit("Decode bls key failed")
	}
	m.blsKey = priKey
}

func (m *Module) Finalize(executor *EntryExecutor) error {
	err := m.chain.ChainReader().BlockerApp().EndBlock(executor.workCtx)
	if err != nil {
		m.logger.Warn("EndBlock entry failed", "number", executor.header.Number, "err", err)
		return err
	}
	block, err := m.chain.ConsensusFinalize().Finalize(m.chain.ChainReader(), &executor.header, executor.statedb, executor.Transactions(), executor.receipts)
	if err != nil {
		m.logger.Warn("Finalize entry failed", "number", executor.header.Number, "err", err)
		return err
	}
	executor.block = block
	core.CommitDB(executor.statedb)
	m.Lock()
	defer m.Unlock()
	m.logger.Info("Add finish executor", "number", block.NumberU64())
	m.bsc.AddFinishExecutor(executor)
	m.SendFinalizeBlock(NewFinalizeBlockEvent{block, executor.statedb, executor.receipts})
	m.logger.Info("Finalize success", "cost", time.Since(executor.start), "number", block.NumberU64(), "sealHash", block.Header().SealHash(), "txs", block.Transactions().Len(), "header", block.Header(), "epoch", executor.epoch, "view", executor.view)
	return nil
}

func (m *Module) genEnv(header *types.Header, parentSealHash, parentHash common.Hash, parentNumber uint64) (sdk.StateDB, *pevm.PEVM, sdk.WorkerContext, error) {
	statedb, _ := m.bsc.FindStateDB(parentSealHash)
	parentBlock := m.chain.ChainReader().Engine().GetBlockByHashAndNum(parentHash, parentNumber)
	if parentBlock == nil {
		return nil, nil, nil, fmt.Errorf("no header")
	}
	if statedb == nil {
		m.logger.Info("Find statedb from cache failed", "parent", parentNumber)
		var err error
		if statedb, err = m.chain.StateDbMaker().MakeStateDB(parentBlock.Header()); err != nil {
			return nil, nil, nil, err
		}
	} else {
		statedb = core.CloneStateDB(statedb)
	}
	chainConfig := m.chain.ChainReader().Config()
	chainContext := m.chain.ChainReader()
	vmConfig := *m.chain.ChainReader().GetVMConfig()
	workCtx := NewAsyncWorkerContext(
		parentBlock, statedb, m.backend, header, chainConfig, &vmConfig)
	if err := m.chain.ChainReader().BlockerApp().BeginBlock(workCtx); err != nil {
		return nil, nil, nil, err
	}
	return statedb, pevm.NewPEVM(
		m.forceSequential,
		m.concurrencyLevel,
		m.txsBatch,
		m.logger,
		&pevm.Env{
			Header:        header,
			StateDB:       statedb,
			ChainConfig:   chainConfig,
			ChainContext:  chainContext,
			VMConfig:      vmConfig,
			IsWorker:      false,
			BlockDeadline: time.Time{},
		}, m.chain.ChainReader().ContractApp()), workCtx, nil
}

func (m *Module) CreateBlockExecutor(ctx sdk.BlockchainContext) (sdk.BlockExecutor, error) {
	m.chain = ctx
	return m, nil
}

func (m *Module) SplitEntries(blockNumber *big.Int, root common.Hash) bool {
	return m.bsc.IsSplit(blockNumber.Uint64())
}

// 执行区块
func (m *Module) ExecuteBlock(ctx sdk.BlockExecutorContext, block *types.Block, parent *types.Block) (sdk.StateDB, types.Receipts, error) {
	logger := m.logger.New("targetNumber", block.NumberU64(), "targetSeal", block.Header().SealHash(), "parentSeal", parent.Header().SealHash())
	logger.Debug("Execute block", "txs", block.Transactions().Len(), "header", block.Header())
	findBlock := func(ctx context.Context) (sdk.StateDB, types.Receipts, error) {
		if len(block.Transactions()) == 0 {
			return nil, nil, errors.New("empty block")
		}

		hash := block.Header().SealHash()
		logger.Debug("Try find block from state cache")
		m.Lock()

		if statedb, receipts := m.bsc.FindStateDB(hash); statedb != nil {
			m.Unlock()
			logger.Debug("Try find block success")
			return statedb, receipts, nil
		} else {
			if !m.bsc.HadFullEntry(block.NumberU64()) {
				m.Unlock()
				logger.Debug("No Full block entry, end the search")
				return nil, nil, errors.New("no full block entry")
			}

			ev := make(chan NewFinalizeBlockEvent, 10)
			sub := m.SubscribeNewFinalizeBlockEvent(ev)
			m.Unlock()
			logger.Info("Subscribe finish entry block")

			for {
				select {
				case <-ctx.Done():
					sub.Unsubscribe()
					logger.Warn("Find block failed, context done", "err", ctx.Err())
					return nil, nil, ctx.Err()
				case b := <-ev:
					if b.Block.NumberU64() == block.NumberU64() && (hash != b.Block.Header().SealHash() || b.Block.ReceiptHash() != block.ReceiptHash() || b.Block.Bloom() != block.Bloom()) {
						logger.Warn("Receive subscribe block, but is not target block", "number", b.Block.NumberU64(), "seal", b.Block.Header().SealHash(), "targetHeader", block.Header(), "header", b.Block.Header())
					}
					if hash == b.Block.Header().SealHash() {
						sub.Unsubscribe()
						m.logger.Debug("Receive subscribe block, is target block")
						return b.StateDB, b.Receipts, nil
					}
				}
			}
		}
	}
	deadline, _ := context.WithTimeout(context.Background(), time.Millisecond*1000)
	logger.Info("Try Find finish executor")
	statedb, receipts, err := findBlock(deadline)
	if err == nil {
		logger.Info("Find finish executor success")
		return statedb, receipts, nil
	}
	logger.Info("Try to execute block")
	statedb, vm, workCtx, err := m.genEnv(block.Header(), parent.Header().SealHash(), parent.Hash(), parent.NumberU64())
	if err != nil {
		return nil, nil, err
	}
	result, err := vm.Run(block.Transactions(), false)
	if err != nil {
		m.logger.Warn("Execute block failed", "err", err)
		return nil, nil, err
	}
	if err := m.chain.ChainReader().BlockerApp().EndBlock(workCtx); err != nil {
		m.logger.Warn("End block failed", "err", err)
		return nil, nil, err
	}
	if _, err := m.chain.ConsensusFinalize().Finalize(m.chain.ChainReader(), block.Header(), statedb, block.Transactions(), result.Receipts); err != nil {
		m.logger.Warn("Finalize block failed", "err", err)
		return nil, nil, err
	}
	if err := m.chain.ValidateBlock().ValidateState(block, statedb, result.Receipts, result.GasUsed); err != nil {
		m.logger.Error("Validate state failed", "err", err)
		return nil, nil, err
	}
	return statedb, result.Receipts, nil
}

func (m *Module) BlockBody(ctx sdk.BlockExecutorContext, blockNumber *big.Int, epoch, view uint64) (types.Transactions, error) {
	m.logger.Debug("Find transaction", "number", blockNumber, "epoch", epoch, "view", view)
	return m.bsc.TryFindBody(view, blockNumber.Uint64())
}

func (m *Module) FillTransactions(ctx sdk.WorkerContext, cb sdk.TxApplyCallbackApp) (types.Transactions, types.Receipts, error) {
	var (
		allReceipts = make(types.Receipts, 0)
		allTxs      = make(types.Transactions, 0)
		usedGas     uint64
		begin       = time.Now()
		pevm        = pevm.NewPEVM(
			m.forceSequential,
			m.concurrencyLevel,
			m.txsBatch,
			m.logger,
			&pevm.Env{
				Header:        ctx.Header(),
				StateDB:       ctx.StateDB(),
				ChainConfig:   ctx.ChainConfig(),
				ChainContext:  ctx.Backend().ChainContext(),
				VMConfig:      *ctx.VMConfig(),
				IsWorker:      ctx.IsWorker(),
				BlockDeadline: ctx.BlockDeadline(),
			}, cb)
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
		m.logger.Debug("Add system transactions", "sysTxs", len(sysTxs))
		result, err := pevm.Run(sysTxs, true)
		if err != nil {
			m.logger.Error("Failed to execute system transactions",
				"blockNumber", ctx.Header().Number,
				"parentHash", ctx.Header().ParentHash,
				"err", err)
			return allTxs, allReceipts, err
		}
		allTxs = append(allTxs, result.Transactions...)
		allReceipts = append(allReceipts, result.Receipts...)
		usedGas = result.GasUsed
	}
	epoch, view := m.cs.EpochView()
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
	m.logger.Debug("Get all transaction", "epoch", epoch, "view", view, "number", ctx.Header().Number.Uint64(), "sortedTxs", sortedTxs.Len())
	if len(sortedTxs) > m.splitEntryThreshold {
		index := 0
		entryNumber := uint32(0)
		for index < len(sortedTxs) {
			end := index + m.entrySizeLimit
			if end > len(sortedTxs) {
				end = len(sortedTxs)
			}
			result, err := pevm.Run(sortedTxs[index:end], false)
			if err != nil {
				m.logger.Error("Failed to execute sorted transactions",
					"blockNumber", ctx.Header().Number,
					"parentHash", ctx.Header().ParentHash,
					"err", err)
				return allTxs, allReceipts, err
			}
			allTxs = append(allTxs, result.Transactions...)
			allReceipts = append(allReceipts, result.Receipts...)
			usedGas = result.GasUsed
			entry := Entry{
				Epoch:          epoch,
				View:           view,
				BlockNumber:    ctx.Header().Number.Uint64(),
				ParentSealHash: ctx.ParentBlock().Header().SealHash(),
				Header:         nil,
				EntryNumber:    entryNumber,
				Transactions:   result.Transactions,
				Ending:         0,
				Signature:      nil,
			}

			if entryNumber == 0 {
				m.bsc.AddSplit(ctx.Header().Number.Uint64())
				entry.Transactions = allTxs[0:len(allTxs)]
				entry.Header = ctx.Header()
			}
			if end == sortedTxs.Len() || result.Timeout {
				entry.Ending = 1
			}
			entryNumber++
			entry.Sign(m.blsKey)
			m.p2p.BroadcastEntry(&EntryMsg{Entry: entry})
			log.Info("Broadcast fill", "blockNumber", entry.BlockNumber, "entryNumber", entry.EntryNumber)
			index = end
			if result.Timeout {
				m.logger.Warn("Fill transaction timeout", "blockNumber", entry.BlockNumber, "entryNumber", entry.EntryNumber)
				break
			}
		}
	} else if len(sortedTxs) > 0 {
		result, err := pevm.Run(sortedTxs, false)
		if err != nil {
			m.logger.Error("Failed to execute sorted transactions",
				"blockNumber", ctx.Header().Number,
				"parentHash", ctx.Header().ParentHash,
				"err", err)
			return allTxs, allReceipts, err
		}
		allTxs = append(allTxs, result.Transactions...)
		allReceipts = append(allReceipts, result.Receipts...)
		usedGas = result.GasUsed
	}
	// NOTE: need set gas used to header
	ctx.Header().GasUsed = usedGas
	m.logger.Info("Fill transactions success",
		"blockNumber", ctx.Header().Number,
		"parentHash", ctx.Header().ParentHash,
		"gasUsed", usedGas,
		"count", len(allTxs),
		"elapsed", time.Since(begin))
	return allTxs, allReceipts, nil
}
func (m *Module) ExecuteTxs(ctx sdk.WorkerContext, cApp sdk.ContractsApp, txs types.Transactions) (types.Receipts, uint64, error) {
	var (
		header = ctx.Header()
		now    = time.Now()
		pevm   = pevm.NewPEVM(
			m.forceSequential,
			m.concurrencyLevel,
			m.txsBatch,
			m.logger,
			&pevm.Env{
				Header:        ctx.Header(),
				StateDB:       ctx.StateDB(),
				ChainConfig:   ctx.ChainConfig(),
				ChainContext:  ctx.Backend().ChainContext(),
				VMConfig:      *ctx.VMConfig(),
				IsWorker:      ctx.IsWorker(),
				BlockDeadline: ctx.BlockDeadline(),
			}, cApp)
	)

	result, err := pevm.Run(txs, false)
	m.logger.Info("Execute transactions finished",
		"blockNumber", header.Number,
		"blockHash", header.Hash(),
		"count", len(txs),
		"err", err,
		"elapsed", time.Since(now))
	return result.Receipts, result.GasUsed, err
}

func (m *Module) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {
	m.cs.Update(ctx.Epoch(), ctx.View(), validators, int(ctx.View())%len(validators))
	m.p2p.Update(validators)
	go m.cleanState(ctx.Epoch(), ctx.View())
}
func (m *Module) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	m.bsc.RemoteSplit(block.NumberU64())
	return nil
}
func (m *Module) Protocols() []p2p.Protocol {
	return m.p2p.Protocols()
}

func (m *Module) cleanState(epoch, view uint64) {
	m.logger.Debug("Clean state", "epoch", epoch, "view", view)
	m.bsc.Clean(epoch, view)
	//todo clean blockstate
}
func (m *Module) executeLoop() {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-m.stateChangeCh:
			log.Info("state change ch")
			go m.bsc.FindExecutableEntry()
		case <-tick:
			log.Info("state change tick")
			//todo find next executable entry
			go m.bsc.FindExecutableEntry()
		}
	}
}

func (m *Module) handleEntry(peer sdkp2p.Peer, msg *EntryMsg) error {
	m.logger.Debug("Handle receive entry ", "epoch", msg.Epoch, "view", msg.View, "blockNumber", msg.BlockNumber, "entryNumber", msg.EntryNumber)
	//verify signature
	if err := m.cs.VerifyEntry(&msg.Entry); err != nil {
		m.logger.Error("Verify entry failed", "err", err)
		return nil
	}

	m.bsc.AddEntry(&msg.Entry)
	m.computerSender.AddEntry(&msg.Entry)
	m.stateChangeCh <- struct{}{}

	m.logger.Debug("Add entry success", "epoch", msg.Epoch, "view", msg.View, "blockNumber", msg.BlockNumber, "entryNumber", msg.EntryNumber)
	return nil
}

func (m *Module) handleGetEntry(peer sdkp2p.Peer, msg *GetEntryMsg) error {
	entry := m.bsc.GetEntry(msg.BlockNumber, msg.View, msg.EntryNumber)
	if entry != nil {
		m.p2p.Send(peer, &EntryMsg{Entry: *entry})
	}
	return nil
}

func (m *Module) HandleMsg(peer sdkp2p.Peer, msg sdkp2p.Message) error {
	switch s := msg.(type) {
	case *EntryMsg:
		return m.handleEntry(peer, s)
	case *GetEntryMsg:
		m.handleGetEntry(peer, s)
	default:
		return errors.New("unknown message type")
	}
	return nil
}

func (m *Module) fillSender(txs types.Transactions) {
	start := time.Now()

	end := uint32(txs.Len())
	index := atomic.Uint32{}
	group := sync.WaitGroup{}
	group.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer group.Done()
			for {
				pos := index.Add(1)
				if pos > end {
					return
				}
				types.Sender(m.signer, txs[pos-1])
			}
		}()
	}
	group.Wait()
	m.logger.Debug("Fill sender", "cost", time.Since(start))
}
