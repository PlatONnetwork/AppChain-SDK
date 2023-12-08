package statesync

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/utils"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/AppChain-SDK/x"
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
	"gopkg.in/urfave/cli.v1"
)

// 同步 L1 事件， 提供 ExtraData，验证 ExtraData
type StateSync struct {
	logger       log.Logger
	rpcAddress   string
	keystoreFile string
	passwordFile string
	startBlock   *big.Int
	store        store.Store
	privateKey   *ecdsa.PrivateKey
	extraDb      *extravote.ExtraVoteDB
	l1Sync       *sync.L1Sync
	eventProofDb *EventProofDB
	backend      sdk.Backend
	p2p          *SyncP2P
	syncUpdateCh chan struct{}
}

// TODO 启动查询合约执行的ID序号，定位同步的起始点
// TODO 动态的清理数据库数据
func NewStateSync(ctx *cli.Context, store store.Store, extraDb *extravote.ExtraVoteDB) (*StateSync, error) {
	var start *big.Int
	if ctx.GlobalIsSet(StartBlockFlag.Name) {
		start = new(big.Int).SetUint64(ctx.GlobalUint64(StartBlockFlag.Name))
	}
	eventProofDb := NewEventProofDB(store)
	return &StateSync{
		logger:       log.New("module", "statesync"),
		rpcAddress:   ctx.GlobalString(x.RootchainNodeRPCFlag.Name),
		keystoreFile: ctx.GlobalString(utils.KeystoreFlag.Name),
		passwordFile: ctx.GlobalString(utils.PasswordFlag.Name),
		startBlock:   start,
		store:        store,
		eventProofDb: eventProofDb,
		extraDb:      extraDb,
		p2p:          NewSyncP2P(),
		syncUpdateCh: make(chan struct{}),
	}, nil
}

func (s *StateSync) Name() string {
	return "statesync"
}

func (s *StateSync) Init() error {
	if s.rpcAddress == "" {
		return fmt.Errorf("node rpc address not set")
	}

	key, err := utils.DecodePrivateKey(s.keystoreFile, s.passwordFile)
	if err != nil {
		return err
	}
	s.privateKey = key.PrivateKey

	l1Sync, err := sync.NewL1Sync(constants.StateSyncAddress, s.rpcAddress, s.startBlock, s.store, s.syncUpdateCh)
	if err != nil {
		return err
	}
	s.l1Sync = l1Sync
	go s.p2p.Run(context.Background())
	go s.p2p.Run(context.Background())
	return nil
}

func (s *StateSync) listen(ctx context.Context) error {
	for {
		select {
		case <-s.syncUpdateCh:
			id, err := s.l1Sync.SyncDB().GetMaxSyncId()
			if err != nil || id == nil {
				s.logger.Warn("get max sync id failed", "err", err)
			}
			number, err := s.l1Sync.SyncDB().LastBlockNumber()
			if err != nil || id == nil {
				s.logger.Warn("get last block number failed", "err", err)
			}
			s.p2p.SetSyncStatus(&SyncStatus{Id: id, BlockNumber: number})
			s.logger.Debug("set sync statue", "id", id, "number", number)
		case <-ctx.Done():
			break
		}
	}
}

func (s *StateSync) Address() common.Address {
	return constants.StateSyncAddress
}

func (s *StateSync) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	stateReceiver, _ := contracts.NewStateReceiver(evm, contract, readOnly)
	return stateReceiver.Run(input)
}

func (s *StateSync) Protocols() []p2p.Protocol {
	return s.p2p.Protocols()
}

func (s *StateSync) ExtendData(ctx sdk.ConsensusContext) []byte {
	return s.ExtendDataImpl(ctx, ctx.Epoch(), ctx.View(), ctx.BlockIndex(), ctx.Header())
}

func (s *StateSync) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) error {
	return s.VerifyExtendDataImpl(ctx.Epoch(), ctx.View(), ctx.BlockIndex(), ctx.Header(), data)
}

func (s *StateSync) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	s.PrepareQCImpl(block, votes)
}
func (s *StateSync) AddTxs(ctx sdk.WorkerContext, local, remote map[common.Address]types.Transactions) (map[common.Address]types.Transactions, map[common.Address]types.Transactions) {
	//创建commitment
	receiver, err := s.newStateSyncCallContract(ctx, ctx.Header())
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

	block := ctx.Backend().GetBlockByHash(blockHash)
	_, qc, err := types2.DecodeExtra(block.ExtraData())
	if err != nil {
		return local, remote
	}
	index, voteProof, err := s.extraDb.GetProof(qc.Epoch, qc.ViewNumber, qc.BlockIndex, match.Root[:])
	from := crypto.PubkeyToAddress(s.privateKey.PublicKey)
	nonce, err := ctx.Backend().GetPoolNonce(from)
	if err != nil {
		return local, remote
	}
	cmtx, err := s.createCommitTx(ctx, match, index, qc, voteProof, nonce)
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
	exTxs, err := s.createExecuteTxs(ctx, proofs, events, nonce+1)
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
