package reward

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/contracts"
	rewarddb "github.com/PlatONnetwork/AppChain-SDK/x/reward/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/reward/types"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	MODULE_NAME_REWARD = "reward"
)

type RewardModule struct {
	logger         log.Logger
	nodePrivateKey *ecdsa.PrivateKey
	stageModule    types.StageModuler
	stakeModule    types.StakeModuler
}

func NewRewardModule(ctx *cli.Context, stage types.StageModuler) *RewardModule {
	return &RewardModule{
		logger:      log.New("module", MODULE_NAME_REWARD),
		stageModule: stage,
	}
}

func (r *RewardModule) SetStakeModule(stake types.StakeModuler) {
	r.stakeModule = stake
}

func (r *RewardModule) Name() string {
	return MODULE_NAME_REWARD
}

func (r *RewardModule) Init(ctx sdk.InitContext) error {
	r.nodePrivateKey = ctx.NodeKey()
	return nil
}

func (r *RewardModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {

	configParams := config.DefaultRewardNetworkParams()
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON RewardNetworkParams bytes", "error", err)
		return err
	}

	var conf config.RewardNetworkParams
	if err := json.Unmarshal(raw, &conf); nil != err {
		log.Error("Failed UnmarshalJSON RewardNetworkParams", "error", err)
		return err
	} else {
		configParams = &conf
	}
	// init reward manager account nonce
	initAccountNonce(db, r.Address())
	// set config params
	initConfigParams(db, r.Address(), configParams)

	log.Info("Succeed init genesis", "module", r.Name(), "RewardNetworkParams", configParams.String())
	return nil
}

func (r *RewardModule) Address() basecommon.Address {
	return constants.RewardManagerAddress
}

func (r *RewardModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	rewardManager, _ := contracts.NewRewardManager(evm, contract, readOnly)
	rewardManager.SetStageModule(r.stageModule)
	rewardManager.SetStakeModule(r.stakeModule)
	rewardManager.SetRewardModule(r)
	return rewardManager.Run(input)
}

func (r *RewardModule) BeginBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}
	// distribute blocks reward (with round)
	if r.stageModule.IsBeginOfCurrentRound(ctx.StateDB(), currentBlock) {
		if err := r.handleBlocksRewardForPreviousRound(ctx.StateDB(), currentBlock); nil != err {
			return fmt.Errorf("can not handle blocks reward for previous round, %s, currentRound: %d", err, r.stageModule.GetCurrentRound(ctx.StateDB()))
		}
	}
	return nil
}

func (r *RewardModule) EndBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}

	// distribute epoch reward
	if r.stageModule.IsEndOfCurrentEpoch(ctx.StateDB(), currentBlock) {
		if err := r.handleEpochReward(ctx.StateDB(), currentBlock); nil != err {
			return fmt.Errorf("can not handle epoch reward, %s, currentEpoch: %d", err, r.stageModule.GetCurrentEpoch(ctx.StateDB()))
		}
	}

	return nil
}

func (r *RewardModule) handleBlocksRewardForPreviousRound(stateDB sdk.StateDB, blockNumber uint64) error {

	if r.stageModule.IsNotBeginOfCurrentRound(stateDB, blockNumber) {
		return errors.New("block is not endBlock of current epoch")
	}

	currentRound := r.stageModule.GetCurrentRound(stateDB)
	if currentRound == 0 {
		return nil
	}

	handleRound := currentRound - 1
	previousRoundValidatorIds := r.stakeModule.GetRoundValidatorIds(stateDB, handleRound)

	totalPaidReward := basecommon.Big0
	for _, validatorAddr := range previousRoundValidatorIds {

		numberOfBlocks := r.stakeModule.GetNumberOfBlocksForRoundValidator(stateDB, validatorAddr, handleRound)
		// got it !
		blocksReward := new(big.Int).Mul(r.GetRewardPerBlock(stateDB), big.NewInt(int64(numberOfBlocks)))

		// increment validatorEpochReward to validator rewards
		rewarddb.IncrementPendingValidatorReward(stateDB, r.Address(), validatorAddr, blocksReward)

		totalPaidReward = new(big.Int).Add(totalPaidReward, blocksReward)

		r.logger.Debug("Finished distribute blocks reward", "currentRound", currentRound, "handleRound", handleRound, "validatorAddr", validatorAddr.Hex(),
			"blocksReward", blocksReward, "numberOfBlocks", numberOfBlocks, "blockNumber", blockNumber)
	}

	rewarddb.IncrementPaidRewardPerEpoch(stateDB, r.Address(), r.stageModule.GetCurrentEpoch(stateDB), totalPaidReward)
	return nil
}

