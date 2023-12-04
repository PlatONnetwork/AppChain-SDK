package types

import (
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

type Staking interface {
	GetRoundValidatorIds(stateDB sdk.StateDBReader, round uint64) []basecommon.Address
	GetEpochValidatorIds(stateDB sdk.StateDBReader, epoch uint64) []basecommon.Address
	GetValidatorCommissionRate(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) uint64
	GetValidatorStakeAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int
	GetValidatorDelegateAmount(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *big.Int
	GetValidatorOwner(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) basecommon.Address
	GetCurrentRound(stateDB sdk.StateDBReader) uint64
	GetCurrentEpoch(stateDB sdk.StateDBReader) uint64
	IsBeginOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsBeginOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsEndOfCurrentRound(stateDB sdk.StateDBReader, blockNumber uint64) bool
	IsEndOfCurrentEpoch(stateDB sdk.StateDBReader, blockNumber uint64) bool
	GetNumberOfBlocksForRoundValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address, round uint64) uint64
}
