package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

func (c *StakeHandler) addLogSlashedEvent(exitId *big.Int, validators []common.Address, amounts []*big.Int) error {
	c.EmitSlashedEvent(exitId, validators, amounts)
	return nil
}

func (c *StakeHandler) addLogDelegationAddedEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	c.EmitDelegationAddedEvent(delegator, validator, amount)
	return nil
}

func (c *StakeHandler) addLogStakeAddedEvent(validator common.Address, amount *big.Int) error {
	c.EmitStakeAddedEvent(validator, amount)
	return nil
}

func (c *StakeHandler) addLogUnDelegatedEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	c.EmitUnDelegatedEvent(delegator, validator, amount)
	return nil
}

func (c *StakeHandler) addLogUnStakedEvent(validator common.Address, amount *big.Int) error {
	c.EmitUnStakedEvent(validator, amount)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalEvent(validator common.Address, amount *big.Int) error {
	c.EmitStakeWithdrawalEvent(validator, amount)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalRegisteredEvent(validator common.Address, amount *big.Int) error {
	c.EmitStakeWithdrawalRegisteredEvent(validator, amount)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	c.EmitDelegateWithdrawalEvent(delegator, validator, amount)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalRegisteredEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	c.EmitDelegateWithdrawalRegisteredEvent(delegator, validator, amount)
	return nil
}

func (c *StakeHandler) addLogUpdateValidatorStatusEvent(validator common.Address, status *big.Int) error {
	c.EmitUpdateValidatorStatusEvent(validator, status)
	return nil
}

func (c *StakeHandler) addLogValidatorRegisteredEvent(validator common.Address, owner common.Address, commissionRate *big.Int, pubKey []byte, blsKey []byte) error {
	c.EmitValidatorRegisteredEvent(validator, owner, commissionRate, pubKey, blsKey)
	return nil
}
