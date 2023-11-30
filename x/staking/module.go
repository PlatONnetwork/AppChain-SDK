package staking

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlayaNetwork/Alaya-Go/x/staking"
	"github.com/AlayaNetwork/Alaya-Go/x/xcom"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	stakecommon "github.com/PlatONnetwork/AppChain-SDK/x/staking/common"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	stakingp2p "github.com/PlatONnetwork/AppChain-SDK/x/staking/p2p"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
	"reflect"
)

type StakeModule struct {
	p2p    *stakingp2p.StakingP2P
	logger log.Logger
}

func NewStakeModule() *StakeModule {
	return &StakeModule{
		logger: log.New("module", "staking"),
	}
}

func (s *StakeModule) Name() string {
	return "staking"
}

func (s *StakeModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) {
	if err := initStakeHandler(db); nil != err {
		log.Error("Failed initialize StakeHandler", "error", err)
		panic(err)
	}

}

func (s *StakeModule) Address() common.Address {
	return address.StakeHandlerAddres
}

func (s *StakeModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stateReceiver, _ := contracts.NewStakeHandler(evm, contract, readOnly)
	return stateReceiver.Run(input)
}

func (s *StakeModule) IsEndOfRound(ctx sdk.Context, blockNumber uint64) bool {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return false
	}

	return isEndOfRound(wctx, blockNumber)
}
func (s *StakeModule) GetRoundValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error) {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return nil, errors.New("unexpeced sdk context")
	}

	var (
		round      uint64
		startBlock uint64
	)

	currentRound := db.GetCurrentRound(wctx.StateDB(), address.StakeHandlerAddres)
	rounds, queue := db.GetRoundQueueAndIndexFromTail(wctx.StateDB(), address.StakeHandlerAddres, currentRound, 100)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			round = rounds[i]
			startBlock = item.StartBlock
			break
		}
	}
	validatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(wctx.StateDB(), address.StakeHandlerAddres, round)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found round validators", "blockNumber", blockNumber, "round", round)
		return nil, errors.New("round validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorSnapQueue))

	for i, snap := range validatorSnapQueue {
		v := db.GetValidator(wctx.StateDB(), address.StakeHandlerAddres, snap.ValidatorAddr)
		if v.IsInvalid() {
			continue
		}
		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   common.NodeAddress(snap.ValidatorAddr),
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
func (s *StakeModule) BlocksOfRound(ctx sdk.Context) uint64 {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return 0
	}
	round := db.GetCurrentRound(wctx.StateDB(), address.StakeHandlerAddres)
	roundItem := db.GetRoundItem(wctx.StateDB(), address.StakeHandlerAddres, round)
	if roundItem.IsEmpty() {
		return 0
	}

	return roundItem.EndBlock - roundItem.StartBlock + 1
}
func (s *StakeModule) IsEndOfEpoch(ctx sdk.Context, blockNumber uint64) bool {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return false
	}

	return isEndOfEpoch(wctx, blockNumber)
}
func (s *StakeModule) GetEpochValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error) {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return nil, errors.New("unexpeced sdk context")
	}

	var (
		epoch      uint64
		startBlock uint64
	)

	currentEpoch := db.GetCurrentEpoch(wctx.StateDB(), address.StakeHandlerAddres)
	epochs, queue := db.GetEpochQueueAndIndexUtil(wctx.StateDB(), address.StakeHandlerAddres, currentEpoch, 100)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			epoch = epochs[i]
			startBlock = item.StartBlock
			break
		}
	}
	validatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(wctx.StateDB(), address.StakeHandlerAddres, epoch)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found epoch validators", "blockNumber", blockNumber, "epoch", epoch)
		return nil, errors.New("epoch validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorSnapQueue))

	for i, snap := range validatorSnapQueue {
		v := db.GetValidator(wctx.StateDB(), address.StakeHandlerAddres, snap.ValidatorAddr)
		if v.IsInvalid() {
			continue
		}
		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   common.NodeAddress(snap.ValidatorAddr),
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
func (s *StakeModule) BlocksOfEpoch(ctx sdk.Context) uint64 {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return 0
	}
	epoch := db.GetCurrentEpoch(wctx.StateDB(), address.StakeHandlerAddres)
	epochItem := db.GetEpochItem(wctx.StateDB(), address.StakeHandlerAddres, epoch)
	if epochItem.IsEmpty() {
		return 0
	}

	return epochItem.EndBlock - epochItem.StartBlock + 1
}

func (s *StakeModule) NewHeader(ctx sdk.Context, header *types.Header) error { return nil }
func (s *StakeModule) GetLastNumber(ctx sdk.Context, blockNumber uint64) uint64 {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return 0
	}

	var endBlock uint64
	currentRound := db.GetCurrentRound(wctx.StateDB(), address.StakeHandlerAddres)
	queue := db.GetRoundQueueUtil(wctx.StateDB(), address.StakeHandlerAddres, currentRound, 100)
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
func (s *StakeModule) GetValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error) {
	return s.GetRoundValidator(ctx, blockNumber)
}

