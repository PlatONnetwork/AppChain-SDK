package contracts

import (
	"bytes"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/big"
)

var (
	ErrStoreFailed = errors.New("store failed")
	ErrRlpEncode   = errors.New("rlp encode failed")
	ErrRlpDecode   = errors.New("rlp decode failed")
	ErrNotFound    = errors.New("not found")
	ErrMisMatching = errors.New("mismatching")
)

var (
	//validatorDelegaterCountKeyPrefix    = []byte("validatorDelegaterCount")
	validatorNonceKey          = []byte("validatorNonce")        // "validatorNonce" => nonce (It is a self increasing stake index number)
	validatorKeyPrefix         = []byte("validator")             // "validator":validatorAddr => validator
	delegationKeyPrefix        = []byte("delegation")            // "delegater":delegaterAddr:validatorAddr:stakeBlockNumber => delegation
	currentEpochKey            = []byte("currentEpoch")          // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey            = []byte("currentRound")          // "currentRound" => currentRound (It is a number)
	epochValidatorIdsKeyPrefix = []byte("epochValidatorIds")     // "epochValidatorIds":epochId => []validatorAddr  (For settlement epoch)
	roundValidatorIdsKeyPrefix = []byte("roundValidatorIds")     // "roundValidatorIds":roundId => []validatorAddr  (For consensus round)
	PriorityValidatorHeadKey   = []byte("priorityValidatorHead") // "priorityValidatorHead" => priorityValidator(head)
	PriorityValidatorTailKey   = []byte("priorityValidatorTail") // "priorityValidatorTail" => priorityValidator(tail)
	priorityValidatorKeyPrefix = []byte("priorityValidator")     // "priorityValidator":shares(stakeAmount+delegataionAmount):blockNumber:stakeIndex => priorityValidator{preKey, nextKey, validatorAddr}

	stakeWithdrawalQueueItemKeyPrefix = []byte("stakeWithdrawalQueueItem") // "stakeWithdrawalQueueItem":validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}

	delegateWithdrawalQueueItemKeyPrefix = []byte("delegateWithdrawalQueueItem") // "delegateWithdrawalQueueItem":delegaterAddr:validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}

	validatorRcKeyPrefix         = []byte("validatorRc")         // "validatorRc":validatorAddr:stakeBlockNumber => delegation count
	unStakeDelegationRcKeyPrefix = []byte("unStakeDelegationRc") // "unStakeDelegationRc":validatorAddr:stakeBlockNumber => unStakeDelegationRcItem{preStakeBlock, nextStakeBlock, delegation count}

	epochBoundKeyPrefix      = []byte("epochBound")      // "epochBound":epochId => {start, end}  maybe add block root range start and end ??????
	blockRangeEpochKeyPrefix = []byte("blockRangeEpoch") // "blockRangeEpoch":start:end => epoch
	slashProcessedKeyPrefix  = []byte("slashProcessed")  // "slashProcessed":handleEventId => []SlashValidatorWithdrawItem{validatorAddr, amount}
)

func EncodeValidatorKey(validatorAddr common.Address) []byte {
	return append(validatorKeyPrefix, validatorAddr.Bytes()...)
}

func EncodeDelegaterKey(delegaterAddr, validatorAddr common.Address, stakeBlock uint64) []byte {
	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()
	stakeBlockBytes := common.Uint64ToBytes(stakeBlock)

	keyPrefixSize := len(delegationKeyPrefix)
	appendDelegaterSize := keyPrefixSize + len(delegaterAddrBytes)
	appendVlidatorAddrSize := appendDelegaterSize + len(validatorAddrBytes)
	size := appendVlidatorAddrSize + len(stakeBlockBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegationKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterSize], delegaterAddrBytes)
	copy(key[appendDelegaterSize:appendVlidatorAddrSize], validatorAddrBytes)
	copy(key[appendVlidatorAddrSize:], stakeBlockBytes)

	return key
}

