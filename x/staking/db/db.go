package db

import (
	"bytes"
	"errors"
	stagedb "github.com/PlatONnetwork/AppChain-SDK/x/stage/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

var (
	ErrStoreFailed  = errors.New("store failed")
	ErrRlpEncode    = errors.New("rlp encode failed")
	ErrRlpDecode    = errors.New("rlp decode failed")
	ErrNotFound     = errors.New("not found")
	ErrExist        = errors.New("already exist")
	ErrMisMatching  = errors.New("mismatching")
	ErrInvalidValue = errors.New("invalid value")
)

var (
	validatorNonceKey = []byte("validatorNonce") // "validatorNonce" => nonce (It is a self increasing stake index number)

	priorityValidatorHeadKey   = []byte("priorityValidatorHead") // "priorityValidatorHead" => priorityValidator(head)
	priorityValidatorTailKey   = []byte("priorityValidatorTail") // "priorityValidatorTail" => priorityValidator(tail)
	priorityValidatorKeyPrefix = []byte("priorityValidator")     // "priorityValidator":shares(stakeAmount+delegataionAmount):stakeEpoch:stakeIndex => priorityValidator{preKey, nextKey, validatorAddr}
	validatorKeyPrefix         = []byte("validator")             // "validator":validatorAddr => validator

	delegationKeyPrefix                  = []byte("delegation")                  // "delegater":delegaterAddr:validatorAddr:stakeEpoch => delegation
	stakeWithdrawalQueueItemKeyPrefix    = []byte("stakeWithdrawalQueueItem")    // "stakeWithdrawalQueueItem":validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}
	delegateWithdrawalQueueItemKeyPrefix = []byte("delegateWithdrawalQueueItem") // "delegateWithdrawalQueueItem":delegaterAddr:validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}
	validatorDelegationRcKeyPrefix       = []byte("validatorDelegationRc")       // "validatorDelegationRc":validatorAddr:stakeEpoch => unStakeDelegationRcItem{preStakeEpoch, nextStakeEpoch, delegation count}
	slashProcessedKeyPrefix              = []byte("slashProcessed")              // "slashProcessed":handleEventId => []SlashValidatorWithdrawItem{validatorAddr, amount}

	epochValidatorSharesSnapshotQueueKeyPrefix = []byte("epochValidatorSharesSnapshotQueue") // "epochValidatorSharesSnapshotQueue":epochId => []validatorSharesSnapshot  (For settlement epoch)
	roundValidatorSharesSnapshotQueueKeyPrefix = []byte("roundValidatorSharesSnapshotQueue") // "roundValidatorSharesSnapshotQueue":roundId => []validatorSharesSnapshot  (For consensus round)

	numberOfBlocksForRoundValidatorKeyPrefix = []byte("numberOfBlocksForRoundValidator") // "numberOfBlocksForRoundValidator":validatorAddr:round => numberOfBlocks
)

func EncodePriorityValidatorHeadKey() []byte {
	return priorityValidatorHeadKey
}

func EncodePriorityValidatorTailKey() []byte {
	return priorityValidatorTailKey
}

func encodePriorityValidatorKey(epoch, stakeIndex uint64, shares *big.Int) []byte {

	sharesSub := new(big.Int).Sub(math.MaxBig104, shares)
	zeros := make([]byte, len(math.MaxBig104.Bytes()))
	sharesPriority := append(zeros, sharesSub.Bytes()...)

	stakeEpoch := common.Uint64ToBytes(epoch)
	index := common.Uint64ToBytes(stakeIndex)

	// some index of pivots
	keyPrefixSize := len(priorityValidatorKeyPrefix)
	appendSharePrioritySize := keyPrefixSize + len(sharesPriority)
	appendStakeEpochSize := appendSharePrioritySize + len(stakeEpoch)
	size := appendStakeEpochSize + len(index)

	// build key
	key := make([]byte, size)
	copy(key[:keyPrefixSize], priorityValidatorKeyPrefix)
	copy(key[keyPrefixSize:appendSharePrioritySize], sharesPriority)
	copy(key[appendSharePrioritySize:appendStakeEpochSize], stakeEpoch)
	copy(key[appendStakeEpochSize:], index)

	return key
}

func encodeValidatorKey(validatorAddr common.Address) []byte {
	return append(validatorKeyPrefix, validatorAddr.Bytes()...)
}

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

func encodeEpochValidatorSharesSnapshotQueueKey(epoch uint64) []byte {
	return append(epochValidatorSharesSnapshotQueueKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func encodeRoundValidatorSharesSnapshotQueueKey(round uint64) []byte {
	return append(roundValidatorSharesSnapshotQueueKeyPrefix, common.Uint64ToBytes(round)...)
}

func encodeNumberOfBlocksForRoundValidatorKey(validatorAddr common.Address, round uint64) []byte {

	validatorAddrBytes := validatorAddr.Bytes()
	roundBytes := common.Uint64ToBytes(round)

	keyPrefixSize := len(numberOfBlocksForRoundValidatorKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(roundBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], numberOfBlocksForRoundValidatorKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], roundBytes)

	return key
}

// ------------------------------------------------------ db methods ------------------------------------------------------

func GetValidatorPriority(db sdk.StateDBReader, addr common.Address, epoch, stakeIndex uint64, shares *big.Int) *types.PriorityValidator {
	return getValidatorPriorityByKey(db, addr, encodePriorityValidatorKey(epoch, stakeIndex, shares))
}

func getValidatorPriorityByKey(db sdk.StateDBReader, addr common.Address, key []byte) *types.PriorityValidator {
	value := db.GetState(addr, key)
	if len(value) == 0 {
		return nil
	}
	var priority types.PriorityValidator
	if err := rlp.DecodeBytes(value, &priority); nil == err {
		return &priority
	}
	return nil
}

func setValidatorPriorityByKey(db sdk.StateDB, addr common.Address, key []byte, priority *types.PriorityValidator) error {
	value, err := rlp.EncodeToBytes(priority)
	if nil != err {
		return ErrRlpEncode
	}

	db.SetState(addr, key, value)
	return nil
}

func SetValidatorPriority(db sdk.StateDB, addr, validatorAddr common.Address, epoch, stakeIndex uint64, shares *big.Int) error {

	indexKey := EncodePriorityValidatorHeadKey()
	indexItem := getValidatorPriorityByKey(db, addr, EncodePriorityValidatorHeadKey())

	priorityKey := encodePriorityValidatorKey(epoch, stakeIndex, shares)
	priority := types.NewPriorityValidator(
		[]byte{},
		[]byte{},
		validatorAddr,
	)
	// first insert
	if bytes.Compare(indexItem.PreKey, indexItem.NextKey) == 0 && bytes.Compare(indexItem.PreKey, EncodePriorityValidatorTailKey()) == 0 {

		// if  tail -> head -> tail -> head
		//
		// then: tail -> head -> priority -> tail -> head

		next := getValidatorPriorityByKey(db, addr, indexItem.NextKey)

		priority.UpdatePreKey(indexKey)
		priority.UpdateNextKey(indexItem.NextKey)
		indexItem.UpdateNextKey(priorityKey)
		next.UpdatePreKey(priorityKey)

		if err := setValidatorPriorityByKey(db, addr, indexKey, indexItem); nil != err {
			return err
		}
		if err := setValidatorPriorityByKey(db, addr, priorityKey, priority); nil != err {
			return err
		}
		if err := setValidatorPriorityByKey(db, addr, priority.NextKey, next); nil != err {
			return err
		}

		return nil
	}

	for bytes.Compare(indexItem.NextKey, EncodePriorityValidatorHeadKey()) != 0 { // not as  tail
		if bytes.Compare(indexItem.PreKey, EncodePriorityValidatorTailKey()) != 0 { // not as head

			if bytes.Compare(indexKey, priorityKey) == 0 {
				return ErrInvalidValue
			}

			if bytes.Compare(indexKey, priorityKey) < 0 {
				if bytes.Compare(indexItem.NextKey, EncodePriorityValidatorTailKey()) == 0 {
					// if  tail -> head -> index -> tail -> head
					// and index < priority
					//
					// then:  tail -> head -> index -> priority -> tail -> head

					next := getValidatorPriorityByKey(db, addr, indexItem.NextKey) // tail

					priority.UpdatePreKey(indexKey)
					priority.UpdateNextKey(indexItem.NextKey)
					indexItem.UpdateNextKey(priorityKey)
					next.UpdatePreKey(priorityKey)

					if err := setValidatorPriorityByKey(db, addr, indexKey, indexItem); nil != err {
						return err
					}
					if err := setValidatorPriorityByKey(db, addr, priorityKey, priority); nil != err {
						return err
					}
					if err := setValidatorPriorityByKey(db, addr, priority.NextKey, next); nil != err {
						return err
					}
					break
				}

				// if  tail -> head -> index -> next -> ... tail -> head
				// and index < priority
				// maybe next < priority
				//
				// then:  continue next become new index
			} else {
				// if  tail -> head -> index -> (next) tail -> head
				// and index > priority
				//
				// then:  tail -> head -> priority -> index -> (next)  tail -> head

				pre := getValidatorPriorityByKey(db, addr, indexItem.PreKey)

				priority.UpdatePreKey(indexItem.PreKey)
				priority.UpdateNextKey(indexKey)
				indexItem.UpdatePreKey(priorityKey)
				pre.UpdateNextKey(priorityKey)

				if err := setValidatorPriorityByKey(db, addr, priority.PreKey, pre); nil != err {
					return err
				}
				if err := setValidatorPriorityByKey(db, addr, priorityKey, priority); nil != err {
					return err
				}
				if err := setValidatorPriorityByKey(db, addr, indexKey, indexItem); nil != err {
					return err
				}

				break
			}
		}

		indexKey = indexItem.NextKey
		indexItem = getValidatorPriorityByKey(db, addr, indexKey)
	}

	return nil
}

func RemoveValidatorPriority(db sdk.StateDB, addr common.Address, epoch, stakeIndex uint64, shares *big.Int) error {

	priorityKey := encodePriorityValidatorKey(epoch, stakeIndex, shares)
	priority := getValidatorPriorityByKey(db, addr, priorityKey)

	preKey := priority.PreKey
	nextKey := priority.NextKey

	pre := getValidatorPriorityByKey(db, addr, preKey)
	next := getValidatorPriorityByKey(db, addr, nextKey)

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
	db.SetState(addr, preKey, pvalue)
	db.SetState(addr, nextKey, nvalue)
	db.SetState(addr, priorityKey, []byte{})

	return nil
}

func RankPriorityValidatorIds(db sdk.StateDBReader, addr common.Address, size uint64) types.ValidatorIds {

	arr := make(types.ValidatorIds, size)
	var count uint64 = 0

	headItem := getValidatorPriorityByKey(db, addr, EncodePriorityValidatorHeadKey())
	item := getValidatorPriorityByKey(db, addr, headItem.NextKey)

	for bytes.Compare(item.NextKey, EncodePriorityValidatorHeadKey()) != 0 && count < size { // not as tail  and count less size
		arr[count] = item.ValidatorAddr
		item = getValidatorPriorityByKey(db, addr, item.NextKey)
		count++
	}
	return arr[:count]
}

func SetValidator(db sdk.StateDB, addr, validatorAddr common.Address, validator *types.Validator) error {
	value, err := rlp.EncodeToBytes(validator)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeValidatorKey(validatorAddr), value)
	return nil
}

func GetValidator(db sdk.StateDBReader, addr, validatorAddr common.Address) *types.Validator {
	value := db.GetState(addr, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return nil
	}
	var validator types.Validator
	if err := rlp.DecodeBytes(value, &validator); nil == err {
		return &validator
	}
	return nil
}

func UpdateValidatorStatus(db sdk.StateDB, addr, validatorAddr common.Address, status types.ValidatorStatus) error {
	validator := GetValidator(db, addr, validatorAddr)
	if validator.IsEmpty() {
		return ErrInvalidValue
	}
	validator.AppendStatus(status)
	return SetValidator(db, addr, validatorAddr, validator)
}

func HasValidator(db sdk.StateDBReader, addr, validatorAddr common.Address) bool {
	value := db.GetState(addr, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return false
	}
	return true
}
func HasNotValidator(db sdk.StateDBReader, addr, validatorAddr common.Address) bool {
	return !HasValidator(db, addr, validatorAddr)
}

func RemoveValidator(db sdk.StateDB, addr, validatorAddr common.Address) {
	db.SetState(addr, encodeValidatorKey(validatorAddr), []byte{})
}

// -------

func IncrementValidatorNonce(db sdk.StateDB, addr common.Address) uint64 {
	nonce := GetValidatorNonce(db, addr)
	old := nonce
	nonce++
	db.SetState(addr, validatorNonceKey, common.Uint64ToBytes(nonce))
	return old
}

func GetValidatorNonce(db sdk.StateDBReader, addr common.Address) uint64 {
	value := db.GetState(addr, validatorNonceKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func SetDelegation(db sdk.StateDB, addr, delegaterAddr, validatorAddr common.Address, stakeEpoch uint64, delegation *types.Delegation) error {
	value, err := rlp.EncodeToBytes(delegation)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeDelegaterKey(delegaterAddr, validatorAddr, stakeEpoch), value)
	return nil
}

func RemoveDelegation(db sdk.StateDB, addr, delegaterAddr, validatorAddr common.Address, stakeEpoch uint64) {
	db.SetState(addr, encodeDelegaterKey(delegaterAddr, validatorAddr, stakeEpoch), []byte{})
}

func GetDelegation(db sdk.StateDBReader, addr, delegaterAddr, validatorAddr common.Address, stakeEpoch uint64) *types.Delegation {
	value := db.GetState(addr, encodeDelegaterKey(delegaterAddr, validatorAddr, stakeEpoch))
	if len(value) == 0 {
		return nil
	}
	var delegation types.Delegation
	if err := rlp.DecodeBytes(value, &delegation); nil == err {
		return &delegation
	}
	return nil
}

func AppendStakeWithdrawal(db sdk.StateDB, addr, validatorAddr common.Address, epoch uint64, amount *big.Int) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

	// first insert
	if indexItem.IsEmpty() {

		preEpoch := uint64(0)
		// tail -> head -> epoch -> tail -> head

		preItem := types.NewStakeWithdrawalItem(indexEpoch, epoch, common.Big0) // head
		epochItem := types.NewStakeWithdrawalItem(preEpoch, indexEpoch, amount) // item
		indexItem = types.NewStakeWithdrawalItem(epoch, preEpoch, common.Big0)  // tail

		if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, preEpoch, preItem); nil != err {
			return err
		}
		if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
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
			if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			break
		} else if indexEpoch < epoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			next := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexItem.NextEpoch)
			epochItem := types.NewStakeWithdrawalItem(indexEpoch, indexItem.NextEpoch, amount)

			indexItem.UpdateNextEpoch(epoch)
			next.UpdatePreEpoch(epoch)

			if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, epoch, epochItem); nil != err {
				return err
			}
			if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, epochItem.NextEpoch, next); nil != err {
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

				pre := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexItem.PreEpoch) // head
				epochItem := types.NewStakeWithdrawalItem(indexItem.PreEpoch, epoch, amount)

				pre.UpdateNextEpoch(epoch)
				indexItem.UpdatePreEpoch(epoch)

				if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, epochItem.PreEpoch, pre); nil != err {
					return err
				}
				if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, epoch, epochItem); nil != err {
					return err
				}
				if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
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
		indexItem = GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)
	}
	return nil
}

