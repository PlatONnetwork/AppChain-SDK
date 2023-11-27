package types

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

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

func (d *Delegation) IncrementAmount(amount *big.Int) {
	d.Amount = new(big.Int).Add(d.Amount, amount)
}

func (d *Delegation) DecrementAmount(amount *big.Int) {
	if d.Amount.Cmp(amount) < 0 {
		d.Amount = common.Big0
	} else {
		d.Amount = new(big.Int).Sub(d.Amount, amount)
	}
}

func (d *Delegation) IsEmpty() bool {
	return nil == d
}

func (d *Delegation) IsNotEmpty() bool {
	return !d.IsEmpty()
}
