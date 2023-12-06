package staking

import (
	stakingdb "github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initStakeHandler(statedb sdk.StateDB, addr common.Address) error {

	if err := initValidatorPriority(statedb, addr); nil != err {
		return err
	}

	if err := initValidators(statedb, addr); nil != err {
		return err
	}

	return nil
}

func initValidatorPriority(statedb sdk.StateDB, addr common.Address) error {
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

	statedb.SetState(addr, stakingdb.EncodePriorityValidatorHeadKey(), hvalue)
	statedb.SetState(addr, stakingdb.EncodePriorityValidatorTailKey(), tvalue)
	return nil
}

func initValidators(statedb sdk.StateDB, addr common.Address) error {

	// TODO 读取创世快， 添加创世的 priority / validator / epochValidatorSnapshotQueue(epoch:1)/ roundValidatorSnapshotQueue(round:1)

	return nil
}
