package deploy

import (
	"github.com/PlatONnetwork/PlatON-Go/log"
	"gopkg.in/urfave/cli.v1"
	"os"
)

func initLog(lv int) {
	glogger := log.NewGlogHandler(log.StreamHandler(os.Stdout, log.TerminalFormat(false)))
	glogger.Verbosity(log.Lvl(lv))
	log.Root().SetHandler(glogger)
}
func CreateApp() *cli.App {
	app := cli.NewApp()
	app.HideVersion = true
	app.Commands = []cli.Command{}
	app.Before = func(ctx *cli.Context) error {
		initLog(ctx.Int(LogLevelFlag.Name))
		return nil
	}
	app.Flags = []cli.Flag{
		LogLevelFlag,
	}
	return app
}
