package types

import "github.com/PlatONnetwork/PlatON-Go/common"

type ExtraVote interface {
	GetProof(epoch, view uint64, leaf common.Hash) (uint64, []common.Hash, error)
}
