package reward

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	rewardcommon "github.com/PlatONnetwork/AppChain-SDK/x/reward/common"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/contracts"
	rewarddb "github.com/PlatONnetwork/AppChain-SDK/x/reward/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

type RewardModule struct {
	logger  log.Logger
	staking types.Staking
	stage   types.Stage
}

func NewRewardModule(staking types.Staking, stage types.Stage) *RewardModule {
	return &RewardModule{
		logger:  log.New("module", "reward"),
		staking: staking,
		stage:   stage,
	}
}

func (r *RewardModule) Name() string {
	return "reward"
}

func (r *RewardModule) Address() basecommon.Address {
	return constants.RewardManagerAddress
}

func (r *RewardModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	rewardManager, _ := contracts.NewRewardManager(evm, contract, readOnly)
	return rewardManager.Run(input)
}

func (r *RewardModule) BeginBlock(ctx sdk.WorkerContext) {
	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()
	// distribute blocks reward (with round)
	if r.stage.IsBeginOfCurrentRound(ctx.StateDB(), currentBlock) {
		r.handleBlocksRewardForPreviousRound(ctx.StateDB())
	}
}

func (r *RewardModule) EndBlock(ctx sdk.WorkerContext) {

	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()

	// distribute epoch reward
	if r.stage.IsEndOfCurrentEpoch(ctx.StateDB(), currentBlock) {
		if err := r.handleEpochReward(ctx.StateDB()); nil != err {
			panic(fmt.Sprintf("Failed to handle epoch reward, %s", err))
		}
	}
}

func (r *RewardModule) handleBlocksRewardForPreviousRound(stateDB sdk.StateDB) error {
	currentRound := r.stage.GetCurrentRound(stateDB)
	previousRoundValidatorIds := r.staking.GetRoundValidatorIds(stateDB, currentRound-1)

	for _, validatorAddr := range previousRoundValidatorIds {

		numberOfBlocks := r.staking.GetNumberOfBlocksForRoundValidator(stateDB, validatorAddr, currentRound)
		// got it !
		blocksReward := new(big.Int).Mul(rewardcommon.REWARD_PER_BLOCK, big.NewInt(int64(numberOfBlocks)))

		// increment validatorEpochReward to validator rewards
		rewarddb.IncrementPendingValidatorReward(stateDB, r.Address(), validatorAddr, blocksReward)

	}
	return nil
}

func (r *RewardModule) handleEpochReward(stateDB sdk.StateDB) error {
	currentEpoch := r.stage.GetCurrentEpoch(stateDB)
	epochValidatorIds := r.staking.GetEpochValidatorIds(stateDB, currentEpoch)

	perValidatorEpochReward := new(big.Int).Div(rewardcommon.REWARD_PER_EPOCH, big.NewInt(int64(len(epochValidatorIds))))

	for _, validatorAddr := range epochValidatorIds {

		stakeAmount := r.staking.GetValidatorStakeAmount(stateDB, validatorAddr)
		delegateAmount := r.staking.GetValidatorDelegateAmount(stateDB, validatorAddr)
		totalShares := new(big.Int).Add(stakeAmount, delegateAmount)

		commissionRate := r.staking.GetValidatorCommissionRate(stateDB, validatorAddr)
		// commission == perValidatorEpochReward * (commissionRate/ 100) == (perValidatorEpochReward * commissionRate)/ 100
		commission := new(big.Int).Div(new(big.Int).Mul(perValidatorEpochReward, big.NewInt(int64(commissionRate))), basecommon.Big100)
		// nonCommission == perValidatorEpochReward - commission
		nonCommission := new(big.Int).Sub(perValidatorEpochReward, commission)
		// stakeEpochReward = nonCommission * (stakeAmount/totalShares) == (nonCommission * stakeAmount) / totalShares
		stakeEpochReward := new(big.Int).Div(new(big.Int).Mul(nonCommission, stakeAmount), totalShares)

		// got it !
		validatorEpochReward := new(big.Int).Add(commission, stakeEpochReward)
		delegateEpochReward := new(big.Int).Sub(perValidatorEpochReward, validatorEpochReward)
		delegaterEpochPerShareReward := new(big.Int).Div(delegateEpochReward, delegateAmount)

		// increment validatorEpochReward to validator rewards
		rewarddb.IncrementPendingValidatorReward(stateDB, r.Address(), validatorAddr, stakeEpochReward)
		// store delegaterEpochTotalReward and delegaterEpochPerShareReward
		if err := rewarddb.SetEpochDelegationRewardPerShareItem(stateDB, r.Address(), validatorAddr, currentEpoch,
			types.NewEpochDelegationRewardPerShareItem(delegateEpochReward, delegaterEpochPerShareReward)); nil != err {
			r.logger.Error("Set epoch  delegation reward for per share", "currentEpoch", currentEpoch, "error", err)
			return err
		}

		// update owner of validator (for with validator reward)
		newOwner := r.staking.GetValidatorOwner(stateDB, validatorAddr)
		oldOwner := rewarddb.GetValidatorRewardOwner(stateDB, r.Address(), validatorAddr)
		if newOwner != oldOwner {
			rewarddb.SetValidatorRewardOwner(stateDB, r.Address(), validatorAddr, newOwner)
		}

	}
	return nil
}
