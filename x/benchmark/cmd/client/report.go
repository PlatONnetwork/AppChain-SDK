package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"gopkg.in/urfave/cli.v1"
	"math/big"
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
	var report [][]string
	startTime := uint64(0)
	txTotal := int64(0)
	lastTxLength := 0
	for i := start; i <= end; i++ {
		beginTime := time.Now()
		block, err := clis[0].BlockByNumber(context.Background(), new(big.Int).SetUint64(i))
		if err != nil {
			return nil, err
		}
		var info benchmark.BlockInfo
		info.Number = block.NumberU64()
		info.ProduceTime = block.Time()
		for _, cli := range clis {
			f, err := cli.GetBlockState(context.Background(), i)
			if err != nil {
				return nil, err
			}
			if f.Number != block.NumberU64() || f.ProduceTime != block.Time() {
				return nil, fmt.Errorf("get block info failed, expect:%d,%d, actual:%d,%d", block.NumberU64(), block.Time(), f.Number, f.ProduceTime)
			}

			info.TimeUse += f.TimeUse
			info.TxLength += f.TxLength
		}
		fmt.Println("get block finish", info.Number, "cost", time.Since(beginTime))

		if i != start {
			tps := float64(txTotal) / (float64(block.Time()-startTime) / float64(1000))
			latency := float64(0)
			if info.TxLength != 0 {
				latency = (float64(info.TimeUse) / float64(1000)) / float64(info.TxLength)
			}
			report = append(report, []string{
				time.UnixMilli(int64(info.ProduceTime)).Format("2006-01-02 15:04:05.123"),
				fmt.Sprintf("%d", lastTxLength),
				fmt.Sprintf("%.3f", tps),
				fmt.Sprintf("%.3f", latency),
			})
		} else {
			startTime = block.Time()
		}
		lastTxLength = block.Transactions().Len()
		txTotal += int64(block.Transactions().Len())

	}
	return report, nil
}
