package db

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

var (
	ErrStoreFailed  = errors.New("store failed")
	ErrRlpEncode    = errors.New("rlp encode failed")
	ErrRlpDecode    = errors.New("rlp decode failed")
	ErrNotFound     = errors.New("not found")
	ErrMisMatching  = errors.New("mismatching")
	ErrInvalidValue = errors.New("invalid value")
)

var (
	validatorKeyPrefix         = []byte("validator")         // "validator":validatorAddr => validator
	currentEpochKey            = []byte("currentEpoch")      // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey            = []byte("currentRound")      // "currentRound" => currentRound (It is a number)
	epochValidatorIdsKeyPrefix = []byte("epochValidatorIds") // "epochValidatorIds":epochId => []validatorAddr  (For settlement epoch)
	roundValidatorIdsKeyPrefix = []byte("roundValidatorIds") // "roundValidatorIds":roundId => []validatorAddr  (For consensus round)

	epochItemKeyPrefix = []byte("epochItem") // "epochItem":epochId => {preEpoch, nextEpoch, startBlock, endBlock}    todo maybe add block root range start and end
	roundItemKeyPrefix = []byte("roundItem") // "roundItem":roundId => {preRound, nextRound, startBlock, endBlock}

)

func encodeValidatorKey(validatorAddr common.Address) []byte {
	return append(validatorKeyPrefix, validatorAddr.Bytes()...)
}