func GetStakeWithdrawal(db sdk.StateDBReader, addr, validatorAddr common.Address) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

	for indexItem.NextEpoch != uint64(0) { // not as tail
		amount = new(big.Int).Add(amount, indexItem.Amount)
		indexEpoch = indexItem.NextEpoch
		indexItem = GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)
	}
	return amount
}

func GetStakeWithdrawalByEpoch(db sdk.StateDBReader, addr, validatorAddr common.Address, epoch uint64) *big.Int {
	item := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, epoch)
	if nil != item {
		return item.Amount
	}
	return common.Big0
}

func GetStakeWithdrawalLastEpoch(db sdk.StateDBReader, addr, validatorAddr common.Address) uint64 {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return 0
	}

	return indexItem.PreEpoch
}

// Total of all rewards until epoch
func GetStakeWithdrawable(db sdk.StateDBReader, addr, validatorAddr common.Address, epoch uint64) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

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
		indexItem = GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)
	}
	return amount
}

func ApplyStakeWithdrawable(db sdk.StateDB, addr, validatorAddr common.Address, epoch uint64) (*big.Int, error) {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

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
			removeStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)
		}

		indexEpoch = indexItem.NextEpoch
		indexItem = GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)
	}

	// remove head and tail
	if indexItem.NextEpoch == uint64(0) { // as tail
		removeStakeWithdrawalQueueItem(db, addr, validatorAddr, 0)
		removeStakeWithdrawalQueueItem(db, addr, validatorAddr, math.MaxUint64)
	} else { // update start index
		head := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, uint64(0))
		head.UpdateNextEpoch(indexEpoch)
		if err := SetStakeWithdrawalQueueItem(db, addr, validatorAddr, uint64(0), head); nil != err { // update head
			return common.Big0, err
		}
	}

	return amount, nil
}

