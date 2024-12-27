package lottery

import (
	"encoding/json"
	"fmt"
	common2 "github.com/PlatONnetwork/AppChain-SDK/example/common"
	"github.com/PlatONnetwork/AppChain-SDK/example/lottery/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/types/module"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/vrf"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"github.com/stretchr/testify/require"
	"golang.org/x/net/context"
	"math/big"
	"sync"
	"testing"
	"time"
)

func TestVRF(t *testing.T) {
	sk, _ := common2.UserPrivateKeys[common2.UserAddrs[0]]
	nonceProof, err := vrf.Prove(sk, common.Hash{}.Bytes())
	require.Nil(t, err)
	success, err := vrf.Verify(&sk.PublicKey, nonceProof, common.Hash{}.Bytes())
	require.True(t, success)
	fmt.Println(hexutil.Encode((nonceProof)))
	addr, _ := common.Bech32ToAddress("lat1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqqqq542u6a")
	fmt.Println(addr.Hex())
}

func TestLottery(t *testing.T) {
	lotteryModule := Module{
		key:     testutil.DefaultAccount[0].NodePrivateKey(),
		address: crypto.PubkeyToAddress(testutil.DefaultAccount[0].NodePrivateKey().PublicKey),
	}

	vals, err := testutil.NewValidator(testutil.DefaultAccount[0:1])
	require.Nil(t, err)
	vals.SetCoinbase(testutil.DefaultAccount[0].NodePrivateKey())
	manager := module.NewManager(vals, &lotteryModule)

	manager.SetElection(vals.Name())
	manager.SetOrderGenesis(lotteryModule.Name())
	app := testutil.NewApp(manager)
	require.Nil(t, err)
	config := GenesisConfig{
		Name:   "Token",
		Symbol: "USDC",
		Nonce:  "0x03f3b376f00863de14440eff826835d16ffa3c8b0fc7ad1402beee7ccf076aa9282f795c3e92d2f61e45e85abe5fcfef134ae6700a50b0885942a92d92b9c88a280450a416880be13a23e449f41ef12f43",
	}
	s, _ := json.Marshal(config)
	stack, backend, err := testutil.CreateCluster(testutil.DefaultAccount[0:1], []sdk.App{app}, map[string]json.RawMessage{
		lotteryModule.Name(): s,
	}, []common.Address{common.BigToAddress(big.NewInt(1)), common.BigToAddress(big.NewInt(2))})
	require.Nil(t, err)
	require.Nil(t, stack[0].Start())
	require.Nil(t, backend[0].Start())
	var sg sync.WaitGroup
	sg.Add(1)
	sg.Wait()
}

func TestClient(t *testing.T) {
	cli, err := ethclient.Dial("http://127.0.0.1:8801")
	require.Nil(t, err)
	builder, err := contracts.NewLotteryTxBuilder(LotteryAddr, testutil.DefaultAccount[1].NodePrivateKey(), testutil.DefaultGenesisTemplate.Config.ChainID)
	require.Nil(t, err)
	for i := 0; i < 1; i++ {
		//blockNumber, err := cli.BlockNumber(context.Background())
		//require.Nil(t, err)
		gasPrice, err := cli.SuggestGasPrice(context.Background())
		require.Nil(t, err)
		nonce, err := cli.NonceAt(context.Background(), testutil.DefaultAccount[1].NodeAddress(), nil)
		tx, err := builder.WithNonce(nonce).WithGasLimit(10000000).WithGasPrice(gasPrice).Guessing(big.NewInt(2))
		require.Nil(t, err)
		//cli.CallContract(context.Background(), platon.CallMsg{
		//	From:       common.Address{},
		//	To:         nil,
		//	Gas:        1000000,
		//	GasPrice:   nil,
		//	GasFeeCap:  nil,
		//	GasTipCap:  nil,
		//	Value:      nil,
		//	Data:       nil,
		//	AccessList: nil,
		//})
		cli.SendTransaction(context.Background(), tx)

		for i := 0; i < 5; i++ {
			receipt, _ := cli.TransactionReceipt(context.Background(), tx.Hash())
			if receipt != nil {
				fmt.Println("success:", tx.Hash(), receipt.Status)
				break
			}
			time.Sleep(time.Second)
		}
		//input, err := builder.PackGetNonce(new(big.Int).SetUint64(blockNumber))
		//require.Nil(t, err)
		//res, err := cli.CallContract(context.Background(), platon.CallMsg{
		//	From:       testutil.DefaultAccount[0].NodeAddress(),
		//	To:         &VRFStorageAddr,
		//	Gas:        1100000,
		//	GasPrice:   gasPrice,
		//	GasFeeCap:  nil,
		//	GasTipCap:  nil,
		//	Value:      nil,
		//	Data:       input,
		//	AccessList: nil,
		//}, nil)
		//if err != nil {
		//	if res != nil {
		//		fmt.Println(string(res))
		//	}
		//	fmt.Println(err.Error())
		//}
		//fmt.Println(len(res))
		time.Sleep(time.Second)
	}

}
func TestA(t *testing.T) {
	//addr, _ := common.Bech32ToAddress("lat1qqqqqqqqqqqqqqqqqqqqqqqqqqqqqq2d9s932a")
	//fmt.Println(addr.Hex())
	//h := vrf.ProofToHash([]byte("abc"))
	//fmt.Println(len(h))
	for _, s := range []string{"0x70d207c1322ccb9069d3790d6768866dabff1035",
		"0x343972bf63d1062761aaaa891d2750f03cb4b2f7",
		"0x97ab3d4f7f5051f127b0e9f8d10772125d94d65b",
		"0xa7429ae04d8bb89cfd572963adea8cbf8219609a",
		"0xd503c89dc82e0d43688860e9bdf3ef3efe07dae2"} {
		addr := common.HexToAddress(s)
		fmt.Println(s, "==>", addr.String())
	}
}
