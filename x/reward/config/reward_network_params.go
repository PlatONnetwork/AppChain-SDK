package config

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"math/big"
)

type RewardNetworkParams struct {
	RewardPerBlock *big.Int `json:"rewardPerBlock"` // The validator receives rewards for each block builded
	RewardPerEpoch *big.Int `json:"rewardPerEpoch"` // The validator receives rewards based on 'stack shares' for each epoch
}

func DefaultRewardNetworkParams() *RewardNetworkParams {
	return &RewardNetworkParams{
		RewardPerBlock: constants.REWARD_PER_BLOCK,
		RewardPerEpoch: constants.REWARD_PER_EPOCH,
	}
}
