package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
	"gopkg.in/urfave/cli.v1"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	outputFlag = cli.StringFlag{
		Name:   "output",
		Usage:  "output tps report",
		EnvVar: "BENCHMARK_OUTPUT",
	}
	serverFlag = cli.StringFlag{Name: "web", EnvVar: "BENCHMARK_SERVER", Usage: "web server address"}
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
	addr := ctx.String(serverFlag.Name)
	if len(addr) != 0 {
		server := &Server{
			clis: clis,
			addr: addr,
		}
		server.start()
	}
	return nil
}

func generateReport(clis []*Connection, start, end uint64) ([][]string, error) {
	var report [][]string
	startTime := uint64(0)
	txTotal := int64(0)
	lastTxLength := 0
	for i := start; i <= end; i++ {
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

type Server struct {
	clis []*Connection
	addr string
}

func (s *Server) start() {
	fmt.Println("Start server:", s.addr)
	s.route()
	http.ListenAndServe(s.addr, nil)
}
func (s *Server) route() {
	http.HandleFunc("/line", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		start, err := strconv.Atoi(query.Get("start"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		end, err := strconv.Atoi(query.Get("end"))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		report, err := generateReport(s.clis, uint64(start), uint64(end))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		var data [4][]string
		for _, r := range report {
			data[0] = append(data[0], r[0])
			data[1] = append(data[1], r[1])
			data[2] = append(data[2], r[2])
			data[3] = append(data[3], r[3])
		}
		b, err := GeneratePhotoBytes(data[0], data[1], data[2], data[3])
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		w.Write(b)
	})
}
func GeneratePhotoBytes(timestamp []string, txs []string, tps []string, latency []string) ([]byte, error) {
	// create a new line instance
	line := charts.NewLine()

	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{
			Theme: types.ThemeInfographic,
		}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Line chart in Go",
			Subtitle: "",
		}),
		charts.WithGridOpts(opts.Grid{Width: "800", Height: "400"}),
		charts.WithYAxisOpts(opts.YAxis{
			SplitLine: &opts.SplitLine{
				Show: opts.Bool(true),
			},
		}),
	)
	// Put data into instance
	line.SetXAxis(timestamp)
	makeLindeData := func(v []string) []opts.LineData {
		var items []opts.LineData
		for _, i := range v {
			items = append(items, opts.LineData{Value: i})
		}
		return items
	}
	line.AddSeries("transactions", makeLindeData(txs), charts.WithLabelOpts(
		opts.Label{Show: opts.Bool(true)},
	))

	line.AddSeries("tps", makeLindeData(tps), charts.WithLabelOpts(
		opts.Label{Show: opts.Bool(true)},
	))

	line.AddSeries("latency", makeLindeData(latency), charts.WithLabelOpts(
		opts.Label{Show: opts.Bool(true)},
	))

	buffer := bytes.NewBuffer([]byte{})
	if err := line.Render(buffer); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
