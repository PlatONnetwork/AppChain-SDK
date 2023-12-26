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

// Will be governed in the future
var (
	roundValidatorElectionDistanceKey = []byte("roundValidatorElectionDistance")
	roundSizeKey                      = []byte("roundSize")
	epochSizeKey                      = []byte("epochSize")
)

var (
	currentEpochKey    = []byte("currentEpoch") // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey    = []byte("currentRound") // "currentRound" => currentRound (It is a number)
	epochItemKeyPrefix = []byte("epochItem")    // "epochItem":epochId => {preEpoch, nextEpoch, startBlock, endBlock, roundCount}
	roundItemKeyPrefix = []byte("roundItem")    // "roundItem":roundId => {preRound, nextRound, startBlock, endBlock}
)

// ------
func EncodeRoundValidatorElectionDistanceKey() []byte {
	return roundValidatorElectionDistanceKey
}
func EncodeRoundSizeKey() []byte {
	return roundSizeKey
}
func EncodeEpochSizeKey() []byte { return epochSizeKey }

// ------
func EncodeEpochItemKey(epoch uint64) []byte {
	return append(epochItemKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func EncodeRoundItemKey(round uint64) []byte {
	return append(roundItemKeyPrefix, common.Uint64ToBytes(round)...)
}

// ------------------------------------------------------ db methods ------------------------------------------------------

func SetRoundValidatorElectionDistance(db sdk.StateDB, addr common.Address, value uint64) {
	db.SetState(addr, EncodeRoundValidatorElectionDistanceKey(), common.Uint64ToBytes(value))
}
func GetRoundValidatorElectionDistance(db sdk.StateDBReader, addr common.Address) uint64 {
	value := db.GetState(addr, EncodeRoundValidatorElectionDistanceKey())
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func SetRoundSize(db sdk.StateDB, addr common.Address, value uint64) {
	db.SetState(addr, EncodeRoundSizeKey(), common.Uint64ToBytes(value))
}
func GetRoundSize(db sdk.StateDBReader, addr common.Address) uint64 {
	value := db.GetState(addr, EncodeRoundSizeKey())
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

func SetEpochSize(db sdk.StateDB, addr common.Address, value uint64) {
	db.SetState(addr, EncodeEpochSizeKey(), common.Uint64ToBytes(value))
}
func GetEpochSize(db sdk.StateDBReader, addr common.Address) uint64 {
	value := db.GetState(addr, EncodeEpochSizeKey())
	if len(value) == 0 {
		return 0
	}
	return common.BytesToUint64(value)
}

// ------
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

	var count uint64 = 0

	currentEpoch := GetCurrentEpoch(db, addr)
	index := epoch

	queue := types.NewEpochQueue(0)
	for index != currentEpoch+1 && count < size {
		item := GetEpochItem(db, addr, index)
		if item.IsEmpty() {
			break
		}
		queue = append(queue, item)
		index++
		count++
	}
	return queue
}

func GetEpochQueueFromTail(db sdk.StateDBReader, addr common.Address, epoch, size uint64) types.EpochQueue {

	var count uint64 = 0

	index := epoch

	queue := types.NewEpochQueue(0)
	for index != 0 && count < size {
		item := GetEpochItem(db, addr, index)
		if item.IsEmpty() {
			break
		}
		queue = append(queue, item)
		index--
		count++
	}
	return queue
}

func GetEpochQueueAndIndexSince(db sdk.StateDBReader, addr common.Address, epoch, size uint64) ([]uint64, types.EpochQueue) {

	var count uint64 = 0

	currentEpoch := GetCurrentEpoch(db, addr)
	index := epoch

	epochs := make([]uint64, 0)
	queue := types.NewEpochQueue(0)
	for index != currentEpoch+1 && count < size {
		item := GetEpochItem(db, addr, index)
		if item.IsEmpty() {
			break
		}

		queue = append(queue, item)
		epochs = append(epochs, index)
		index++
		count++
	}
	return epochs, queue
}

func GetEpochQueueAndIndexFromTail(db sdk.StateDBReader, addr common.Address, epoch, size uint64) ([]uint64, types.EpochQueue) {

	var count uint64 = 0

	index := epoch

	epochs := make([]uint64, 0)
	queue := types.NewEpochQueue(0)
	for index != 0 && count < size {
		item := GetEpochItem(db, addr, index)
		if item.IsEmpty() {
			break
		}
		queue = append(queue, item)
		epochs = append(epochs, index)
		index--
		count++
	}
	return epochs, queue
}

func GetEpochItemAndIndexByBlockNumber(db sdk.StateDBReader, addr common.Address, blockNumber uint64) (uint64, *types.EpochItem) {

	// #### NOTE ####
	// Starting from the next epoch of data collection,
	// it is mainly for many scenarios to go back and obtain
	// the next epoch of data in advance,
	// and the next epoch of data will be generated in advance
	currentEpoch := GetCurrentEpoch(db, addr)
	index := currentEpoch + 1
	item := GetEpochItem(db, addr, index)
	for index != 0 {
		if item.IsEmpty() {
			index--
			item = GetEpochItem(db, addr, index)
			continue
		}
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {
			break
		}
		index--
		item = GetEpochItem(db, addr, index)
	}
	return index, item
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

	var count uint64 = 0

	currentEpoch := GetCurrentRound(db, addr)
	index := round

	queue := types.NewRoundQueue(0)
	for index != currentEpoch+1 && count < size {
		item := GetRoundItem(db, addr, index)
		if item.IsEmpty() {
			break
		}
		queue = append(queue, item)
		index++
		count++
	}
	return queue
}

func GetRoundQueueFromTail(db sdk.StateDBReader, addr common.Address, round, size uint64) types.RoundQueue {

	var count uint64 = 0

	index := round

	queue := types.NewRoundQueue(0)
	for index != 0 && count < size {
		item := GetRoundItem(db, addr, index)
		if item.IsEmpty() {
			break
		}
		queue = append(queue, item)
		index--
		count++
	}
	return queue
}

func GetRoundQueueAndIndexSince(db sdk.StateDBReader, addr common.Address, round, size uint64) ([]uint64, types.RoundQueue) {

	var count uint64 = 0

	currentRound := GetCurrentRound(db, addr)
	index := round

	rounds := make([]uint64, 0)
	queue := types.NewRoundQueue(0)
	for index != currentRound+1 && count < size {
		item := GetRoundItem(db, addr, index)
		if item.IsEmpty() {
			break
		}

		queue = append(queue, item)
		rounds = append(rounds, index)
		index++
		count++
	}
	return rounds, queue
}

func GetRoundQueueAndIndexFromTail(db sdk.StateDBReader, addr common.Address, round, size uint64) ([]uint64, types.RoundQueue) {

	var count uint64 = 0

	index := round

	rounds := make([]uint64, 0)
	queue := types.NewRoundQueue(0)
	for index != 0 && count < size {
		item := GetRoundItem(db, addr, index)
		if item.IsEmpty() {
			break
		}
		queue = append(queue, item)
		rounds = append(rounds, index)
		index--
		count++
	}
	return rounds, queue
}

func GetRoundItemAndIndexByBlockNumber(db sdk.StateDBReader, addr common.Address, blockNumber uint64) (uint64, *types.RoundItem) {

	// #### NOTE ####
	// Starting from the next round of data collection,
	// it is mainly for many scenarios to go back and obtain
	// the next round of data in advance,
	// and the next round of data will be generated in advance
	currentRound := GetCurrentRound(db, addr)
	index := currentRound + 1
	item := GetRoundItem(db, addr, index)
	for index != 0 {
		if item.IsEmpty() {
			index--
			item = GetRoundItem(db, addr, index)
			continue
		}
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {
			break
		}
		index--
		item = GetRoundItem(db, addr, index)
	}
	return index, item
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
	// #### NOTE ####
	// Starting from the next round of data collection,
	// it is mainly for many scenarios to go back and obtain
	// the next round of data in advance,
	// and the next round of data will be generated in advance
	currentRound := GetCurrentRound(db, addr)
	index := currentRound + 1
	queue := GetRoundQueueFromTail(db, addr, index, size)
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
	// #### NOTE ####
	// Starting from the next round of data collection,
	// it is mainly for many scenarios to go back and obtain
	// the next round of data in advance,
	// and the next round of data will be generated in advance
	currentRound := GetCurrentRound(db, addr)
	index := currentRound + 1
	queue := GetRoundQueueFromTail(db, addr, index, size)
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

	// #### NOTE ####
	// Starting from the next epoch of data collection,
	// it is mainly for many scenarios to go back and obtain
	// the next epoch of data in advance,
	// and the next epoch of data will be generated in advance
	currentEpoch := GetCurrentEpoch(db, addr)
	index := currentEpoch + 1
	queue := GetEpochQueueFromTail(db, addr, index, size)
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
	// #### NOTE ####
	// Starting from the next epoch of data collection,
	// it is mainly for many scenarios to go back and obtain
	// the next epoch of data in advance,
	// and the next epoch of data will be generated in advance
	currentEpoch := GetCurrentEpoch(db, addr)
	index := currentEpoch + 1
	queue := GetEpochQueueFromTail(db, addr, index, size)
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
