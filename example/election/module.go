package election

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/contracts/proxy"
	"github.com/PlatONnetwork/AppChain-SDK/example/election/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

const ModuleVersion uint64 = 0
const ModuleName = "election"

var ProxyAddress = common.BigToAddress(big.NewInt(101))
var ElectionAddress = common.BigToAddress(big.NewInt(102))
var CallerAddress = common.BigToAddress(big.NewInt(133))

type Nodes []Node

func (ns Nodes) toContractNode() []contracts.Node {
	var res []contracts.Node
	for _, n := range ns {
		res = append(res, contracts.Node{
			Name:        n.Name,
			Owner:       n.Owner,
			Desc:        n.Desc,
			PublicKey:   n.PublicKey,
			BlsPubKey:   n.BlsPubKey,
			HostAddress: n.HostAddress,
			RpcPort:     n.RpcPort,
			P2pPort:     n.P2pPort,
		})
	}
	return res
}

type Node struct {
	Name        string
	Owner       common.Address
	Desc        string
	PublicKey   hexutil.Bytes
	BlsPubKey   hexutil.Bytes
	HostAddress string
	RpcPort     uint16
	P2pPort     uint16
}

type GenesisConfig struct {
	module.ModuleGenesisConfig
	InitialNodes     Nodes          `json:"initialNodes"`
	AdminAddress     common.Address `json:"adminAddress"`
	EpochSize        uint64         `json:"epochSize"`
	ElectionDistance uint64         `json:"electionDistance"`
}

type Module struct {
	chainConfig    *params.ChainConfig
	nodePrivateKey *ecdsa.PrivateKey
}

func NewModule() *Module {
	return &Module{}
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}

func (m *Module) Init(ctx sdk.InitContext) error {
	m.chainConfig = ctx.Backend().ChainConfig()
	m.nodePrivateKey = ctx.NodeKey()
	return nil
}

func (m *Module) InitGenesis(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig, data json.RawMessage) error {
	var genesis GenesisConfig
	if err := json.Unmarshal(data, &genesis); err != nil {
		return err
	}

	proxyCaller, err := proxy.NewInitializableTransparentUpgradeableProxyGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
	if err = proxyCaller.WithCaller(genesis.AdminAddress).WithTo(ProxyAddress).DeployInitializableTransparentUpgradeableProxy(); err != nil {
		return err
	}
	electionCaller, err := contracts.NewElectionGenesisCaller(ctx, db, chainConfig)
	if err != nil {
		return err
	}
	if err = electionCaller.WithCaller(genesis.AdminAddress).WithTo(ElectionAddress).DeployElection(); err != nil {
		return err
	}

	input, _ := electionCaller.PackInitialize(genesis.InitialNodes.toContractNode(), genesis.EpochSize, genesis.ElectionDistance)
	if err = proxyCaller.WithCaller(genesis.AdminAddress).WithTo(ProxyAddress).Initialize(ElectionAddress, genesis.AdminAddress, input); err != nil {
		return err
	}
	return nil
}

func (m *Module) EndBlock(ctx sdk.WorkerContext) error {
	caller, err := contracts.NewElectionBackendBackendCaller(ProxyAddress)
	if err != nil {
		return err
	}
	return caller.WithCaller(CallerAddress).ChangeEpoch(ctx)
}

func (m *Module) NewHeader(ctx sdk.ConsensusContext, header *types.Header) error {
	if ctx.IsProposer() {
		currentValidatorPubKey := hex.EncodeToString(crypto.FromECDSAPub(&m.nodePrivateKey.PublicKey))
		caller, _ := contracts.NewElectionGenesisCaller(ctx, types.NewStateDBWrapper(ctx.ParentStateDB()), m.chainConfig)
		nodeInfo, _ := caller.WithCaller(CallerAddress).WithTo(ProxyAddress).GetByPubKey(currentValidatorPubKey)
		if len(nodeInfo.Name) == 0 {
			return errors.New("not found validator")
		}
		header.Coinbase = nodeInfo.Owner
	}
	return nil
}

func (m *Module) GetLastNumber(ctx sdk.ConsensusContext, blockNumber uint64) uint64 {
	caller, _ := contracts.NewElectionGenesisCaller(ctx, types.NewStateDBWrapper(ctx.ParentStateDB()), m.chainConfig)
	epochOfBlocks, _ := caller.WithCaller(CallerAddress).WithTo(ProxyAddress).EpochSize()
	return (blockNumber + epochOfBlocks - 1) / epochOfBlocks * epochOfBlocks
}

func (m *Module) GetValidator(ctx sdk.ConsensusContext, blockNumber uint64) (*cbfttypes.Validators, error) {
	caller, _ := contracts.NewElectionGenesisCaller(ctx, types.NewStateDBWrapper(ctx.ParentStateDB()), m.chainConfig)
	caller = caller.WithCaller(CallerAddress).WithTo(ProxyAddress)
	rvn, _ := caller.GetCurrentRoundValidator()
	if rvn.End.Uint64() < blockNumber {
		rvn, _ = caller.GetNextRoundValidator()
	}
	vnm := make(cbfttypes.ValidateNodeMap, len(rvn.Nodes))
	for i, name := range rvn.Nodes {
		vd, _ := caller.GetByName(name)
		var blsKey bls.PublicKey
		pubKey, err := crypto.UnmarshalPubkey(vd.PublicKey)
		if err != nil {
			return nil, err
		}
		err = blsKey.Deserialize(vd.BlsPubKey)
		if err != nil {
			return nil, err
		}
		nodeId := enode.PubkeyToIDV4(pubKey)
		addr := crypto.PubkeyToNodeAddress(*pubKey)
		vn := &cbfttypes.ValidateNode{
			Index:     uint32(i),
			Address:   addr,
			PubKey:    pubKey,
			BlsPubKey: &blsKey,
			NodeID:    nodeId,
		}
		vnm[nodeId] = vn
	}

	cvd := &cbfttypes.Validators{
		Nodes:            vnm,
		ValidBlockNumber: rvn.Start.Uint64(),
	}
	return cvd, nil
}

func (m *Module) IsCandidateNode(ctx sdk.ConsensusContext, nodeID enode.IDv0) bool {
	return false
}
