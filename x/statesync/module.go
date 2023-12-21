package statesync

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/utils"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/l1"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rlp"

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

type ElectionValidator interface {
	GetRoundValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error)
}

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
	l1Module     *l1.L1Module
	eventProofDb *EventProofDB
	backend      sdk.Backend
	p2p          *SyncP2P
	validator    ElectionValidator
	syncUpdateCh chan struct{}
}

func NewStateSync(ctx *cli.Context, l1Module *l1.L1Module, validator ElectionValidator, store store.Store, extraDb *extravote.ExtraVoteDB) (*StateSync, error) {
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
		l1Module:     l1Module,
		validator:    validator,
		p2p:          NewSyncP2P(),
		syncUpdateCh: make(chan struct{}),
	}, nil
}

func (s *StateSync) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	db.SetNonce(constants.StateSyncAddress, 1)
	s.logger.Info("Set StateSync Nonce", "nonce", 1)
	return nil
}

func (s *StateSync) Name() string {
	return "statesync"
}

func (s *StateSync) Init(ctx sdk.InitContext) error {
	if s.rpcAddress == "" {
		return fmt.Errorf("node rpc address not set")
	}

	key, err := utils.DecodePrivateKey(s.keystoreFile, s.passwordFile)
	if err != nil {
		return err
	}
	s.privateKey = key.PrivateKey
	stateAddress, err := s.l1Module.GetStateAddress()
	if err != nil {
		return err
	}
	l1Sync, err := sync.NewL1Sync(stateAddress, s.rpcAddress, s.startBlock, s.store, s.syncUpdateCh)
	if err != nil {
		return err
	}
	s.l1Sync = l1Sync
	go s.p2p.Run(context.Background())
	go s.l1Sync.Run(context.Background())
	go s.listen(context.Background())
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
			s.logger.Debug("set sync status", "id", id, "number", number)
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
func (s *StateSync) AddTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions) (map[common.Address]types.Transactions, error) {
	block := ctx.Backend().GetBlock(ctx.Header().ParentHash, ctx.Header().Number.Uint64()-1)
	if block == nil {
		s.logger.Warn("Get block failed", "number", ctx.Header().Number.Uint64()-1)
		return local, nil
	}
	receiver, err := s.newStateSyncCallContract(ctx, block.Header())
	if err != nil {
		s.logger.Warn("New state sync caller failed", "err", err)
		return local, nil
	}
	from := crypto.PubkeyToAddress(s.privateKey.PublicKey)

	nonce := common2.EnableNonce(local[from], func() uint64 {
		return ctx.StateDB().GetNonce(from)
	})
	if err != nil {
		s.logger.Warn("Get pool nonce failed", "nonce", nonce, "err", err)
		return local, nil
	}
	cmtx, err := s.addCommitTx(ctx, receiver, nonce)
	if err == nil {
		if local[from] == nil {
			local[from] = types.Transactions{}
		}
		local[from] = append(local[from], cmtx)
		nonce += 1
	}
	//创建 event proof
	exTxs, err := s.addExecutedTx(ctx, receiver, nonce)
	if err != nil {
		s.logger.Warn("Get executed txs failed", "err", err)
		return local, nil
	}
	if local[from] == nil {
		local[from] = types.Transactions{}
	}
	local[from] = append(local[from], exTxs...)
	return local, nil
}

func (s *StateSync) addCommitTx(ctx sdk.WorkerContext, receiver *contracts.StateReceiver, nonce uint64) (*types.Transaction, error) {
	syncId, err := receiver.GetStateSyncId()
	if err != nil {
		s.logger.Warn("Get state sync id failed", "err", err)
		return nil, err
	}
	start := new(big.Int).Add(syncId, big.NewInt(1))
	if syncId.Cmp(big.NewInt(0)) != 0 {
		commitment, err := receiver.GetCommitmentByStateSyncId(syncId)
		if err != nil {
			s.logger.Warn("Get commitment state sync id failed", "err", err)
			return nil, err
		}
		start = new(big.Int).Add(commitment.EndId, big.NewInt(1))
	}
	match, err := s.eventProofDb.FindProofRoot(start)
	if match == nil || err != nil {
		s.logger.Warn("Find proof root failed", "start", start, "err", err)
		return nil, errors.New(fmt.Sprintf("find proof failed start:%d", start.Uint64()))
	}

	s.logger.Debug("Find proof root", "proof", match)
	blockHash := s.eventProofDb.GetRootBlock(match.Root)

	block := ctx.Backend().GetBlockByHash(blockHash)
	if block == nil {
		s.logger.Warn("Get block failed", "hash", blockHash.Hex())
		return nil, errors.New(fmt.Sprintf("get block failed:%s", blockHash.Hex()))
	}
	_, qc, err := types2.DecodeExtra(block.ExtraData())
	if err != nil {
		s.logger.Warn("Decode extra failed", "start", start, "err", err)
		return nil, err
	}
	leaf, _ := rlp.EncodeToBytes(match)
	index, voteProof, err := s.extraDb.GetProof(qc.Epoch, qc.ViewNumber, qc.BlockIndex, leaf)
	if err != nil {
		s.logger.Warn("Extra get proof failed", "qc", qc, "err", err)
		return nil, err
	}

	cmtx, err := s.createCommitTx(ctx, match, index, qc, voteProof, nonce)
	if err != nil {
		s.logger.Warn("Create commit tx failed", "err", err)
		return nil, err
	}
	return cmtx, nil
}

func (s *StateSync) addExecutedTx(ctx sdk.WorkerContext, receiver *contracts.StateReceiver, nonce uint64) ([]*types.Transaction, error) {
	syncId, err := receiver.GetStateSyncId()
	if err != nil {
		return nil, err
	}
	executedId, err := receiver.GetExecutedId()
	if err != nil {
		s.logger.Warn("Get executed id failed", "err", err)
		return nil, err
	}
	if syncId.Cmp(big.NewInt(0)) == 0 && executedId.Cmp(big.NewInt(0)) == 0 {
		return nil, errors.New("contract commitment is empty")
	}
	eventId := new(big.Int).Add(executedId, big.NewInt(1))
	commitment, err := receiver.GetCommitmentByStateSyncId(eventId)
	if err != nil {
		s.logger.Warn("Get commitment state sync id failed", "err", err)
		return nil, err
	}
	var events []*sync.StateSender
	var proofs [][]common.Hash
	for {
		event, err := s.l1Sync.SyncDB().GetStateSenderEvent(eventId)
		if event == nil || err != nil {
			break
		}
		proof, err := s.eventProofDb.GetProof(commitment.Root, eventId)
		if err != nil {
			s.logger.Warn("Get proof failed", "root", commitment.Root, "eventid", eventId, "err", err)
			break
		}
		events = append(events, event)
		proofs = append(proofs, proof)
		eventId = eventId.Add(eventId, big.NewInt(1))
		s.logger.Debug("Get executed event", "eventid", eventId)
	}
	exTxs, err := s.createExecuteTxs(ctx, proofs, events, nonce)
	if err != nil {
		return nil, err
	}
	return exTxs, nil
}
