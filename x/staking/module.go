package staking

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	stakingp2p "github.com/PlatONnetwork/AppChain-SDK/x/staking/p2p"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	stakewrap "github.com/PlatONnetwork/AppChain-SDK/x/staking/wrap"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName    = "staking"
	ModuleVersion = 1
)

var _ module.ContractModule = (*StakeModule)(nil)

type StakeModule struct {
	p2p            *stakingp2p.StakingP2P
	logger         log.Logger
	nodePrivateKey *ecdsa.PrivateKey
	l1Module       staketypes.L1Moduler
	stageModule    staketypes.StageModuler
	vrfModule      staketypes.VRFModuler
	rewardModule   staketypes.RewardModuler
}

func NewModule(ctx *cli.Context, l1Module staketypes.L1Moduler, stage staketypes.StageModuler) *StakeModule {
	return &StakeModule{
		p2p:         stakingp2p.NewStakingP2P(),
		logger:      log.New("module", ModuleName),
		l1Module:    l1Module,
		stageModule: stage,
	}
}

func (s *StakeModule) SetRewardModule(reward staketypes.RewardModuler) {
	s.rewardModule = reward
}

func (s *StakeModule) SetVRFModule(vrf staketypes.VRFModuler) {
	s.vrfModule = vrf
}

func (s *StakeModule) Name() string {
	return ModuleName
}

func (s *StakeModule) Version() uint64 {
	return ModuleVersion
}

func (s *StakeModule) Init(ctx sdk.InitContext) error {
	s.nodePrivateKey = ctx.NodeKey()
	return nil
}

func (s *StakeModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	configParams := config.DefualtStakeNetworkParams()
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON StakeNetworkParams bytes", "error", err)
		return err
	}

	var conf config.StakeNetworkParams
	if err := json.Unmarshal(raw, &conf); nil != err {
		log.Error("Failed UnmarshalJSON StakeNetworkParams", "error", err)
		return err
	} else {
		configParams = &conf
	}

	// init staking handler account nonce
	initAccountNonce(db, s.Address())
	// store configParms
	initStakeConfigParams(db, s.Address(), configParams)

	if err := initValidators(db, s.Address(), chainConfig, configParams); nil != err {
		log.Error("Failed initialize genesis validators", "error", err)
		return err
	}

	// TODO: set create block

	log.Info("Succeed init genesis", "module", s.Name(), "StakeNetworkParams", configParams.String())
	return nil
}

func (s *StakeModule) Protocols() []p2p.Protocol {
	return s.p2p.Protocols()
}

func (s *StakeModule) Address() basecommon.Address {
	return constants.StakeHandlerAddress
}

func (s *StakeModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stakeHandler, _ := contracts.NewStakeHandler(evm, contract, readOnly)
	stakeHandler.SetL1Module(s.l1Module)
	stakeHandler.SetStageModule(s.stageModule)
	stakeHandler.SetStakeModule(s)
	stakeHandler.SetRewardModule(s.rewardModule)
	return stakeHandler.Run(input)
}

func (s *StakeModule) ContractCreateBlockNumber(statedb sdk.StateDB) uint64 {
	// TODO: implement me
	return 0
}

func (s *StakeModule) AddTxs(ctx sdk.WorkerContext, local map[basecommon.Address]types.Transactions) (map[basecommon.Address]types.Transactions, error) {

	blockNumber := ctx.Header().Number.Uint64()
	if blockNumber == 0 {
		return local, nil
	}

	if s.stageModule.IsNotBeginOfCurrentRound(ctx.StateDB(), blockNumber) {
		return local, nil
	}

	currentRound := s.stageModule.GetCurrentRound(ctx.StateDB())
	minBlocksOfRoundValidator := s.GetMinBlocksOfRoundValidator(ctx.StateDB())

	// check low blocks validtors of round, and send slash tx
	if db.HasNotLowBlocksValidator(ctx.StateDB(), s.Address(), currentRound, minBlocksOfRoundValidator) {
		return local, nil
	}

	from := crypto.PubkeyToAddress(s.nodePrivateKey.PublicKey)
	var (
		err error
	)
	// find previous nonce from local tx queue
	txNonce := common.EnableNonce(local[from], func() uint64 {
		return ctx.StateDB().GetNonce(from)
	})

	slashTx, err := s.createSlashTx(ctx, txNonce)
	if nil != err {
		s.logger.Error("Failed to create slash tx", "blockNumber", blockNumber, "error", err)
		return local, err
	}
	if nil == local[from] {
		local[from] = make(types.Transactions, 0)
	}
	local[from] = append(local[from], slashTx)
	s.logger.Debug("create Slash tx", "blockNumber", blockNumber, "txHash", slashTx.Hash().Hex(), "from", from.Hex(), "currentRound", currentRound, "minBlocksOfRoundValidator", minBlocksOfRoundValidator)
	return local, nil
}

