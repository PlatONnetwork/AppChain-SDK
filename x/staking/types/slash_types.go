package types

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

type SlashValidatorWithdrawItem struct {
	ValidatorAddr common.Address
	Amount        *big.Int
}

func NewSlashValidatorWithdrawItem(validatorAddr common.Address, amount *big.Int) *SlashValidatorWithdrawItem {
	return &SlashValidatorWithdrawItem{
		ValidatorAddr: validatorAddr,
		Amount:        amount,
	}
}

type SlashValidatorWithdrawItemQueue []*SlashValidatorWithdrawItem

func (queue SlashValidatorWithdrawItemQueue) UnSafeAppend(item *SlashValidatorWithdrawItem) SlashValidatorWithdrawItemQueue {
	return append(queue, item)
}

func (queue SlashValidatorWithdrawItemQueue) Append(item *SlashValidatorWithdrawItem) SlashValidatorWithdrawItemQueue {
	return append(queue, item)
}

func (queue SlashValidatorWithdrawItemQueue) Has(item *SlashValidatorWithdrawItem) bool {
	for _, v := range queue {
		if v.ValidatorAddr == item.ValidatorAddr && v.Amount.Cmp(item.Amount) == 0 {
			return true
		}
	}
	return false
}

func (queue SlashValidatorWithdrawItemQueue) IndexOf(item *SlashValidatorWithdrawItem) int {
	for i, v := range queue {
		if v.ValidatorAddr == item.ValidatorAddr && v.Amount.Cmp(item.Amount) == 0 {
			return i
		}
	}
	return 0
}

func (queue SlashValidatorWithdrawItemQueue) Empty() bool {
	return len(queue) == 0
}

func (queue SlashValidatorWithdrawItemQueue) NotEmpty() bool {
	return !queue.Empty()
}


