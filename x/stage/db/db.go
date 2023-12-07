package db

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
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
	currentEpochKey    = []byte("currentEpoch") // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey    = []byte("currentRound") // "currentRound" => currentRound (It is a number)
	epochItemKeyPrefix = []byte("epochItem")    // "epochItem":epochId => {preEpoch, nextEpoch, startBlock, endBlock, roundCount}
	roundItemKeyPrefix = []byte("roundItem")    // "roundItem":roundId => {preRound, nextRound, startBlock, endBlock}
)

func EncodeEpochItemKey(epoch uint64) []byte {
	return append(epochItemKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func EncodeRoundItemKey(round uint64) []byte {
	return append(roundItemKeyPrefix, common.Uint64ToBytes(round)...)
}

// ---------

func IncrementCurrentEpoch(db sdk.StateDB, addr common.Address) {
	epoch := GetCurrentEpoch(db, addr)
	db.SetState(addr, currentEpochKey, common.Uint64ToBytes(epoch+1))
}

func GetCurrentEpoch(db sdk.StateDBReader, addr common.Address) uint64 {
	value := db.GetState(addr, currentEpochKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func InrementCurrentRound(db sdk.StateDB, addr common.Address) {
	round := GetCurrentRound(db, addr)
	db.SetState(addr, currentRoundKey, common.Uint64ToBytes(round+1))
}

func GetCurrentRound(db sdk.StateDBReader, addr common.Address) uint64 {
	value := db.GetState(addr, currentRoundKey)
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

// ------

func AppendEpochItem(db sdk.StateDB, addr common.Address, epoch uint64, startBlock, endBlock, roundCount uint64) error {
	epochItem := GetEpochItem(db, addr, epoch)
	if epochItem.IsNotEmpty() {
		return ErrExist
	}
	epochItem = types.NewEpochItem(startBlock, endBlock, roundCount)
	return SetEpochItem(db, addr, epoch, epochItem)
}

func GetEpochQueueSince(db sdk.StateDBReader, addr common.Address, epoch, size uint64) types.EpochQueue {
	queue := types.NewEpochQueue(size)

	var count uint64 = 0

	currentEpoch := GetCurrentEpoch(db, addr)
	index := epoch

	for index != currentEpoch+1 && count < size {
		queue[count] = GetEpochItem(db, addr, index)
		index++
		count++
	}
	return queue[:count]
}

func GetEpochQueueFromTail(db sdk.StateDBReader, addr common.Address, epoch, size uint64) types.EpochQueue {
	queue := types.NewEpochQueue(size)

	var count uint64 = 0

	index := epoch
	for index != 0 && count < size {
		queue[count] = GetEpochItem(db, addr, index)
		index--
		count++
	}
	return queue[:count]
}

func GetEpochQueueAndIndexSince(db sdk.StateDBReader, addr common.Address, epoch, size uint64) ([]uint64, types.EpochQueue) {
	queue := types.NewEpochQueue(size)
	epochs := make([]uint64, size)

	var count uint64 = 0

	currentEpoch := GetCurrentEpoch(db, addr)
	index := epoch

	for index != currentEpoch+1 && count < size {
		queue[count] = GetEpochItem(db, addr, index)
		epochs[count] = index
		index++
		count++
	}
	return epochs[:count], queue[:count]
}

func GetEpochQueueAndIndexFromTail(db sdk.StateDBReader, addr common.Address, epoch, size uint64) ([]uint64, types.EpochQueue) {
	queue := types.NewEpochQueue(size)
	epochs := make([]uint64, size)

	var count uint64 = 0

	index := epoch
	for index != 0 && count < size {
		queue[count] = GetEpochItem(db, addr, index)
		epochs[count] = index
		index--
		count++
	}
	return epochs[:count], queue[:count]
}

func SetEpochItem(db sdk.StateDB, addr common.Address, epoch uint64, item *types.EpochItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, EncodeEpochItemKey(epoch), value)
	return nil
}

func GetEpochItem(db sdk.StateDBReader, addr common.Address, epoch uint64) *types.EpochItem {
	value := db.GetState(addr, EncodeEpochItemKey(epoch))
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

func AppendRoundItem(db sdk.StateDB, addr common.Address, round uint64, startBlock, endBlock uint64) error {
	roundItem := GetRoundItem(db, addr, round)
	if roundItem.IsNotEmpty() {
		return ErrExist
	}
	roundItem = types.NewRoundItem(startBlock, endBlock)
	return SetRoundItem(db, addr, round, roundItem)
}

func GetRoundQueueSince(db sdk.StateDBReader, addr common.Address, round, size uint64) types.RoundQueue {
	queue := types.NewRoundQueue(size)

	var count uint64 = 0

	currentEpoch := GetCurrentRound(db, addr)
	index := round

	for index != currentEpoch+1 && count < size {
		queue[count] = GetRoundItem(db, addr, index)
		index++
		count++
	}
	return queue[:count]
}

func GetRoundQueueFromTail(db sdk.StateDBReader, addr common.Address, round, size uint64) types.RoundQueue {
	queue := types.NewRoundQueue(size)

	var count uint64 = 0

	index := round
	for index != 0 && count < size {
		queue[count] = GetRoundItem(db, addr, index)
		index--
		count++
	}
	return queue[:count]
}

func GetRoundQueueAndIndexSince(db sdk.StateDBReader, addr common.Address, round, size uint64) ([]uint64, types.RoundQueue) {
	queue := types.NewRoundQueue(size)
	rounds := make([]uint64, size)

	var count uint64 = 0

	currentRound := GetCurrentRound(db, addr)
	index := round

	for index != currentRound+1 && count < size {
		queue[count] = GetRoundItem(db, addr, index)
		rounds[count] = index
		index++
		count++
	}
	return rounds[:count], queue[:count]
}

func GetRoundQueueAndIndexFromTail(db sdk.StateDBReader, addr common.Address, round, size uint64) ([]uint64, types.RoundQueue) {
	queue := types.NewRoundQueue(size)
	rounds := make([]uint64, size)

	var count uint64 = 0

	index := round
	for index != 0 && count < size {
		queue[count] = GetRoundItem(db, addr, index)
		rounds[count] = index
		index--
		count++
	}
	return rounds[:count], queue[:count]
}

func SetRoundItem(db sdk.StateDB, addr common.Address, round uint64, item *types.RoundItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, EncodeRoundItemKey(round), value)
	return nil
}

func GetRoundItem(db sdk.StateDBReader, addr common.Address, round uint64) *types.RoundItem {
	value := db.GetState(addr, EncodeRoundItemKey(round))
	if len(value) == 0 {
		return nil
	}

	var item types.RoundItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

// ---------

func IsElectionBlockOnCurrentRound(db sdk.StateDBReader, addr common.Address, blockNumber, distance uint64) bool {

	currentRound := GetCurrentRound(db, addr)
	currentRoundItem := GetRoundItem(db, addr, currentRound)
	if currentRoundItem.IsEmpty() {
		return false
	}

	tmp := blockNumber + distance
	if tmp == currentRoundItem.EndBlock {
		return true
	}
	return false
}

func IsNotElectionBlockOnCurrentRound(db sdk.StateDBReader, addr common.Address, blockNumber, distance uint64) bool {
	return !IsElectionBlockOnCurrentRound(db, addr, blockNumber, distance)
}

func IsBeginOfRound(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {

	currentRound := GetCurrentRound(db, addr)
	queue := GetRoundQueueFromTail(db, addr, currentRound, size)
	for _, item := range queue {
		if item.StartBlock == blockNumber {
			return true
		}
	}
	return false
}

func IsNotBeginOfRound(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {
	return !IsBeginOfRound(db, addr, blockNumber, size)
}

func IsBeginOfCurrentRound(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	currentRound := GetCurrentRound(db, addr)
	currentRoundItem := GetRoundItem(db, addr, currentRound)
	if currentRoundItem.StartBlock == blockNumber {
		return true
	}
	return false
}

func IsNotBeginOfCurrentRound(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	return !IsBeginOfCurrentRound(db, addr, blockNumber)
}

func IsBeginOfNextRound(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	currentRound := GetCurrentRound(db, addr)
	currentRoundItem := GetRoundItem(db, addr, currentRound)
	if currentRoundItem.EndBlock+1 == blockNumber {
		return true
	}
	return false
}

func IsNotBeginOfNextRound(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	return !IsBeginOfNextRound(db, addr, blockNumber)
}

func IsEndOfRound(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {

	currentRound := GetCurrentRound(db, addr)
	queue := GetRoundQueueFromTail(db, addr, currentRound, size)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
}

func IsNotEndOfRound(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {
	return !IsEndOfRound(db, addr, blockNumber, size)
}

func IsEndOfCurrentRound(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	currentRound := GetCurrentRound(db, addr)
	currentRoundItem := GetRoundItem(db, addr, currentRound)
	if currentRoundItem.IsEmpty() {
		return false
	}
	if currentRoundItem.EndBlock == blockNumber {
		return true
	}
	return false
}

func IsNotEndOfCurrentRound(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	return !IsEndOfCurrentRound(db, addr, blockNumber)
}

// ------

func IsElectionBlockOnCurrentEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {

	currentEpoch := GetCurrentEpoch(db, addr)
	currentEpochItem := GetEpochItem(db, addr, currentEpoch)
	if currentEpochItem.IsEmpty() {
		return false
	}

	if currentEpochItem.EndBlock == blockNumber {
		return true
	}
	return false
}

func IsNotElectionBlockOnCurrentEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	return !IsElectionBlockOnCurrentEpoch(db, addr, blockNumber)
}

func IsBeginOfEpoch(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {

	currentEpoch := GetCurrentEpoch(db, addr)
	queue := GetEpochQueueFromTail(db, addr, currentEpoch, size)
	for _, item := range queue {
		if item.StartBlock == blockNumber {
			return true
		}
	}
	return false
}

func IsNotBeginOfEpoch(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {
	return !IsBeginOfEpoch(db, addr, blockNumber, size)
}

func IsBeginOfCurrentEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	currentEpoch := GetCurrentEpoch(db, addr)
	currentEpochItem := GetEpochItem(db, addr, currentEpoch)
	if currentEpochItem.StartBlock == blockNumber {
		return true
	}
	return false
}

func IsNotBeginOfCurrentEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	return !IsBeginOfCurrentEpoch(db, addr, blockNumber)
}

func IsBeginOfNextEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	currentEpoch := GetCurrentEpoch(db, addr)
	currentEpochItem := GetEpochItem(db, addr, currentEpoch)
	if currentEpochItem.EndBlock+1 == blockNumber {
		return true
	}
	return false
}

func IsNotBeginOfNextEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	return !IsBeginOfNextEpoch(db, addr, blockNumber)
}

func IsEndOfEpoch(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {
	currentEpoch := GetCurrentEpoch(db, addr)
	queue := GetEpochQueueFromTail(db, addr, currentEpoch, size)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
}

func IsNotEndOfEpoch(db sdk.StateDBReader, addr common.Address, blockNumber, size uint64) bool {
	return !IsEndOfEpoch(db, addr, blockNumber, size)
}

func IsEndOfCurrentEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {

	currentEpoch := GetCurrentEpoch(db, addr)
	currentEpochItem := GetEpochItem(db, addr, currentEpoch)
	if currentEpochItem.IsEmpty() {
		return false
	}
	if currentEpochItem.EndBlock == blockNumber {
		return true
	}
	return false
}

func IsNotEndOfCurrentEpoch(db sdk.StateDBReader, addr common.Address, blockNumber uint64) bool {
	return !IsEndOfCurrentEpoch(db, addr, blockNumber)
}

func BuildNextRound(db sdk.StateDB, addr common.Address, roundSize uint64) error {

	currentRound := GetCurrentRound(db, addr)
	currentRoundItem := GetRoundItem(db, addr, currentRound)

	startBlock := currentRoundItem.EndBlock + 1
	endBlock := currentRoundItem.EndBlock + roundSize
	return AppendRoundItem(db, addr, currentRound+1, startBlock, endBlock)
}

func BuildNextEpoch(db sdk.StateDB, addr common.Address, epochSize, roundSize uint64) error {

	currentEpoch := GetCurrentEpoch(db, addr)
	currentEpochItem := GetEpochItem(db, addr, currentEpoch)

	startBlock := currentEpochItem.EndBlock + 1
	endBlock := currentEpochItem.EndBlock + epochSize
	roundCount := epochSize / roundSize
	return AppendEpochItem(db, addr, currentEpoch+1, startBlock, endBlock, roundCount)
}
