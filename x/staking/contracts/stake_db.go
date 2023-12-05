package contracts

import (
	db "github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"math/big"
)

// ------------------------------------------------------ db methods ------------------------------------------------------

func (c *StakeHandler) incrementValidatorNonce() uint64 {
	return db.IncrementValidatorNonce(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) getValidatorNonce() uint64 {
	return db.GetValidatorNonce(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) setValidatorByPriority(validatorAddr common.Address, validator *types.Validator) error {
	if err := db.SetValidatorPriority(c.evm.StateDB, c.contract.Address(), validatorAddr, validator.Epoch, validator.StakeIndex, validator.Shares()); nil != err {
		return err
	}
	return db.SetValidator(c.evm.StateDB, c.contract.Address(), validatorAddr, validator)
}

func (c *StakeHandler) updateValidatorRemovePriority(validatorAddr common.Address, validator *types.Validator) error {
	old := c.getValidator(validatorAddr)
	if nil == old { // maybe short circuit
		return nil
	}
	// delete old priority
	if db.GetValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares()).ValidatorAddr != validatorAddr {
		return db.ErrMisMatching
	}
	if err := db.RemoveValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares()); nil != err {
		return err
	}
	// set new priority only

	return db.SetValidator(c.evm.StateDB, c.contract.Address(), validatorAddr, validator)
}

func (c *StakeHandler) updateValidatorByPriority(validatorAddr common.Address, validator *types.Validator) error {

	old := c.getValidator(validatorAddr)
	if nil == old { // maybe short circuit
		return nil
	}
	// delete old priority

	if db.GetValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares()).ValidatorAddr != validatorAddr {
		return db.ErrMisMatching
	}

	if err := db.RemoveValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares()); nil != err {
		return err
	}
	// set new priority and validator
	return c.setValidatorByPriority(validatorAddr, validator)
}

func (c *StakeHandler) getValidator(validatorAddr common.Address) *types.Validator {
	return db.GetValidator(c.evm.StateDB, c.contract.Address(), validatorAddr)
}

func (c *StakeHandler) hasValidator(validatorAddr common.Address) bool {
	return db.HasValidator(c.evm.StateDB, c.contract.Address(), validatorAddr)
}
func (c *StakeHandler) hasNotValidator(validatorAddr common.Address) bool {
	return db.HasNotValidator(c.evm.StateDB, c.contract.Address(), validatorAddr)
}

func (c *StakeHandler) removeValidator(validatorAddr common.Address) {
	db.RemoveValidator(c.evm.StateDB, c.contract.Address(), validatorAddr)
}

func (c *StakeHandler) setDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch uint64, delegation *types.Delegation) error {
	return db.SetDelegation(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, stakeEpoch, delegation)
}

func (c *StakeHandler) removeDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch uint64) {
	db.RemoveDelegation(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, stakeEpoch)
}

func (c *StakeHandler) getDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch uint64) *types.Delegation {
	return db.GetDelegation(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, stakeEpoch)
}

func (c *StakeHandler) incrementDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch, delegateEpoch uint64, amount *big.Int) (bool, error) {
	del := c.getDelegation(delegaterAddr, validatorAddr, stakeEpoch)

	var build bool

	if nil != del {
		del.UpdateEpoch(delegateEpoch)
		del.IncrementAmount(amount)
	} else {
		del = types.NewDelegation(delegateEpoch, amount)
		build = true
	}
	return build, c.setDelegation(delegaterAddr, validatorAddr, stakeEpoch, del)
}

func (c *StakeHandler) decrementDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch, delegateEpoch uint64, amount *big.Int) (bool, error) {
	del := c.getDelegation(delegaterAddr, validatorAddr, stakeEpoch)
	if nil == del {
		return false, db.ErrNotFound
	}
	del.UpdateEpoch(delegateEpoch)
	del.DecrementAmount(amount)

	var remove bool
	var err error
	if del.Amount.Cmp(common.Big0) == 0 {

		c.removeDelegation(delegaterAddr, validatorAddr, stakeEpoch)
		remove = true
	} else {
		err = c.setDelegation(delegaterAddr, validatorAddr, stakeEpoch, del)
	}
	return remove, err
}

func (c *StakeHandler) getCurrentEpoch() uint64 {
	return c.stage.GetCurrentEpoch(c.evm.StateDB)
}

func (c *StakeHandler) getCurrentRound() uint64 {
	return c.stage.GetCurrentRound(c.evm.StateDB)
}