func CleanStakeWithdrawable(db sdk.StateDB, addr, validatorAddr common.Address) {

	indexEpoch := uint64(0)
	indexItem := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

	// empty queue
	if indexItem.IsEmpty() {
		return
	}
	for indexItem.IsNotEmpty() {

		removeStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

		indexEpoch = indexItem.NextEpoch
		indexItem = GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)
	}

	return
}

// Total of all rewards since epoch
func GetStakeWithdrawalPending(db sdk.StateDBReader, addr, validatorAddr common.Address, epoch uint64) *big.Int {

	amount := common.Big0

	indexEpoch := uint64(math.MaxUint64)
	indexItem := GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)

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
		indexItem = GetStakeWithdrawalQueueItem(db, addr, validatorAddr, indexEpoch)
	}
	return amount
}

func SetStakeWithdrawalQueueItem(db sdk.StateDB, addr, validatorAddr common.Address, epoch uint64, item *types.StakeWithdrawalItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeStakeWithdrawalQueueItemKey(validatorAddr, epoch), value)
	return nil
}

func removeStakeWithdrawalQueueItem(db sdk.StateDB, addr, validatorAddr common.Address, epoch uint64) {
	db.SetState(addr, encodeStakeWithdrawalQueueItemKey(validatorAddr, epoch), []byte{})
}

