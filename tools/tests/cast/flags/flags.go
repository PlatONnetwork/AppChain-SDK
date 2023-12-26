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
	StartFlags = cli.Uint64Flag{
		Name:   "start",
		Usage:  "start block",
		Hidden: true,
		Value:  0,
	}
	EndFlags = cli.Uint64Flag{
		Name:   "end",
		Usage:  "end block",
		Hidden: true,
		Value:  0,
	}
)
