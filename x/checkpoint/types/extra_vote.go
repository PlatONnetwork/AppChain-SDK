package types

import "github.com/PlatONnetwork/PlatON-Go/common"

type ExtraVote interface {
	GetProof(epoch, view uint64, index uint32, leaf []byte) (uint64, []common.Hash, error)
}
