package contracts

import (
	"context"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/mocks"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	basemock "github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/test-go/testify/assert"
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

func newStakeHandler(from common.Address, blockNumber *big.Int, statedb vm.StateDB) *StakeHandler {

	evm := &vm.EVM{Context: vm.BlockContext{
		CanTransfer: func(db vm.StateDB, addr common.Address, amount *big.Int) bool {
			return db.GetBalance(addr).Cmp(amount) >= 0
		},
		Transfer: func(db vm.StateDB, sender, recipient common.Address, amount *big.Int) {
			db.SubBalance(sender, amount)
			db.AddBalance(recipient, amount)
		},
		Ctx:         context.TODO(),
		BlockNumber: blockNumber,
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

func getStakeForData() ([][]byte, map[common.Address]struct {
	Owner          common.Address
	StakeAmount    *big.Int
	CommissionRate *big.Int
	PubKey         []byte
	BlsKey         []byte
}) {

	bytes := [][]byte{
		// 	validator: 0x18D80f8B8302D8Fc41D4239d06b18b6f201640A8
		//	{
		//	    "blsKey":"0x8af81b3645b53841345d883389b78f32f943d57f6255ecb7130ef38866d333ef52a1dc92601f217e313975453c37839b",
		//	    "commissionRate":10,
		//	    "owner":"0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B",
		//	    "pubKey":"0x16863d2464e176212ed546ddcc13b42430a7e269c56e4d31ceedd6ddb3045e787052b334da1ccec549df45bea44101bf941e4f446b5b4caa04cf3a3675ab9df9",
		//	    "stakeAmount":200000000000
		//	}
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a6900000000000000000000000018d80f8b8302d8fc41d4239d06b18b6f201640a80000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000308af81b3645b53841345d883389b78f32f943d57f6255ecb7130ef38866d333ef52a1dc92601f217e313975453c37839b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004016863d2464e176212ed546ddcc13b42430a7e269c56e4d31ceedd6ddb3045e787052b334da1ccec549df45bea44101bf941e4f446b5b4caa04cf3a3675ab9df9"),
		// 	validator: 0x987CAcA3842e1C13645172373bFbDEd25A6F59f2
		//	{
		//	    "blsKey":"0x8cf17658d276b579cb12cb6910a2d348f6b88a8d07ee2297ef1b532ae5cf0185cbe76686321bb949b0b67aae5fdf6ff1",
		//	    "commissionRate":10,
		//	    "owner":"0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B",
		//	    "pubKey":"0xa2dc4a45559340d4b9051d9797b0d8a656c79036518590c634a003511c4f66d68b5fa13987b605a5c9793abf3fc29f2334cb87d4607ded49d17286f581b6e0fe",
		//	    "stakeAmount":200000000000
		//	}
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a69000000000000000000000000987caca3842e1c13645172373bfbded25a6f59f20000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000308cf17658d276b579cb12cb6910a2d348f6b88a8d07ee2297ef1b532ae5cf0185cbe76686321bb949b0b67aae5fdf6ff1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040a2dc4a45559340d4b9051d9797b0d8a656c79036518590c634a003511c4f66d68b5fa13987b605a5c9793abf3fc29f2334cb87d4607ded49d17286f581b6e0fe"),
		// 	validator: 0x35E1EC7b136DeF2E2d561F7E003deE2CBEf3C4c9
		//	{
		//	    "blsKey":"0xb969678ef2cf458b49b8c568d95e63221efe0f30383a9b0c5eb683bf2e23d118664631ce992d81ec4b6127ec0a760f86",
		//	    "commissionRate":10,
		//	    "owner":"0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B",
		//	    "pubKey":"0x77aefbb9d99a7924eaf53259050e1ba148a56712e00bc6a354b723cf707fa21aa329a4fead45d5c880d25a5f45adc344cbf4f9971d0009a7200416179a181972",
		//	    "stakeAmount":200000000000
		//	}
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a6900000000000000000000000035e1ec7b136def2e2d561f7e003dee2cbef3c4c90000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000030b969678ef2cf458b49b8c568d95e63221efe0f30383a9b0c5eb683bf2e23d118664631ce992d81ec4b6127ec0a760f8600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004077aefbb9d99a7924eaf53259050e1ba148a56712e00bc6a354b723cf707fa21aa329a4fead45d5c880d25a5f45adc344cbf4f9971d0009a7200416179a181972"),
		// 	validator: 0x5c8Fb7c7746417b551B953DEac5dE42E06871B2f
		//	{
		//	    "blsKey":"0xa33240d0c07e9a1abf747dfe4e87e54d25b2314bc33b8a700357e54a1fd4ced34f18ffe6ddc97e67a4f23655ecc38915",
		//	    "commissionRate":10,
		//	    "owner":"0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B",
		//	    "pubKey":"0x6b31afac957edd9a70c2671ff97493988620f605dbbd58ac47d9f2368f6eab6d3408ce00ebec3ad3bbcc1ef6931878d11f9974589726458995d6f90503ae1114",
		//	    "stakeAmount":200000000000
		//	}
		common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a690000000000000000000000005c8fb7c7746417b551b953deac5de42e06871b2f0000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000030a33240d0c07e9a1abf747dfe4e87e54d25b2314bc33b8a700357e54a1fd4ced34f18ffe6ddc97e67a4f23655ecc389150000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000406b31afac957edd9a70c2671ff97493988620f605dbbd58ac47d9f2368f6eab6d3408ce00ebec3ad3bbcc1ef6931878d11f9974589726458995d6f90503ae1114"),
	}
	cache := map[common.Address]struct {
		Owner          common.Address
		StakeAmount    *big.Int
		CommissionRate *big.Int
		PubKey         []byte
		BlsKey         []byte
	}{
		common.HexToAddress("0x18D80f8B8302D8Fc41D4239d06b18b6f201640A8"): {
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("16863d2464e176212ed546ddcc13b42430a7e269c56e4d31ceedd6ddb3045e787052b334da1ccec549df45bea44101bf941e4f446b5b4caa04cf3a3675ab9df9"),
			BlsKey:         common.Hex2Bytes("8af81b3645b53841345d883389b78f32f943d57f6255ecb7130ef38866d333ef52a1dc92601f217e313975453c37839b"),
		},
		common.HexToAddress("0x987CAcA3842e1C13645172373bFbDEd25A6F59f2"): {
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("a2dc4a45559340d4b9051d9797b0d8a656c79036518590c634a003511c4f66d68b5fa13987b605a5c9793abf3fc29f2334cb87d4607ded49d17286f581b6e0fe"),
			BlsKey:         common.Hex2Bytes("8cf17658d276b579cb12cb6910a2d348f6b88a8d07ee2297ef1b532ae5cf0185cbe76686321bb949b0b67aae5fdf6ff1"),
		},
		common.HexToAddress("0x35E1EC7b136DeF2E2d561F7E003deE2CBEf3C4c9"): {
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("77aefbb9d99a7924eaf53259050e1ba148a56712e00bc6a354b723cf707fa21aa329a4fead45d5c880d25a5f45adc344cbf4f9971d0009a7200416179a181972"),
			BlsKey:         common.Hex2Bytes("b969678ef2cf458b49b8c568d95e63221efe0f30383a9b0c5eb683bf2e23d118664631ce992d81ec4b6127ec0a760f86"),
		},
		common.HexToAddress("0x5c8Fb7c7746417b551B953DEac5dE42E06871B2f"): {
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("6b31afac957edd9a70c2671ff97493988620f605dbbd58ac47d9f2368f6eab6d3408ce00ebec3ad3bbcc1ef6931878d11f9974589726458995d6f90503ae1114"),
			BlsKey:         common.Hex2Bytes("a33240d0c07e9a1abf747dfe4e87e54d25b2314bc33b8a700357e54a1fd4ced34f18ffe6ddc97e67a4f23655ecc38915"),
		},
	}
	return bytes, cache
}

func Test_StakeFor(t *testing.T) {

	datas, cache := getStakeForData()
	epoch := uint64(1)
	from := common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10")
	statedb := stakeHandlerTestConfig.StateDB
	stakeHandler := newStakeHandler(from, new(big.Int).SetUint64(2), statedb.(vm.StateDB))
	stakeHandlerTestConfig.StageModule.MockCurrentEpoch(epoch)
	err := stakeHandlerTestConfig.StakeModule.MockInitValidatorGenesisPriority()
	assert.Nil(t, err, "Fialed to call MockInitValidatorGenesisPriority")
	initStakeHandler(stakeHandler, stakeHandlerTestConfig.StageModule, stakeHandlerTestConfig.StakeModule, stakeHandlerTestConfig.RewardModule)

	for _, data := range datas {
		err = stakeHandler.onStake(data)
		//if nil != err {
		//	t.Error(err)
		//}
		assert.Nil(t, err, "Fialed to call onStake")
	}
	next, queue, err := stakeHandler.GetValidators([]byte{}, common.Big100)
	//t.Log("next", hexutils.BytesToHex(next))
	//t.Log("queue", fmt.Sprintf("%v", queue))
	//t.Log("err", err)
	assert.Nil(t, err, "Fialed to call GetValidators")
	assert.Equal(t, 4, len(queue), "mismatching validator arr size")
	assert.Equal(t, []byte("priorityValidatorTail"), next, "mismatching next priority key")

	for i, validator := range queue {
		item, ok := cache[validator.ValidatorAddr]
		assert.True(t, ok, "Not found validator")
		assert.Equal(t, item.Owner, validator.Owner, fmt.Sprintf("mismatching owner, expect: %s, actual: %s", item.Owner.Hex(), validator.Owner.Hex()))
		assert.Equal(t, item.StakeAmount, validator.StakeAmount, fmt.Sprintf("mismatching stakeAmount, expect: %d, actual: %d", item.StakeAmount, validator.StakeAmount))
		assert.Equal(t, item.CommissionRate, validator.CommissionRate, fmt.Sprintf("mismatching commissionRate, expect: %d, actual: %d", item.CommissionRate, validator.CommissionRate))
		assert.Equal(t, common.Big0, validator.Status, fmt.Sprintf("mismatching status, expect: %d, actual: %d", common.Big0, validator.Status))
		assert.Equal(t, epoch, validator.Epoch.Uint64(), fmt.Sprintf("mismatching epoch, expect: %d, actual: %d", epoch, validator.Epoch))
		assert.Equal(t, uint64(i), validator.StakeIndex.Uint64(), fmt.Sprintf("mismatching stakeIndex, expect: %d, actual: %d", i, validator.StakeIndex))
		assert.Equal(t, item.PubKey, validator.PubKey, fmt.Sprintf("mismatching pubKey, expect: %s, actual: %s", hexutils.BytesToHex(item.PubKey), hexutils.BytesToHex(validator.PubKey)))
		assert.Equal(t, item.BlsKey, validator.BlsKey, fmt.Sprintf("mismatching blsKey, expect: %s, actual: %s", hexutils.BytesToHex(item.BlsKey), hexutils.BytesToHex(validator.BlsKey)))
	}
}
