package contracts

import (
	db "github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/big"
)

var (
	//validatorDelegaterCountKeyPrefix    = []byte("validatorDelegaterCount")
	validatorNonceKey = []byte("validatorNonce") // "validatorNonce" => nonce (It is a self increasing stake index number)
	//validatorKeyPrefix         = []byte("validator")             // "validator":validatorAddr => validator
	delegationKeyPrefix = []byte("delegation") // "delegater":delegaterAddr:validatorAddr:stakeEpoch => delegation

	stakeWithdrawalQueueItemKeyPrefix = []byte("stakeWithdrawalQueueItem") // "stakeWithdrawalQueueItem":validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}

	delegateWithdrawalQueueItemKeyPrefix = []byte("delegateWithdrawalQueueItem") // "delegateWithdrawalQueueItem":delegaterAddr:validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}

	validatorDelegationRcKeyPrefix = []byte("validatorDelegationRc") // "validatorDelegationRc":validatorAddr:stakeEpoch => unStakeDelegationRcItem{preStakeEpoch, nextStakeEpoch, delegation count}

	slashProcessedKeyPrefix = []byte("slashProcessed") // "slashProcessed":handleEventId => []SlashValidatorWithdrawItem{validatorAddr, amount}
)

//func encodeValidatorKey(validatorAddr common.Address) []byte {
//	return append(validatorKeyPrefix, validatorAddr.Bytes()...)
//}

func encodeDelegaterKey(delegaterAddr, validatorAddr common.Address, stakeEpoch uint64) []byte {
	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()
	stakeEpochBytes := common.Uint64ToBytes(stakeEpoch)

	keyPrefixSize := len(delegationKeyPrefix)
	appendDelegaterSize := keyPrefixSize + len(delegaterAddrBytes)
	appendVlidatorAddrSize := appendDelegaterSize + len(validatorAddrBytes)
	size := appendVlidatorAddrSize + len(stakeEpochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegationKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterSize], delegaterAddrBytes)
	copy(key[appendDelegaterSize:appendVlidatorAddrSize], validatorAddrBytes)
	copy(key[appendVlidatorAddrSize:], stakeEpochBytes)

	return key
}

func encodeStakeWithdrawalQueueItemKey(validatorAddr common.Address, unlockEpoch uint64) []byte {

	validatorAddrBytes := validatorAddr.Bytes()
	unlockEpochBytes := common.Uint64ToBytes(unlockEpoch)

	keyPrefixSize := len(stakeWithdrawalQueueItemKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(unlockEpochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], stakeWithdrawalQueueItemKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], unlockEpochBytes)

	return key
}

func encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr common.Address, unlockEpoch uint64) []byte {

	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()
	unlockEpochBytes := common.Uint64ToBytes(unlockEpoch)

	keyPrefixSize := len(delegateWithdrawalQueueItemKeyPrefix)
	appendDelegaterSize := keyPrefixSize + len(delegaterAddrBytes)
	appendValidatorAddrSize := appendDelegaterSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(unlockEpochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegateWithdrawalQueueItemKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterSize], delegaterAddrBytes)
	copy(key[appendDelegaterSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], unlockEpochBytes)

	return key
}

func encodeValidatorDelegationRcKey(validatorAddr common.Address, stakeEpoch uint64) []byte {
	validatorAddrBytes := validatorAddr.Bytes()
	stakeEpochBytes := common.Uint64ToBytes(stakeEpoch)

	keyPrefixSize := len(validatorDelegationRcKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(stakeEpochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], validatorDelegationRcKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], stakeEpochBytes)

	return key
}

func encodeSlashProcessedKey(handleEventId *big.Int) []byte {
	return append(slashProcessedKeyPrefix, handleEventId.Bytes()...)
}

// ------------------------------------------------------ db methods ------------------------------------------------------

func (c *StakeHandler) incrementValidatorNonce() uint64 {
	value := c.evm.StateDB.GetState(c.contract.Address(), validatorNonceKey)
	var v uint64
	if len(value) != 0 {
		v = common.BytesToUint64(value)
	}
	old := v
	v++
	c.evm.StateDB.SetState(c.contract.Address(), validatorNonceKey, common.Uint64ToBytes(v))
	return old
}

