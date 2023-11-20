package checkpoint

import "gopkg.in/urfave/cli.v1"

var (
	KeystoreFlag = cli.StringFlag{
		Name:  "checkpoint.keystore",
		Usage: "Keystore for signing checkpoint transaction",
	}
	PasswordFlag = cli.StringFlag{
		Name:  "checkpoint.password",
		Usage: "Password for keystore",
		EnvVar: "CHECKPOINT_PASSWORD",
	}
)
