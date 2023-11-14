package extravote_test

import (
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestDB(t *testing.T) {
	store := memorydb.New()
	db := extravote.NewExtraVoteDB(store)
	leaves := make([][]byte, 0, 10)
	for i := 0; i < 10; i++ {
		leaves = append(leaves, common.BigToAddress(big.NewInt(int64(i))).Bytes())
	}
	tree, err := merkle.NewMerkleTree(leaves)
	require.Nil(t, err)

	err = db.InsertProof(1, 1, 1, leaves, tree)
	require.Nil(t, err)
	for i, leaf := range leaves {
		index, proof, err := db.GetProof(1, 1, 1, leaf)
		require.Nil(t, err)
		require.Equal(t, i, int(index))
		//var ps []common.Hash
		//for _, p := range proof {
		//	ps = append(ps, common.BytesToHash(p))
		//}
		require.Nil(t, merkle.VerifyProof(index, leaf, proof, tree.Hash()))
	}
}
