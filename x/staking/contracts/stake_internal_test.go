package contracts

import (
	"context"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/mocks"
	"github.com/PlatONnetwork/PlatON-Go/common"
	basemock "github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"math/big"
	"testing"
)

func newStakeHandler(from common.Address) *StakeHandler {
	statedb := basemock.NewMockStateDB()
	evm := &vm.EVM{Context: vm.BlockContext{
		CanTransfer: func(db vm.StateDB, addr common.Address, amount *big.Int) bool {
			return db.GetBalance(addr).Cmp(amount) >= 0
		},
		Transfer: func(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
			db.SubBalance(sender, amount)
			db.AddBalance(recipient, amount)
		},
		Ctx: context.TODO(),
	},
		StateDB: statedb,
	}
	stakeHandler, _ := NewStakeHandler(evm, vm.NewContract(vm.AccountRef(from), vm.AccountRef(constants.StakeHandlerAddress), big.NewInt(0), 1000000), false)
	// new mock modules
	l1Module := mocks.NewMockMockL1Moduler(statedb)
	stageModule := mocks.NewMockStageModuler(statedb)
	stakeModule := mocks.NewMockStakeModuler(statedb, l1Module, stageModule)
	rewardModule := mocks.NewMockRewardModuler(statedb, stageModule, stakeModule)
	vrfModule := mocks.NewMockVRFModuler(statedb, stageModule, stakeModule)
	stakeModule.SetRewardModule(rewardModule)
	stakeModule.SetVRFModule(vrfModule)
	// set mock modules
	stakeHandler.SetStageModule(stageModule)
	stakeHandler.SetStakeModule(stakeModule)
	stakeHandler.SetRewardModule(rewardModule)
	//stakeHandler.Set
	return stakeHandler
}

func Test_Delegate(t *testing.T) {

	validatorAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	delegatorAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")
	ownerAddr := common.HexToAddress("0x3333333333333333333333333333333333333333")
	stakeHandler := newStakeHandler(common.ZeroAddr)
	var err error
	err = stakeHandler.stake(validatorAddr, ownerAddr, common.Big100, 10, []byte{}, enode.IDv0{})
	if nil != err {
		t.Error(err)
	}
	err = stakeHandler.delegate(validatorAddr, delegatorAddr, common.Big32)
	if nil != err {
		t.Error(err)
	}

	delegations, err := stakeHandler.GetDelegationsWithValidator([]common.Address{validatorAddr}, delegatorAddr)
	if nil != err {
		t.Error(err)
	}

	t.Log("delegation size", len(delegations))
}
