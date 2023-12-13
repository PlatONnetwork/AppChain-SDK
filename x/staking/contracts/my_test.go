package contracts

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/umbracle/ethgo"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umbracle/ethgo/abi"
)

func TestEncodeStakeData(t *testing.T) {

	ADDSTAKE_SIG := crypto.Keccak256Hash([]byte("ADDSTAKE"))
	addr := common.HexToAddress("0xFA66dAa530328D0d914B6652e4B64B00d84e3a1a")
	amount := 99
	//abiType := abi.MustNewType("tuple(bytes32, address, uint256)")
	abiType := abi.MustNewType("tuple(bytes32 STAKE_SIG, address addr, uint256 amount)")
	input, err := abiType.Encode([]interface{}{ADDSTAKE_SIG, addr, amount})
	if nil != err {
		t.Error(err)
	}
	fmt.Printf("SIGN:\n %s \n", ADDSTAKE_SIG.Hex())
	fmt.Printf("data:\n %s \n", common.Bytes2Hex(input))
}

// 0x7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73
// 0x7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73

//0x7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73000000000000000000000000fa66daa530328d0d914b6652e4b64b00d84e3a1a0000000000000000000000000000000000000000000000000000000000000063
// 7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73000000000000000000000000fa66daa530328d0d914b6652e4b64b00d84e3a1a0000000000000000000000000000000000000000000000000000000000000063

func TestDecodeStakeData(t *testing.T) {
	ADDSTAKE_SIG := crypto.Keccak256Hash([]byte("ADDSTAKE"))
	//addr := common.HexToAddress("0xFA66dAa530328D0d914B6652e4B64B00d84e3a1a")
	//amount := 99
	//
	data := common.Hex2Bytes("7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73000000000000000000000000fa66daa530328d0d914b6652e4b64b00d84e3a1a0000000000000000000000000000000000000000000000000000000000000063")
	fmt.Printf("input len: %d \n", len(data))
	//abiType := abi.MustNewType("tuple(bytes32, address, uint256)")
	//abiType := abi.MustNewType("tuple(address, uint256)")
	abiType := abi.MustNewType("tuple(address addr, uint256 amount)")

	assert.True(t, bytes.Compare(ADDSTAKE_SIG.Bytes(), data[:32]) == 0, "no equals sign")
	//decoded, err := abiType.Decode(data)
	decoded, err := abiType.Decode(data[32:])
	//decoded, err := abi.Decode(abiType, data)
	//decoded, err := abi.Decode(abiType, data[32:])
	if nil != err {
		t.Error(err)
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		t.Error("err......")
	}

	b, e := json.Marshal(res)
	if e != nil {
		t.Error(e)
	}
	fmt.Printf("%s \n", string(b))

	addr, ok := res["addr"].(ethgo.Address)
	if !ok {
		t.Error("failed......")
	}
	fmt.Printf("%s \n", addr.String())
	addr1 := common.Address(addr)

	fmt.Printf("%s \n", addr1.Hex())
}

func TestEncodeMethod(t *testing.T) {
	SyncState := crypto.Keccak256Hash([]byte("SyncState(address, bytes)"))
	syncState := crypto.Keccak256Hash([]byte("syncState(address, bytes)"))
	fmt.Printf("SyncState: %s \n", SyncState.Hex())
	fmt.Printf("syncState: %s \n", syncState.Hex())
}

func TestInsertItem(t *testing.T) {

	a := []uint64{1, 2, 4, 5, 6}

	//a := []uint64{}

	size := len(a)

	item := uint64(2)

	if size != 0 {
		if a[size-1] < item {
			a = append(a, item)
		} else {
			for i := 0; i < size; i++ {
				if a[i] == item {
					break
				}

				if a[i] > item {
					a = append(a, 0)
					copy(a[i+1:], a[i:])
					a[i] = item
					break
				}
			}
		}
	} else {
		a = []uint64{item}
	}

	fmt.Printf("a len : %d, \n a: %+v \n", len(a), a)
}

func TestInsertSlice(t *testing.T) {

	a := []uint64{1, 2, 6, 7, 8}

	b := []uint64{3, 4, 5}

	a = append(a, b...)

	copy(a[2+len(b):], a[2:])
	copy(a[2:], b)
	fmt.Printf("a len : %d, \n a: %+v \n", len(a), a)
}

