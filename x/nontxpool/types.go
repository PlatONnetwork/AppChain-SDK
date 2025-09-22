package nontxpool

import (
	"errors"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"sync"
)

var (
	CheckQueue = false
)

type TxQueue struct {
	sync.Mutex
	logger        log.Logger
	signer        types.Signer
	counter       uint64
	conf          TxQueueConfig
	localTxsList  map[common.Address][]*types.Transaction
	remoteTxsList map[common.Address][]*types.Transaction
}

func NewTxQueue(signer types.Signer, conf TxQueueConfig) *TxQueue {
	return &TxQueue{
		logger:        log.New("module", ModuleName, "component", "txqueue"),
		signer:        signer,
		conf:          conf,
		localTxsList:  make(map[common.Address][]*types.Transaction),
		remoteTxsList: make(map[common.Address][]*types.Transaction),
	}
}

func (t *TxQueue) Local(addr common.Address) bool {
	t.Lock()
	defer t.Unlock()
	return t.localTxsList[addr] != nil
}

func (t *TxQueue) Remote(addr common.Address) bool {
	t.Lock()
	defer t.Unlock()
	return t.remoteTxsList[addr] != nil
}
func (t *TxQueue) Total() uint64 {
	t.Lock()
	defer t.Unlock()
	return t.counter
}
func (t *TxQueue) truncate() {
	for t.counter > t.conf.GlobalTxCount && len(t.remoteTxsList) != 0 {
		for addr, v := range t.remoteTxsList {
			t.counter -= uint64(len(v))
			delete(t.remoteTxsList, addr)
			break
		}
	}
}
func (t *TxQueue) addLocalTxs(addr common.Address, txs ...*types.Transaction) {
	var txsList []*types.Transaction
	if txsList = t.localTxsList[addr]; txsList == nil {
		if txsList = t.remoteTxsList[addr]; txsList != nil {
			delete(t.remoteTxsList, addr)
		}
	}
	t.localTxsList[addr] = append(txsList, txs...)
	t.counter += uint64(len(txs))
	t.truncate()
}

func (t *TxQueue) addRemoteTxs(addr common.Address, txs ...*types.Transaction) {
	t.remoteTxsList[addr] = append(t.remoteTxsList[addr], txs...)
	t.counter += uint64(len(txs))
	t.truncate()
}
func (t *TxQueue) AddLocalBatch(addr common.Address, txs []*types.Transaction) {
	t.Lock()
	defer t.Unlock()
	if pendingNonce, err := t.pendingNonce(addr); err == nil {
		nonce := txs[0].Nonce()
		if nonce != pendingNonce {
			if nonce > pendingNonce || pendingNonce-nonce < uint64(len(txs)) {
				return
			}
			txs = txs[pendingNonce-nonce:]
		}
	}

	t.addLocalTxs(addr, txs...)
	t.checkQueue()

}

