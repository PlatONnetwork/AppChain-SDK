package txrelayer

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/log"
)

const (
	ModuleName = "txrelayer"

	DefaultRPCAddress     = "http://127.0.0.1:6789"
	DefaultReceiptTimeout = 50 * time.Millisecond
	DefaultNumRetries     = 1000
)

var (
	_ module.Module = (*Module)(nil)
)

type Signer interface {
	SignTx(txn *types.Transaction) (*types.Transaction, error)
	Address() common.Address
}

type Module struct {
	rpcAddress     string
	client         *ethclient.Client
	receiptTimeout time.Duration
	numRetries     int

	lock sync.Mutex
}

func NewModule(rpcAddress string, receiptTimeout time.Duration, numRetries int) (*Module, error) {
	m := &Module{
		rpcAddress:     rpcAddress,
		receiptTimeout: receiptTimeout,
		numRetries:     numRetries,
	}

	client, err := ethclient.Dial(rpcAddress)
	if err != nil {
		return nil, err
	}

	m.client = client
	return m, nil
}

func (m *Module) Name() string {
	return ModuleName
}

func (m *Module) Call(from common.Address, to common.Address, data []byte) ([]byte, error) {
	result, err := m.client.CallContract(context.Background(), platon.CallMsg{
		From: from,
		To:   &to,
		Data: data,
	}, big.NewInt(-1))
	return result, err
}

func (m *Module) SendTransaction(txn *types.Transaction, signer Signer) (*types.Receipt, error) {
	txnHash, err := m.sendTransactionLocked(txn, signer)
	if err != nil {
		return nil, err
	}
	return m.waitForReceipt(txnHash)
}

func (m *Module) sendTransactionLocked(txn *types.Transaction, signer Signer) (common.Hash, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	var (
		err      error
		nonce    uint64
		gasPrice *big.Int
		gasLimit uint64
	)

	nonce = txn.Nonce()
	gasPrice = txn.GasPrice()
	gasLimit = txn.Gas()

	if nonce == 0 {
		nonce, err = m.client.NonceAt(context.Background(), signer.Address(), big.NewInt(-1))
		if err != nil {
			return common.ZeroHash, err
		}
	}
	if gasPrice == nil || gasPrice.Uint64() == 0 {
		gasPrice, err = m.client.SuggestGasPrice(context.Background())
		if err != nil {
			return common.ZeroHash, err
		}
	}
	if gasLimit == 0 {
		gasLimit, err = m.client.EstimateGas(context.Background(), platon.CallMsg{
			From: signer.Address(),
			To:   txn.To(),
			Data: txn.Data(),
		})
		if err != nil {
			return common.ZeroHash, err
		}
	}

	newTxn := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       txn.To(),
		Value:    txn.Value(),
		Data:     txn.Data()})
	signedTxn, err := signer.SignTx(newTxn)
	if err != nil {
		return common.ZeroHash, err
	}

	log.Info("Send transaction", "module", m.Name(), "from", signer.Address(), "to", signedTxn.To(), "gasPrice", gasPrice, "gasLimit", gasLimit, "txnHash", signedTxn.Hash())
	if err = m.client.SendTransaction(context.Background(), signedTxn); err != nil {
		return common.ZeroHash, err
	}
	return signedTxn.Hash(), nil
}

func (m *Module) waitForReceipt(hash common.Hash) (*types.Receipt, error) {
	if m.numRetries < 0 {
		return nil, nil
	}

	for count := 0; count < m.numRetries; count++ {
		receipt, err := m.client.TransactionReceipt(context.Background(), hash)
		if err != nil {
			if err.Error() != "not found" {
				return nil, err
			}
		}

		if receipt != nil {
			return receipt, nil
		}

		time.Sleep(m.receiptTimeout)
	}

	return nil, fmt.Errorf("timeout while waiting for transaction %s to be processed", hash)
}

func (m *Module) Client() *ethclient.Client {
	return m.client
}