func EncodeEpochValidatorIdsKey(epoch uint64) []byte {
	return append(epochValidatorIdsKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func EncodeRoundValidatorIdsKey(round uint64) []byte {
	return append(roundValidatorIdsKeyPrefix, common.Uint64ToBytes(round)...)
}

func EncodePriorityValidatorKey(blockNumber, stakeIndex uint64, shares *big.Int) []byte {

	sharesSub := new(big.Int).Sub(math.MaxBig104, shares)
	zeros := make([]byte, len(math.MaxBig104.Bytes()))
	sharesPriority := append(zeros, sharesSub.Bytes()...)

	stakeBlock := common.Uint64ToBytes(blockNumber)
	index := common.Uint64ToBytes(stakeIndex)

	// some index of pivots
	keyPrefixSize := len(priorityValidatorKeyPrefix)
	appendSharePrioritySize := keyPrefixSize + len(sharesPriority)
	appendstakeBlockSize := appendSharePrioritySize + len(stakeBlock)
	size := appendstakeBlockSize + len(index)

	// build key
	key := make([]byte, size)
	copy(key[:keyPrefixSize], priorityValidatorKeyPrefix)
	copy(key[keyPrefixSize:appendSharePrioritySize], sharesPriority)
	copy(key[appendSharePrioritySize:appendstakeBlockSize], stakeBlock)
	copy(key[appendstakeBlockSize:], index)

	return key
}

func EncodeStakeWithdrawalQueueItemKey(validatorAddr common.Address, unlockEpoch uint64) []byte {

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

func EncodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr common.Address, unlockEpoch uint64) []byte {

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

func EncodeValidatorRcKey(validatorAddr common.Address, stakeBlock uint64) []byte {

	validatorAddrBytes := validatorAddr.Bytes()
	indexBytes := common.Uint64ToBytes(stakeBlock)

	keyPrefixSize := len(validatorRcKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(indexBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], validatorRcKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], indexBytes)

	return key
}

func EncodeUnStakeDelegationRcKey(validatorAddr common.Address, stakeBlock uint64) []byte {
	validatorAddrBytes := validatorAddr.Bytes()
	stakeBlockBytes := common.Uint64ToBytes(stakeBlock)

	keyPrefixSize := len(unStakeDelegationRcKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(stakeBlockBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], unStakeDelegationRcKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], stakeBlockBytes)

	return key
}

func EncodeEpochBoundKey(epoch uint64) []byte {
	return append(epochBoundKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func EncodeBlockRangeEpochKey(startBlock, endBlock uint64) []byte {
	startBlockBytes := common.Uint64ToBytes(startBlock)
	endBlockBytes := common.Uint64ToBytes(endBlock)

	keyPrefixSize := len(blockRangeEpochKeyPrefix)
	appendStartBlockSize := keyPrefixSize + len(startBlockBytes)
	size := appendStartBlockSize + len(endBlockBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], blockRangeEpochKeyPrefix)
	copy(key[keyPrefixSize:appendStartBlockSize], startBlockBytes)
	copy(key[appendStartBlockSize:], endBlockBytes)

	return key
}

func EncodeSlashProcessedKey(handleEventId *big.Int) []byte {
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

func (c *StakeHandler) setValidator(validatorAddr common.Address, validator *types.Validator) error {
	value, err := rlp.EncodeToBytes(validator)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeValidatorKey(validatorAddr), value)
	return nil
}

func (c *StakeHandler) setValidatorByPriority(validatorAddr common.Address, validator *types.Validator) error {
	if err := c.setValidatorPriority(validatorAddr, validator.BlockNumber, validator.StakeIndex, validator.Shares()); nil != err {
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
	if c.getValidatorPriority(old.BlockNumber, old.StakeIndex, old.Shares()).ValidatorAddr != validatorAddr {
		return ErrMisMatching
	}
	if err := c.removeValidatorPriority(old.BlockNumber, old.StakeIndex, old.Shares()); nil != err {
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
	if c.getValidatorPriority(old.BlockNumber, old.StakeIndex, old.Shares()).ValidatorAddr != validatorAddr {
		return ErrMisMatching
	}
	if err := c.removeValidatorPriority(old.BlockNumber, old.StakeIndex, old.Shares()); nil != err {
		return err
	}
	// set new priority and validator
	return c.setValidatorByPriority(validatorAddr, validator)
}

func (c *StakeHandler) GetValidator(validatorAddr common.Address) *types.Validator {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return nil
	}
	var validator types.Validator
	if err := rlp.DecodeBytes(value, &validator); nil == err {
		return &validator
	}
	return nil
}

func (c *StakeHandler) hasValidator(validatorAddr common.Address) bool {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return false
	}
	return true
}
func (c *StakeHandler) hasNotValidator(validatorAddr common.Address) bool {
	return !c.hasValidator(validatorAddr)
}

func (c *StakeHandler) removeValidator(validatorAddr common.Address) {
	c.evm.StateDB.SetState(c.contract.Address(), EncodeValidatorKey(validatorAddr), []byte{})
}

func (c *StakeHandler) setDelegation(delegaterAddr, validatorAddr common.Address, stakeBlock uint64, delegation *types.Delegation) error {
	value, err := rlp.EncodeToBytes(delegation)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeDelegaterKey(delegaterAddr, validatorAddr, stakeBlock), value)
	return nil
}

func (c *StakeHandler) updateDelegation(delegaterAddr, validatorAddr common.Address, stakeBlock uint64, delegation *types.Delegation) error {
	del := c.GetDelegation(delegaterAddr, validatorAddr, stakeBlock)
	if nil != del {
		del.UpdateBlockNumber(delegation.BlockNumber)
		del.AddAmount(delegation.Amount)
	} else {
		del = delegation
		c.incrementValidatorRc(validatorAddr, stakeBlock, 1)
	}
	return c.setDelegation(delegaterAddr, validatorAddr, stakeBlock, del)
}

func (c *StakeHandler) GetDelegation(delegaterAddr, validatorAddr common.Address, stakeBlock uint64) *types.Delegation {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeDelegaterKey(delegaterAddr, validatorAddr, stakeBlock))
	if len(value) == 0 {
		return nil
	}
	var delegation types.Delegation
	if err := rlp.DecodeBytes(value, &delegation); nil == err {
		return &delegation
	}
	return nil
}