func (s *StakeModule) BeginBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}

	// increase the number of validator blocks generated from the previous block
	parentBlock := currentBlock - 1
	if parentBlock != 0 {
		parentHeader := ctx.ParentBlock().Header()
		// @TODO for debug ...
		s.logger.Debug("Start call setNumberOfBlocksForRoundValidator", "currentBlock", currentBlock, "parentBlock", parentHeader.Number.Uint64())
		if err := s.setNumberOfBlocksForRoundValidator(ctx.StateDB(), parentHeader); nil != err {
			return fmt.Errorf("can not set number of blocks for round validators, %s, parentBlock: %d", err, parentBlock)
		}
	}

	if s.stageModule.IsBeginOfCurrentRound(ctx.StateDB(), currentBlock) {
		currentRound := s.stageModule.GetCurrentRound(ctx.StateDB())
		minBlocksOfRoundValidator := s.GetMinBlocksOfRoundValidator(ctx.StateDB())
		// check low blocks validators
		lowBlocksValidatorAddrQueue := db.CheckLowBlocksValidatorForPreviousRound(ctx.StateDB(), s.Address(), currentRound, minBlocksOfRoundValidator)
		// update validator status
		for _, validatorAddr := range lowBlocksValidatorAddrQueue {
			if err := s.updateValidatorStatus(ctx.StateDB(), validatorAddr, staketypes.Invalided|staketypes.LowBlocks); nil != err {
				return fmt.Errorf("can not update validator status to [lowBlocks], %s, validator: %s", err, validatorAddr.Hex())
			}
		}
	}
	return nil
}
func (s *StakeModule) EndBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}

	// election next round validators (at cuurent round electionBlock)
	if s.stageModule.IsElectionBlockOnCurrentRound(ctx.StateDB(), currentBlock) {
		if err := s.electionRoundValidators(ctx, currentBlock); nil != err {
			return fmt.Errorf("can not elected round validators, %s", err)
		}
	}

	// election next epoch validators (at current epoch endBlock)
	// and store next epochItem
	if s.stageModule.IsEndOfCurrentEpoch(ctx.StateDB(), currentBlock) {
		if err := s.electionEpochValidators(ctx, currentBlock); nil != err {
			return fmt.Errorf("can not elected epoch validators, %s", err)
		}
	}
	return nil
}

func (s *StakeModule) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {

	// ###### NOTE ######
	// only pre connect (p2p) the list of different validators for the next and current rounds
	// when the height of the block is selected by the validators of the round.
	if s.stageModule.IsNotElectionBlockOnCurrentRound(ctx.StateDB(), block.NumberU64()) {
		return nil
	}

	currentRound := s.stageModule.GetCurrentRound(ctx.StateDB())

	currentValidatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentRound)
	nextValidatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentRound+1)

	cache := make(map[basecommon.Address]struct{}, len(currentValidatorSnapQueue))

	for _, snap := range currentValidatorSnapQueue {
		cache[snap.ValidatorAddr] = struct{}{}
	}

	diffIds := make([]basecommon.Address, 0)

	for _, snap := range nextValidatorSnapQueue {
		if _, ok := cache[snap.ValidatorAddr]; !ok {
			diffIds = append(diffIds, snap.ValidatorAddr)
		}
	}

	if len(diffIds) == 0 {
		return nil
	}

	for _, id := range diffIds {
		v := db.GetValidator(ctx.StateDB(), s.Address(), id)
		pubkey, _ := v.PubKey.Pubkey()
		s.p2p.Addnode(enode.NewV4(pubkey, nil, 0, 0).URLv4())
	}

	return nil
}

