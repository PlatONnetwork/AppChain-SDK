package testutil

import (
	"encoding/hex"
	"encoding/json"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type MockApp struct {
	validator *cbfttypes.Validators
}

func NewMockApp(accounts []*Account) (*MockApp, error) {
	validators := &cbfttypes.Validators{}
	nodeMap := make(cbfttypes.ValidateNodeMap)
	for index, acc := range accounts {
		var blsKey bls.SecretKey
		keyBuf, _ := hex.DecodeString(acc.BlsKey)
		blsKey.Deserialize(keyBuf)
		nodePrivate, err := crypto.HexToECDSA(acc.NodeKey)
		if err != nil {
			return nil, err
		}
		node := enode.NewV4(&nodePrivate.PublicKey, acc.Host, acc.P2PPort, acc.P2PPort)
		nodeAddr := crypto.PubkeyToAddress(nodePrivate.PublicKey)
		nodeMap[node.ID()] = &cbfttypes.ValidateNode{
			Index:     uint32(index),
			Address:   common.NodeAddress(nodeAddr),
			PubKey:    &nodePrivate.PublicKey,
			NodeID:    node.ID(),
			BlsPubKey: blsKey.GetPublicKey(),
		}
	}
	validators.Nodes = nodeMap
	validators.ValidBlockNumber = 10000000
	return &MockApp{validator: validators}, nil
}

func (m *MockApp) Start() error {
	return nil
}

func (m *MockApp) Stop() error {
	return nil
}

func (m *MockApp) Contracts(statedb sdk.StateDBReader, blockNumber uint64) []sdk.SDKContract {
	return []sdk.SDKContract{}
}

func (m *MockApp) CheckTx(ctx sdk.Context, tx *types.Transaction) error {
	return nil
}

func (m *MockApp) FilterPendingTxs(ctx sdk.Context, txs map[common.Address]types.Transactions) map[common.Address]types.Transactions {
	return txs
}

func (m *MockApp) APIs() []rpc.API {

	return []rpc.API{}
}

func (m *MockApp) Protocols() []p2p.Protocol {

	return []p2p.Protocol{}
}

func (m *MockApp) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	return nil
}

func (m *MockApp) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return 10000000
}

func (m *MockApp) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return m.validator, nil
}

func (m *MockApp) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return true
}

func (m *MockApp) OnCommit(ctx sdk.ConsensusContext, block *types.Block) error {
	return nil
}

func (m *MockApp) ExtendData(ctx sdk.ConsensusContext) []byte {
	return nil
}

func (m *MockApp) VerifyExtendData(ctx sdk.ConsensusContext, data []byte) (common.Hash, error) {
	return common.Hash{}, nil
}

func (m *MockApp) PrepareQC(ctx sdk.ConsensusContext, block *protocols.PrepareBlock, votes map[uint32]*protocols.PrepareVote) {

}

func (m *MockApp) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {

}

func (m *MockApp) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data map[string]json.RawMessage) error {
	return nil
}

func (m *MockApp) BeginBlock(ctx sdk.WorkerContext) error {
	return nil
}

func (m *MockApp) EndBlock(ctx sdk.WorkerContext) error {
	return nil
}

func (m *MockApp) AddTxs(ctx sdk.WorkerContext) (types.Transactions, error) {
	return []*types.Transaction{}, nil
}

func (m *MockApp) SortTxs(ctx sdk.WorkerContext, local map[common.Address]types.Transactions, remote map[common.Address]types.Transactions) (types.Transactions, error) {
	return []*types.Transaction{}, nil
}

func (m *MockApp) InitChain(ctx sdk.InitContext) error {

	return nil
}
