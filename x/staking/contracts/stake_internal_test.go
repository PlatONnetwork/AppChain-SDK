package contracts

import (
	"context"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/mocks"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	basemock "github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"sync"
	"testing"
)

var (
	stakeHandlerTestConfig        *StakeHandlerTestConfig
	newStakeHandlerTestConfigOnce sync.Once
)

type StakeHandlerTestConfig struct {
	StateDB      sdk.StateDB
	L1Module     *mocks.MockL1Module
	StageModule  *mocks.MockStageModule
	StakeModule  *mocks.MockStakeModule
	RewardModule *mocks.MockRewardModule
	VrfModule    *mocks.MockVRFModule
}

func newStakeHandlerTestConfig() {
	newStakeHandlerTestConfigOnce.Do(
		func() {
			statedb := basemock.NewMockStateDB()
			// new mock modules
			l1Module := mocks.NewMockMockL1Module(statedb)
			stageModule := mocks.NewMockStageModule(statedb)
			stakeModule := mocks.NewMockStakeModule(statedb, l1Module, stageModule)
			rewardModule := mocks.NewMockRewardModule(statedb, stageModule, stakeModule)
			vrfModule := mocks.NewMockVRFModule(statedb, stageModule, stakeModule)
			stakeModule.SetRewardModule(rewardModule)
			stakeModule.SetVRFModule(vrfModule)

			stakeHandlerTestConfig = &StakeHandlerTestConfig{
				StateDB:      statedb,
				L1Module:     l1Module,
				StageModule:  stageModule,
				StakeModule:  stakeModule,
				RewardModule: rewardModule,
				VrfModule:    vrfModule,
			}
		},
	)
}

func init() {
	newStakeHandlerTestConfig()
}

func newStakeHandler(from common.Address, statedb vm.StateDB) *StakeHandler {

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

	//stakeHandler.Set
	return stakeHandler
}

func initStakeHandler(stakeHandler *StakeHandler, stageModule staketypes.StageModuler, stakeModule staketypes.StakeModuler, rewardModule staketypes.RewardModuler) {

	// set mock modules
	stakeHandler.SetStageModule(stageModule)
	stakeHandler.SetStakeModule(stakeModule)
	stakeHandler.SetRewardModule(rewardModule)
}

func Test_Delegate(t *testing.T) {

	//validatorAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	//delegatorAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")
	//ownerAddr := common.HexToAddress("0x3333333333333333333333333333333333333333")
	//stakeHandler := newStakeHandler(common.ZeroAddr)
	//var err error
	//err = stakeHandler.stake(validatorAddr, ownerAddr, common.Big100, 10, []byte{}, enode.IDv0{})
	//if nil != err {
	//	t.Error(err)
	//}
	//err = stakeHandler.delegate(validatorAddr, delegatorAddr, common.Big32)
	//if nil != err {
	//	t.Error(err)
	//}
	//
	//delegations, err := stakeHandler.GetDelegationsWithValidator([]common.Address{validatorAddr}, delegatorAddr)
	//if nil != err {
	//	t.Error(err)
	//}
	//
	//t.Log("delegation size", len(delegations))
}

func getStakeForData() [][]byte {
	return [][]byte{
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a6900000000000000000000000018d80f8b8302d8fc41d4239d06b18b6f201640a80000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000000005f5e100000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000308af81b3645b53841345d883389b78f32f943d57f6255ecb7130ef38866d333ef52a1dc92601f217e313975453c37839b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004016863d2464e176212ed546ddcc13b42430a7e269c56e4d31ceedd6ddb3045e787052b334da1ccec549df45bea44101bf941e4f446b5b4caa04cf3a3675ab9df9"),
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a69000000000000000000000000987caca3842e1c13645172373bfbded25a6f59f20000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000000005f5e100000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000308cf17658d276b579cb12cb6910a2d348f6b88a8d07ee2297ef1b532ae5cf0185cbe76686321bb949b0b67aae5fdf6ff1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040a2dc4a45559340d4b9051d9797b0d8a656c79036518590c634a003511c4f66d68b5fa13987b605a5c9793abf3fc29f2334cb87d4607ded49d17286f581b6e0fe"),
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a6900000000000000000000000035e1ec7b136def2e2d561f7e003dee2cbef3c4c90000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000000005f5e100000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000030b969678ef2cf458b49b8c568d95e63221efe0f30383a9b0c5eb683bf2e23d118664631ce992d81ec4b6127ec0a760f8600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004077aefbb9d99a7924eaf53259050e1ba148a56712e00bc6a354b723cf707fa21aa329a4fead45d5c880d25a5f45adc344cbf4f9971d0009a7200416179a181972"),
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a690000000000000000000000005c8fb7c7746417b551b953deac5de42e06871b2f0000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000000005f5e100000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000030a33240d0c07e9a1abf747dfe4e87e54d25b2314bc33b8a700357e54a1fd4ced34f18ffe6ddc97e67a4f23655ecc389150000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000406b31afac957edd9a70c2671ff97493988620f605dbbd58ac47d9f2368f6eab6d3408ce00ebec3ad3bbcc1ef6931878d11f9974589726458995d6f90503ae1114"),
	}
}

func Test_StakeFor(t *testing.T) {
	datas := getStakeForData()
	from := common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10")
	statedb := stakeHandlerTestConfig.StateDB
	stakeHandler := newStakeHandler(from, statedb.(vm.StateDB))
	initStakeHandler(stakeHandler, stakeHandlerTestConfig.StageModule, stakeHandlerTestConfig.StakeModule, stakeHandlerTestConfig.RewardModule)
	var err error
	for _, data := range datas {
		err = stakeHandler.onStake(data)
		if nil != err {
			t.Error(err)
		}
	}
}