func GetStakeWithdrawalQueueItem(db sdk.StateDBReader, addr, validatorAddr common.Address, epoch uint64) *types.StakeWithdrawalItem {
	value := db.GetState(addr, encodeStakeWithdrawalQueueItemKey(validatorAddr, epoch))

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

func AppendDelegateWithdrawal(db sdk.StateDB, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64, amount *big.Int) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)

	// first insert
	if indexItem.IsEmpty() {

		preEpoch := uint64(0)
		// tail -> head -> epoch -> tail -> head

		preItem := types.NewDelegateWithdrawalItem(indexEpoch, epoch, common.Big0) // head
		epochItem := types.NewDelegateWithdrawalItem(preEpoch, indexEpoch, amount) // item
		indexItem = types.NewDelegateWithdrawalItem(epoch, preEpoch, common.Big0)  // tail

		if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, preEpoch, preItem); nil != err {
			return err
		}
		if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
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
			if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			break
		} else if indexEpoch < epoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			next := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexItem.NextEpoch)
			epochItem := types.NewDelegateWithdrawalItem(indexEpoch, indexItem.NextEpoch, amount)

			indexItem.UpdateNextEpoch(epoch)
			next.UpdatePreEpoch(epoch)

			if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
				return err
			}
			if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, epochItem.NextEpoch, next); nil != err {
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

				pre := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexItem.PreEpoch) // head
				epochItem := types.NewDelegateWithdrawalItem(indexItem.PreEpoch, epoch, amount)

				pre.UpdateNextEpoch(epoch)
				indexItem.UpdatePreEpoch(epoch)

				if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, epochItem.PreEpoch, pre); nil != err {
					return err
				}
				if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, epoch, epochItem); nil != err {
					return err
				}
				if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch, indexItem); nil != err {
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
		indexItem = getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)
	}
	return nil
}

