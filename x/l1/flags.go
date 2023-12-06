package l1

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/crypto"

	"gopkg.in/urfave/cli.v1"
)

func DecodeNodePrivateKey(ctx *cli.Context) *ecdsa.PrivateKey {
	var (
		hex  = ctx.GlobalString(utils.NodeKeyHexFlag.Name)
		file = ctx.GlobalString(utils.NodeKeyFileFlag.Name)
		key  *ecdsa.PrivateKey
		err  error
	)
	switch {
	case file != "" && hex != "":
		utils.Fatalf("Options %q and %q are mutually exclusive", utils.NodeKeyFileFlag.Name, utils.NodeKeyHexFlag.Name)
	case file != "":
		if key, err = crypto.LoadECDSA(file); err != nil {
			utils.Fatalf("Option %q: %v", utils.NodeKeyFileFlag.Name, err)
		}

	case hex != "":
		if key, err = crypto.HexToECDSA(hex); err != nil {
			utils.Fatalf("Option %q: %v", utils.NodeKeyHexFlag.Name, err)
		}

	}
	return key
}