func (c *StakeHandler) GetValidatorNonce() uint64 {
	value := c.evm.StateDB.GetState(c.contract.Address(), validatorNonceKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func (c *StakeHandler) setValidatorByPriority(validatorAddr common.Address, validator *types.Validator) error {
	if err := c.setValidatorPriority(validatorAddr, validator.Epoch, validator.StakeIndex, validator.Shares()); nil != err {
		return err
	}
	return c.setValidator(validatorAddr, validator)
}

func (c *StakeHandler) updateValidatorRemovePriority(validatorAddr common.Address, validator *types.Validator) error {
	old := c.GetValidator(validatorAddr)
	if nil == old { // maybe short circuit
		return nil
	}
	// delete old priority
	if c.getValidatorPriority(old.Epoch, old.StakeIndex, old.Shares()).ValidatorAddr != validatorAddr {
		return db.ErrMisMatching
	}
	if err := c.removeValidatorPriority(old.Epoch, old.StakeIndex, old.Shares()); nil != err {
		return err
	}
	// set new priority only
	return c.setValidator(validatorAddr, validator)
}

func (c *StakeHandler) updateValidatorByPriority(validatorAddr common.Address, validator *types.Validator) error {

	old := c.GetValidator(validatorAddr)
	if nil == old { // maybe short circuit
		return nil
	}
	// delete old priority
	if c.getValidatorPriority(old.Epoch, old.StakeIndex, old.Shares()).ValidatorAddr != validatorAddr {
		return db.ErrMisMatching
	}
	if err := c.removeValidatorPriority(old.Epoch, old.StakeIndex, old.Shares()); nil != err {
		return err
	}
	// set new priority and validator
	return c.setValidatorByPriority(validatorAddr, validator)
}

func (c *StakeHandler) getValidatorPriority(epoch, stakeIndex uint64, shares *big.Int) *types.PriorityValidator {
	return db.GetValidatorPriority(c.evm.StateDB, c.contract.Address(), epoch, stakeIndex, shares)
}

func (c *StakeHandler) setValidatorPriority(validatorAddr common.Address, epoch, stakeIndex uint64, shares *big.Int) error {
	return db.SetValidatorPriority(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch, stakeIndex, shares)
}

func (c *StakeHandler) removeValidatorPriority(epoch, stakeIndex uint64, shares *big.Int) error {
	return db.RemoveValidatorPriority(c.evm.StateDB, c.contract.Address(), epoch, stakeIndex, shares)
}

func (c *StakeHandler) setValidator(validatorAddr common.Address, validator *types.Validator) error {
	return db.SetValidator(c.evm.StateDB, c.contract.Address(), validatorAddr, validator)
}

func (c *StakeHandler) GetValidator(validatorAddr common.Address) *types.Validator {
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
	value, err := rlp.EncodeToBytes(delegation)
	if nil != err {
		return db.ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegaterKey(delegaterAddr, validatorAddr, stakeEpoch), value)
	return nil
}

func (c *StakeHandler) removeDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegaterKey(delegaterAddr, validatorAddr, stakeEpoch), []byte{})
}

func (c *StakeHandler) GetDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch uint64) *types.Delegation {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeDelegaterKey(delegaterAddr, validatorAddr, stakeEpoch))
	if len(value) == 0 {
		return nil
	}
	var delegation types.Delegation
	if err := rlp.DecodeBytes(value, &delegation); nil == err {
		return &delegation
	}
	return nil
}

