package reward

import (
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/contracts"
	rewarddb "github.com/PlatONnetwork/AppChain-SDK/x/reward/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
)

type RewardModule struct {
	logger log.Logger
	stage  types.Stage
	stake  types.Stake
}

func NewRewardModule(ctx *cli.Context, stage types.Stage) *RewardModule {
	return &RewardModule{
		logger: log.New("module", "reward"),
		stage:  stage,
	}
}

func (r *RewardModule) SetStakeModule(stake types.Stake) {
	r.stake = stake
}

func (r *RewardModule) Name() string {
	return "reward"
}

func (r *RewardModule) Address() basecommon.Address {
	return constants.RewardManagerAddress
}

func (r *RewardModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	rewardManager, _ := contracts.NewRewardManager(evm, contract, readOnly)
	rewardManager.SetStageModule(r.stage)
	rewardManager.SetStakeModule(r.stake)
	rewardManager.SetRewardModule(r)
	return rewardManager.Run(input)
}

func (r *RewardModule) BeginBlock(ctx sdk.WorkerContext) {
	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()
	// distribute blocks reward (with round)
	if r.stage.IsBeginOfCurrentRound(ctx.StateDB(), currentBlock) {
		if err := r.handleBlocksRewardForPreviousRound(ctx.StateDB(), currentBlock); nil != err {
			panic(fmt.Sprintf("Failed to handle blocks reward for previous round, currentRound: %d, blockNumber: %d, error: %s", r.stage.GetCurrentRound(ctx.StateDB()), currentBlock, err))
		}
	}
}

func (r *RewardModule) EndBlock(ctx sdk.WorkerContext) {

	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()

	// distribute epoch reward
	if r.stage.IsEndOfCurrentEpoch(ctx.StateDB(), currentBlock) {
		if err := r.handleEpochReward(ctx.StateDB(), currentBlock); nil != err {
			panic(fmt.Sprintf("Failed to handle epoch reward, currentEpoch: %d, blockNumber: %d, error: %s", r.stage.GetCurrentEpoch(ctx.StateDB()), currentBlock, err))
		}
	}
}

func (r *RewardModule) handleBlocksRewardForPreviousRound(stateDB sdk.StateDB, blockNumber uint64) error {

	if r.stage.IsNotBeginOfCurrentRound(stateDB, blockNumber) {
		return errors.New("block is not endBlock of current epoch")
	}

	currentRound := r.stage.GetCurrentRound(stateDB)
	handleRound := currentRound - 1
	previousRoundValidatorIds := r.stake.GetRoundValidatorIds(stateDB, handleRound)

	totalPaidReward := basecommon.Big0
	for _, validatorAddr := range previousRoundValidatorIds {

		numberOfBlocks := r.stake.GetNumberOfBlocksForRoundValidator(stateDB, validatorAddr, handleRound)
		// got it !
		blocksReward := new(big.Int).Mul(constants.REWARD_PER_BLOCK, big.NewInt(int64(numberOfBlocks)))

		// increment validatorEpochReward to validator rewards
		rewarddb.IncrementPendingValidatorReward(stateDB, r.Address(), validatorAddr, blocksReward)

		totalPaidReward = new(big.Int).Add(totalPaidReward, blocksReward)

		r.logger.Debug("Finished distribute blocks reward", "currentRound", currentRound, "handle round", handleRound, "validatorAddr", validatorAddr.Hex(),
			"blocksReward", blocksReward, "numberOfBlocks", numberOfBlocks, "blockNumber", blockNumber)
	}

	rewarddb.IncrementPaidRewardPerEpoch(stateDB, r.Address(), r.stage.GetCurrentEpoch(stateDB), totalPaidReward)
	return nil
}

