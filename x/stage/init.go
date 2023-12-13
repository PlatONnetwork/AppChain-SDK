package stage

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/config"
	stagedb "github.com/PlatONnetwork/AppChain-SDK/x/stage/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initConfigParams(statedb sdk.StateDB, addr common.Address, params *config.StageNetworkParams) {
	statedb.SetState(addr, stagedb.EncodeRoundValidatorElectionDistanceKey(), common.Uint64ToBytes(params.RoundValidatorElectionDistance))
	statedb.SetState(addr, stagedb.EncodeRoundSizeKey(), common.Uint64ToBytes(params.RoundSize))
	statedb.SetState(addr, stagedb.EncodeEpochSizeKey(), common.Uint64ToBytes(params.EpochSize))
}

func initGenesisEpochItem(statedb sdk.StateDB, addr common.Address, params *config.StageNetworkParams) error {

	fmt.Printf("addr: %s", addr.Hex())

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

	fmt.Printf("addr: %s", addr.Hex())

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
