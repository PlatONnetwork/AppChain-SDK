package statesync

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/sync"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	types2 "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
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

func (s *StateSync) Name() string {
	return "statesync"
}
func (s *StateSync) ContractAddress() common.Address {
	return contracts.StateSyncAddress
}
func (s *StateSync) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stateReceiver, _ := contracts.NewStateReceiver(evm, contract, readOnly)
	return stateReceiver.Run(input)
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
func (s *StateSync) AddTxs(ctx sdk.Context, local, remote map[common.Address]types.Transactions) (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	cc := ctx.(sdk.WorkerContext)
	//创建commitment
	receiver, err := s.newStateSyncCallContract(cc.Header().ParentHash)
	if err != nil {
		return local, remote
	}
	syncId, err := receiver.GetStateSyncId()
	if err != nil {
		return local, remote
	}
	commitment, err := receiver.GetCommitmentByStateSyncId(syncId)
	if err != nil {
		return local, remote
	}
	start := new(big.Int).Add(commitment.EndId, big.NewInt(1))
	match, err := s.eventProofDb.FindProofRoot(start)
	if err != nil {
		return local, remote
	}
	blockHash := s.eventProofDb.GetRootBlock(match.Root)

	block := s.backend.GetBlockByHash(blockHash)
	_, qc, err := types2.DecodeExtra(block.ExtraData())
	if err != nil {
		return local, remote
	}
	index, voteProof, err := s.extraDb.GetProof(qc.Epoch, qc.ViewNumber, qc.BlockIndex, match.Root[:])
	from := crypto.PubkeyToAddress(s.privateKey.PublicKey)
	nonce, err := s.backend.GetPoolNonce(from)
	if err != nil {
		return local, remote
	}
	cmtx, err := s.createCommitTx(match, index, qc, voteProof, nonce)
	if err != nil {
		return local, remote
	}

	//创建 event proof
	eventId := new(big.Int).Add(syncId, big.NewInt(1))
	var events []*sync.StateSender
	var proofs [][]common.Hash
	for {
		event, err := s.l1Sync.SyncDB().GetStateSenderEvent(eventId)
		if event == nil || err != nil {
			break
		}
		proof, err := s.eventProofDb.GetProof(match.Root, eventId)
		if event == nil || err != nil {
			break
		}
		events = append(events, event)
		proofs = append(proofs, proof)
	}
	exTxs, err := s.createExecuteTxs(proofs, events, nonce+1)
	if err != nil {
		return local, remote
	}
	if local[from] == nil {
		local[from] = types.Transactions{}
	}
	local[from] = append(local[from], cmtx)
	local[from] = append(local[from], exTxs...)
	return local, remote
}
