package mocks

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
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
	// for stage
	roundSizeKey       = []byte("roundSize")
	epochSizeKey       = []byte("epochSize")
	currentEpochKey    = []byte("currentEpoch") // "currentEpoch" => currentEpoch (It is a number)
	currentRoundKey    = []byte("currentRound") // "currentRound" => currentRound (It is a number)
	epochItemKeyPrefix = []byte("epochItem")    // "epochItem":epochId => {preEpoch, nextEpoch, startBlock, endBlock, epochCount}
	roundItemKeyPrefix = []byte("roundItem")    // "roundItem":roundId => {preRound, nextRound, startBlock, endBlock}
	// for stage module mock
	blockRoundKeyPrefix              = []byte("blockRound")              // "blockRound":block => round
	blockEpochKeyPrefix              = []byte("blockEpoch")              // "blockEpoch":block => epoch
	blockElectionOfRoundKeyPrefix    = []byte("blockElectionOfRound")    // "blockElectionOfRound":block => round
	blockElectionOfEpochKeyPrefix    = []byte("blockElectionOfEpoch")    // "blockElectionOfEpoch":block => epoch
	blockEndOfRoundKeyPrefix         = []byte("blockEndOfRound")         // "blockEndOfRound":block => round
	blockEndOfEpochKeyPrefix         = []byte("blockEndOfEpoch")         // "blockEndOfEpoch":block => epoch
	blockBeginOfRoundKeyPrefix       = []byte("blockBeginOfRound")       // "blockBeginOfRound":block => round
	blockBeginOfEpochKeyPrefix       = []byte("blockBeginOfEpoch")       // "blockBeginOfEpoch":block => epoch
	blockBeginOfNextRoundKeyPrefix   = []byte("blockBeginOfNextRound")   // "blockBeginOfNextRound":block => round
	blockBeginOfNextEpochKeyPrefix   = []byte("blockBeginOfNextEpoch")   // "blockBeginOfNextEpoch":block => epoch
	blocksOfRoundKeyPrefix           = []byte("blocksOfRound")           // "blocksOfRound":round => numbers of block
	blocksOfEpochKeyPrefix           = []byte("blocksOfEpoch")           // "blocksOfEpoch":epoch => numbers of block
	lastBlockOfRoundKeyPrefix        = []byte("lastBlockOfRound")        // "lastBlockOfRound":block => lastBlock
	lastBlockOfEpochKeyPrefix        = []byte("lastBlockOfEpoch")        // "lastBlockOfEpoch":block => lastBlock
	roundAndBlockBoundBeginKeyPrefix = []byte("roundAndBlockBoundBegin") // "roundAndBlockBoundBegin":block = > begin block
	roundAndBlockBoundEndKeyPrefix   = []byte("roundAndBlockBoundEnd")   // "roundAndBlockBoundEnd":block = > begin block
	epochAndBlockBoundBeginKeyPrefix = []byte("epochAndBlockBoundBegin") // "epochAndBlockBoundBegin":block = > begin block
	epochAndBlockBoundEndKeyPrefix   = []byte("epochAndBlockBoundEnd")   // "epochAndBlockBoundEnd":block = > begin block
	// for stake module
	stakeWithdrawalWaitPeriodKey               = []byte("stakeWithdrawalWaitPeriod")
	delegateWithdrawalWaitPeriodKey            = []byte("delegateWithdrawalWaitPeriod")
	slashingPercentageKey                      = []byte("slashingPercentage")
	slashIncentivePercentageKey                = []byte("slashIncentivePercentage")
	maxRoundValidatorsSizeKey                  = []byte("maxRoundValidatorsSize")
	maxEpochValidatorsSizeKey                  = []byte("maxEpochValidatorsSize")
	minBlocksOfRoundValidatorKey               = []byte("minBlocksOfRoundValidator")
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
	rewardPerBlockKey                          = []byte("rewardPerBlock")
	rewardPerEpochKey                          = []byte("rewardPerEpoch")
	paidRewardPerEpochKeyPrefix                = []byte("paidRewardPerEpoch")                // "paidRewardPerEpoch":epochId => paidRewards
	pendingValidatorRewardKeyPrefix            = []byte("pendingValidatorReward")            // "pendingValidatorReward":validatorAddr => pendingRewards
	pendingDelegatorRewardKeyPrefix            = []byte("pendingDelegatorReward")            // "pendingDelegatorReward":delegatorAddr:validatorAddr => pendingRewards
	epochDelegationRewardPerShareItemKeyPrefix = []byte("epochDelegationRewardPerShareItem") // "epochDelegationRewardPerShareItem":validatorAddr:stakeEpoch:rewardEpoch => epochRewardPerDelegationShareItem{ preRewardEpoch, nextRewardEpoch, totalReward, perShareReward}
	// for stake module mock
	validatorDelegationRcForMockKeyPrefix                = []byte("validatorDelegationRcForMock")                // "validatorDelegationRcForMock"validatorAddr:stakeEpoch => numberOfRcs
	validatorDelegationRcStakeEpochQueueForMockKeyPrefix = []byte("validatorDelegationRcStakeEpochQueueForMock") // "validatorDelegationRcStakeEpochQueueForMock":validatorAddr => []uint64{stakeEpoch,...,stakeEpoch}
	// for vrf module
	nonceAndProofKey = []byte("nonceAndProof") // "nonce":blockNumber => nonceAndProof
	// for vrf module mock
	nonceOfVRFForMockKeyPrefix = []byte("nonceOfVRFForMock") // "nonceOfVRFForMock":blockNumber => nonce
)

func encodeEpochItemKey(epoch uint64) []byte {
	return append(epochItemKeyPrefix, common.Uint64ToBytes(epoch)...)
}