func (s *StakeModule) IsEndOfRound(ctx sdk.ConsensusContext, blockNumber uint64) bool {
	return s.stageModule.IsEndOfRound(ctx.StateDB(), blockNumber)
}
func (s *StakeModule) GetRoundValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	round, startBlock, _ := s.stageModule.GetRoundAndBlockBoundByBlockNumber(ctx.StateDB(), blockNumber)

	validatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), round)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found round validators", "blockNumber", blockNumber, "round", round)
		return nil, errors.New("round validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorSnapQueue))

	for i, snap := range validatorSnapQueue {
		v := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)
		if v.IsEmpty() {
			continue
		}
		pubkey, _ := v.PubKey.Pubkey()
		blsKey := bls.PublicKey{}
		(&blsKey).Deserialize(v.BlsKey)

		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   basecommon.NodeAddress(snap.ValidatorAddr),
			PubKey:    pubkey,
			NodeID:    enode.PubkeyToIDV4(pubkey),
			BlsPubKey: &blsKey,
		}
		valMap[validator.NodeID] = validator
	}

	return &cbfttypes.Validators{
		Nodes:            valMap,
		ValidBlockNumber: startBlock,
	}, nil
}
func (s *StakeModule) BlocksOfRound(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	round := s.stageModule.GetRoundByBlockNumber(ctx.StateDB(), blockNumber)
	return s.stageModule.BlocksOfRound(ctx.StateDB(), round)
}
func (s *StakeModule) IsEndOfEpoch(ctx sdk.ConsensusContext, blockNumber uint64) bool {
	return s.stageModule.IsEndOfEpoch(ctx.StateDB(), blockNumber)
}
func (s *StakeModule) GetEpochValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	epoch, startBlock, _ := s.stageModule.GetEpochAndBlockBoundByBlockNumber(ctx.StateDB(), blockNumber)
	validatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), epoch)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found epoch validators", "blockNumber", blockNumber, "epoch", epoch)
		return nil, errors.New("epoch validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorSnapQueue))

	for i, snap := range validatorSnapQueue {
		v := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)
		if v.IsEmpty() {
			continue
		}

		pubkey, _ := v.PubKey.Pubkey()
		blsKey := bls.PublicKey{}
		(&blsKey).Deserialize(v.BlsKey)

		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   basecommon.NodeAddress(snap.ValidatorAddr),
			PubKey:    pubkey,
			NodeID:    enode.PubkeyToIDV4(pubkey),
			BlsPubKey: &blsKey,
		}
		valMap[validator.NodeID] = validator
	}

	return &cbfttypes.Validators{
		Nodes:            valMap,
		ValidBlockNumber: startBlock,
	}, nil
}
func (s *StakeModule) BlocksOfEpoch(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	epoch := s.stageModule.GetEpochByBlockNumber(ctx.StateDB(), blockNumber)
	return s.stageModule.BlocksOfEpoch(ctx.StateDB(), epoch)
}

func (s *StakeModule) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {

	if ctx.IsProposer() {
		currentValidatorAddr := crypto.PubkeyToAddress(s.nodePrivateKey.PublicKey)

		currentValidator := db.GetValidator(ctx.ParentStateDB(), s.Address(), currentValidatorAddr)
		if currentValidator.IsEmpty() {
			return errors.New("not found validator")
		}

		header.Coinbase = currentValidator.Owner
	}
	return nil
}
func (s *StakeModule) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return s.stageModule.GetLastNumber(ctx.StateDB(), blockNumber)
}

// round validator
func (s *StakeModule) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return s.GetRoundValidator(ctx, blockNumber)
}

