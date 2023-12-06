package types

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type StageModuler interface {
	GetCurrentRound(stateDB sdk.StateDBReader) uint64
	GetCurrentEpoch(stateDB sdk.StateDBReader) uint64

	IsElectionBlockOnCurrentRound(db sdk.StateDBReader, blockNumber uint64) bool
	IsElectionBlockOnCurrentEpoch(db sdk.StateDBReader, blockNumber uint64) bool
	IsBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsBeginOfNextRound(db sdk.StateDBReader, blockNumber uint64) bool
	IsBeginOfNextEpoch(db sdk.StateDBReader, blockNumber uint64) bool
	IsEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsEndOfRound(db sdk.StateDBReader, blockNumber, size uint64) bool
	IsEndOfEpoch(db sdk.StateDBReader, blockNumber, size uint64) bool

	IsNotElectionBlockOnCurrentRound(db sdk.StateDBReader, blockNumber uint64) bool
	IsNotElectionBlockOnCurrentEpoch(db sdk.StateDBReader, blockNumber uint64) bool
	IsNotBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsNotBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsNotBeginOfNextRound(db sdk.StateDBReader, blockNumber uint64) bool
	IsNotBeginOfNextEpoch(db sdk.StateDBReader, blockNumber uint64) bool
	IsNotEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsNotEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsNotEndOfRound(db sdk.StateDBReader, blockNumber, size uint64) bool
	IsNotEndOfEpoch(db sdk.StateDBReader, blockNumber, size uint64) bool

	BlocksOfRound(stateDB sdk.StateDBReader, round uint64) uint64
	BlocksOfEpoch(stateDB sdk.StateDBReader, epoch uint64) uint64

	// for stakeModule
	GetLastNumber(stateDB sdk.StateDBReader, blockNumber uint64) uint64

	GetRoundAndBlockBoundByBlockNumber(db sdk.StateDBReader, blockNumber, size uint64) (uint64, uint64, uint64)
	GetEpochAndBlockBoundByBlockNumber(db sdk.StateDBReader, blockNumber, size uint64) (uint64, uint64, uint64)
}

type StakeModuler interface {
}

type RewardModuler interface {
	UpdateDelegationRewardsByStakeEpoch(stateDB sdk.StateDB, delegaterAddr, validatorAddr basecommon.Address, stakeEpoch uint64) error
}

type VRFModuler interface {
	GetNonceQueueUtil(stateDB sdk.StateDBReader, blockNumber, size uint64) ([]basecommon.Hash, error)
	GetCurrentNonce(stateDB sdk.StateDBReader, blockNumber uint64) (basecommon.Hash, error)
}
