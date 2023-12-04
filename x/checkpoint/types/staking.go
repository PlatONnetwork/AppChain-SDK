package types

import (
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Staking interface {
	IsEndOfRound(ctx sdk.ConsensusContext, blockNumber uint64) bool
	GetRoundValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error)
	BlocksOfRound(ctx sdk.ConsensusContext) uint64
}