func (r *RewardModule) handleEpochReward(stateDB sdk.StateDB, blockNumber uint64) error {

	if r.stageModule.IsNotEndOfCurrentEpoch(stateDB, blockNumber) {
		return errors.New("block is not endBlock of current epoch")
	}

	currentEpoch := r.stageModule.GetCurrentEpoch(stateDB)
	epochValidatorIds := r.stakeModule.GetEpochValidatorIds(stateDB, currentEpoch)

	perValidatorEpochReward := new(big.Int).Div(r.GetRewardPerEpoch(stateDB), big.NewInt(int64(len(epochValidatorIds))))

	for _, validatorAddr := range epochValidatorIds {

		if r.stakeModule.IsInvalidValidator(stateDB, validatorAddr) {
			continue
		}

		stakeAmount := r.stakeModule.GetValidatorStakeAmount(stateDB, validatorAddr)
		delegateAmount := r.stakeModule.GetValidatorDelegateAmount(stateDB, validatorAddr)
		commissionRate := r.stakeModule.GetValidatorCommissionRate(stateDB, validatorAddr)

		realDelegateEpochReward := basecommon.Big0
		realValidatorEpochReward := basecommon.Big0
		perShareDelegatorEpochReward := basecommon.Big0

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
			perShareDelegatorEpochReward = new(big.Int).Div(delegateEpochReward, delegateAmount)
			// ###### NOTE ######
			// Rolling calculation eliminates the situation where the total `delegateEpochReward` is not evenly divided,
			// preventing `delegateEpochReward` from not being reduced to zero.
			realDelegateEpochReward = new(big.Int).Mul(perShareDelegatorEpochReward, delegateAmount)

			// store delegatorEpochTotalReward and delegaterEpochPerShareReward
			if err := rewarddb.AppendEpochDelegationRewardPerShareItem(stateDB, r.Address(), validatorAddr,
				r.stakeModule.GetValidatorStakeEpoch(stateDB, validatorAddr), currentEpoch,
				realDelegateEpochReward, perShareDelegatorEpochReward); nil != err {

				r.logger.Error("Set epoch  delegation reward for per share", "currentEpoch", currentEpoch, "error", err)
				return err
			}
		}
		realValidatorEpochReward = new(big.Int).Sub(perValidatorEpochReward, realDelegateEpochReward)

		// increment validatorEpochReward to validator rewards
		rewarddb.IncrementPendingValidatorReward(stateDB, r.Address(), validatorAddr, realValidatorEpochReward)

		r.logger.Debug("Finished distribute epoch reward", "currentEpoch", currentEpoch, "validatorAddr", validatorAddr.Hex(), "validatorEpochReward", realValidatorEpochReward,
			"delegateEpochReward", realDelegateEpochReward, "perShareDelegatorEpochReward", perShareDelegatorEpochReward, "blockNumber", blockNumber)
	}

	rewarddb.IncrementPaidRewardPerEpoch(stateDB, r.Address(), currentEpoch, r.GetRewardPerEpoch(stateDB))

	return nil
}

func (r *RewardModule) getDelegateSnapshot(stateDB sdk.StateDBReader, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) *types.DelegationSnapshot {
	delegateEpoch, delegateAmount := r.stakeModule.GetDelegationFlatten(stateDB, delegaterAddr, validatorAddr, stakeEpoch)
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
		rewarddb.IncrementPendingDelegatorReward(stateDB, r.Address(), delegaterAddr, validatorAddr, totalRewards)
	}
	return nil
}

// extern

func (r *RewardModule) UpdateDelegationRewards(stateDB sdk.StateDB, delegaterAddr, validatorAddr basecommon.Address) error {

	currentEpoch := r.stageModule.GetCurrentEpoch(stateDB)
	if currentEpoch == 1 {
		r.logger.Warn("No epoch reward has been assigned yet", "currentEpoch", currentEpoch)
		return nil
	}

	stakeEpochQueue := r.stakeModule.GetEpochByValidatorDelegationRcPending(stateDB, validatorAddr)

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
		if err := r.stakeModule.UpdateDelegationEpoch(stateDB, delegaterAddr, validatorAddr, snap.Delegation.StakeEpoch, currentEpoch); nil != err {
			r.logger.Error("Failed to update delegation epoch", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "stakeEpoch", snap.Delegation.StakeEpoch, "error", err)
			return err
		}
	}

	return nil
}
func (r *RewardModule) UpdateDelegationRewardsByStakeEpoch(stateDB sdk.StateDB, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) error {

	currentEpoch := r.stageModule.GetCurrentEpoch(stateDB)
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
		if err := r.stakeModule.UpdateDelegationEpoch(stateDB, delegaterAddr, validatorAddr, snap.Delegation.StakeEpoch, currentEpoch); nil != err {
			r.logger.Error("Failed to update delegation epoch", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "stakeEpoch", snap.Delegation.StakeEpoch, "error", err)
			return err
		}
	}

	return nil
}

func (r *RewardModule) GetRewardPerBlock(stateDB sdk.StateDBReader) *big.Int {
	return rewarddb.GetRewardPerBlock(stateDB, r.Address())
}
func (r *RewardModule) GetRewardPerEpoch(stateDB sdk.StateDBReader) *big.Int {
	return rewarddb.GetRewardPerEpoch(stateDB, r.Address())
}
