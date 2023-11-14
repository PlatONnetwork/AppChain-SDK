package statesync

import (
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestDB(t *testing.T) {
	store := memorydb.New()
	db := NewEventProofDB(store)
	var trieNodes [][]byte

	leaves := make(map[*big.Int]common.Hash)
	for i := 0; i < 10; i++ {
		leaves[big.NewInt(int64(i))] = common.BigToHash(big.NewInt(int64(i)))
		trieNodes = append(trieNodes, common.BigToHash(big.NewInt(int64(i))).Bytes())
	}
	tree, err := merkle.NewMerkleTree(trieNodes)
	require.Nil(t, err)
	err = db.InsertProof(1, 1, 1, big.NewInt(0), big.NewInt(10), leaves, tree)
	require.Nil(t, err)
	for i := 0; i < 10; i++ {
		proof, err := db.GetProof(tree.Hash(), big.NewInt(int64(i)))
		require.Nil(t, err)
		err = merkle.VerifyProof(uint64(i), common.BigToHash(big.NewInt(int64(i))).Bytes(), proof, tree.Hash())
		require.Nil(t, err)
	}
	root := db.GetProofRoot(1, 1, 1)
	require.Equal(t, tree.Hash(), root)
	cm, err := db.FindProofRoot(big.NewInt(0))
	require.Nil(t, err)
	require.Equal(t, tree.Hash(), common.BytesToHash(cm.Root[:]))
}
