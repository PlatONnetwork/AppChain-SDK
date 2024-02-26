package contracts

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/mocks"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/test-go/testify/assert"
	"math/big"
	"testing"
)

func Test_RemoveInvalidValidatorPriorityWhenLowBlocks(t *testing.T) {
	stakeHandlerTestConfig := newStakeHandlerTestConfig()
	numberOfBlocks := uint64(2)
	epoch := uint64(1) // 24 blocks is a epoch
	round := uint64(1) // 8 blocks is a round
	from := common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10")
	statedb := (stakeHandlerTestConfig.StateDB).(vm.StateDB)

	// init stakeHandler
	stakeHandler := newStakeHandler(statedb, from, new(big.Int).SetUint64(2))
	initStakeHandler(stakeHandler, stakeHandlerTestConfig.L1Module, stakeHandlerTestConfig.StageModule, stakeHandlerTestConfig.StakeModule, stakeHandlerTestConfig.RewardModule)

	// mock init genesis
	stakeHandlerTestConfig.StageModule.MockCurrentRound(round)
	stakeHandlerTestConfig.StageModule.MockCurrentEpoch(epoch)
	stakeHandlerTestConfig.StakeModule.MockInitValidatorGenesisPriority()
	// get mock data
	datas, _ := getStakeForData()
	stakeFor(t, stakeHandler, epoch, datas)

	_, queue, _ := stakeHandler.GetValidators([]byte{}, common.Big100)

	snapshotQueue := mocks.NewMockValidatorSnapshotQueue(uint64(len(queue)))

	var lowBlocksValidatorAddr common.Address
	for i, validator := range queue {
		if i != 0 {
			// mock numberOfBlocks for round validators
			setNumberOfBlocksForRoundValidator(statedb, validator.ValidatorAddr, round, numberOfBlocks)
		} else {
			lowBlocksValidatorAddr = validator.ValidatorAddr
		}
		snapshotQueue[i] = &mocks.MockValidatorSnapshot{
			ValidatorAddr:  validator.ValidatorAddr,
			Epoch:          validator.Epoch.Uint64(),
			StakeIndex:     validator.StakeIndex.Uint64(),
			CommissionRate: validator.CommissionRate.Uint64(),
			StakeAmount:    validator.StakeAmount,
			DelegateAmount: validator.DelegateAmount,
		}
	}

	stakeHandlerTestConfig.StakeModule.MockRoundValidatorSnapQueue(round, snapshotQueue)

	// mock check lowBlocks
	err := checkLowBlocksValidatorForPreviousRound(statedb, 2, 1)
	assert.Nil(t, err, "Failed to call checkLowBlocksValidatorForPreviousRound")

	for _, validator := range queue {
		v := stakeHandler.getValidator(validator.ValidatorAddr)
		blocks := getNumberOfBlocksForRoundValidator(statedb, validator.ValidatorAddr, round)
		if lowBlocksValidatorAddr == validator.ValidatorAddr {
			assert.True(t, v.IsInvalidLowBlocks(), "had not lowBlocks status")
			assert.Equal(t, uint64(0), blocks, "must be zero blocks")
		} else {
			assert.False(t, v.IsInvalidLowBlocks(), "must be not lowBlocks status")
			assert.Equal(t, numberOfBlocks, blocks, "must be not zero blocks")
		}
	}
	validatorIds := rankPriorityValidatorIds(statedb, 201)
	assert.Equal(t, 3, len(validatorIds), "invalid validators count")

	validatorAddrCache := make(map[common.Address]struct{}, 0)
	for _, validatorAddr := range validatorIds {
		validatorAddrCache[validatorAddr] = struct{}{}
	}
	_, ok := validatorAddrCache[lowBlocksValidatorAddr]
	assert.False(t, ok, "must be not found")
}

