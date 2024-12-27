package coupon

import (
	"context"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
)

type Priority interface {
	Set([]common.Address)
	Get() []common.Address
}

type RPC struct {
	p Priority
}

func NewRPC(priority Priority) *RPC {
	return &RPC{p: priority}
}

func (r *RPC) SetPriority(address []common.Address) error {
	r.p.Set(address)
	return nil
}

func (r *RPC) GetPriority() ([]common.Address, error) {
	return r.p.Get(), nil
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

func (c *Client) SetPriority(ctx context.Context, address []common.Address) error {
	return c.rpc.CallContext(ctx, nil, "coupon_setPriority", address)
}

func (c *Client) GetPriority(ctx context.Context) ([]common.Address, error) {
	var address []common.Address
	err := c.rpc.CallContext(ctx, &address, "coupon_getPriority")
	return address, err
}
