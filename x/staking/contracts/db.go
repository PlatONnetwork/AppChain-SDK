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
	validatorNonceKey                     = []byte("validatorNonce")               // "validatorNonce" => nonce (It is a self increasing stake index number)
	validatorKeyPrefix                    = []byte("validator")                    // "validator":validatorAddr => validator
	delegationKeyPrefix                   = []byte("delegation")                   // "delegater":delegaterAddr:validatorAddr => delegation
	currentEpochKey                       = []byte("currentEpoch")                 // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey                       = []byte("currentRound")                 // "currentRound" => currentRound (It is a number)
	epochValidatorIdsKeyPrefix            = []byte("epochValidatorIds")            // "epochValidatorIds":epochId => []validatorAddr  (For settlement epoch)
	roundValidatorIdsKeyPrefix            = []byte("roundValidatorIds")            // "roundValidatorIds":roundId => []validatorAddr  (For consensus round)
	priorityValidatorHeadKey              = []byte("priorityValidatorHead")        // "priorityValidatorHead" => priorityValidator(head)
	priorityValidatorTailKey              = []byte("priorityValidatorTail")        // "priorityValidatorTail" => priorityValidator(tail)
	priorityValidatorKeyPrefix            = []byte("priorityValidator")            // "priorityValidator":shares(stakeAmount+delegataionAmount):blockNumber:stakeIndex => priorityValidator
	stakeWithdrawalQueueBoundKeyPrefix    = []byte("stakeWithdrawalQueueBound")    // "stakeWithdrawalQueueBound":validatorAddr => StakeWithdrawalBound{head, tail}
	stakeWithdrawalQueueItemKeyPrefix     = []byte("stakeWithdrawalQueueItem")     // "stakeWithdrawalQueueItem":validatorAddr:index => {amount, (unlock)epoch}
	delegateWithdrawalQueueBoundKeyPrefix = []byte("delegateWithdrawalQueueBound") // "delegateWithdrawalQueueBound":delegaterAddr:validatorAddr => {head, tail}
	delegateWithdrawalQueueItemKeyPrefix  = []byte("delegateWithdrawalQueueItem")  // "delegateWithdrawalQueueItem":delegaterAddr:validatorAddr:index => {amount, (unlock)epoch}
	epochBoundKeyPrefix                   = []byte("epochBound")                   // "epochBound":epochId => {start, end}  maybe add block root range start and end ??????
	blockRangeEpochKeyPrefix              = []byte("blockRangeEpoch")              // "blockRangeEpoch":start:end => epoch
	slashProcessedKeyPrefix               = []byte("slashProcessed")               // "slashProcessed":handleEventId => []SlashValidatorWithdrawItem{validatorAddr, amount}
)

func encodeValidatorKey(validatorAddr common.Address) []byte {
	return append(validatorKeyPrefix, validatorAddr.Bytes()...)
}

func encodeDelegaterKey(delegaterAddr, validatorAddr common.Address) []byte {
	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()

	keyPrefixSize := len(delegationKeyPrefix)
	appendDelegaterSize := keyPrefixSize + len(delegaterAddrBytes)
	size := appendDelegaterSize + len(validatorAddrBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegationKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterSize], delegaterAddrBytes)
	copy(key[appendDelegaterSize:], validatorAddrBytes)

	return key
}

func encodeEpochValidatorIdsKey(epoch uint64) []byte {
	return append(epochValidatorIdsKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func encodeRoundValidatorIdsKey(round uint64) []byte {
	return append(roundValidatorIdsKeyPrefix, common.Uint64ToBytes(round)...)
}

func encodePriorityValidatorKey(blockNumber, stakeIndex uint64, shares *big.Int) []byte {

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

func encodeStakeWithdrawalQueueBoundKey(validatorAddr common.Address) []byte {
	return append(stakeWithdrawalQueueBoundKeyPrefix, validatorAddr.Bytes()...)
}

func encodeStakeWithdrawalQueueItemKey(validatorAddr common.Address, index uint64) []byte {

	validatorAddrBytes := validatorAddr.Bytes()
	indexBytes := common.Uint64ToBytes(index)

	keyPrefixSize := len(stakeWithdrawalQueueItemKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(indexBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], stakeWithdrawalQueueItemKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], indexBytes)

	return key
}

func encodeDelegateWithdrawalQueueBoundKey(delegaterAddr, validatorAddr common.Address) []byte {

	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()

	keyPrefixSize := len(delegateWithdrawalQueueBoundKeyPrefix)
	appendDelegaterSize := keyPrefixSize + len(delegaterAddrBytes)
	size := appendDelegaterSize + len(validatorAddrBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegateWithdrawalQueueBoundKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterSize], delegaterAddrBytes)
	copy(key[appendDelegaterSize:], validatorAddrBytes)

	return key
}

func encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr common.Address, index uint64) []byte {

	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()
	indexBytes := common.Uint64ToBytes(index)

	keyPrefixSize := len(delegateWithdrawalQueueItemKeyPrefix)
	appendDelegaterSize := keyPrefixSize + len(delegaterAddrBytes)
	appendValidatorAddrSize := appendDelegaterSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(indexBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegateWithdrawalQueueItemKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterSize], delegaterAddrBytes)
	copy(key[appendDelegaterSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], indexBytes)

	return key
}

