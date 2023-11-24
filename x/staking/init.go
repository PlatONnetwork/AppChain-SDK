package staking

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initStakeHandler(db sdk.StateDB) error {

	if err := initValidatorPriority(db); nil != err {
		return err
	}

	if err := initEpochItem(db); nil != err {
		return err
	}

	return nil
}

func initValidatorPriority(db sdk.StateDB) error {
	head := types.NewPriorityValidator(
		contracts.PriorityValidatorTailKey,
		contracts.PriorityValidatorTailKey,
		common.ZeroAddr,
	)
	tail := types.NewPriorityValidator(
		contracts.PriorityValidatorHeadKey,
		contracts.PriorityValidatorHeadKey,
		common.ZeroAddr,
	)
	hvalue, err := rlp.EncodeToBytes(head)
	if nil != err {
		return contracts.ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tail)
	if nil != err {
		return contracts.ErrRlpEncode
	}

	db.SetState(address.StakeHandlerAddres, contracts.PriorityValidatorHeadKey, hvalue)
	db.SetState(address.StakeHandlerAddres, contracts.PriorityValidatorTailKey, tvalue)
	return nil
}

func initEpochItem(db sdk.StateDB) error {

	// TODO 需要根据配置读取 epochSize 计算第一轮的边界
	//
	// tail -> head -> first -> tail -> head
	headEpoch := types.NewEpochItem(math.MaxUint64, 1, 0, 0)
	firstEpoch := types.NewEpochItem(0, math.MaxUint64, 1, 100) // todo 需要重新计算边界
	tailEpoch := types.NewEpochItem(1, 0, 0, 0)

	hvalue, err := rlp.EncodeToBytes(headEpoch)
	if nil != err {
		return contracts.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(firstEpoch)
	if nil != err {
		return contracts.ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tailEpoch)
	if nil != err {
		return contracts.ErrRlpEncode
	}
	db.SetState(address.StakeHandlerAddres, contracts.EncodeEpochItemKey(0), hvalue)
	db.SetState(address.StakeHandlerAddres, contracts.EncodeEpochItemKey(1), value)
	db.SetState(address.StakeHandlerAddres, contracts.EncodeEpochItemKey(math.MaxUint64), tvalue)
	return nil
}
