package staking

import (
	"encoding/json"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	stakecommon "github.com/PlatONnetwork/AppChain-SDK/x/staking/common"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	stakingp2p "github.com/PlatONnetwork/AppChain-SDK/x/staking/p2p"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
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

	return s.isEndOfRound(wctx, blockNumber)
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

	rounds, queue := db.GetRoundQueueAndIndexFromTail(wctx.StateDB(), address.StakeHandlerAddres, 100)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			round = rounds[i]
			startBlock = item.StartBlock
			break
		}
	}
	validatorIds := db.GetRoundValidatorIds(wctx.StateDB(), address.StakeHandlerAddres, round)
	if len(validatorIds) == 0 {
		s.logger.Error("Not found round validators", "blockNumber", blockNumber, "round", round)
		return nil, errors.New("round validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorIds))

	for i, validatorId := range validatorIds {
		v := db.GetValidator(wctx.StateDB(), address.StakeHandlerAddres, validatorId)
		if v.IsInvalid() {
			continue
		}
		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   common.NodeAddress(validatorId),
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

	return s.isEndOfEpoch(wctx, blockNumber)
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

	epochs, queue := db.GetEpochQueueAndIndexFromTail(wctx.StateDB(), address.StakeHandlerAddres, 100)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			epoch = epochs[i]
			startBlock = item.StartBlock
			break
		}
	}
	validatorIds := db.GetEpochValidatorIds(wctx.StateDB(), address.StakeHandlerAddres, epoch)
	if len(validatorIds) == 0 {
		s.logger.Error("Not found epoch validators", "blockNumber", blockNumber, "epoch", epoch)
		return nil, errors.New("epoch validators not found")
	}

	valMap := make(cbfttypes.ValidateNodeMap, len(validatorIds))

	for i, validatorId := range validatorIds {
		v := db.GetValidator(wctx.StateDB(), address.StakeHandlerAddres, validatorId)
		if v.IsInvalid() {
			continue
		}
		validator := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   common.NodeAddress(validatorId),
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
	return 0
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

	validatorIds := db.GetEpochValidatorIds(wctx.StateDB(), address.StakeHandlerAddres, epoch)
	if len(validatorIds) == 0 {
		s.logger.Error("Not found epoch validators", "epoch", epoch)
		return false
	}

	for _, validatorId := range validatorIds {
		v := db.GetValidator(wctx.StateDB(), address.StakeHandlerAddres, validatorId)
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

	// todo change current round AND current epoch

}
func (s *StakeModule) EndBlock(ctx sdk.Context) {

	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return
	}

	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()

	// election next round validators (at cuurent round electionBlock)
	if s.isCurrentElectionBlock(wctx, currentBlock) {
		if err := s.electionRoundValidators(wctx, currentBlock); nil != err {
			s.logger.Error("Failed to election round validators", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// election next epoch validators (at current epoch endBlock)
	if s.isEndOfEpoch(wctx, currentBlock) {
		if err := s.electionEpochValidators(wctx, currentBlock); nil != err {
			s.logger.Error("Failed to election epoch validators", "blockNumber", currentBlock, "error", err)
			return
		}
	}

	// todo calculation reward

	// todo record signBlocks of validator in round

	// todo write next round AND write next epoch

}

func (s *StakeModule) OnCommit(ctx sdk.Context, block *types.Block) error {
	wctx, ok := ctx.(sdk.WorkerContext)
	if !ok {
		s.logger.Error("Unexpeced sdk context", "ctx", reflect.TypeOf(ctx).String())
		return errors.New("unexpeced sdk context")
	}

	if s.isNotCurrentElectionBlock(wctx, block.NumberU64()) {
		return nil
	}

	currentRound := db.GetCurrentRound(wctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(wctx.StateDB(), address.StakeHandlerAddres, currentRound)

	if currentRoundItem.NextRound == math.MaxUint64 { // had not next round
		return errors.New("not found next round validators")
	}

	currentValidatorIds := db.GetRoundValidatorIds(wctx.StateDB(), address.StakeHandlerAddres, currentRound)
	nextValidatorIds := db.GetRoundValidatorIds(wctx.StateDB(), address.StakeHandlerAddres, currentRoundItem.NextRound)

	cache := make(map[common.Address]struct{}, len(currentValidatorIds))

	for _, id := range currentValidatorIds {
		cache[id] = struct{}{}
	}

	diffIds := make([]common.Address, 0)

	for _, id := range nextValidatorIds {
		if _, ok := cache[id]; !ok {
			diffIds = append(diffIds, id)
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

func (s *StakeModule) isCurrentElectionBlock(ctx sdk.WorkerContext, blockNumber uint64) bool {

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

func (s *StakeModule) isNotCurrentElectionBlock(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isCurrentElectionBlock(ctx, blockNumber)
}

func (s *StakeModule) isStartOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {

	// NOTE: Only search for the most recent 100 rounds to save resource consumption
	queue := db.GetRoundQueueFromTail(ctx.StateDB(), address.StakeHandlerAddres, 100)
	for _, item := range queue {
		if item.StartBlock == blockNumber {
			return true
		}
	}
	return false
}

func (s *StakeModule) isNotStartOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isStartOfRound(ctx, blockNumber)
}

func (s *StakeModule) isStartOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	if currentRoundItem.StartBlock == blockNumber {
		return true
	}
	return false
}

func (s *StakeModule) isNotStartOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isStartOfCurrentRound(ctx, blockNumber)
}

func (s *StakeModule) isStartOfNextRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	if currentRoundItem.EndBlock+1 == blockNumber {
		return true
	}
	return false
}

func (s *StakeModule) isNotStartOfNextRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isStartOfNextRound(ctx, blockNumber)
}

func (s *StakeModule) isEndOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {

	// NOTE: Only search for the most recent 100 rounds to save resource consumption
	queue := db.GetRoundQueueFromTail(ctx.StateDB(), address.StakeHandlerAddres, 100)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
}

func (s *StakeModule) isNotEndOfRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isEndOfRound(ctx, blockNumber)
}

func (s *StakeModule) isEndOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentRound := db.GetCurrentRound(ctx.StateDB(), address.StakeHandlerAddres)
	currentRoundItem := db.GetRoundItem(ctx.StateDB(), address.StakeHandlerAddres, currentRound)
	if currentRoundItem.EndBlock == blockNumber {
		return true
	}
	return false
}

func (s *StakeModule) isNotEndOfCurrentRound(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isEndOfCurrentRound(ctx, blockNumber)
}

func (s *StakeModule) isStartOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	// NOTE: Only search for the most recent 100 epochs to save resource consumption
	queue := db.GetEpochQueueFromTail(ctx.StateDB(), address.StakeHandlerAddres, 100)
	for _, item := range queue {
		if item.StartBlock == blockNumber {
			return true
		}
	}
	return false
}

func (s *StakeModule) isNotStartOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isStartOfEpoch(ctx, blockNumber)
}

