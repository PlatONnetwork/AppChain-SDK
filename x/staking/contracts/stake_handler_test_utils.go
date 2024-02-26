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
	"math/big"
)

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