func (c *StakeHandler) removeDelegation(delegaterAddr, validatorAddr common.Address, stakeBlock uint64) {
	del := c.GetDelegation(delegaterAddr, validatorAddr, stakeBlock)
	if nil != del {
		c.decrementValidatorRc(validatorAddr, stakeBlock, 1)
		// todo unstakeRC --
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeDelegaterKey(delegaterAddr, validatorAddr, stakeBlock), []byte{})
}

func (c *StakeHandler) setCurrentEpoch(epoch uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), currentEpochKey, common.Uint64ToBytes(epoch))
}

func (c *StakeHandler) GetCurrentEpoch() uint64 {
	value := c.evm.StateDB.GetState(c.contract.Address(), currentEpochKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func (c *StakeHandler) setCurrentRound(round uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), currentRoundKey, common.Uint64ToBytes(round))
}

func (c *StakeHandler) GetCurrentRound() uint64 {
	value := c.evm.StateDB.GetState(c.contract.Address(), currentRoundKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func (c *StakeHandler) setEpochValidatorIds(epoch uint64, ids types.ValidatorIds) error {
	value, err := rlp.EncodeToBytes(ids)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeEpochValidatorIdsKey(epoch), value)
	return nil
}

func (c *StakeHandler) GetEpochValidatorIds(epoch uint64) types.ValidatorIds {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeEpochValidatorIdsKey(epoch))
	if len(value) == 0 {
		return nil
	}
	var ids types.ValidatorIds
	if err := rlp.DecodeBytes(value, &ids); nil == err {
		return ids
	}
	return nil
}

func (c *StakeHandler) setRoundValidatorIds(round uint64, ids types.ValidatorIds) error {
	value, err := rlp.EncodeToBytes(ids)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeRoundValidatorIdsKey(round), value)
	return nil
}