func (c *StakeHandler) incrementDelegation(delegaterAddr, validatorAddr common.Address, stakeEpoch, delegateEpoch uint64, amount *big.Int) (bool, error) {
	del := c.GetDelegation(delegaterAddr, validatorAddr, stakeEpoch)

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
	del := c.GetDelegation(delegaterAddr, validatorAddr, stakeEpoch)
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
	return db.GetCurrentEpoch(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) getCurrentRound() uint64 {
	return db.GetCurrentRound(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) getEpochValidatorIds(epoch uint64) types.ValidatorIds {
	return db.GetEpochValidatorIds(c.evm.StateDB, c.contract.Address(), epoch)
}

func (c *StakeHandler) getRoundValidatorIds(round uint64) types.ValidatorIds {
	return db.GetRoundValidatorIds(c.evm.StateDB, c.contract.Address(), round)
}

func (c *StakeHandler) appendStakeWithdrawal(validatorAddr common.Address, epoch uint64, amount *big.Int) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// first insert
	if indexItem.IsEmpty() {

		preEpoch := uint64(0)
		// tail -> head -> epoch -> tail -> head

		preItem := types.NewStakeWithdrawalItem(indexEpoch, epoch, common.Big0) // head
		epochItem := types.NewStakeWithdrawalItem(preEpoch, indexEpoch, amount) // item
		indexItem = types.NewStakeWithdrawalItem(epoch, preEpoch, common.Big0)  // tail

		if err := c.setStakeWithdrawalQueueItem(validatorAddr, preEpoch, preItem); nil != err {
			return err
		}
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexEpoch, indexItem); nil != err {
			return err
		}

		return nil
	}

	// range from tail to head
	for indexItem.PreEpoch != math.MaxUint64 { // not as head
		if indexEpoch == epoch {
			// pre -> index(epoch) -> next
			// index == epoch

			indexItem.IncrementAmount(amount)
			if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			break
		} else if indexEpoch < epoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			next := c.getStakeWithdrawalQueueItem(validatorAddr, indexItem.NextEpoch)
			epochItem := types.NewStakeWithdrawalItem(indexEpoch, indexItem.NextEpoch, amount)

			indexItem.UpdateNextEpoch(epoch)
			next.UpdatePreEpoch(epoch)

			if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := c.setStakeWithdrawalQueueItem(validatorAddr, epoch, epochItem); nil != err {
				return err
			}
			if err := c.setStakeWithdrawalQueueItem(validatorAddr, epochItem.NextEpoch, next); nil != err {
				return err
			}

			break

		} else {

			if indexItem.PreEpoch == uint64(0) {
				// if  head(min) -> index -> tail(max)
				// and epoch < index
				//
				// then: head(min) -> epoch -> index<last one> -> ... -> tail(max)
				// head < epoch < index < ... < tail

				pre := c.getStakeWithdrawalQueueItem(validatorAddr, indexItem.PreEpoch) // head
				epochItem := types.NewStakeWithdrawalItem(indexItem.PreEpoch, epoch, amount)

				pre.UpdateNextEpoch(epoch)
				indexItem.UpdatePreEpoch(epoch)

				if err := c.setStakeWithdrawalQueueItem(validatorAddr, epochItem.PreEpoch, pre); nil != err {
					return err
				}
				if err := c.setStakeWithdrawalQueueItem(validatorAddr, epoch, epochItem); nil != err {
					return err
				}
				if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexEpoch, indexItem); nil != err {
					return err
				}

				break

			}

			// if  head -> ... -> pre -> index -> ... -> max
			// (head < ... <  pre < index < ... < tail(max) )
			// and epoch < index
			// maybe epoch < pre
			//
			// then: continue pre become new index
			//

		}

		indexEpoch = indexItem.PreEpoch
		indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
	}
	return nil
}

func (c *StakeHandler) GetStakeWithdrawal(validatorAddr common.Address) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	for indexItem.NextEpoch != uint64(0) { // not as tail
		amount = new(big.Int).Add(amount, indexItem.Amount)
		indexEpoch = indexItem.NextEpoch
		indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
	}
	return amount
}

func (c *StakeHandler) GetStakeWithdrawalByEpoch(validatorAddr common.Address, epoch uint64) *big.Int {
	item := c.getStakeWithdrawalQueueItem(validatorAddr, epoch)
	if nil != item {
		return item.Amount
	}
	return common.Big0
}

func (c *StakeHandler) getStakeWithdrawalLastEpoch(validatorAddr common.Address) uint64 {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return 0
	}

	return indexItem.PreEpoch
}

// Total of all rewards until epoch
func (c *StakeHandler) GetStakeWithdrawable(validatorAddr common.Address, epoch uint64) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return amount
	}

	for indexItem.NextEpoch != uint64(0) {

		if indexEpoch > epoch {
			break
		}
		amount = new(big.Int).Add(amount, indexItem.Amount)
		indexEpoch = indexItem.NextEpoch
		indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
	}
	return amount
}

