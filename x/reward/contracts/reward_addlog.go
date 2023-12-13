package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	baselog "github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
)

func (c *RewardManager) addLogEpochRewardEvent(epochId *big.Int, validators []basecommon.Address, amounts []*big.Int) error {
	log, err := c.EmitEpochRewardEvent(epochId, validators, amounts)
	if nil != err {
		baselog.Error("Failed to emit EpochRewardEvent", "epochId", epochId, "validators size", len(validators), "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("RewardManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *RewardManager) addLogRewardDistributedEvent(epochId *big.Int, totalReward *big.Int) error {
	log, err := c.EmitRewardDistributedEvent(epochId, totalReward)
	if nil != err {
		baselog.Error("Failed to emit RewardDistributedEvent", "epochId", epochId, "totalReward", totalReward, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("RewardManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *RewardManager) addLogEmitBlockRewardEvent(epochId *big.Int, validators []basecommon.Address, amounts []*big.Int) error {
	log, err := c.EmitBlockRewardEvent(epochId, validators, amounts)
	if nil != err {
		baselog.Error("Failed to emit BlockRewardEvent", "epochId", epochId, "validators size", len(validators), "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("RewardManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *RewardManager) addLogEmitDelegatorRewardWithdrawalEvent(validator basecommon.Address, amount *big.Int, caller basecommon.Address) error {
	log, err := c.EmitDelegatorRewardWithdrawalEvent(validator, amount, caller)
	if nil != err {
		baselog.Error("Failed to emit DelegatorRewardWithdrawalEvent", "validator", validator.Hex(), "amount", amount, "caller", caller.Hex(), "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("RewardManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *RewardManager) addLogEmitValidatorRewardWithdrawalEvent(validator basecommon.Address, amount *big.Int, caller basecommon.Address) error {
	log, err := c.EmitValidatorRewardWithdrawalEvent(validator, amount, caller)
	if nil != err {
		baselog.Error("Failed to emit ValidatorRewardWithdrawalEvent", "validator", validator.Hex(), "amount", amount, "caller", caller.Hex(), "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("RewardManager: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
