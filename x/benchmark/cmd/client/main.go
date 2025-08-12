package main

import (
	"context"
	"errors"
	"fmt"
	deploytools "github.com/PlatONnetwork/AppChain-SDK/tools/deploy"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"gopkg.in/urfave/cli.v1"
	"os"
	"strings"
	"sync"
	"time"
)

var (
	urlsFlag          = cli.StringFlag{Name: "urls", EnvVar: "BENCHMARK_URLS", Usage: "Node http rpc url, example:\"127.0.0.1:8801,127.0.0.2:8802\""}
	addrsFlag         = cli.IntFlag{Name: "addrs", EnvVar: "BENCHMARK_ADDRS", Usage: "Server ip, example:\"127.0.0.1,127.0.0.2\""}
	rawTxPercentFlag  = cli.IntFlag{Name: "rawtx", EnvVar: "BENCHMARK_RAWTX", Value: 100, Usage: "The proportion of transactions to the total number of transactions"}
	countFlag         = cli.Uint64Flag{Name: "count", EnvVar: "BENCHMARK_COUNT", Value: 10000, Usage: "Generate transaction count"}
	tpsFlag           = cli.Uint64Flag{Name: "tps", EnvVar: "BENCHMARK_TPS", Value: 1000, Usage: "Sent txs per second"}
	txsPerAccountFlag = cli.IntFlag{Name: "txsperaccount", EnvVar: "BENCHMARK_TXSPERACCOUNT", Value: 256, Usage: "Sent txs per account"}
	sendTxPoolFlag    = cli.BoolFlag{Name: "txpool", Usage: "Transactions send to the txpool"}
	startBlockFlag    = cli.Uint64Flag{Name: "start", EnvVar: "BENCHMARK_STARTBLOCK", Usage: "Start block number"}
	endBlockFlag      = cli.Uint64Flag{Name: "end", EnvVar: "BENCHMARK_ENDBLOCK", Usage: "End block number"}
	GenTxCommand      = cli.Command{
		Name:   "gentx",
		Action: GenTx,
		Flags: []cli.Flag{
			urlsFlag,
			addrsFlag,
			rawTxPercentFlag,
			countFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	StartCommand = cli.Command{
		Name:   "start",
		Action: Start,
		Flags: []cli.Flag{
			urlsFlag,
			tpsFlag,
			txsPerAccountFlag,
			sendTxPoolFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	StopCommand = cli.Command{
		Name:   "stop",
		Action: Stop,
		Flags: []cli.Flag{
			urlsFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	StatusCommand = cli.Command{
		Name:   "status",
		Action: Status,
		Flags: []cli.Flag{
			urlsFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	ReportCommand = cli.Command{
		Name:   "report",
		Action: Report,
		Flags: []cli.Flag{
			urlsFlag,
			startBlockFlag,
			endBlockFlag,
			outputFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
	GenerateCommand = cli.Command{
		Name:   "generate",
		Action: Generate,
		Flags: []cli.Flag{
			hostsFlag,
			deploytools.OutputFlag,
			deploytools.AnsibleDirFlag,
			binFlag,
			passwordFlag,
			userFlag,
			deploytools.StartArgsFlag,
			extraArgsFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func main() {
	app := deploytools.CreateApp()
	app.Commands = []cli.Command{
		GenTxCommand,
		StartCommand,
		StopCommand,
		StatusCommand,
		ReportCommand,
		deploytools.CreateAnsibleCommand,
		GenerateCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func GenTx(ctx *cli.Context) error {
	clis, err := createClient(ctx)
	if err != nil {
		return err
	}
	addrs := ctx.Int(addrsFlag.Name)
	rawTxPercent := ctx.Int(rawTxPercentFlag.Name)
	contractTxPercent := 100 - rawTxPercent
	amount := ctx.Uint64(countFlag.Name)
	var s sync.WaitGroup
	s.Add(len(clis))
	for i, cli := range clis {
		fmt.Println(cli.url, i*addrs, (i+1)*addrs-1, rawTxPercent, contractTxPercent, amount)
		go func(cli *Connection, start, end int) {
			ctx, _ := context.WithTimeout(context.Background(), time.Hour)
			err = cli.GenTxs(ctx, start, end, rawTxPercent, contractTxPercent, amount)
			if err != nil {
				log.Error("gen txs failed", "url", cli.url, "err", err)
			}
			fmt.Println(cli.url, "gen txs success")
			s.Done()
		}(cli, i*addrs, (i+1)*addrs-1)
	}
	s.Wait()
	return nil
}
func Start(ctx *cli.Context) error {
	clis, err := createClient(ctx)
	if err != nil {
		return err
	}
	tps := ctx.Uint64(tpsFlag.Name)
	txsPerAccount := ctx.Int(txsPerAccountFlag.Name)
	sendTxPool := ctx.Bool(sendTxPoolFlag.Name)
	for _, cli := range clis {
		err = cli.Start(context.Background(), tps, txsPerAccount, sendTxPool)
		if err != nil {
			log.Error("start failed", "url", cli.url, "err", err)
		}
		fmt.Println(fmt.Sprintf("[%s]", cli.url), "start success")
	}
	return nil
}
func Stop(ctx *cli.Context) error {
	clis, err := createClient(ctx)
	if err != nil {
		return err
	}
	for _, cli := range clis {
		err = cli.Stop(context.Background())
		if err != nil {
			log.Error("stop failed", "url", cli.url, "err", err)
		}
		fmt.Println(fmt.Sprintf("[%s]", cli.url), "stop success")

	}
	return nil
}
func Status(ctx *cli.Context) error {
	clis, err := createClient(ctx)
	if err != nil {
		return err
	}
	for _, cli := range clis {
		status, err := cli.Status(context.Background())
		if err != nil {
			log.Error("stop failed", "url", cli.url, "err", err)
			continue
		}
		fmt.Println(fmt.Sprintf("[%s]", cli.url),
			"sent", status.Sent, "tps", status.Tps, "cache", status.CacheTx,
			"rawTxPercent", fmt.Sprintf("%d%%", status.RawTxPercent),
			"contractTxPercent", fmt.Sprintf("%d%%", status.ContractTxPercent))
	}
	return nil
}

type Connection struct {
	*benchmark.Client
	url string
}

func createClient(ctx *cli.Context) ([]*Connection, error) {
	urls := strings.Split(ctx.String(urlsFlag.Name), ",")
	if len(urls) == 0 {
		return nil, errors.New("url is empty")
	}
	var clis []*Connection
	for _, url := range urls {
		cli, err := benchmark.NewClient(url)
		if err != nil {
			return nil, errors.Join(err, errors.New(url))
		}
		clis = append(clis, &Connection{cli, url})
	}
	return clis, nil
}
