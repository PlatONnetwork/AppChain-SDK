package consensusnetwork

import (
	"github.com/PlatONnetwork/PlatON-Go/consensus/cbft/protocols"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"sort"
	"sync"
	"time"

	"gopkg.in/urfave/cli.v1"

	"github.com/PlatONnetwork/AppChain-SDK/x/consensusnetwork/network"
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	ModuleName    = "consensusnetwork"
	ModuleVersion = 0
)

type ValidatorNodeList struct {
	sync.Mutex
	ids []enode.ID
}

func (vn *ValidatorNodeList) Update(validators []*cbfttypes.ValidateNode) {
	vn.Lock()
	defer vn.Unlock()
	var sortedValidatorNode cbfttypes.SortedValidatorNode
	for _, v := range validators {
		sortedValidatorNode = append(sortedValidatorNode, v)
	}
	sort.Sort(sortedValidatorNode)
	ids := make([]enode.ID, 0, len(validators))
	for _, v := range sortedValidatorNode {
		ids = append(ids, v.NodeID)
	}
	vn.ids = ids
}
func (vn *ValidatorNodeList) Len() int {
	vn.Lock()
	defer vn.Unlock()
	return len(vn.ids)
}
func (vn *ValidatorNodeList) Get(index int) enode.ID {
	vn.Lock()
	defer vn.Unlock()
	if len(vn.ids) > index {

		return vn.ids[index]
	}
	return enode.ID{}
}

type Module struct {
	logger     log.Logger
	network    *network.EngineManager
	engine     network.ConsensusNetworkEngine
	validators ValidatorNodeList
}

func NewModule(ctx *cli.Context) *Module {
	return &Module{
		logger:  log.New("module", ModuleName),
		network: network.NewEngineManger(),
	}
}

func (m *Module) Init(ctx sdk.InitContext) error {
	//consensusEngine := ctx.Backend().ChainContext().Engine()
	//networkEngine, ok := consensusEngine.(network.ConsensusNetworkEngine)
	//if ok {
	//	m.network.SetEngine(networkEngine)
	//} else {
	//	panic("Illegal consensus engine")
	//}
	return nil
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Version() uint64 {
	return ModuleVersion
}
func (m *Module) CreateConsensusNetworkEngine(ctx sdk.ConsensusNetworkContext) (sdk.ConsensusNetworkEngine, error) {
	m.engine = ctx
	m.network.SetEngine(ctx)
	return m, nil
}

func (m *Module) Protocols() []p2p.Protocol {
	return m.network.Protocols()
}

func (m *Module) StartNetworkEngine() {
	m.network.Start()
}
func (m *Module) Publish(message ctypes.Message) {
	switch msg := message.(type) {
	case *protocols.PrepareBlock, *protocols.PrepareHeader:
		m.logger.Debug("Publish consensus msg", "msg", msg.String())
		m.network.Broadcast(message)
	case *protocols.PrepareVote:
		m.logger.Debug("Publish vote msg", "msg", msg.String())

		leader := m.validators.Get(int(msg.ViewNumber) % m.validators.Len()).TerminalString()
		m.logger.Debug("Publish vote msg", "leader", leader)

		m.network.Send(leader, msg)
	}
}
func (m *Module) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {
	m.validators.Update(validators)
}
func (m *Module) Broadcast(msg ctypes.Message) {
	m.network.Broadcast(msg)
}
func (m *Module) PartBroadcast(msg ctypes.Message) {
	m.network.PartBroadcast(msg)
}
func (m *Module) Forwarding(nodeID string, msg ctypes.Message) error {
	return m.network.Forwarding(nodeID, msg)
}
func (m *Module) Send(peerID string, msg ctypes.Message) {
	m.network.Send(peerID, msg)
}
func (m *Module) AvgLatency() time.Duration {
	return m.network.AvgLatency()
}
func (m *Module) PeerSetting(peerID string, bType uint64, blockNumber uint64) error {
	return m.network.PeerSetting(peerID, bType, blockNumber)
}
func (m *Module) RemovePeer(id string) {
	m.network.RemovePeer(id)
}
