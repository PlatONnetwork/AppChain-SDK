package types

import "math/big"

type Delegation struct {
	BlockNumber uint64
	Amount      *big.Int
}

func NewDelegation(blockNumber uint64, amount *big.Int) *Delegation {
	return &Delegation{
		BlockNumber: blockNumber,
		Amount:      amount,
	}
}

func (d *Delegation) UpdateBlockNumber(blockNumber uint64) {
	d.BlockNumber = blockNumber
}

func (d *Delegation) AddAmount(amount *big.Int) {
	d.Amount = new(big.Int).Add(d.Amount, amount)
}

func (d *Delegation) SubAmount(amount *big.Int) {
	d.Amount = new(big.Int).Sub(d.Amount, amount)
}
