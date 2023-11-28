package staking

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initStakeHandler(statedb sdk.StateDB) error {

	if err := initValidatorPriority(statedb); nil != err {
		return err
	}

	if err := initEpochItem(statedb); nil != err {
		return err
	}
	if err := initRoundItem(statedb); nil != err {
		return err
	}

	return nil
}

func initValidatorPriority(statedb sdk.StateDB) error {
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
		return db.ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tail)
	if nil != err {
		return db.ErrRlpEncode
	}

	statedb.SetState(address.StakeHandlerAddres, contracts.PriorityValidatorHeadKey, hvalue)
	statedb.SetState(address.StakeHandlerAddres, contracts.PriorityValidatorTailKey, tvalue)
	return nil
}

func initEpochItem(statedb sdk.StateDB) error {

	// TODO 需要根据配置读取 epochSize 计算第一轮的边界
	//
	// tail -> head -> first -> tail -> head
	headEpoch := types.NewEpochItem(math.MaxUint64, 1, 0, 0, 0)
	firstEpoch := types.NewEpochItem(0, math.MaxUint64, 1, 100, 10) // todo 需要重新计算边界
	tailEpoch := types.NewEpochItem(1, 0, 0, 0, 0)

	hvalue, err := rlp.EncodeToBytes(headEpoch)
	if nil != err {
		return db.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(firstEpoch)
	if nil != err {
		return db.ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tailEpoch)
	if nil != err {
		return db.ErrRlpEncode
	}
	statedb.SetState(address.StakeHandlerAddres, db.EncodeEpochItemKey(0), hvalue)
	statedb.SetState(address.StakeHandlerAddres, db.EncodeEpochItemKey(1), value)
	statedb.SetState(address.StakeHandlerAddres, db.EncodeEpochItemKey(math.MaxUint64), tvalue)
	return nil
}

func initRoundItem(statedb sdk.StateDB) error {

	// TODO 需要根据配置读取 roundSize 计算第一轮的边界
	//
	// tail -> head -> first -> tail -> head
	headRound := types.NewRoundItem(math.MaxUint64, 1, 0, 0)
	firstRound := types.NewRoundItem(0, math.MaxUint64, 1, 100) // todo 需要重新计算边界
	tailRound := types.NewRoundItem(1, 0, 0, 0)

	hvalue, err := rlp.EncodeToBytes(headRound)
	if nil != err {
		return db.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(firstRound)
	if nil != err {
		return db.ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tailRound)
	if nil != err {
		return db.ErrRlpEncode
	}
	statedb.SetState(address.StakeHandlerAddres, db.EncodeRoundItemKey(0), hvalue)
	statedb.SetState(address.StakeHandlerAddres, db.EncodeRoundItemKey(1), value)
	statedb.SetState(address.StakeHandlerAddres, db.EncodeRoundItemKey(math.MaxUint64), tvalue)
	return nil
}
