package stage

import (
	"encoding/json"
	"fmt"
	sdkcontracts "github.com/PlatONnetwork/AppChain-SDK/contracts"
	contracts2 "github.com/PlatONnetwork/AppChain-SDK/x/stage/contracts"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/config"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/db"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"gopkg.in/urfave/cli.v1"
)

const (
	ModuleName    = "stage"
	ModuleVersion = 0
)

var _ module.ContractModule = (*StageModule)(nil)

type StageModule struct {
	logger log.Logger
}

func NewModule(ctx *cli.Context) *StageModule {
	return &StageModule{
		logger: log.New("module", ModuleName),
	}
}

func (s *StageModule) Name() string {
	return ModuleName
}

func (s *StageModule) Version() uint64 {
	return ModuleVersion
}

func (s *StageModule) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {

	configParams := config.DefualtStageNetworkParams()
	raw, err := data.MarshalJSON()
	if nil != err {
		log.Error("Failed MarshalJSON StageNetworkParams bytes", "error", err)
		return err
	}

	var conf config.StageNetworkParams
	if err := json.Unmarshal(raw, &conf); nil != err {
		log.Error("Failed UnmarshalJSON StageNetworkParams", "error", err)
		return err
	} else {
		configParams = &conf
	}

	stage, _ := contracts2.NewStageManager(sdkcontracts.NewEVM(db, big.NewInt(0)), sdkcontracts.NewContract(s, s), false)
	stage.InitGenesis(0)
	// set config params
	initConfigParams(db, s.Address(), configParams)

	if err := initGenesisRoundItem(db, s.Address(), configParams); nil != err {
		log.Error("Failed initialize genesis round", "error", err)
		return err
	}

	if err := initGenesisEpochItem(db, s.Address(), configParams); nil != err {
		log.Error("Failed initialize genesis epoch", "error", err)
		return err
	}

	log.Info("Succeed init genesis", "module", s.Name(), "StageNetworkParams", configParams.String())
	return nil
}

func (s *StageModule) Address() basecommon.Address {
	return constants.StageManagerAddress
}

func (s *StageModule) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	return nil, nil
}

func (s *StageModule) ContractCreateBlockNumber(statedb sdk.StateDBReader) uint64 {
	// TODO: implement me
	return 0
}

func (s *StageModule) BeginBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}
	// NOTE: change current round at new round startBlock
	if db.IsBeginOfNextRound(ctx.StateDB(), s.Address(), currentBlock) {
		oldRound := s.GetCurrentRound(ctx.StateDB())
		db.InrementCurrentRound(ctx.StateDB(), s.Address())
		newRound := s.GetCurrentRound(ctx.StateDB())
		log.Debug("Succeed InrementCurrentRound", "oldRound", oldRound, "newRound", newRound, "currentBlock", currentBlock)
	}
	// NOTE: change current epoch at new epoch startBlock
	if db.IsBeginOfNextEpoch(ctx.StateDB(), s.Address(), currentBlock) {
		oldEpoch := s.GetCurrentEpoch(ctx.StateDB())
		db.IncrementCurrentEpoch(ctx.StateDB(), s.Address())
		newEpoch := s.GetCurrentEpoch(ctx.StateDB())
		log.Debug("Succeed IncrementCurrentEpoch", "oldEpoch", oldEpoch, "newEpoch", newEpoch, "currentBlock", currentBlock)
	}
	return nil
}

func (s *StageModule) EndBlock(ctx sdk.WorkerContext) error {

	currentBlock := ctx.Header().Number.Uint64()
	if currentBlock == 0 {
		return nil
	}

	// store next epochItem (at current round endBlock)
	if db.IsEndOfCurrentRound(ctx.StateDB(), s.Address(), currentBlock) {
		if err := db.BuildNextRound(ctx.StateDB(), s.Address(), s.GetRoundSize(ctx.StateDB())); nil != err {
			return fmt.Errorf("can not build next round, %s, currentRound: %d, currentBlock: %d", err, s.GetCurrentRound(ctx.StateDB()), currentBlock)
		}
		log.Debug("Succeed BuildNextRound", "currentRound", s.GetCurrentRound(ctx.StateDB()), "currentBlock", currentBlock)
	}

	// election next epoch validators (at current epoch endBlock)
	// and store next epochItem
	if db.IsEndOfCurrentEpoch(ctx.StateDB(), s.Address(), currentBlock) {
		if err := db.BuildNextEpoch(ctx.StateDB(), s.Address(), s.GetEpochSize(ctx.StateDB()), s.GetRoundSize(ctx.StateDB())); nil != err {
			return fmt.Errorf("can not build next epoch, %s, currentEpoch: %d, currentBlock: %d", err, s.GetCurrentEpoch(ctx.StateDB()), currentBlock)
		}
		log.Debug("Succeed BuildNextEpoch", "currentEpoch", s.GetCurrentEpoch(ctx.StateDB()), "currentBlock", currentBlock)
	}
	return nil
}