func (r *RewardModule) handleEpochReward(stateDB sdk.StateDB, blockNumber uint64) error {

	if r.stage.IsNotEndOfCurrentEpoch(stateDB, blockNumber) {
		return errors.New("block is not endBlock of current epoch")
	}

	currentEpoch := r.stage.GetCurrentEpoch(stateDB)
	epochValidatorIds := r.stake.GetEpochValidatorIds(stateDB, currentEpoch)

	perValidatorEpochReward := new(big.Int).Div(constants.REWARD_PER_EPOCH, big.NewInt(int64(len(epochValidatorIds))))

	for _, validatorAddr := range epochValidatorIds {

		if r.stake.IsInvalidValidator(stateDB, validatorAddr) {
			continue
		}

		stakeAmount := r.stake.GetValidatorStakeAmount(stateDB, validatorAddr)
		delegateAmount := r.stake.GetValidatorDelegateAmount(stateDB, validatorAddr)
		commissionRate := r.stake.GetValidatorCommissionRate(stateDB, validatorAddr)

		realDelegateEpochReward := basecommon.Big0
		realValidatorEpochReward := basecommon.Big0
		perShareDelegaterEpochReward := basecommon.Big0

		totalShares := new(big.Int).Add(stakeAmount, delegateAmount)
		// commissionAmount == perValidatorEpochReward * (commissionRate/ 100) == (perValidatorEpochReward * commissionRate)/ 100
		commissionAmount := new(big.Int).Div(new(big.Int).Mul(perValidatorEpochReward, big.NewInt(int64(commissionRate))), basecommon.Big100)
		// nonCommissionAmount == perValidatorEpochReward - commissionAmount
		nonCommissionAmount := new(big.Int).Sub(perValidatorEpochReward, commissionAmount)
		// stakeEpochReward = nonCommissionAmount * (stakeAmount/totalShares) == (nonCommissionAmount * stakeAmount) / totalShares
		stakeEpochReward := new(big.Int).Div(new(big.Int).Mul(nonCommissionAmount, stakeAmount), totalShares)

		// got it !
		validatorEpochReward := new(big.Int).Add(commissionAmount, stakeEpochReward)
		delegateEpochReward := new(big.Int).Sub(perValidatorEpochReward, validatorEpochReward)

		if delegateAmount.Cmp(basecommon.Big0) != 0 { // has delegate amount
			perShareDelegaterEpochReward = new(big.Int).Div(delegateEpochReward, delegateAmount)
			// ###### NOTE ######
			// Rolling calculation eliminates the situation where the total `delegateEpochReward` is not evenly divided,
			// preventing `delegateEpochReward` from not being reduced to zero.
			realDelegateEpochReward = new(big.Int).Mul(perShareDelegaterEpochReward, delegateAmount)

			// store delegaterEpochTotalReward and delegaterEpochPerShareReward
			if err := rewarddb.AppendEpochDelegationRewardPerShareItem(stateDB, r.Address(), validatorAddr,
				r.stake.GetValidatorStakeEpoch(stateDB, validatorAddr), currentEpoch,
				realDelegateEpochReward, perShareDelegaterEpochReward); nil != err {

				r.logger.Error("Set epoch  delegation reward for per share", "currentEpoch", currentEpoch, "error", err)
				return err
			}
		}
		realValidatorEpochReward = new(big.Int).Sub(perValidatorEpochReward, realDelegateEpochReward)

		// increment validatorEpochReward to validator rewards
		rewarddb.IncrementPendingValidatorReward(stateDB, r.Address(), validatorAddr, realValidatorEpochReward)

		r.logger.Debug("Finished distribute epoch reward", "currentEpoch", currentEpoch, "validatorAddr", validatorAddr.Hex(), "validatorEpochReward", realValidatorEpochReward,
			"delegateEpochReward", realDelegateEpochReward, "perShareDelegaterEpochReward", perShareDelegaterEpochReward, "blockNumber", blockNumber)

		// update owner of validator (for with validator reward)
		newOwner := r.stake.GetValidatorOwner(stateDB, validatorAddr)
		oldOwner := rewarddb.GetValidatorRewardOwner(stateDB, r.Address(), validatorAddr)
		if newOwner != basecommon.ZeroAddr && newOwner != oldOwner {
			rewarddb.SetValidatorRewardOwner(stateDB, r.Address(), validatorAddr, newOwner)
		}
	}

	rewarddb.IncrementPaidRewardPerEpoch(stateDB, r.Address(), currentEpoch, constants.REWARD_PER_EPOCH)

	return nil
}

func (r *RewardModule) getDelegateSnapshot(stateDB sdk.StateDBReader, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) *types.DelegationSnapshot {
	delegateEpoch, delegateAmount := r.stake.GetDelegationFlatten(stateDB, delegaterAddr, validatorAddr, stakeEpoch)
	if delegateEpoch == 0 && delegateAmount == basecommon.Big0 {
		return nil
	}
	return types.NewDelegationSnapshot(stakeEpoch, delegateEpoch, delegateAmount)
}

func (r *RewardModule) getEpochDelegationRewardQueue(stateDB sdk.StateDBReader, validatorAddr basecommon.Address, stakeEpoch, fromRewardEpoch, toRewardEpoch uint64) types.EpochDelegationRewardPerShareWithEpochQueue {

	queue := types.NewEpochDelegationRewardPerShareWithEpochQueue(0)

	for rewardEpoch := fromRewardEpoch; rewardEpoch <= toRewardEpoch; rewardEpoch++ {
		item := rewarddb.GetEpochDelegationRewardPerShareItem(stateDB, r.Address(), validatorAddr, stakeEpoch, rewardEpoch)
		if item.IsEmpty() || item.IsZeroTotalReward() {
			continue
		}
		queue = append(queue, types.NewEpochDelegationRewardPerShareWithEpochItem(rewardEpoch, item))
	}

	return queue
}

