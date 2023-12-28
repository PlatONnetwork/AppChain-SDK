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
		Usage: "send,call,logs,abi",
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
	CommandHelpTemplate = `NAME:
   {{.HelpName}} - {{.Usage}}

USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

CATEGORY:
   {{.Category}}{{end}}{{if .Description}}

DESCRIPTION:
   {{.Description}}{{end}}{{if .VisibleFlags}}

OPTIONS:
   {{range .Flags}}{{.}}
   {{end}}{{end}}
`
)
