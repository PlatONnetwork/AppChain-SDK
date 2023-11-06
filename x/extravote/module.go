package extravote

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/hashicorp/golang-lru/simplelru"
)

const (
	cacheSize         = 100
	ExtraVoteDatabase = "extravote"
)

// TODO 处理扩展投票，对每个子模块进行扩展，生成投票 Merkle 证明
type ExtraVote struct {
	modules []module.ConsensusExtendModule
	cache   simplelru.LRUCache
	db      *ExtraVoteDB
}

func NewExtraVote(store store.Store, ms []module.ConsensusExtendModule) *ExtraVote {
	db := NewExtraVoteDB(store)
	cache, _ := simplelru.NewLRU(cacheSize, nil)
	return &ExtraVote{modules: ms, cache: cache, db: db}
}

func (e *ExtraVote) Name() string {
	return "extravote"
}

func (e *ExtraVote) ExtendData(ctx sdk.Context, epoch, view uint64, header *types.Header) []byte {
	data := make([][]byte, len(e.modules))
	for i, m := range e.modules {
		data[i] = m.ExtendData(ctx, epoch, view, header)
	}
	extraData, _ := rlp.EncodeToBytes(data)
	return extraData
}

func (e *ExtraVote) VerifyExtendData(ctx sdk.Context, epoch, view uint64, header *types.Header, data []byte) (common.Hash, error) {
	var extraData [][]byte
	err := rlp.DecodeBytes(data, &extraData)
	if err != nil {
		return common.Hash{}, err
	}
	if len(extraData) != len(e.modules) {
		return common.Hash{}, errors.New("invalid extra data length")
	}
	var trieNodes [][]byte
	for i, m := range e.modules {
		leaf, err := m.VerifyExtendData(ctx, epoch, view, header, extraData[i])
		if err != nil {
			return common.Hash{}, err
		}
		trieNodes = append(trieNodes, leaf.Bytes())
	}
	trie, err := merkle.NewMerkleTree(trieNodes)
	if err != nil {
		return common.Hash{}, err
	}
	e.db.InsertProof(epoch, view, trieNodes, trie)

	e.cache.Add(trie.Hash(), trie)
	return trie.Hash(), nil
}

func (e *ExtraVote) PrepareQC(ctx sdk.Context, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	for _, m := range e.modules {
		m.PrepareQC(ctx, block, votes)
	}
}
