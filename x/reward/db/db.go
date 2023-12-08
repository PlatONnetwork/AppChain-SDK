package db

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
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

// Will be governed in the future
var (
	rewardPerBlockKey = []byte("rewardPerBlock")
	rewardPerEpochKey = []byte("rewardPerEpoch")
)

var (
	paidRewardPerEpochKeyPrefix     = []byte("paidRewardPerEpoch")     // "paidRewardPerEpoch":epochId => paidRewards
	pendingValidatorRewardKeyPrefix = []byte("pendingValidatorReward") // "pendingValidatorReward":validatorAddr => pendingRewards
	pendingDelegaterRewardKeyPrefix = []byte("pendingDelegaterReward") // "pendingDelegaterReward":delegaterAddr:validatorAddr => pendingRewards

	validatorRewardOwnerKeyPrefix = []byte("validatorRewardOwner") // "validatorRewardOwner":validatorAddr => ownerAddr

	//delegaterRewardPendingIndexKeyPrefix       = []byte("delegaterRewardPendingIndex")       // "delegaterRewardPendingIndex":delegaterAddr:validatorAddr:stakeEpoch => needRewardEpoch
	epochDelegationRewardPerShareItemKeyPrefix = []byte("epochDelegationRewardPerShareItem") // "epochDelegationRewardPerShareItem":validatorAddr:stakeEpoch:rewardEpoch => epochRewardPerDelegationShareItem{ preRewardEpoch, nextRewardEpoch, totalReward, perShareReward}
)

// ------

func EncodeRewardPerBlock() []byte {
	return rewardPerBlockKey
}
func EncodeRewardPerEpoch() []byte {
	return rewardPerEpochKey
}

// ------
func encodePaidRewardPerEpochKey(epoch uint64) []byte {
	return append(paidRewardPerEpochKeyPrefix, basecommon.Uint64ToBytes(epoch)...)
}

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

func encodeEpochDelegationRewardPerShareItemKey(validatorAddr basecommon.Address, stakeEpoch, rewardEpoch uint64) []byte {

	validatorAddrBytes := validatorAddr.Bytes()
	stakeEpochBytes := basecommon.Uint64ToBytes(stakeEpoch)
	rewardEpochBytes := basecommon.Uint64ToBytes(rewardEpoch)

	keyPrefixSize := len(epochDelegationRewardPerShareItemKeyPrefix)
	appendValidatorAddrSize := keyPrefixSize + len(validatorAddrBytes)
	appendStakeEpochSize := appendValidatorAddrSize + len(stakeEpochBytes)
	size := appendStakeEpochSize + len(rewardEpochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], epochDelegationRewardPerShareItemKeyPrefix)
	copy(key[keyPrefixSize:appendValidatorAddrSize], validatorAddrBytes)
	copy(key[appendValidatorAddrSize:appendStakeEpochSize], stakeEpochBytes)
	copy(key[appendStakeEpochSize:], rewardEpochBytes)

	return key
}

// ------------------------------------------------------ db methods ------------------------------------------------------

func SetRewardPerBlock(db sdk.StateDB, addr basecommon.Address, value *big.Int) {
	db.SetState(addr, EncodeRewardPerBlock(), value.Bytes())
}
func GetRewardPerBlock(db sdk.StateDBReader, addr basecommon.Address) *big.Int {
	value := db.GetState(addr, EncodeRewardPerBlock())
	if len(value) == 0 {
		return basecommon.Big0
	}
	return new(big.Int).SetBytes(value)
}

func SetRewardPerEpoch(db sdk.StateDB, addr basecommon.Address, value *big.Int) {
	db.SetState(addr, EncodeRewardPerEpoch(), value.Bytes())
}
func GetRewardPerEpoch(db sdk.StateDBReader, addr basecommon.Address) *big.Int {
	value := db.GetState(addr, EncodeRewardPerEpoch())
	if len(value) == 0 {
		return basecommon.Big0
	}
	return new(big.Int).SetBytes(value)
}

// ------

func GetPaidRewardPerEpoch(db sdk.StateDB, addr basecommon.Address, epoch uint64) *big.Int {
	value := db.GetState(addr, encodePaidRewardPerEpochKey(epoch))
	number := basecommon.Big0
	if len(value) != 0 {
		number = new(big.Int).SetBytes(value)
	}
	return number
}

func IncrementPaidRewardPerEpoch(db sdk.StateDB, addr basecommon.Address, epoch uint64, increment *big.Int) {
	number := GetPaidRewardPerEpoch(db, addr, epoch)
	number = new(big.Int).Add(number, increment)
	db.SetState(addr, encodePaidRewardPerEpochKey(epoch), number.Bytes())
}

func IncrementPendingValidatorReward(db sdk.StateDB, addr, validatorAddr basecommon.Address, increment *big.Int) {
	number := GetPendingValidatorReward(db, addr, validatorAddr)
	number = new(big.Int).Add(number, increment)
	db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), number.Bytes())
}

func DecrementPendingValidatorReward(db sdk.StateDB, addr, validatorAddr basecommon.Address, decrement *big.Int) {
	number := GetPendingValidatorReward(db, addr, validatorAddr)

	if number.Cmp(decrement) <= 0 {
		db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), []byte{})
	} else {
		number = new(big.Int).Sub(number, decrement)
		db.SetState(addr, encodePendingValidatorRewardKey(validatorAddr), number.Bytes())
	}
}

