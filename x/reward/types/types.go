package types

import (
	"math/big"

	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
)

type EpochDelegationRewardPerShareItem struct {
	PreRewardEpoch  uint64
	NextRewardEpoch uint64
	TotalReward     *big.Int
	PerShareReward  *big.Int
}

func NewEpochDelegationRewardPerShareItem(preRewardEpoch, nextRewardEpoch uint64, totalReward, perShareReward *big.Int) *EpochDelegationRewardPerShareItem {
	return &EpochDelegationRewardPerShareItem{
		PreRewardEpoch:  preRewardEpoch,
		NextRewardEpoch: nextRewardEpoch,
		TotalReward:     totalReward,
		PerShareReward:  perShareReward,
	}
}

func (item *EpochDelegationRewardPerShareItem) IsEmpty() bool {
	return nil == item
}

func (item *EpochDelegationRewardPerShareItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

func (item *EpochDelegationRewardPerShareItem) IsZeroTotalReward() bool {
	return item.TotalReward.Cmp(basecommon.Big0) == 0
}

func (item *EpochDelegationRewardPerShareItem) DecrementTotalReward(amount *big.Int) {
	if item.TotalReward.Cmp(amount) < 0 {
		item.TotalReward = basecommon.Big0
	} else {
		item.TotalReward = new(big.Int).Sub(item.TotalReward, amount)
	}
}

func (item *EpochDelegationRewardPerShareItem) UpdatePreRewardEpoch(epoch uint64) {
	item.PreRewardEpoch = epoch
}

func (item *EpochDelegationRewardPerShareItem) UpdateNextRewardEpoch(epoch uint64) {
	item.NextRewardEpoch = epoch
}

type EpochDelegationRewardPerShareQueue []*EpochDelegationRewardPerShareItem

func NewEpochDelegationRewardPerShareQueue(size uint64) EpochDelegationRewardPerShareQueue {
	queue := make(EpochDelegationRewardPerShareQueue, size)
	return queue
}

func (queue EpochDelegationRewardPerShareQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue EpochDelegationRewardPerShareQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

// ------

type EpochDelegationRewardPerShareWithEpochItem struct {
	RewardEpoch uint64
	Data        *EpochDelegationRewardPerShareItem
}

func NewEpochDelegationRewardPerShareWithEpochItem(rewardEpoch uint64, data *EpochDelegationRewardPerShareItem) *EpochDelegationRewardPerShareWithEpochItem {
	return &EpochDelegationRewardPerShareWithEpochItem{
		RewardEpoch: rewardEpoch,
		Data:        data,
	}
}

func (item *EpochDelegationRewardPerShareWithEpochItem) IsEmpty() bool {
	return nil == item
}

func (item *EpochDelegationRewardPerShareWithEpochItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type EpochDelegationRewardPerShareWithEpochQueue []*EpochDelegationRewardPerShareWithEpochItem

func NewEpochDelegationRewardPerShareWithEpochQueue(size uint64) EpochDelegationRewardPerShareWithEpochQueue {
	queue := make(EpochDelegationRewardPerShareWithEpochQueue, size)
	return queue
}

func (queue EpochDelegationRewardPerShareWithEpochQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue EpochDelegationRewardPerShareWithEpochQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

// ------

type DelegationSnapshot struct {
	StakeEpoch     uint64   // stake epoch
	DelegateEpoch  uint64   // delegate epoch (Update the value when withdraw delegation rewards or when the delegated amount changes)
	PreEpochAmount *big.Int // total delegate amounts on pre epoch
	Amount         *big.Int
}

func NewDelegationSnapshot(stakeEpoch, delegateEpoch uint64, preEpochAmount, amount *big.Int) *DelegationSnapshot {
	return &DelegationSnapshot{
		StakeEpoch:     stakeEpoch,
		DelegateEpoch:  delegateEpoch,
		PreEpochAmount: preEpochAmount,
		Amount:         amount,
	}
}

func (d *DelegationSnapshot) IsEmpty() bool {
	return nil == d
}

func (d *DelegationSnapshot) IsNotEmpty() bool {
	return !d.IsEmpty()
}

type DelegationRewardSnapshot struct {
	Delegation  *DelegationSnapshot
	RewardQueue EpochDelegationRewardPerShareWithEpochQueue
}

func NewDelegationRewardSnapshot(delegation *DelegationSnapshot, rewardQueue EpochDelegationRewardPerShareWithEpochQueue) *DelegationRewardSnapshot {
	return &DelegationRewardSnapshot{
		Delegation:  delegation,
		RewardQueue: rewardQueue,
	}
}

func (d *DelegationRewardSnapshot) IsEmpty() bool {
	return nil == d
}

func (d *DelegationRewardSnapshot) IsNotEmpty() bool {
	return !d.IsEmpty()
}

type DelegationRewardSnapshotQueue []*DelegationRewardSnapshot

func NewDelegationRewardSnapshotQueue(size uint64) DelegationRewardSnapshotQueue {
	queue := make(DelegationRewardSnapshotQueue, size)
	return queue
}

func (queue DelegationRewardSnapshotQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue DelegationRewardSnapshotQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}
