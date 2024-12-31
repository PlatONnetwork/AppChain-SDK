package testutil

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Validator struct {
	validator *cbfttypes.Validators
	coinbase  *ecdsa.PrivateKey
}

func NewValidator(accounts []*Account) (*Validator, error) {
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
		node.IDv0()
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
	return &Validator{validator: validators}, nil
}
func (v *Validator) SetCoinbase(coinbase *ecdsa.PrivateKey) {
	v.coinbase = coinbase
}
func (v Validator) Name() string {
	return "validator"
}

func (v Validator) Version() uint64 {
	return 0
}

func (v Validator) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	if v.coinbase != nil {
		header.Coinbase = crypto.PubkeyToAddress(v.coinbase.PublicKey)
	}
	return nil
}

func (v Validator) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	return 10000000
}

func (v Validator) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	return v.validator, nil
}

func (v Validator) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	//TODO implement me
	panic("implement me")
}
