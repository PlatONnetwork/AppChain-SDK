package deploy

import "gopkg.in/urfave/cli.v1"

var (
	LogLevelFlag = cli.IntFlag{Name: "loglevel", Usage: "Log level, 1:error,2:warn,3:info,4:debug", Value: 4}

	OutputFlag = cli.StringFlag{
		Name:   "output",
		Usage:  "Output result directory",
		EnvVar: "OUTPUT",
		Value:  ".",
	}
	GenesisNodeFlag = cli.StringFlag{
		Name:   "genesisnode",
		Usage:  "Chain genesis node ip, separate with commas",
		EnvVar: "GENESIS_NODE",
	}
	NormalNodeFlag = cli.StringFlag{
		Name:   "node",
		Usage:  "Chain normal node ip, separate with commas",
		EnvVar: "NORMAL_NODE",
	}
	NodeConfigFileFlag = cli.StringFlag{
		Name:   "nodeconfigfile",
		Usage:  "Chain config file name",
		EnvVar: "NODE_CONFIG_FILE",
		Value:  "nodes.toml",
	}
	NodePProfPortFlag = cli.IntFlag{
		Name:   "nodepprofport",
		Usage:  "Chain pprof port",
		EnvVar: "NODE_PPROF_PORT",
		Value:  7701,
	}

	NodeRpcPortFlag = cli.IntFlag{
		Name:   "noderpcport",
		Usage:  "Chain rpc port",
		EnvVar: "NODE_RPC_PORT",
		Value:  8801,
	}

	NodeP2pPortFlag = cli.IntFlag{
		Name:   "nodep2pport",
		Usage:  "Chain p2p port",
		EnvVar: "NODE_P2P_PORT",
		Value:  18001,
	}
	NodePortIncFlag = cli.BoolFlag{
		Name:   "nodeportinc",
		Usage:  "Chain port increment",
		EnvVar: "NODE_PORT_INC",
	}
	GenesisFileFlag = cli.StringFlag{
		Name:   "genesisfile",
		Usage:  "Chain genesis file name",
		EnvVar: "GENESIS_FILE",
		Value:  "genesis.json",
	}
	AnsibleDirFlag = cli.StringFlag{
		Name:   "ansibledir",
		Usage:  "Remote ansible directory",
		EnvVar: "ANSIBLE_DIR",
		Value:  "ansible",
	}
	RemoteAnsibleDirFlag = cli.StringFlag{
		Name:   "remoteansibledir",
		Usage:  "Remote ansible directory",
		EnvVar: "REMOTE_ANSIBLE_DIR",
		Value:  "~/opt",
	}

	BinFlag = cli.StringFlag{
		Name:   "bin",
		Usage:  "Binary path",
		EnvVar: "BIN",
		Value:  "./node",
	}

	UsernameFlag = cli.StringFlag{
		Name:   "username",
		Usage:  "Remote server ssh username",
		EnvVar: "USERNAME",
		Value:  "",
	}
	PasswordFlag = cli.StringFlag{
		Name:   "password",
		Usage:  "Remote server ssh password",
		EnvVar: "PASSWORD",
		Value:  "",
	}

	PProfFlag = cli.IntFlag{
		Name:   "pprof",
		Usage:  "Pprof address, empty is disable, 0:close, 1:local, 2:server, 3:any",
		EnvVar: "PPROF",
		Value:  3,
	}
	HttpFlag = cli.IntFlag{
		Name:   "http",
		Usage:  "Http address, empty is disable 0:close, 1:local, 2:server, 3:any",
		EnvVar: "HTTP",
		Value:  3,
	}
	StartArgsFlag = cli.StringFlag{
		Name:   "startargs",
		Usage:  "Command start args",
		EnvVar: "START_ARGS",
		Value:  "--identity node --verbosity 4 --http.api platon,debug,personal,admin,net,web3,txpool,benchmark --http.vhosts \"*\" --cache 256 --metrics --ipcdisable  --maxpeers 100 --maxconsensuspeers 75 --txpool.globalslots 1000000 --txpool.accountslots 1000000 --txpool.globalqueue 1000000 --txpool.accountqueue 1000000 --txpool.cacheSize 1000000 --txpool.globaltxcount 1000000 --nodiscover  --networkid 102 --allow-insecure-unlock",
	}
)
