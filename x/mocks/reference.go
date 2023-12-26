package mocks

import (
	"crypto/ecdsa"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

var (
	ErrStoreFailed  = errors.New("store failed")
	ErrRlpEncode    = errors.New("rlp encode failed")
	ErrRlpDecode    = errors.New("rlp decode failed")
	ErrNotFound     = errors.New("not found")
	ErrExist        = errors.New("already exist")
	ErrMisMatching  = errors.New("mismatching")
	ErrInvalidValue = errors.New("invalid value")
)

var (
	// for L1 module
	chainIdKey               = []byte("chainId")
	stateAddressKey          = []byte("stateAddress")
	checkpointAddressKey     = []byte("checkpointAddress")
	stakeManagerAddressKey   = []byte("stakeManagerAddress")
	depositManagerAddressKey = []byte("depositManagerAddress")
	// for stage module
	currentEpochKey    = []byte("currentEpoch") // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey    = []byte("currentRound") // "currentRound" => currentRound (It is a number)
	epochItemKeyPrefix = []byte("epochItem")    // "epochItem":epochId => {preEpoch, nextEpoch, startBlock, endBlock, roundCount}
	roundItemKeyPrefix = []byte("roundItem")
	// for stake module
	stakeWithdrawalWaitPeriodKey               = []byte("stakeWithdrawalWaitPeriod")
	delegateWithdrawalWaitPeriodKey            = []byte("delegateWithdrawalWaitPeriod")
	slashingPercentageKey                      = []byte("slashingPercentage")
	slashIncentivePercentageKey                = []byte("slashIncentivePercentage")
	maxRoundValidatorsSizeKey                  = []byte("maxRoundValidatorsSize")
	maxEpochValidatorsSizeKey                  = []byte("maxEpochValidatorsSize")
	minRoundValidatorBlockNumberKey            = []byte("minRoundValidatorBlockNumber")
	validatorNonceKey                          = []byte("validatorNonce")                    // "validatorNonce" => nonce (It is a self increasing stake index number)
	priorityValidatorHeadKey                   = []byte("priorityValidatorHead")             // "priorityValidatorHead" => priorityValidator(head)
	priorityValidatorTailKey                   = []byte("priorityValidatorTail")             // "priorityValidatorTail" => priorityValidator(tail)
	priorityValidatorKeyPrefix                 = []byte("priorityValidator")                 // "priorityValidator":shares(stakeAmount+delegataionAmount):stakeEpoch:stakeIndex => priorityValidator{preKey, nextKey, validatorAddr}
	validatorKeyPrefix                         = []byte("validator")                         // "validator":validatorAddr => validator
	validatorOwnerKeyPrefix                    = []byte("validatorOwner")                    // "validatorOwner":validatorAddr => ownerAddr
	delegationKeyPrefix                        = []byte("delegation")                        // "delegator":delegatorAddr:validatorAddr:stakeEpoch => delegation
	stakeWithdrawalQueueItemKeyPrefix          = []byte("stakeWithdrawalQueueItem")          // "stakeWithdrawalQueueItem":validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}
	delegateWithdrawalQueueItemKeyPrefix       = []byte("delegateWithdrawalQueueItem")       // "delegateWithdrawalQueueItem":delegatorAddr:validatorAddr:(unlock)epoch => {preEpoch, nextEpoch, amount}
	validatorDelegationRcKeyPrefix             = []byte("validatorDelegationRc")             // "validatorDelegationRc":validatorAddr:stakeEpoch => unStakeDelegationRcItem{preStakeEpoch, nextStakeEpoch, delegation count}
	slashProcessedKeyPrefix                    = []byte("slashProcessed")                    // "slashProcessed":exitEventId => []SlashValidatorWithdrawItem{validatorAddr, amount}
	epochValidatorSharesSnapshotQueueKeyPrefix = []byte("epochValidatorSharesSnapshotQueue") // "epochValidatorSharesSnapshotQueue":epochId => []validatorSharesSnapshot  (For settlement epoch)
	roundValidatorSharesSnapshotQueueKeyPrefix = []byte("roundValidatorSharesSnapshotQueue") // "roundValidatorSharesSnapshotQueue":roundId => []validatorSharesSnapshot  (For consensus round)
	numberOfBlocksForRoundValidatorKeyPrefix   = []byte("numberOfBlocksForRoundValidator")   // "numberOfBlocksForRoundValidator":validatorAddr:round => numberOfBlocks
	//for reward module
	paidRewardPerEpochKeyPrefix                = []byte("paidRewardPerEpoch")                // "paidRewardPerEpoch":epochId => paidRewards
	pendingValidatorRewardKeyPrefix            = []byte("pendingValidatorReward")            // "pendingValidatorReward":validatorAddr => pendingRewards
	pendingDelegatorRewardKeyPrefix            = []byte("pendingDelegatorReward")            // "pendingDelegatorReward":delegatorAddr:validatorAddr => pendingRewards
	epochDelegationRewardPerShareItemKeyPrefix = []byte("epochDelegationRewardPerShareItem") // "epochDelegationRewardPerShareItem":validatorAddr:stakeEpoch:rewardEpoch => epochRewardPerDelegationShareItem{ preRewardEpoch, nextRewardEpoch, totalReward, perShareReward}
	// for vrf module
	nonceAndProofKey = []byte("nonceAndProof") // "nonce":blockNumber => nonceAndProof
)

type MockL1Module struct {
	addr    common.Address
	statedb *mock.MockStateDB
}

func NewMockMockL1Moduler(statedb *mock.MockStateDB) *MockL1Module {
	return &MockL1Module{
		addr:    common.HexToAddress("0xeeeeeeeeeeeeeeeee1111fffffffffffffffffff"),
		statedb: statedb,
	}
}
func (l1 *MockL1Module) MockChainID(chainId *big.Int) {
	l1.statedb.SetState(l1.addr, chainIdKey, chainId.Bytes())
}
func (l1 *MockL1Module) MockStateAddress(stateAddr common.Address) {
	l1.statedb.SetState(l1.addr, stateAddressKey, stateAddr.Bytes())
}
func (l1 *MockL1Module) MockCheckpointAddress(checkpointAddr common.Address) {
	l1.statedb.SetState(l1.addr, checkpointAddressKey, checkpointAddr.Bytes())
}
func (l1 *MockL1Module) MockStakeManagerAddress(stakeManagerAddr common.Address) {
	l1.statedb.SetState(l1.addr, stakeManagerAddressKey, stakeManagerAddr.Bytes())
}
func (l1 *MockL1Module) MockDepositManagerAddress(depositManagerAddr common.Address) {
	l1.statedb.SetState(l1.addr, depositManagerAddressKey, depositManagerAddr.Bytes())
}
func (l1 *MockL1Module) GetChainID() (*big.Int, error) {
	return new(big.Int).SetBytes(l1.statedb.GetState(l1.addr, chainIdKey)), nil
}
func (l1 *MockL1Module) GetStateAddress() (common.Address, error) {
	value := l1.statedb.GetState(l1.addr, stateAddressKey)
	if len(value) == 0 {
		return common.ZeroAddr, ErrNotFound
	}
	return common.BytesToAddress(value), nil
}
func (l1 *MockL1Module) GetCheckpointAddress() (common.Address, error) {
	value := l1.statedb.GetState(l1.addr, checkpointAddressKey)
	if len(value) == 0 {
		return common.ZeroAddr, ErrNotFound
	}
	return common.BytesToAddress(value), nil
}
func (l1 *MockL1Module) GetStakeManagerAddress() (common.Address, error) {
	value := l1.statedb.GetState(l1.addr, stakeManagerAddressKey)
	if len(value) == 0 {
		return common.ZeroAddr, ErrNotFound
	}
	return common.BytesToAddress(value), nil
}
func (l1 *MockL1Module) GetDepositManagerAddress() (common.Address, error) {
	value := l1.statedb.GetState(l1.addr, depositManagerAddressKey)
	if len(value) == 0 {
		return common.ZeroAddr, ErrNotFound
	}
	return common.BytesToAddress(value), nil
}

type MockStageModule struct {
	addr    common.Address
	statedb *mock.MockStateDB
}

func NewMockStageModuler(statedb *mock.MockStateDB) *MockStageModule {
	return &MockStageModule{
		addr:    constants.StageManagerAddress,
		statedb: statedb,
	}
}

func (stage *MockStageModule) GetCurrentRound(stateDB sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stage.statedb.GetState(stage.addr, currentEpochKey))
}
func (stage *MockStageModule) GetCurrentEpoch(stateDB sdk.StateDBReader) uint64 { return 1 }
func (stage *MockStageModule) GetRoundByBlockNumber(stateDB sdk.StateDBReader, blockNumber uint64) uint64 {
	return 1
}
func (stage *MockStageModule) GetEpochByBlockNumber(stateDB sdk.StateDBReader, blockNumber uint64) uint64 {
	return 1
}
func (stage *MockStageModule) IsElectionBlockOnCurrentRound(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsElectionBlockOnCurrentEpoch(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsBeginOfNextRound(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsBeginOfNextEpoch(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsEndOfRound(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsEndOfEpoch(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotElectionBlockOnCurrentRound(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotElectionBlockOnCurrentEpoch(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotBeginOfNextRound(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotBeginOfNextEpoch(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotEndOfRound(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) IsNotEndOfEpoch(db sdk.StateDBReader, blockNumber uint64) bool {
	return true
}
func (stage *MockStageModule) BlocksOfRound(stateDB sdk.StateDBReader, round uint64) uint64 {
	return 1
}
func (stage *MockStageModule) BlocksOfEpoch(stateDB sdk.StateDBReader, epoch uint64) uint64 {
	return 1
}
func (stage *MockStageModule) GetLastNumber(stateDB sdk.StateDBReader, blockNumber uint64) uint64 {
	return 1
}
func (stage *MockStageModule) GetRoundAndBlockBoundByBlockNumber(db sdk.StateDBReader, blockNumber uint64) (uint64, uint64, uint64) {
	return 1, 1, 1
}
func (stage *MockStageModule) GetEpochAndBlockBoundByBlockNumber(db sdk.StateDBReader, blockNumber uint64) (uint64, uint64, uint64) {
	return 1, 1, 1
}

type MockStakeModule struct {
	addr         common.Address
	statedb      *mock.MockStateDB
	l1Module     *MockL1Module
	stageModule  *MockStageModule
	rewardModule *MockRewardModule
	vrfModule    *MockVRFModule
}

func NewMockStakeModuler(statedb *mock.MockStateDB, l1Module *MockL1Module, stageModule *MockStageModule) *MockStakeModule {
	return &MockStakeModule{
		addr:        constants.StakeHandlerAddress,
		statedb:     statedb,
		l1Module:    l1Module,
		stageModule: stageModule,
	}
}
func (stake *MockStakeModule) SetRewardModule(rewardModule *MockRewardModule) {
	stake.rewardModule = rewardModule
}
func (stake *MockStakeModule) SetVRFModule(vrfModule *MockVRFModule) {
	stake.vrfModule = vrfModule
}

func (stake *MockStakeModule) GetStakeWithdrawalWaitPeriod(stateDB sdk.StateDBReader) uint64 {
	return 1
}
func (stake *MockStakeModule) GetDelegateWithdrawalWaitPeriod(stateDB sdk.StateDBReader) uint64 {
	return 1
}
func (stake *MockStakeModule) GetSlashingPercentage(stateDB sdk.StateDBReader) uint64 { return 1 }
func (stake *MockStakeModule) GetSlashIncentivePercentage(stateDB sdk.StateDBReader) uint64 {
	return 1
}
func (stake *MockStakeModule) GetMaxRoundValidatorsSize(stateDB sdk.StateDBReader) uint64 { return 1 }
func (stake *MockStakeModule) GetMaxEpochValidatorsSize(stateDB sdk.StateDBReader) uint64 { return 1 }
func (stake *MockStakeModule) GetMinRoundValidatorBlockNumber(stateDB sdk.StateDBReader) uint64 {
	return 1
}
func (stake *MockStakeModule) GetRoundValidatorIds(stateDB sdk.StateDBReader, round uint64) []common.Address {
	return nil
}
func (stake *MockStakeModule) GetEpochValidatorIds(stateDB sdk.StateDBReader, epoch uint64) []common.Address {
	return nil
}
func (stake *MockStakeModule) GetEpochValidatorSnapQueueFlatten(stateDB sdk.StateDBReader, epoch uint64) ([]common.Address, []*big.Int, []*big.Int, []uint64, []uint64, []uint64, []uint64) {
	return nil, nil, nil, nil, nil, nil, nil
}
func (stake *MockStakeModule) IsValidValidator(stateDB sdk.StateDBReader, validatorAddr common.Address) bool {
	return true
}
func (stake *MockStakeModule) IsInvalidValidator(stateDB sdk.StateDBReader, validatorAddr common.Address) bool {
	return true
}
func (stake *MockStakeModule) IsEmptyValidator(stateDB sdk.StateDBReader, validatorAddr common.Address) bool {
	return true
}
func (stake *MockStakeModule) GetValidatorCommissionRate(stateDB sdk.StateDBReader, validatorAddr common.Address) uint64 {
	return 1
}
func (stake *MockStakeModule) GetValidatorStakeEpoch(stateDB sdk.StateDBReader, validatorAddr common.Address) uint64 {
	return 1
}
func (stake *MockStakeModule) GetValidatorStakeAmount(stateDB sdk.StateDBReader, validatorAddr common.Address) *big.Int {
	return common.Big0
}
func (stake *MockStakeModule) GetValidatorDelegateAmount(stateDB sdk.StateDBReader, validatorAddr common.Address) *big.Int {
	return common.Big0
}
func (stake *MockStakeModule) GetValidatorOwner(stateDB sdk.StateDBReader, validatorAddr common.Address) common.Address {
	return common.ZeroAddr
}
func (stake *MockStakeModule) GetNumberOfBlocksForRoundValidator(stateDB sdk.StateDBReader, validatorAddr common.Address, round uint64) uint64 {
	return 1
}
func (stake *MockStakeModule) GetEpochByValidatorDelegationRcPending(stateDB sdk.StateDBReader, validatorAddr common.Address) []uint64 {
	return nil
}
func (stake *MockStakeModule) GetDelegationFlatten(stateDB sdk.StateDBReader, delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) (uint64, *big.Int) {
	return 1, common.Big0
}
func (stake *MockStakeModule) UpdateDelegationEpoch(stateDB sdk.StateDB, delegatorAddr, validatorAddr common.Address, stakeEpoch, delegateEpoch uint64) error {
	return nil
}
func (stake *MockStakeModule) GetValidatorECDSAPubKey(stateDB sdk.StateDBReader, validatorAddr common.Address) *ecdsa.PublicKey {
	return nil
}

type MockRewardModule struct {
	addr        common.Address
	statedb     *mock.MockStateDB
	stageModule *MockStageModule
	stakeMdoule *MockStakeModule
}

func NewMockRewardModuler(statedb *mock.MockStateDB, stageModule *MockStageModule, stakeMdoule *MockStakeModule) *MockRewardModule {
	return &MockRewardModule{
		addr:        constants.RewardManagerAddress,
		statedb:     statedb,
		stageModule: stageModule,
		stakeMdoule: stakeMdoule,
	}
}

func (reward *MockRewardModule) UpdateDelegationRewardsByStakeEpoch(stateDB sdk.StateDB, delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) error {
	return nil
}
func (reward *MockRewardModule) UpdateDelegationRewards(stateDB sdk.StateDB, delegatorAddr, validatorAddr common.Address) error {
	return nil
}
func (reward *MockRewardModule) GetRewardPerBlock(stateDB sdk.StateDBReader) *big.Int {
	return common.Big0
}
func (reward *MockRewardModule) GetRewardPerEpoch(stateDB sdk.StateDBReader) *big.Int {
	return common.Big0
}

type MockVRFModule struct {
	addr        common.Address
	statedb     *mock.MockStateDB
	stageModule *MockStageModule
	stakeMdoule *MockStakeModule
}

func NewMockVRFModuler(statedb *mock.MockStateDB, stageModule *MockStageModule, stakeMdoule *MockStakeModule) *MockVRFModule {
	return &MockVRFModule{
		addr:        constants.VRFManagerAddress,
		statedb:     statedb,
		stageModule: stageModule,
		stakeMdoule: stakeMdoule,
	}
}

func (vrf *MockVRFModule) GetNonceQueueFromTail(stateDB sdk.StateDBReader, blockNumber, size uint64) ([]common.Hash, error) {
	return nil, nil
}
func (vrf *MockVRFModule) GetCurrentNonce(stateDB sdk.StateDBReader, blockNumber uint64) (common.Hash, error) {
	return common.ZeroHash, nil
}
