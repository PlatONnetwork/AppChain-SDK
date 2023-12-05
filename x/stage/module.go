package stage

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/db"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type StageModule struct {
	logger log.Logger
}

func NewStageModule() *StageModule {
	return &StageModule{
		logger: log.New("module", "stage"),
	}
}

func (s *StageModule) Name() string {
	return "stage"
}

func (s *StageModule) Address() basecommon.Address {
	return constants.RewardManagerAddress
}

func (s *StageModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	return nil, nil
}

func (s *StageModule) BeginBlock(ctx sdk.WorkerContext) {
	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()
	// NOTE: change current round at new round startBlock
	if db.IsBeginOfNextRound(ctx.StateDB(), s.Address(), currentBlock) {
		db.InrementCurrentRound(ctx.StateDB(), s.Address())
	}
	// NOTE: change current epoch at new epoch startBlock
	if db.IsBeginOfNextEpoch(ctx.StateDB(), s.Address(), currentBlock) {
		db.IncrementCurrentEpoch(ctx.StateDB(), s.Address())
	}
}

func (s *StageModule) EndBlock(ctx sdk.WorkerContext) {

	currentBlock := ctx.Backend().CurrentHeader().Number.Uint64()

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
	}
}

func (s *StageModule) GetCurrentRound(stateDB sdk.StateDBReader) uint64 {
	return db.GetCurrentRound(stateDB, s.Address())
}
func (s *StageModule) GetCurrentEpoch(stateDB sdk.StateDBReader) uint64 {
	return db.GetCurrentEpoch(stateDB, s.Address())
}

func (s *StageModule) IsElectionBlockOnCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsElectionBlockOnCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsElectionBlockOnCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsElectionBlockOnCurrentEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsBeginOfCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsBeginOfCurrentEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsBeginOfNextRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsBeginOfNextRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsBeginOfNextEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsBeginOfNextEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsEndOfCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsEndOfCurrentEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsEndOfRound(stateDB sdk.StateDBReader, blockNumber, size uint64) bool {
	return db.IsEndOfRound(stateDB, s.Address(), blockNumber, size)
}
func (s *StageModule) IsEndOfEpoch(stateDB sdk.StateDBReader, blockNumber, size uint64) bool {
	return db.IsEndOfEpoch(stateDB, s.Address(), blockNumber, size)
}

func (s *StageModule) IsNotElectionBlockOnCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotElectionBlockOnCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsNotElectionBlockOnCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotElectionBlockOnCurrentEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsNotBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotBeginOfCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsNotBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsBeginOfCurrentEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsNotBeginOfNextRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotBeginOfNextRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsNotBeginOfNextEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotBeginOfNextEpoch(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsNotEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotEndOfCurrentRound(stateDB, s.Address(), blockNumber)
}
func (s *StageModule) IsNotEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotEndOfCurrentEpoch(stateDB, s.Address(), blockNumber)
}

func (s *StageModule) IsNotEndOfRound(stateDB sdk.StateDBReader, blockNumber, size uint64) bool {
	return db.IsNotEndOfRound(stateDB, s.Address(), blockNumber, size)
}
func (s *StageModule) IsNotEndOfEpoch(stateDB sdk.StateDBReader, blockNumber, size uint64) bool {
	return db.IsNotEndOfEpoch(stateDB, s.Address(), blockNumber, size)
}

func (s *StageModule) BlocksOfRound(stateDB sdk.StateDBReader, round uint64) uint64 {
	roundItem := db.GetRoundItem(stateDB, s.Address(), round)
	if roundItem.IsEmpty() {
		return 0
	}

	return roundItem.EndBlock - roundItem.StartBlock + 1
}

func (s *StageModule) BlocksOfEpoch(stateDB sdk.StateDBReader, epoch uint64) uint64 {
	epochItem := db.GetEpochItem(stateDB, s.Address(), epoch)
	if epochItem.IsEmpty() {
		return 0
	}

	return epochItem.EndBlock - epochItem.StartBlock + 1
}

func (s *StageModule) GetLastNumber(stateDB sdk.StateDBReader, blockNumber uint64) uint64 {
	var endBlock uint64
	currentRound := db.GetCurrentRound(stateDB, s.Address())
	queue := db.GetRoundQueueUtil(stateDB, s.Address(), currentRound, 100)
	for _, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			endBlock = item.EndBlock

			break
		}
	}
	return endBlock
}

func (s *StageModule) GetRoundAndBlockBoundByBlockNumber(stateDB sdk.StateDBReader, blockNumber, size uint64) (uint64, uint64, uint64) {

	var (
		round      uint64
		startBlock uint64
		endBlock   uint64
	)
	currentRound := db.GetCurrentRound(stateDB, s.Address())
	rounds, queue := db.GetRoundQueueAndIndexFromTail(stateDB, s.Address(), currentRound, size)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			round = rounds[i]
			startBlock = item.StartBlock
			endBlock = item.EndBlock
			break
		}
	}
	return round, startBlock, endBlock
}

func (s *StageModule) GetEpochAndBlockBoundByBlockNumber(stateDB sdk.StateDBReader, blockNumber, size uint64) (uint64, uint64, uint64) {

	var (
		epoch      uint64
		startBlock uint64
		endBlock   uint64
	)
	currentEpoch := db.GetCurrentEpoch(stateDB, s.Address())
	epochs, queue := db.GetEpochQueueAndIndexUtil(stateDB, s.Address(), currentEpoch, size)
	for i, item := range queue {
		// [startBlock, endBlock) || (startBlock, endBlock]
		if item.StartBlock <= blockNumber && item.EndBlock >= blockNumber {

			epoch = epochs[i]
			startBlock = item.StartBlock
			endBlock = item.EndBlock
			break
		}
	}

	return epoch, startBlock, endBlock
}
