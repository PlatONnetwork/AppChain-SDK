package nontxpool

import (
	"context"
	"errors"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/rand"
	"sync"
	"time"

	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
)

const (
	ModuleVersion uint64 = 0
	ModuleName           = "nontxpool"

	txChanSize = 4096

	// txSlotSize is used to calculate how many data slots a single transaction
	// takes up based on its size. The slots are used as DoS protection, ensuring
	// that validating a new transaction remains a constant operation (in reality
	// O(maxslots), where max slots are 4 currently).
	txSlotSize = 32 * 1024

	txMaxSize = 32 * txSlotSize // 128KB
)

var (
	// ErrTipVeryHigh is a sanity error to avoid extremely big numbers specified
	// in the tip field.
	ErrTipVeryHigh = errors.New("max priority fee per gas higher than 2^256-1")

	// ErrTipAboveFeeCap is a sanity error to ensure no one is able to specify a
	// transaction with a tip higher than the total fee cap.
	ErrTipAboveFeeCap = errors.New("max priority fee per gas higher than max fee per gas")

	// ErrFeeCapVeryHigh is a sanity error to avoid extremely big numbers specified
	// in the fee cap field.
	ErrFeeCapVeryHigh = errors.New("max fee per gas higher than 2^256-1")

	// ErrFeeCapTooLow is returned if the transaction fee cap is less than the
	// the base fee of the block.
	ErrFeeCapTooLow = errors.New("max fee per gas less than block base fee")

	// ErrNegativeValue is a sanity error to ensure no one is able to specify a
	// transaction with a negative value.
	ErrNegativeValue = errors.New("negative value")
	// ErrOversizedData is returned if the input data of a transaction is greater
	// than some meaningful limit a user might use. This is not a consensus error
	// making the transaction invalid, rather a DOS protection.
	ErrOversizedData = errors.New("oversized data")

	// ErrInvalidSender is returned if the transaction contains an invalid signature.
	ErrInvalidSender = errors.New("invalid sender")
)

// NonTxPoolModule
type ConsensusState struct {
	sync.Mutex
	Self enode.ID

	Leader     enode.ID
	Epoch      uint64
	ViewNumber uint64
	Validator  []*cbfttypes.ValidateNode
}

func (c *ConsensusState) Update(epoch, viewNumber uint64, validator []*cbfttypes.ValidateNode, leader enode.ID) {
	c.Lock()
	defer c.Unlock()
	c.Epoch = epoch
	c.ViewNumber = viewNumber
	c.Validator = validator
	c.Leader = leader
}
func (c *ConsensusState) LeaderId() string {
	c.Lock()
	defer c.Unlock()
	return c.Leader.String()
}
func (c *ConsensusState) IsLeader() bool {
	c.Lock()
	defer c.Unlock()
	if c.Leader == c.Self {
		return true
	} else {
		return false
	}
}

func (c *ConsensusState) IsValidator(enode.ID) bool {
	c.Lock()
	defer c.Unlock()
	for _, node := range c.Validator {
		if node.NodeID == c.Self {
			return true
		}
	}
	return false
}
func (c *ConsensusState) ValidatorsNodeId() map[string]struct{} {
	c.Lock()
	defer c.Unlock()
	ids := make(map[string]struct{})
	for _, node := range c.Validator {
		ids[node.NodeID.String()] = struct{}{}
	}
	return ids
}

func NewModule(ctx *cli.Context) *Module {
	logger := log.New("module", ModuleName)
	txsCacheSize := ctx.GlobalInt(TxsCacheSizeFlag.Name)
	broadcastInterval := ctx.GlobalInt(BroadcastIntervalFlag.Name)
	logger.Debug("Get params", "TxsCacheSize", txsCacheSize, "BroadcastInterval", broadcastInterval)
	m := &Module{
		logger: logger,

		txsCacheSize:      txsCacheSize,
		broadcastInterval: time.Duration(broadcastInterval) * time.Millisecond,
		remoteTxCh:        make(chan *TransactionsPacket, txChanSize),
		txsCache:          make([]*types.Transaction, 0),

		txBroadcast: make(chan []*types.Transaction),
	}
	validateTx := func(tx *types.Transaction) error { return nil }
	if ctx.GlobalBool(ValidateTxFlag.Name) {
		validateTx = m.validateTx
	}
	m.NonTxPoolP2P = NewNonTxPoolP2P(validateTx, m.remoteTxCh)

	return m
}