func Test_RemoveInvalidValidatorPriorityWhenSlash(t *testing.T) {
	stakeHandlerTestConfig := newStakeHandlerTestConfig()
	numberOfBlocks := uint64(2)
	epoch := uint64(1) // 24 blocks is a epoch
	round := uint64(1) // 8 blocks is a round
	from := common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10")
	statedb := (stakeHandlerTestConfig.StateDB).(vm.StateDB)

	// init stakeHandler
	stakeHandler := newStakeHandler(statedb, from, new(big.Int).SetUint64(2))
	initStakeHandler(stakeHandler, stakeHandlerTestConfig.L1Module, stakeHandlerTestConfig.StageModule, stakeHandlerTestConfig.StakeModule, stakeHandlerTestConfig.RewardModule)

	// mock init genesis
	stakeHandlerTestConfig.L1Module.MockStakeManagerAddress(common.HexToAddress("0x10506aB975D36aa781B77C1Ce204F46e8f87dA57"))
	stakeHandlerTestConfig.StageModule.MockCurrentRound(round)
	stakeHandlerTestConfig.StageModule.MockCurrentEpoch(epoch)
	stakeHandlerTestConfig.StakeModule.MockInitValidatorGenesisPriority()
	// get mock data
	datas, _ := getStakeForData()
	stakeFor(t, stakeHandler, epoch, datas)

	_, queue, _ := stakeHandler.GetValidators([]byte{}, common.Big100)

	snapshotQueue := mocks.NewMockValidatorSnapshotQueue(uint64(len(queue)))

	var lowBlocksValidatorAddr common.Address
	for i, validator := range queue {
		if i != 0 {
			// mock numberOfBlocks for round validators
			setNumberOfBlocksForRoundValidator(statedb, validator.ValidatorAddr, round, numberOfBlocks)
		} else {
			lowBlocksValidatorAddr = validator.ValidatorAddr
		}
		snapshotQueue[i] = &mocks.MockValidatorSnapshot{
			ValidatorAddr:  validator.ValidatorAddr,
			Epoch:          validator.Epoch.Uint64(),
			StakeIndex:     validator.StakeIndex.Uint64(),
			CommissionRate: validator.CommissionRate.Uint64(),
			StakeAmount:    validator.StakeAmount,
			DelegateAmount: validator.DelegateAmount,
		}
	}

	stakeHandlerTestConfig.StageModule.MockCurrentRound(2)
	stakeHandlerTestConfig.StakeModule.MockRoundValidatorSnapQueue(round, snapshotQueue)
	stakeHandlerTestConfig.StakeModule.MockEpochValidatorSnapQueue(round, snapshotQueue)
	stakeHandlerTestConfig.StakeModule.MockMinBlocksOfRoundValidator(1)

	// mock check lowBlocks
	checkLowBlocksValidatorForPreviousRound(statedb, 2, 1)

	// mock slash
	stakeHandler.Slash()

	//
	_, queue, _ = stakeHandler.GetValidators([]byte{}, common.Big100)
	for _, validator := range queue {
		v := stakeHandler.getValidator(validator.ValidatorAddr)
		blocks := getNumberOfBlocksForRoundValidator(statedb, validator.ValidatorAddr, round)
		if lowBlocksValidatorAddr == validator.ValidatorAddr {
			assert.True(t, v.IsInvalidLowBlocks() && v.IsInvalidSlashing(), "had not lowBlocks and slashing status")
			assert.Equal(t, uint64(0), blocks, "must be zero blocks")
		} else {
			assert.False(t, v.IsInvalidLowBlocks() && v.IsInvalidSlashing(), "must be not lowBlocks and slashing status")
			assert.Equal(t, numberOfBlocks, blocks, "must be not zero blocks")
		}
	}
	validatorIds := rankPriorityValidatorIds(statedb, 201)
	assert.Equal(t, 3, len(validatorIds), "invalid validators count")

	validatorAddrCache := make(map[common.Address]struct{}, 0)
	for _, validatorAddr := range validatorIds {
		validatorAddrCache[validatorAddr] = struct{}{}
	}
	_, ok := validatorAddrCache[lowBlocksValidatorAddr]
	assert.False(t, ok, "must be not found")
}
