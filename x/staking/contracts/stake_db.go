package contracts

import (
	"bytes"
	"math/big"

	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	db "github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
)

// ------------------------------------------------------ db methods ------------------------------------------------------

func (c *StakeHandler) incrementValidatorNonce() uint64 {
	return db.IncrementValidatorNonce(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) getValidatorNonce() uint64 {
	return db.GetValidatorNonce(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) getValidatorsByPriorityKey(start []byte, size uint64) ([]byte, []common.Address, []*types.Validator) {

	validatorAddrQueue := make([]common.Address, 0)
	validatorQueue := make([]*types.Validator, 0)

	if len(start) == 0 {
		headItem := db.GetValidatorPriorityByKey(c.evm.StateDB, c.contract.Address(), db.EncodePriorityValidatorHeadKey())
		if headItem.IsEmpty() {
			return nil, nil, nil
		}
		start = headItem.NextKey
	}
	var (
		count uint64
		index = start
	)
	item := db.GetValidatorPriorityByKey(c.evm.StateDB, c.contract.Address(), index)
	// it is not tail or not enough count
	for item.IsNotEmpty() && bytes.Compare(item.NextKey, db.EncodePriorityValidatorHeadKey()) != 0 && count < size {

		validator := c.getValidator(item.ValidatorAddr)
		if validator.IsEmpty() {

			index = item.NextKey
			item = db.GetValidatorPriorityByKey(c.evm.StateDB, c.contract.Address(), index)
			continue
		}

		validatorAddrQueue = append(validatorAddrQueue, item.ValidatorAddr)
		validatorQueue = append(validatorQueue, validator)
		index = item.NextKey
		item = db.GetValidatorPriorityByKey(c.evm.StateDB, c.contract.Address(), index)
		count++

	}
	return index, validatorAddrQueue, validatorQueue
}

func (c *StakeHandler) setValidatorByPriority(validatorAddr common.Address, validator *types.Validator) error {
	if err := db.SetValidator(c.evm.StateDB, c.contract.Address(), validatorAddr, validator); nil != err {
		return err
	}
	db.SetValidatorOwner(c.evm.StateDB, c.contract.Address(), validatorAddr, validator.Owner)
	return db.SetValidatorPriority(c.evm.StateDB, c.contract.Address(), validatorAddr, validator.Epoch, validator.StakeIndex, validator.Shares())
}

func (c *StakeHandler) updateValidatorRemovePriority(validatorAddr common.Address, validator *types.Validator) error {
	old := c.getValidator(validatorAddr)
	if old.IsEmpty() { // maybe short circuit
		return nil
	}
	// delete old priority
	priority := db.GetValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares())
	if priority.IsNotEmpty() {
		if priority.ValidatorAddr != validatorAddr {
			return db.ErrMisMatching
		}
		if err := db.RemoveValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares()); nil != err {
			return err
		}
	}
	// set new validator information only
	return db.SetValidator(c.evm.StateDB, c.contract.Address(), validatorAddr, validator)
}