func (s *StakeModule) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {

	currentEpoch := s.stageModule.GetCurrentEpoch(ctx.StateDB())

	validatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentEpoch)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found epoch validators", "epoch", currentEpoch)
		return false
	}

	for _, snap := range validatorSnapQueue {
		v := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)
		if v.IsEmpty() {
			continue
		}
		pubkey, _ := v.PubKey.Pubkey()
		if enode.PubkeyToIDV4(pubkey) == nodeID.ID() {
			return true
		}
	}

	return false
}

// internal

func (s *StakeModule) setNumberOfBlocksForRoundValidator(stateDB sdk.StateDB, header *types.Header) error {
	// Extract the validator public key of the build block based on the signature in the block header
	sign := header.Signature()
	sealhash := header.SealHash().Bytes()
	pk, err := crypto.SigToPub(sealhash, sign)
	if nil != err {
		return fmt.Errorf("can not sigToPub %s", err)
	}
	round := s.stageModule.GetRoundByBlockNumber(stateDB, header.Number.Uint64())

	// @TODO for debug ...
	validatorAddr := crypto.PubkeyToAddress(*pk)
	numberOfBlocks := db.GetNumberOfBlocksForRoundValidator(stateDB, s.Address(), validatorAddr, round)
	s.logger.Debug("setNumberOfBlocksForRoundValidator", "round", round, "header blockNumber", header.Number.Uint64(), "validatorAddr", validatorAddr.Hex(), "old numberOfBlocks", numberOfBlocks)
	db.IncrementNumberOfBlocksForRoundValidator(stateDB, s.Address(), validatorAddr, round, 1)

	return nil
}

