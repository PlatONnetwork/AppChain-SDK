package stage

import (
	stagedb "github.com/PlatONnetwork/AppChain-SDK/x/stage/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initEpochItem(statedb sdk.StateDB, addr common.Address) error {

	// TODO 需要根据配置读取 epochSize 计算第一轮的边界
	//
	zero := types.NewEpochItem(0, 0, 0)
	epoch := types.NewEpochItem(1, 25000, 10) // todo 需要重新计算边界
	zvalue, err := rlp.EncodeToBytes(zero)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(epoch)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	statedb.SetState(addr, stagedb.EncodeEpochItemKey(0), zvalue)
	statedb.SetState(addr, stagedb.EncodeEpochItemKey(1), value)

	return nil
}

func initRoundItem(statedb sdk.StateDB, addr common.Address) error {

	// TODO 需要根据配置读取 roundSize 计算第一轮的边界
	//
	zero := types.NewRoundItem(0, 0)
	round := types.NewRoundItem(1, 250) // todo 需要重新计算边界
	zvalue, err := rlp.EncodeToBytes(zero)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(round)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	statedb.SetState(addr, stagedb.EncodeRoundItemKey(0), zvalue)
	statedb.SetState(addr, stagedb.EncodeRoundItemKey(1), value)

	return nil
}
