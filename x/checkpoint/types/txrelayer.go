package types

import (
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type TxRelayer interface {
	// Init initializing TxRelayer module
	Init(sdk.InitContext) error
	// Call executes a message call immediately without creating a transaction on the blockchain
	Call(from common.Address, to common.Address, data []byte) ([]byte, error)
	// SendTransaction signs given transaction by provided key and sends it to the blockchain
	SendTransaction(txn *types.Transaction, key *keystore.Key) (*types.Receipt, error)
	// Client returns platon client
	Client() *ethclient.Client
}