func (s *StakeModule) electionRoundValidators(ctx sdk.WorkerContext, blockNumber uint64) error {

	currentRound := s.stageModule.GetCurrentRound(ctx.StateDB())
	if s.stageModule.IsNotElectionBlockOnCurrentRound(ctx.StateDB(), blockNumber) {
		return errors.New("block is not round electionBlock of current round")
	}

	currentEpoch := s.stageModule.GetCurrentEpoch(ctx.StateDB())

	currentRoundValidatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentRound)
	currentEpochValidatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentEpoch)

	if currentRoundValidatorSnapQueue.IsEmpty() {
		panic(fmt.Sprintf("the current round validators is empty, blockNumber: %d, round: %d", blockNumber, currentRound))
	}

	if currentEpochValidatorSnapQueue.IsEmpty() {
		panic(fmt.Sprintf("the current epoch validators is empty, blockNumber: %d, epoch: %d", blockNumber, currentEpoch))
	}

	unstakeValidatorAddrCache := make(map[basecommon.Address]struct{}, 0)
	maybeRemoveValidatorStatusCache := make(staketypes.MaybeRemoveValidatorStatusCache, 0)

	type validatorSharesItem struct {
		StakeAmount    *big.Int
		DelegateAmount *big.Int
	}
	validatorSharesCache := make(map[basecommon.Address]*validatorSharesItem, 0)

	for _, snap := range currentRoundValidatorSnapQueue {

		validator := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)

		// The validator details may have been removed due to being slashed
		if validator.IsEmpty() {
			maybeRemoveValidatorStatusCache[snap.ValidatorAddr] = staketypes.NotExist
		}

		if validator.IsInvalid() {
			maybeRemoveValidatorStatusCache[snap.ValidatorAddr] = validator.Status

			if validator.IsInvalidUnstaked() {
				unstakeValidatorAddrCache[snap.ValidatorAddr] = struct{}{}
			}
		}
		// Used to update the shares in the validator snapshot for the next rounds
		validatorSharesCache[snap.ValidatorAddr] = &validatorSharesItem{StakeAmount: snap.StakeAmount, DelegateAmount: snap.DelegateAmount}
	}

	diffValidatorSnapshotQueue := make(staketypes.ValidatorSortSnapshotQueue, 0)
	for _, snap := range currentEpochValidatorSnapQueue {

		// If a validator with a status of `unstaked` in the current round validatorSharesSnapshotQueue that
		// it still exists in the validatorSharesSnapshotQueue of the current epoch,
		// then continue to participate in the next round of validatorSharesSnapshotQueue elections
		if _, ok := unstakeValidatorAddrCache[snap.ValidatorAddr]; ok {
			delete(unstakeValidatorAddrCache, snap.ValidatorAddr)
			delete(maybeRemoveValidatorStatusCache, snap.ValidatorAddr)
		}

		// Extract the shares snapshot value from the current epoch validatorSnapshot to update the next round of validatorSnapshot shares
		if _, ok := validatorSharesCache[snap.ValidatorAddr]; ok {
			validatorSharesCache[snap.ValidatorAddr].StakeAmount = snap.StakeAmount
			validatorSharesCache[snap.ValidatorAddr].DelegateAmount = snap.DelegateAmount
			continue
		}

		validator := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)

		// Skip invalid validators (include status `unstake`)
		if validator.IsEmptyOrInvalid() {
			continue
		}

		// Collect validators that do not exist in the current round of validatorSharesSnapshotQueue
		// but exist in the current epoch validatorSharesSnapshotQueue,
		// for the election of the next round of validators
		diffValidatorSnapshotQueue = append(diffValidatorSnapshotQueue, snap)
	}

	maxRoundValidatorsSize := s.GetMaxRoundValidatorsSize(ctx.StateDB())

	shuffle := func(invalidLen int, currentRoundValidatorQueue, vrfValidatorSnapshotQueue staketypes.ValidatorSortSnapshotQueue, blockNumber uint64) (staketypes.ValidatorSortSnapshotQueue, error) {

		// increase term and use new shares  one by one
		for i, v := range currentRoundValidatorQueue {
			v.ValidatorTerm++
			v.StakeAmount = validatorSharesCache[v.ValidatorAddr].StakeAmount
			v.DelegateAmount = validatorSharesCache[v.ValidatorAddr].DelegateAmount
			currentRoundValidatorQueue[i] = v
		}

		// sort the validator by del rule
		currentRoundValidatorQueue.ValidatorSort(maybeRemoveValidatorStatusCache, staketypes.CompareForRemoveFromHead)
		// Increase term of validator
		copyQueue := make(staketypes.ValidatorSortSnapshotQueue, len(currentRoundValidatorQueue)-invalidLen)
		// Remove the invalid validators
		copy(copyQueue, currentRoundValidatorQueue[invalidLen:])
		return stakewrap.ShuffleQueue(ctx.StateDB(), s.vrfModule, copyQueue, vrfValidatorSnapshotQueue, blockNumber, maxRoundValidatorsSize)
	}

	var vrfValidatorSnapshotQueue staketypes.ValidatorSortSnapshotQueue
	var vrfQueueSize uint64
	// #### NOTE ####
	// The size of diffValidatorSnapshotQueue may be zero
	// (when the status of all validators in the currentEpochValidatorSnapQueue has 'unstake')
	if uint64(len(diffValidatorSnapshotQueue)) > maxRoundValidatorsSize {
		vrfQueueSize = maxRoundValidatorsSize
	} else {
		vrfQueueSize = uint64(len(diffValidatorSnapshotQueue))
	}

	if vrfQueueSize != 0 {

		if queue, err := stakewrap.ElectionValidatorByVRF(ctx.StateDB(), s.vrfModule, diffValidatorSnapshotQueue, blockNumber, vrfQueueSize); nil != err {
			s.logger.Error("Failed to call ElectionValidatorByVRF", "blockNumber", blockNumber, "err", err)
			return err
		} else {
			vrfValidatorSnapshotQueue = queue
		}
	}

	s.logger.Debug("Call electionRoundValidators statistics",
		"maybe remove current round validator count", len(maybeRemoveValidatorStatusCache), "unstake invalid validator count",
		unstakeValidatorAddrCache, "current round validators count", len(currentRoundValidatorSnapQueue), "MAX ROUND VALIDATOR SIZE ", maxRoundValidatorsSize,
		"maybe shift validator count", (maxRoundValidatorsSize-1)/3, "diff queue", len(diffValidatorSnapshotQueue),
		"vrf queue", len(vrfValidatorSnapshotQueue))

	nextRoundValidatorQueue, err := shuffle(len(maybeRemoveValidatorStatusCache), currentRoundValidatorSnapQueue, vrfValidatorSnapshotQueue, blockNumber)
	if nil != err {
		return err
	}

	if len(nextRoundValidatorQueue) == 0 {
		panic(fmt.Sprintf("The Next Round Validator is empty, blockNumber: %d, round: %d", blockNumber, currentRound))
	}

	if err := db.SetRoundValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentRound+1, nextRoundValidatorQueue); nil != err {
		s.logger.Error("Failed to call SetRoundValidatorSharesSnapshotQueue", "blockNumber", blockNumber, "err", err)
		return err
	}

	s.logger.Debug("Succeed to elected next round validators", "blockNumber", blockNumber, "round", currentRound, "validators size", len(nextRoundValidatorQueue))
	return nil
}