func (c *StakeHandler) applyStakeWithdrawable(validatorAddr common.Address, epoch uint64) (*big.Int, error) {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return amount, nil
	}
	for indexItem.NextEpoch != uint64(0) { // not sa tail

		if indexEpoch > epoch {
			break
		}
		amount = new(big.Int).Add(amount, indexItem.Amount)

		// remove item
		if indexEpoch != uint64(0) { // not as haed
			c.removeStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
		}

		indexEpoch = indexItem.NextEpoch
		indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
	}

	// remove head and tail
	if indexItem.NextEpoch == uint64(0) { // as tail
		c.removeStakeWithdrawalQueueItem(validatorAddr, 0)
		c.removeStakeWithdrawalQueueItem(validatorAddr, math.MaxUint64)
	} else { // update start index
		head := c.getStakeWithdrawalQueueItem(validatorAddr, uint64(0))
		head.UpdateNextEpoch(indexEpoch)
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, uint64(0), head); nil != err { // update head
			return common.Big0, err
		}
	}

	return amount, nil
}

func (c *StakeHandler) cleanStakeWithdrawable(validatorAddr common.Address) {

	indexEpoch := uint64(0)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return
	}
	for indexItem.IsNotEmpty() {

		c.removeStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

		indexEpoch = indexItem.NextEpoch
		indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
	}

	return
}

// Total of all rewards since epoch
func (c *StakeHandler) GetStakeWithdrawalPending(validatorAddr common.Address, epoch uint64) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(math.MaxUint64)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return amount
	}

	for indexItem.PreEpoch != math.MaxUint64 {

		if indexEpoch <= epoch {
			break
		}
		amount = new(big.Int).Add(amount, indexItem.Amount)

		indexEpoch = indexItem.PreEpoch
		indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
	}
	return amount
}

func (c *StakeHandler) setStakeWithdrawalQueueItem(validatorAddr common.Address, epoch uint64, item *types.StakeWithdrawalItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return db.ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeStakeWithdrawalQueueItemKey(validatorAddr, epoch), value)
	return nil
}

func (c *StakeHandler) removeStakeWithdrawalQueueItem(validatorAddr common.Address, epoch uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), encodeStakeWithdrawalQueueItemKey(validatorAddr, epoch), []byte{})
}

func (c *StakeHandler) getStakeWithdrawalQueueItem(validatorAddr common.Address, epoch uint64) *types.StakeWithdrawalItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeStakeWithdrawalQueueItemKey(validatorAddr, epoch))

	if len(value) == 0 {
		return nil
	}
	var item types.StakeWithdrawalItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

// -------------

func (c *StakeHandler) appendDelegateWithdrawal(delegaterAddr, validatorAddr common.Address, epoch uint64, amount *big.Int) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)

	// first insert
	if indexItem.IsEmpty() {

		preEpoch := uint64(0)
		// tail -> head -> epoch -> tail -> head

		preItem := types.NewDelegateWithdrawalItem(indexEpoch, epoch, common.Big0) // head
		epochItem := types.NewDelegateWithdrawalItem(preEpoch, indexEpoch, amount) // item
		indexItem = types.NewDelegateWithdrawalItem(epoch, preEpoch, common.Big0)  // tail

		if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, preEpoch, preItem); nil != err {
			return err
		}
		if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
			return err
		}

		return nil
	}

	// range from tail to head
	for indexItem.PreEpoch != math.MaxUint64 { // not as head
		if indexEpoch == epoch {
			// pre -> index(epoch) -> next
			// index == epoch

			indexItem.IncrementAmount(amount)
			if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			break
		} else if indexEpoch < epoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			next := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexItem.NextEpoch)
			epochItem := types.NewDelegateWithdrawalItem(indexEpoch, indexItem.NextEpoch, amount)

			indexItem.UpdateNextEpoch(epoch)
			next.UpdatePreEpoch(epoch)

			if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
				return err
			}
			if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epochItem.NextEpoch, next); nil != err {
				return err
			}

			break

		} else {

			if indexItem.PreEpoch == uint64(0) {
				// if  head(min) -> index -> tail(max)
				// and epoch < index
				//
				// then: head(min) -> epoch -> index<last one> -> ... -> tail(max)
				// head < epoch < index < ... < tail

				pre := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexItem.PreEpoch) // head
				epochItem := types.NewDelegateWithdrawalItem(indexItem.PreEpoch, epoch, amount)

				pre.UpdateNextEpoch(epoch)
				indexItem.UpdatePreEpoch(epoch)

				if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epochItem.PreEpoch, pre); nil != err {
					return err
				}
				if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
					return err
				}
				if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
					return err
				}

				break

			}

			// if  head -> ... -> pre -> index -> ... -> max
			// (head < ... <  pre < index < ... < tail(max) )
			// and epoch < index
			// maybe epoch < pre
			//
			// then: continue pre become new index
			//

		}

		indexEpoch = indexItem.PreEpoch
		indexItem = c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)
	}
	return nil
}

