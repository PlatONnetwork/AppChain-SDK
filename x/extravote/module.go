package extravote

import (
	"errors"

	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	cacheSize         = 100
	ExtraVoteDatabase = "extravote"
)

// TODO 处理扩展投票，对每个子模块进行扩展，生成投票 Merkle 证明
type ExtraVote struct {
	modules []module.ConsensusExtendModule
	db      *ExtraVoteDB
}

func NewExtraVote(store store.Store, ms []module.ConsensusExtendModule) *ExtraVote {
	db := NewExtraVoteDB(store)
	return &ExtraVote{modules: ms, db: db}
}

func (e *ExtraVote) Name() string {
	return "extravote"
}

func (e *ExtraVote) ExtendData(ctx sdk.Context) []byte {
	data := make([][]byte, len(e.modules))
	for i, m := range e.modules {
		data[i] = m.ExtendData(ctx)
		log.Info("Extend data for module", "module", m.Name(), "data", len(data[i]))
	}
	extraData, _ := rlp.EncodeToBytes(data)
	return extraData
}

func (e *ExtraVote) VerifyExtendData(ctx sdk.Context, data []byte) (common.Hash, error) {
	cc := ctx.(sdk.ConsensusContext)

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
		leaf, err := m.VerifyExtendData(ctx, extraData[i])
		if err != nil {
			log.Error("Failed to verify extend data", "i", i, "module", m.Name(), "err", err)
			return common.Hash{}, err
		}
		trieNodes = append(trieNodes, leaf.Bytes())
	}
	trie, err := merkle.NewMerkleTree(trieNodes)
	if err != nil {
		return common.Hash{}, err
	}
	e.db.InsertProof(cc.Epoch(), cc.View(), cc.BlockIndex(), trieNodes, trie)

	return trie.Hash(), nil
}

func (e *ExtraVote) PrepareQC(ctx sdk.Context, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	for _, m := range e.modules {
		m.PrepareQC(ctx, block, votes)
	}
}