func (s *StakeModule) electionEpochValidators(ctx sdk.WorkerContext, blockNumber uint64) error {

	currentEpoch := s.stageModule.GetCurrentEpoch(ctx.StateDB())

	if s.stageModule.IsNotElectionBlockOnCurrentEpoch(ctx.StateDB(), blockNumber) {
		return errors.New("block is not endBlock of current epoch")
	}

	validatorIds := db.RankPriorityValidatorIds(ctx.StateDB(), s.Address(), s.GetMaxEpochValidatorsSize(ctx.StateDB()))

	if len(validatorIds) == 0 {
		return errors.New("not found validatorIds")
	}

	queue := staketypes.NewValidatorSharesSnapshotQueue(uint64(len(validatorIds)))

	for i, id := range validatorIds {

		validator := db.GetValidator(ctx.StateDB(), s.Address(), id)
		if validator.IsEmptyOrInvalid() {
			return errors.New("invalid validator")
		}
		queue[i] = staketypes.NewValidatorSharesSnapshot(id, validator.Epoch, validator.StakeIndex, validator.CommissionRate, validator.StakeAmount, validator.DelegateAmount)
	}

	if err := db.SetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentEpoch+1, queue); nil != err {
		s.logger.Error("Failed to store next epoch validators", "blockNumber", blockNumber, "epoch", currentEpoch, "error", err)
		return errors.New("store next epoch failed")
	}

	s.logger.Debug("Succeed to elected next epoch validators", "blockNumber", blockNumber, "epoch", currentEpoch, "validators size", len(queue))
	return nil
}

func (s *StakeModule) createSlashTx(ctx sdk.Context, txNonce uint64) (*types.Transaction, error) {

	method := contracts.Abi.Methods["slash"]

	input, err := method.Inputs.Pack()
	if nil != err {
		return nil, err
	}
	input = append(method.ID, input...)

	tx := types.NewTransaction(txNonce, s.Address(), nil, 100000, big.NewInt(0), input)
	chainId, _ := ctx.Backend().ChainId()
	signer := types.NewEIP155Signer(chainId)
	tx, err = types.SignTx(tx, signer, s.nodePrivateKey)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *StakeModule) updateValidatorStatus(stateDB sdk.StateDB, validatorAddr basecommon.Address, status staketypes.ValidatorStatus) error {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return nil
	}

	validator.AppendStatus(status)

	if status.IsInvalid() {

		// delete old priority
		priority := db.GetValidatorPriority(stateDB, s.Address(), validator.Epoch, validator.StakeIndex, validator.Shares())
		if priority.IsNotEmpty() {
			if priority.ValidatorAddr != validatorAddr {
				return db.ErrMisMatching
			}
			if err := db.RemoveValidatorPriority(stateDB, s.Address(), validator.Epoch, validator.StakeIndex, validator.Shares()); nil != err {
				return err
			}
		}
	}
	// set new priority only
	return db.SetValidator(stateDB, s.Address(), validatorAddr, validator)
}

// --- extern

