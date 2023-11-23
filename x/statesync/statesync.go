package statesync

import (
	"errors"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/merkle"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/statesync/sync"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	types2 "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func (s *StateSync) ExtendDataImpl(ctx sdk.Context, epoch, view uint64, index uint32, header *types.Header) []byte {
	receiver, err := s.newStateSyncCallContract(ctx, header)
	if err != nil {
		log.Error("Failed to new state sync call contract", "err", err)
		return nil
	}
	syncId, err := receiver.GetStateSyncId()
	if err != nil {
		return nil
	}
	commitment, err := receiver.GetCommitmentByStateSyncId(syncId)
	if err != nil {
		return nil
	}
	start := new(big.Int).Add(commitment.EndId, big.NewInt(1))
	end := s.MaxSyncId()
	if start.Cmp(end) > 0 {
		return nil
	}
	root, err := s.GenProof(epoch, view, index, start, end)
	if err != nil {
		return nil
	}

	raw, err := rlp.EncodeToBytes(&contracts.StateSyncCommitment{StartId: start, EndId: end, Root: root})
	if err != nil {
		return nil
	}
	return raw
}

func (s *StateSync) VerifyExtendDataImpl(epoch, view uint64, index uint32, header *types.Header, data []byte) (common.Hash, error) {
	if len(data) == 0 {
		return crypto.Keccak256Hash(nil), nil
	}

	var commitment contracts.StateSyncCommitment
	rlp.DecodeBytes(data, &commitment)
	root, err := s.GenProof(epoch, view, index, commitment.StartId, commitment.EndId)
	if err != nil {
		return common.Hash{}, err
	}
	if root != commitment.Root {
		return common.Hash{}, errors.New("state sync root hash is invalid")
	}
	return crypto.Keccak256Hash(data), nil
}

func (s *StateSync) PrepareQCImpl(block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {
	//TODO clear history event proofs
	root := s.eventProofDb.GetProofRoot(block.Epoch, block.ViewNumber, block.BlockIndex)
	if root == common.ZeroHash {
		return
	}

	s.eventProofDb.InsertRootBlock(root, block.Block.Hash())
}

func (s *StateSync) MaxSyncId() *big.Int {
	//TODO 获取验证人列表
	quorumId := s.p2p.GetQuorumSyncId(nil)
	id, _ := s.l1Sync.SyncDB().GetMaxSyncId()
	if id.Cmp(quorumId) > 0 {
		return quorumId
	}
	return id
}

func (s *StateSync) GenProof(epoch, view uint64, index uint32, start, end *big.Int) (common.Hash, error) {
	if root := s.eventProofDb.GetProofRoot(epoch, view, index); root != common.ZeroHash {
		return root, nil
	}
	events, err := s.l1Sync.SyncDB().FindStateSenderEvent(start, end)
	if err != nil {
		return common.Hash{}, nil
	}
	leafId := make(map[*big.Int]common.Hash)
	var trieNodes [][]byte
	for _, event := range events {
		value, err := rlp.EncodeToBytes(event)
		if err != nil {
			return common.Hash{}, nil
		}
		hash := crypto.Keccak256Hash(value)
		trieNodes = append(trieNodes, hash.Bytes())
		leafId[event.Id] = hash
	}
	trie, err := merkle.NewMerkleTree(trieNodes)
	if err := s.eventProofDb.InsertProof(epoch, view, index, start, end, leafId, trie); err != nil {
		return common.Hash{}, nil
	}
	return trie.Hash(), nil
}

func (s *StateSync) createCommitTx(ctx sdk.Context, cm *contracts.StateSyncCommitment, index uint64, qc *types2.QuorumCert, voteProof []common.Hash, nonce uint64) (*types.Transaction, error) {
	input, err := contracts.Abi.Methods["commit"].Inputs.Pack(cm, index, voteProof, &contracts.QuorumCert{
		Epoch:       qc.Epoch,
		ViewNumber:  qc.ViewNumber,
		BlockHash:   qc.BlockHash,
		BlockNumber: qc.BlockNumber,
		BlockIndex:  qc.BlockIndex,
		ExtendHash:  qc.ExtendHash,
	}, qc.Signature[:], qc.ValidatorSet.Bits, qc.ValidatorSet.Bytes())
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(nonce, contracts.StateSyncAddress, nil, 100000, big.NewInt(0), input)
	chainId, _ := ctx.Backend().ChainId()
	signer := types.NewEIP155Signer(chainId)
	tx, err = types.SignTx(tx, signer, s.privateKey)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *StateSync) createExecuteTxs(ctx sdk.Context, proofs [][]common.Hash, events []*sync.StateSender, nonce uint64) ([]*types.Transaction, error) {
	var txs []*types.Transaction
	for i, proof := range proofs {
		input, err := contracts.Abi.Methods["execute"].Inputs.Pack(proof, events[i])
		if err != nil {
			return nil, err
		}
		tx := types.NewTransaction(nonce, contracts.StateSyncAddress, nil, 100000, big.NewInt(0), input)
		chainId, _ := ctx.Backend().ChainId()
		signer := types.NewEIP155Signer(chainId)
		tx, err = types.SignTx(tx, signer, s.privateKey)
		if err != nil {
			return nil, err
		}
		txs = append(txs, tx)
	}
	return txs, nil
}

func (s *StateSync) newStateSyncCallContract(ctx sdk.Context, header *types.Header) (*contracts.StateReceiver, error) {
	from := crypto.PubkeyToAddress(s.privateKey.PublicKey)
	evm, _, err := ctx.Backend().GetEVM(NewOnlyCallMessage(from), header)
	if err != nil {
		return nil, err
	}
	return contracts.NewStateReceiver(evm, vm.NewContract(vm.AccountRef(from), vm.AccountRef(contracts.StateSyncAddress), big.NewInt(0), 1000000), true)
}