func encodeEpochBoundKey(epoch uint64) []byte {
	return append(epochBoundKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func encodeBlockRangeEpochKey(startBlock, endBlock uint64) []byte {
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

func (c *StakeHandler) setValidator(validatorAddr common.Address, validator *types.Validator) error {
	value, err := rlp.EncodeToBytes(validator)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeValidatorKey(validatorAddr), value)
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
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeValidatorKey(validatorAddr))
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
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return false
	}
	return true
}
func (c *StakeHandler) hasNotValidator(validatorAddr common.Address) bool {
	return !c.hasValidator(validatorAddr)
}

func (c *StakeHandler) removeValidator(validatorAddr common.Address) {
	c.evm.StateDB.SetState(c.contract.Address(), encodeValidatorKey(validatorAddr), []byte{})
}

func (c *StakeHandler) setDelegation(delegaterAddr, validatorAddr common.Address, delegation *types.Delegation) error {
	value, err := rlp.EncodeToBytes(delegation)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegaterKey(delegaterAddr, validatorAddr), value)
	return nil
}

func (c *StakeHandler) updateDelegation(delegaterAddr, validatorAddr common.Address, delegation *types.Delegation) error {
	del := c.GetDelegation(delegaterAddr, validatorAddr)
	if nil != del {
		del.UpdateBlockNumber(delegation.BlockNumber)
		del.AddAmount(delegation.Amount)
	} else {
		del = delegation
	}
	return c.setDelegation(delegaterAddr, validatorAddr, del)
}

func (c *StakeHandler) GetDelegation(delegaterAddr, validatorAddr common.Address) *types.Delegation {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeDelegaterKey(delegaterAddr, validatorAddr))
	if len(value) == 0 {
		return nil
	}
	var delegation types.Delegation
	if err := rlp.DecodeBytes(value, &delegation); nil == err {
		return &delegation
	}
	return nil
}

func (c *StakeHandler) removeDelegation(delegaterAddr, validatorAddr common.Address) {
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegaterKey(delegaterAddr, validatorAddr), []byte{})
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
	c.evm.StateDB.SetState(c.contract.Address(), encodeEpochValidatorIdsKey(epoch), value)
	return nil
}

func (c *StakeHandler) GetEpochValidatorIds(epoch uint64) types.ValidatorIds {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeEpochValidatorIdsKey(epoch))
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
	c.evm.StateDB.SetState(c.contract.Address(), encodeRoundValidatorIdsKey(round), value)
	return nil
}

func (c *StakeHandler) GetRoundValidatorIds(round uint64) types.ValidatorIds {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeRoundValidatorIdsKey(round))
	if len(value) == 0 {
		return nil
	}
	var ids types.ValidatorIds
	if err := rlp.DecodeBytes(value, &ids); nil == err {
		return ids
	}
	return nil
}

func (c *StakeHandler) initValidatorPriority() error {
	head := types.NewPriorityValidator(
		priorityValidatorTailKey,
		priorityValidatorTailKey,
		common.ZeroAddr,
	)
	tail := types.NewPriorityValidator(
		priorityValidatorHeadKey,
		priorityValidatorHeadKey,
		common.ZeroAddr,
	)
	hvalue, err := rlp.EncodeToBytes(head)
	if nil != err {
		return ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tail)
	if nil != err {
		return ErrRlpEncode
	}

	c.evm.StateDB.SetState(c.contract.Address(), priorityValidatorHeadKey, hvalue)
	c.evm.StateDB.SetState(c.contract.Address(), priorityValidatorTailKey, tvalue)
	return nil
}