func (s *StakeModule) GetRoundValidatorIds(stateDB sdk.StateDBReader, round uint64) []basecommon.Address {
	queue := db.GetRoundValidatorIds(stateDB, s.Address(), round)
	return queue
}
func (s *StakeModule) GetEpochValidatorIds(stateDB sdk.StateDBReader, epoch uint64) []basecommon.Address {
	queue := db.GetEpochValidatorIds(stateDB, s.Address(), epoch)
	return queue
}
func (s *StakeModule) GetRoundValidatorSnapQueueFlatten(stateDB sdk.StateDBReader, epoch uint64) ([]basecommon.Address, []*big.Int, []*big.Int, []uint64, []uint64, []uint64, []uint64) {
	queue := db.GetRoundValidatorSharesSnapshotQueue(stateDB, s.Address(), epoch)
	validatorAddrQueue := make([]basecommon.Address, len(queue))
	stakeAmountQueue := make([]*big.Int, len(queue))
	delegateAmountQueue := make([]*big.Int, len(queue))
	stakeEpochQueue := make([]uint64, len(queue))
	stakeIndexQueue := make([]uint64, len(queue))
	commissionRateQueue := make([]uint64, len(queue))
	validatorTermQueue := make([]uint64, len(queue))
	for i, _ := range queue {
		validatorAddrQueue[i] = queue[i].ValidatorAddr
		stakeAmountQueue[i] = queue[i].StakeAmount
		delegateAmountQueue[i] = queue[i].DelegateAmount
		stakeEpochQueue[i] = queue[i].Epoch
		stakeIndexQueue[i] = queue[i].StakeIndex
		commissionRateQueue[i] = queue[i].CommissionRate
		validatorTermQueue[i] = queue[i].ValidatorTerm
	}
	return validatorAddrQueue, stakeAmountQueue, delegateAmountQueue, commissionRateQueue, stakeEpochQueue, stakeIndexQueue, validatorTermQueue
}
func (s *StakeModule) GetEpochValidatorSnapQueueFlatten(stateDB sdk.StateDBReader, epoch uint64) ([]basecommon.Address, []*big.Int, []*big.Int, []uint64, []uint64, []uint64, []uint64) {
	queue := db.GetEpochValidatorSharesSnapshotQueue(stateDB, s.Address(), epoch)
	validatorAddrQueue := make([]basecommon.Address, len(queue))
	stakeAmountQueue := make([]*big.Int, len(queue))
	delegateAmountQueue := make([]*big.Int, len(queue))
	stakeEpochQueue := make([]uint64, len(queue))
	stakeIndexQueue := make([]uint64, len(queue))
	commissionRateQueue := make([]uint64, len(queue))
	validatorTermQueue := make([]uint64, len(queue))
	for i, _ := range queue {
		validatorAddrQueue[i] = queue[i].ValidatorAddr
		stakeAmountQueue[i] = queue[i].StakeAmount
		delegateAmountQueue[i] = queue[i].DelegateAmount
		stakeEpochQueue[i] = queue[i].Epoch
		stakeIndexQueue[i] = queue[i].StakeIndex
		commissionRateQueue[i] = queue[i].CommissionRate
		validatorTermQueue[i] = queue[i].ValidatorTerm
	}
	return validatorAddrQueue, stakeAmountQueue, delegateAmountQueue, commissionRateQueue, stakeEpochQueue, stakeIndexQueue, validatorTermQueue
}

func (s *StakeModule) IsValidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	return validator.IsValid()
}
func (s *StakeModule) IsInvalidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	return validator.IsEmptyOrInvalid()
}
func (s *StakeModule) IsOnlyInvalidUnstakeValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	return validator.IsOnlyInvalidUnstaked()
}

func (s *StakeModule) IsEmptyValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	return validator.IsEmpty()
}

