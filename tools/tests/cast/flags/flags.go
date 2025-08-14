package flags

import "gopkg.in/urfave/cli.v1"

var (
	RootchainRPCFlags = cli.StringFlag{
		Name:   "rootchainrpc",
		Usage:  "rpc url",
		EnvVar: "ROOTCHAIN_RPC_URL",
	}
	AppchainRPCFlag = cli.StringFlag{
		Name:   "rpc",
		Usage:  "AppChain RPC endpoint",
		EnvVar: "SDK_RPC_URL",
	}
	RoundFlags = cli.Uint64Flag{
		Name:  "round",
		Usage: "round, stage roundSize",
	}
	EpochFlags = cli.Uint64Flag{
		Name:  "epoch",
		Usage: "epoch, stage epochSize",
	}
	KeyFlags = cli.StringFlag{
		Name:   "key",
		Usage:  "private key used to send tx",
		EnvVar: "SDK_KEY",
	}
	NonceFlags = cli.Uint64Flag{
		Name:   "nonce",
		Usage:  "transaction nonce",
		EnvVar: "NONCE",
	}
	GasLimitFlags = cli.Uint64Flag{
		Name:   "gas-limit",
		Usage:  "transaction gas limit",
		EnvVar: "GASLIMIT",
		Value:  2000000,
	}
	GasPriceFlags = cli.Uint64Flag{
		Name:   "gas-price",
		Usage:  "transaction gas price",
		EnvVar: "GASPRICE",
		Value:  2000000000,
	}
	ModuleFlags = cli.StringFlag{
		Name:  "module",
		Usage: "method name",
	}
	MethodFlags = cli.StringFlag{
		Name:  "method",
		Usage: "method name",
	}
	AbiFileFlags = cli.StringFlag{
		Name:   "abifile",
		Usage:  "abi file",
		EnvVar: "ABI_FILE",
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
