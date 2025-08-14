package checkpoint

import (
	"context"
	"github.com/PlatONnetwork/AppChain-SDK/x/checkpoint/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
)

type RpcService struct {
	checkpoint *Module
}

func NewRpcService(m *Module) *RpcService {
	return &RpcService{checkpoint: m}
}

func (s *RpcService) GenerateExitProof(exitID uint64) (types.Proof, error) {
	return s.checkpoint.generateExitProof(exitID)
}

type Client struct {
	*ethclient.Client
}

func NewClient(client *ethclient.Client) *Client {
	return &Client{
		Client: client,
	}
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
