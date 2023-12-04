package staking

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/l2"
	stakecommon "github.com/PlatONnetwork/AppChain-SDK/x/staking/common"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	stakingp2p "github.com/PlatONnetwork/AppChain-SDK/x/staking/p2p"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	stakewrap "github.com/PlatONnetwork/AppChain-SDK/x/staking/wrap"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"reflect"
)

type StakeModule struct {
	p2p          *stakingp2p.StakingP2P
	logger       log.Logger
	privateKey   *ecdsa.PrivateKey
	keystoreFile string
	passwordFile string
}

func NewStakeModule(ctx *cli.Context) *StakeModule {
	return &StakeModule{
		logger:       log.New("module", "staking"),
		keystoreFile: ctx.GlobalString(l2.KeystoreFlag.Name),
		passwordFile: ctx.GlobalString(l2.PasswordFlag.Name),
	}
}

func (s *StakeModule) Name() string {
	return "staking"
}

func (s *StakeModule) Init() error {

	key, err := l2.DecodePrivateKey(s.keystoreFile, s.passwordFile)
	if err != nil {
		return err
	}
	s.privateKey = key.PrivateKey
	return nil
}

func (s *StakeModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) {
	if err := initStakeHandler(db, s.Address()); nil != err {
		log.Error("Failed initialize StakeHandler", "error", err)
		panic(err)
	}
}

func (s *StakeModule) Address() basecommon.Address {
	return address.StakeHandlerAddress
}

func (s *StakeModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stakeHandler, _ := contracts.NewStakeHandler(evm, contract, readOnly)
	return stakeHandler.Run(input)
}