func (c *StakeHandler) GetRoundValidatorIds(round uint64) types.ValidatorIds {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeRoundValidatorIdsKey(round))
	if len(value) == 0 {
		return nil
	}
	var ids types.ValidatorIds
	if err := rlp.DecodeBytes(value, &ids); nil == err {
		return ids
	}
	return nil
}

func (c *StakeHandler) getValidatorPriority(blockNumber, stakeIndex uint64, shares *big.Int) *types.PriorityValidator {
	return c.getValidatorPriorityByKey(EncodePriorityValidatorKey(blockNumber, stakeIndex, shares))
}

func (c *StakeHandler) getValidatorPriorityByKey(key []byte) *types.PriorityValidator {
	value := c.evm.StateDB.GetState(c.contract.Address(), key)
	if len(value) == 0 {
		return nil
	}
	var priority types.PriorityValidator
	if err := rlp.DecodeBytes(value, &priority); nil == err {
		return &priority
	}
	return nil
}

func (c *StakeHandler) setValidatorPriorityByKey(key []byte, priority *types.PriorityValidator) error {
	value, err := rlp.EncodeToBytes(priority)
	if nil != err {
		return ErrRlpEncode
	}

	c.evm.StateDB.SetState(c.contract.Address(), key, value)
	return nil
}

func (c *StakeHandler) setValidatorPriority(validatorAddr common.Address, blockNumber, stakeIndex uint64, shares *big.Int) error {

	// ---------------------------------------------------------------------------------------------------------------------------------------------------------------

	preKey := PriorityValidatorHeadKey
	nextKey := PriorityValidatorTailKey
	pre := c.getValidatorPriorityByKey(PriorityValidatorHeadKey)
	next := c.getValidatorPriorityByKey(PriorityValidatorTailKey)

	priorityKey := EncodePriorityValidatorKey(blockNumber, stakeIndex, shares)
	priority := types.NewPriorityValidator(
		[]byte{},
		[]byte{},
		validatorAddr,
	)
	// first insert
	if bytes.Compare(pre.PreKey, pre.NextKey) == 0 && bytes.Compare(next.PreKey, next.NextKey) == 0 {
		priority.UpdatePreKey(preKey)
		priority.UpdateNextKey(nextKey)
		pre.UpdateNextKey(priorityKey)
		next.UpdatePreKey(priorityKey)
	} else {

		// TODO 感觉这里会有 bug 啊
		if bytes.Compare(priorityKey, pre.NextKey) < 0 { // add on head

			next = c.getValidatorPriorityByKey(pre.NextKey)

			// head<pre> -> priority -> next -> ... -> tail -> head
			next.UpdatePreKey(priorityKey)
			priority.UpdatePreKey(preKey)
			priority.UpdateNextKey(pre.NextKey)
			pre.UpdateNextKey(priorityKey)

		} else if bytes.Compare(priorityKey, next.NextKey) > 0 { // add on tail

			pre = c.getValidatorPriorityByKey(next.PreKey)

			// head -> ... -> pre -> priority -> tail<next> -> head
			pre.UpdateNextKey(priorityKey)
			priority.UpdatePreKey(next.PreKey)
			priority.UpdateNextKey(nextKey)
			next.UpdatePreKey(priorityKey)

		} else { // start from head
			nextKey = pre.NextKey
			next = c.getValidatorPriorityByKey(nextKey)
			for bytes.Compare(next.NextKey, PriorityValidatorHeadKey) != 0 && bytes.Compare(priorityKey, nextKey) < 0 { // if next is not tail && priorityKey < next(key)
				preKey = nextKey
				nextKey = next.NextKey
				pre = next
				next = c.getValidatorPriorityByKey(nextKey)
			}
			// insert front of next
			priority.UpdatePreKey(next.PreKey)
			priority.UpdateNextKey(pre.NextKey)
			pre.UpdateNextKey(priorityKey)
			next.UpdatePreKey(priorityKey)

		}
	}

	pvalue, err := rlp.EncodeToBytes(pre)
	if nil != err {
		return ErrRlpEncode
	}
	nvalue, err := rlp.EncodeToBytes(next)
	if nil != err {
		return ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(priority)
	if nil != err {
		return ErrRlpEncode
	}

	c.evm.StateDB.SetState(c.contract.Address(), preKey, pvalue)
	c.evm.StateDB.SetState(c.contract.Address(), nextKey, nvalue)
	c.evm.StateDB.SetState(c.contract.Address(), priorityKey, value)
	return nil
}

func (c *StakeHandler) removeValidatorPriority(blockNumber, stakeIndex uint64, shares *big.Int) error {

	priorityKey := EncodePriorityValidatorKey(blockNumber, stakeIndex, shares)
	priority := c.getValidatorPriorityByKey(priorityKey)

	preKey := priority.PreKey
	nextKey := priority.NextKey

	pre := c.getValidatorPriorityByKey(preKey)
	next := c.getValidatorPriorityByKey(nextKey)

	pre.UpdateNextKey(nextKey)
	next.UpdatePreKey(preKey)

	pvalue, err := rlp.EncodeToBytes(pre)
	if nil != err {
		return ErrRlpEncode
	}
	nvalue, err := rlp.EncodeToBytes(next)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), preKey, pvalue)
	c.evm.StateDB.SetState(c.contract.Address(), nextKey, nvalue)
	c.evm.StateDB.SetState(c.contract.Address(), priorityKey, []byte{})

	return nil
}

