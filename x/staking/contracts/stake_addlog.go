package contracts

import (
	"fmt"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	baselog "github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
)

func (c *StakeHandler) addLogSlashedEvent(exitId *big.Int, validators []common.Address, amounts []*big.Int) error {
	log, err := c.EmitSlashedEvent(exitId, validators, amounts)
	if nil != err {
		baselog.Error("Failed to emit SlashedEvent", "exitId", exitId, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegationAddedEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegationAddedEvent(delegator, validator, amount)
	if nil != err {
		baselog.Error("Failed to emit DelegationAddedEvent", "delegator", delegator.Hex(), "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeAddedEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitStakeAddedEvent(validator, amount)
	if nil != err {
		baselog.Error("Failed to emit StakeAddedEvent", "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogUnDelegatedEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitUnDelegatedEvent(delegator, validator, amount)
	if nil != err {
		baselog.Error("Failed to emit UnDelegatedEvent", "delegator", delegator.Hex(), "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogUnStakedEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitUnStakedEvent(validator, amount)
	if nil != err {
		baselog.Error("Failed to emit UnStakedEvent", "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitStakeWithdrawalEvent(validator, amount)
	if nil != err {
		baselog.Error("Failed to emit StakeWithdrawalEvent", "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogStakeWithdrawalRegisteredEvent(validator common.Address, amount *big.Int) error {
	log, err := c.EmitStakeWithdrawalRegisteredEvent(validator, amount)
	if nil != err {
		baselog.Error("Failed to emit StakeWithdrawalRegisteredEvent", "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegateWithdrawalEvent(delegator, validator, amount)
	if nil != err {
		baselog.Error("Failed to emit DelegateWithdrawalEvent", "delegator", delegator.Hex(), "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogDelegateWithdrawalRegisteredEvent(delegator common.Address, validator common.Address, amount *big.Int) error {
	log, err := c.EmitDelegateWithdrawalRegisteredEvent(delegator, validator, amount)
	if nil != err {
		baselog.Error("Failed to emit DelegateWithdrawalRegisteredEvent", "delegator", delegator.Hex(), "validator", validator.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}

func (c *StakeHandler) addLogUpdateValidatorStatusEvent(validator common.Address, status *big.Int) error {
	log, err := c.EmitUpdateValidatorStatusEvent(validator, status)
	if nil != err {
		baselog.Error("Failed to emit UpdateValidatorStatusEvent", "validator", validator.Hex(), "status", status.Uint64(), "error", err)
		return typesdk.NewRevertError(fmt.Sprintf("StakeHandler: %s", err))
	}
	c.evm.StateDB.AddLog(log)
	return nil
}