func (r *RewardModule) aggregationEpochDelegationRewards(stateDB sdk.StateDB, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64, delegateAmount *big.Int, rewardQueue types.EpochDelegationRewardPerShareWithEpochQueue) error {

	totalRewards := basecommon.Big0

	for _, item := range rewardQueue {

		shareReward := new(big.Int).Mul(item.Data.PerShareReward, delegateAmount)

		// update epoch delegation reward item
		item.Data.TotalReward = new(big.Int).Sub(item.Data.TotalReward, shareReward)
		if err := rewarddb.SetEpochDelegationRewardPerShareItem(stateDB, r.Address(), validatorAddr, stakeEpoch, item.RewardEpoch, item.Data); nil != err {
			return err
		}
		totalRewards = new(big.Int).Add(totalRewards, shareReward)
	}

	// increment delegater rewards
	if totalRewards.Cmp(basecommon.Big0) != 0 {
		rewarddb.IncrementPendingDelegaterReward(stateDB, r.Address(), delegaterAddr, validatorAddr, totalRewards)
	}
	return nil
}

// extern

func (r *RewardModule) UpdateDelegationRewards(stateDB sdk.StateDB, delegaterAddr, validatorAddr basecommon.Address) error {

	currentEpoch := r.stage.GetCurrentEpoch(stateDB)
	if currentEpoch == 1 {
		r.logger.Warn("No epoch reward has been assigned yet", "currentEpoch", currentEpoch)
		return nil
	}

	stakeEpochQueue := r.stake.GetEpochByValidatorDelegationRcPending(stateDB, validatorAddr)

	previousEpoch := currentEpoch - 1
	delegateRewardSnapshotQueue := types.NewDelegationRewardSnapshotQueue(0)
	// get epoch delegate rewards
	for _, stakeEpoch := range stakeEpochQueue {
		// get delegation
		delegation := r.getDelegateSnapshot(stateDB, delegaterAddr, validatorAddr, stakeEpoch)
		rewardQueue := r.getEpochDelegationRewardQueue(stateDB, validatorAddr, stakeEpoch, delegation.DelegateEpoch, previousEpoch)
		if rewardQueue.IsEmpty() {
			continue
		}
		delegateRewardSnapshotQueue = append(delegateRewardSnapshotQueue, types.NewDelegationRewardSnapshot(delegation, rewardQueue))
	}

	if len(delegateRewardSnapshotQueue) == 0 {
		return nil
	}

	for _, snap := range delegateRewardSnapshotQueue {
		if err := r.aggregationEpochDelegationRewards(stateDB, delegaterAddr, validatorAddr, snap.Delegation.StakeEpoch, snap.Delegation.Amount, snap.RewardQueue); nil != err {
			r.logger.Error("Failed to aggregate epoch delegation rewards", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "stakeEpoch", snap.Delegation.StakeEpoch, "error", err)
			return err
		}
		// update delegateEpoch of delegation to currentEpoch
		if err := r.stake.UpdateDelegationEpoch(stateDB, delegaterAddr, validatorAddr, snap.Delegation.StakeEpoch, currentEpoch); nil != err {
			r.logger.Error("Failed to update delegation epoch", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "stakeEpoch", snap.Delegation.StakeEpoch, "error", err)
			return err
		}
	}

	return nil
}
func (r *RewardModule) UpdateDelegationRewardsByStakeEpoch(stateDB sdk.StateDB, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) error {

	currentEpoch := r.stage.GetCurrentEpoch(stateDB)
	if currentEpoch == 1 {
		r.logger.Warn("No epoch reward has been assigned yet", "currentEpoch", currentEpoch)
		return nil
	}

	previousEpoch := currentEpoch - 1
	delegateRewardSnapshotQueue := types.NewDelegationRewardSnapshotQueue(0)

	// get delegation
	delegation := r.getDelegateSnapshot(stateDB, delegaterAddr, validatorAddr, stakeEpoch)
	rewardQueue := r.getEpochDelegationRewardQueue(stateDB, validatorAddr, stakeEpoch, delegation.DelegateEpoch, previousEpoch)
	if rewardQueue.IsEmpty() {
		return nil
	}
	delegateRewardSnapshotQueue = append(delegateRewardSnapshotQueue, types.NewDelegationRewardSnapshot(delegation, rewardQueue))

	if len(delegateRewardSnapshotQueue) == 0 {
		return nil
	}

	for _, snap := range delegateRewardSnapshotQueue {
		if err := r.aggregationEpochDelegationRewards(stateDB, delegaterAddr, validatorAddr, snap.Delegation.StakeEpoch, snap.Delegation.Amount, snap.RewardQueue); nil != err {
			r.logger.Error("Failed to aggregate epoch delegation rewards", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "stakeEpoch", snap.Delegation.StakeEpoch, "error", err)
			return err
		}
		// update delegateEpoch of delegation to currentEpoch
		if err := r.stake.UpdateDelegationEpoch(stateDB, delegaterAddr, validatorAddr, snap.Delegation.StakeEpoch, currentEpoch); nil != err {
			r.logger.Error("Failed to update delegation epoch", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "stakeEpoch", snap.Delegation.StakeEpoch, "error", err)
			return err
		}
	}

	return nil
}