func GetDelegateWithdrawal(db sdk.StateDBReader, addr common.Address, delegaterAddr, validatorAddr common.Address) *big.Int {
	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)

	for indexItem.NextEpoch != uint64(0) { // not as tail
		amount = new(big.Int).Add(amount, indexItem.Amount)
		indexEpoch = indexItem.NextEpoch
		indexItem = getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)
	}
	return amount
}

func GetDelegateWithdrawalByEpoch(db sdk.StateDBReader, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	item := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, epoch)
	if nil != item {
		return item.Amount
	}
	return common.Big0
}

// Total of all rewards until epoch
func GetDelegateWithdrawable(db sdk.StateDBReader, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)

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
		indexItem = getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)
	}
	return amount
}

func ApplyDelegateWithdrawable(db sdk.StateDB, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64) (*big.Int, error) {

	amount := common.Big0

	indexEpoch := uint64(0)
	indexItem := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)

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
			removeDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)
		}

		indexEpoch = indexItem.NextEpoch
		indexItem = getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)
	}

	// remove head and tail
	if indexItem.NextEpoch == uint64(0) {
		removeDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, 0)
		removeDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, math.MaxUint64)
	} else { // update start index
		head := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, uint64(0))
		head.UpdateNextEpoch(indexEpoch)
		if err := setDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, uint64(0), head); nil != err {
			return common.Big0, err
		}
	}

	return amount, nil
}

// Total of all rewards since epoch
func GetDelegateWithdrawalPending(db sdk.StateDBReader, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64) *big.Int {
	amount := common.Big0

	indexEpoch := uint64(math.MaxUint64)
	indexItem := getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)

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
		indexItem = getDelegateWithdrawalQueueItem(db, addr, delegaterAddr, validatorAddr, indexEpoch)
	}
	return amount
}

func setDelegateWithdrawalQueueItem(db sdk.StateDB, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64, item *types.DelegateWithdrawalItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch), value)
	return nil
}

func removeDelegateWithdrawalQueueItem(db sdk.StateDB, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64) {
	db.SetState(addr, encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch), []byte{})
}