func (c *StakeHandler) updateValidatorByPriority(validatorAddr common.Address, validator *types.Validator) error {

	old := c.getValidator(validatorAddr)
	if old.IsEmpty() { // maybe short circuit
		return nil
	}
	// delete old priority
	priority := db.GetValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares())
	if priority.IsNotEmpty() {
		if priority.ValidatorAddr != validatorAddr {
			return db.ErrMisMatching
		}
		if err := db.RemoveValidatorPriority(c.evm.StateDB, c.contract.Address(), old.Epoch, old.StakeIndex, old.Shares()); nil != err {
			return err
		}
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

func (c *StakeHandler) setDelegation(delegatorAddr, validatorAddr common.Address, stakeEpoch uint64, delegation *types.Delegation) error {
	return db.SetDelegation(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, stakeEpoch, delegation)
}

func (c *StakeHandler) removeDelegation(delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) {
	db.RemoveDelegation(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, stakeEpoch)
}

func (c *StakeHandler) getDelegation(delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) *types.Delegation {
	return db.GetDelegation(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, stakeEpoch)
}

func (c *StakeHandler) incrementDelegation(delegatorAddr, validatorAddr common.Address, stakeEpoch, delegateEpoch uint64, amount *big.Int) (bool, error) {
	del := c.getDelegation(delegatorAddr, validatorAddr, stakeEpoch)

	var build bool

	if nil != del {
		switch {
		case del.Epoch > delegateEpoch:
			return false, db.ErrInvalidValue
		case del.Epoch < delegateEpoch:
			del.UpdateEpoch(delegateEpoch)
			del.SnapPreEpochAmount()
		}
		del.IncrementAmount(amount)
	} else {
		del = types.NewDelegation(delegateEpoch, amount)
		build = true
	}
	return build, c.setDelegation(delegatorAddr, validatorAddr, stakeEpoch, del)
}

func (c *StakeHandler) decrementDelegation(delegatorAddr, validatorAddr common.Address, stakeEpoch, delegateEpoch uint64, amount *big.Int) (bool, error) {
	del := c.getDelegation(delegatorAddr, validatorAddr, stakeEpoch)
	if nil == del {
		return false, db.ErrNotFound
	}
	switch {
	case del.Epoch > delegateEpoch:
		return false, db.ErrInvalidValue
	case del.Epoch < delegateEpoch:
		del.UpdateEpoch(delegateEpoch)
		del.SnapPreEpochAmount()
	}
	del.DecrementAmount(amount)

	var remove bool
	var err error
	if del.Amount.Cmp(common.Big0) == 0 {

		c.removeDelegation(delegatorAddr, validatorAddr, stakeEpoch)
		remove = true
	} else {
		err = c.setDelegation(delegatorAddr, validatorAddr, stakeEpoch, del)
	}
	return remove, err
}

func (c *StakeHandler) getCurrentEpoch() uint64 {
	return c.stageModule.GetCurrentEpoch(c.evm.StateDB)
}

func (c *StakeHandler) getCurrentRound() uint64 {
	return c.stageModule.GetCurrentRound(c.evm.StateDB)
}

func (c *StakeHandler) getEpochValidatorIds(epoch uint64) types.ValidatorAddrQueue {
	return db.GetEpochValidatorIds(c.evm.StateDB, c.contract.Address(), epoch)
}

func (c *StakeHandler) getRoundValidatorIds(round uint64) types.ValidatorAddrQueue {
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

func (c *StakeHandler) appendDelegateWithdrawal(delegatorAddr, validatorAddr common.Address, epoch uint64, amount *big.Int) error {
	return db.AppendDelegateWithdrawal(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, epoch, amount)
}

func (c *StakeHandler) getDelegateWithdrawalByEpoch(delegatorAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetDelegateWithdrawalByEpoch(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, epoch)
}

// Total of all rewards until epoch
func (c *StakeHandler) getDelegateWithdrawable(delegatorAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetDelegateWithdrawable(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, epoch)
}

func (c *StakeHandler) applyDelegateWithdrawable(delegatorAddr, validatorAddr common.Address, epoch uint64) (*big.Int, error) {
	return db.ApplyDelegateWithdrawable(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, epoch)
}

// Total of all rewards since epoch
func (c *StakeHandler) getDelegateWithdrawalPending(delegatorAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	return db.GetDelegateWithdrawalPending(c.evm.StateDB, c.contract.Address(), delegatorAddr, validatorAddr, epoch)
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

// -------

func (c *StakeHandler) setSlashProcessed(exitEventId *big.Int, queue types.SlashValidatorWithdrawItemQueue) error {
	return db.SetSlashProcessed(c.evm.StateDB, c.contract.Address(), exitEventId, queue)
}

func (c *StakeHandler) hasSlashProcessed(exitEventId *big.Int) bool {
	return db.HasSlashProcessed(c.evm.StateDB, c.contract.Address(), exitEventId)
}

func (c *StakeHandler) hasNotSlashProcessed(exitEventId *big.Int) bool {
	return db.HasNotSlashProcessed(c.evm.StateDB, c.contract.Address(), exitEventId)
}

// ------

func (c *StakeHandler) getBlocksOfValidatorsForRound(round uint64) ([]BlocksOfValidator, error) {

	validatorAddrQueue := c.getRoundValidatorIds(round)
	queue := make([]BlocksOfValidator, len(validatorAddrQueue))
	for i := 0; i < len(validatorAddrQueue); i++ {
		validatorAddr := validatorAddrQueue[i]
		blocks := db.GetNumberOfBlocksForRoundValidator(c.evm.StateDB, c.contract.Address(), validatorAddr, round)

		queue[i] = BlocksOfValidator{
			ValidatorAddr: validatorAddr,
			Blocks:        new(big.Int).SetUint64(blocks),
		}
	}
	return queue, nil
}

func (c *StakeHandler) getBlocksOfValidatorsForEpoch(epoch uint64) ([]BlocksOfValidator, error) {

	epochStartBlockNumber, epochEndBlockNumber, roundCount := c.stageModule.GetEpochFlatten(c.evm.StateDB, epoch)

	firstRound := c.stageModule.GetRoundByBlockNumber(c.evm.StateDB, epochStartBlockNumber)
	lastRound := firstRound + roundCount - 1

	// check block number end edge
	_, lastRoundEndBlockNumber := c.stageModule.GetRoundFlatten(c.evm.StateDB, lastRound)
	if lastRoundEndBlockNumber != epochEndBlockNumber {
		return nil, typesdk.NewRevertError("StakeHandler: SYSTEM ERROR")
	}

	validatorAddrQueue := c.getEpochValidatorIds(epoch)
	queue := make([]BlocksOfValidator, len(validatorAddrQueue))

	for i := 0; i < len(validatorAddrQueue); i++ {

		validatorAddr := validatorAddrQueue[i]

		var accumulateBlocks uint64

		for roundIndex := firstRound; roundIndex <= lastRound; roundIndex++ {
			accumulateBlocks += db.GetNumberOfBlocksForRoundValidator(c.evm.StateDB, c.contract.Address(), validatorAddr, roundIndex)
		}

		queue[i] = BlocksOfValidator{
			ValidatorAddr: validatorAddr,
			Blocks:        new(big.Int).SetUint64(accumulateBlocks),
		}
	}
	return queue, nil
}
