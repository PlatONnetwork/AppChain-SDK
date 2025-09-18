package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
	"time"
)

func createTransactionOpts(sk *ecdsa.PrivateKey, chainid *big.Int, gaslimit uint64, gasprice *big.Int) *bind.TransactOpts {
	return &bind.TransactOpts{
		From: crypto.PubkeyToAddress(sk.PublicKey),
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, types.NewLondonSigner(chainid), sk)
		},
		Value:    big.NewInt(0),
		GasPrice: gasprice,
		GasLimit: gaslimit,
	}
}

func waitTx(cli *ethclient.Client, hash common.Hash) *types.Receipt {
	log.Debug("Wait tx", "hash", hash.Hex())
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*60))
	var receipt *types.Receipt
	var err error
	for {
		receipt, err = cli.TransactionReceipt(ctx, hash)
		if receipt == nil {
			time.Sleep(time.Millisecond * 500)
			continue
		}
		if err != nil {
			log.Error("Wait tx failed", "err", err, "hash", hash.Hex())
		}
		break
	}
	return receipt
}

func deploy(fn func() (common.Address, *types.Transaction, error), client *ethclient.Client, msg string) (common.Address, error) {
	addr, tx, err := fn()
	if err != nil {
		log.Error(fmt.Sprintf("Deploy %s failed", msg), "err", err)
		return common.Address{}, err
	}
	log.Debug(fmt.Sprintf("Deploy %s", msg), "tx", tx.Hash().Hex())
	receipt := waitTx(client, tx.Hash())
	if receipt == nil || receipt.Status == 0x00 {
		reason := "receipt is nil"
		if receipt.Status == 0x00 {
			err := callTx(client, tx)
			if err != nil {
				reason = err.Error()
			}
		}
		log.Error(fmt.Sprintf("Deploy %s", msg), "reason", reason)
		return common.Address{}, fmt.Errorf("Deploy %s failed, receipt is nil or status is false", msg)
	}
	log.Debug(fmt.Sprintf("Deploy %s success", msg), "addr", addr.Hex())
	return addr, nil
}

func callTx(client *ethclient.Client, tx *types.Transaction) error {
	chainid, _ := client.ChainID(context.Background())
	_, err := client.CallContract(context.Background(), platon.CallMsg{
		From:       tx.FromAddr(types.NewLondonSigner(chainid)),
		To:         tx.To(),
		Gas:        tx.Gas(),
		GasPrice:   tx.GasPrice(),
		GasFeeCap:  tx.GasFeeCap(),
		GasTipCap:  tx.GasTipCap(),
		Value:      tx.Value(),
		Data:       tx.Data(),
		AccessList: tx.AccessList(),
	}, nil)
	return err
}
func send(fn func() (*types.Transaction, error), client *ethclient.Client, msg string) (*types.Receipt, error) {
	tx, err := fn()
	if err != nil {
		log.Error("Send failed", "err", err)
		return nil, err
	}
	receipt := waitTx(client, tx.Hash())
	if receipt == nil {
		return nil, fmt.Errorf("Send %s failed, receipt is nil", msg)
	}
	if receipt.Status == 0x00 {
		err = callTx(client, tx)
		log.Debug(fmt.Sprintf("Send %s failed", msg), "err", err)
		return nil, err
	}
	log.Debug(fmt.Sprintf("Send %s success", msg), "hash", tx.Hash().Hex())
	return receipt, err
}