func (c *StakeHandler) getEpochValidatorIds(epoch uint64) types.ValidatorIds {
	return db.GetEpochValidatorIds(c.evm.StateDB, c.contract.Address(), epoch)
}

func (c *StakeHandler) getRoundValidatorIds(round uint64) types.ValidatorIds {
	return db.GetRoundValidatorIds(c.evm.StateDB, c.contract.Address(), round)
}

func (c *StakeHandler) appendStakeWithdrawal(validatorAddr common.Address, epoch uint64, amount *big.Int) error {
	return db.AppendStakeWithdrawal(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch, amount)
}

func (c *StakeHandler) getStakeWithdrawalLastEpoch(validatorAddr common.Address) uint64 {
	return db.GetStakeWithdrawalLastEpoch(c.evm.StateDB, c.contract.Address(), validatorAddr)
}

// Total of all rewards until epoch
func (c *StakeHandler) getStakeWithdrawable(validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetStakeWithdrawable(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch)
}

func (c *StakeHandler) applyStakeWithdrawable(validatorAddr common.Address, epoch uint64) (*big.Int, error) {
	return db.ApplyStakeWithdrawable(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch)
}

func (c *StakeHandler) cleanStakeWithdrawable(validatorAddr common.Address) {
	db.CleanStakeWithdrawable(c.evm.StateDB, c.contract.Address(), validatorAddr)
}

// Total of all rewards since epoch
func (c *StakeHandler) getStakeWithdrawalPending(validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetStakeWithdrawalPending(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch)
}

func (c *StakeHandler) setStakeWithdrawalQueueItem(validatorAddr common.Address, epoch uint64, item *types.StakeWithdrawalItem) error {
	return db.SetStakeWithdrawalQueueItem(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch, item)
}

// -------------

func (c *StakeHandler) appendDelegateWithdrawal(delegaterAddr, validatorAddr common.Address, epoch uint64, amount *big.Int) error {
	return db.AppendDelegateWithdrawal(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, epoch, amount)
}

func (c *StakeHandler) getDelegateWithdrawalByEpoch(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetDelegateWithdrawalByEpoch(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, epoch)
}

// Total of all rewards until epoch
func (c *StakeHandler) getDelegateWithdrawable(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetDelegateWithdrawable(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, epoch)
}

func (c *StakeHandler) applyDelegateWithdrawable(delegaterAddr, validatorAddr common.Address, epoch uint64) (*big.Int, error) {
	return db.ApplyDelegateWithdrawable(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, epoch)
}

// Total of all rewards since epoch
func (c *StakeHandler) getDelegateWithdrawalPending(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetDelegateWithdrawalPending(c.evm.StateDB, c.contract.Address(), delegaterAddr, validatorAddr, epoch)
}

// -------

func (c *StakeHandler) appendValidatorDelegationRc(validatorAddr common.Address, epoch, rc uint64) error {
	return db.AppendValidatorDelegationRc(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch, rc)
}

func (c *StakeHandler) getValidatorDelegationRcPending(validatorAddr common.Address, size uint64) types.ValidatorDelegationRcQueue {
	return db.GetValidatorDelegationRcPending(c.evm.StateDB, c.contract.Address(), validatorAddr, size)
}

func (c *StakeHandler) getValidatorDelegationRcPendingAndEpoch(validatorAddr common.Address, size uint64) ([]uint64, types.ValidatorDelegationRcQueue) {
	return db.GetValidatorDelegationRcPendingAndEpoch(c.evm.StateDB, c.contract.Address(), validatorAddr, size)
}

func (c *StakeHandler) releaseValidatorDelegationRcItem(validatorAddr common.Address, stakeEpoch, decrement uint64) error {
	return db.ReleaseValidatorDelegationRcItem(c.evm.StateDB, c.contract.Address(), validatorAddr, stakeEpoch, decrement)
}

// ----

func (c *StakeHandler) setSlashProcessed(handleEventId *big.Int, queue types.SlashValidatorWithdrawItemQueue) error {
	return db.SetSlashProcessed(c.evm.StateDB, c.contract.Address(), handleEventId, queue)
}

func (c *StakeHandler) hasSlashProcessed(handleEventId *big.Int) bool {
	return db.HasSlashProcessed(c.evm.StateDB, c.contract.Address(), handleEventId)
}

func (c *StakeHandler) hasNotSlashProcessed(handleEventId *big.Int) bool {
	return db.HasNotSlashProcessed(c.evm.StateDB, c.contract.Address(), handleEventId)
}