func (c *StakeHandler) GetDelegateWithdrawal(delegaterAddr, validatorAddr common.Address) *big.Int {
	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)

	for indexItem.NextEpoch != uint64(0) { // not as tail
		amount = new(big.Int).Add(amount, indexItem.Amount)
		indexEpoch = indexItem.NextEpoch
		indexItem = c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)
	}
	return amount
}

func (c *StakeHandler) GetDelegateWithdrawalByEpoch(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	item := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epoch)
	if nil != item {
		return item.Amount
	}
	return common.Big0
}

// Total of all rewards until epoch
func (c *StakeHandler) GetDelegateWithdrawable(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return amount
	}

	for indexItem.NextEpoch != uint64(0) {

		if indexEpoch > epoch {
			break
		}
		amount = new(big.Int).Add(amount, indexItem.Amount)
		indexEpoch = indexItem.NextEpoch
		indexItem = c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)
	}
	return amount
}

func (c *StakeHandler) applyDelegateWithdrawable(delegaterAddr, validatorAddr common.Address, epoch uint64) (*big.Int, error) {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return amount, nil
	}
	for indexItem.NextEpoch != uint64(0) {

		if indexEpoch > epoch {
			break
		}
		amount = new(big.Int).Add(amount, indexItem.Amount)

		// remove item
		if indexEpoch != uint64(0) {
			c.removeDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)
		}

		indexEpoch = indexItem.NextEpoch
		indexItem = c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)
	}

	// remove head and tail
	if indexItem.NextEpoch == uint64(0) {
		c.removeDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, 0)
		c.removeDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, math.MaxUint64)
	} else { // update start index
		head := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, uint64(0))
		head.UpdateNextEpoch(indexEpoch)
		if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, uint64(0), head); nil != err {
			return common.Big0, err
		}
	}

	return amount, nil
}

// Total of all rewards since epoch
func (c *StakeHandler) GetDelegateWithdrawalPending(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	amount := common.Big0

	indexEpoch := uint64(math.MaxUint64)
	indexItem := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return amount
	}

	for indexItem.PreEpoch != math.MaxUint64 {

		if indexEpoch <= epoch {
			break
		}
		amount = new(big.Int).Add(amount, indexItem.Amount)

		indexEpoch = indexItem.PreEpoch
		indexItem = c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)
	}
	return amount
}

func (c *StakeHandler) setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr common.Address, epoch uint64, item *types.DelegateWithdrawalItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return db.ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch), value)
	return nil
}

func (c *StakeHandler) removeDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr common.Address, epoch uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch), []byte{})
}

func (c *StakeHandler) getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr common.Address, epoch uint64) *types.DelegateWithdrawalItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch))

	if len(value) == 0 {
		return nil
	}
	var item types.DelegateWithdrawalItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

// -------