func encodeRoundItemKey(round uint64) []byte {
	return append(roundItemKeyPrefix, common.Uint64ToBytes(round)...)
}

// for stage module mock
func encodeBlockRoundKey(blockNumber uint64) []byte {
	return append(blockRoundKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockEpochKey(blockNumber uint64) []byte {
	return append(blockEpochKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockElectionOfRoundKey(blockNumber uint64) []byte {
	return append(blockElectionOfRoundKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockElectionOfEpochKey(blockNumber uint64) []byte {
	return append(blockElectionOfEpochKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockEndOfRoundKey(blockNumber uint64) []byte {
	return append(blockEndOfRoundKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockEndOfEpochKey(blockNumber uint64) []byte {
	return append(blockEndOfEpochKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockBeginOfRoundKey(blockNumber uint64) []byte {
	return append(blockBeginOfRoundKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockBeginOfEpochKey(blockNumber uint64) []byte {
	return append(blockBeginOfEpochKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockBeginOfNextRoundKey(blockNumber uint64) []byte {
	return append(blockBeginOfNextRoundKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlockBeginOfNextEpochKey(blockNumber uint64) []byte {
	return append(blockBeginOfNextEpochKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeBlocksOfRoundKey(round uint64) []byte {
	return append(blocksOfRoundKeyPrefix, common.Uint64ToBytes(round)...)
}
func encodeBlocksOfEpochKey(epoch uint64) []byte {
	return append(blocksOfEpochKeyPrefix, common.Uint64ToBytes(epoch)...)
}
func encodeLastBlockOfRoundKey(blockNumber uint64) []byte {
	return append(lastBlockOfRoundKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeLastBlockOfEpochKey(blockNumber uint64) []byte {
	return append(lastBlockOfEpochKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeRoundAndBlockBoundBeginKey(blockNumber uint64) []byte {
	return append(roundAndBlockBoundBeginKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeEpochAndBlockBoundBeginKey(blockNumber uint64) []byte {
	return append(epochAndBlockBoundBeginKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeRoundAndBlockBoundEndKey(blockNumber uint64) []byte {
	return append(roundAndBlockBoundEndKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}
func encodeEpochAndBlockBoundEndKey(blockNumber uint64) []byte {
	return append(epochAndBlockBoundEndKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}

// for stake module mock
func encodeValidatorKey(validatorAddr common.Address) []byte {
	return append(validatorKeyPrefix, validatorAddr.Bytes()...)

}
func encodeRoundValidatorSharesSnapshotQueueKey(round uint64) []byte {
	return append(roundValidatorSharesSnapshotQueueKeyPrefix, common.Uint64ToBytes(round)...)
}
func encodeEpochValidatorSharesSnapshotQueueKey(epoch uint64) []byte {
	return append(epochValidatorSharesSnapshotQueueKeyPrefix, common.Uint64ToBytes(epoch)...)
}
func encodeNumberOfBlocksForRoundValidatorKey(validatorAddr common.Address, round uint64) []byte {
	return append(append(numberOfBlocksForRoundValidatorKeyPrefix, validatorAddr.Bytes()...), common.Uint64ToBytes(round)...)
}
func encodeDelegatorKey(delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) []byte {
	delegatorAddrBytes := delegatorAddr.Bytes()
	validatorAddrBytes := validatorAddr.Bytes()
	stakeEpochBytes := common.Uint64ToBytes(stakeEpoch)

	keyPrefixSize := len(delegationKeyPrefix)
	appendDelegatorSize := keyPrefixSize + len(delegatorAddrBytes)
	appendVlidatorAddrSize := appendDelegatorSize + len(validatorAddrBytes)
	size := appendVlidatorAddrSize + len(stakeEpochBytes)

	key := make([]byte, size)

	copy(key[:keyPrefixSize], delegationKeyPrefix)
	copy(key[keyPrefixSize:appendDelegatorSize], delegatorAddrBytes)
	copy(key[appendDelegatorSize:appendVlidatorAddrSize], validatorAddrBytes)
	copy(key[appendVlidatorAddrSize:], stakeEpochBytes)

	return key
}

func encodeValidatorDelegationRcForMockKey(validatorAddr common.Address, stakeEpoch uint64) []byte {
	return append(append(validatorDelegationRcForMockKeyPrefix, validatorAddr.Bytes()...), common.Uint64ToBytes(stakeEpoch)...)
}
func encodeValidatorDelegationRcStakeEpochQueueForMockKey(validatorAddr common.Address) []byte {
	return append(validatorDelegationRcStakeEpochQueueForMockKeyPrefix, validatorAddr.Bytes()...)
}

// for reward module

// for vrf module
func encodeNonceOfVRFForMockKey(blockNumber uint64) []byte {
	return append(nonceOfVRFForMockKeyPrefix, common.Uint64ToBytes(blockNumber)...)
}

// -----------------------------------------------
type MockL1Module struct {
	addr    common.Address
	statedb sdk.StateDB
}

func NewMockMockL1Module(statedb sdk.StateDB) *MockL1Module {
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
	statedb sdk.StateDB
}

func NewMockStageModule(statedb sdk.StateDB) *MockStageModule {
	return &MockStageModule{
		addr:    constants.StageManagerAddress,
		statedb: statedb,
	}
}

func (stage *MockStageModule) MockRoundSize(size uint64) {
	stage.statedb.SetState(stage.addr, roundSizeKey, common.Uint64ToBytes(size))
}
func (stage *MockStageModule) MockEpochSize(size uint64) {
	stage.statedb.SetState(stage.addr, epochSizeKey, common.Uint64ToBytes(size))
}

func (stage *MockStageModule) MockCurrentRound(round uint64) {
	stage.statedb.SetState(stage.addr, currentRoundKey, common.Uint64ToBytes(round))
}
func (stage *MockStageModule) MockCurrentEpoch(epoch uint64) {
	stage.statedb.SetState(stage.addr, currentEpochKey, common.Uint64ToBytes(epoch))
}

func (stage *MockStageModule) MockRound(round uint64, item []byte) {
	stage.statedb.SetState(stage.addr, encodeRoundItemKey(round), item)
}

func (stage *MockStageModule) MockEpoch(epoch uint64, item []byte) {
	stage.statedb.SetState(stage.addr, encodeEpochItemKey(epoch), item)
}

func (stage *MockStageModule) MockBlockRound(blockNumber, round uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockRoundKey(blockNumber), common.Uint64ToBytes(round))
}
func (stage *MockStageModule) MockBlockEpoch(blockNumber, epoch uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockEpochKey(blockNumber), common.Uint64ToBytes(epoch))
}

func (stage *MockStageModule) MockBlockElectionOfRound(blockNumber, round uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockElectionOfRoundKey(blockNumber), common.Uint64ToBytes(round))
}
func (stage *MockStageModule) MockBlockElectionOfEpoch(blockNumber, epoch uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockElectionOfEpochKey(blockNumber), common.Uint64ToBytes(epoch))
}

func (stage *MockStageModule) MockBlockEndOfRound(blockNumber, round uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockEndOfRoundKey(blockNumber), common.Uint64ToBytes(round))
}
func (stage *MockStageModule) MockBlockEndOfEpoch(blockNumber, epoch uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockEndOfEpochKey(blockNumber), common.Uint64ToBytes(epoch))
}

func (stage *MockStageModule) MockBlockBeginOfRound(blockNumber, round uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockBeginOfRoundKey(blockNumber), common.Uint64ToBytes(round))
}
func (stage *MockStageModule) MockBlockBeginOfEpoch(blockNumber, epoch uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockBeginOfEpochKey(blockNumber), common.Uint64ToBytes(epoch))
}

func (stage *MockStageModule) MockBlockBeginOfNextRound(blockNumber, round uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockBeginOfNextRoundKey(blockNumber), common.Uint64ToBytes(round))
}
func (stage *MockStageModule) MockBlockBeginOfNextEpoch(blockNumber, epoch uint64) {
	stage.statedb.SetState(stage.addr, encodeBlockBeginOfNextEpochKey(blockNumber), common.Uint64ToBytes(epoch))
}

func (stage *MockStageModule) MockBlocksOfRound(blockNumber, blocks uint64) {
	stage.statedb.SetState(stage.addr, encodeBlocksOfRoundKey(blockNumber), common.Uint64ToBytes(blocks))
}
func (stage *MockStageModule) MockBlocksOfEpoch(blockNumber, blocks uint64) {
	stage.statedb.SetState(stage.addr, encodeBlocksOfEpochKey(blockNumber), common.Uint64ToBytes(blocks))
}

func (stage *MockStageModule) MockLastBlockOfRound(blockNumber, lastBlock uint64) {
	stage.statedb.SetState(stage.addr, encodeLastBlockOfRoundKey(blockNumber), common.Uint64ToBytes(lastBlock))
}
func (stage *MockStageModule) MockLastBlockOfEpoch(blockNumber, lastBlock uint64) {
	stage.statedb.SetState(stage.addr, encodeLastBlockOfEpochKey(blockNumber), common.Uint64ToBytes(lastBlock))
}

func (stage *MockStageModule) MockRoundAndBlockBoundBegin(blockNumber, beginBlock uint64) {
	stage.statedb.SetState(stage.addr, encodeRoundAndBlockBoundBeginKey(blockNumber), common.Uint64ToBytes(beginBlock))
}
func (stage *MockStageModule) MockEpochAndBlockBoundBegin(blockNumber, beginBlock uint64) {
	stage.statedb.SetState(stage.addr, encodeEpochAndBlockBoundBeginKey(blockNumber), common.Uint64ToBytes(beginBlock))
}

func (stage *MockStageModule) MockRoundAndBlockBoundEnd(blockNumber, endBlock uint64) {
	stage.statedb.SetState(stage.addr, encodeRoundAndBlockBoundEndKey(blockNumber), common.Uint64ToBytes(endBlock))
}
func (stage *MockStageModule) MockEpochAndBlockBoundEnd(blockNumber, endBlock uint64) {
	stage.statedb.SetState(stage.addr, encodeEpochAndBlockBoundEndKey(blockNumber), common.Uint64ToBytes(endBlock))
}

func (stage *MockStageModule) GetRoundSize(stateDB sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stage.statedb.GetState(stage.addr, roundSizeKey))
}
func (stage *MockStageModule) GetEpochSize(stateDB sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stage.statedb.GetState(stage.addr, epochSizeKey))
}

func (stage *MockStageModule) GetCurrentRound(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stage.statedb.GetState(stage.addr, currentRoundKey))
}
func (stage *MockStageModule) GetCurrentEpoch(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stage.statedb.GetState(stage.addr, currentEpochKey))
}
func (stage *MockStageModule) GetRoundByBlockNumber(statedb sdk.StateDBReader, blockNumber uint64) uint64 {
	// append(blockRoundKeyPrefix, common.Uint64ToBytes(blockNumber)...)
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockRoundKey(blockNumber)))
}
func (stage *MockStageModule) GetEpochByBlockNumber(statedb sdk.StateDBReader, blockNumber uint64) uint64 {
	// append(blockEpochKeyPrefix, common.Uint64ToBytes(blockNumber)...)
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockEpochKey(blockNumber)))
}
func (stage *MockStageModule) IsElectionBlockOnCurrentRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockElectionOfRoundKey(blockNumber))) == stage.GetCurrentRound(statedb)
}
func (stage *MockStageModule) IsElectionBlockOnCurrentEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockElectionOfEpochKey(blockNumber))) == stage.GetCurrentEpoch(statedb)
}
func (stage *MockStageModule) IsBeginOfCurrentRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockBeginOfRoundKey(blockNumber))) == stage.GetCurrentRound(statedb)
}
func (stage *MockStageModule) IsBeginOfCurrentEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockBeginOfEpochKey(blockNumber))) == stage.GetCurrentEpoch(statedb)
}
func (stage *MockStageModule) IsBeginOfNextRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockBeginOfNextRoundKey(blockNumber))) == stage.GetCurrentRound(statedb)+1
}
func (stage *MockStageModule) IsBeginOfNextEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockBeginOfNextEpochKey(blockNumber))) == stage.GetCurrentEpoch(statedb)+1
}
func (stage *MockStageModule) IsEndOfCurrentRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockEndOfRoundKey(blockNumber))) == stage.GetCurrentRound(statedb)
}
func (stage *MockStageModule) IsEndOfCurrentEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockEndOfEpochKey(blockNumber))) == stage.GetCurrentEpoch(statedb)
}
func (stage *MockStageModule) IsEndOfRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockEndOfRoundKey(blockNumber))) != 0
}
func (stage *MockStageModule) IsEndOfEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockEndOfEpochKey(blockNumber))) != 0
}
func (stage *MockStageModule) IsNotElectionBlockOnCurrentRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsElectionBlockOnCurrentRound(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotElectionBlockOnCurrentEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsElectionBlockOnCurrentEpoch(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotBeginOfCurrentRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsBeginOfCurrentRound(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotBeginOfCurrentEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsBeginOfCurrentEpoch(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotBeginOfNextRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsBeginOfNextRound(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotBeginOfNextEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsBeginOfNextEpoch(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotEndOfCurrentRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsEndOfCurrentRound(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotEndOfCurrentEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsEndOfCurrentEpoch(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotEndOfRound(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsEndOfRound(statedb, blockNumber)
}
func (stage *MockStageModule) IsNotEndOfEpoch(statedb sdk.StateDBReader, blockNumber uint64) bool {
	return !stage.IsEndOfEpoch(statedb, blockNumber)
}
func (stage *MockStageModule) BlocksOfRound(statedb sdk.StateDBReader, round uint64) uint64 {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlocksOfRoundKey(round)))
}
func (stage *MockStageModule) BlocksOfEpoch(statedb sdk.StateDBReader, epoch uint64) uint64 {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlocksOfEpochKey(epoch)))
}
func (stage *MockStageModule) GetLastNumber(statedb sdk.StateDBReader, blockNumber uint64) uint64 {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeLastBlockOfRoundKey(blockNumber)))
}
func (stage *MockStageModule) GetRoundAndBlockBoundByBlockNumber(statedb sdk.StateDBReader, blockNumber uint64) (uint64, uint64, uint64) {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockRoundKey(blockNumber))),
		common.BytesToUint64(statedb.GetState(stage.addr, encodeRoundAndBlockBoundBeginKey(blockNumber))),
		common.BytesToUint64(statedb.GetState(stage.addr, encodeRoundAndBlockBoundEndKey(blockNumber)))
}
func (stage *MockStageModule) GetEpochAndBlockBoundByBlockNumber(statedb sdk.StateDBReader, blockNumber uint64) (uint64, uint64, uint64) {
	return common.BytesToUint64(statedb.GetState(stage.addr, encodeBlockEpochKey(blockNumber))),
		common.BytesToUint64(statedb.GetState(stage.addr, encodeEpochAndBlockBoundBeginKey(blockNumber))),
		common.BytesToUint64(statedb.GetState(stage.addr, encodeEpochAndBlockBoundEndKey(blockNumber)))
}

