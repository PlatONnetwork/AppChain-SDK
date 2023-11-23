package types

import (
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Staking interface {
	IsEndOfRound(blockNumber uint64) bool
	GetValidator(ctx sdk.Context, blockNumber uint64) (*cbfttypes.Validators, error)
	BlocksOfRound() uint64
}
