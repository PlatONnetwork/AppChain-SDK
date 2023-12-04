package l2

import (
	"fmt"

	"github.com/PlatONnetwork/AppChain-SDK/common"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"gopkg.in/urfave/cli.v1"
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