func (s *StakeModule) AddTxs(ctx sdk.WorkerContext, local, remote map[basecommon.Address]types.Transactions) (map[basecommon.Address]types.Transactions, map[basecommon.Address]types.Transactions) {

	blockNumber := ctx.Backend().CurrentHeader().Number.Uint64()

	if db.IsNotStartOfCurrentRound(ctx.StateDB(), s.Address(), blockNumber) {
		return local, remote
	}

	// check low blocks validtors of round, and send slash tx
	if db.HasNotLowBlocksValidator(ctx.StateDB(), s.Address()) {
		return local, remote
	}

	from := crypto.PubkeyToAddress(s.privateKey.PublicKey)

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

func (s *StakeModule) BeginBlock(ctx sdk.WorkerContext) {

	blockNumber := ctx.Backend().CurrentHeader().Number.Uint64()
	// NOTE: change current round at new round startBlock
	if db.IsBeginOfNextRound(ctx.StateDB(), s.Address(), blockNumber) {
		db.InrementCurrentRound(ctx.StateDB(), s.Address())
	}
	// NOTE: change current epoch at new epoch startBlock
	if db.IsBeginOfNextEpoch(ctx.StateDB(), s.Address(), blockNumber) {
		db.IncrementCurrentEpoch(ctx.StateDB(), s.Address())
	}

	// increase the number of validator blocks generated from the previous block
	parentNumber := blockNumber - 1
	parentHash := ctx.Backend().CurrentHeader().ParentHash
	if parentNumber != 0 {
		parentHeader := ctx.Backend().GetBlock(parentHash, parentNumber).Header()
		if err := stakewrap.SetNumberOfBlocksForRoundValidator(ctx.StateDB(), parentHeader); nil != err {
			panic(err)
		}
	}

}
func (s *StakeModule) EndBlock(ctx sdk.WorkerContext) {

	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()

	// election next round validators (at cuurent round electionBlock)
	if db.IsElectionBlockOnCurrentRound(ctx.StateDB(), s.Address(), currentBlock) {
		if err := s.electionRoundValidators(ctx, currentBlock); nil != err {
			s.logger.Error("Failed to election round validators", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// store next epochItem (at current round endBlock)
	// NOTE: Only search for the most recent 100 rounds to save resource consumption
	if db.IsEndOfCurrentRound(ctx.StateDB(), s.Address(), currentBlock) {
		if err := db.BuildNextRound(ctx.StateDB(), s.Address()); nil != err {
			s.logger.Error("Failed to build next round", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// election next epoch validators (at current epoch endBlock)
	// and store next epochItem
	// NOTE: Only search for the most recent 100 epochs to save resource consumption
	if db.IsEndOfCurrentEpoch(ctx.StateDB(), s.Address(), currentBlock) {
		if err := db.BuildNextEpoch(ctx.StateDB(), s.Address()); nil != err {
			s.logger.Error("Failed to build next epoch", "blockNumber", currentBlock, "error", err)
			return
		}
		if err := s.electionEpochValidators(ctx, currentBlock); nil != err {
			s.logger.Error("Failed to election epoch validators", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// todo calculation reward

}

func (s *StakeModule) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {

	if db.IsNotElectionBlockOnCurrentRound(ctx.StateDB(), s.Address(), block.NumberU64()) {
		return nil
	}

	currentRound := db.GetCurrentRound(ctx.StateDB(), s.Address())

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
		s.p2p.Addnode(enode.NewV4(v.PubKey, nil, 0, 0).URLv4())
	}

	return nil
}

func (s *StakeModule) IsEndOfRound(ctx sdk.ConsensusContext, blockNumber uint64) bool {
	// NOTE: Only search for the most recent 100 rounds to save resource consumption
	return db.IsEndOfRound(ctx.StateDB(), s.Address(), blockNumber, 100)
}
func (s *StakeModule) GetRoundValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {

	var (
		round      uint64
		startBlock uint64
	)

	currentRound := db.GetCurrentRound(ctx.StateDB(), s.Address())
	rounds, queue := db.GetRoundQueueAndIndexFromTail(ctx.StateDB(), s.Address(), currentRound, 100)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			round = rounds[i]
			startBlock = item.StartBlock
			break
		}
	}
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
		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   basecommon.NodeAddress(snap.ValidatorAddr),
			PubKey:    v.PubKey,
			NodeID:    enode.PubkeyToIDV4(v.PubKey),
			BlsPubKey: v.BlsKey,
		}
		valMap[validator.NodeID] = validator
	}

	return &cbfttypes.Validators{
		Nodes:            valMap,
		ValidBlockNumber: startBlock,
	}, nil
}
func (s *StakeModule) BlocksOfRound(ctx sdk.ConsensusContext) uint64 {
	round := db.GetCurrentRound(ctx.StateDB(), s.Address())
	roundItem := db.GetRoundItem(ctx.StateDB(), s.Address(), round)
	if roundItem.IsEmpty() {
		return 0
	}

	return roundItem.EndBlock - roundItem.StartBlock + 1
}
func (s *StakeModule) IsEndOfEpoch(ctx sdk.ConsensusContext, blockNumber uint64) bool {
	// NOTE: Only search for the most recent 100 epochs to save resource consumption
	return db.IsEndOfEpoch(ctx.StateDB(), s.Address(), blockNumber, 100)
}
func (s *StakeModule) GetEpochValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {

	var (
		epoch      uint64
		startBlock uint64
	)

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), s.Address())
	epochs, queue := db.GetEpochQueueAndIndexUtil(ctx.StateDB(), s.Address(), currentEpoch, 100)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			epoch = epochs[i]
			startBlock = item.StartBlock
			break
		}
	}
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
		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   basecommon.NodeAddress(snap.ValidatorAddr),
			PubKey:    v.PubKey,
			NodeID:    enode.PubkeyToIDV4(v.PubKey),
			BlsPubKey: v.BlsKey,
		}
		valMap[validator.NodeID] = validator
	}

	return &cbfttypes.Validators{
		Nodes:            valMap,
		ValidBlockNumber: startBlock,
	}, nil
}
func (s *StakeModule) BlocksOfEpoch(ctx sdk.ConsensusContext) uint64 {

	epoch := db.GetCurrentEpoch(ctx.StateDB(), s.Address())
	epochItem := db.GetEpochItem(ctx.StateDB(), s.Address(), epoch)
	if epochItem.IsEmpty() {
		return 0
	}

	return epochItem.EndBlock - epochItem.StartBlock + 1
}

func (s *StakeModule) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error { return nil }
func (s *StakeModule) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {

	var endBlock uint64
	currentRound := db.GetCurrentRound(ctx.StateDB(), s.Address())
	queue := db.GetRoundQueueUtil(ctx.StateDB(), s.Address(), currentRound, 100)
	for _, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			endBlock = item.EndBlock

			break
		}
	}
	return endBlock
}

// round validator
func (s *StakeModule) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return s.GetRoundValidator(ctx, blockNumber)
}

func (s *StakeModule) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {

	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return false
	}

	epoch := db.GetCurrentEpoch(wctx.StateDB(), s.Address())

	validatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(wctx.StateDB(), s.Address(), epoch)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found epoch validators", "epoch", epoch)
		return false
	}

	for _, snap := range validatorSnapQueue {
		v := db.GetValidator(wctx.StateDB(), s.Address(), snap.ValidatorAddr)
		if v.IsInvalid() {
			continue
		}

		if enode.PubkeyToIDV4(v.PubKey) == nodeID.ID() {
			return true
		}
	}

	return false
}

