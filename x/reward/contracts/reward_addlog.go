package contracts

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

func (c *RewardManager) addLogEpochRewardEvent(epochId *big.Int, validators []basecommon.Address, amounts []*big.Int) error {
	c.EmitEpochRewardEvent(epochId, validators, amounts)
	return nil
}

func (c *RewardManager) addLogRewardDistributedEvent(epochId *big.Int, totalReward *big.Int) error {
	c.EmitRewardDistributedEvent(epochId, totalReward)
	return nil
}

func (c *RewardManager) addLogEmitBlockRewardEvent(epochId *big.Int, validators []basecommon.Address, amounts []*big.Int) error {
	c.EmitBlockRewardEvent(epochId, validators, amounts)
	return nil
}

func (c *RewardManager) addLogEmitDelegatorRewardWithdrawalEvent(validator basecommon.Address, amount *big.Int, caller basecommon.Address) error {
	c.EmitDelegatorRewardWithdrawalEvent(validator, amount, caller)
	return nil
}

func (c *RewardManager) addLogEmitValidatorRewardWithdrawalEvent(validator basecommon.Address, amount *big.Int, caller basecommon.Address) error {
	c.EmitValidatorRewardWithdrawalEvent(validator, amount, caller)
	return nil
}
