package statesync

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"gopkg.in/urfave/cli.v1"
)

var (
	KeystoreFlag = cli.StringFlag{
		Name:  "statesync.keystore",
		Usage: "Keystore for signing statesync transaction",
	}
	PasswordFlag = cli.StringFlag{
		Name:   "statesync.password",
		Usage:  "Password for keystore",
		EnvVar: "STATESYNC_PASSWORD",
	}
	StartBlockFlag = cli.Uint64Flag{
		Name:  "statesync.startblock",
		Usage: "Starting block height for scanning",
	}
)

func AddModuleInitFlags(app *cli.App) {
	app.Flags = append(app.Flags, KeystoreFlag)
	app.Flags = append(app.Flags, PasswordFlag)
	app.Flags = append(app.Flags, StartBlockFlag)

}

func decodePrivateKey(ctx *cli.Context) (*keystore.Key, error) {
	if !ctx.GlobalIsSet(KeystoreFlag.Name) {
		return nil, fmt.Errorf("statesync.keystore not set")
	}
	if !ctx.GlobalIsSet(PasswordFlag.Name) {
		return nil, fmt.Errorf("statesync.password not set")

	}
	ksFile := ctx.GlobalString(KeystoreFlag.Name)
	ksPaswordFile := ctx.GlobalString(PasswordFlag.Name)
	key, err := common.DecryptKey(ksFile, ksPaswordFile)
	if err != nil {
		return nil, err
	}
	return key, nil
}
