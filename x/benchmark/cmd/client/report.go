package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"gopkg.in/urfave/cli.v1"
	"os"
	"time"
)

var (
	outputFlag = cli.StringFlag{
		Name:   "output",
		Usage:  "output tps report",
		EnvVar: "BENCHMARK_OUTPUT",
	}
)

func Report(ctx *cli.Context) error {
	clis, err := createClient(ctx)
	if err != nil {
		return err
	}
	output := ctx.String(outputFlag.Name)
	if len(output) != 0 {
		start := ctx.Uint64(startBlockFlag.Name)
		end := ctx.Uint64(endBlockFlag.Name)
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		data, err := generateReport(clis, start, end)
		err = csv.NewWriter(file).WriteAll(data)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateReport(clis []*Connection, start, end uint64) ([][]string, error) {
	report := [][]string{
		[]string{
			"number", "time", "tps", "interval",
		},
	}
	startTime := uint64(0)
	txTotal := int64(0)
	lastTxLength := 0
	lastProduceTime := uint64(0)
	for i := start; i <= end; i++ {
		beginTime := time.Now()
		var info benchmark.BlockInfo
		for _, cli := range clis {
			f, err := cli.GetBlockState(context.Background(), i)
			if err != nil {
				return nil, err
			}
			info.Number = f.Number
			info.ProduceTime = f.ProduceTime
			info.TotalLength = f.TotalLength
			info.TimeUse += f.TimeUse
			info.TxLength += f.TxLength
		}
		fmt.Println("get block finish", info.Number, "cost", time.Since(beginTime))

		if i != start {
			tps := float64(txTotal) / (float64(info.ProduceTime-startTime) / float64(1000))
			interval := float64(info.ProduceTime-lastProduceTime) / float64(1000)
			report = append(report, []string{
				fmt.Sprintf("%d", i),
				time.UnixMilli(int64(info.ProduceTime)).Format(time.RFC3339Nano),
				fmt.Sprintf("%d", info.TotalLength),
				fmt.Sprintf("%.3f", tps),
				fmt.Sprintf("%.3f", interval),
			})
		} else {
			startTime = info.ProduceTime
			report = append(report, []string{
				fmt.Sprintf("%d", i),
				time.UnixMilli(int64(info.ProduceTime)).Format(time.RFC3339Nano),
				fmt.Sprintf("%d", lastTxLength),
				fmt.Sprintf("%d", 0),
				fmt.Sprintf("%d", 0),
			})
		}
		lastTxLength = info.TotalLength
		lastProduceTime = info.ProduceTime
		txTotal += int64(info.TotalLength)

	}
	return report, nil
}