type Module struct {
	logger log.Logger

	txsCacheSize      int
	broadcastInterval time.Duration
	signer            types.Signer

	txpool sdk.TxPool
	cs     ConsensusState
	*NonTxPoolP2P

	txsCache []*types.Transaction

	txsSub     event.Subscription    // Subscription for new transaction event
	localtxCh  chan core.NewTxsEvent // Channel to receive new transactions event
	remoteTxCh chan *TransactionsPacket

	txBroadcast chan []*types.Transaction // Channel used to queue transaction propagation requests

}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) Init(ctx sdk.InitContext) error {

	m.txpool = ctx.Backend().TxPool()
	chainID, err := ctx.Backend().ChainId()
	if err != nil {
		return err
	}
	m.signer = types.NewLondonSigner(chainID)

	id := enode.PubkeyToIDV4(&ctx.NodeKey().PublicKey)
	m.cs = ConsensusState{Self: id}

	ctx.Backend().SetNoTxBroadcast(true)
	m.localtxCh = make(chan core.NewTxsEvent, txChanSize)
	m.txsSub = m.txpool.SubscribeNewTxsEvent(m.localtxCh)

	go m.txLoop(context.Background())
	go m.broadcastTransactions(context.Background())

	if err := m.NonTxPoolP2P.Run(); err != nil {
		m.logger.Error("Start p2p failed", "err", err)
		return err
	}
	return nil
}

func (m *Module) txLoop(ctx context.Context) {

	timer := time.NewTimer(m.broadcastInterval)

	for {
		select {
		// 来自本地的交易
		case ev := <-m.localtxCh:
			if m.cs.IsLeader() {
				continue
			}
			m.txsCache = append(m.txsCache, ev.Txs...)
			if len(ev.Txs) > m.txsCacheSize {
				m.txBroadcast <- m.txsCache
				m.txsCache = make([]*types.Transaction, 0)
				timer.Reset(m.broadcastInterval)
			}

		case ev := <-m.remoteTxCh:
			if m.cs.IsLeader() {
				m.logger.Debug("I'm leader, add remote txs to the txpool", "account", len(ev.txs))
				for _, tx := range ev.txs {
					if err := m.txpool.AddRemote(tx); err != nil {
						m.logger.Warn("Add remote tx failed", "err", err, "tx", tx.Hash())
					}
				}
			} else {
				m.txsCache = append(m.txsCache, ev.txs...)
				if len(ev.txs) > m.txsCacheSize {
					m.txBroadcast <- m.txsCache
					m.txsCache = make([]*types.Transaction, 0)
					timer.Reset(m.broadcastInterval)
				}
			}
		case <-timer.C:
			if len(m.txsCache) > 0 {
				m.txBroadcast <- m.txsCache
				m.txsCache = make([]*types.Transaction, 0)
			}
			timer.Reset(m.broadcastInterval)
		case <-m.txsSub.Err():
			return
		case <-ctx.Done():
			return
		}
	}
}

func (m *Module) sendTx(txs []*types.Transaction) {
	//选择leader转发交易
	var sendPeer sdkp2p.Peer
	peers := m.NonTxPoolP2P.p2p.Peers()
	leader := m.cs.LeaderId()
	for _, peer := range peers {
		if leader == peer.Id() {
			sendPeer = peer
			break
		}
	}

	if sendPeer == nil {
		var tmp []sdkp2p.Peer
		validatorIds := m.cs.ValidatorsNodeId()
		for _, peer := range peers {
			if _, ok := validatorIds[peer.Id()]; ok {
				tmp = append(tmp, peer)
			}
		}

		if len(tmp) > 0 {
			randomIndex := rand.Intn(len(tmp))
			sendPeer = tmp[randomIndex]
		}
		m.logger.Debug("Didn't found leader connection, select a validator connection", "leader", leader, "peer", sendPeer.Id())

	}
	if sendPeer != nil {
		m.logger.Debug("Send TransactionsMessage", "peer", sendPeer.Id(), "len", len(txs))
		m.NonTxPoolP2P.p2p.Send(sendPeer, &TransactionsMessage{
			Txs: txs,
		})
	}
}