func (stage *MockStageModule) GetRoundFlatten(db sdk.StateDBReader, round uint64) (uint64, uint64) {

	value := db.GetState(stage.addr, encodeRoundItemKey(round))
	if len(value) == 0 {
		return 0, 0
	}

	var item types.RoundItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return (&item).StartBlock, (&item).EndBlock
	}

	return 0, 0
}

func (stage *MockStageModule) GetEpochFlatten(db sdk.StateDBReader, epoch uint64) (uint64, uint64, uint64) {

	value := db.GetState(stage.addr, encodeEpochItemKey(epoch))
	if len(value) == 0 {
		return 0, 0, 0
	}

	var item types.EpochItem
	if err := rlp.DecodeBytes(value, &item); nil == err {
		return (&item).StartBlock, (&item).EndBlock, (&item).RoundCount
	}

	return 0, 0, 0
}

type MockStakeModule struct {
	addr         common.Address
	statedb      sdk.StateDB
	l1Module     *MockL1Module
	stageModule  *MockStageModule
	rewardModule *MockRewardModule
	vrfModule    *MockVRFModule
}

func NewMockStakeModule(statedb sdk.StateDB, l1Module *MockL1Module, stageModule *MockStageModule) *MockStakeModule {
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

func (stake *MockStakeModule) MockStakeWithdrawalWaitPeriod(period uint64) {
	stake.statedb.SetState(stake.addr, stakeWithdrawalWaitPeriodKey, common.Uint64ToBytes(period))
}
func (stake *MockStakeModule) MockDelegateWithdrawalWaitPeriod(period uint64) {
	stake.statedb.SetState(stake.addr, delegateWithdrawalWaitPeriodKey, common.Uint64ToBytes(period))
}
func (stake *MockStakeModule) MockSlashingPercentage(percentage uint64) {
	stake.statedb.SetState(stake.addr, slashingPercentageKey, common.Uint64ToBytes(percentage))
}
func (stake *MockStakeModule) MockSlashIncentivePercentage(percentage uint64) {
	stake.statedb.SetState(stake.addr, slashIncentivePercentageKey, common.Uint64ToBytes(percentage))
}
func (stake *MockStakeModule) MockMaxRoundValidatorsSize(size uint64) {
	stake.statedb.SetState(stake.addr, maxRoundValidatorsSizeKey, common.Uint64ToBytes(size))
}
func (stake *MockStakeModule) MockMaxEpochValidatorsSize(size uint64) {
	stake.statedb.SetState(stake.addr, maxEpochValidatorsSizeKey, common.Uint64ToBytes(size))
}
func (stake *MockStakeModule) MockMinBlocksOfRoundValidator(blocks uint64) {
	stake.statedb.SetState(stake.addr, minBlocksOfRoundValidatorKey, common.Uint64ToBytes(blocks))
}

func (stake *MockStakeModule) MockInitValidatorGenesisPriority() error {
	head := NewMockPriorityValidator(
		priorityValidatorTailKey,
		priorityValidatorTailKey,
		common.ZeroAddr,
	)
	tail := NewMockPriorityValidator(
		priorityValidatorHeadKey,
		priorityValidatorHeadKey,
		common.ZeroAddr,
	)
	hvalue, err := rlp.EncodeToBytes(head)
	if nil != err {
		return fmt.Errorf("rlp encode head validator priority %s", err)
	}
	tvalue, err := rlp.EncodeToBytes(tail)
	if nil != err {
		return fmt.Errorf("rlp encode tail validator priority %s", err)
	}

	stake.statedb.SetState(stake.addr, priorityValidatorHeadKey, hvalue)
	stake.statedb.SetState(stake.addr, priorityValidatorTailKey, tvalue)
	return nil
}

func (stake *MockStakeModule) MockValidator(validatorAddr common.Address, validator *MockValidatorSnapshot) error {

	value, err := rlp.EncodeToBytes(validator)
	if nil != err {
		return err
	}
	stake.statedb.SetState(stake.addr, encodeValidatorKey(validatorAddr), value)
	return nil
}
func (stake *MockStakeModule) MockRoundValidatorSnapQueue(round uint64, queue MockValidatorSnapshotQueue) error {

	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return err
	}
	stake.statedb.SetState(stake.addr, encodeRoundValidatorSharesSnapshotQueueKey(round), value)
	return nil
}
func (stake *MockStakeModule) MockEpochValidatorSnapQueue(epoch uint64, queue MockValidatorSnapshotQueue) error {

	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return err
	}
	stake.statedb.SetState(stake.addr, encodeEpochValidatorSharesSnapshotQueueKey(epoch), value)
	return nil
}
func (stake *MockStakeModule) MockNumberOfBlocksForRoundValidator(validatorAddr common.Address, round, numberOfBlocks uint64) {
	stake.statedb.SetState(stake.addr, encodeNumberOfBlocksForRoundValidatorKey(validatorAddr, round), common.Uint64ToBytes(numberOfBlocks))
}
func (stake *MockStakeModule) MockDelegation(delegatorAddr, validatorAddr common.Address, stakeEpoch uint64, delegation *MockDelegation) error {

	value, err := rlp.EncodeToBytes(delegation)
	if nil != err {
		return err
	}
	stake.statedb.SetState(stake.addr, encodeDelegatorKey(delegatorAddr, validatorAddr, stakeEpoch), value)
	return nil
}

