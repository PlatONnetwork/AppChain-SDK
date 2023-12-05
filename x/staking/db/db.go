package db

import (
	"bytes"
	"errors"
	stakecommon "github.com/PlatONnetwork/AppChain-SDK/x/constants"
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
	priorityValidatorHeadKey   = []byte("priorityValidatorHead") // "priorityValidatorHead" => priorityValidator(head)
	priorityValidatorTailKey   = []byte("priorityValidatorTail") // "priorityValidatorTail" => priorityValidator(tail)
	priorityValidatorKeyPrefix = []byte("priorityValidator")     // "priorityValidator":shares(stakeAmount+delegataionAmount):stakeEpoch:stakeIndex => priorityValidator{preKey, nextKey, validatorAddr}
	validatorKeyPrefix         = []byte("validator")             // "validator":validatorAddr => validator

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

func SetValidatorPriority(db sdk.StateDB, addr common.Address, validatorAddr common.Address, epoch, stakeIndex uint64, shares *big.Int) error {

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

func SetValidator(db sdk.StateDB, addr common.Address, validatorAddr common.Address, validator *types.Validator) error {
	value, err := rlp.EncodeToBytes(validator)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeValidatorKey(validatorAddr), value)
	return nil
}

func GetValidator(db sdk.StateDBReader, addr common.Address, validatorAddr common.Address) *types.Validator {
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

func HasValidator(db sdk.StateDBReader, addr common.Address, validatorAddr common.Address) bool {
	value := db.GetState(addr, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return false
	}
	return true
}
func HasNotValidator(db sdk.StateDBReader, addr common.Address, validatorAddr common.Address) bool {
	return !HasValidator(db, addr, validatorAddr)
}

func RemoveValidator(db sdk.StateDB, addr common.Address, validatorAddr common.Address) {
	db.SetState(addr, encodeValidatorKey(validatorAddr), []byte{})
}

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

func IncrementNumberOfBlocksForRoundValidator(db sdk.StateDB, addr common.Address, validatorAddr common.Address, round, increment uint64) {
	number := GetNumberOfBlocksForRoundValidator(db, addr, validatorAddr, round)
	number += increment
	db.SetState(addr, encodeNumberOfBlocksForRoundValidatorKey(validatorAddr, round), common.Uint64ToBytes(number))
}

func GetNumberOfBlocksForRoundValidator(db sdk.StateDBReader, addr common.Address, validatorAddr common.Address, round uint64) uint64 {
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

func HasLowBlocksValidator(db sdk.StateDB, addr common.Address) bool {
	currentRound := stagedb.GetCurrentRound(db, addr)
	if currentRound == 1 {
		return false
	}

	previousRound := currentRound - 1
	previousRoundValidatorAddrQueue := GetRoundValidatorIds(db, addr, previousRound)

	cache := getNumberOfBlocksForRoundValidatorsMap(db, addr, previousRoundValidatorAddrQueue, previousRound)

	for _, number := range cache {
		if number < stakecommon.MIN_ROUND_VALIDATOR_BLOCK_NUMBER {
			return true
		}
	}
	return false
}

func HasNotLowBlocksValidator(db sdk.StateDB, addr common.Address) bool {
	return !HasLowBlocksValidator(db, addr)
}

func CheckLowBlocksValidator(db sdk.StateDB, addr common.Address) types.ValidatorIds {
	currentRound := stagedb.GetCurrentRound(db, addr)
	if currentRound == 1 {
		return nil
	}

	previousRound := currentRound - 1
	previousRoundValidatorAddrQueue := GetRoundValidatorIds(db, addr, previousRound)

	cache := getNumberOfBlocksForRoundValidatorsMap(db, addr, previousRoundValidatorAddrQueue, previousRound)
	validators := make(types.ValidatorIds, 0)

	for validatorAddr, number := range cache {
		if number < stakecommon.MIN_ROUND_VALIDATOR_BLOCK_NUMBER {
			validators = append(validators, validatorAddr)
		}
	}
	return validators
}
