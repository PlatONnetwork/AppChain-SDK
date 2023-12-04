package types

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

type EpochDelegationRewardPerShareItem struct {
	TotalReward    *big.Int
	PerShareReward *big.Int
}

func NewEpochDelegationRewardPerShareItem(totalReward, perShareReward *big.Int) *EpochDelegationRewardPerShareItem {
	return &EpochDelegationRewardPerShareItem{
		TotalReward:    totalReward,
		PerShareReward: perShareReward,
	}
}

func (item *EpochDelegationRewardPerShareItem) IsEmpty() bool {
	return nil == item
}

func (item *EpochDelegationRewardPerShareItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

func (item *EpochDelegationRewardPerShareItem) DecrementTotalReward(amount *big.Int) {
	if item.TotalReward.Cmp(amount) < 0 {
		item.TotalReward = basecommon.Big0
	} else {
		item.TotalReward = new(big.Int).Sub(item.TotalReward, amount)
	}
}