func (c *StakeHandler) rankPriorityValidatorIds(size uint64) types.ValidatorIds {

	arr := make(types.ValidatorIds, size)
	var count uint64 = 0

	headItem := c.getValidatorPriorityByKey(PriorityValidatorHeadKey)
	itemKey := headItem.NextKey
	item := c.getValidatorPriorityByKey(itemKey)

	for bytes.Compare(item.NextKey, PriorityValidatorHeadKey) != 0 && count < size {
		arr[count] = item.ValidatorAddr
		itemKey = item.NextKey
		item = c.getValidatorPriorityByKey(itemKey)
		count++
	}
	return arr[:count]
}

func (c *StakeHandler) appendStakeWithdrawal(validatorAddr common.Address, epoch uint64, amount *big.Int) error {

	indexEpoch := uint64(0)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// first insert
	if nil == indexItem {

		nextEpoch := uint64(math.MaxUint64)
		// tail -> head -> epoch -> tail -> head

		indexItem = types.NewStakeWithdrawalItem(nextEpoch, epoch, common.Big0)
		epochItem := types.NewStakeWithdrawalItem(indexEpoch, nextEpoch, amount)
		nextItem := types.NewStakeWithdrawalItem(epoch, indexEpoch, common.Big0)

		if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexEpoch, indexItem); nil != err {
			return err
		}
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, nextEpoch, nextItem); nil != err {
			return err
		}
	} else {
		// range from head to tail
		for indexItem.NextEpoch != uint64(0) {

			if indexEpoch == epoch { // pre -> index(epoch) -> next

				indexItem.IncrementAmount(amount)
				if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexEpoch, indexItem); nil != err {
					return err
				}
				break
			} else if indexEpoch > epoch { // pre -> epoch -> index -> ... -> max

				pre := c.getStakeWithdrawalQueueItem(validatorAddr, indexItem.PreEpoch)
				epochItem := types.NewStakeWithdrawalItem(indexItem.PreEpoch, indexEpoch, amount)

				pre.UpdateNextEpoch(epoch)
				if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexItem.PreEpoch, pre); nil != err {
					return err
				}

				indexItem.UpdatePreEpoch(epoch)
				if err := c.setStakeWithdrawalQueueItem(validatorAddr, indexEpoch, indexItem); nil != err {
					return err
				}

				if err := c.setStakeWithdrawalQueueItem(validatorAddr, epoch, epochItem); nil != err {
					return err
				}

				break
			} else { // index -> epoch -> next -> ... -> max | index -> epoch -> max
				indexEpoch = indexItem.NextEpoch
				indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
			}
		}
	}

	return nil
}

