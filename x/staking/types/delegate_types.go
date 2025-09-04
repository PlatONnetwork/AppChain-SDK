package types

import (
	"math/big"

	"github.com/PlatONnetwork/PlatON-Go/common"
)

type Delegation struct {
	Epoch          uint64 // delegate epoch
	PreEpochAmount *big.Int
	Amount         *big.Int
}

func NewDelegation(epoch uint64, amount *big.Int) *Delegation {
	return &Delegation{
		Epoch:          epoch,
		PreEpochAmount: common.Big0,
		Amount:         amount,
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

func (d *Delegation) SnapPreEpochAmount() {
	d.PreEpochAmount = new(big.Int).SetBytes(d.Amount.Bytes())
}

func (d *Delegation) IsEmpty() bool {
	return nil == d
}

func (d *Delegation) IsNotEmpty() bool {
	return !d.IsEmpty()
}
