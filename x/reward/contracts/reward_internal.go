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

func (c *RewardManager) updateDelegationRewards(delegatorAddr, validatorAddr basecommon.Address) error {
	return c.rewardModule.UpdateDelegationRewards(c.evm.StateDB, delegatorAddr, validatorAddr)
}

func (c *RewardManager) withdrawDelegationRewards(delegatorAddr, validatorAddr basecommon.Address) (*big.Int, error) {

	rewards := db.GetPendingDelegatorReward(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr)

	// check rewards
	if rewards.Cmp(basecommon.Big0) == 0 {
		log.Error("has no delegation rewards", "delegatorAddr", delegatorAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
			"rewards", rewards, "currentEpoch", c.stageModule.GetCurrentEpoch(c.evm.StateDB), "blockNumber", c.evm.Context.BlockNumber)
		return basecommon.Big0, typesdk.NewRevertError("RewardManager: HAS NO DELEGATION REWARDS")
	}

	// check reward pool balance
	rewardPoolBalance := c.evm.StateDB.GetBalance(c.contract.Address())
	if rewardPoolBalance.Cmp(basecommon.Big0) == 0 || rewardPoolBalance.Cmp(rewards) < 0 {
		log.Error("insufficient balance on reward pool", "will withdraw delegate reward", rewards, "reward pool balance", rewardPoolBalance)
		return basecommon.Big0, typesdk.NewRevertError("RewardManager: insufficient balance on reward pool")
	}

	// decrement pending delegator rewards
	db.DecrementPendingDelegatorReward(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, rewards)

	// transfer rewards from reward pool to delegator
	c.evm.Context.Transfer(c.evm.StateDB, c.contract.Address(), delegatorAddr, rewards)

	return rewards, nil
}

func (c *RewardManager) withdrawValidatorRewrads(owner, validatorAddr basecommon.Address) (*big.Int, error) {

	rewards := db.GetPendingValidatorReward(c.evm.StateDB, c.contract.Address(), validatorAddr)

	// check rewards
	if rewards.Cmp(basecommon.Big0) == 0 {
		log.Error("has no validator rewards", "validatorAddr", validatorAddr.Hex(),
			"rewards", rewards, "currentEpoch", c.stageModule.GetCurrentEpoch(c.evm.StateDB), "blockNumber", c.evm.Context.BlockNumber)
		return basecommon.Big0, typesdk.NewRevertError("RewardManager: HAS NO VALIDATOR REWARDS")
	}

	// check reward pool balance
	rewardPoolBalance := c.evm.StateDB.GetBalance(c.contract.Address())
	if rewardPoolBalance.Cmp(basecommon.Big0) == 0 || rewardPoolBalance.Cmp(rewards) < 0 {
		log.Error("insufficient balance on reward pool", "will withdraw validator reward", rewards, "reward pool balance", rewardPoolBalance)
		return basecommon.Big0, typesdk.NewRevertError("RewardManager: insufficient balance on reward pool")
	}

	// decrement pending validator rewards
	db.DecrementPendingValidatorReward(c.evm.StateDB, c.contract.Address(), validatorAddr, rewards)

	// transfer rewards from reward pool to owner of validator
	c.evm.Context.Transfer(c.evm.StateDB, c.contract.Address(), owner, rewards)

	return rewards, nil
}
