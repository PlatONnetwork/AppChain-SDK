package txrelayer

import (
	"math/big"
	"testing"

	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/google/uuid"
	"github.com/test-go/testify/assert"
)

var (
	privateKeyHex = "68e7f80b51de2c0761fcdccc43d4c07d94e19dff399feafc06c2db6af1610c2d"
	CounterAddr   = common.HexToAddress("0xdaCB886C825aF81d4361dF81Dc556ba017772ee7")
)

func newTxRelayer(rpcAddr string) *Module {
	txRelayer := NewModule(rpcAddr, DefaultReceiptTimeout, DefaultNumRetries)
	txRelayer.Init(nil)
	return txRelayer
}

func TestCall(t *testing.T) {
	m := newTxRelayer("https://devnet2openapi2.platon.network/rpc")

	data := common.Hex2Bytes("10c2044e")
	wlat := common.HexToAddress("0xcc9fbab49c29b3ff536a3d94873e988cc4a572af")
	to := common.HexToAddress("0x71E7Fdf12f8aEA0671f1dB0fdC27B3B22e5d0e44")
	result, err := m.Call(common.ZeroAddr, to, data)
	assert.Nil(t, err)
	addr := common.BytesToAddress(result[12:])
	assert.Equal(t, wlat, addr)
}

func TestSendTransaction(t *testing.T) {
	m := newTxRelayer("https://devnet2openapi2.platon.network/rpc")

	privateKey := crypto.HexMustToECDSA(privateKeyHex)
	id, err := uuid.NewRandom()
	assert.NoError(t, err)
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}

	numberData := common.Hex2Bytes("8381f58a")
	result, err := m.Call(key.Address, CounterAddr, numberData)
	assert.Nil(t, err)
	number := big.NewInt(0).SetBytes([]byte(result))

	incrementData := common.Hex2Bytes("d09de08a")
	receipt, err := m.SendTransaction(types.NewTx(&types.LegacyTx{
		Nonce:    0,
		GasPrice: nil,
		Gas:      0,
		Data:     incrementData,
		To:       &CounterAddr,
		Value:    nil,
	}), key)
	assert.Nil(t, err)
	assert.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)

	result, err = m.Call(common.ZeroAddr, CounterAddr, numberData)
	assert.Nil(t, err)
	number2 := big.NewInt(0).SetBytes(result)
	assert.Equal(t, number2, number.Add(number, big.NewInt(1)))
}
