package contracts

import (
	"context"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/mocks"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	basemock "github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/test-go/testify/assert"
	"math/big"
	"testing"
)

func stakePrepare(t *testing.T, genesisEpoch uint64, caller common.Address, check bool) *StakeHandler {
	stakeHandlerTestConfig := newStakeHandlerTestConfig()
	// init stakeHandler
	stakeHandler := newStakeHandler((stakeHandlerTestConfig.StateDB).(vm.StateDB), caller, new(big.Int).SetUint64(2))
	initStakeHandler(stakeHandler, stakeHandlerTestConfig.L1Module, stakeHandlerTestConfig.StageModule, stakeHandlerTestConfig.StakeModule, stakeHandlerTestConfig.RewardModule)

	// mock init genesis
	stakeHandlerTestConfig.StageModule.MockCurrentEpoch(genesisEpoch)
	err := stakeHandlerTestConfig.StakeModule.MockInitValidatorGenesisPriority()
	if check {
		assert.Nil(t, err, "Failed to call MockInitValidatorGenesisPriority")
	}
	return stakeHandler
}

type StakeHandlerTestConfig struct {
	StateDB      sdk.StateDB
	L1Module     *mocks.MockL1Module
	StageModule  *mocks.MockStageModule
	StakeModule  *mocks.MockStakeModule
	RewardModule *mocks.MockRewardModule
	VrfModule    *mocks.MockVRFModule
}

func newStakeHandlerTestConfig() *StakeHandlerTestConfig {
	statedb := basemock.NewMockStateDB()
	// new mock modules
	l1Module := mocks.NewMockMockL1Module(statedb)
	stageModule := mocks.NewMockStageModule(statedb)
	stakeModule := mocks.NewMockStakeModule(statedb, l1Module, stageModule)
	rewardModule := mocks.NewMockRewardModule(statedb, stageModule, stakeModule)
	vrfModule := mocks.NewMockVRFModule(statedb, stageModule, stakeModule)
	stakeModule.SetRewardModule(rewardModule)
	stakeModule.SetVRFModule(vrfModule)

	return &StakeHandlerTestConfig{
		StateDB:      statedb,
		L1Module:     l1Module,
		StageModule:  stageModule,
		StakeModule:  stakeModule,
		RewardModule: rewardModule,
		VrfModule:    vrfModule,
	}
}

