package consensus

import (
	"gopkg.in/urfave/cli.v1"

	"github.com/PlatONnetwork/AppChain-SDK/x/consensus/network"
	"github.com/PlatONnetwork/PlatON-Go/core/cbfttypes"
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

func (s *ConsensusNetworkModule) ViewChange(ctx sdk.ConsensusContext, validators []*cbfttypes.ValidateNode) {
	// TODO
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
