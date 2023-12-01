package staking

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	stakingdb "github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
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

	if err := initValidators(statedb); nil != err {
		return err
	}

	return nil
}

func initValidatorPriority(statedb sdk.StateDB) error {
	head := types.NewPriorityValidator(
		stakingdb.EncodePriorityValidatorTailKey(),
		stakingdb.EncodePriorityValidatorTailKey(),
		common.ZeroAddr,
	)
	tail := types.NewPriorityValidator(
		stakingdb.EncodePriorityValidatorHeadKey(),
		stakingdb.EncodePriorityValidatorHeadKey(),
		common.ZeroAddr,
	)
	hvalue, err := rlp.EncodeToBytes(head)
	if nil != err {
		return stakingdb.ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tail)
	if nil != err {
		return stakingdb.ErrRlpEncode
	}

	statedb.SetState(address.StakeHandlerAddress, stakingdb.EncodePriorityValidatorHeadKey(), hvalue)
	statedb.SetState(address.StakeHandlerAddress, stakingdb.EncodePriorityValidatorTailKey(), tvalue)
	return nil
}

func initEpochItem(statedb sdk.StateDB) error {

	// TODO 需要根据配置读取 epochSize 计算第一轮的边界
	//
	zero := types.NewEpochItem(0, 0, 0)
	epoch := types.NewEpochItem(1, 25000, 10) // todo 需要重新计算边界
	zvalue, err := rlp.EncodeToBytes(zero)
	if nil != err {
		return stakingdb.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(epoch)
	if nil != err {
		return stakingdb.ErrRlpEncode
	}
	statedb.SetState(address.StakeHandlerAddress, stakingdb.EncodeEpochItemKey(0), zvalue)
	statedb.SetState(address.StakeHandlerAddress, stakingdb.EncodeEpochItemKey(1), value)

	return nil
}

func initRoundItem(statedb sdk.StateDB) error {

	// TODO 需要根据配置读取 roundSize 计算第一轮的边界
	//
	zero := types.NewRoundItem(0, 0)
	round := types.NewRoundItem(1, 250) // todo 需要重新计算边界
	zvalue, err := rlp.EncodeToBytes(zero)
	if nil != err {
		return stakingdb.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(round)
	if nil != err {
		return stakingdb.ErrRlpEncode
	}
	statedb.SetState(address.StakeHandlerAddress, stakingdb.EncodeRoundItemKey(0), zvalue)
	statedb.SetState(address.StakeHandlerAddress, stakingdb.EncodeRoundItemKey(1), value)

	return nil
}

func initValidators(statedb sdk.StateDB) error {

	// TODO 读取创世快， 添加创世的 priority / validator / epochValidatorSnapshotQueue(epoch:1)/ roundValidatorSnapshotQueue(round:1)

	return nil
}
