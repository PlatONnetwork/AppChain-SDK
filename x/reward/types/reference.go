package types

import (
	"math/big"

	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type StageModuler interface {
	GetCurrentRound(stateDB sdk.StateDBReader) uint64
	GetCurrentEpoch(stateDB sdk.StateDBReader) uint64
	IsBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool

	IsNotBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsNotBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsNotEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsNotEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
}

type StakeModuler interface {
	GetRoundValidatorIds(stateDB sdk.StateDBReader, round uint64) []basecommon.Address
	GetEpochValidatorIds(stateDB sdk.StateDBReader, epoch uint64) []basecommon.Address
	GetEpochValidatorSnapQueueFlatten(stateDB sdk.StateDBReader, epoch uint64) ([]basecommon.Address, []*big.Int, []*big.Int, []uint64, []uint64, []uint64, []uint64)
	IsValidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool
	IsInvalidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool
	//IsOnlyInvalidUnstakeValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool
	//IsEmptyValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool
	GetValidatorCommissionRate(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) uint64
	GetValidatorStakeEpoch(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) uint64
	GetValidatorStakeAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int
	GetValidatorDelegateAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int
	GetValidatorOwner(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) basecommon.Address
	GetNumberOfBlocksForRoundValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address, round uint64) uint64

	// for reward contract
	GetEpochByValidatorDelegationRcPending(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) []uint64
	GetDelegationFlatten(stateDB sdk.StateDBReader, delegatorAddr, validatorAddr basecommon.Address, stakeEpoch uint64) (uint64, *big.Int, *big.Int)
	UpdateDelegationEpoch(stateDB sdk.StateDB, delegatorAddr, validatorAddr basecommon.Address, stakeEpoch, delegateEpoch uint64) error
}

type RewardModuler interface {
	UpdateDelegationRewards(stateDB sdk.StateDB, delegatorAddr, validatorAddr basecommon.Address) error

	GetRewardPerBlock(stateDB sdk.StateDBReader) *big.Int
	GetRewardPerEpoch(stateDB sdk.StateDBReader) *big.Int
}
