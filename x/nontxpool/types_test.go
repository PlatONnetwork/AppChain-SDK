package nontxpool

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/test-go/testify/assert"
	"math/big"
	"testing"
)

// --- Tests ---
var signer = types.NewLondonSigner(big.NewInt(1))

func createAddr(i int64) common.Address {
	return common.BigToAddress(big.NewInt(i))
}
func createTx(addr common.Address, nonce uint64, gasPrice *big.Int) *types.Transaction {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		Value:    big.NewInt(1),
		Gas:      21000,
		GasPrice: gasPrice,
	})
	tx.CacheFromAddr(signer, addr)
	return tx
}
func TestAddLocalRemote(t *testing.T) {
	queue := NewTxQueue(signer, TxQueueConfig{GlobalTxCount: 1})
	queue.AddLocal([]*types.Transaction{createTx(createAddr(1), 0, big.NewInt(1))})
	assert.Equal(t, 1, len(queue.localTxsList))
	assert.Equal(t, 1, len(queue.localTxsList[createAddr(1)]))
	assert.Equal(t, uint64(1), queue.counter)
	queue.AddLocal([]*types.Transaction{createTx(createAddr(1), 1, big.NewInt(1))})
	assert.Equal(t, 1, len(queue.localTxsList))
	assert.Equal(t, 2, len(queue.localTxsList[createAddr(1)]))
	queue.AddRemote([]*types.Transaction{createTx(createAddr(2), 1, big.NewInt(1))})
	assert.Equal(t, 0, len(queue.remoteTxsList))
	assert.Equal(t, uint64(2), queue.counter)
	queue.localTxsList = make(map[common.Address][]*types.Transaction)
	queue.counter = 0
	queue.AddRemote([]*types.Transaction{createTx(createAddr(2), 1, big.NewInt(1))})
	assert.Equal(t, 1, len(queue.remoteTxsList))
	assert.Equal(t, uint64(1), queue.counter)
	queue.AddLocal([]*types.Transaction{createTx(createAddr(1), 0, big.NewInt(1))})
	assert.Equal(t, 1, len(queue.localTxsList))
	assert.Equal(t, 0, len(queue.remoteTxsList))

	assert.Equal(t, 1, len(queue.localTxsList[createAddr(1)]))
}
func TestReset(t *testing.T) {
	queue := NewTxQueue(signer, TxQueueConfig{GlobalTxCount: 1})
	queue.AddLocal([]*types.Transaction{createTx(createAddr(1), 1, big.NewInt(1))})

	queue.Reset(func(addr common.Address) uint64 {
		return 1
	})
	assert.Equal(t, 1, len(queue.localTxsList))
	queue.Reset(func(addr common.Address) uint64 {
		return 2
	})
	assert.Equal(t, 0, len(queue.localTxsList))
	assert.Equal(t, uint64(0), queue.counter)

	queue.AddRemote([]*types.Transaction{createTx(createAddr(2), 1, big.NewInt(1))})
	queue.Reset(func(addr common.Address) uint64 {
		return 1
	})
	assert.Equal(t, 1, len(queue.remoteTxsList))
	queue.Reset(func(addr common.Address) uint64 {
		return 2
	})
	assert.Equal(t, 0, len(queue.remoteTxsList))

}

func TestPending(t *testing.T) {
	queue := NewTxQueue(signer, TxQueueConfig{GlobalTxCount: 1})
	for i := 0; i < 10; i++ {
		queue.AddLocal([]*types.Transaction{createTx(createAddr(1), uint64(i), big.NewInt(1))})
		queue.AddLocal([]*types.Transaction{createTx(createAddr(2), uint64(i), big.NewInt(1))})
		queue.AddLocal([]*types.Transaction{createTx(createAddr(3), uint64(i), big.NewInt(1))})

	}
	txs := queue.Pending(func(addr common.Address) uint64 {
		return 0
	}, 10, 2)
	assert.Equal(t, 6, len(txs))
	txs = queue.Pending(func(addr common.Address) uint64 {
		return 0
	}, 10, 4)
	assert.Equal(t, 10, len(txs))
	txs = queue.Pending(func(addr common.Address) uint64 {
		return 9
	}, 10, 4)
	assert.Equal(t, 3, len(txs))
	txs = queue.Pending(func(addr common.Address) uint64 {
		return 10
	}, 10, 4)
	assert.Equal(t, 0, len(txs))

}