func (c *StakeHandler) getValidatorPriority(blockNumber, stakeIndex uint64, shares *big.Int) *types.PriorityValidator {
	return c.getValidatorPriorityByKey(encodePriorityValidatorKey(blockNumber, stakeIndex, shares))
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

func (c *StakeHandler) setValidatorPriority(validatorAddr common.Address, blockNumber, stakeIndex uint64, shares *big.Int) error {
	preKey := priorityValidatorHeadKey
	nextKey := priorityValidatorTailKey
	pre := c.getValidatorPriorityByKey(priorityValidatorHeadKey)
	next := c.getValidatorPriorityByKey(priorityValidatorTailKey)

	priorityKey := encodePriorityValidatorKey(blockNumber, stakeIndex, shares)
	priority := types.NewPriorityValidator(
		[]byte{},
		[]byte{},
		validatorAddr,
	)
	// first insert
	if bytes.Compare(pre.Previous, pre.Next) == 0 && bytes.Compare(next.Previous, next.Next) == 0 {
		priority.UpdatePrevious(preKey)
		priority.UpdateNext(nextKey)
		pre.UpdateNext(priorityKey)
		next.UpdatePrevious(priorityKey)
	} else {

		if bytes.Compare(priorityKey, pre.Next) < 0 { // add on head

			next = c.getValidatorPriorityByKey(pre.Next)

			// head<pre> -> priority -> next -> ... -> tail -> head
			next.UpdatePrevious(priorityKey)
			priority.UpdatePrevious(preKey)
			priority.UpdateNext(pre.Next)
			pre.UpdateNext(priorityKey)

		} else if bytes.Compare(priorityKey, next.Next) > 0 { // add on tail

			pre = c.getValidatorPriorityByKey(next.Previous)

			// head -> ... -> pre -> priority -> tail<next> -> head
			pre.UpdateNext(priorityKey)
			priority.UpdatePrevious(next.Previous)
			priority.UpdateNext(nextKey)
			next.UpdatePrevious(priorityKey)

		} else { // start from head
			nextKey = pre.Next
			next = c.getValidatorPriorityByKey(nextKey)
			for bytes.Compare(next.Next, priorityValidatorHeadKey) != 0 && bytes.Compare(priorityKey, nextKey) < 0 { // if next is not tail && priorityKey < next(key)
				preKey = nextKey
				nextKey = next.Next
				pre = next
				next = c.getValidatorPriorityByKey(nextKey)
			}
			// insert front of next
			priority.UpdatePrevious(next.Previous)
			priority.UpdateNext(pre.Next)
			pre.UpdateNext(priorityKey)
			next.UpdatePrevious(priorityKey)

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

	priorityKey := encodePriorityValidatorKey(blockNumber, stakeIndex, shares)
	priority := c.getValidatorPriorityByKey(priorityKey)

	preKey := priority.Previous
	nextKey := priority.Next

	pre := c.getValidatorPriorityByKey(preKey)
	next := c.getValidatorPriorityByKey(nextKey)

	pre.UpdateNext(nextKey)
	next.UpdatePrevious(preKey)

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

	headItem := c.getValidatorPriorityByKey(priorityValidatorHeadKey)
	itemKey := headItem.Next
	item := c.getValidatorPriorityByKey(itemKey)

	for bytes.Compare(item.Next, priorityValidatorHeadKey) != 0 && count < size {
		arr[count] = item.ValidatorAddr
		itemKey = item.Next
		item = c.getValidatorPriorityByKey(itemKey)
		count++
	}
	return arr[:count]
}

func (c *StakeHandler) appendStakeWithdrawal(validatorAddr common.Address, epoch uint64, amount *big.Int) error {

	bound := c.getStakeWithdrawalQueueBound(validatorAddr)
	var withdraw *types.StakeWithdrawalItem
	var index uint64
	if nil == bound {
		bound = types.NewStakeWithdrawalBound(0, 1)
		withdraw = types.NewStakeWithdrawalItem(epoch, amount)
	} else {

		lastWithdraw := c.getStakeWithdrawalQueueItem(validatorAddr, bound.Tail-1)
		lastEpoch := lastWithdraw.Epoch

		if epoch > lastEpoch {
			// new withdrawal for next epoch
			index = bound.Tail
			bound.IncrementTail(1)
			withdraw = types.NewStakeWithdrawalItem(epoch, amount)
		} else if epoch == lastEpoch {
			index = bound.Tail - 1
			withdraw = lastWithdraw
			withdraw.IncrementAmount(amount)
		} else {
			return ErrNotFound
		}
	}
	if err := c.setStakeWithdrawalQueueBound(validatorAddr, bound); nil != err {
		return err
	}
	if err := c.setStakeWithdrawalQueueItem(validatorAddr, index, withdraw); nil != err {
		return err
	}
	return nil
}

func (c *StakeHandler) GetStakeWithdrawal(validatorAddr common.Address) *big.Int {
	bound := c.getStakeWithdrawalQueueBound(validatorAddr)
	amount := common.Big0
	for i := bound.Head; i < bound.Tail; i++ {
		withdraw := c.getStakeWithdrawalQueueItem(validatorAddr, i)
		amount = new(big.Int).Add(amount, withdraw.Amount)
	}
	return amount
}

func (c *StakeHandler) GetStakeWithdrawalByEpoch(validatorAddr common.Address, epoch uint64) *big.Int {
	bound := c.getStakeWithdrawalQueueBound(validatorAddr)
	for i := bound.Head; i < bound.Tail; i++ {
		withdraw := c.getStakeWithdrawalQueueItem(validatorAddr, i)
		if withdraw.Epoch == epoch {
			return withdraw.Amount
		}
	}
	return common.Big0
}

// Total of all rewards until epoch
func (c *StakeHandler) GetStakeWithdrawable(validatorAddr common.Address, epoch uint64) (*big.Int, uint64) {
	bound := c.getStakeWithdrawalQueueBound(validatorAddr)
	amount := common.Big0
	var newHead uint64
	for newHead = bound.Head; newHead < bound.Tail; newHead++ {
		withdraw := c.getStakeWithdrawalQueueItem(validatorAddr, newHead)
		if withdraw.Epoch > epoch {
			break
		}
		amount = new(big.Int).Add(amount, withdraw.Amount)
	}
	return amount, newHead
}

// Total of all rewards since epoch
func (c *StakeHandler) GetStakeWithdrawalPending(validatorAddr common.Address, epoch uint64) *big.Int {
	bound := c.getStakeWithdrawalQueueBound(validatorAddr)
	amount := common.Big0
	for i := bound.Tail; i >= bound.Head; i-- {
		withdraw := c.getStakeWithdrawalQueueItem(validatorAddr, i)
		if withdraw.Epoch <= epoch {
			break
		}
		amount = new(big.Int).Add(amount, withdraw.Amount)
	}
	return amount
}

func (c *StakeHandler) setStakeWithdrawalQueueBound(validatorAddr common.Address, bound *types.StakeWithdrawalBound) error {
	value, err := rlp.EncodeToBytes(bound)
	if nil != err {
		return err
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeStakeWithdrawalQueueBoundKey(validatorAddr), value)
	return nil
}

func (c *StakeHandler) updateStakeWithdrawQueueBoundHead(validatorAddr common.Address, head uint64) error {
	bound := c.getStakeWithdrawalQueueBound(validatorAddr)
	if nil == bound {
		return ErrNotFound
	}
	bound.UpdateHead(head)
	return c.setStakeWithdrawalQueueBound(validatorAddr, bound)
}

func (c *StakeHandler) getStakeWithdrawalQueueBound(validatorAddr common.Address) *types.StakeWithdrawalBound {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeStakeWithdrawalQueueBoundKey(validatorAddr))

	if len(value) == 0 {
		return nil
	}
	var bound types.StakeWithdrawalBound
	if err := rlp.DecodeBytes(value, &bound); nil == err {
		return &bound
	}
	return nil
}

func (c *StakeHandler) setStakeWithdrawalQueueItem(validatorAddr common.Address, index uint64, withdraw *types.StakeWithdrawalItem) error {
	value, err := rlp.EncodeToBytes(withdraw)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeStakeWithdrawalQueueItemKey(validatorAddr, index), value)
	return nil
}

func (c *StakeHandler) getStakeWithdrawalQueueItem(validatorAddr common.Address, index uint64) *types.StakeWithdrawalItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeStakeWithdrawalQueueItemKey(validatorAddr, index))

	if len(value) == 0 {
		return nil
	}
	var withdraw types.StakeWithdrawalItem
	if err := rlp.DecodeBytes(value, &withdraw); nil == err {
		return &withdraw
	}
	return nil
}

