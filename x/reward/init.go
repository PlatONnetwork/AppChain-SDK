package reward

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/config"
	rewarddb "github.com/PlatONnetwork/AppChain-SDK/x/reward/db"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initConfigParams(statedb sdk.StateDB, addr common.Address, params *config.RewardNetworkParams) {
	statedb.SetState(addr, rewarddb.EncodeRewardPerBlock(), params.RewardPerBlock.Bytes())
	statedb.SetState(addr, rewarddb.EncodeRewardPerEpoch(), params.RewardPerEpoch.Bytes())
}