func (m *Module) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {

	// length := cbft.validatorPool.Len(cbft.state.Epoch())
	//	currentProposer := cbft.state.ViewNumber() % uint64(length)
	leader := validators[int(ctx.View())%len(validators)].NodeID
	m.cs.Update(ctx.Epoch(), ctx.View(), validators, leader)
	m.logger.Info("View change", "proposer", ctx.IsProposer(), "epoch", ctx.Epoch(), "view", ctx.View(), "leader", leader)
	if !ctx.IsProposer() {
		go func() {
			txs := m.txpool.RemoteTxs()
			m.logger.Debug("I'm not a leader, try to remove remote txs", "accounts", len(txs))
			for _, transactions := range txs {
				for _, transaction := range transactions {
					m.txpool.RemoveTx(transaction.Hash(), false)
				}
			}

			localTxs := m.txpool.LocalTxs()
			m.logger.Debug("I'm not a leader, try to send local tx", "accounts", len(localTxs))

			for _, transactions := range localTxs {
				m.sendTx(transactions)
			}
		}()
	}
}

func (m *Module) Protocols() []p2p.Protocol {
	return m.p2p.Protocol()
}

// broadcastTransactions is a write loop that schedules transaction broadcasts
// to the remote peer. The goal is to have an async writer that does not lock up
// node internals and at the same time rate limits queued data.
func (m *Module) broadcastTransactions(ctx context.Context) {
	var (
		queue  []*types.Transaction  // Queue of hashes to broadcast as full transactions
		done   chan struct{}         // Non-nil if background broadcaster is running
		fail   = make(chan error, 1) // Channel used to receive network error
		failed bool                  // Flag whether a send failed, discard everything onward
	)
	for {
		// If there's no in-flight broadcast running, check if a new one is needed
		if done == nil && len(queue) > 0 {
			// Pile transaction until we reach our allowed network limit
			var (
				hashesCount uint64
				txs         []*types.Transaction
				size        common.StorageSize
			)
			for i := 0; i < len(queue) && size < maxTxPacketSize; i++ {

				tx := queue[i]
				txs = append(txs, tx)
				size += tx.Size()

				hashesCount++
			}
			// Shift and trim queue
			queue = queue[:copy(queue, queue[hashesCount:])]

			// If there's anything available to transfer, fire up an async writer
			if len(txs) > 0 {
				done = make(chan struct{})
				go func() {
					m.sendTx(txs)
					close(done)
					m.logger.Trace("Send transactions", "count", len(txs))
				}()
			}
		}
		// Transfer goroutine may or may not have been started, listen for events
		select {
		case txs := <-m.txBroadcast:
			// If the connection failed, discard all transaction events
			if failed {
				continue
			}
			// New batch of transactions to be broadcast, queue them (with cap)
			queue = append(queue, txs...)
			if len(queue) > maxQueuedTxs {
				// Fancy copy and resize to ensure buffer doesn't grow indefinitely
				queue = queue[:copy(queue, queue[len(queue)-maxQueuedTxs:])]
			}

		case <-done:
			done = nil

		case <-fail:
			failed = true
		case <-ctx.Done():
			return
		}
	}
}

func (m *Module) validateTx(tx *types.Transaction) error {
	// Reject transactions over defined size to prevent DOS attacks
	if uint64(tx.Size()) > txMaxSize {
		return ErrOversizedData
	}
	// Transactions can't be negative. This may never happen using RLP decoded
	// transactions but may occur if you create a transaction using the RPC.
	if tx.Value().Sign() < 0 {
		return ErrNegativeValue
	}
	// Sanity check for extremely large numbers
	if tx.GasFeeCap().BitLen() > 256 {
		return ErrFeeCapVeryHigh
	}
	if tx.GasTipCap().BitLen() > 256 {
		return ErrTipVeryHigh
	}
	// Ensure gasFeeCap is greater than or equal to gasTipCap.
	if tx.GasFeeCapIntCmp(tx.GasTipCap()) < 0 {
		return ErrTipAboveFeeCap
	}
	// Make sure the transaction is signed properly.
	_, err := types.Sender(m.signer, tx)
	if err != nil {
		log.Debug("validateTx  fail", "tx", tx.Hash(), "err", err)
		return ErrInvalidSender
	}

	return nil
}
