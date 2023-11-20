package types

import (
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
)

type Staking interface {
	IsEndOfEpoch(blockNumber uint64) bool
	GetValidator(blockNumber uint64) *cbfttypes.Validators
	BlocksOfEpoch() uint64
}
