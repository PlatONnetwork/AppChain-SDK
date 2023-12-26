package transaction

import (
	"context"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"time"
)

func WaitTx(client *ethclient.Client, hash common.Hash) (*types.Receipt, error) {
	for i := 0; i < 5; i++ {
		receipt, _ := client.TransactionReceipt(context.Background(), hash)
		if receipt != nil {
			return receipt, nil
		}
		time.Sleep(time.Second)
	}
	return nil, nil
}