func (s *StakeModule) GetValidatorECDSAPubKey(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *ecdsa.PublicKey {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return nil
	}
	pubkey, _ := validator.PubKey.Pubkey()
	return pubkey
}
func (s *StakeModule) GetValidatorBLSPubKey(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *bls.PublicKey {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return nil
	}
	blsKey := bls.PublicKey{}
	(&blsKey).Deserialize(validator.BlsKey)
	return &blsKey
}
func (s *StakeModule) GetValidatorCommissionRate(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) uint64 {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return 0
	}
	return validator.CommissionRate
}
func (s *StakeModule) GetValidatorStakeEpoch(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) uint64 {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return 0
	}
	return validator.Epoch
}
func (s *StakeModule) GetValidatorStakeAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return basecommon.Big0
	}
	return validator.StakeAmount
}
func (s *StakeModule) GetValidatorDelegateAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return basecommon.Big0
	}
	return validator.DelegateAmount
}
func (s *StakeModule) GetValidatorOwner(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) basecommon.Address {
	return db.GetValidatorOwner(stateDB, s.Address(), validatorAddr)
}
func (s *StakeModule) GetNumberOfBlocksForRoundValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address, round uint64) uint64 {
	return db.GetNumberOfBlocksForRoundValidator(stateDB, s.Address(), validatorAddr, round)
}
func (s *StakeModule) GetEpochByValidatorDelegationRcPending(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) []uint64 {
	epochs, _ := db.GetValidatorDelegationRcPendingAndEpoch(stateDB, s.Address(), validatorAddr, math.MaxUint64)
	return epochs
}
func (s *StakeModule) GetDelegationFlatten(stateDB sdk.StateDBReader, delegatorAddr, validatorAddr basecommon.Address, stakeEpoch uint64) (uint64, *big.Int) {
	delegation := db.GetDelegation(stateDB, s.Address(), delegatorAddr, validatorAddr, stakeEpoch)
	if delegation.IsEmpty() {
		return 0, basecommon.Big0
	}
	return delegation.Epoch, delegation.Amount
}
func (s *StakeModule) UpdateDelegationEpoch(stateDB sdk.StateDB, delegatorAddr, validatorAddr basecommon.Address, stakeEpoch, delegateEpoch uint64) error {
	delegation := db.GetDelegation(stateDB, s.Address(), delegatorAddr, validatorAddr, stakeEpoch)
	if delegation.IsEmpty() {
		return db.ErrNotFound
	}
	if delegation.Epoch == delegateEpoch {
		return nil
	}
	if delegation.Epoch > delegateEpoch {
		return fmt.Errorf("new delegate epoch not greater than old, old epoch: %d, new epoch: %d", delegation.Epoch, delegateEpoch)
	}
	delegation.UpdateEpoch(delegateEpoch)
	return db.SetDelegation(stateDB, s.Address(), delegatorAddr, validatorAddr, stakeEpoch, delegation)
}

// ------
func (s *StakeModule) GetStakeWithdrawalWaitPeriod(stateDB sdk.StateDBReader) uint64 {
	return db.GetStakeWithdrawalWaitPeriod(stateDB, s.Address())
}
func (s *StakeModule) GetDelegateWithdrawalWaitPeriod(stateDB sdk.StateDBReader) uint64 {
	return db.GetDelegateWithdrawalWaitPeriod(stateDB, s.Address())
}
func (s *StakeModule) GetSlashingPercentage(stateDB sdk.StateDBReader) uint64 {
	return db.GetSlashingPercentage(stateDB, s.Address())
}
func (s *StakeModule) GetSlashIncentivePercentage(stateDB sdk.StateDBReader) uint64 {
	return db.GetSlashIncentivePercentage(stateDB, s.Address())
}
func (s *StakeModule) GetMaxRoundValidatorsSize(stateDB sdk.StateDBReader) uint64 {
	return db.GetMaxRoundValidatorsSize(stateDB, s.Address())
}
func (s *StakeModule) GetMaxEpochValidatorsSize(stateDB sdk.StateDBReader) uint64 {
	return db.GetMaxEpochValidatorsSize(stateDB, s.Address())
}
func (s *StakeModule) GetMinBlocksOfRoundValidator(stateDB sdk.StateDBReader) uint64 {
	return db.GetMinBlocksOfRoundValidator(stateDB, s.Address())
}
