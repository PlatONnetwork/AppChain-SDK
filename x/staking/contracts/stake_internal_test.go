package contracts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/umbracle/ethgo"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/umbracle/ethgo/abi"
)

func TestEncodeData(t *testing.T) {

	_ADDSTAKE_SIG := crypto.Keccak256Hash([]byte("ADDSTAKE"))
	addr := common.HexToAddress("0xFA66dAa530328D0d914B6652e4B64B00d84e3a1a")
	amount := 99
	//abiType := abi.MustNewType("tuple(bytes32, address, uint256)")
	abiType := abi.MustNewType("tuple(bytes32 STAKE_SIG, address addr, uint256 amount)")
	input, err := abiType.Encode([]interface{}{_ADDSTAKE_SIG, addr, amount})
	if nil != err {
		t.Error(err)
	}
	fmt.Printf("SIGN:\n %s \n", _ADDSTAKE_SIG.Hex())
	fmt.Printf("data:\n %s \n", common.Bytes2Hex(input))
}

// 0x7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73
// 0x7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73

//0x7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73000000000000000000000000fa66daa530328d0d914b6652e4b64b00d84e3a1a0000000000000000000000000000000000000000000000000000000000000063
// 7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73000000000000000000000000fa66daa530328d0d914b6652e4b64b00d84e3a1a0000000000000000000000000000000000000000000000000000000000000063

func TestDecodeData(t *testing.T) {
	_ADDSTAKE_SIG := crypto.Keccak256Hash([]byte("ADDSTAKE"))
	//addr := common.HexToAddress("0xFA66dAa530328D0d914B6652e4B64B00d84e3a1a")
	//amount := 99

	data := common.Hex2Bytes("7f629647b0cf8231fa5380e25f7c9bf0685fecbdc41360b93da5b447cef9ee73000000000000000000000000fa66daa530328d0d914b6652e4b64b00d84e3a1a0000000000000000000000000000000000000000000000000000000000000063")
	fmt.Printf("input len: %d \n", len(data))
	//abiType := abi.MustNewType("tuple(bytes32, address, uint256)")
	//abiType := abi.MustNewType("tuple(address, uint256)")
	abiType := abi.MustNewType("tuple(address addr, uint256 amount)")

	assert.True(t, bytes.Compare(_ADDSTAKE_SIG.Bytes(), data[:32]) == 0, "no equals sign")
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
