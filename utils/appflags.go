package utils

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"gopkg.in/urfave/cli.v1"
	"strings"
)

var (
	KeystoreFlag = cli.StringFlag{
		Name:  "l2.keystore",
		Usage: "Keystore for signing l2 transaction",
	}
	PasswordFlag = cli.StringFlag{
		Name:   "l2.password",
		Usage:  "Password for keystore",
		EnvVar: "L2_PASSWORD",
	}
)

func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, KeystoreFlag)
	app.Flags = append(app.Flags, PasswordFlag)
}

func DecodePrivateKey(keystoreFile, passwordFile string) (*keystore.Key, error) {
	if keystoreFile == "" {
		return nil, fmt.Errorf("l2.keystore not set")
	}
	if passwordFile == "" {
		return nil, fmt.Errorf("l2.password not set")

	}

	key, err := common.DecryptKey(keystoreFile, passwordFile)
	if err != nil {
		return nil, err
	}
	return key, nil
}

func SetBootNodes(ctx *cli.Context, mainnetBootnodes, testnetBootnodes []string) {
	if !ctx.GlobalIsSet(utils.BootnodesFlag.Name) && !ctx.GlobalIsSet(utils.BootnodesV4Flag.Name) {
		if ctx.GlobalBool(utils.TestnetFlag.Name) && testnetBootnodes != nil {
			ctx.GlobalSet(utils.BootnodesV4Flag.Name, strings.Join(testnetBootnodes, ","))
		} else if mainnetBootnodes != nil {
			ctx.GlobalSet(utils.BootnodesV4Flag.Name, strings.Join(mainnetBootnodes, ","))
		}
	}
}
