package nontxpool

import (
	"context"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
)

type RPC struct {
	m *Module
}

func NewRPC(m *Module) *RPC {
	return &RPC{m: m}
}

func (r *RPC) PendingNonce(addr common.Address) (uint64, error) {
	return r.m.PendingNonce(addr), nil
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

func (c *Client) PendingNonce(ctx context.Context, addr common.Address) (uint64, error) {
	var res uint64
	err := c.rpc.CallContext(ctx, &res, "nontxpool_pengingNonce", addr)
	return res, err
}
