package upgrade

import (
	"encoding/json"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/election"
	"github.com/PlatONnetwork/AppChain-SDK/example/oracle"
	oracle2 "github.com/PlatONnetwork/AppChain-SDK/example/upgrade/oracle"
	"github.com/PlatONnetwork/AppChain-SDK/store/memorydb"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/AppChain-SDK/x/extravote"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade"
	"github.com/PlatONnetwork/AppChain-SDK/x/upgrade/types"

	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/eth"
	"github.com/PlatONnetwork/PlatON-Go/node"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/stretchr/testify/require"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestM(t *testing.T) {
	store := memorydb.New()
	oracleModule := oracle2.NewModule(store, testutil.DefaultAccount[0].NodePrivateKey(), oracle.NewTimerRate(time.Second*2))
	extraVote := extravote.NewExtraVote(store, []extravote.ExtraVerifier{oracleModule})
	oracleModule.SetExtraVote(extraVote)
	vals := election.NewModule()
	upgradeModule := upgrade.NewModule(store)
	upgradeOracleModule := NewModule()
	manager := module.NewManager(vals, extraVote, upgradeModule, upgradeOracleModule)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(vals.Name(), extraVote.Name(), upgradeModule.Name(), upgradeOracleModule.Name())
	manager.SetConsensusExtend(extraVote.Name())
	manager.RegisterUpgradeHandler(upgradeModule)
	manager.SetModuleValidChecker(upgradeModule.IsModuleValid)
	upgradeModule.SetIsContractModule(manager.IsContractModule)
	oracleModule.RegistryUpgradeHandler(upgradeModule)
	app := testutil.NewApp(manager)

	config := election.GenesisConfig{
		InitialNodes: election.Nodes{
			election.Node{
				Name:        "node1",
				Owner:       testutil.DefaultAccount[0].NodeAddress(),
				Desc:        "node1",
				PublicKey:   crypto.FromECDSAPub(&testutil.DefaultAccount[0].NodePrivateKey().PublicKey),
				BlsPubKey:   testutil.DefaultAccount[0].BlsSecretKey().GetPublicKey().Serialize(),
				HostAddress: "127.0.0.1",
				RpcPort:     uint16(testutil.DefaultAccount[0].HTTP),
				P2pPort:     uint16(testutil.DefaultAccount[0].P2PPort),
			},
		},
		AdminAddress:     common2.UserAddrs[0],
		EpochSize:        250,
		ElectionDistance: 20,
	}
	s, _ := json.Marshal(config)
	upgradeConfig, _ := json.Marshal(&types.GenesisConfig{
		ModuleGenesisConfig: module.ModuleGenesisConfig{},
		Owner:               testutil.DefaultAccount[2].CoinBaseAddress(),
	})
	var stack []*node.Node
	var backend []*eth.Ethereum
	var err error
	if stack, backend, err = testutil.CreateCluster(testutil.DefaultAccount[0:1], testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		vals.Name():          s,
		upgradeModule.Name(): upgradeConfig,
	}, common2.UserAddrs); err != nil {
		require.Nil(t, err)
	}

	stack[0].Start()
	backend[0].Start()
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
}