func (t *TxQueue) AddLocal(txs []*types.Transaction) {
	t.Lock()
	defer t.Unlock()
	t.addLocal(txs)
	t.checkQueue()

}
func (t *TxQueue) addLocal(txs []*types.Transaction) {
	for _, tx := range txs {
		from := tx.FromAddr(t.signer)
		if pendingNonce, err := t.pendingNonce(from); err == nil && pendingNonce != tx.Nonce() {
			t.logger.Debug("Add local tx failed, nonce doesn't match", "from", from.Hex(), "pending", pendingNonce, "tx", tx.Nonce())
			continue
		}
		t.addLocalTxs(from, tx)
	}
	t.truncate()
}
func (t *TxQueue) AddRemote(txs []*types.Transaction) {
	t.Lock()
	defer t.Unlock()
	for _, tx := range txs {
		from := tx.FromAddr(t.signer)
		if txs, ok := t.localTxsList[from]; ok {
			if txs[len(txs)-1].Nonce()+1 == tx.Nonce() {
				t.addLocalTxs(from, tx)
			}
			continue
		}
		if txs, ok := t.remoteTxsList[from]; ok {
			if txs[len(txs)-1].Nonce()+1 != tx.Nonce() {
				continue
			}
		}
		t.addRemoteTxs(from, tx)
	}
	t.checkQueue()

}
func (t *TxQueue) Reset(getNonce func(addr common.Address) uint64) {
	t.Lock()
	defer t.Unlock()
	resetFn := func(txsList map[common.Address][]*types.Transaction) {
		deletedMap := make(map[common.Address]int)
		for addr, txs := range txsList {
			nonce := getNonce(addr)
			if len(txs) != 0 {
				if txs[0].Nonce() < nonce {
					pos := nonce - txs[0].Nonce()
					if pos >= uint64(len(txs)) {
						deletedMap[addr] = len(txs)
						continue
					}
					txsList[addr] = txs[pos:]
					t.counter -= pos
				}
			}
		}
		for addr, amount := range deletedMap {
			t.counter -= uint64(amount)
			delete(txsList, addr)
		}
	}

	resetFn(t.localTxsList)
	resetFn(t.remoteTxsList)
	t.checkQueue()

}
func (t *TxQueue) CleanRemote() {
	t.Lock()
	defer t.Unlock()
	for _, txs := range t.remoteTxsList {
		t.counter -= uint64(len(txs))
	}
	t.remoteTxsList = make(map[common.Address][]*types.Transaction)
	t.checkQueue()

}
func (t *TxQueue) checkQueue() {
	if !CheckQueue {
		return
	}
	sum := 0
	for _, v := range t.localTxsList {
		sum += len(v)
	}
	for _, v := range t.remoteTxsList {
		sum += len(v)
	}
	if uint64(sum) != t.counter {
		panic(fmt.Sprintf("sum:%d, counter:%d", sum, t.counter))
	}
}
func (t *TxQueue) Locals(handle func([]*types.Transaction)) {
	t.Lock()
	defer t.Unlock()

	for _, txsList := range t.localTxsList {
		handle(txsList)
	}
	t.checkQueue()
}
func (t *TxQueue) pendingNonce(addr common.Address) (uint64, error) {
	if txs := t.localTxsList[addr]; len(txs) != 0 {
		return txs[len(txs)-1].Nonce() + 1, nil
	}
	if txs := t.remoteTxsList[addr]; len(txs) != 0 {
		return txs[len(txs)-1].Nonce() + 1, nil
	}
	return 0, errors.New("unknown address")
}
func (t *TxQueue) PendingNonce(addr common.Address) (uint64, error) {
	t.Lock()
	defer t.Unlock()
	return t.pendingNonce(addr)
}

func (t *TxQueue) Pending(getNonce func(addr common.Address) uint64, limit int, txsPerAccount int) types.Transactions {
	t.Lock()
	defer t.Unlock()
	var txs []types.Transactions
	sum := 0
	getTxs := func(txsList map[common.Address][]*types.Transaction) {
		t.logger.Debug("Pending Txs", "len", len(txsList))
		for k, v := range txsList {
			if len(v) == 0 {
				continue
			}
			if sum >= limit {
				t.logger.Debug("Enough Txs", "address", k, "len", len(v), "limit", limit, "txsPerAccount", txsPerAccount, "sum", sum)
				break
			}
			nonce := getNonce(k)
			if nonce < v[0].Nonce() {
				t.logger.Debug("Nonce too high", "nonce", nonce, "firstNonce", v[0].Nonce())
				continue
			}
			start := int(nonce - v[0].Nonce())
			if start > len(v) {
				t.logger.Debug("Nonce too high", "nonce", nonce, "firstNonce", v[0].Nonce(), "len", len(v))
				continue
			}
			end := start + txsPerAccount
			if end > len(v) {
				end = len(v)
			}
			if sum+end-start > limit {
				end = start + limit - sum
			}
			t.logger.Debug("Pending Txs", "address", k, "len", len(v), "nonce", nonce, "firstNonce", v[0].Nonce(), "start", start, "end", end, "limit", limit, "txsPerAccount", txsPerAccount, "sum", sum)
			txs = append(txs, v[start:end])

			sum += end - start

		}
	}
	getTxs(t.localTxsList)
	getTxs(t.remoteTxsList)
	res := make(types.Transactions, 0, sum)
	for _, s := range txs {
		res = append(res, s...)
	}
	return res
}