func (s *StakeModule) IsCandidateNode(ctx sdk.Context, nodeID enode.IDv0) bool {

	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return false
	}

	epoch := db.GetCurrentEpoch(wctx.StateDB(), address.StakeHandlerAddres)

	validatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(wctx.StateDB(), address.StakeHandlerAddres, epoch)
	if len(validatorSnapQueue) == 0 {
		s.logger.Error("Not found epoch validators", "epoch", epoch)
		return false
	}

	for _, snap := range validatorSnapQueue {
		v := db.GetValidator(wctx.StateDB(), address.StakeHandlerAddres, snap.ValidatorAddr)
		if v.IsInvalid() {
			continue
		}

		if enode.PubkeyToIDV4(v.PubKey) == nodeID.ID() {
			return true
		}
	}

	return false
}

func (s *StakeModule) BeginBlock(ctx sdk.Context) {

	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return
	}

	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()
	// change current round at new round startBlock
	if isStartOfNextRound(wctx, currentBlock) {
		currentRound := db.GetCurrentRound(wctx.StateDB(), address.StakeHandlerAddres)
		db.SetCurrentRound(wctx.StateDB(), address.StakeHandlerAddres, currentRound+1)
	}
	// change current epoch at new epoch startBlock
	if isStartOfNextEpoch(wctx, currentBlock) {
		currentEpoch := db.GetCurrentEpoch(wctx.StateDB(), address.StakeHandlerAddres)
		db.SetCurrentEpoch(wctx.StateDB(), address.StakeHandlerAddres, currentEpoch+1)
	}

}
func (s *StakeModule) EndBlock(ctx sdk.Context) {

	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return
	}

	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()

	// election next round validators (at cuurent round electionBlock)
	if isCurrentElectionBlock(wctx, currentBlock) {
		if err := s.electionRoundValidators(wctx, currentBlock); nil != err {
			s.logger.Error("Failed to election round validators", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// store next epochItem (at current round endBlock)
	if isEndOfRound(wctx, currentBlock) {
		if err := buildNextRound(wctx); nil != err {
			s.logger.Error("Failed to build next round", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// election next epoch validators (at current epoch endBlock)
	// and store next epochItem
	if isEndOfEpoch(wctx, currentBlock) {
		if err := s.electionEpochValidators(wctx, currentBlock); nil != err {
			s.logger.Error("Failed to election epoch validators", "blockNumber", currentBlock, "error", err)
			return
		}
		if err := buildNextEpoch(wctx); nil != err {
			s.logger.Error("Failed to build next epoch", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// todo calculation reward

	// todo record signBlocks of validator in round

}

func (s *StakeModule) OnCommit(ctx sdk.Context, block *types.Block) error {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return errors.New("unexpeced sdk context")
	}

	if isNotCurrentElectionBlock(wctx, block.NumberU64()) {
		return nil
	}

	currentRound := db.GetCurrentRound(wctx.StateDB(), address.StakeHandlerAddres)

	currentValidatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(wctx.StateDB(), address.StakeHandlerAddres, currentRound)
	nextValidatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(wctx.StateDB(), address.StakeHandlerAddres, currentRound+1)

	cache := make(map[common.Address]struct{}, len(currentValidatorSnapQueue))

	for _, snap := range currentValidatorSnapQueue {
		cache[snap.ValidatorAddr] = struct{}{}
	}

	diffIds := make([]common.Address, 0)

	for _, snap := range nextValidatorSnapQueue {
		if _, ok := cache[snap.ValidatorAddr]; !ok {
			diffIds = append(diffIds, snap.ValidatorAddr)
		}
	}

	if len(diffIds) == 0 {
		return nil
	}

	for _, id := range diffIds {
		v := db.GetValidator(wctx.StateDB(), address.StakeHandlerAddres, id)
		s.p2p.Addnode(enode.NewV4(v.PubKey, nil, 0, 0).URLv4())
	}

	return nil
}

func isCurrentElectionBlock(ctx sdk.WorkerContext, blockNumber uint64) bool {

	round := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	roundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, round)
	if roundItem.IsEmpty() {
		return false
	}

	tmp := blockNumber + stakecommon.ROUND_VALIDATOR_ELECTION_DISTANCE
	if tmp == roundItem.EndBlock {
		return true
	}
	return false
}

func isNotCurrentElectionBlock(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isCurrentElectionBlock(ctx, blockNumber)
}

func isStartOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {

	// NOTE: Only search for the most recent 100 rounds to save resource consumption
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	queue := db.GetRoundQueueUtil(ctx.StateDB(), address.StakeHandlerAddres, currentRound, 100)
	for _, item := range queue {
		if item.StartBlock == blockNumber {
			return true
		}
	}
	return false
}

func isNotStartOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isStartOfRound(ctx, blockNumber)
}

func isStartOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	if currentRoundItem.StartBlock == blockNumber {
		return true
	}
	return false
}

func isNotStartOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isStartOfCurrentRound(ctx, blockNumber)
}

func isStartOfNextRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	if currentRoundItem.EndBlock+1 == blockNumber {
		return true
	}
	return false
}

func isNotStartOfNextRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isStartOfNextRound(ctx, blockNumber)
}

func isEndOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {

	// NOTE: Only search for the most recent 100 rounds to save resource consumption
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	queue := db.GetRoundQueueUtil(ctx.StateDB(), address.StakeHandlerAddres, currentRound, 100)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
}

func isNotEndOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isEndOfRound(ctx, blockNumber)
}

func isEndOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	if currentRoundItem.EndBlock == blockNumber {
		return true
	}
	return false
}

func isNotEndOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isEndOfCurrentRound(ctx, blockNumber)
}

func isStartOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	// NOTE: Only search for the most recent 100 epochs to save resource consumption
	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	queue := db.GetEpochQueueUtil(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch, 100)
	for _, item := range queue {
		if item.StartBlock == blockNumber {
			return true
		}
	}
	return false
}

func isNotStartOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isStartOfEpoch(ctx, blockNumber)
}

func isStartOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)
	if currentEpochItem.StartBlock == blockNumber {
		return true
	}
	return false
}

func isNotStartOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isStartOfCurrentEpoch(ctx, blockNumber)
}

func isStartOfNextEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)
	if currentEpochItem.EndBlock+1 == blockNumber {
		return true
	}
	return false
}

func isNotStartOfNextEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isStartOfNextEpoch(ctx, blockNumber)
}

func isEndOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {

	// NOTE: Only search for the most recent 100 epochs to save resource consumption
	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	queue := db.GetEpochQueueUtil(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch, 100)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
}

func isNotEndOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isEndOfEpoch(ctx, blockNumber)
}

func isEndOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)
	if currentEpochItem.EndBlock == blockNumber {
		return true
	}
	return false
}

func isNotEndOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !isEndOfCurrentEpoch(ctx, blockNumber)
}

func buildNextRound(ctx sdk.WorkerContext) error {

	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)

	startBlock := currentRoundItem.EndBlock + 1
	endBlock := currentRoundItem.EndBlock + stakecommon.ROUND_SIZE
	return db.AppendRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound+1, startBlock, endBlock)
}

func buildNextEpoch(ctx sdk.WorkerContext) error {

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)

	startBlock := currentEpochItem.EndBlock + 1
	endBlock := currentEpochItem.EndBlock + stakecommon.EPOCH_SIZE
	roundCount := stakecommon.EPOCH_SIZE / stakecommon.ROUND_SIZE
	return db.AppendEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch+1, startBlock, endBlock, roundCount)
}

