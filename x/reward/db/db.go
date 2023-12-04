package db

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
	"github.com/PlatONnetwork/PlatON-Go/rlp"

	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
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
	pendingValidatorRewardKeyPrefix = []byte("pendingValidatorReward") // "pendingValidatorReward":validatorAddr => pendingReward
	pendingDelegaterRewardKeyPrefix = []byte("pendingDelegaterReward") // "pendingDelegaterReward":delegaterAddr => pendingReward

	validatorRewardOwnerKeyPrefix = []byte("validatorRewardOwner") // "validatorRewardOwner":validatorAddr => ownerAddr

	delegaterRewardPendingIndexKeyPrefix       = []byte("delegaterRewardPendingIndex")       // "delegaterRewardPendingIndex":delegaterAddr:validatorAddr => epochId
	epochRewardPerDelegationShareItemKeyPrefix = []byte("epochRewardPerDelegationShareItem") // "epochRewardPerDelegationShareItem":validatorAddr:epochId => epochRewardPerDelegationShareItem{totalReward, perShareReward}
)

func encodePendingValidatorRewardKey(validatorAddr basecommon.Address) []byte {
	return append(pendingValidatorRewardKeyPrefix, validatorAddr.Bytes()...)
}

func encodePendingDelegaterRewardKey(delegaterAddr basecommon.Address) []byte {
	return append(pendingDelegaterRewardKeyPrefix, delegaterAddr.Bytes()...)
}

func encodeValidatorRewardOwnerKey(validatorAddr basecommon.Address) []byte {
	return append(validatorRewardOwnerKeyPrefix, validatorAddr.Bytes()...)
}

func encodeDelegaterRewardPendingIndexKey(delegaterAddr, validatorAddr basecommon.Address) []byte {

	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()

	keyPrefixSize := len(delegaterRewardPendingIndexKeyPrefix)
	appendDelegaterAddrSize := keyPrefixSize + len(delegaterAddrBytes)
	size := appendDelegaterAddrSize + len(validatorAddrBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegaterRewardPendingIndexKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterAddrSize], delegaterAddrBytes)
	copy(key[appendDelegaterAddrSize:], validatorAddrBytes)

	return key
}

func encodeEpochRewardPerDelegationShareItemKey(validatorAddr basecommon.Address, epoch uint64) []byte {

	validatorAddrBytes := validatorAddr.Bytes()
	epochBytes := basecommon.Uint64ToBytes(epoch)

	keyPrefixSize := len(epochRewardPerDelegationShareItemKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(epochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], epochRewardPerDelegationShareItemKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], epochBytes)

	return key
}

func IncrementPendingValidatorReward(db sdk.StateDB, addr basecommon.Address, validatorAddr basecommon.Address, increment uint64) {
	number := GetPendingValidatorReward(db, addr, validatorAddr)
	number += increment
	db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), basecommon.Uint64ToBytes(number))
}

func DecrementPendingValidatorReward(db sdk.StateDB, addr basecommon.Address, validatorAddr basecommon.Address, decrement uint64) {
	number := GetPendingValidatorReward(db, addr, validatorAddr)

	if number <= decrement {
		db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), []byte{})
	} else {
		number -= decrement
		db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), basecommon.Uint64ToBytes(number))
	}
}

func GetPendingValidatorReward(db sdk.StateDBReader, addr basecommon.Address, validatorAddr basecommon.Address) uint64 {
	value := db.GetState(addr, encodePendingValidatorRewardKey(validatorAddr))
	var number uint64
	if len(value) != 0 {
		number = basecommon.BytesToUint64(value)
	}
	return number
}

func IncrementPendingDelegaterReward(db sdk.StateDB, addr basecommon.Address, delegaterAddr basecommon.Address, increment uint64) {
	number := GetPendingDelegaterReward(db, addr, delegaterAddr)
	number += increment
	db.SetState(addr, encodePendingDelegaterRewardKey(delegaterAddr), basecommon.Uint64ToBytes(number))
}

func DecrementPendingDelegaterReward(db sdk.StateDB, addr basecommon.Address, delegaterAddr basecommon.Address, decrement uint64) {
	number := GetPendingDelegaterReward(db, addr, delegaterAddr)

	if number <= decrement {
		db.SetState(addr, encodePendingDelegaterRewardKey(delegaterAddr), []byte{})
	} else {
		number -= decrement
		db.SetState(addr, encodePendingDelegaterRewardKey(delegaterAddr), basecommon.Uint64ToBytes(number))
	}
}

func GetPendingDelegaterReward(db sdk.StateDBReader, addr basecommon.Address, delegaterAddr basecommon.Address) uint64 {
	value := db.GetState(addr, encodePendingDelegaterRewardKey(delegaterAddr))
	var number uint64
	if len(value) != 0 {
		number = basecommon.BytesToUint64(value)
	}
	return number
}

func SetValidatorRewardOwner(db sdk.StateDB, addr basecommon.Address, validatorAddr, ownerAddr basecommon.Address) {
	db.SetState(addr, encodeValidatorRewardOwnerKey(validatorAddr), ownerAddr.Bytes())
}

func GetValidatorRewardOwner(db sdk.StateDB, addr basecommon.Address, validatorAddr basecommon.Address) basecommon.Address {
	ownerAddrBytes := db.GetState(addr, encodeValidatorRewardOwnerKey(validatorAddr))
	if len(ownerAddrBytes) == 0 {
		return basecommon.ZeroAddr
	}
	return basecommon.BytesToAddress(ownerAddrBytes)
}

func GetDelegaterRewardPendingIndex(db sdk.StateDB, addr, delegaterAddr, validatorAddr basecommon.Address) uint64 {

	value := db.GetState(addr, encodeDelegaterRewardPendingIndexKey(delegaterAddr, validatorAddr))

	if len(value) == 0 {
		return 0
	}
	return basecommon.BytesToUint64(value)
}

func SetDelegaterRewardPendingIndex(db sdk.StateDB, addr, delegaterAddr, validatorAddr basecommon.Address, epoch uint64) {
	db.SetState(addr, encodeDelegaterRewardPendingIndexKey(delegaterAddr, validatorAddr), basecommon.Uint64ToBytes(epoch))
}

func GetEpochRewardPerDelegationShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, epoch uint64) *types.EpochRewardPerDelegationShareItem {
	value := db.GetState(addr, encodeEpochRewardPerDelegationShareItemKey(validatorAddr, epoch))
	if len(value) == 0 {
		return nil
	}

	var item types.EpochRewardPerDelegationShareItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

func SetEpochRewardPerDelegationShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, epoch uint64, item *types.EpochRewardPerDelegationShareItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeEpochRewardPerDelegationShareItemKey(validatorAddr, epoch), value)
	return nil
}
