package contracts

import (
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/db"
	rewardtypes "github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
)

// internal

func (c *RewardManager) SetStageModule(stage rewardtypes.StageModuler) {
	c.stageModule = stage
}

func (c *RewardManager) SetStakeModule(stake rewardtypes.StakeModuler) {
	c.stakeModule = stake
}

func (c *RewardManager) SetRewardModule(reward rewardtypes.RewardModuler) {
	c.rewardModule = reward
}

func (c *RewardManager) updateDelegationRewards(delegaterAddr, validatorAddr basecommon.Address) error {
	return c.rewardModule.UpdateDelegationRewards(c.evm.StateDB, delegaterAddr, validatorAddr)
}

func (c *RewardManager) withdrawDelegationRewards(delegaterAddr, validatorAddr basecommon.Address) (*big.Int, error) {
	rewards := db.GetPendingDelegaterReward(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr)
	if rewards.Cmp(basecommon.Big0) != 0 {
		return basecommon.Big0, nil
	}

	rewardPoolBalance := c.evm.StateDB.GetBalance(c.contract.Address())
	if rewardPoolBalance.Cmp(rewards) < 0 {
		log.Error("insufficient balance on reward pool", "will withdraw delegate reward", rewards, "reward pool balance", rewardPoolBalance)
		return basecommon.Big0, typesdk.NewRevertError("RewardManager: insufficient balance on reward pool")
	}

	db.DecrementPendingDelegaterReward(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, rewards)
	c.evm.Context.Transfer(c.evm.StateDB, c.contract.Address(), delegaterAddr, rewards)

	return rewards, nil
}

func (c *RewardManager) withdrawValidatorRewrads(owner, validatorAddr basecommon.Address) (*big.Int, error) {
	rewards := db.GetPendingValidatorReward(c.evm.StateDB, c.contract.Address(), validatorAddr)

	if rewards.Cmp(basecommon.Big0) != 0 {
		return basecommon.Big0, nil
	}

	rewardPoolBalance := c.evm.StateDB.GetBalance(c.contract.Address())
	if rewardPoolBalance.Cmp(rewards) < 0 {
		log.Error("insufficient balance on reward pool", "will withdraw validator reward", rewards, "reward pool balance", rewardPoolBalance)
		return basecommon.Big0, typesdk.NewRevertError("RewardManager: insufficient balance on reward pool")
	}

	db.DecrementPendingValidatorReward(c.evm.StateDB, c.contract.Address(), validatorAddr, rewards)
	c.evm.Context.Transfer(c.evm.StateDB, c.contract.Address(), owner, rewards)

	return rewards, nil
}