func (s *StakeModule) electionRoundValidators(ctx sdk.WorkerContext, blockNumber uint64) error {

	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	if currentRoundItem.EndBlock-stakecommon.ROUND_VALIDATOR_ELECTION_DISTANCE != blockNumber {
		return errors.New("block is not round electionBlock of current round")
	}

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)

	currentRoundValidatorSnapQueue := db.GetRoundValidatorSharesSnapshotQueue(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	currentEpochValidatorSnapQueue := db.GetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)

	if currentRoundValidatorSnapQueue.IsEmpty() {
		panic(fmt.Sprintf("the current round validators is empty, blockNumber: %d, round: %d", blockNumber, currentRound))
	}

	if currentEpochValidatorSnapQueue.IsEmpty() {
		panic(fmt.Sprintf("the current epoch validators is empty, blockNumber: %d, epoch: %d", blockNumber, currentEpoch))
	}

	unstakeValidatorAddrCache := make(map[common.Address]struct{}, 0)
	maybeRemoveValidatorStatusCache := make(staketypes.MaybeRemoveValidatorStatusCache, 0)

	type validatorSharesItem struct {
		StakeAmount    *big.Int
		DelegateAmount *big.Int
	}
	validatorSharesCache := make(map[common.Address]*validatorSharesItem, 0)

	for _, snap := range currentRoundValidatorSnapQueue {

		validator := db.GetValidator(ctx.StateDB(), address.StakeHandlerAddres, snap.ValidatorAddr)

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

	diffValidatorSnapshotQueue := make(staketypes.ValidatorSharesSnapshotQueue, 0)
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

		validator := db.GetValidator(ctx.StateDB(), address.StakeHandlerAddres, snap.ValidatorAddr)

		// Skip invalid validators
		if validator.IsEmpty() || validator.IsInvalid() {
			continue
		}

		// Collect validators that do not exist in the current round of validatorSharesSnapshotQueue
		// but exist in the current epoch validatorSharesSnapshotQueue,
		// for the election of the next round of validators
		diffValidatorSnapshotQueue = append(diffValidatorSnapshotQueue, snap)
	}

	//shuffle := func(invalidLen int, currentRoundValidatorQueue, vrfQueue staketypes.ValidatorSharesSnapshotQueue, blockNumber uint64, parentHash common.Hash) (staking.ValidatorQueue, error) {
	//
	//	// increase term and use new shares  one by one
	//	for i, v := range currentRoundValidatorQueue {
	//		v.ValidatorTerm++
	//		v.Shares = currMap[v.NodeId]
	//		currentRoundValidatorQueue[i] = v
	//	}
	//
	//	// sort the validator by del rule
	//	currentRoundValidatorQueue.ValidatorSort(maybeRemoveValidatorStatusCache, staketypes.CompareForRemoveFromHead)
	//	// Increase term of validator
	//	copyQueue := make(staking.ValidatorQueue, len(currentRoundValidatorQueue)-invalidLen)
	//	// Remove the invalid validators
	//	copy(copyQueue, currentRoundValidatorQueue[invalidLen:])
	//	return shuffleQueue(copyQueue, vrfQueue, blockNumber, parentHash)
	//}

	return nil
}

func shuffleQueue(remainCurrQueue, vrfQueue staking.ValidatorQueue, blockNumber uint64, parentHash common.Hash) (staking.ValidatorQueue, error) {

	remainLen := len(remainCurrQueue)
	totalQueue := append(remainCurrQueue, vrfQueue...)

	for remainLen > int(xcom.MaxConsensusVals()-xcom.ShiftValidatorNum()) && len(totalQueue) > int(xcom.MaxConsensusVals()) {
		totalQueue = totalQueue[1:]
		remainLen--
	}

	if len(totalQueue) > int(xcom.MaxConsensusVals()) {
		totalQueue = totalQueue[:xcom.MaxConsensusVals()]
	}

	next := make(staking.ValidatorQueue, len(totalQueue))

	copy(next, totalQueue)

	// Divide all consensus nodes into two groups, the front and back positions of each group are not changed,
	// but random ordering is performed in each group
	// The first group: the first f nodes
	// The second group: the last 2f + 1 nodes
	next, err := randomOrderValidatorQueue(blockNumber, parentHash, next)
	if nil != err {
		return nil, err
	}
	return next, nil
}

func (s *StakeModule) electionEpochValidators(ctx sdk.WorkerContext, blockNumber uint64) error {

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)
	if currentEpochItem.IsEmpty() {
		return errors.New("not found currentEpochItem")
	}

	if currentEpochItem.EndBlock != blockNumber {
		return errors.New("block is not endBlock of current epoch")
	}

	validatorIds := db.RankPriorityValidatorIds(ctx.StateDB(), address.StakeHandlerAddres, stakecommon.MAX_EPOCH_VALIDATORS_SIZE)

	if len(validatorIds) == 0 {
		return errors.New("not found validatorIds")
	}

	queue := make(staketypes.ValidatorSharesSnapshotQueue, len(validatorIds))
	for i, id := range validatorIds {

		validator := db.GetValidator(ctx.StateDB(), address.StakeHandlerAddres, id)
		if validator.IsInvalid() {
			return errors.New("invalid validator")
		}
		queue[i] = staketypes.NewValidatorSharesSnapshot(id, validator.Epoch, validator.StakeIndex, validator.StakeAmount, validator.DelegateAmount)
	}

	if err := db.SetEpochValidatorSharesSnapshotQueue(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch+1, queue); nil != err {
		s.logger.Error("Failed to store next epoch validators", "blockNumber", blockNumber, "epoch", currentEpoch, "error", err)
		return errors.New("store next epoch failed")
	}

	s.logger.Debug("Succeed to elected next epoch validators", "blockNumber", blockNumber, "epoch", currentEpoch, "validators size", len(queue))
	return nil
}
