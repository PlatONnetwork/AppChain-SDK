package staking

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
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
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
)

const (
	MODULE_NAME_STAKING = "staking"
)

type StakeModule struct {
	p2p            *stakingp2p.StakingP2P
	logger         log.Logger
	nodePrivateKey *ecdsa.PrivateKey
	l1Module       staketypes.L1Moduler
	stageModule    staketypes.StageModuler
	vrfModule      staketypes.VRFModuler
	rewardModule   staketypes.RewardModuler
}

func NewStakeModule(ctx *cli.Context, l1Module staketypes.L1Moduler, stage staketypes.StageModuler) *StakeModule {
	return &StakeModule{
		logger:      log.New("module", MODULE_NAME_STAKING),
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
	return MODULE_NAME_STAKING
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

	log.Info("Succeed init genesis", "module", s.Name(), "StakeNetworkParams", configParams.String())
	return nil
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

func (s *StakeModule) AddTxs(ctx sdk.WorkerContext, local, remote map[basecommon.Address]types.Transactions) (map[basecommon.Address]types.Transactions, map[basecommon.Address]types.Transactions) {

	blockNumber := ctx.Header().Number.Uint64()
	if blockNumber == 0 {
		return local, remote
	}

	if s.stageModule.IsNotBeginOfCurrentRound(ctx.StateDB(), blockNumber) {
		return local, remote
	}

	// check low blocks validtors of round, and send slash tx
	if db.HasNotLowBlocksValidator(ctx.StateDB(), s.Address(), s.GetMinRoundValidatorBlockNumber(ctx.StateDB())) {
		return local, remote
	}

	from := crypto.PubkeyToAddress(s.nodePrivateKey.PublicKey)

	slashTx, err := s.createSlashTx(ctx)
	if nil != err {
		s.logger.Error("Failed to create slash tx", "blockNumber", blockNumber, "error", err)
		return local, remote
	}
	if nil == local[from] {
		local[from] = make(types.Transactions, 0)
	}
	local[from] = append(local[from], slashTx)
	return local, remote
}

func (s *StakeModule) BeginBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}

	// increase the number of validator blocks generated from the previous block
	parentBlock := currentBlock - 1
	parentHash := ctx.Header().ParentHash
	if parentBlock != 0 {
		parentHeader := ctx.Backend().GetBlock(parentHash, parentBlock).Header()
		if err := s.setNumberOfBlocksForRoundValidator(ctx.StateDB(), parentHeader); nil != err {
			return fmt.Errorf("can not set number of blocks for round validators, %s, parentBlock: %d", err, parentBlock)
		}
	}

	if s.stageModule.IsBeginOfCurrentRound(ctx.StateDB(), currentBlock) {
		// check low blocks validators
		lowBlocksValidatorAddrQueue := db.CheckLowBlocksValidatorForPreviousRound(ctx.StateDB(), s.Address(), s.GetMinRoundValidatorBlockNumber(ctx.StateDB()))
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
	// NOTE: Optimization of queries, search for the validator list for the last 100 rounds
	return s.stageModule.IsEndOfRound(ctx.StateDB(), blockNumber, 100)
}
func (s *StakeModule) GetRoundValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {

	// NOTE: Optimization of queries, search for the validator list for the last 100 rounds
	round, startBlock, _ := s.stageModule.GetRoundAndBlockBoundByBlockNumber(ctx.StateDB(), blockNumber, 100)

	validatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), round)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found round validators", "blockNumber", blockNumber, "round", round)
		return nil, errors.New("round validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorSnapQueue))

	for i, snap := range validatorSnapQueue {
		v := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)
		if v.IsInvalid() {
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
func (s *StakeModule) BlocksOfRound(ctx sdk.ConsensusContext) uint64 {
	round := s.stageModule.GetCurrentRound(ctx.StateDB())
	return s.stageModule.BlocksOfRound(ctx.StateDB(), round)
}
func (s *StakeModule) IsEndOfEpoch(ctx sdk.ConsensusContext, blockNumber uint64) bool {
	// NOTE: Optimization of queries, search for the validator list for the last 100 epochs
	return s.stageModule.IsEndOfEpoch(ctx.StateDB(), blockNumber, 100)
}
func (s *StakeModule) GetEpochValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {

	// NOTE: Optimization of queries, search for the validator list for the last 100 epochs
	epoch, startBlock, _ := s.stageModule.GetEpochAndBlockBoundByBlockNumber(ctx.StateDB(), blockNumber, 100)
	validatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), epoch)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found epoch validators", "blockNumber", blockNumber, "epoch", epoch)
		return nil, errors.New("epoch validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorSnapQueue))

	for i, snap := range validatorSnapQueue {
		v := db.GetValidator(ctx.StateDB(), s.Address(), snap.ValidatorAddr)
		if v.IsInvalid() {
			continue
		}

		pubkey, _ := v.PubKey.Pubkey()
		blsKey := bls.PublicKey{}
		(&blsKey).DeserializeUncompressed(v.BlsKey)

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
func (s *StakeModule) BlocksOfEpoch(ctx sdk.ConsensusContext) uint64 {
	epoch := s.stageModule.GetCurrentEpoch(ctx.StateDB())
	return s.stageModule.BlocksOfEpoch(ctx.StateDB(), epoch)
}

func (s *StakeModule) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {

	if ctx.IsProposer() {
		currentValidatorAddr := crypto.PubkeyToAddress(s.nodePrivateKey.PublicKey)

		currentValidator := db.GetValidator(ctx.ParentStateDB(), s.Address(), currentValidatorAddr)
		if currentValidator.IsInvalid() {
			return errors.New("invalida validator")
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
		if v.IsInvalid() {
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
	round, err := s.stageModule.GetRoundByBlockNumber(stateDB, header.Number.Uint64())
	if nil != err {
		return fmt.Errorf("get round by block %s", err)
	}

	db.IncrementNumberOfBlocksForRoundValidator(stateDB, s.Address(), crypto.PubkeyToAddress(*pk), round, 1)

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

		// Skip invalid validators
		if validator.IsEmpty() || validator.IsInvalid() {
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

	queue := make(staketypes.ValidatorSortSnapshotQueue, len(validatorIds))
	for i, id := range validatorIds {

		validator := db.GetValidator(ctx.StateDB(), s.Address(), id)
		if validator.IsInvalid() {
			return errors.New("invalid validator")
		}
		queue[i] = staketypes.NewValidatorSharesSnapshot(id, validator.Epoch, validator.StakeIndex, validator.StakeAmount, validator.DelegateAmount)
	}

	if err := db.SetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), s.Address(), currentEpoch+1, queue); nil != err {
		s.logger.Error("Failed to store next epoch validators", "blockNumber", blockNumber, "epoch", currentEpoch, "error", err)
		return errors.New("store next epoch failed")
	}

	s.logger.Debug("Succeed to elected next epoch validators", "blockNumber", blockNumber, "epoch", currentEpoch, "validators size", len(queue))
	return nil
}

func (s *StakeModule) createSlashTx(ctx sdk.Context) (*types.Transaction, error) {

	input, err := contracts.Abi.Methods["slash"].Inputs.Pack()
	if nil != err {
		return nil, err
	}
	from := crypto.PubkeyToAddress(s.nodePrivateKey.PublicKey)
	txNonce, err := ctx.Backend().GetPoolNonce(from)
	if nil != err {
		return nil, err
	}

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
	old := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if old.IsEmpty() {
		return errors.New("has not validator")
	}
	old.AppendStatus(status)

	if status.IsInvalid() {

		// delete old priority
		if db.GetValidatorPriority(stateDB, s.Address(), old.Epoch, old.StakeIndex, old.Shares()).ValidatorAddr != validatorAddr {
			return db.ErrMisMatching
		}
		if err := db.RemoveValidatorPriority(stateDB, s.Address(), old.Epoch, old.StakeIndex, old.Shares()); nil != err {
			return err
		}
	}

	return db.SetValidator(stateDB, s.Address(), validatorAddr, old)
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
func (s *StakeModule) IsValidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	return validator.IsValid()
}
func (s *StakeModule) IsInvalidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool {
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	return validator.IsInvalid()
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
	(&blsKey).DeserializeUncompressed(validator.BlsKey)
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
	validator := db.GetValidator(stateDB, s.Address(), validatorAddr)
	if validator.IsEmpty() {
		return basecommon.ZeroAddr
	}
	return validator.Owner
}
func (s *StakeModule) GetNumberOfBlocksForRoundValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address, round uint64) uint64 {
	return db.GetNumberOfBlocksForRoundValidator(stateDB, s.Address(), validatorAddr, round)
}
func (s *StakeModule) GetEpochByValidatorDelegationRcPending(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) []uint64 {
	epochs, _ := db.GetValidatorDelegationRcPendingAndEpoch(stateDB, s.Address(), validatorAddr, math.MaxUint64)
	return epochs
}
func (s *StakeModule) GetDelegationFlatten(stateDB sdk.StateDBReader, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) (uint64, *big.Int) {
	delegation := db.GetDelegation(stateDB, s.Address(), delegaterAddr, validatorAddr, stakeEpoch)
	if delegation.IsEmpty() {
		return 0, basecommon.Big0
	}
	return delegation.Epoch, delegation.Amount
}
func (s *StakeModule) UpdateDelegationEpoch(stateDB sdk.StateDB, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch, delegateEpoch uint64) error {
	del := db.GetDelegation(stateDB, s.Address(), delegaterAddr, validatorAddr, stakeEpoch)
	if nil == del {
		return db.ErrNotFound
	}
	if del.Epoch == delegateEpoch {
		return nil
	}
	if del.Epoch > delegateEpoch {
		return fmt.Errorf("new delegate epoch not greater than old, old epoch: %d, new epoch: %d", del.Epoch, delegateEpoch)
	}
	del.UpdateEpoch(delegateEpoch)
	return db.SetDelegation(stateDB, s.Address(), delegaterAddr, validatorAddr, stakeEpoch, del)
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
func (s *StakeModule) GetMinRoundValidatorBlockNumber(stateDB sdk.StateDBReader) uint64 {
	return db.GetMinRoundValidatorBlockNumber(stateDB, s.Address())
}
