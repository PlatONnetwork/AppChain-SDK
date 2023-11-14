package types

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
)

type Signer interface {
	SignTx(txn *types.Transaction) (*types.Transaction, error)
}

type TxRelayer interface {
	// Call executes a message call immediately without creating a transaction on the blockchain
	Call(from common.Address, to common.Address, data []byte) ([]byte, error)
	// SendTransaction signs given transaction by provided key and sends it to the blockchain
	SendTransaction(txn *types.Transaction, signer Signer) (*types.Receipt, error)
	// Client returns platon client
	Client() *ethclient.Client
}
