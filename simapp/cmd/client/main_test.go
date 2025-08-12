package main

import (
	"context"
	"encoding/hex"
	"github.com/PlatONnetwork/AppChain-SDK/simapp/cmd/client/contracts/registrymanager"
	"github.com/PlatONnetwork/AppChain-SDK/x/benchmark"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/pelletier/go-toml/v2"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

type BB struct {
	E string `toml:"e"                     comment:"RegistryManager implement contract address"`
	F string `toml:"f"                     comment:"RegistryManager implement contract address"`

	G string `toml:"g"                     comment:"RegistryManager implement contract address"`
}
type Con struct {
	A string      `toml:"a"                     comment:"RegistryManager implement contract address"`
	B interface{} `toml:"b"                     comment:"RegistryManager implement contract address"`
}

func TestB(t *testing.T) {
	a, err := toml.Marshal(&Con{
		A: "abc",
		B: &BB{
			E: "eee",
			F: "fff",
			G: "ggg",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(a))
	var con Con
	toml.NewDecoder()
	err = toml.Unmarshal(a, &con)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(con)
}
func TestA(t *testing.T) {
	key, _ := keystore.DecryptKey([]byte("{\"crypto\":{\"cipher\":\"aes-128-ctr\",\"cipherparams\":{\"iv\":\"50108431ff56c4ca9b2c196a2a4f3b25\"},\"ciphertext\":\"5e93d704bf0de41f0a41d9ea63c61bbf9422f1e380e5a31cf8052bcfdba71fd1\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":8192,\"p\":1,\"r\":8,\"salt\":\"266eb0d13fecd7e6dc07ecde64103b33c78c3a4320d9ee9aa550f3aaf16a2738\"},\"mac\":\"d3517ba0f3ca997b267b4523603a351d2be5306577c309052941c9f9f02ee52e\"},\"id\":\"62d63dd4-bfb3-451e-b4d7-b9c75b02a894\",\"version\":3}"), "123456")

	//testnet := "https://devnet2openapi.platon.network/rpc"
	localnet := "http://127.0.0.1:8801"
	//testnetCli, _ := ethclient.Dial(testnet)
	localCli, _ := ethclient.Dial(localnet)
	nonce, _ := localCli.NonceAt(context.Background(), crypto.PubkeyToAddress(key.PrivateKey.PublicKey), nil)
	hash := common.HexToHash("0xcf7f8f1e35aab04affa259ba8cef66ff4a300311a55c0d6819927ce82ef89c00")
	//testnetChainId, _ := testnetCli.ChainID(context.Background())
	localChainId, _ := localCli.ChainID(context.Background())
	//authOpts := createTransactionOpts(key.PrivateKey, localChainId, 0, nil)
	tx, _, _ := localCli.TransactionByHash(context.Background(), hash)
	tx = types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		GasPrice: tx.GasPrice(),
		Gas:      tx.Gas(),
		To:       tx.To(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	})
	tx, _ = types.SignTx(tx, types.NewLondonSigner(localChainId), key.PrivateKey)
	localCli.SendTransaction(context.Background(), tx)
	waitTx(localCli, tx.Hash())
}

//func TestDe(t *testing.T) {
//	deploytools.MustDecodePubKey("45edf2fe5bef8be6cee13ef809bb60817f9008f66849a25558cfa25d1742f498702710955ffc3c1b6f0a3c3cbaa0466e5408cdc8eb3b2e5b4c6af9ae81b34d9c")
//	common.SetAddressHRP("lat")
//	key, err := keystore.DecryptKey([]byte("{\"crypto\":{\"cipher\":\"aes-128-ctr\",\"cipherparams\":{\"iv\":\"50108431ff56c4ca9b2c196a2a4f3b25\"},\"ciphertext\":\"5e93d704bf0de41f0a41d9ea63c61bbf9422f1e380e5a31cf8052bcfdba71fd1\",\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":8192,\"p\":1,\"r\":8,\"salt\":\"266eb0d13fecd7e6dc07ecde64103b33c78c3a4320d9ee9aa550f3aaf16a2738\"},\"mac\":\"d3517ba0f3ca997b267b4523603a351d2be5306577c309052941c9f9f02ee52e\"},\"id\":\"62d63dd4-bfb3-451e-b4d7-b9c75b02a894\",\"version\":3}"), "123456")
//	if err != nil {
//		panic(err)
//	}
//	t.Log(hex.EncodeToString(crypto.FromECDSA(key.PrivateKey)))
//	//fmt.Println(fmt.Sprintf(commandScript, "aaa"))
//	buffer := new(bytes.Buffer)
//
//	tmpl := template.Must(template.New("").Parse(commandScript))
//
//	if err := tmpl.Execute(buffer, &main2.ScriptParams{
//		Dir: "abcddf",
//	}); err != nil {
//		panic(err)
//	}
//	// For Go bindings pass the code through gofmt to clean it up
//	fmt.Println(string(buffer.Bytes()))
//}

func TestDeploy(t *testing.T) {
	accs, _ := benchmark.GenKeys(1000)
	sk := accs[999]
	t.Log(hex.EncodeToString(crypto.FromECDSA(sk)))
	return
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	cli, err := ethclient.Dial("http://10.1.1.33:8801")
	require.Nil(t, err)
	nonce, err := cli.PendingNonceAt(context.Background(), addr)
	require.Nil(t, err)
	t.Log(nonce)
	chainid, err := cli.ChainID(context.Background())
	require.Nil(t, err)

	t.Log(chainid)
	addr, tx, manager, err := registrymanager.DeployRegistryManager(&bind.TransactOpts{
		From: addr,
		Signer: func(address common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, types.NewLondonSigner(chainid), sk)
		},
		Value:     big.NewInt(0),
		GasPrice:  big.NewInt(1000000000),
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  10000000,
		Context:   nil,
		NoSend:    false,
	}, cli)
	require.Nil(t, err)
	for {
		receipt, err := cli.TransactionReceipt(context.Background(), tx.Hash())
		if receipt == nil && err != nil {
			continue
		}
		require.Nil(t, err)
		t.Log(receipt)
		break
	}
	addr2, err := manager.GetCheckpointManager(&bind.CallOpts{
		Pending:     false,
		From:        addr,
		BlockNumber: nil,
		Context:     nil,
	}, big.NewInt(1))
	t.Log(addr2)
}

func A() (int, string, error) {
	return 1, "abc", nil
}

func B() (string, string, error) {
	return "efg", "def", nil
}
