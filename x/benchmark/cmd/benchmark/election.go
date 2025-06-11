package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

var ElectionAddress = common.HexToAddress("0x1300000000000000000000000000000000000001")
var key = []byte("initnode")
var maxValidNumber = uint64(100000000000)

type Election struct {
	validator cbfttypes.Validators
	coinbase  *ecdsa.PrivateKey
}

func NewElection() *Election {
	return &Election{}
}
func (e *Election) SetCoinbase(coinbase *ecdsa.PrivateKey) {
	e.coinbase = coinbase
}
func (e Election) Name() string {
	return "election"
}

func (e Election) Version() uint64 {
	return 0
}
func (e *Election) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	data, err := json.Marshal(chainConfig.Cbft.InitialNodes)
	if err != nil {
		return err
	}
	db.SetState(ElectionAddress, key, data)
	return nil
}
func (e *Election) Init(ctx sdk.InitContext) error {
	return nil
}
func (e Election) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	if e.coinbase != nil {
		header.Coinbase = crypto.PubkeyToAddress(e.coinbase.PublicKey)
	}
	return nil
}

func (e Election) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return maxValidNumber
}

func (e Election) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	statedb, err := ctx.Backend().State()
	if err != nil {
		return nil, err
	}
	data := statedb.GetState(ElectionAddress, key)
	var nodes []params.CbftNode
	err = json.Unmarshal(data, &nodes)
	if err != nil {
		return nil, err
	}
	var validator cbfttypes.Validators
	nodeMap := make(cbfttypes.ValidateNodeMap)
	for i, n := range nodes {
		var blsPubKey bls.PublicKey
		blsPubKey.Deserialize(n.BlsPubKey.Serialize())
		nodeMap[n.Node.ID()] = &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   crypto.PubkeyToNodeAddress(*n.Node.Pubkey()),
			PubKey:    n.Node.Pubkey(),
			NodeID:    n.Node.ID(),
			BlsPubKey: &blsPubKey,
			Shares:    big.NewInt(1),
		}
	}
	validator.Nodes = nodeMap
	validator.ValidBlockNumber = maxValidNumber
	return &validator, nil
}

func (e Election) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return false
}
