package flags

import "gopkg.in/urfave/cli.v1"

var (
	RPCFlags = cli.StringFlag{
		Name:   "rpc",
		Usage:  "rpc url",
		EnvVar: "SDK_RPC_URL",
	}
	KeyFlags = cli.StringFlag{
		Name:   "key",
		Usage:  "private key used to send tx",
		EnvVar: "SDK_KEY",
	}
	MethodFlags = cli.StringFlag{
		Name:  "method",
		Usage: "method name",
	}
	AddressFlags = cli.StringFlag{
		Name:  "address",
		Usage: "contract address",
	}
	TypeFlags = cli.StringFlag{
		Name:  "type",
		Usage: "send or call",
	}
)