func (c *StakeHandler) GetStakeWithdrawal(validatorAddr common.Address) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	for indexItem.NextEpoch != uint64(0) {
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

// Total of all rewards until epoch
func (c *StakeHandler) GetStakeWithdrawable(validatorAddr common.Address, epoch uint64) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// empty queue
	if nil == indexItem {
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
	if nil == indexItem {
		return amount, nil
	}
	for indexItem.NextEpoch != uint64(0) {

		if indexEpoch > epoch {
			break
		}
		amount = new(big.Int).Add(amount, indexItem.Amount)

		// remove item
		if indexEpoch != uint64(0) {
			c.removeStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
		}

		indexEpoch = indexItem.NextEpoch
		indexItem = c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)
	}

	// remove head and tail
	if indexItem.NextEpoch == uint64(0) {
		c.removeStakeWithdrawalQueueItem(validatorAddr, uint64(0))
		c.removeStakeWithdrawalQueueItem(validatorAddr, uint64(math.MaxUint64))
	} else { // update start index
		head := c.getStakeWithdrawalQueueItem(validatorAddr, uint64(0))
		head.UpdateNextEpoch(indexEpoch)
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, uint64(0), head); nil != err {
			return common.Big0, err
		}
	}

	return amount, nil
}

// Total of all rewards since epoch
func (c *StakeHandler) GetStakeWithdrawalPending(validatorAddr common.Address, epoch uint64) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(math.MaxUint64)
	indexItem := c.getStakeWithdrawalQueueItem(validatorAddr, indexEpoch)

	// empty queue
	if nil == indexItem {
		return amount
	}

	for indexItem.PreEpoch != uint64(math.MaxUint64) {

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
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeStakeWithdrawalQueueItemKey(validatorAddr, epoch), value)
	return nil
}

func (c *StakeHandler) removeStakeWithdrawalQueueItem(validatorAddr common.Address, epoch uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), EncodeStakeWithdrawalQueueItemKey(validatorAddr, epoch), []byte{})
}

func (c *StakeHandler) getStakeWithdrawalQueueItem(validatorAddr common.Address, epoch uint64) *types.StakeWithdrawalItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeStakeWithdrawalQueueItemKey(validatorAddr, epoch))

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

	indexEpoch := uint64(0)
	indexItem := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)

	// first insert
	if nil == indexItem {

		nextEpoch := uint64(math.MaxUint64)
		// tail -> head -> epoch -> tail -> head

		indexItem = types.NewDelegateWithdrawalItem(nextEpoch, epoch, common.Big0)
		epochItem := types.NewDelegateWithdrawalItem(indexEpoch, nextEpoch, amount)
		nextItem := types.NewDelegateWithdrawalItem(epoch, indexEpoch, common.Big0)

		if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
			return err
		}
		if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, nextEpoch, nextItem); nil != err {
			return err
		}
	} else {
		// range from head to tail
		for indexItem.NextEpoch != uint64(0) {

			if indexEpoch == epoch { // pre -> index(epoch) -> next

				indexItem.IncrementAmount(amount)
				if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
					return err
				}
				break
			} else if indexEpoch > epoch { // pre -> epoch -> index -> ... -> max

				pre := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexItem.PreEpoch)
				epochItem := types.NewDelegateWithdrawalItem(indexItem.PreEpoch, indexEpoch, amount)

				pre.UpdateNextEpoch(epoch)
				if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexItem.PreEpoch, pre); nil != err {
					return err
				}

				indexItem.UpdatePreEpoch(epoch)
				if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
					return err
				}

				if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
					return err
				}

				break
			} else { // index -> epoch -> next -> ... -> max | index -> epoch -> max
				indexEpoch = indexItem.NextEpoch
				indexItem = c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)
			}
		}
	}

	return nil
}

func (c *StakeHandler) GetDelegateWithdrawal(delegaterAddr, validatorAddr common.Address) *big.Int {
	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, indexEpoch)

	for indexItem.NextEpoch != uint64(0) {
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
	if nil == indexItem {
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
	if nil == indexItem {
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
		c.removeDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, uint64(0))
		c.removeDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, uint64(math.MaxUint64))
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
	if nil == indexItem {
		return amount
	}

	for indexItem.PreEpoch != uint64(math.MaxUint64) {

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
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch), value)
	return nil
}

