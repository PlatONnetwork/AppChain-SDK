package checkpoint

import "gopkg.in/urfave/cli.v1"

var (
	RootchainNodeRPCFlag = cli.StringFlag{
		Name:  "checkpoint.rootchain-node-rpc",
		Usage: "Rootchain node RPC endpoint",
	}
	KeystoreFlag = cli.StringFlag{
		Name:  "checkpoint.keystore",
		Usage: "Keystore for signing checkpoint transaction",
	}
	PasswordFlag = cli.StringFlag{
		Name:  "checkpoint.password",
		Usage: "",
	}
)