// -------------

func (c *StakeHandler) appendDelegateWithdrawal(delegaterAddr, validatorAddr common.Address, epoch uint64, amount *big.Int) error {

	bound := c.getDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr)
	var withdraw *types.DelegateWithdrawalItem
	var index uint64
	if nil == bound {
		bound = types.NewDelegateWithdrawalBound(0, 1)
		withdraw = types.NewDelegateWithdrawalItem(epoch, amount)
	} else {

		lastWithdraw := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, bound.Tail-1)
		lastEpoch := lastWithdraw.Epoch

		if epoch > lastEpoch {
			// new withdrawal for next epoch
			index = bound.Tail
			bound.IncrementTail(1)
			withdraw = types.NewDelegateWithdrawalItem(epoch, amount)
		} else if epoch == lastEpoch {
			index = bound.Tail - 1
			withdraw = lastWithdraw
			withdraw.IncrementAmount(amount)
		} else {
			return ErrNotFound
		}
	}
	if err := c.setDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr, bound); nil != err {
		return err
	}
	if err := c.setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, index, withdraw); nil != err {
		return err
	}
	return nil
}

func (c *StakeHandler) GetDelegateWithdrawal(delegaterAddr, validatorAddr common.Address) *big.Int {
	bound := c.getDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr)
	amount := common.Big0
	for i := bound.Head; i < bound.Tail; i++ {
		withdraw := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, i)
		amount = new(big.Int).Add(amount, withdraw.Amount)
	}
	return amount
}