func (c *StakeHandler) removeDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr common.Address, epoch uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), EncodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch), []byte{})
}

func (c *StakeHandler) getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr common.Address, epoch uint64) *types.DelegateWithdrawalItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch))

	if len(value) == 0 {
		return nil
	}
	var item types.DelegateWithdrawalItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

// ---

func (c *StakeHandler) incrementValidatorRc(validatorAddr common.Address, stakeBlock, increment uint64) {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeValidatorRcKey(validatorAddr, stakeBlock))
	var rc uint64
	if len(value) != 0 {
		rc = common.BytesToUint64(value)
	}
	rc += increment
	c.evm.StateDB.SetState(c.contract.Address(), EncodeValidatorRcKey(validatorAddr, stakeBlock), common.Uint64ToBytes(rc))
}

func (c *StakeHandler) decrementValidatorRc(validatorAddr common.Address, stakeBlock, decrement uint64) {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeValidatorRcKey(validatorAddr, stakeBlock))
	var rc uint64
	if len(value) != 0 {
		rc = common.BytesToUint64(value)
	}
	if rc < decrement {
		rc = 0
	} else {
		rc -= decrement
	}
	if rc == 0 {
		value = []byte{}
	} else {
		value = common.Uint64ToBytes(rc)
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeValidatorRcKey(validatorAddr, stakeBlock), value)
}

// ------- encodeUnStakeRcBoundKey

func (c *StakeHandler) appendUnStakeDelegationRc(validatorAddr common.Address, stakeBlock, rc uint64) error {

	indexStakeBlock := uint64(0)
	indexItem := c.getUnStakeDelegationRcItem(validatorAddr, indexStakeBlock)

	// first insert
	if nil == indexItem {

		nextStakeBlock := uint64(math.MaxUint64)
		// tail -> head -> stakeBlock -> tail -> head

		indexItem = types.NewUnStakeDelegationRcItem(nextStakeBlock, stakeBlock, 0)
		stakeBlockItem := types.NewUnStakeDelegationRcItem(indexStakeBlock, nextStakeBlock, rc)
		nextItem := types.NewUnStakeDelegationRcItem(stakeBlock, indexStakeBlock, 0)

		if err := c.setUnStakeDelegationRcItem(validatorAddr, indexStakeBlock, indexItem); nil != err {
			return err
		}
		if err := c.setUnStakeDelegationRcItem(validatorAddr, stakeBlock, stakeBlockItem); nil != err {
			return err
		}
		if err := c.setUnStakeDelegationRcItem(validatorAddr, nextStakeBlock, nextItem); nil != err {
			return err
		}
	} else {
		// range from head to tail
		for indexItem.NextStakeBlock != uint64(0) {

			if indexStakeBlock == stakeBlock { // pre -> index(stakeBlock) -> next

				indexItem.IncrementRc(rc)
				if err := c.setUnStakeDelegationRcItem(validatorAddr, indexStakeBlock, indexItem); nil != err {
					return err
				}
				break
			} else if indexStakeBlock > stakeBlock { // pre -> stakeBlock -> index -> ... -> max

				pre := c.getUnStakeDelegationRcItem(validatorAddr, indexItem.PreStakeBlock)
				epochItem := types.NewUnStakeDelegationRcItem(indexItem.PreStakeBlock, indexStakeBlock, rc)

				pre.UpdateNextStakeBlock(stakeBlock)
				if err := c.setUnStakeDelegationRcItem(validatorAddr, indexItem.PreStakeBlock, pre); nil != err {
					return err
				}

				indexItem.UpdatePreStakeBlock(stakeBlock)
				if err := c.setUnStakeDelegationRcItem(validatorAddr, indexStakeBlock, indexItem); nil != err {
					return err
				}

				if err := c.setUnStakeDelegationRcItem(validatorAddr, stakeBlock, epochItem); nil != err {
					return err
				}

				break
			} else { // index -> stakeBlock -> next -> ... -> max | index -> stakeBlock -> max
				indexStakeBlock = indexItem.NextStakeBlock
				indexItem = c.getUnStakeDelegationRcItem(validatorAddr, indexStakeBlock)
			}
		}
	}

	return nil
}

