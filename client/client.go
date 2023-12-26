package client

import (
	"context"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
)

// client defines typed wrapped for the AppChain RPC API.
type Client struct {
	*ethclient.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	ethcli, err := ethclient.Dial(rawurl)
	if err != nil {
		return nil, err
	}
	return &Client{
		Client: ethcli,
	}, nil
}

type ProofMetadata struct {
	LeafIndex       uint64
	ExitEvent       string
	CheckpointBlock uint64
}
type ExitProof struct {
	Data     []common.Hash
	Metadata ProofMetadata
}

func (c *Client) GenerateExitProof(ctx context.Context, exitID uint64) (*ExitProof, error) {
	var proof ExitProof
	err := c.RpcClient().CallContext(ctx, &proof, "checkpoint_generateExitProof", exitID)
	if err != nil {
		return nil, err
	}
	return &proof, nil
}
