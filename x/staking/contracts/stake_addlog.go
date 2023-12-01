package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

func (c *StakeHandler) addLogSlashedEvent(exitId *big.Int, validators []common.Address, amounts []*big.Int) error {
	log, err := c.EmitSlashedEvent(exitId, validators, amounts)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegationAddedEvent(delegater common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegationAddedEvent(delegater, validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeAddedEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitStakeAddedEvent(validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogUnDelegatedEvent(delegater common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitUnDelegatedEvent(delegater, validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogUnStakedEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitUnStakedEvent(validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitStakeWithdrawalEvent(validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalRegisteredEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitStakeWithdrawalRegisteredEvent(validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalEvent(delegater common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegateWithdrawalEvent(delegater, validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalRegisteredEvent(delegater common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegateWithdrawalRegisteredEvent(delegater, validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