func (stake *MockStakeModule) MockValidatorDelegationRc(validatorAddr common.Address, stakeEpoch, numberOfRcs uint64) error {

	var queue []uint64
	value := stake.statedb.GetState(stake.addr, encodeValidatorDelegationRcStakeEpochQueueForMockKey(validatorAddr))
	if len(value) != 0 {
		err := rlp.DecodeBytes(value, &queue)
		if nil != err {
			return err
		}

	} else {
		queue = make([]uint64, 0)
	}

	queue = append(queue, stakeEpoch)
	value, err := rlp.EncodeToBytes(queue)
	if nil != err {
		return err
	}
	stake.statedb.SetState(stake.addr, encodeValidatorDelegationRcStakeEpochQueueForMockKey(validatorAddr), value)
	stake.statedb.SetState(stake.addr, encodeValidatorDelegationRcForMockKey(validatorAddr, stakeEpoch), common.Uint64ToBytes(numberOfRcs))
	return nil
}

func (stake *MockStakeModule) GetStakeWithdrawalWaitPeriod(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, stakeWithdrawalWaitPeriodKey))
}
func (stake *MockStakeModule) GetDelegateWithdrawalWaitPeriod(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, delegateWithdrawalWaitPeriodKey))
}
func (stake *MockStakeModule) GetSlashingPercentage(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, slashingPercentageKey))
}
func (stake *MockStakeModule) GetSlashIncentivePercentage(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, slashIncentivePercentageKey))
}
func (stake *MockStakeModule) GetMaxRoundValidatorsSize(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, maxRoundValidatorsSizeKey))
}
func (stake *MockStakeModule) GetMaxEpochValidatorsSize(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, maxEpochValidatorsSizeKey))
}
func (stake *MockStakeModule) GetMinBlocksOfRoundValidator(statedb sdk.StateDBReader) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, minBlocksOfRoundValidatorKey))
}
func (stake *MockStakeModule) GetRoundValidatorIds(statedb sdk.StateDBReader, round uint64) []common.Address {

	value := stake.statedb.GetState(stake.addr, encodeRoundValidatorSharesSnapshotQueueKey(round))
	if len(value) == 0 {
		return nil
	}
	var queue MockValidatorSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil != err {
		return nil
	}

	arr := make([]common.Address, len(queue))
	for i, snap := range queue {
		arr[i] = snap.ValidatorAddr
	}

	return arr
}
func (stake *MockStakeModule) GetEpochValidatorIds(statedb sdk.StateDBReader, epoch uint64) []common.Address {

	value := stake.statedb.GetState(stake.addr, encodeEpochValidatorSharesSnapshotQueueKey(epoch))
	if len(value) == 0 {
		return nil
	}
	var queue MockValidatorSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil != err {
		return nil
	}

	arr := make([]common.Address, len(queue))
	for i, snap := range queue {
		arr[i] = snap.ValidatorAddr
	}

	return arr
}
func (stake *MockStakeModule) GetEpochValidatorSnapQueueFlatten(statedb sdk.StateDBReader, epoch uint64) ([]common.Address, []*big.Int, []*big.Int, []uint64, []uint64, []uint64, []uint64) {

	value := stake.statedb.GetState(stake.addr, encodeEpochValidatorSharesSnapshotQueueKey(epoch))
	if len(value) == 0 {
		return nil, nil, nil, nil, nil, nil, nil
	}
	var queue MockValidatorSnapshotQueue
	if err := rlp.DecodeBytes(value, &queue); nil != err {
		return nil, nil, nil, nil, nil, nil, nil
	}

	validatorAddrQueue := make([]common.Address, len(queue))
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

func (stake *MockStakeModule) getValidator(validatorAddr common.Address) *MockValidator {
	value := stake.statedb.GetState(stake.addr, encodeValidatorKey(validatorAddr))
	if len(value) == 0 {
		return nil
	}
	var validator MockValidator
	if err := rlp.DecodeBytes(value, &validator); nil != err {
		return nil
	}
	return &validator
}
func (stake *MockStakeModule) getDelegation(delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) *MockDelegation {
	value := stake.statedb.GetState(stake.addr, encodeDelegatorKey(delegatorAddr, validatorAddr, stakeEpoch))
	if len(value) == 0 {
		return nil
	}
	var delegation MockDelegation
	if err := rlp.DecodeBytes(value, &delegation); nil != err {
		return nil
	}
	return &delegation
}
func (stake *MockStakeModule) IsValidValidator(statedb sdk.StateDBReader, validatorAddr common.Address) bool {
	validator := stake.getValidator(validatorAddr)
	return validator.IsNotEmpty() && validator.Status.IsValid()
}
func (stake *MockStakeModule) IsInvalidValidator(statedb sdk.StateDBReader, validatorAddr common.Address) bool {
	return !stake.IsValidValidator(statedb, validatorAddr)
}
func (stake *MockStakeModule) IsEmptyValidator(statedb sdk.StateDBReader, validatorAddr common.Address) bool {
	validator := stake.getValidator(validatorAddr)
	return validator.IsEmpty()
}
func (stake *MockStakeModule) GetValidatorCommissionRate(statedb sdk.StateDBReader, validatorAddr common.Address) uint64 {
	validator := stake.getValidator(validatorAddr)
	if validator.IsEmpty() {
		return 0
	}
	return validator.CommissionRate
}
func (stake *MockStakeModule) GetValidatorStakeEpoch(statedb sdk.StateDBReader, validatorAddr common.Address) uint64 {
	validator := stake.getValidator(validatorAddr)
	if validator.IsEmpty() {
		return 0
	}
	return validator.Epoch
}
func (stake *MockStakeModule) GetValidatorStakeAmount(statedb sdk.StateDBReader, validatorAddr common.Address) *big.Int {
	validator := stake.getValidator(validatorAddr)
	if validator.IsEmpty() {
		return big.NewInt(0)
	}
	return validator.StakeAmount
}
func (stake *MockStakeModule) GetValidatorDelegateAmount(statedb sdk.StateDBReader, validatorAddr common.Address) *big.Int {
	validator := stake.getValidator(validatorAddr)
	if validator.IsEmpty() {
		return big.NewInt(0)
	}
	return validator.DelegateAmount
}
func (stake *MockStakeModule) GetValidatorOwner(statedb sdk.StateDBReader, validatorAddr common.Address) common.Address {
	validator := stake.getValidator(validatorAddr)
	if validator.IsEmpty() {
		return common.ZeroAddr
	}
	return validator.Owner
}
func (stake *MockStakeModule) GetValidatorECDSAPubKey(statedb sdk.StateDBReader, validatorAddr common.Address) *ecdsa.PublicKey {
	validator := stake.getValidator(validatorAddr)
	if validator.IsEmpty() {
		return nil
	}
	pubkey, _ := validator.PubKey.Pubkey()
	return pubkey
}
func (stake *MockStakeModule) GetNumberOfBlocksForRoundValidator(statedb sdk.StateDBReader, validatorAddr common.Address, round uint64) uint64 {
	return common.BytesToUint64(stake.statedb.GetState(stake.addr, encodeNumberOfBlocksForRoundValidatorKey(validatorAddr, round)))
}
func (stake *MockStakeModule) GetEpochByValidatorDelegationRcPending(statedb sdk.StateDBReader, validatorAddr common.Address) []uint64 {
	var queue []uint64
	value := stake.statedb.GetState(stake.addr, encodeValidatorDelegationRcStakeEpochQueueForMockKey(validatorAddr))
	if len(value) != 0 {
		err := rlp.DecodeBytes(value, &queue)
		if nil != err {
			return nil
		}
	}
	return queue
}
func (stake *MockStakeModule) GetDelegationFlatten(statedb sdk.StateDBReader, delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) (uint64, *big.Int, *big.Int) {
	delegation := stake.getDelegation(delegatorAddr, validatorAddr, stakeEpoch)
	if delegation.IsEmpty() {
		return 0, big.NewInt(0), big.NewInt(0)
	}
	return delegation.Epoch, delegation.PreEpochAmount, delegation.Amount
}
func (stake *MockStakeModule) UpdateDelegationEpoch(statedb sdk.StateDB, delegatorAddr, validatorAddr common.Address, stakeEpoch, delegateEpoch uint64) error {
	delegation := stake.getDelegation(delegatorAddr, validatorAddr, stakeEpoch)
	if delegation.IsEmpty() {
		return db.ErrNotFound
	}
	if delegation.Epoch == delegateEpoch {
		return nil
	}
	if delegation.Epoch > delegateEpoch {
		return fmt.Errorf("new delegate epoch not greater than old, old epoch: %d, new epoch: %d", delegation.Epoch, delegateEpoch)
	}
	delegation.Epoch = delegateEpoch
	return stake.MockDelegation(delegatorAddr, validatorAddr, stakeEpoch, delegation)
}

const (
	/**
	######   ######   ######   ######
	#	  THE VALIDATOR  STATUS     #
	######   ######   ######   ######
	*/
	Invalided    MockValidatorStatus = 1 << iota // 0001: The validator is deactivated
	LowBlocks                                    // 0010: The validator was low block rate
	LowThreshold                                 // 0100: The validator's stake was lower than minimum stake threshold
	Duplicated                                   // 1000: The validator was duplicate block or duplicate signature
	Unstaked                                     // 0010,0000: The validator was unstaked
	Slashing                                     // 0100,0000: The validator is being slashed
	Valided      = 0                             // 0000: The validator was activated
	NotExist     = 1 << 31                       // 1000,xxxx,... : The validator is not exist
)

type MockValidatorStatus uint32

func (status MockValidatorStatus) IsValid() bool {
	return !status.IsInvalid()
}
func (status MockValidatorStatus) IsInvalid() bool {
	return status&Invalided == Invalided
}
func (status MockValidatorStatus) IsOnlyInvalid() bool {
	return status&Invalided == status|Invalided
}

func (status MockValidatorStatus) IsLowBlocks() bool {
	return status&LowBlocks == LowBlocks
}
func (status MockValidatorStatus) IsOnlyLowBlocks() bool {
	return status&LowBlocks == status|LowBlocks
}
func (status MockValidatorStatus) IsInvalidLowBlocks() bool {
	return status&(Invalided|LowBlocks) == (Invalided | LowBlocks)
}

func (status MockValidatorStatus) IsLowThreshold() bool {
	return status&LowThreshold == LowThreshold
}
func (status MockValidatorStatus) IsOnlyLowThreshold() bool {
	return status&LowThreshold == status|LowThreshold
}
func (status MockValidatorStatus) IsInvalidLowThreshold() bool {
	return status&(Invalided|LowThreshold) == (Invalided | LowThreshold)
}

func (status MockValidatorStatus) IsDuplicated() bool {
	return status&Duplicated == Duplicated
}
func (status MockValidatorStatus) IsInvalidDuplicated() bool {
	return status&(Duplicated|Invalided) == (Duplicated | Invalided)
}

func (status MockValidatorStatus) IsUnstaked() bool { return status&Unstaked == Unstaked }
func (status MockValidatorStatus) IsOnlyUnstaked() bool {
	return status&Unstaked == status|Unstaked
}
func (status MockValidatorStatus) IsInvalidUnstaked() bool {
	return status&(Invalided|Unstaked) == (Invalided | Unstaked)
}
func (status MockValidatorStatus) IsOnlyInvalidUnstaked() bool {
	return status&(Invalided|Unstaked) == status|(Invalided|Unstaked)
}

func (status MockValidatorStatus) IsSlashing() bool     { return status&Slashing == Slashing }
func (status MockValidatorStatus) IsOnlySlashing() bool { return status&Slashing == status|Slashing }
func (status MockValidatorStatus) IsInvalidSlashing() bool {
	return status&(Invalided|Slashing) == (Invalided | Slashing)
}
func (status MockValidatorStatus) IsOnlyInvalidSlashing() bool {
	return status&(Invalided|Slashing) == status|(Invalided|Slashing)
}

func (status MockValidatorStatus) IsNotExist() bool {
	return status&NotExist == NotExist
}

type MockValidator struct {
	Status         MockValidatorStatus
	CommissionRate uint64
	Epoch          uint64
	StakeIndex     uint64
	Owner          common.Address
	StakeAmount    *big.Int
	DelegateAmount *big.Int
	PubKey         enode.IDv0
	BlsKey         []byte
}

func NewMockValidator(owner common.Address, epoch, index, commissionRate uint64, status MockValidatorStatus, stakeAmount, delegateAmount *big.Int, pubKey enode.IDv0, blsKey []byte) *MockValidator {
	return &MockValidator{
		Status:         status,
		CommissionRate: commissionRate,
		Epoch:          epoch,
		StakeIndex:     index,
		Owner:          owner,
		StakeAmount:    stakeAmount,
		DelegateAmount: delegateAmount,
		PubKey:         pubKey,
		BlsKey:         blsKey,
	}
}

func (item *MockValidator) IsEmpty() bool {
	return nil == item
}
func (item *MockValidator) IsNotEmpty() bool {
	return !item.IsEmpty()
}

func (item *MockValidator) AppendStatus(status MockValidatorStatus) {
	item.Status |= status
}

type MockPriorityValidator struct {
	PreKey        []byte // previous priority validator key in statedb
	NextKey       []byte // next priority validator key in statedb
	ValidatorAddr common.Address
}

func NewMockPriorityValidator(preKey, nextKey []byte, addr common.Address) *MockPriorityValidator {
	return &MockPriorityValidator{
		PreKey:        preKey,
		NextKey:       nextKey,
		ValidatorAddr: addr,
	}
}

func (item *MockPriorityValidator) IsEmpty() bool {
	return nil == item
}

func (item *MockPriorityValidator) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type MockValidatorSnapshot struct {
	ValidatorAddr  common.Address
	Epoch          uint64
	StakeIndex     uint64
	ValidatorTerm  uint64
	CommissionRate uint64
	StakeAmount    *big.Int
	DelegateAmount *big.Int
}

func NewMockValidatorSnapshot(validatorAddr common.Address, epoch, stakeIndex, commissionRate uint64, stakeAmount, delegateAmount *big.Int) *MockValidatorSnapshot {
	return &MockValidatorSnapshot{
		ValidatorAddr:  validatorAddr,
		Epoch:          epoch,
		StakeIndex:     stakeIndex,
		CommissionRate: commissionRate,
		StakeAmount:    stakeAmount,
		DelegateAmount: delegateAmount,
	}
}

func (item *MockValidatorSnapshot) IsEmpty() bool {
	return nil == item
}
func (item *MockValidatorSnapshot) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type MockValidatorSnapshotQueue []*MockValidatorSnapshot

func NewMockValidatorSnapshotQueue(size uint64) MockValidatorSnapshotQueue {
	queue := make(MockValidatorSnapshotQueue, size)
	return queue
}

func (queue MockValidatorSnapshotQueue) IsEmpty() bool {
	return len(queue) == 0
}
func (queue MockValidatorSnapshotQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

type MockDelegation struct {
	Epoch          uint64 // delegate epoch
	PreEpochAmount *big.Int
	Amount         *big.Int
}

func NewMockDelegation(epoch uint64, amount *big.Int) *MockDelegation {
	return &MockDelegation{
		Epoch:          epoch,
		PreEpochAmount: big.NewInt(0),
		Amount:         amount,
	}
}
func (item *MockDelegation) IsEmpty() bool {
	return nil == item
}

func (item *MockDelegation) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type MockRewardModule struct {
	addr        common.Address
	statedb     sdk.StateDB
	stageModule *MockStageModule
	stakeMdoule *MockStakeModule
}

func NewMockRewardModule(statedb sdk.StateDB, stageModule *MockStageModule, stakeMdoule *MockStakeModule) *MockRewardModule {
	return &MockRewardModule{
		addr:        constants.RewardManagerAddress,
		statedb:     statedb,
		stageModule: stageModule,
		stakeMdoule: stakeMdoule,
	}
}

func (reward *MockRewardModule) MockRewardPerBlockKey(amount *big.Int) {
	reward.statedb.SetState(reward.addr, rewardPerBlockKey, amount.Bytes())
}

func (reward *MockRewardModule) MockRewardPerEpochKey(amount *big.Int) {
	reward.statedb.SetState(reward.addr, rewardPerEpochKey, amount.Bytes())
}

func (reward *MockRewardModule) UpdateDelegationRewardsByStakeEpoch(statedb sdk.StateDB, delegatorAddr, validatorAddr common.Address, stakeEpoch uint64) error {
	return nil
}
func (reward *MockRewardModule) UpdateDelegationRewards(statedb sdk.StateDB, delegatorAddr, validatorAddr common.Address) error {
	return nil
}
func (reward *MockRewardModule) GetRewardPerBlock(statedb sdk.StateDBReader) *big.Int {
	return new(big.Int).SetBytes(reward.statedb.GetState(reward.addr, rewardPerBlockKey))
}
func (reward *MockRewardModule) GetRewardPerEpoch(statedb sdk.StateDBReader) *big.Int {
	return new(big.Int).SetBytes(reward.statedb.GetState(reward.addr, rewardPerEpochKey))
}

type MockVRFModule struct {
	addr        common.Address
	statedb     sdk.StateDB
	stageModule *MockStageModule
	stakeMdoule *MockStakeModule
}

func NewMockVRFModule(statedb sdk.StateDB, stageModule *MockStageModule, stakeMdoule *MockStakeModule) *MockVRFModule {
	return &MockVRFModule{
		addr:        constants.VRFManagerAddress,
		statedb:     statedb,
		stageModule: stageModule,
		stakeMdoule: stakeMdoule,
	}
}

func (vrf *MockVRFModule) getNonceByBlock(blockNumber uint64) common.Hash {
	value := vrf.statedb.GetState(vrf.addr, encodeNonceOfVRFForMockKey(blockNumber))
	if len(value) == 0 {
		return common.ZeroHash
	}
	return common.BytesToHash(value)
}

func (vrf *MockVRFModule) MockNonceOfVRF(blockNumber uint64, nonce common.Hash) {
	vrf.statedb.SetState(vrf.addr, encodeNonceOfVRFForMockKey(blockNumber), nonce.Bytes())
}

func (vrf *MockVRFModule) GetNonceQueueFromTail(statedb sdk.StateDBReader, blockNumber, size uint64) ([]common.Hash, error) {

	arr := make([]common.Hash, 0)

	index := blockNumber
	nonce := vrf.getNonceByBlock(index)
	var count uint64
	for nonce != common.ZeroHash && count < size {
		arr = append(arr, nonce)
		index--
		nonce = vrf.getNonceByBlock(index)
		count++
	}

	return arr, nil
}
func (vrf *MockVRFModule) GetCurrentNonce(statedb sdk.StateDBReader, blockNumber uint64) (common.Hash, error) {
	nonce := vrf.getNonceByBlock(blockNumber)
	if nonce == common.ZeroHash {
		return common.ZeroHash, ErrNotFound
	}
	return nonce, nil
}
