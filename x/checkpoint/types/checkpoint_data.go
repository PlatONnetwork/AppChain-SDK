package types

import (
	"fmt"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
)

type CheckpointData struct {
	EpochNumber           uint64
	ViewNumber            uint64
	BlockIndex            uint32
	BlockNumber           uint64
	BlockHash             common.Hash
	CurrentValidatorsHash common.Hash
	NextValidatorsHash    common.Hash
	EventRoot             common.Hash
}

type StorageCheckpointData struct {
	*CheckpointData
	ExtendRoot common.Hash
	Signature  []byte
	Bitmap     []byte
}

func (s *StorageCheckpointData) String() string {
	return fmt.Sprintf("{epoch:%d,view:%d,index:%d,number:%d,hash:%s,current:%s,next:%s,root:%s}",
		s.EpochNumber,
		s.ViewNumber,
		s.BlockIndex,
		s.BlockNumber,
		s.BlockHash.String(),
		s.CurrentValidatorsHash.String(),
		s.NextValidatorsHash.String(),
		s.EventRoot.String())
}

func (cd *CheckpointData) MarshalRLP() []byte {
	bs, _ := rlp.EncodeToBytes(cd)
	return bs
}

func (cd *CheckpointData) UnmarshalRLP(data []byte) error {
	return rlp.DecodeBytes(data, cd)
}

func (cd *CheckpointData) Hash() (common.Hash, error) {
	return crypto.Keccak256Hash(cd.MarshalRLP()), nil
}
