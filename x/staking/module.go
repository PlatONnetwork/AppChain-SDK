package staking

import (
	"encoding/json"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/PlatON-Go/common"
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

	// NOTE: Only search for the most recent 100 rounds to save resource consumption
	queue := db.GetRoundQueueFromTail(wctx.StateDB(), address.StakeHandlerAddres, 100)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
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

	// NOTE: Only search for the most recent 100 epochs to save resource consumption
	queue := db.GetEpochQueueFromTail(wctx.StateDB(), address.StakeHandlerAddres, 100)
	for _, item := range queue {
		if item.EndBlock == blockNumber {
			return true
		}
	}
	return false
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

}
func (s *StakeModule) EndBlock(ctx sdk.Context) {

}

func (s *StakeModule) OnCommit(ctx sdk.Context, block *types.Block) error {
	ctx.Backend()
	return nil
}
