package erc20

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	state2 "github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestErc20(t *testing.T) {
	db := state2.NewDatabase(rawdb.NewMemoryDatabase())
	statedb, _ := state2.New(common.Hash{}, db, nil)
	evm := &vm.EVM{
		Context: vm.BlockContext{
			BlockNumber: big.NewInt(10),
		},
		TxContext: vm.TxContext{
			Origin:   common.BigToAddress(big.NewInt(1)),
			GasPrice: nil,
		},
		StateDB: statedb,
		Config:  vm.Config{},
	}
	caller := vm.AccountRef(common.BigToAddress(big.NewInt(2)))
	object := vm.AccountRef(common.BigToAddress(big.NewInt(3)))
	contract := vm.NewContract(caller, object, nil, 0)
	erc20, err := NewErc20(evm, contract, false)

	input, err := Abi.Pack("decimals")
	require.Nil(t, err)
	output, err := erc20.Run(input)
	require.Nil(t, err)
	res, err := Abi.Unpack("decimals", output)
	require.Equal(t, res[0].(uint8), uint8(18))

	input, err = Abi.Pack("name")
	require.Nil(t, err)
	output, err = erc20.Run(input)
	require.Nil(t, err)
	res, err = Abi.Unpack("name", output)
	require.Equal(t, res[0].(string), "TestName")

	input, err = Abi.Pack("symbol")
	require.Nil(t, err)
	output, err = erc20.Run(input)
	require.Nil(t, err)
	res, err = Abi.Unpack("symbol", output)
	require.Equal(t, res[0].(string), "TEU")

	input, err = Abi.Pack("totalSupply")
	require.Nil(t, err)
	output, err = erc20.Run(input)
	require.Nil(t, err)
	res, err = Abi.Unpack("totalSupply", output)
	require.Equal(t, res[0].(*big.Int).Uint64(), uint64(0))

	input, err = Abi.Pack("mint", common.BigToAddress(big.NewInt(2)), big.NewInt(100))
	require.Nil(t, err)
	output, err = erc20.Run(input)
	require.Nil(t, err)
	require.Nil(t, output)

	input, err = Abi.Pack("totalSupply")
	require.Nil(t, err)
	output, err = erc20.Run(input)
	require.Nil(t, err)
	res, err = Abi.Unpack("totalSupply", output)
	require.Equal(t, res[0].(*big.Int).Uint64(), uint64(100))

	input, err = Abi.Pack("balanceOf", common.BigToAddress(big.NewInt(2)))
	output, err = erc20.Run(input)
	require.Nil(t, err)
	res, err = Abi.Unpack("balanceOf", output)
	require.Equal(t, res[0].(*big.Int).Uint64(), uint64(100))

	input, err = Abi.Pack("approve", common.BigToAddress(big.NewInt(5)), big.NewInt(1000))
	require.Nil(t, err)
	require.NotNil(t, input)
	output, err = erc20.Run(input)
	res, err = Abi.Unpack("approve", output)
	require.True(t, res[0].(bool))

	input, err = Abi.Pack("allowance", common.BigToAddress(big.NewInt(2)), common.BigToAddress(big.NewInt(5)))
	require.Nil(t, err)
	require.NotNil(t, input)
	output, err = erc20.Run(input)
	res, err = Abi.Unpack("allowance", output)
	require.Equal(t, res[0].(*big.Int).Uint64(), uint64(1000))

	input, err = Abi.Pack("transfer", common.BigToAddress(big.NewInt(3)), big.NewInt(50))
	require.Nil(t, err)
	require.NotNil(t, input)
	output, err = erc20.Run(input)
	res, err = Abi.Unpack("transfer", output)
	require.True(t, res[0].(bool))

	input, err = Abi.Pack("transfer", common.BigToAddress(big.NewInt(3)), big.NewInt(10000))
	require.Nil(t, err)
	require.NotNil(t, input)
	output, err = erc20.Run(input)
	res, err = Abi.Unpack("transfer", output)
	require.NotNil(t, err)

	contract = vm.NewContract(vm.AccountRef(common.BigToAddress(big.NewInt(5))), object, nil, 0)
	erc20, _ = NewErc20(evm, contract, false)
	input, err = Abi.Pack("transferFrom", common.BigToAddress(big.NewInt(2)), common.BigToAddress(big.NewInt(3)), big.NewInt(5))
	require.Nil(t, err)
	require.NotNil(t, input)
	output, err = erc20.Run(input)
	res, err = Abi.Unpack("transferFrom", output)
	require.True(t, res[0].(bool))
}
