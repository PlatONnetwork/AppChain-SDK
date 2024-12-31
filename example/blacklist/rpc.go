package blacklist

import (
	"context"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/rpc"
)

type Blacklist interface {
	Set([]common.Address)
	Get() []common.Address
}

type RPC struct {
	p Blacklist
}

func NewRPC(blacklist Blacklist) *RPC {
	return &RPC{p: blacklist}
}

func (r *RPC) SetBlacklist(address []common.Address) error {
	r.p.Set(address)
	return nil
}

func (r *RPC) GetBlacklist() ([]common.Address, error) {
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

func (c *Client) SetBlacklist(ctx context.Context, address []common.Address) error {
	return c.rpc.CallContext(ctx, nil, "blacklist_setBlacklist", address)
}

func (c *Client) GetBlacklist(ctx context.Context) ([]common.Address, error) {
	var address []common.Address
	err := c.rpc.CallContext(ctx, &address, "blacklist_getBlacklist")
	return address, err
}