func (c *StakeHandler) getUnStakeDelegationRcItem(validatorAddr common.Address, stakeBlock uint64) *types.UnStakeDelegationRcItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeUnStakeDelegationRcKey(validatorAddr, stakeBlock))

	if len(value) != 0 {
		return nil
	}
	var item types.UnStakeDelegationRcItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

func (c *StakeHandler) setUnStakeDelegationRcItem(validatorAddr common.Address, stakeBlock uint64, item *types.UnStakeDelegationRcItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeUnStakeDelegationRcKey(validatorAddr, stakeBlock), value)
	return nil
}

func (c *StakeHandler) removeUnStakeDelegationRcItem(validatorAddr common.Address, stakeBlock uint64) {
	c.evm.StateDB.SetState(c.contract.Address(), EncodeUnStakeDelegationRcKey(validatorAddr, stakeBlock), []byte{})
}

func (c *StakeHandler) incrementUnStakeDelegationRc(validatorAddr common.Address, stakeBlock, increment uint64) error {
	item := c.getUnStakeDelegationRcItem(validatorAddr, stakeBlock)
	if nil == item {
		return ErrNotFound
	}
	item.IncrementRc(increment)
	return c.setUnStakeDelegationRcItem(validatorAddr, stakeBlock, item)
}

func (c *StakeHandler) decrementUnStakeDelegationRc(validatorAddr common.Address, stakeBlock, decrement uint64) error {
	item := c.getUnStakeDelegationRcItem(validatorAddr, stakeBlock)
	if nil == item {
		return ErrNotFound
	}

	item.DecrementRc(decrement)

	// change index
	if item.Rc == 0 {
		pre := c.getUnStakeDelegationRcItem(validatorAddr, item.PreStakeBlock)
		next := c.getUnStakeDelegationRcItem(validatorAddr, item.NextStakeBlock)

		// remove the last one   tail -> head -> lastone(remove) -> tail -> head
		if pre.PreStakeBlock == uint64(math.MaxUint64) && next.NextStakeBlock == uint64(0) {
			c.removeUnStakeDelegationRcItem(validatorAddr, item.PreStakeBlock)
			c.removeUnStakeDelegationRcItem(validatorAddr, item.NextStakeBlock)
		} else {
			pre.UpdateNextStakeBlock(item.NextStakeBlock)
			next.UpdatePreStakeBlock(item.PreStakeBlock)
			if err := c.setUnStakeDelegationRcItem(validatorAddr, item.PreStakeBlock, pre); nil != err {
				return err
			}
			if err := c.setUnStakeDelegationRcItem(validatorAddr, item.NextStakeBlock, next); nil != err {
				return err
			}
		}
		c.removeUnStakeDelegationRcItem(validatorAddr, stakeBlock)
	} else {
		if err := c.setUnStakeDelegationRcItem(validatorAddr, stakeBlock, item); nil != err {
			return err
		}
	}
	return nil
}

func (c *StakeHandler) getUnStakeDelegationRc(validatorAddr common.Address, stakeBlock uint64) uint64 {
	item := c.getUnStakeDelegationRcItem(validatorAddr, stakeBlock)
	if nil == item {
		return 0
	}
	return item.Rc
}

// ----

func (c *StakeHandler) setSlashProcessed(handleEventId *big.Int, queue types.SlashValidatorWithdrawItemQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), EncodeSlashProcessedKey(handleEventId), value)
	return nil
}

func (c *StakeHandler) getSlashProcessed(handleEventId *big.Int) types.SlashValidatorWithdrawItemQueue {
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeSlashProcessedKey(handleEventId))
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
	value := c.evm.StateDB.GetState(c.contract.Address(), EncodeSlashProcessedKey(handleEventId))
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
