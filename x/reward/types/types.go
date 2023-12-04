package types

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

type EpochRewardPerDelegationShareItem struct {
	TotalReward    *big.Int
	PerShareReward *big.Int
}

func NewEpochRewardPerDelegationShareItem(totalReward, perShareReward *big.Int) *EpochRewardPerDelegationShareItem {
	return &EpochRewardPerDelegationShareItem{
		TotalReward:    totalReward,
		PerShareReward: perShareReward,
	}
}

func (item *EpochRewardPerDelegationShareItem) IsEmpty() bool {
	return nil == item
}

func (item *EpochRewardPerDelegationShareItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

func (item *EpochRewardPerDelegationShareItem) DecrementTotalReward(amount *big.Int) {
	if item.TotalReward.Cmp(amount) < 0 {
		item.TotalReward = basecommon.Big0
	} else {
		item.TotalReward = new(big.Int).Sub(item.TotalReward, amount)
	}
}