func getDelegateWithdrawalQueueItem(db sdk.StateDBReader, addr common.Address, delegaterAddr, validatorAddr common.Address, epoch uint64) *types.DelegateWithdrawalItem {
	value := db.GetState(addr, encodeDelegateWithdrawalQueueItemKey(delegaterAddr, validatorAddr, epoch))

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

func AppendValidatorDelegationRc(db sdk.StateDB, addr, validatorAddr common.Address, epoch, rc uint64) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := getValidatorDelegationRcItem(db, addr, validatorAddr, indexEpoch)

	// first insert
	if indexItem.IsEmpty() {

		preEpoch := uint64(0)
		// tail -> head -> epoch -> tail -> head

		preItem := types.NewValidatorDelegationRcItem(indexEpoch, epoch, 0)       // head
		epochItem := types.NewValidatorDelegationRcItem(preEpoch, indexEpoch, rc) // item
		indexItem = types.NewValidatorDelegationRcItem(epoch, preEpoch, 0)        // tail

		if err := setValidatorDelegationRcItem(db, addr, validatorAddr, preEpoch, preItem); nil != err {
			return err
		}
		if err := setValidatorDelegationRcItem(db, addr, validatorAddr, epoch, epochItem); nil != err {
			return err
		}
		if err := setValidatorDelegationRcItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
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
			if err := setValidatorDelegationRcItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			break
		} else if indexEpoch < epoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			nextItem := getValidatorDelegationRcItem(db, addr, validatorAddr, indexItem.NextStakeEpoch)
			epochItem := types.NewValidatorDelegationRcItem(indexEpoch, indexItem.NextStakeEpoch, rc)

			indexItem.UpdateNextStakeEpoch(epoch)
			nextItem.UpdatePreStakeEpoch(epoch)

			if err := setValidatorDelegationRcItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := setValidatorDelegationRcItem(db, addr, validatorAddr, epoch, epochItem); nil != err {
				return err
			}
			if err := setValidatorDelegationRcItem(db, addr, validatorAddr, epochItem.NextStakeEpoch, nextItem); nil != err {
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

				preItem := getValidatorDelegationRcItem(db, addr, validatorAddr, indexItem.PreStakeEpoch) // head
				epochItem := types.NewValidatorDelegationRcItem(indexItem.PreStakeEpoch, epoch, rc)

				preItem.UpdateNextStakeEpoch(epoch)
				indexItem.UpdatePreStakeEpoch(epoch)

				if err := setValidatorDelegationRcItem(db, addr, validatorAddr, epochItem.PreStakeEpoch, preItem); nil != err {
					return err
				}
				if err := setValidatorDelegationRcItem(db, addr, validatorAddr, epoch, epochItem); nil != err {
					return err
				}
				if err := setValidatorDelegationRcItem(db, addr, validatorAddr, indexEpoch, indexItem); nil != err {
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
		indexItem = getValidatorDelegationRcItem(db, addr, validatorAddr, indexEpoch)
	}
	return nil
}

func GetValidatorDelegationRcPending(db sdk.StateDBReader, addr, validatorAddr common.Address, size uint64) types.ValidatorDelegationRcQueue {
	indexStakeEpoch := uint64(0)
	indexItem := getValidatorDelegationRcItem(db, addr, validatorAddr, indexStakeEpoch)

	if indexItem.IsEmpty() {
		return nil
	}
	queue := types.NewValidatorDelegationRcQueue(size)
	count := uint64(0)
	for indexItem.NextStakeEpoch != uint64(0) && count < size { // not tail or count less size
		queue[count] = indexItem

		indexStakeEpoch = indexItem.NextStakeEpoch
		indexItem = getValidatorDelegationRcItem(db, addr, validatorAddr, indexStakeEpoch)
		count++
	}
	return queue[:count]
}

func GetValidatorDelegationRcPendingAndEpoch(db sdk.StateDBReader, addr, validatorAddr common.Address, size uint64) ([]uint64, types.ValidatorDelegationRcQueue) {
	indexStakeEpoch := uint64(0)
	indexItem := getValidatorDelegationRcItem(db, addr, validatorAddr, indexStakeEpoch)

	if indexItem.IsEmpty() {
		return nil, nil
	}

	// the first one
	indexStakeEpoch = indexItem.NextStakeEpoch
	indexItem = getValidatorDelegationRcItem(db, addr, validatorAddr, indexStakeEpoch)

	indexStakeEpochQueue := make([]uint64, size)
	queue := types.NewValidatorDelegationRcQueue(size)
	count := uint64(0)
	for indexItem.NextStakeEpoch != uint64(0) && count < size { // not tail or count less size
		indexStakeEpochQueue[count] = indexStakeEpoch
		queue[count] = indexItem

		indexStakeEpoch = indexItem.NextStakeEpoch
		indexItem = getValidatorDelegationRcItem(db, addr, validatorAddr, indexStakeEpoch)
		count++
	}
	return indexStakeEpochQueue[:count], queue[:count]
}

func getValidatorDelegationRcItem(db sdk.StateDBReader, addr, validatorAddr common.Address, stakeEpoch uint64) *types.ValidatorDelegationRcItem {
	value := db.GetState(addr, encodeValidatorDelegationRcKey(validatorAddr, stakeEpoch))

	if len(value) != 0 {
		return nil
	}
	var item types.ValidatorDelegationRcItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

func setValidatorDelegationRcItem(db sdk.StateDB, addr, validatorAddr common.Address, stakeEpoch uint64, item *types.ValidatorDelegationRcItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeValidatorDelegationRcKey(validatorAddr, stakeEpoch), value)
	return nil
}

func removeValidatorDelegationRcItem(db sdk.StateDB, addr, validatorAddr common.Address, stakeEpoch uint64) {
	db.SetState(addr, encodeValidatorDelegationRcKey(validatorAddr, stakeEpoch), []byte{})
}

func ReleaseValidatorDelegationRcItem(db sdk.StateDB, addr, validatorAddr common.Address, stakeEpoch, decrement uint64) error {
	item := getValidatorDelegationRcItem(db, addr, validatorAddr, stakeEpoch)
	if nil == item {
		return ErrNotFound
	}

	item.DecrementRc(decrement)

	// change index
	if item.Rc == 0 {
		pre := getValidatorDelegationRcItem(db, addr, validatorAddr, item.PreStakeEpoch)
		next := getValidatorDelegationRcItem(db, addr, validatorAddr, item.NextStakeEpoch)

		// remove the last one   tail -> head -> lastone(remove) -> tail -> head
		if pre.PreStakeEpoch == math.MaxUint64 && next.NextStakeEpoch == 0 {
			removeValidatorDelegationRcItem(db, addr, validatorAddr, item.PreStakeEpoch)
			removeValidatorDelegationRcItem(db, addr, validatorAddr, item.NextStakeEpoch)
		} else {
			pre.UpdateNextStakeEpoch(item.NextStakeEpoch)
			next.UpdatePreStakeEpoch(item.PreStakeEpoch)
			if err := setValidatorDelegationRcItem(db, addr, validatorAddr, item.PreStakeEpoch, pre); nil != err {
				return err
			}
			if err := setValidatorDelegationRcItem(db, addr, validatorAddr, item.NextStakeEpoch, next); nil != err {
				return err
			}
		}
		removeValidatorDelegationRcItem(db, addr, validatorAddr, stakeEpoch)
	} else {
		if err := setValidatorDelegationRcItem(db, addr, validatorAddr, stakeEpoch, item); nil != err {
			return err
		}
	}
	return nil
}

func GetValidatorDelegationRc(db sdk.StateDBReader, addr, validatorAddr common.Address, stakeEpoch uint64) uint64 {
	item := getValidatorDelegationRcItem(db, addr, validatorAddr, stakeEpoch)
	if nil == item {
		return 0
	}
	return item.Rc
}

// ----

func SetSlashProcessed(db sdk.StateDB, addr common.Address, handleEventId *big.Int, queue types.SlashValidatorWithdrawItemQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeSlashProcessedKey(handleEventId), value)
	return nil
}