func (c *StakeHandler) appendValidatorDelegationRc(validatorAddr common.Address, epoch, rc uint64) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := c.getValidatorDelegationRcItem(validatorAddr, indexEpoch)

	// first insert
	if indexItem.IsEmpty() {

		preEpoch := uint64(0)
		// tail -> head -> epoch -> tail -> head

		preItem := types.NewValidatorDelegationRcItem(indexEpoch, epoch, 0)       // head
		epochItem := types.NewValidatorDelegationRcItem(preEpoch, indexEpoch, rc) // item
		indexItem = types.NewValidatorDelegationRcItem(epoch, preEpoch, 0)        // tail

		if err := c.setValidatorDelegationRcItem(validatorAddr, preEpoch, preItem); nil != err {
			return err
		}
		if err := c.setValidatorDelegationRcItem(validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := c.setValidatorDelegationRcItem(validatorAddr, indexEpoch, indexItem); nil != err {
			return err
		}

		return nil
	}

	// range from tail to head
	for indexItem.PreStakeEpoch != math.MaxUint64 { // not as head
		if indexEpoch == epoch {
			// pre -> index(epoch) -> next
			// index == epoch

			indexItem.IncrementRc(rc)
			if err := c.setValidatorDelegationRcItem(validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			break
		} else if indexEpoch < epoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			next := c.getValidatorDelegationRcItem(validatorAddr, indexItem.NextStakeEpoch)
			epochItem := types.NewValidatorDelegationRcItem(indexEpoch, indexItem.NextStakeEpoch, rc)

			indexItem.UpdateNextStakeEpoch(epoch)
			next.UpdatePreStakeEpoch(epoch)

			if err := c.setValidatorDelegationRcItem(validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := c.setValidatorDelegationRcItem(validatorAddr, epoch, epochItem); nil != err {
				return err
			}
			if err := c.setValidatorDelegationRcItem(validatorAddr, epochItem.NextStakeEpoch, next); nil != err {
				return err
			}

			break

		} else {

			if indexItem.PreStakeEpoch == uint64(0) {
				// if  head(min) -> index -> tail(max)
				// and epoch < index
				//
				// then: head(min) -> epoch -> index<last one> -> ... -> tail(max)
				// head < epoch < index < ... < tail

				pre := c.getValidatorDelegationRcItem(validatorAddr, indexItem.PreStakeEpoch) // head
				epochItem := types.NewValidatorDelegationRcItem(indexItem.PreStakeEpoch, epoch, rc)

				pre.UpdateNextStakeEpoch(epoch)
				indexItem.UpdatePreStakeEpoch(epoch)

				if err := c.setValidatorDelegationRcItem(validatorAddr, epochItem.PreStakeEpoch, pre); nil != err {
					return err
				}
				if err := c.setValidatorDelegationRcItem(validatorAddr, epoch, epochItem); nil != err {
					return err
				}
				if err := c.setValidatorDelegationRcItem(validatorAddr, indexEpoch, indexItem); nil != err {
					return err
				}

				break

			}

			// if  head -> ... -> pre -> index -> ... -> max
			// (head < ... <  pre < index < ... < tail(max) )
			// and epoch < index
			// maybe epoch < pre
			//
			// then: continue pre become new index
			//

		}

		indexEpoch = indexItem.PreStakeEpoch
		indexItem = c.getValidatorDelegationRcItem(validatorAddr, indexEpoch)
	}
	return nil
}

func (c *StakeHandler) getValidatorDelegationRcPending(validatorAddr common.Address, size uint64) types.ValidatorDelegationRcQueue {
	indexStakeEpoch := uint64(0)
	indexItem := c.getValidatorDelegationRcItem(validatorAddr, indexStakeEpoch)

	if indexItem.IsEmpty() {
		return nil
	}
	queue := types.NewValidatorDelegationRcQueue(size)
	count := uint64(0)
	for indexItem.NextStakeEpoch != uint64(0) && count < size { // not tail or count less size
		queue[count] = indexItem

		indexStakeEpoch = indexItem.NextStakeEpoch
		indexItem = c.getValidatorDelegationRcItem(validatorAddr, indexStakeEpoch)
		count++
	}
	return queue[:count]
}

func (c *StakeHandler) getValidatorDelegationRcItem(validatorAddr common.Address, stakeEpoch uint64) *types.ValidatorDelegationRcItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeValidatorDelegationRcKey(validatorAddr, stakeEpoch))

	if len(value) != 0 {
		return nil
	}
	var item types.ValidatorDelegationRcItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

func (c *StakeHandler) setValidatorDelegationRcItem(validatorAddr common.Address, stakeEpoch uint64, item *types.ValidatorDelegationRcItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return db.ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeValidatorDelegationRcKey(validatorAddr, stakeEpoch), value)
	return nil
}

func (c *StakeHandler) removeValidatorDelegationRcItem(validatorAddr common.Address, stakeEpoch uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), encodeValidatorDelegationRcKey(validatorAddr, stakeEpoch), []byte{})
}

func (c *StakeHandler) incrementValidatorDelegationRcItem(validatorAddr common.Address, stakeEpoch, increment uint64) error {
	item := c.getValidatorDelegationRcItem(validatorAddr, stakeEpoch)
	if nil == item {
		return db.ErrNotFound
	}
	item.IncrementRc(increment)
	return c.setValidatorDelegationRcItem(validatorAddr, stakeEpoch, item)
}

