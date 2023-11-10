package extravote

import (
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
)

type ExtraDataProof struct {
	Index uint64
	Proof []common.Hash
}

func encodeEpochViewHash(epoch, view uint64, index uint32, hash []byte) []byte {
	return append(append(utils.EncodeUint64ToBytes(epoch), append(utils.EncodeUint64ToBytes(view), utils.EncodeUint32ToBytes(index)...)...), hash...)
}

type ExtraVoteDB struct {
	db store.KVStore
}

func NewExtraVoteDB(store store.Store) *ExtraVoteDB {
	db := store.GetKVStore(ExtraVoteDatabase)
	return &ExtraVoteDB{
		db: db,
	}
}

func (db *ExtraVoteDB) InsertProof(epoch, view uint64, index uint32, leaves [][]byte, tree *merkle.MerkleTree) error {
	for _, leaf := range leaves {
		proof, err := tree.GenerateProof(leaf)
		if err != nil {
			return err
		}
		leafIndex, err := tree.LeafIndex(leaf)
		if err != nil {
			return err
		}

		raw, err := rlp.EncodeToBytes(&ExtraDataProof{
			Index: leafIndex,
			Proof: proof,
		})
		if err != nil {
			return err
		}
		db.db.Set(encodeEpochViewHash(epoch, view, index, leaf), raw)
	}
	return nil
}

func (db *ExtraVoteDB) GetProof(epoch, view uint64, index uint32, leaf []byte) (uint64, []common.Hash, error) {
	raw, err := db.db.Get(encodeEpochViewHash(epoch, view, index, leaf))
	if err != nil {
		return 0, nil, err
	}
	var extraProof ExtraDataProof
	err = rlp.DecodeBytes(raw, &extraProof)
	if err != nil {
		return 0, nil, err
	}
	return extraProof.Index, extraProof.Proof, nil
}
