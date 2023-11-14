package statesync

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/sync"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

// 同步 L1 事件， 提供 ExtraData，验证 ExtraData
type StateSync struct {
	privateKey   *ecdsa.PrivateKey
	extraDb      *extravote.ExtraVoteDB
	l1Sync       *sync.L1Sync
	eventProofDb *EventProofDB
	backend      sdk.Backend
	p2p          *SyncP2P
}

// TODO 启动查询合约执行的ID序号，定位同步的起始点
// TODO 动态的清理数据库数据
func NewStateSync(url string, start *big.Int, store store.Store, privateKey *ecdsa.PrivateKey, extraDb *extravote.ExtraVoteDB, backend sdk.Backend) (*StateSync, error) {
	l1Sync, err := sync.NewL1Sync(contracts.StateSyncAddress, url, start, store)
	if err != nil {
		return nil, err
	}
	eventProofDb := NewEventProofDB(store)
	return &StateSync{
		l1Sync:       l1Sync,
		eventProofDb: eventProofDb,
		privateKey:   privateKey,
		extraDb:      extraDb,
		backend:      backend,
		p2p:          NewSyncP2P(),
	}, nil
}

func (s *StateSync) Protocols() []p2p.Protocol {
	return s.p2p.Protocols()
}

func (s *StateSync) ExtendData(ctx sdk.Context) []byte {
	cc := ctx.Context().(sdk.ConsensusContext)
	return s.ExtendDataImpl(cc.Epoch(), cc.View(), cc.BlockIndex(), cc.Header())
}

func (s *StateSync) VerifyExtendData(ctx sdk.Context, data []byte) (common.Hash, error) {
	cc := ctx.Context().(sdk.ConsensusContext)
	return s.VerifyExtendDataImpl(cc.Epoch(), cc.View(), cc.BlockIndex(), cc.Header(), data)
}

func (s *StateSync) PrepareQC(ctx sdk.Context, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	s.PrepareQCImpl(block, votes)
}