type testItem struct {
	Amount uint64
}

type testQueue []*testItem

func newTestQueue(size uint64) testQueue {
	queue := make(testQueue, size)
	return queue
}
func (queue testQueue) Append(item *testItem) testQueue {
	queue = append(queue, item)
	return queue
}

func (queue testQueue) IsEmpty() bool {
	return len(queue) == 0
}
func (queue testQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

func TestQueueAppend(t *testing.T) {
	queue := newTestQueue(0)

	for _, v := range []uint64{1, 3, 8, 9, 5, 6} {
		item := &testItem{
			v,
		}
		//queue.Append(item) // queue: []
		// queue = queue.Append(item) // queue: [{"Amount":1},{"Amount":3},{"Amount":8},{"Amount":9},{"Amount":5},{"Amount":6}]
		queue = append(queue, item) // queue: [{"Amount":1},{"Amount":3},{"Amount":8},{"Amount":9},{"Amount":5},{"Amount":6}]
	}

	v, err := json.Marshal(queue)
	if nil != err {
		t.Error(err)
	}

	fmt.Printf("queue: %s \n", string(v))
	fmt.Printf("queue empty: %v \n", queue.IsEmpty())
	fmt.Printf("queue not empty: %v \n", queue.IsNotEmpty())

	q := newTestQueue(0)

	fmt.Printf("q empty: %v \n", q.IsEmpty())
	fmt.Printf("q not empty: %v \n", q.IsNotEmpty())
}

func TestEncodeSlashData(t *testing.T) {

	SLASH_SIG := crypto.Keccak256Hash([]byte("SLASH"))
	addrs := []common.Address{common.HexToAddress("0xFA66dAa530328D0d914B6652e4B64B00d84e3a1a"), common.HexToAddress("0x4d21D80BA135AD50f515861a5b334A2B4FF8271D")}
	amounts := []uint64{66, 72}

	abiType := abi.MustNewType("tuple(bytes32 SLASH_SIG, uint256 handleEventId, address[] addrs, uint256[] amounts)")
	input, err := abiType.Encode([]interface{}{SLASH_SIG, uint64(12), addrs, amounts})
	if nil != err {
		t.Error(err)
	}
	fmt.Printf("SIGN:\n %s \n", SLASH_SIG.Hex())
	fmt.Printf("data:\n %s \n", common.Bytes2Hex(input))
}

func TestDecodeSlashData(t *testing.T) {
	SLASH_SIG := crypto.Keccak256Hash([]byte("SLASH"))

	data := common.Hex2Bytes("117f1d6f44fd34ccb7a58f1261fa59e5c4bf68e2712d65f246a8805167a93344000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000008000000000000000000000000000000000000000000000000000000000000000e00000000000000000000000000000000000000000000000000000000000000002000000000000000000000000fa66daa530328d0d914b6652e4b64b00d84e3a1a0000000000000000000000004d21d80ba135ad50f515861a5b334a2b4ff8271d000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000420000000000000000000000000000000000000000000000000000000000000048")

	abiType := abi.MustNewType("tuple(bytes32 SLASH_SIG, uint256 id, address[] addrs, uint256[] amounts)")

	//abiType := abi.MustNewType("tuple(uint256 id, address[] addrs, uint256[] amounts)")

	assert.True(t, bytes.Compare(SLASH_SIG.Bytes(), data[:32]) == 0, "no equals sign")

	//decoded, err := abiType.Decode(data[32:])

	decoded, err := abiType.Decode(data)

	if nil != err {
		t.Error(err)
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		t.Error("err......")
	}

	//b, e := json.Marshal(res)
	//if e != nil {
	//	t.Error(e)
	//}
	//fmt.Printf("%s \n", string(b))

	id, ok := res["id"].(*big.Int)
	if !ok {
		t.Error("failed......")
	}
	fmt.Printf("id: %s \n", fmt.Sprintf("%+v", id))

	addrs, ok := res["addrs"].([]ethgo.Address)
	if !ok {
		t.Error("failed......")
	}
	addrQueue := make([]common.Address, len(addrs))
	for i, v := range addrs {
		addrQueue[i] = common.Address(v)
		fmt.Printf("addr: %s \n", addrQueue[i].Hex())
	}

	amounts, ok := res["amounts"].([]*big.Int)

	fmt.Printf("amounts: %s \n", fmt.Sprintf("%+v", amounts))
}

func TestMapRange(t *testing.T) {

	cache := make(map[string]struct{}, 0)

	// current round queue {"a", "b", "c", "d", "e", "f", "g", "h"}

	// unstake queue {"b", "c", "g"}

	// currentepoch queue {"c", "d", "g", "m", "n"}

	for _, c := range []string{"b", "c", "g"} {
		cache[c] = struct{}{}
	}

	for _, c := range []string{"c", "d", "g", "m", "n"} {
		if _, ok := cache[c]; ok {
			delete(cache, c)
		}
	}

	fmt.Printf("cache: %+v \n", cache)

}

func TestMapRange2(t *testing.T) {

	removeCache := make(map[string]struct{}, 0)
	cache := make(map[string]struct{}, 0)
	queue := make([]string, 0)

	// current round queue {"a", "b", "c", "d", "e", "f", "g", "h"}

	// unstake queue {"b", "c", "g"}

	// currentepoch queue {"c", "d", "g", "m", "n"}

	for _, c := range []string{"b", "c", "g"} {
		cache[c] = struct{}{}
		queue = append(queue, c)
	}

	for _, c := range []string{"c", "d", "g", "m", "n"} {
		if _, ok := cache[c]; ok {
			delete(cache, c)
		}
	}

	for i := 0; i < len(queue); i++ {

		c := queue[i]
		if v, ok := cache[c]; !ok {
			// remove the can on withdrewqueue
			queue = append(queue[:i], queue[i+1:]...)
			i--
		} else {
			// append to the collection that needs to be removed
			removeCache[c] = v
		}
	}

	fmt.Printf("cache: %+v \nremoveCache: %+v \nqueue: %+v \n", cache, removeCache, queue)

}

func TestPrivateKey(t *testing.T) {
	// address: 0x5AC3D0154832bA77CB34c5de3Ef67f1acfbd4730
	// address: 0x5c8Fb7c7746417b551B953DEac5dE42E06871B2f
	// address: 0x35E1EC7b136DeF2E2d561F7E003deE2CBEf3C4c9
	// 0x72776072703AD32b29A3546063c801b3Af7a28e9
	priKeyArr := []*ecdsa.PrivateKey{
		crypto.HexMustToECDSA("f55e740c5099295ca16bccd53b868cdf648645812e36ba1e45442db433f99580"),
		crypto.HexMustToECDSA("3f3775f651d33432fa32f52f7a37876992feca6f64bb5f4ae21d5673c9ea5e75"),
		crypto.HexMustToECDSA("8ffc52602366ca88ac29743732984eb58b5df03758e717fa7d160a73333d238d"),
		crypto.HexMustToECDSA("eb7a0b9f4bf4927049b29ce4c717cf4956e3e644437a6b28fe90b17b7b1d4b4a"),
	}

	for i, privateKey := range priKeyArr {
		fmt.Printf("index: %d\n", i)
		fmt.Printf("privateKey: %s\n", hex.EncodeToString(crypto.FromECDSA(privateKey)))
		fmt.Printf("publicKey: %s\n", hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)[1:]))
		fmt.Printf("address: %s\n", crypto.PubkeyToAddress(privateKey.PublicKey).Hex())
	}
}

func TestGenesisVRFNonceHash(t *testing.T) {
	fmt.Printf("genesis vrf nonce hash: %s", common.BytesToHash([]byte("genesisVRFNonce")).Hex())
}

// 0xf22c7e0702483b876d22ebC6Ac0542C7ffF9A4Eb
func TestAddressToBech32(t *testing.T) {
	//addr := common.HexToAddress("0xf22c7e0702483b876d22ebC6Ac0542C7ffF9A4Eb")
	//fmt.Printf("address: %s\nbech32: %s\n", addr.Hex(), addr.Bech32())
	//assert.Equal(t, addr, common.MustBech32ToAddress(addr.Bech32()), "mismatching")

	addr := common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10")
	fmt.Printf("address: %s\nbech32: %s\n", addr.Hex(), addr.Bech32())
}