func (s *StakeModule) electionRoundValidators(ctx sdk.WorkerContext, blockNumber uint64) error {

	currentRound := db.GetCurrentRound(ctx.StateDB(), s.Address())
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), s.Address(), currentRound)
	if currentRoundItem.EndBlock-stakecommon.ROUND_VALIDATOR_ELECTION_DISTANCE != blockNumber {
		return errors.New("block is not round electionBlock of current round")
	}

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), s.Address())

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
		return stakewrap.ShuffleQueue(ctx.StateDB(), copyQueue, vrfValidatorSnapshotQueue, blockNumber)
	}

	var vrfValidatorSnapshotQueue staketypes.ValidatorSortSnapshotQueue
	var vrfQueueSize uint64
	if uint64(len(diffValidatorSnapshotQueue)) > stakecommon.MAX_ROUND_VALIDATORS_SIZE {
		vrfQueueSize = stakecommon.MAX_ROUND_VALIDATORS_SIZE
	} else {
		vrfQueueSize = uint64(len(diffValidatorSnapshotQueue))
	}

	if vrfQueueSize != 0 {
		if queue, err := stakewrap.ElectionValidatorByVRF(ctx.StateDB(), diffValidatorSnapshotQueue, blockNumber, vrfQueueSize); nil != err {
			s.logger.Error("Failed to call ElectionValidatorByVRF", "blockNumber", blockNumber, "err", err)
			return err
		} else {
			vrfValidatorSnapshotQueue = queue
		}
	}

	s.logger.Debug("Call electionRoundValidators statistics",
		"maybe remove current round validator count", len(maybeRemoveValidatorStatusCache), "unstake invalid validator count",
		unstakeValidatorAddrCache, "current round validators count", len(currentRoundValidatorSnapQueue), "MAX ROUND VALIDATOR SIZE ", stakecommon.MAX_ROUND_VALIDATORS_SIZE,
		"maybe shift validator count", (stakecommon.MAX_ROUND_VALIDATORS_SIZE-1)/3, "diff queue", len(diffValidatorSnapshotQueue),
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

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), s.Address())
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), s.Address(), currentEpoch)
	if currentEpochItem.IsEmpty() {
		return errors.New("not found currentEpochItem")
	}

	if currentEpochItem.EndBlock != blockNumber {
		return errors.New("block is not endBlock of current epoch")
	}

	validatorIds := db.RankPriorityValidatorIds(ctx.StateDB(), s.Address(), stakecommon.MAX_EPOCH_VALIDATORS_SIZE)

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
	from := crypto.PubkeyToAddress(s.privateKey.PublicKey)
	txNonce, err := ctx.Backend().GetPoolNonce(from)
	if nil != err {
		return nil, err
	}

	tx := types.NewTransaction(txNonce, s.Address(), nil, 100000, big.NewInt(0), input)
	chainId, _ := ctx.Backend().ChainId()
	signer := types.NewEIP155Signer(chainId)
	tx, err = types.SignTx(tx, signer, s.privateKey)
	if err != nil {
		return nil, err
	}

	return tx, nil
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
func (s *StakeModule) GetValidatorCommissionRate(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) uint64 {
	return db.GetValidator(stateDB, s.Address(), validatorAddr).CommissionRate
}
func (s *StakeModule) GetValidatorStakeAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int {
	return db.GetValidator(stateDB, s.Address(), validatorAddr).StakeAmount
}
func (s *StakeModule) GetValidatorDelegateAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int {
	return db.GetValidator(stateDB, s.Address(), validatorAddr).DelegateAmount
}
func (s *StakeModule) GetValidatorOwner(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) basecommon.Address {
	return db.GetValidator(stateDB, s.Address(), validatorAddr).Owner
}
func (s *StakeModule) GetCurrentRound(stateDB sdk.StateDBReader) uint64 {
	return db.GetCurrentRound(stateDB, s.Address())
}
func (s *StakeModule) GetCurrentEpoch(stateDB sdk.StateDBReader) uint64 {
	return db.GetCurrentEpoch(stateDB, s.Address())
}
func (s *StakeModule) IsBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsBeginOfCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StakeModule) IsBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsBeginOfCurrentEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StakeModule) IsEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsEndOfCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StakeModule) IsEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsEndOfCurrentEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StakeModule) GetNumberOfBlocksForRoundValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address, round uint64) uint64 {
	return db.GetNumberOfBlocksForRoundValidator(stateDB, s.Address(), validatorAddr, round)
}
