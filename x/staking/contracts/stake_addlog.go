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

func (c *StakeHandler) addLogUnDelegatedEvent(validator common.Address, delegater common.Address) error {
	log, err := c.EmitUnDelegatedEvent(validator, delegater)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogUnStakedEvent(validator common.Address) error {
	log, err := c.EmitUnStakedEvent(validator)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalEvent(account common.Address, amount *big.Int) error {
	log, err := c.EmitStakeWithdrawalEvent(account, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalRegisteredEvent(account common.Address, amount *big.Int) error {
	log, err := c.EmitStakeWithdrawalRegisteredEvent(account, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalEvent(account common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegateWithdrawalEvent(account, validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalRegisteredEvent(account common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegateWithdrawalRegisteredEvent(account, validator, amount)
	if nil != err {
		return nil
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