func encodeEpochValidatorIdsKey(epoch uint64) []byte {
	return append(epochValidatorIdsKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func encodeRoundValidatorIdsKey(round uint64) []byte {
	return append(roundValidatorIdsKeyPrefix, common.Uint64ToBytes(round)...)
}

func EncodeEpochItemKey(epoch uint64) []byte {
	return append(epochItemKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func EncodeRoundItemKey(round uint64) []byte {
	return append(roundItemKeyPrefix, common.Uint64ToBytes(round)...)
}

// ------------------------------------------------------ db methods ------------------------------------------------------

func SetValidator(db sdk.StateDB, address common.Address, validatorAddr common.Address, validator *types.Validator) error {
	value, err := rlp.EncodeToBytes(validator)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address, encodeValidatorKey(validatorAddr), value)
	return nil
}

func GetValidator(db sdk.StateDB, address common.Address, validatorAddr common.Address) *types.Validator {
	value := db.GetState(address, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return nil
	}
	var validator types.Validator
	if err := rlp.DecodeBytes(value, &validator); nil == err {
		return &validator
	}
	return nil
}

func HasValidator(db sdk.StateDB, address common.Address, validatorAddr common.Address) bool {
	value := db.GetState(address, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return false
	}
	return true
}
func HasNotValidator(db sdk.StateDB, address common.Address, validatorAddr common.Address) bool {
	return !HasValidator(db, address, validatorAddr)
}

func RemoveValidator(db sdk.StateDB, address common.Address, validatorAddr common.Address) {
	db.SetState(address, encodeValidatorKey(validatorAddr), []byte{})
}

func SetCurrentEpoch(db sdk.StateDB, address common.Address, epoch uint64) {
	db.SetState(address, currentEpochKey, common.Uint64ToBytes(epoch))
}

func GetCurrentEpoch(db sdk.StateDB, address common.Address) uint64 {
	value := db.GetState(address, currentEpochKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func SetCurrentRound(db sdk.StateDB, address common.Address, round uint64) {
	db.SetState(address, currentRoundKey, common.Uint64ToBytes(round))
}

func GetCurrentRound(db sdk.StateDB, address common.Address) uint64 {
	value := db.GetState(address, currentRoundKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func SetEpochValidatorIds(db sdk.StateDB, address common.Address, epoch uint64, ids types.ValidatorIds) error {
	value, err := rlp.EncodeToBytes(ids)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address, encodeEpochValidatorIdsKey(epoch), value)
	return nil
}

func GetEpochValidatorIds(db sdk.StateDB, address common.Address, epoch uint64) types.ValidatorIds {
	value := db.GetState(address, encodeEpochValidatorIdsKey(epoch))
	if len(value) == 0 {
		return nil
	}
	var ids types.ValidatorIds
	if err := rlp.DecodeBytes(value, &ids); nil == err {
		return ids
	}
	return nil
}

func SetRoundValidatorIds(db sdk.StateDB, address common.Address, round uint64, ids types.ValidatorIds) error {
	value, err := rlp.EncodeToBytes(ids)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address, encodeRoundValidatorIdsKey(round), value)
	return nil
}

func GetRoundValidatorIds(db sdk.StateDB, address common.Address, round uint64) types.ValidatorIds {
	value := db.GetState(address, encodeRoundValidatorIdsKey(round))
	if len(value) == 0 {
		return nil
	}
	var ids types.ValidatorIds
	if err := rlp.DecodeBytes(value, &ids); nil == err {
		return ids
	}
	return nil
}

// ----

func AppendEpochItem(db sdk.StateDB, address common.Address, epoch uint64, startBlock, endBlock, roundCount uint64) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := GetEpochItem(db, address, indexEpoch)

	// range from tail to head
	for indexItem.PreEpoch != math.MaxUint64 { // not as head
		if indexEpoch == epoch {
			// pre -> index(epoch) -> next
			// index == epoch

			return ErrInvalidValue
		}

		if indexEpoch < epoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			next := GetEpochItem(db, address, indexItem.NextEpoch)
			epochItem := types.NewEpochItem(indexEpoch, indexItem.NextEpoch, startBlock, endBlock, roundCount)

			indexItem.UpdateNextEpoch(epoch)
			next.UpdatePreEpoch(epoch)

			if err := SetEpochItem(db, address, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := SetEpochItem(db, address, epoch, epochItem); nil != err {
				return err
			}
			if err := SetEpochItem(db, address, epochItem.NextEpoch, next); nil != err {
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

				pre := GetEpochItem(db, address, indexItem.PreEpoch) // head
				epochItem := types.NewEpochItem(indexItem.PreEpoch, epoch, startBlock, endBlock, roundCount)

				pre.UpdateNextEpoch(epoch)
				indexItem.UpdatePreEpoch(epoch)

				if err := SetEpochItem(db, address, epochItem.PreEpoch, pre); nil != err {
					return err
				}
				if err := SetEpochItem(db, address, epoch, epochItem); nil != err {
					return err
				}
				if err := SetEpochItem(db, address, indexEpoch, indexItem); nil != err {
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
		indexItem = GetEpochItem(db, address, indexEpoch)
	}
	return nil
}

func GetLastEpochItem(db sdk.StateDB, address common.Address) *types.EpochItem {
	tail := GetEpochItem(db, address, math.MaxUint64)
	item := GetEpochItem(db, address, tail.PreEpoch)
	return item
}

func GetLastEpoch(db sdk.StateDB, address common.Address) uint64 {
	tail := GetEpochItem(db, address, math.MaxUint64)
	return tail.PreEpoch
}

func GetEpochQueueFromHead(db sdk.StateDB, address common.Address, size uint64) types.EpochQueue {
	queue := types.NewEpochQueue(size)

	var count uint64 = 0

	item := GetEpochItem(db, address, 1)

	for item.NextEpoch != 0 && count < size { // not as tail<nextEpoch == 0>  and count less size
		queue[count] = item
		item = GetEpochItem(db, address, item.NextEpoch)
		count++
	}
	return queue[:count]
}

func GetEpochQueueFromTail(db sdk.StateDB, address common.Address, size uint64) types.EpochQueue {
	queue := types.NewEpochQueue(size)

	var count uint64 = 0

	tail := GetEpochItem(db, address, math.MaxUint64)
	item := GetEpochItem(db, address, tail.PreEpoch)

	for item.PreEpoch != math.MaxUint64 && count < size { // not as head<preEpoch == MaxUint64>  and count less size
		queue[count] = item
		item = GetEpochItem(db, address, item.PreEpoch)
		count++
	}
	return queue[:count]
}

func GetEpochQueueAndIndexFromHead(db sdk.StateDB, address common.Address, size uint64) ([]uint64, types.EpochQueue) {
	queue := types.NewEpochQueue(size)
	epochs := make([]uint64, size)
	var count uint64 = 0

	epoch := uint64(1)
	item := GetEpochItem(db, address, epoch)

	for item.NextEpoch != 0 && count < size { // not as tail<nextEpoch == 0>  and count less size
		queue[count] = item
		epochs[count] = epoch
		item = GetEpochItem(db, address, item.NextEpoch)
		epoch = item.NextEpoch
		count++
	}
	return epochs[:count], queue[:count]
}

func GetEpochQueueAndIndexFromTail(db sdk.StateDB, address common.Address, size uint64) ([]uint64, types.EpochQueue) {
	queue := types.NewEpochQueue(size)
	epochs := make([]uint64, size)
	var count uint64 = 0

	tail := GetEpochItem(db, address, math.MaxUint64)
	epoch := tail.PreEpoch
	item := GetEpochItem(db, address, epoch)

	for item.PreEpoch != math.MaxUint64 && count < size { // not as head<preEpoch == MaxUint64>  and count less size
		queue[count] = item
		epochs[count] = epoch
		item = GetEpochItem(db, address, item.PreEpoch)
		epoch = item.NextEpoch
		count++
	}
	return epochs[:count], queue[:count]
}

func SetEpochItem(db sdk.StateDB, address common.Address, epoch uint64, item *types.EpochItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address, EncodeEpochItemKey(epoch), value)
	return nil
}

func GetEpochItem(db sdk.StateDB, address common.Address, epoch uint64) *types.EpochItem {
	value := db.GetState(address, EncodeEpochItemKey(epoch))
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

func AppendRoundItem(db sdk.StateDB, address common.Address, round uint64, startBlock, endBlock uint64) error {

	indexRound := uint64(math.MaxUint64)
	indexItem := GetRoundItem(db, address, indexRound)

	// range from tail to head
	for indexItem.PreRound != math.MaxUint64 { // not as head
		if indexRound == round {
			// pre -> index(round) -> next
			// index == round

			return ErrInvalidValue
		}

		if indexRound < round {
			// pre -> index -> round -> next... -> tail(max)
			// pre < index < round < next ... < tail(max)

			next := GetRoundItem(db, address, indexItem.NextRound)
			roundItem := types.NewRoundItem(indexRound, indexItem.NextRound, startBlock, endBlock)

			indexItem.UpdateNextRound(round)
			next.UpdatePreRound(round)

			if err := SetRoundItem(db, address, indexRound, indexItem); nil != err {
				return err
			}
			if err := SetRoundItem(db, address, round, roundItem); nil != err {
				return err
			}
			if err := SetRoundItem(db, address, roundItem.NextRound, next); nil != err {
				return err
			}

			break

		} else {

			if indexItem.PreRound == uint64(0) {
				// if  head(min) -> index -> tail(max)
				// and round < index
				//
				// then: head(min) -> round -> index<last one> -> ... -> tail(max)
				// head < round < index < ... < tail

				pre := GetRoundItem(db, address, indexItem.PreRound) // head
				epochItem := types.NewRoundItem(indexItem.PreRound, round, startBlock, endBlock)

				pre.UpdateNextRound(round)
				indexItem.UpdatePreRound(round)

				if err := SetRoundItem(db, address, epochItem.PreRound, pre); nil != err {
					return err
				}
				if err := SetRoundItem(db, address, round, epochItem); nil != err {
					return err
				}
				if err := SetRoundItem(db, address, indexRound, indexItem); nil != err {
					return err
				}

				break

			}

			// if  head -> ... -> pre -> index -> ... -> max
			// (head < ... <  pre < index < ... < tail(max) )
			// and round < index
			// maybe round < pre
			//
			// then: continue pre become new index
			//

		}

		indexRound = indexItem.PreRound
		indexItem = GetRoundItem(db, address, indexRound)
	}
	return nil
}

func GetLastRoundItem(db sdk.StateDB, address common.Address) *types.RoundItem {
	tail := GetRoundItem(db, address, math.MaxUint64)
	item := GetRoundItem(db, address, tail.PreRound)
	return item
}

func GetLastRound(db sdk.StateDB, address common.Address) uint64 {
	tail := GetRoundItem(db, address, math.MaxUint64)
	return tail.PreRound
}

func GetRoundQueueFromHead(db sdk.StateDB, address common.Address, size uint64) types.RoundQueue {
	queue := types.NewRoundQueue(size)

	var count uint64 = 0

	item := GetRoundItem(db, address, 1)

	for item.NextRound != 0 && count < size { // not as tail<nextRound == 0>  and count less size
		queue[count] = item
		item = GetRoundItem(db, address, item.NextRound)
		count++
	}
	return queue[:count]
}

func GetRoundQueueFromTail(db sdk.StateDB, address common.Address, size uint64) types.RoundQueue {
	queue := types.NewRoundQueue(size)

	var count uint64 = 0

	tail := GetRoundItem(db, address, math.MaxUint64)
	item := GetRoundItem(db, address, tail.PreRound)

	for item.PreRound != math.MaxUint64 && count < size { // not as head<preRound == MaxUint64>  and count less size
		queue[count] = item
		item = GetRoundItem(db, address, item.PreRound)
		count++
	}
	return queue[:count]
}

func GetRoundQueueAndIndexFromHead(db sdk.StateDB, address common.Address, size uint64) ([]uint64, types.RoundQueue) {
	queue := types.NewRoundQueue(size)
	rounds := make([]uint64, size)
	var count uint64 = 0

	round := uint64(1)
	item := GetRoundItem(db, address, round)

	for item.NextRound != 0 && count < size { // not as tail<nextRound == 0>  and count less size
		queue[count] = item
		rounds[count] = round
		item = GetRoundItem(db, address, item.NextRound)
		round = item.NextRound
		count++
	}
	return rounds[:count], queue[:count]
}

func GetRoundQueueAndIndexFromTail(db sdk.StateDB, address common.Address, size uint64) ([]uint64, types.RoundQueue) {
	queue := types.NewRoundQueue(size)
	rounds := make([]uint64, size)
	var count uint64 = 0

	tail := GetRoundItem(db, address, math.MaxUint64)
	round := tail.PreRound
	item := GetRoundItem(db, address, round)

	for item.PreRound != math.MaxUint64 && count < size { // not as head<preRound == MaxUint64>  and count less size
		queue[count] = item
		rounds[count] = round
		item = GetRoundItem(db, address, item.PreRound)
		round = item.NextRound
		count++
	}
	return rounds[:count], queue[:count]
}

func SetRoundItem(db sdk.StateDB, address common.Address, round uint64, item *types.RoundItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(address, EncodeRoundItemKey(round), value)
	return nil
}

func GetRoundItem(db sdk.StateDB, address common.Address, round uint64) *types.RoundItem {
	value := db.GetState(address, EncodeRoundItemKey(round))
	if len(value) == 0 {
		return nil
	}

	var item types.RoundItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}