func (c *StakeHandler) GetDelegateWithdrawalByEpoch(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	bound := c.getDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr)
	for i := bound.Head; i < bound.Tail; i++ {
		withdraw := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, i)
		if withdraw.Epoch == epoch {
			return withdraw.Amount
		}
	}
	return common.Big0
}

// Total of all rewards until epoch
func (c *StakeHandler) GetDelegateWithdrawable(delegaterAddr, validatorAddr common.Address, epoch uint64) (*big.Int, uint64) {
	bound := c.getDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr)
	amount := common.Big0
	var newHead uint64
	for newHead = bound.Head; newHead < bound.Tail; newHead++ {
		withdraw := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, newHead)
		if withdraw.Epoch > epoch {
			break
		}
		amount = new(big.Int).Add(amount, withdraw.Amount)
	}
	return amount, newHead
}

// Total of all rewards since epoch
func (c *StakeHandler) GetDelegateWithdrawalPending(delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	bound := c.getDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr)
	amount := common.Big0
	for i := bound.Tail; i >= bound.Head; i-- {
		withdraw := c.getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr, i)
		if withdraw.Epoch <= epoch {
			break
		}
		amount = new(big.Int).Add(amount, withdraw.Amount)
	}
	return amount
}

func (c *StakeHandler) setDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr common.Address, bound *types.DelegateWithdrawalBound) error {
	value, err := rlp.EncodeToBytes(bound)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegateWithdrawalQueueBoundKey(delegaterAddr, validatorAddr), value)
	return nil
}

func (c *StakeHandler) getDelegateWithdrawalQueueBound(delegaterAddr, validatorAddr common.Address) *types.DelegateWithdrawalBound {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeDelegateWithdrawalQueueBoundKey(delegaterAddr, validatorAddr))

	if len(value) == 0 {
		return nil
	}
	var bound types.DelegateWithdrawalBound
	if err := rlp.DecodeBytes(value, &bound); nil == err {
		return &bound
	}
	return nil
}

func (c *StakeHandler) setDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr common.Address, index uint64, withdraw *types.DelegateWithdrawalItem) error {
	value, err := rlp.EncodeToBytes(withdraw)
	if nil != err {
		return ErrRlpEncode
	}
	c.evm.StateDB.SetState(c.contract.Address(), encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, index), value)
	return nil
}

func (c *StakeHandler) getDelegateWithdrawalQueueItem(delegaterAddr, validatorAddr common.Address, index uint64) *types.DelegateWithdrawalItem {
	value := c.evm.StateDB.GetState(c.contract.Address(), encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, index))

	if len(value) == 0 {
		return nil
	}
	var withdraw types.DelegateWithdrawalItem
	if err := rlp.DecodeBytes(value, &withdraw); nil == err {
		return &withdraw
	}
	return nil
}

// ----

func (c *StakeHandler) setSlashProcessed(handleEventId *big.Int, queue types.SlashValidatorWithdrawItemQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return ErrRlpEncode
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
