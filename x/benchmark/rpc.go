package benchmark

import (
	"context"

	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
	"time"
)

type Status struct {
	Sent              uint64  `json:"sent"`
	Tps               float64 `json:"tps"`
	CacheTx           int     `json:"cacheTx"`
	RawTxPercent      int     `json:"rawTxPercent"`
	ContractTxPercent int     `json:"contractTxPercent"`
}
type BlockInfo struct {
	ProduceTime uint64 `json:"id_time"`
	Number      uint64 `json:"block"`
	TxLength    int    `json:"tx_length"`
	TimeUse     uint64 `json:"time_use"`
}
type RPC struct {
	m *Module
}

func NewRPC(m *Module) *RPC {
	return &RPC{m: m}
}

func (r *RPC) GenTxs(accountBeginIndex, accountEndIndex, rawTxPercent, contractTxPercent int, totalTx uint64) error {
	r.m.Lock()
	defer r.m.Unlock()
	r.m.startIndex = accountBeginIndex
	r.m.endIndex = accountEndIndex
	r.m.rawTxPercent = rawTxPercent
	r.m.contractTxPercent = contractTxPercent
	return r.m.createTransactions(totalTx)
}
func (r *RPC) Start(tps uint64, sendTxPool bool) error {
	return r.m.start(tps, sendTxPool)
}
func (r *RPC) Stop() error {
	return r.m.stop()
}
func (r *RPC) Status() (*Status, error) {
	now := time.Now().Unix()
	r.m.Lock()
	defer r.m.Unlock()
	total := 0
	for _, v := range r.m.txCache {
		total += len(v)
	}
	tps := float64(r.m.confirm.Load()) / float64(now-r.m.Statistics.start.Unix())
	return &Status{
		Sent:              r.m.send.Load(),
		Tps:               tps,
		CacheTx:           total,
		RawTxPercent:      r.m.rawTxPercent,
		ContractTxPercent: r.m.contractTxPercent,
	}, nil
}
func (r *RPC) GetBlockState(blockNumber uint64) (*BlockInfo, error) {
	return r.m.db.Get(blockNumber)
}

type Client struct {
	*ethclient.Client
	rpc *rpc.Client
}

func NewClient(rawurl string) (*Client, error) {
	c, err := rpc.DialContext(context.Background(), rawurl)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: ethclient.NewClient(c),
		rpc:    c,
	}, nil
}

func (c *Client) GenTxs(ctx context.Context, accountBeginIndex, accountEndIndex, rawTxPercent, contractTxPercent int, totalTx uint64) error {
	return c.rpc.CallContext(ctx, nil, "benchmark_genTxs", accountBeginIndex, accountEndIndex, rawTxPercent, contractTxPercent, totalTx)
}

func (c *Client) Start(ctx context.Context, tps uint64, sendTxPool bool) error {
	err := c.rpc.CallContext(ctx, nil, "benchmark_start", tps, sendTxPool)
	return err
}

func (c *Client) Stop(ctx context.Context) error {
	err := c.rpc.CallContext(ctx, nil, "benchmark_stop")
	return err
}
func (c *Client) Status(ctx context.Context) (*Status, error) {
	var res Status
	err := c.rpc.CallContext(ctx, &res, "benchmark_status")
	return &res, err
}

func (c *Client) GetBlockState(ctx context.Context, blockNumber uint64) (*BlockInfo, error) {
	var res BlockInfo
	err := c.rpc.CallContext(ctx, &res, "benchmark_getBlockState", blockNumber)
	return &res, err
}
