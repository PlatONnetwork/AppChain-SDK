package benchmark

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark/contracts"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"math/big"
)

var (
	gasPrice    = big.NewInt(1000000000)
	tokenAbi, _ = contracts.BenchTokenMetaData.GetAbi()
)

func createRawTransfer(signer types.Signer, from *ecdsa.PrivateKey, to common.Address, nonce uint64) (*types.Transaction, error) {
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &to,
		Value:    big.NewInt(1),
		Gas:      200000,
		GasPrice: gasPrice,
		Data:     nil,
	})
	return types.SignTx(tx, signer, from)
}

func createTokenTransfer(signer types.Signer, from *ecdsa.PrivateKey, contract common.Address, to common.Address, nonce uint64) (*types.Transaction, error) {
	input, _ := tokenAbi.Pack("transfer", to, big.NewInt(1))
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &contract,
		Value:    big.NewInt(0),
		Gas:      200000,
		GasPrice: gasPrice,
		Data:     input,
	})
	return types.SignTx(tx, signer, from)
}
