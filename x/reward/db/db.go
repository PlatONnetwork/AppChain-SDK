package db

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/big"

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
	pendingDelegaterRewardKeyPrefix = []byte("pendingDelegaterReward") // "pendingDelegaterReward":delegaterAddr:validatorAddr => pendingReward

	validatorRewardOwnerKeyPrefix = []byte("validatorRewardOwner") // "validatorRewardOwner":validatorAddr => ownerAddr

	//delegaterRewardPendingIndexKeyPrefix       = []byte("delegaterRewardPendingIndex")       // "delegaterRewardPendingIndex":delegaterAddr:validatorAddr:stakeEpoch => needRewardEpoch
	epochDelegationRewardPerShareItemKeyPrefix = []byte("epochDelegationRewardPerShareItem") // "epochDelegationRewardPerShareItem":validatorAddr:stakeEpoch:epochId => epochRewardPerDelegationShareItem{ previousEpoch, nextEpoch, totalReward, perShareReward}
)

func encodePendingValidatorRewardKey(validatorAddr basecommon.Address) []byte {
	return append(pendingValidatorRewardKeyPrefix, validatorAddr.Bytes()...)
}

func encodePendingDelegaterRewardKey(delegaterAddr, validatorAddr basecommon.Address) []byte {

	delegaterAddrBytes := delegaterAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()

	keyPrefixSize := len(pendingDelegaterRewardKeyPrefix)
	appendDelegaterAddrSize := keyPrefixSize + len(delegaterAddrBytes)
	size := appendDelegaterAddrSize + len(validatorAddrBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], pendingDelegaterRewardKeyPrefix)
	copy(key[keyPrefixSize:appendDelegaterAddrSize], delegaterAddrBytes)
	copy(key[appendDelegaterAddrSize:], validatorAddrBytes)

	return key
}

func encodeValidatorRewardOwnerKey(validatorAddr basecommon.Address) []byte {
	return append(validatorRewardOwnerKeyPrefix, validatorAddr.Bytes()...)
}

//func encodeDelegaterRewardPendingIndexKey(delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) []byte {
//	delegaterAddrBytes := delegaterAddr.Bytes()
//	validatorAddrBytes := validatorAddr.Bytes()
//	stakeEpochBytes := common.Uint64ToBytes(stakeEpoch)
//
//	keyPrefixSize := len(delegaterRewardPendingIndexKeyPrefix)
//	appendDelegaterSize := keyPrefixSize + len(delegaterAddrBytes)
//	appendVlidatorAddrSize := appendDelegaterSize + len(validatorAddrBytes)
//	size := appendVlidatorAddrSize + len(stakeEpochBytes)
//
//	key := make([]byte, size)
//
//	copy(key[:keyPrefixSize], delegaterRewardPendingIndexKeyPrefix)
//	copy(key[keyPrefixSize:appendDelegaterSize], delegaterAddrBytes)
//	copy(key[appendDelegaterSize:appendVlidatorAddrSize], validatorAddrBytes)
//	copy(key[appendVlidatorAddrSize:], stakeEpochBytes)
//
//	return key
//}

func encodeEpochDelegationRewardPerShareItemKey(validatorAddr basecommon.Address, epoch uint64) []byte {

	validatorAddrBytes := validatorAddr.Bytes()
	epochBytes := basecommon.Uint64ToBytes(epoch)

	keyPrefixSize := len(epochDelegationRewardPerShareItemKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	size := appendValidatorAddrSize + len(epochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], epochDelegationRewardPerShareItemKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:], epochBytes)

	return key
}

func IncrementPendingValidatorReward(db sdk.StateDB, addr basecommon.Address, validatorAddr basecommon.Address, increment *big.Int) {
	number := GetPendingValidatorReward(db, addr, validatorAddr)
	number = new(big.Int).Add(number, increment)
	db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), number.Bytes())
}

func DecrementPendingValidatorReward(db sdk.StateDB, addr basecommon.Address, validatorAddr basecommon.Address, decrement *big.Int) {
	number := GetPendingValidatorReward(db, addr, validatorAddr)

	if number.Cmp(decrement) <= 0 {
		db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), []byte{})
	} else {
		number = new(big.Int).Sub(number, decrement)
		db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), number.Bytes())
	}
}

func GetPendingValidatorReward(db sdk.StateDBReader, addr basecommon.Address, validatorAddr basecommon.Address) *big.Int {
	value := db.GetState(addr, encodePendingValidatorRewardKey(validatorAddr))
	number := basecommon.Big0
	if len(value) != 0 {
		number = new(big.Int).SetBytes(value)
	}
	return number
}

func IncrementPendingDelegaterReward(db sdk.StateDB, addr basecommon.Address, delegaterAddr, validatorAddr basecommon.Address, increment *big.Int) {
	number := GetPendingDelegaterReward(db, addr, delegaterAddr, validatorAddr)
	number = new(big.Int).Add(number, increment)
	db.SetState(addr, encodePendingDelegaterRewardKey(delegaterAddr, validatorAddr), number.Bytes())
}

func DecrementPendingDelegaterReward(db sdk.StateDB, addr basecommon.Address, delegaterAddr, validatorAddr basecommon.Address, decrement *big.Int) {
	number := GetPendingDelegaterReward(db, addr, delegaterAddr, validatorAddr)

	if number.Cmp(decrement) <= 0 {
		db.SetState(addr, encodePendingDelegaterRewardKey(delegaterAddr, validatorAddr), []byte{})
	} else {
		number = new(big.Int).Sub(number, decrement)
		db.SetState(addr, encodePendingDelegaterRewardKey(delegaterAddr, validatorAddr), number.Bytes())
	}
}

func GetPendingDelegaterReward(db sdk.StateDBReader, addr basecommon.Address, delegaterAddr, validatorAddr basecommon.Address) *big.Int {
	value := db.GetState(addr, encodePendingDelegaterRewardKey(delegaterAddr, validatorAddr))
	number := basecommon.Big0
	if len(value) != 0 {
		number = new(big.Int).SetBytes(value)
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

//func GetDelegaterRewardPendingIndex(db sdk.StateDB, addr, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) uint64 {
//
//	value := db.GetState(addr, encodeDelegaterRewardPendingIndexKey(delegaterAddr, validatorAddr, stakeEpoch))
//
//	if len(value) == 0 {
//		return 0
//	}
//	return basecommon.BytesToUint64(value)
//}
//
//func SetDelegaterRewardPendingIndex(db sdk.StateDB, addr, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch, rewardEpoch uint64) {
//	db.SetState(addr, encodeDelegaterRewardPendingIndexKey(delegaterAddr, validatorAddr, stakeEpoch), basecommon.Uint64ToBytes(rewardEpoch))
//}

func GetEpochDelegationRewardPerShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, epoch uint64) *types.EpochDelegationRewardPerShareItem {
	value := db.GetState(addr, encodeEpochDelegationRewardPerShareItemKey(validatorAddr, epoch))
	if len(value) == 0 {
		return nil
	}

	var item types.EpochDelegationRewardPerShareItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

func SetEpochDelegationRewardPerShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, epoch uint64, item *types.EpochDelegationRewardPerShareItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeEpochDelegationRewardPerShareItemKey(validatorAddr, epoch), value)
	return nil
}