func newStakeHandler(statedb vm.StateDB, from common.Address, blockNumber *big.Int) *StakeHandler {

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

func initStakeHandler(stakeHandler *StakeHandler, l1Module staketypes.L1Moduler, stageModule staketypes.StageModuler, stakeModule staketypes.StakeModuler, rewardModule staketypes.RewardModuler) {

	// set mock modules
	stakeHandler.SetL1Module(l1Module)
	stakeHandler.SetStageModule(stageModule)
	stakeHandler.SetStakeModule(stakeModule)
	stakeHandler.SetRewardModule(rewardModule)
}

func setNumberOfBlocksForRoundValidator(statedb vm.StateDB, validatorAddr common.Address, round, increment uint64) error {
	db.IncrementNumberOfBlocksForRoundValidator(statedb, constants.StakeHandlerAddress, validatorAddr, round, increment)
	return nil
}
func getNumberOfBlocksForRoundValidator(statedb vm.StateDB, validatorAddr common.Address, round uint64) uint64 {
	return db.GetNumberOfBlocksForRoundValidator(statedb, constants.StakeHandlerAddress, validatorAddr, round)
}

func checkLowBlocksValidatorForPreviousRound(statedb vm.StateDB, currentRound, minBlocksOfRoundValidator uint64) error {
	// check low blocks validators
	lowBlocksValidatorAddrQueue := db.CheckLowBlocksValidatorForPreviousRound(statedb, constants.StakeHandlerAddress, currentRound, minBlocksOfRoundValidator)
	// update validator status
	for _, validatorAddr := range lowBlocksValidatorAddrQueue {
		if err := updateValidatorStatus(statedb, validatorAddr, staketypes.Invalided|staketypes.LowBlocks); nil != err {
			return fmt.Errorf("can not update validator status to [lowBlocks], %s, validator: %s", err, validatorAddr.Hex())
		}
	}
	return nil
}

func updateValidatorStatus(statedb vm.StateDB, validatorAddr common.Address, status staketypes.ValidatorStatus) error {
	validator := db.GetValidator(statedb, constants.StakeHandlerAddress, validatorAddr)
	if validator.IsEmpty() {
		return nil
	}

	validator.AppendStatus(status)

	if status.IsInvalid() {

		// delete old priority
		priority := db.GetValidatorPriority(statedb, constants.StakeHandlerAddress, validator.Epoch, validator.StakeIndex, validator.Shares())
		if priority.IsNotEmpty() {
			if priority.ValidatorAddr != validatorAddr {
				return db.ErrMisMatching
			}
			if err := db.RemoveValidatorPriority(statedb, constants.StakeHandlerAddress, validator.Epoch, validator.StakeIndex, validator.Shares()); nil != err {
				return err
			}
		}
	}
	// set new validator information only
	return db.SetValidator(statedb, constants.StakeHandlerAddress, validatorAddr, validator)
}

func rankPriorityValidatorIds(statedb vm.StateDB, maxEpochValidatorsSize uint64) []common.Address {
	return db.RankPriorityValidatorIds(statedb, constants.StakeHandlerAddress, maxEpochValidatorsSize)
}

// -------------------------------------------------------------

func encodeStakeForData(stakeData *(struct {
	ValidatorAddr  common.Address
	Owner          common.Address
	StakeAmount    *big.Int
	CommissionRate *big.Int
	PubKey         []byte
	BlsKey         []byte
})) []byte {

	input, _ := STAKE_PARAMS_TYPE.Encode([]interface{}{STAKE_SIG, stakeData.ValidatorAddr, stakeData.Owner, stakeData.StakeAmount, stakeData.CommissionRate, stakeData.BlsKey, stakeData.PubKey})
	return input
}

func encodeAddSatkeData(addStakeData *(struct {
	ValidatorAddr common.Address
	Amount        *big.Int
	//AddStakeData []byte
})) []byte {

	input, _ := ADDSTAKE_PARAMS_TYPE.Encode([]interface{}{ADDSTAKE_SIG, addStakeData.ValidatorAddr, addStakeData.Amount})
	return input
}

func encodeDelegateData(validator common.Address, delegateData *(struct {
	Delegator common.Address
	Amount    *big.Int
})) []byte {

	input, _ := DELEGATE_PARAMS_TYPE.Encode([]interface{}{DELEGATE_SIG, validator, delegateData.Delegator, delegateData.Amount})
	return input
}

func extractStakeForDataList() ([]common.Address, [][]byte) {

	addrs := make([]common.Address, 0)
	datas := make([][]byte, 0)

	for addr, data := range getStakeForDataCache() {
		addrs = append(addrs, addr)

		datas = append(datas, encodeStakeForData(data))
	}
	return addrs, datas
}

func getStakeForDataCache() map[common.Address]*(struct {
	ValidatorAddr  common.Address
	Owner          common.Address
	StakeAmount    *big.Int
	CommissionRate *big.Int
	PubKey         []byte
	BlsKey         []byte
	//StakeData      []byte
}) {

	return map[common.Address]*(struct {
		ValidatorAddr common.Address
		// fields
		Owner          common.Address
		StakeAmount    *big.Int
		CommissionRate *big.Int
		PubKey         []byte
		BlsKey         []byte
		// for stake
		//StakeData []byte
	}){
		common.HexToAddress("0x18D80f8B8302D8Fc41D4239d06b18b6f201640A8"): {
			ValidatorAddr:  common.HexToAddress("0x18D80f8B8302D8Fc41D4239d06b18b6f201640A8"),
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("16863d2464e176212ed546ddcc13b42430a7e269c56e4d31ceedd6ddb3045e787052b334da1ccec549df45bea44101bf941e4f446b5b4caa04cf3a3675ab9df9"),
			BlsKey:         common.Hex2Bytes("8af81b3645b53841345d883389b78f32f943d57f6255ecb7130ef38866d333ef52a1dc92601f217e313975453c37839b"),
			//StakeData:      common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a6900000000000000000000000018d80f8b8302d8fc41d4239d06b18b6f201640a80000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000308af81b3645b53841345d883389b78f32f943d57f6255ecb7130ef38866d333ef52a1dc92601f217e313975453c37839b00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004016863d2464e176212ed546ddcc13b42430a7e269c56e4d31ceedd6ddb3045e787052b334da1ccec549df45bea44101bf941e4f446b5b4caa04cf3a3675ab9df9"),
		},
		common.HexToAddress("0x987CAcA3842e1C13645172373bFbDEd25A6F59f2"): {
			ValidatorAddr:  common.HexToAddress("0x987CAcA3842e1C13645172373bFbDEd25A6F59f2"),
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("a2dc4a45559340d4b9051d9797b0d8a656c79036518590c634a003511c4f66d68b5fa13987b605a5c9793abf3fc29f2334cb87d4607ded49d17286f581b6e0fe"),
			BlsKey:         common.Hex2Bytes("8cf17658d276b579cb12cb6910a2d348f6b88a8d07ee2297ef1b532ae5cf0185cbe76686321bb949b0b67aae5fdf6ff1"),
			//StakeData:      common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a69000000000000000000000000987caca3842e1c13645172373bfbded25a6f59f20000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e0000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000308cf17658d276b579cb12cb6910a2d348f6b88a8d07ee2297ef1b532ae5cf0185cbe76686321bb949b0b67aae5fdf6ff1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000040a2dc4a45559340d4b9051d9797b0d8a656c79036518590c634a003511c4f66d68b5fa13987b605a5c9793abf3fc29f2334cb87d4607ded49d17286f581b6e0fe"),
		},
		common.HexToAddress("0x35E1EC7b136DeF2E2d561F7E003deE2CBEf3C4c9"): {
			ValidatorAddr:  common.HexToAddress("0x35E1EC7b136DeF2E2d561F7E003deE2CBEf3C4c9"),
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("77aefbb9d99a7924eaf53259050e1ba148a56712e00bc6a354b723cf707fa21aa329a4fead45d5c880d25a5f45adc344cbf4f9971d0009a7200416179a181972"),
			BlsKey:         common.Hex2Bytes("b969678ef2cf458b49b8c568d95e63221efe0f30383a9b0c5eb683bf2e23d118664631ce992d81ec4b6127ec0a760f86"),
			//StakeData:      common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a6900000000000000000000000035e1ec7b136def2e2d561f7e003dee2cbef3c4c90000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000030b969678ef2cf458b49b8c568d95e63221efe0f30383a9b0c5eb683bf2e23d118664631ce992d81ec4b6127ec0a760f8600000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004077aefbb9d99a7924eaf53259050e1ba148a56712e00bc6a354b723cf707fa21aa329a4fead45d5c880d25a5f45adc344cbf4f9971d0009a7200416179a181972"),
		},
		common.HexToAddress("0x5c8Fb7c7746417b551B953DEac5dE42E06871B2f"): {
			ValidatorAddr:  common.HexToAddress("0x5c8Fb7c7746417b551B953DEac5dE42E06871B2f"),
			Owner:          common.HexToAddress("0x1eEEBc3900803a28ca6E68Eb98FDeCf98350D97B"),
			StakeAmount:    new(big.Int).SetUint64(200000000000),
			CommissionRate: new(big.Int).SetUint64(10),
			PubKey:         common.Hex2Bytes("6b31afac957edd9a70c2671ff97493988620f605dbbd58ac47d9f2368f6eab6d3408ce00ebec3ad3bbcc1ef6931878d11f9974589726458995d6f90503ae1114"),
			BlsKey:         common.Hex2Bytes("a33240d0c07e9a1abf747dfe4e87e54d25b2314bc33b8a700357e54a1fd4ced34f18ffe6ddc97e67a4f23655ecc38915"),
			//StakeData:      common.Hex2Bytes("1bcc0f4c3fad314e585165815f94ecca9b96690a26d6417d7876448a9a867a690000000000000000000000005c8fb7c7746417b551b953deac5de42e06871b2f0000000000000000000000001eeebc3900803a28ca6e68eb98fdecf98350d97b0000000000000000000000000000000000000000000000000000002e90edd000000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000000000000000000e000000000000000000000000000000000000000000000000000000000000001400000000000000000000000000000000000000000000000000000000000000030a33240d0c07e9a1abf747dfe4e87e54d25b2314bc33b8a700357e54a1fd4ced34f18ffe6ddc97e67a4f23655ecc389150000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000406b31afac957edd9a70c2671ff97493988620f605dbbd58ac47d9f2368f6eab6d3408ce00ebec3ad3bbcc1ef6931878d11f9974589726458995d6f90503ae1114"),
		},
	}
}

func extractAddStakeDataList() ([]common.Address, [][]byte) {
	addrs := make([]common.Address, 0)
	datas := make([][]byte, 0)

	for addr, data := range getAddStakeDataCache() {
		addrs = append(addrs, addr)
		datas = append(datas, encodeAddSatkeData(data))
	}
	return addrs, datas
}

func getAddStakeDataCache() map[common.Address]*(struct {
	ValidatorAddr common.Address
	Amount        *big.Int
	//AddStakeData []byte
}) {

	return map[common.Address]*(struct {
		ValidatorAddr common.Address
		Amount        *big.Int
		//AddStakeData []byte
	}){
		common.HexToAddress("0x18D80f8B8302D8Fc41D4239d06b18b6f201640A8"): {
			ValidatorAddr: common.HexToAddress("0x18D80f8B8302D8Fc41D4239d06b18b6f201640A8"),
			Amount:        new(big.Int).SetUint64(100000000),
			//AddStakeData: common.Hex2Bytes("7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee7300000000000000000000000018d80f8b8302d8fc41d4239d06b18b6f201640a80000000000000000000000000000000000000000000000000000000005f5e100"),
		},
		common.HexToAddress("0x987CAcA3842e1C13645172373bFbDEd25A6F59f2"): {
			ValidatorAddr: common.HexToAddress("0x987CAcA3842e1C13645172373bFbDEd25A6F59f2"),
			Amount:        new(big.Int).SetUint64(200000000),
			//AddStakeData: common.Hex2Bytes("7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73000000000000000000000000987caca3842e1c13645172373bfbded25a6f59f2000000000000000000000000000000000000000000000000000000000bebc200"),
		},
		common.HexToAddress("0x35E1EC7b136DeF2E2d561F7E003deE2CBEf3C4c9"): {
			ValidatorAddr: common.HexToAddress("0x35E1EC7b136DeF2E2d561F7E003deE2CBEf3C4c9"),
			Amount:        new(big.Int).SetUint64(300000000),
			//AddStakeData: common.Hex2Bytes("7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee7300000000000000000000000035e1ec7b136def2e2d561f7e003dee2cbef3c4c90000000000000000000000000000000000000000000000000000000011e1a300"),
		},
		common.HexToAddress("0x5c8Fb7c7746417b551B953DEac5dE42E06871B2f"): {
			ValidatorAddr: common.HexToAddress("0x5c8Fb7c7746417b551B953DEac5dE42E06871B2f"),
			Amount:        new(big.Int).SetUint64(600000000),
			//AddStakeData: common.Hex2Bytes("7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee730000000000000000000000005c8fb7c7746417b551b953deac5de42e06871b2f0000000000000000000000000000000000000000000000000000000023c34600"),
		},
	}
}

func extractDelegateDatList() []common.Address {
	addrs := make([]common.Address, 0)

	for addr := range getDelegateDataCache() {
		addrs = append(addrs, addr)

	}
	return addrs
}

func getDelegateDataCache() map[common.Address]*(struct {
	Delegator common.Address
	Amount    *big.Int
}) {

	return map[common.Address]*(struct {
		Delegator common.Address
		Amount    *big.Int
	}){
		common.HexToAddress("0xfBE8A5b63aa33195d7C848d0c7adAe5f244Fb251"): {
			Delegator: common.HexToAddress("0xfBE8A5b63aa33195d7C848d0c7adAe5f244Fb251"),
			Amount:    new(big.Int).SetUint64(100000000),
		},
		common.HexToAddress("0xB318affA5d1bd442925F1a269b36cFC7f49f9a82"): {
			Delegator: common.HexToAddress("0xB318affA5d1bd442925F1a269b36cFC7f49f9a82"),
			Amount:    new(big.Int).SetUint64(200000000),
		},
		common.HexToAddress("0x71004B169C727933Ccef48A85eB66a788543854e"): {
			Delegator: common.HexToAddress("0x71004B169C727933Ccef48A85eB66a788543854e"),
			Amount:    new(big.Int).SetUint64(300000000),
		},
		common.HexToAddress("0x1fB7bfae24977dB4A5612D24281b7aCE65175e96"): {
			Delegator: common.HexToAddress("0x1fB7bfae24977dB4A5612D24281b7aCE65175e96"),
			Amount:    new(big.Int).SetUint64(600000000),
		},
		common.HexToAddress("0xB29f773D371952812A8f116441dbaCF68C3dd8c4"): {
			Delegator: common.HexToAddress("0xB29f773D371952812A8f116441dbaCF68C3dd8c4"),
			Amount:    new(big.Int).SetUint64(300000000),
		},
		common.HexToAddress("0xe4b1528d9D72Cdc3f30F7F6dF73721FDA48C2831"): {
			Delegator: common.HexToAddress("0xe4b1528d9D72Cdc3f30F7F6dF73721FDA48C2831"),
			Amount:    new(big.Int).SetUint64(600000000),
		},
	}
}

// ------
func stakeFor(t *testing.T, stakeHandler *StakeHandler, stakeDatas [][]byte, check bool) {

	for _, data := range stakeDatas {
		err := stakeHandler.onStake(data)
		if check {
			assert.Nil(t, err, "Failed to call onStake")
		}
	}

}

func addStake(t *testing.T, stakeHandler *StakeHandler, addStakeDatas [][]byte, check bool) {

	for _, data := range addStakeDatas {
		err := stakeHandler.onAddStake(data)
		if check {
			assert.Nil(t, err, "Failed to call onAddStake")
		}
	}

}

func delegate(t *testing.T, stakeHandler *StakeHandler, delegateData []byte, check bool) {

	err := stakeHandler.onDelegate(delegateData)
	if check {
		assert.Nil(t, err, "Failed to call onDelegate")
	}
}
