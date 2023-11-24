package types

import "math/big"

type Delegation struct {
	Epoch  uint64
	Amount *big.Int
}

func NewDelegation(epoch uint64, amount *big.Int) *Delegation {
	return &Delegation{
		Epoch:  epoch,
		Amount: amount,
	}
}

func (d *Delegation) UpdateEpoch(epoch uint64) {
	d.Epoch = epoch
}

func (d *Delegation) AddAmount(amount *big.Int) {
	d.Amount = new(big.Int).Add(d.Amount, amount)
}

func (d *Delegation) SubAmount(amount *big.Int) {
	d.Amount = new(big.Int).Sub(d.Amount, amount)
}
