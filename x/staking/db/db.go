package db

import (
	"bytes"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
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
	priorityValidatorHeadKey                   = []byte("priorityValidatorHead")             // "priorityValidatorHead" => priorityValidator(head)
	priorityValidatorTailKey                   = []byte("priorityValidatorTail")             // "priorityValidatorTail" => priorityValidator(tail)
	priorityValidatorKeyPrefix                 = []byte("priorityValidator")                 // "priorityValidator":shares(stakeAmount+delegataionAmount):stakeEpoch:stakeIndex => priorityValidator{preKey, nextKey, validatorAddr}
	validatorKeyPrefix                         = []byte("validator")                         // "validator":validatorAddr => validator
	currentEpochKey                            = []byte("currentEpoch")                      // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey                            = []byte("currentRound")                      // "currentRound" => currentRound (It is a number)
	epochValidatorSharesSnapshotQueueKeyPrefix = []byte("epochValidatorSharesSnapshotQueue") // "epochValidatorSharesSnapshotQueue":epochId => []validatorSharesSnapshot  (For settlement epoch)
	roundValidatorSharesSnapshotQueueKeyPrefix = []byte("roundValidatorSharesSnapshotQueue") // "roundValidatorSharesSnapshotQueue":roundId => []validatorSharesSnapshot  (For consensus round)

	epochItemKeyPrefix = []byte("epochItem") // "epochItem":epochId => {preEpoch, nextEpoch, startBlock, endBlock, roundCount}
	roundItemKeyPrefix = []byte("roundItem") // "roundItem":roundId => {preRound, nextRound, startBlock, endBlock}

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

func encodeEpochValidatorSharesSnapshotQueueKey(epoch uint64) []byte {
	return append(epochValidatorSharesSnapshotQueueKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func encodeRoundValidatorSharesSnapshotQueueKey(round uint64) []byte {
	return append(roundValidatorSharesSnapshotQueueKeyPrefix, common.Uint64ToBytes(round)...)
}

func EncodeEpochItemKey(epoch uint64) []byte {
	return append(epochItemKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func EncodeRoundItemKey(round uint64) []byte {
	return append(roundItemKeyPrefix, common.Uint64ToBytes(round)...)
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

func GetValidatorPriority(db sdk.StateDB, epoch, stakeIndex uint64, shares *big.Int) *types.PriorityValidator {
	return getValidatorPriorityByKey(db, encodePriorityValidatorKey(epoch, stakeIndex, shares))
}

func getValidatorPriorityByKey(db sdk.StateDB, key []byte) *types.PriorityValidator {
	value := db.GetState(address.StakeHandlerAddress, key)
	if len(value) == 0 {
		return nil
	}
	var priority types.PriorityValidator
	if err := rlp.DecodeBytes(value, &priority); nil == err {
		return &priority
	}
	return nil
}

func setValidatorPriorityByKey(db sdk.StateDB, key []byte, priority *types.PriorityValidator) error {
	value, err := rlp.EncodeToBytes(priority)
	if nil != err {
		return ErrRlpEncode
	}

	db.SetState(address.StakeHandlerAddress, key, value)
	return nil
}

func SetValidatorPriority(db sdk.StateDB, validatorAddr common.Address, epoch, stakeIndex uint64, shares *big.Int) error {

	indexKey := EncodePriorityValidatorHeadKey()
	indexItem := getValidatorPriorityByKey(db, EncodePriorityValidatorHeadKey())

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

		next := getValidatorPriorityByKey(db, indexItem.NextKey)

		priority.UpdatePreKey(indexKey)
		priority.UpdateNextKey(indexItem.NextKey)
		indexItem.UpdateNextKey(priorityKey)
		next.UpdatePreKey(priorityKey)

		if err := setValidatorPriorityByKey(db, indexKey, indexItem); nil != err {
			return err
		}
		if err := setValidatorPriorityByKey(db, priorityKey, priority); nil != err {
			return err
		}
		if err := setValidatorPriorityByKey(db, priority.NextKey, next); nil != err {
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

					next := getValidatorPriorityByKey(db, indexItem.NextKey) // tail

					priority.UpdatePreKey(indexKey)
					priority.UpdateNextKey(indexItem.NextKey)
					indexItem.UpdateNextKey(priorityKey)
					next.UpdatePreKey(priorityKey)

					if err := setValidatorPriorityByKey(db, indexKey, indexItem); nil != err {
						return err
					}
					if err := setValidatorPriorityByKey(db, priorityKey, priority); nil != err {
						return err
					}
					if err := setValidatorPriorityByKey(db, priority.NextKey, next); nil != err {
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

				pre := getValidatorPriorityByKey(db, indexItem.PreKey)

				priority.UpdatePreKey(indexItem.PreKey)
				priority.UpdateNextKey(indexKey)
				indexItem.UpdatePreKey(priorityKey)
				pre.UpdateNextKey(priorityKey)

				if err := setValidatorPriorityByKey(db, priority.PreKey, pre); nil != err {
					return err
				}
				if err := setValidatorPriorityByKey(db, priorityKey, priority); nil != err {
					return err
				}
				if err := setValidatorPriorityByKey(db, indexKey, indexItem); nil != err {
					return err
				}

				break
			}
		}

		indexKey = indexItem.NextKey
		indexItem = getValidatorPriorityByKey(db, indexKey)
	}

	return nil
}

func RemoveValidatorPriority(db sdk.StateDB, epoch, stakeIndex uint64, shares *big.Int) error {

	priorityKey := encodePriorityValidatorKey(epoch, stakeIndex, shares)
	priority := getValidatorPriorityByKey(db, priorityKey)

	preKey := priority.PreKey
	nextKey := priority.NextKey

	pre := getValidatorPriorityByKey(db, preKey)
	next := getValidatorPriorityByKey(db, nextKey)

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
	db.SetState(address.StakeHandlerAddress, preKey, pvalue)
	db.SetState(address.StakeHandlerAddress, nextKey, nvalue)
	db.SetState(address.StakeHandlerAddress, priorityKey, []byte{})

	return nil
}

func RankPriorityValidatorIds(db sdk.StateDB, size uint64) types.ValidatorIds {

	arr := make(types.ValidatorIds, size)
	var count uint64 = 0

	headItem := getValidatorPriorityByKey(db, EncodePriorityValidatorHeadKey())
	item := getValidatorPriorityByKey(db, headItem.NextKey)

	for bytes.Compare(item.NextKey, EncodePriorityValidatorHeadKey()) != 0 && count < size { // not as tail  and count less size
		arr[count] = item.ValidatorAddr
		item = getValidatorPriorityByKey(db, item.NextKey)
		count++
	}
	return arr[:count]
}

func SetValidator(db sdk.StateDB, validatorAddr common.Address, validator *types.Validator) error {
	value, err := rlp.EncodeToBytes(validator)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address.StakeHandlerAddress, encodeValidatorKey(validatorAddr), value)
	return nil
}

func GetValidator(db sdk.StateDB, validatorAddr common.Address) *types.Validator {
	value := db.GetState(address.StakeHandlerAddress, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return nil
	}
	var validator types.Validator
	if err := rlp.DecodeBytes(value, &validator); nil == err {
		return &validator
	}
	return nil
}

func HasValidator(db sdk.StateDB, validatorAddr common.Address) bool {
	value := db.GetState(address.StakeHandlerAddress, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return false
	}
	return true
}
func HasNotValidator(db sdk.StateDB, validatorAddr common.Address) bool {
	return !HasValidator(db, validatorAddr)
}

func RemoveValidator(db sdk.StateDB, validatorAddr common.Address) {
	db.SetState(address.StakeHandlerAddress, encodeValidatorKey(validatorAddr), []byte{})
}

func SetCurrentEpoch(db sdk.StateDB, epoch uint64) {
	db.SetState(address.StakeHandlerAddress, currentEpochKey, common.Uint64ToBytes(epoch))
}

func GetCurrentEpoch(db sdk.StateDB) uint64 {
	value := db.GetState(address.StakeHandlerAddress, currentEpochKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func SetCurrentRound(db sdk.StateDB, round uint64) {
	db.SetState(address.StakeHandlerAddress, currentRoundKey, common.Uint64ToBytes(round))
}

func GetCurrentRound(db sdk.StateDBReader) uint64 {
	value := db.GetState(address.StakeHandlerAddress, currentRoundKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func SetEpochValidatorSharesSnapshotQueue(db sdk.StateDB, epoch uint64, queue types.ValidatorSortSnapshotQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address.StakeHandlerAddress, encodeEpochValidatorSharesSnapshotQueueKey(epoch), value)
	return nil
}

func GetEpochValidatorSharesSnapshotQueue(db sdk.StateDBReader, epoch uint64) types.ValidatorSortSnapshotQueue {
	value := db.GetState(address.StakeHandlerAddress, encodeEpochValidatorSharesSnapshotQueueKey(epoch))
	if len(value) == 0 {
		return nil
	}
	var queue types.ValidatorSortSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil == err {
		return queue
	}
	return nil
}

func GetEpochValidatorIds(db sdk.StateDBReader, epoch uint64) types.ValidatorIds {
	value := db.GetState(address.StakeHandlerAddress, encodeEpochValidatorSharesSnapshotQueueKey(epoch))
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

func SetRoundValidatorSharesSnapshotQueue(db sdk.StateDB, round uint64, queue types.ValidatorSortSnapshotQueue) error {
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address.StakeHandlerAddress, encodeRoundValidatorSharesSnapshotQueueKey(round), value)
	return nil
}

func GetRoundValidatorSharesSnapshotQueue(db sdk.StateDBReader, round uint64) types.ValidatorSortSnapshotQueue {
	value := db.GetState(address.StakeHandlerAddress, encodeRoundValidatorSharesSnapshotQueueKey(round))
	if len(value) == 0 {
		return nil
	}
	var queue types.ValidatorSortSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil == err {
		return queue
	}
	return nil
}

func GetRoundValidatorIds(db sdk.StateDBReader, round uint64) types.ValidatorIds {
	value := db.GetState(address.StakeHandlerAddress, encodeRoundValidatorSharesSnapshotQueueKey(round))
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

func AppendEpochItem(db sdk.StateDB, epoch uint64, startBlock, endBlock, roundCount uint64) error {
	epochItem := GetEpochItem(db, epoch)
	if epochItem.IsNotEmpty() {
		return ErrExist
	}
	epochItem = types.NewEpochItem(startBlock, endBlock, roundCount)
	return SetEpochItem(db, epoch, epochItem)
}

func GetEpochQueueSince(db sdk.StateDBReader, epoch, size uint64) types.EpochQueue {
	queue := types.NewEpochQueue(size)

	var count uint64 = 0

	currentEpoch := GetCurrentEpoch(db)
	index := epoch

	for index != currentEpoch+1 && count < size {
		queue[count] = GetEpochItem(db, index)
		index++
		count++
	}
	return queue[:count]
}

func GetEpochQueueUtil(db sdk.StateDBReader, epoch, size uint64) types.EpochQueue {
	queue := types.NewEpochQueue(size)

	var count uint64 = 0

	index := epoch
	for index != 0 && count < size {
		queue[count] = GetEpochItem(db, index)
		index--
		count++
	}
	return queue[:count]
}

func GetEpochQueueAndIndexSince(db sdk.StateDBReader, epoch, size uint64) ([]uint64, types.EpochQueue) {
	queue := types.NewEpochQueue(size)
	epochs := make([]uint64, size)

	var count uint64 = 0

	currentEpoch := GetCurrentEpoch(db)
	index := epoch

	for index != currentEpoch+1 && count < size {
		queue[count] = GetEpochItem(db, index)
		epochs[count] = index
		index++
		count++
	}
	return epochs[:count], queue[:count]
}

func GetEpochQueueAndIndexUtil(db sdk.StateDBReader, epoch, size uint64) ([]uint64, types.EpochQueue) {
	queue := types.NewEpochQueue(size)
	epochs := make([]uint64, size)

	var count uint64 = 0

	index := epoch
	for index != 0 && count < size {
		queue[count] = GetEpochItem(db, index)
		epochs[count] = index
		index--
		count++
	}
	return epochs[:count], queue[:count]
}

func SetEpochItem(db sdk.StateDB, epoch uint64, item *types.EpochItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address.StakeHandlerAddress, EncodeEpochItemKey(epoch), value)
	return nil
}

func GetEpochItem(db sdk.StateDB, epoch uint64) *types.EpochItem {
	value := db.GetState(address.StakeHandlerAddress, EncodeEpochItemKey(epoch))
	if len(value) == 0 {
		return nil
	}

	var item types.EpochItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

// ---

func AppendRoundItem(db sdk.StateDB, round uint64, startBlock, endBlock uint64) error {
	roundItem := GetRoundItem(db, round)
	if roundItem.IsNotEmpty() {
		return ErrExist
	}
	roundItem = types.NewRoundItem(startBlock, endBlock)
	return SetRoundItem(db, round, roundItem)
}

func GetRoundQueueSince(db sdk.StateDB, round, size uint64) types.RoundQueue {
	queue := types.NewRoundQueue(size)

	var count uint64 = 0

	currentEpoch := GetCurrentRound(db)
	index := round

	for index != currentEpoch+1 && count < size {
		queue[count] = GetRoundItem(db, index)
		index++
		count++
	}
	return queue[:count]
}

func GetRoundQueueUtil(db sdk.StateDB, round, size uint64) types.RoundQueue {
	queue := types.NewRoundQueue(size)

	var count uint64 = 0

	index := round
	for index != 0 && count < size {
		queue[count] = GetRoundItem(db, index)
		index--
		count++
	}
	return queue[:count]
}

func GetRoundQueueAndIndexSince(db sdk.StateDB, round, size uint64) ([]uint64, types.RoundQueue) {
	queue := types.NewRoundQueue(size)
	rounds := make([]uint64, size)

	var count uint64 = 0

	currentRound := GetCurrentRound(db)
	index := round

	for index != currentRound+1 && count < size {
		queue[count] = GetRoundItem(db, index)
		rounds[count] = index
		index++
		count++
	}
	return rounds[:count], queue[:count]
}

func GetRoundQueueAndIndexFromTail(db sdk.StateDB, round, size uint64) ([]uint64, types.RoundQueue) {
	queue := types.NewRoundQueue(size)
	rounds := make([]uint64, size)

	var count uint64 = 0

	index := round
	for index != 0 && count < size {
		queue[count] = GetRoundItem(db, index)
		rounds[count] = index
		index--
		count++
	}
	return rounds[:count], queue[:count]
}

func SetRoundItem(db sdk.StateDB, round uint64, item *types.RoundItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address.StakeHandlerAddress, EncodeRoundItemKey(round), value)
	return nil
}

func GetRoundItem(db sdk.StateDB, round uint64) *types.RoundItem {
	value := db.GetState(address.StakeHandlerAddress, EncodeRoundItemKey(round))
	if len(value) == 0 {
		return nil
	}

	var item types.RoundItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

func IncrementNumberOfBlocksForRoundValidator(db sdk.StateDB, validatorAddr common.Address, round, increment uint64) {
	value := db.GetState(address.StakeHandlerAddress, encodeNumberOfBlocksForRoundValidatorKey(validatorAddr, round))
	var v uint64
	if len(value) != 0 {
		v = common.BytesToUint64(value)
	}
	v += increment
	db.SetState(address.StakeHandlerAddress, encodeNumberOfBlocksForRoundValidatorKey(validatorAddr, round), common.Uint64ToBytes(v))

}
