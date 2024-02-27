package config

import (
	"fmt"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
)

type RewardNetworkParams struct {
	module.ModuleGenesisConfig
	RewardPerBlock *big.Int `json:"rewardPerBlock"` // The validator receives rewards for each block builded
	RewardPerEpoch *big.Int `json:"rewardPerEpoch"` // The validator receives rewards based on 'stack shares' for each epoch
}

func DefaultRewardNetworkParams() *RewardNetworkParams {
	return &RewardNetworkParams{
		ModuleGenesisConfig: module.ModuleGenesisConfig{CreateBlock: 0},
		RewardPerBlock:      constants.REWARD_PER_BLOCK,
		RewardPerEpoch:      constants.REWARD_PER_EPOCH,
	}
}

func (params *RewardNetworkParams) String() string {
	return fmt.Sprintf(`{"createBlock": %d, "rewardPerBlock": %d,"rewardPerEpoch": %d}`, params.CreateBlock, params.RewardPerBlock, params.RewardPerEpoch)
}
