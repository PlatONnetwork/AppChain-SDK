package common

import (
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"os"

	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
)

func DecryptKey(ksFile, pwdFile string) (*keystore.Key, error) {
	json, err := os.ReadFile(ksFile)
	if err != nil {
		return nil, err
	}
	passphrase, err := os.ReadFile(pwdFile)
	if err != nil {
		return nil, err
	}
	key, err := keystore.DecryptKey(json, string(passphrase))
	if err != nil {
		return nil, err
	}
	return key, nil
}

func MaxNonce(txs types.Transactions) uint64 {
	txNonce := uint64(0)
	for _, tx := range txs {
		if tx.Nonce() > txNonce {
			txNonce = tx.Nonce()
		}
	}
	return txNonce
}

func EnableNonce(txs types.Transactions, getNonce func() uint64) uint64 {
	txNonce := MaxNonce(txs)
	if txNonce != 0 {
		txNonce++
	} else {
		txNonce = getNonce()
	}
	return txNonce
}
