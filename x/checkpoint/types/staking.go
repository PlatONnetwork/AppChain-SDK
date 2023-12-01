package types

import (
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Staking interface {
	IsEndOfEpoch(blockNumber uint64) bool
	GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error)
	BlocksOfEpoch() uint64
}
