package statesync

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/utils"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/contracts"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"math/big"
)

var (
	batchSize = 1024 * 1024
)

func encodeEpochViewIndexHash(epoch, view uint64, index uint32) []byte {
	epochValue := utils.EncodeUint64ToBytes(epoch)
	viewValue := utils.EncodeUint64ToBytes(view)
	indexValue := utils.EncodeUint32ToBytes(index)
	return append([]byte("evi"), append(epochValue, append(viewValue, indexValue...)...)...)
}

func encodeHashId(hash common.Hash, id *big.Int) []byte {
	return append([]byte("hi"), append(hash.Bytes(), math.PaddedBigBytes(id, 32)...)...)
}

func encodeStartEndRoot(start, end *big.Int) []byte {
	return append([]byte("ser"), append(math.PaddedBigBytes(start, 32), math.PaddedBigBytes(end, 32)...)...)
}

func encodeStartRootPrefix(start *big.Int) []byte {
	return append([]byte("ser"), math.PaddedBigBytes(start, 32)...)
}

func encodeRootBlock(root common.Hash) []byte {
	return append([]byte("rb"), root.Bytes()...)

}

type EventProofDB struct {
	db store.KVStore
}

func (s *EventProofDB) FindProofRoot(start *big.Int) (*contracts.StateSyncCommitment, error) {
	it := s.db.NewIterator(encodeStartRootPrefix(start), nil)
	if it.Next() {
		var commitment contracts.StateSyncCommitment
		if err := rlp.DecodeBytes(it.Value(), &commitment); err != nil {
			return nil, err
		}
		return &commitment, nil
	}
	return nil, nil
}

func (s *EventProofDB) InsertRootBlock(root, block common.Hash) error {
	return s.db.Set(encodeRootBlock(root), block.Bytes())
}

func (s *EventProofDB) GetRootBlock(root common.Hash) common.Hash {
	if value, _ := s.db.Get(encodeRootBlock(root)); len(value) != 0 {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

func (s *EventProofDB) GetProofRoot(epoch, view uint64, index uint32) common.Hash {
	if value, _ := s.db.Get(encodeEpochViewIndexHash(epoch, view, index)); len(value) != 0 {
		return common.BytesToHash(value)
	}
	return common.Hash{}
}

func (s *EventProofDB) GetProof(root common.Hash, id *big.Int) ([]common.Hash, error) {
	value, err := s.db.Get(encodeHashId(root, id))
	if err != nil {
		return nil, err
	}
	var proof []common.Hash
	if err := rlp.DecodeBytes(value, &proof); err != nil {
		return nil, err
	}
	return proof, err
}

func (s *EventProofDB) InsertProof(epoch, view uint64, index uint32, start, end *big.Int, leave map[*big.Int]common.Hash, tree *merkle.MerkleTree) error {
	if value, _ := s.db.Get(encodeEpochViewIndexHash(epoch, view, index)); len(value) != 0 {
		return errors.New("repeated insertion proof")
	}
	batch := s.db.NewBatch()
	commitment := &contracts.StateSyncCommitment{
		StartId: start,
		EndId:   end,
		Root:    tree.Hash(),
	}
	value, err := rlp.EncodeToBytes(commitment)
	if err != nil {
		return err
	}
	s.db.Set(encodeStartEndRoot(start, end), value)
	s.db.Set(encodeEpochViewIndexHash(epoch, view, index), tree.Hash().Bytes())
	for k, v := range leave {
		proof, err := tree.GenerateProof(v.Bytes())
		if err != nil {
			return err
		}
		value, err := rlp.EncodeToBytes(proof)
		if err != nil {
			return err
		}
		s.db.Set(encodeHashId(tree.Hash(), k), value)
		if batch.ValueSize() > batchSize {
			batch.Write()
		}
	}
	return batch.Write()
}
