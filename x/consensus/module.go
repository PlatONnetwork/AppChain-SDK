package consensus

import (
	"time"

	"gopkg.in/urfave/cli.v1"

	"github.com/PlatONnetwork/AppChain-SDK/x/consensus/network"
	ctypes "github.com/PlatONnetwork/PlatON-Go/consensus/cbft/types"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/PlatONnetwork/PlatON-Go/p2p"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

const (
	ModuleName    = "consensusNetwork"
	ModuleVersion = 0
)

type ConsensusNetworkModule struct {
	logger  log.Logger
	network *network.EngineManager
}

func NewModule(ctx *cli.Context) *ConsensusNetworkModule {
	return &ConsensusNetworkModule{
		logger:  log.New("module", ModuleName),
		network: network.NewEngineManger(),
	}
}

func (s *ConsensusNetworkModule) Init(ctx sdk.InitContext) error {
	consensusEngine := ctx.Backend().ChainContext().Engine()
	networkEngine, ok := consensusEngine.(network.ConsensusNetworkEngine)
	if ok {
		s.network.SetEngine(networkEngine)
	} else {
		panic("Illegal consensus engine")
	}
	return nil
}

func (s *ConsensusNetworkModule) Name() string {
	return ModuleName
}

func (s *ConsensusNetworkModule) Version() uint64 {
	return ModuleVersion
}

func (s *ConsensusNetworkModule) Protocols() []p2p.Protocol {
	return s.network.Protocols()
}

func (s *ConsensusNetworkModule) StartNetworkEngine() {
	s.network.Start()
}
func (s *ConsensusNetworkModule) Broadcast(msg ctypes.Message) {
	s.network.Broadcast(msg)
}
func (s *ConsensusNetworkModule) PartBroadcast(msg ctypes.Message) {
	s.network.PartBroadcast(msg)
}
func (s *ConsensusNetworkModule) Forwarding(nodeID string, msg ctypes.Message) error {
	return s.network.Forwarding(nodeID, msg)
}
func (s *ConsensusNetworkModule) Send(peerID string, msg ctypes.Message) {
	s.network.Send(peerID, msg)
}
func (s *ConsensusNetworkModule) AvgLatency() time.Duration {
	return s.network.AvgLatency()
}
func (s *ConsensusNetworkModule) PeerSetting(peerID string, bType uint64, blockNumber uint64) error {
	return s.network.PeerSetting(peerID, bType, blockNumber)
}
func (s *ConsensusNetworkModule) RemovePeer(id string) {
	s.network.RemovePeer(id)
}
