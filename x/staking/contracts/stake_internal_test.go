package contracts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
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
	fmt.Printf("%s", string(b))
}
