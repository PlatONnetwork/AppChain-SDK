package common

import (
	"context"
	"errors"
	platon "github.com/PlatONnetwork/PlatON-Go"
	types2 "github.com/PlatONnetwork/PlatON-Go/core/types"
	"time"
)

func WaitTx(cli platon.TransactionReader, tx *types2.Transaction) error {
	var receipt *types2.Receipt
	var err error
	for i := 0; i < 5; i++ {
		receipt, err = cli.TransactionReceipt(context.Background(), tx.Hash())
		if receipt != nil {
			break
		}
		time.Sleep(2 * time.Second)
	}
	if receipt == nil || err != nil {
		return errors.New("get receipt failed")
	}
	if receipt.Status == types2.ReceiptStatusFailed {
		return errors.New("receipt status failed")
	}
	return nil
}
