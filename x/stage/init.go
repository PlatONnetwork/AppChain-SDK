package stage

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/config"
	stagedb "github.com/PlatONnetwork/AppChain-SDK/x/stage/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initGenesisEpochItem(statedb sdk.StateDB, addr common.Address, params *config.StageNetworkParams) error {

	zero := types.NewEpochItem(0, 0, 0)
	epoch := types.NewEpochItem(1, params.EpochSize, params.EpochSize/params.RoundSize)
	zvalue, err := rlp.EncodeToBytes(zero)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(epoch)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	statedb.SetState(addr, stagedb.EncodeEpochItemKey(0), zvalue) // 0 epoch
	statedb.SetState(addr, stagedb.EncodeEpochItemKey(1), value)  // ist epoch

	return nil
}

func initGenesisRoundItem(statedb sdk.StateDB, addr common.Address, params *config.StageNetworkParams) error {

	zero := types.NewRoundItem(0, 0)
	round := types.NewRoundItem(1, params.RoundSize)
	zvalue, err := rlp.EncodeToBytes(zero)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	value, err := rlp.EncodeToBytes(round)
	if nil != err {
		return stagedb.ErrRlpEncode
	}
	statedb.SetState(addr, stagedb.EncodeRoundItemKey(0), zvalue) // 0 round
	statedb.SetState(addr, stagedb.EncodeRoundItemKey(1), value)  // 1st round

	return nil
}
