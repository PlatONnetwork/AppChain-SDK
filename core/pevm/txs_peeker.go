package pevm

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/miner"
)

type TxsPeeker struct {
	peekers  map[common.Address]miner.TxsPeeker
	txsCount int
}

func NewTxsPeeker(txs types.Transactions, signer types.Signer) *TxsPeeker {
	p := &TxsPeeker{
		peekers:  make(map[common.Address]miner.TxsPeeker),
		txsCount: len(txs),
	}
	if len(txs) == 0 {
		return p
	}

	begin := 0
	addr := txs[0].FromAddr(signer)
	for i := 1; i < len(txs); i++ {
		addr1 := txs[i].FromAddr(signer)
		if addr != addr1 || i == len(txs)-1 {
			end := i
			if end == len(txs)-1 {
				end = len(txs)
			}
			p.peekers[addr] = miner.NewAppTxsPeeker(txs[begin:end], signer)
			begin = i
			addr = addr1
		}
	}
	return p
}

func (p *TxsPeeker) Peeks(n int) types.Transactions {
	txs := make(types.Transactions, 0)
	if len(p.peekers) == 0 || p.txsCount <= 0 {
		return txs
	}

	peekTxs := func() {
		for _, peeker := range p.peekers {
			tx := peeker.Peek()
			peeker.Shift()
			if tx != nil {
				txcpy := tx
				txs = append(txs, txcpy)
				p.txsCount--
				if len(txs) >= n || p.txsCount <= 0 {
					break
				}
			}
		}
	}

	for len(txs) < n && p.txsCount > 0 {
		peekTxs()
	}
	return txs
}
