package nontxpool

import (
	"github.com/PlatONnetwork/PlatON-Go/core/types"

	sdkp2p "github.com/PlatONnetwork/AppChain-SDK/p2p"
)

type TransactionsMessage struct {
	sdkp2p.MessageCode
	Txs []*types.Transaction
}
