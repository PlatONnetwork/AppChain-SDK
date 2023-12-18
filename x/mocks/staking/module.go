package staking

import (
	"encoding/json"
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/store"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	ModuleName   = "staking"
	ValidatorKey = "validator"

	NumberBlocksOfEpoch = 250
)

type ValidatorNode struct {
	Index     uint32
	Address   common.NodeAddress
	Node      []byte
	BlsPubKey []byte
}

type ValidatorNodes struct {
	Nodes []ValidatorNode
}

type Module struct {
	store      store.KVStore
	validators *cbfttypes.Validators
}

func NewModule(store store.Store) *Module {
	return &Module{
		store: store.GetKVStore(ModuleName),
	}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) IsEndOfRound(ctx sdk.ConsensusContext, blockNumber uint64) bool {
	return blockNumber%NumberBlocksOfEpoch == 0
}

func (m *Module) BlocksOfRound(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return NumberBlocksOfEpoch
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	nodes := convertToValidatorNodes(chainConfig.Cbft.InitialNodes)
	val, err := rlp.EncodeToBytes(nodes)
	if err != nil {
		return fmt.Errorf("validators rlp error: %v", err)
	}

	m.store.Set([]byte(ValidatorKey), val)
	return nil
}

func (m *Module) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	return nil
}

func (m *Module) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	var lastBlockNumber uint64
	if blockNumber <= NumberBlocksOfEpoch {
		lastBlockNumber = NumberBlocksOfEpoch
	} else {
		vds, err := m.GetRoundValidator(ctx, blockNumber)
		if err != nil {
			log.Error("Get validator fail", "blockNumber", blockNumber)
			return 0
		}

		if vds.ValidBlockNumber == 0 && blockNumber%NumberBlocksOfEpoch == 0 {
			return blockNumber
		}

		// lastNumber = vds.ValidBlockNumber + ia.blocksPerNode * vds.Len() - 1
		lastBlockNumber = vds.ValidBlockNumber + NumberBlocksOfEpoch - 1

		// May be `CurrentValidators ` had not updated, so we need to calcuate `lastBlockNumber`
		// via `blockNumber`.
		if lastBlockNumber < blockNumber {
			blocksPerRound := uint64(NumberBlocksOfEpoch)
			if blockNumber%blocksPerRound == 0 {
				lastBlockNumber = blockNumber
			} else {
				baseNum := blockNumber - (blockNumber % blocksPerRound)
				lastBlockNumber = baseNum + blocksPerRound
			}
		}
	}
	//log.Debug("Get last block number", "blockNumber", blockNumber, "lastBlockNumber", lastBlockNumber)
	return lastBlockNumber
}

func (m *Module) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return m.GetRoundValidator(ctx, blockNumber)
}

func (m *Module) GetRoundValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	if m.validators == nil {
		val, err := m.store.Get([]byte(ValidatorKey))
		if err != nil {
			return nil, err
		}

		var nodes ValidatorNodes
		if err := rlp.DecodeBytes(val, &nodes); err != nil {
			return nil, err
		}

		m.validators = newValidators(&nodes, 1)
	}
	baseNumber := blockNumber
	if blockNumber == 0 {
		baseNumber = 1
	}
	m.validators.ValidBlockNumber = ((baseNumber-1)/uint64(NumberBlocksOfEpoch))*NumberBlocksOfEpoch + 1
	return m.validators, nil
}

func (m *Module) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return false
}

func newValidators(nodes *ValidatorNodes, validBlockNumber uint64) *cbfttypes.Validators {
	vds := &cbfttypes.Validators{
		Nodes:            make(cbfttypes.ValidateNodeMap, len(nodes.Nodes)),
		ValidBlockNumber: validBlockNumber,
	}

	for i, node := range nodes.Nodes {
		var p2pNode enode.Node

		if err := p2pNode.UnmarshalText(node.Node); err != nil {
			panic(err)
		}

		var blsPubKey bls.PublicKey
		if err := blsPubKey.Deserialize(node.BlsPubKey); err != nil {
			panic(err)
		}

		log.Info("validator", "index", node.Index, "enode", p2pNode.String())

		vds.Nodes[p2pNode.ID()] = &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   node.Address,
			PubKey:    p2pNode.Pubkey(),
			NodeID:    p2pNode.ID(),
			BlsPubKey: &blsPubKey,
		}
	}
	return vds
}

func convertToValidatorNodes(nodes []params.CbftNode) *ValidatorNodes {
	var vds ValidatorNodes

	for i, node := range nodes {
		pubkey := node.Node.Pubkey()
		if pubkey == nil {
			panic("pubkey should not nil")
		}

		blsPubKey := node.BlsPubKey

		nodeBuf, err := node.Node.MarshalText()
		if err != nil {
			panic(err)
		}

		vds.Nodes = append(vds.Nodes, ValidatorNode{
			Index:     uint32(i),
			Address:   crypto.PubkeyToNodeAddress(*pubkey),
			Node:      nodeBuf,
			BlsPubKey: blsPubKey.Serialize(),
		})
	}
	return &vds
}