func GetPendingValidatorReward(db sdk.StateDBReader, addr, validatorAddr basecommon.Address) *big.Int {
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

func GetValidatorRewardOwner(db sdk.StateDB, addr, validatorAddr basecommon.Address) basecommon.Address {
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

func GetEpochDelegationRewardPerShareItem(db sdk.StateDBReader, addr, validatorAddr basecommon.Address, stakeEpoch, rewardEpoch uint64) *types.EpochDelegationRewardPerShareItem {
	value := db.GetState(addr, encodeEpochDelegationRewardPerShareItemKey(validatorAddr, stakeEpoch, rewardEpoch))
	if len(value) == 0 {
		return nil
	}

	var item types.EpochDelegationRewardPerShareItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return &item
	}
	return nil
}

func SetEpochDelegationRewardPerShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, stakeEpoch, rewardEpoch uint64, item *types.EpochDelegationRewardPerShareItem) error {
	value, err := rlp.EncodeToBytes(item)
	if nil != err {
		return ErrRlpEncode
	}
	db.SetState(addr, encodeEpochDelegationRewardPerShareItemKey(validatorAddr, stakeEpoch, rewardEpoch), value)
	return nil
}

func removeEpochDelegationRewardPerShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, stakeEpoch, rewardEpoch uint64) {
	db.SetState(addr, encodeEpochDelegationRewardPerShareItemKey(validatorAddr, stakeEpoch, rewardEpoch), []byte{})
}

func AppendEpochDelegationRewardPerShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, stakeEpoch, rewardEpoch uint64, totalReward, perShareReward *big.Int) error {

	indexEpoch := uint64(math.MaxUint64)
	indexItem := GetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexEpoch)

	// first insert
	if indexItem.IsEmpty() {

		preEpoch := uint64(0)
		// tail -> head -> epoch -> tail -> head

		preItem := types.NewEpochDelegationRewardPerShareItem(indexEpoch, rewardEpoch, basecommon.Big0, basecommon.Big0) // head
		epochItem := types.NewEpochDelegationRewardPerShareItem(preEpoch, indexEpoch, totalReward, perShareReward)       // item
		indexItem = types.NewEpochDelegationRewardPerShareItem(rewardEpoch, preEpoch, basecommon.Big0, basecommon.Big0)  // tail

		if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, preEpoch, preItem); nil != err {
			return err
		}
		if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, rewardEpoch, epochItem); nil != err {
			return err
		}
		if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexEpoch, indexItem); nil != err {
			return err
		}

		return nil
	}

	// range from tail to head
	for indexItem.PreRewardEpoch != math.MaxUint64 { // not as head
		if indexEpoch == rewardEpoch {
			// pre -> index(epoch) -> next
			// index == epoch

			indexItem.TotalReward = totalReward
			indexItem.PerShareReward = perShareReward
			if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexEpoch, indexItem); nil != err {
				return err
			}
			break
		} else if indexEpoch < rewardEpoch {
			// pre -> index -> epoch -> next... -> tail(max)
			// pre < index < epoch < next ... < tail(max)

			nextItem := GetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexItem.NextRewardEpoch)
			epochItem := types.NewEpochDelegationRewardPerShareItem(indexEpoch, indexItem.NextRewardEpoch, totalReward, perShareReward)

			indexItem.UpdateNextRewardEpoch(rewardEpoch)
			nextItem.UpdatePreRewardEpoch(rewardEpoch)

			if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexEpoch, indexItem); nil != err {
				return err
			}
			if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, rewardEpoch, epochItem); nil != err {
				return err
			}
			if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, epochItem.NextRewardEpoch, nextItem); nil != err {
				return err
			}

			break

		} else {

			if indexItem.PreRewardEpoch == uint64(0) {
				// if  head(min) -> index -> tail(max)
				// and epoch < index
				//
				// then: head(min) -> epoch -> index<last one> -> ... -> tail(max)
				// head < epoch < index < ... < tail

				preItem := GetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexItem.PreRewardEpoch) // head
				epochItem := types.NewEpochDelegationRewardPerShareItem(indexItem.PreRewardEpoch, rewardEpoch, totalReward, perShareReward)

				preItem.UpdateNextRewardEpoch(rewardEpoch)
				indexItem.UpdatePreRewardEpoch(rewardEpoch)

				if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, epochItem.PreRewardEpoch, preItem); nil != err {
					return err
				}
				if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, rewardEpoch, epochItem); nil != err {
					return err
				}
				if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexEpoch, indexItem); nil != err {
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

		indexEpoch = indexItem.PreRewardEpoch
		indexItem = GetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, indexEpoch)
	}
	return nil
}

func ReleaseEpochDelegationRewardPerShareItem(db sdk.StateDB, addr, validatorAddr basecommon.Address, stakeEpoch, rewardEpoch uint64, decrement *big.Int) error {
	item := GetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, rewardEpoch)
	if nil == item {
		return ErrNotFound
	}

	item.DecrementTotalReward(decrement)

	// change index
	if item.IsZeroTotalReward() {
		pre := GetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, item.PreRewardEpoch)
		next := GetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, item.NextRewardEpoch)

		// remove the last one   tail -> head -> lastone(remove) -> tail -> head
		if pre.PreRewardEpoch == math.MaxUint64 && next.NextRewardEpoch == 0 {
			removeEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, item.PreRewardEpoch)
			removeEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, item.NextRewardEpoch)
		} else {
			pre.UpdateNextRewardEpoch(item.NextRewardEpoch)
			next.UpdatePreRewardEpoch(item.PreRewardEpoch)
			if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, item.PreRewardEpoch, pre); nil != err {
				return err
			}
			if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, item.NextRewardEpoch, next); nil != err {
				return err
			}
		}
		removeEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, rewardEpoch)
	} else {
		if err := SetEpochDelegationRewardPerShareItem(db, addr, validatorAddr, stakeEpoch, rewardEpoch, item); nil != err {
			return err
		}
	}
	return nil
}