func (s *StakeModule) isStartOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)
	if currentEpochItem.StartBlock == blockNumber {
		return true
	}
	return false
}

func (s *StakeModule) isNotStartOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isStartOfCurrentEpoch(ctx, blockNumber)
}

func (s *StakeModule) isStartOfNextEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)
	if currentEpochItem.EndBlock+1 == blockNumber {
		return true
	}
	return false
}

func (s *StakeModule) isNotStartOfNextEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isStartOfNextEpoch(ctx, blockNumber)
}

func (s *StakeModule) isEndOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {

	// NOTE: Only search for the most recent 100 epochs to save resource consumption
	queue := db.GetEpochQueueFromTail(ctx.StateDB(), address.StakeHandlerAddres, 100)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
}

func (s *StakeModule) isNotEndOfEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isEndOfEpoch(ctx, blockNumber)
}

func (s *StakeModule) isEndOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {

	currentEpoch := db.GetCurrentEpoch(ctx.StateDB(), address.StakeHandlerAddres)
	currentEpochItem := db.GetEpochItem(ctx.StateDB(), address.StakeHandlerAddres, currentEpoch)
	if currentEpochItem.EndBlock == blockNumber {
		return true
	}
	return false
}

func (s *StakeModule) isNotEndOfCurrentEpoch(ctx sdk.WorkerContext, blockNumber uint64) bool {
	return !s.isEndOfCurrentEpoch(ctx, blockNumber)
}

func (s *StakeModule) electionRoundValidators(ctx sdk.WorkerContext, blockNumber uint64) error {

	return nil
}

func (s *StakeModule) electionEpochValidators(ctx sdk.WorkerContext, blockNumber uint64) error {

	validatorIds := db.RankPriorityValidatorIds(ctx.StateDB(), address.StakeHandlerAddres, stakecommon.MAX_EPOCH_VALIDATORS_SIZE)

	return nil
}