// extern

func (s *StageModule) GetCurrentRound(stateDB sdk.StateDBReader) uint64 {
	return db.GetCurrentRound(stateDB, s.Address())
}
func (s *StageModule) GetCurrentEpoch(stateDB sdk.StateDBReader) uint64 {
	return db.GetCurrentEpoch(stateDB, s.Address())
}

func (s *StageModule) GetRoundByBlockNumber(stateDB sdk.StateDBReader, blockNumber uint64) uint64 {
	round, _ := db.GetRoundItemAndIndexByBlockNumber(stateDB, s.Address(), blockNumber)
	return round
}

func (s *StageModule) GetEpochByBlockNumber(stateDB sdk.StateDBReader, blockNumber uint64) uint64 {
	epoch, _ := db.GetEpochItemAndIndexByBlockNumber(stateDB, s.Address(), blockNumber)
	return epoch
}

func (s *StageModule) IsElectionBlockOnCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsElectionBlockOnCurrentRound(stateDB, s.Address(), blockNumber, s.GetRoundValidatorElectionDistance(stateDB))
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
func (s *StageModule) IsEndOfRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	_, item := db.GetRoundItemAndIndexByBlockNumber(stateDB, s.Address(), blockNumber)
	return item.EndBlock == blockNumber
}
func (s *StageModule) IsEndOfEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	_, item := db.GetEpochItemAndIndexByBlockNumber(stateDB, s.Address(), blockNumber)
	return item.EndBlock == blockNumber
}

func (s *StageModule) IsNotElectionBlockOnCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return db.IsNotElectionBlockOnCurrentRound(stateDB, s.Address(), blockNumber, s.GetRoundValidatorElectionDistance(stateDB))
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

func (s *StageModule) IsNotEndOfRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return !s.IsEndOfRound(stateDB, blockNumber)
}
func (s *StageModule) IsNotEndOfEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return !s.IsEndOfEpoch(stateDB, blockNumber)
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
	_, item := db.GetRoundItemAndIndexByBlockNumber(stateDB, s.Address(), blockNumber)
	return item.EndBlock
}

func (s *StageModule) GetRoundAndBlockBoundByBlockNumber(stateDB sdk.StateDBReader, blockNumber uint64) (uint64, uint64, uint64) {
	round, item := db.GetRoundItemAndIndexByBlockNumber(stateDB, s.Address(), blockNumber)
	return round, item.StartBlock, item.EndBlock
}

func (s *StageModule) GetEpochAndBlockBoundByBlockNumber(stateDB sdk.StateDBReader, blockNumber uint64) (uint64, uint64, uint64) {
	epoch, item := db.GetEpochItemAndIndexByBlockNumber(stateDB, s.Address(), blockNumber)
	return epoch, item.StartBlock, item.EndBlock
}

func (s *StageModule) GetRoundFlatten(stateDB sdk.StateDBReader, round uint64) (uint64, uint64) {
	item := db.GetRoundItem(stateDB, s.Address(), round)
	if nil == item {
		return 0, 0
	}
	return item.StartBlock, item.EndBlock
}

func (s *StageModule) GetEpochFlatten(stateDB sdk.StateDBReader, epoch uint64) (uint64, uint64, uint64) {
	item := db.GetEpochItem(stateDB, s.Address(), epoch)
	if nil == item {
		return 0, 0, 0
	}
	return item.StartBlock, item.EndBlock, item.RoundCount
}

func (s *StageModule) GetRoundValidatorElectionDistance(stateDB sdk.StateDBReader) uint64 {
	return db.GetRoundValidatorElectionDistance(stateDB, s.Address())
}
func (s *StageModule) GetRoundSize(stateDB sdk.StateDBReader) uint64 {
	return db.GetRoundSize(stateDB, s.Address())
}
func (s *StageModule) GetEpochSize(stateDB sdk.StateDBReader) uint64 {
	return db.GetEpochSize(stateDB, s.Address())
}
