package reward

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type RewardModule struct {
	logger log.Logger
}

func NewRewardModule() *RewardModule {
	return &RewardModule{
		logger: log.New("module", "reward"),
	}
}

func (r *RewardModule) Name() string {
	return "reward"
}

func (r *RewardModule) Address() basecommon.Address {
	return address.RewardManagerAddress
}

func (r *RewardModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	rewardManager, _ := contracts.NewRewardManager(evm, contract, readOnly)
	return rewardManager.Run(input)
}

func (r *RewardModule) BeginBlock(ctx sdk.WorkerContext) {

}

func (r *RewardModule) EndBlock(ctx sdk.WorkerContext) {

}