func GetSlashProcessed(db sdk.StateDBReader, addr common.Address, handleEventId *big.Int) types.SlashValidatorWithdrawItemQueue {
	value := db.GetState(addr, encodeSlashProcessedKey(handleEventId))
	if len(value) == 0 {
		return nil
	}

	var queue types.SlashValidatorWithdrawItemQueue
	if err := rlp.DecodeBytes(value, &queue); nil == err {
		return queue
	}
	return nil
}

func HasSlashProcessed(db sdk.StateDBReader, addr common.Address, handleEventId *big.Int) bool {
	return GetSlashProcessed(db, addr, handleEventId).IsNotEmpty()
}

func HasNotSlashProcessed(db sdk.StateDBReader, addr common.Address, handleEventId *big.Int) bool {
	return !HasSlashProcessed(db, addr, handleEventId)
}

// ---------

func SetEpochValidatorSharesSnapshotQueue(db sdk.StateDB, addr common.Address, epoch uint64, queue types.ValidatorSortSnapshotQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeEpochValidatorSharesSnapshotQueueKey(epoch), value)
	return nil
}

func GetEpochValidatorSharesSnapshotQueue(db sdk.StateDBReader, addr common.Address, epoch uint64) types.ValidatorSortSnapshotQueue {
	value := db.GetState(addr, encodeEpochValidatorSharesSnapshotQueueKey(epoch))
	if len(value) == 0 {
		return nil
	}
	var queue types.ValidatorSortSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil == err {
		return queue
	}
	return nil
}

