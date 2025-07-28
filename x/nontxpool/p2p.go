package nontxpool

import (
	"errors"
	"sync"
	"time"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"

	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
)

const (
	recordsCache = time.Second * 20
)

type TransactionsPacket struct {
	id  string
	txs []*types.Transaction
}

type NonTxPoolP2P struct {
	p2p            *sdkp2p.Protocol
	checkTx        func(transaction *types.Transaction) error
	forwardTxCache *ForwardCache

	txNotify chan<- *TransactionsPacket
}

func NewNonTxPoolP2P(checkTx func(transaction *types.Transaction) error, txNotify chan<- *TransactionsPacket) *NonTxPoolP2P {
	p := &NonTxPoolP2P{}
	protocol := sdkp2p.NewProtocol("nontxpool", 1, 10)
	protocol.RegistryMessageType([]sdkp2p.Message{&TransactionsMessage{}})
	protocol.SetNewPeer(func(p *p2p.Peer, rw p2p.MsgReadWriter) sdkp2p.Peer {
		return &Peer{
			DefaultPeer: sdkp2p.NewDefaultPeer(p, rw),
		}
	})
	protocol.SetUserHandleMsg(func(peer sdkp2p.Peer, msg sdkp2p.Message) error {
		return p.handleMsg(peer, msg)
	})
	p.p2p = protocol
	p.checkTx = checkTx
	p.forwardTxCache = NewForwardCache(recordsCache)
	p.txNotify = txNotify
	return p
}

func (n *NonTxPoolP2P) Run() error {
	//receive txpool transaction event and forward tx message

	go n.forwardTxCache.startCleaner()
	return nil
}

func (n *NonTxPoolP2P) Protocols() []p2p.Protocol {
	return n.p2p.Protocol()
}

func (n *NonTxPoolP2P) handleMsg(peer sdkp2p.Peer, msg sdkp2p.Message) error {
	switch m := msg.(type) {
	case *TransactionsMessage:
		newTxs := make([]*types.Transaction, 0)

		if n.checkTx != nil {
			for i, tx := range m.Txs {
				if !n.forwardTxCache.CanForward(tx.Hash()) {
					log.Trace("Skipping duplicate message", "tx", tx.Hash())
					continue
				}
				if n.checkTx(tx) == nil {
					newTxs = append(newTxs, m.Txs[i])
				}
			}

		} else {
			for i, tx := range m.Txs {
				if !n.forwardTxCache.CanForward(tx.Hash()) {
					log.Trace("Skipping duplicate message", "tx", tx.Hash())
					continue
				}
				newTxs = append(newTxs, m.Txs[i])
			}
		}
		if len(newTxs) > 0 {
			n.txNotify <- &TransactionsPacket{
				id:  peer.Id(),
				txs: newTxs,
			}
		}
	default:
		return errors.New("unknown message type")
	}
	return nil
}

// 转发缓存
type ForwardCache struct {
	mu      sync.Mutex
	records map[common.Hash]time.Time
	ttl     time.Duration
}

func NewForwardCache(ttl time.Duration) *ForwardCache {
	fc := &ForwardCache{
		records: make(map[common.Hash]time.Time),
		ttl:     ttl,
	}
	go fc.startCleaner()
	return fc
}

// 检查是否可以转发，并清理过期项
func (fc *ForwardCache) CanForward(tx common.Hash) bool {
	fc.mu.Lock()
	defer fc.mu.Unlock()

	now := time.Now()
	if t, exists := fc.records[tx]; exists {
		if now.Sub(t) < fc.ttl {
			return false
		}
	}
	// 记录这次转发
	fc.records[tx] = now
	return true
}

// 每秒清理一次过期记录
func (fc *ForwardCache) startCleaner() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for now := range ticker.C {
		fc.mu.Lock()
		for id, t := range fc.records {
			if now.Sub(t) >= fc.ttl {
				delete(fc.records, id)
			}
		}
		fc.mu.Unlock()
	}
}