func (c *StakeHandler) decrementValidatorDelegationRcItem(validatorAddr common.Address, stakeEpoch, decrement uint64) error {
	item := c.getValidatorDelegationRcItem(validatorAddr, stakeEpoch)
	if nil == item {
		return db.ErrNotFound
	}

	item.DecrementRc(decrement)

	// change index
	if item.Rc == 0 {
		pre := c.getValidatorDelegationRcItem(validatorAddr, item.PreStakeEpoch)
		next := c.getValidatorDelegationRcItem(validatorAddr, item.NextStakeEpoch)

		// remove the last one   tail -> head -> lastone(remove) -> tail -> head
		if pre.PreStakeEpoch == math.MaxUint64 && next.NextStakeEpoch == 0 {
			c.removeValidatorDelegationRcItem(validatorAddr, item.PreStakeEpoch)
			c.removeValidatorDelegationRcItem(validatorAddr, item.NextStakeEpoch)
		} else {
			pre.UpdateNextStakeEpoch(item.NextStakeEpoch)
			next.UpdatePreStakeEpoch(item.PreStakeEpoch)
			if err := c.setValidatorDelegationRcItem(validatorAddr, item.PreStakeEpoch, pre); nil != err {
				return err
			}
			if err := c.setValidatorDelegationRcItem(validatorAddr, item.NextStakeEpoch, next); nil != err {
				return err
			}
		}
		c.removeValidatorDelegationRcItem(validatorAddr, stakeEpoch)
	} else {
		if err := c.setValidatorDelegationRcItem(validatorAddr, stakeEpoch, item); nil != err {
			return err
		}
	}
	return nil
}

func (c *StakeHandler) getValidatorDelegationRc(validatorAddr common.Address, stakeEpoch uint64) uint64 {
	item := c.getValidatorDelegationRcItem(validatorAddr, stakeEpoch)
	if nil == item {
		return 0
	}
	return item.Rc
}

// ----

func (c *StakeHandler) appendEpochItem(epoch uint64, startBlock, endBlock, roundCount uint64) error {
	return db.AppendEpochItem(c.evm.StateDB, c.contract.Address(), epoch, startBlock, endBlock, roundCount)
}

func (c *StakeHandler) getLastEpochItem() *types.EpochItem {
	return db.GetLastEpochItem(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) getLastEpoch() uint64 {
	return db.GetLastEpoch(c.evm.StateDB, c.contract.Address())
}

func (c *StakeHandler) getEpochQueueFromHead(size uint64) types.EpochQueue {
	return db.GetEpochQueueFromHead(c.evm.StateDB, c.contract.Address(), size)
}

func (c *StakeHandler) getEpochQueueFromTail(size uint64) types.EpochQueue {
	return db.GetEpochQueueFromTail(c.evm.StateDB, c.contract.Address(), size)
}

func (c *StakeHandler) setEpochItem(epoch uint64, item *types.EpochItem) error {
	return db.SetEpochItem(c.evm.StateDB, c.contract.Address(), epoch, item)
}

func (c *StakeHandler) getEpochItem(epoch uint64) *types.EpochItem {
	return db.GetEpochItem(c.evm.StateDB, c.contract.Address(), epoch)
}

// ----

func (c *StakeHandler) setSlashProcessed(handleEventId *big.Int, queue types.SlashValidatorWithdrawItemQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return db.ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeSlashProcessedKey(handleEventId), value)
	return nil
}

func (c *StakeHandler) getSlashProcessed(handleEventId *big.Int) types.SlashValidatorWithdrawItemQueue {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeSlashProcessedKey(handleEventId))
	if len(value) == 0 {
		return nil
	}

	var queue types.SlashValidatorWithdrawItemQueue
	if err := rlp.DecodeBytes(value, &queue); nil == err {
		return queue
	}
	return nil
}

func (c *StakeHandler) hasSlashProcessed(handleEventId *big.Int) bool {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeSlashProcessedKey(handleEventId))
	if len(value) == 0 {
		return false
	}

	var queue types.SlashValidatorWithdrawItemQueue
	if err := rlp.DecodeBytes(value, &queue); nil == err {
		return queue.NotEmpty()
	}
	return false
}

func (c *StakeHandler) hasNotSlashProcessed(handleEventId *big.Int) bool {
	return !c.hasSlashProcessed(handleEventId)
}