func GetEpochValidatorIds(db sdk.StateDBReader, addr common.Address, epoch uint64) types.ValidatorIds {
	value := db.GetState(addr, encodeEpochValidatorSharesSnapshotQueueKey(epoch))
	if len(value) == 0 {
		return nil
	}

	var queue types.ValidatorSortSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil != err {
		return nil
	}

	ids := make(types.ValidatorIds, len(queue))

	for i, v := range queue {
		ids[i] = v.ValidatorAddr
	}

	return ids
}

func SetRoundValidatorSharesSnapshotQueue(db sdk.StateDB, addr common.Address, round uint64, queue types.ValidatorSortSnapshotQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeRoundValidatorSharesSnapshotQueueKey(round), value)
	return nil
}

func GetRoundValidatorSharesSnapshotQueue(db sdk.StateDBReader, addr common.Address, round uint64) types.ValidatorSortSnapshotQueue {
	value := db.GetState(addr, encodeRoundValidatorSharesSnapshotQueueKey(round))
	if len(value) == 0 {
		return nil
	}
	var queue types.ValidatorSortSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil == err {
		return queue
	}
	return nil
}

func GetRoundValidatorIds(db sdk.StateDBReader, addr common.Address, round uint64) types.ValidatorIds {
	value := db.GetState(addr, encodeRoundValidatorSharesSnapshotQueueKey(round))
	if len(value) == 0 {
		return nil
	}
	var queue types.ValidatorSortSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil != err {
		return nil
	}

	ids := make(types.ValidatorIds, len(queue))

	for i, v := range queue {
		ids[i] = v.ValidatorAddr
	}

	return ids
}

// ----

func IncrementNumberOfBlocksForRoundValidator(db sdk.StateDB, addr, validatorAddr common.Address, round, increment uint64) {
	number := GetNumberOfBlocksForRoundValidator(db, addr, validatorAddr, round)
	number += increment
	db.SetState(addr, encodeNumberOfBlocksForRoundValidatorKey(validatorAddr, round), common.Uint64ToBytes(number))
}

func GetNumberOfBlocksForRoundValidator(db sdk.StateDBReader, addr, validatorAddr common.Address, round uint64) uint64 {
	value := db.GetState(addr, encodeNumberOfBlocksForRoundValidatorKey(validatorAddr, round))
	var number uint64
	if len(value) != 0 {
		number = common.BytesToUint64(value)
	}
	return number
}

func getNumberOfBlocksForRoundValidatorsMap(db sdk.StateDBReader, addr common.Address, validatorAddrQueue []common.Address, round uint64) map[common.Address]uint64 {
	cache := make(map[common.Address]uint64, 0)

	for _, validatorAddr := range validatorAddrQueue {
		cache[validatorAddr] = GetNumberOfBlocksForRoundValidator(db, addr, validatorAddr, round)
	}
	return cache
}

func HasLowBlocksValidator(db sdk.StateDBReader, addr common.Address, minRoundValidatorBlockNumber uint64) bool {
	currentRound := stagedb.GetCurrentRound(db, addr)
	if currentRound == 1 {
		return false
	}

	previousRound := currentRound - 1
	previousRoundValidatorAddrQueue := GetRoundValidatorIds(db, addr, previousRound)

	cache := getNumberOfBlocksForRoundValidatorsMap(db, addr, previousRoundValidatorAddrQueue, previousRound)

	for _, number := range cache {
		if number < minRoundValidatorBlockNumber {
			return true
		}
	}
	return false
}

func HasNotLowBlocksValidator(db sdk.StateDBReader, addr common.Address, minRoundValidatorBlockNumber uint64) bool {
	return !HasLowBlocksValidator(db, addr, minRoundValidatorBlockNumber)
}

func CheckLowBlocksValidatorForPreviousRound(db sdk.StateDBReader, addr common.Address, minRoundValidatorBlockNumber uint64) types.ValidatorIds {
	currentRound := stagedb.GetCurrentRound(db, addr)
	if currentRound == 1 {
		return nil
	}

	previousRound := currentRound - 1
	previousRoundValidatorAddrQueue := GetRoundValidatorIds(db, addr, previousRound)

	cache := getNumberOfBlocksForRoundValidatorsMap(db, addr, previousRoundValidatorAddrQueue, previousRound)
	validators := make(types.ValidatorIds, 0)

	for validatorAddr, number := range cache {
		if number < minRoundValidatorBlockNumber {
			validators = append(validators, validatorAddr)
		}
	}
	return validators
}
