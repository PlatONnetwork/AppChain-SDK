package coupon

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/example/coupon/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/stretchr/testify/require"
	"math/big"
	"sync"
	"testing"
	"time"
)

func TestCoupon(t *testing.T) {
	couponModule := Module{}

	vals, err := testutil.NewValidator(testutil.DefaultAccount[0:1])
	require.Nil(t, err)
	manager := module.NewManager(vals, &couponModule)
	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(couponModule.Name())
	app := testutil.NewApp(manager)
	require.Nil(t, err)
	config := GenesisConfig{
		Name:   "Token",
		Symbol: "USDC",
	}
	s, _ := json.Marshal(config)
	stack, backend, err := testutil.CreateCluster(testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		couponModule.Name(): json.RawMessage(s),
	}, []common.Address{common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2))})
	require.Nil(t, err)
	require.Nil(t, stack[0].Start())
	require.Nil(t, backend[0].Start())
	var sg sync.WaitGroup
	sg.Add(1)
	sg.Wait()
}

func TestClient(t *testing.T) {
	sk, err := crypto.HexToECDSA("8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd")
	//sk, err := crypto.GenerateKey()
	require.Nil(t, err)
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	require.Nil(t, err)
	coupon, err := contracts.NewContracts(CouponAddress, cli)
	require.Nil(t, err)
	from := crypto.PubkeyToAddress(sk.PublicKey)
	nonce, err := cli.NonceAt(context.Background(), from, nil)
	require.Nil(t, err)
	var txs types.Transactions
	for i := 0; i < 15; i++ {
		tx, err := coupon.ApplyCoupon(&bind.TransactOpts{
			From:  from,
			Nonce: new(big.Int).SetUint64(nonce),
			Signer: func(address common.Address, tx *types.Transaction) (*types.Transaction, error) {
				signer := types.NewLondonSigner(big.NewInt(123083))
				signature, err := crypto.Sign(signer.Hash(tx, nil).Bytes(), sk)
				if err != nil {
					return nil, err
				}
				return tx.WithSignature(signer, signature)
			},
			Value:     nil,
			GasPrice:  big.NewInt(11000000000),
			GasFeeCap: nil,
			GasTipCap: nil,
			GasLimit:  0,
			Context:   nil,
			NoSend:    false,
		})
		require.Nil(t, err)
		fmt.Println("tx:", tx.Hash().String())
		txs = append(txs, tx)
		nonce++
	}
	time.Sleep(time.Second)
	for _, tx := range txs {
		receipt, _ := cli.TransactionReceipt(context.Background(), tx.Hash())
		fmt.Println("logs:", len(receipt.Logs))
		for _, log := range receipt.Logs {
			event, err := coupon.ParseApply(*log)
			if err == nil {
				fmt.Println("event:", event.To.String(), event.Value, event.Balance)
				break
			}
		}
	}
}

func TestSK(t *testing.T) {
	for i := 0; i < 10; i++ {
		sk, err := crypto.GenerateKey()
		require.Nil(t, err)
		addr := crypto.PubkeyToAddress(sk.PublicKey)
		skbuf := crypto.FromECDSA(sk)
		fmt.Printf("common.HexToAddress(\"%s\"):crypto.HexMustToECDSA(\"%s\"),\n", addr.Hex(), hex.EncodeToString(skbuf))
	}
}
