package erc20

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/test"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var (
	evm      *vm.EVM
	contract *vm.Contract
	erc20    *ERC20
)

func init() {
	initContract()
}
func initContract() {
	evm = test.NewEVM(test.NewMemoryStateDB(), vm.Config{}, vm.TxContext{}, test.NewBlockContext(), nil)
	contract = vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(test.To), big.NewInt(0), 1000000)
	erc20, _ = NewERC20(evm, contract, false)
	erc20.Init("token", "token", test.From)
}

func TestFlow(t *testing.T) {
	initContract()
	TestMint(t)
}

func TestMint(t *testing.T) {
	oldTotal, err := erc20.TotalSupply()
	require.Nil(t, err)

	err = erc20.Mint(test.ToAddress(1), big.NewInt(10))
	require.Nil(t, err)
	newTotal, err := erc20.TotalSupply()
	require.Equal(t, oldTotal.Add(oldTotal, big.NewInt(10)), newTotal)
	require.Panics(t, func() {
		erc20.Mint(common.Address{}, big.NewInt(10))
	})
	balance, err := erc20.BalanceOf(test.ToAddress(1))
	require.Equal(t, big.NewInt(10), balance)
}

func TestBurn(t *testing.T) {
	oldBalance, err := erc20.BalanceOf(test.ToAddress(1))
	require.Nil(t, err)
	err = erc20.Mint(test.ToAddress(1), big.NewInt(10))
	require.Nil(t, err)
	err = erc20.Burn(test.ToAddress(1), big.NewInt(10))
	require.Nil(t, err)

	newBalance, err := erc20.BalanceOf(test.ToAddress(1))
	require.Equal(t, oldBalance, newBalance)
}

func TestApprove(t *testing.T) {
	contract := vm.NewContract(vm.AccountRef(test.ToAddress(1)), vm.AccountRef(test.To), big.NewInt(0), 100000)
	c, _ := NewERC20(evm, contract, false)
	flag, err := c.Approve(test.ToAddress(2), big.NewInt(100))
	require.True(t, flag)
	require.Nil(t, err)
	allowance, err := erc20.Allowance(test.ToAddress(1), test.ToAddress(2))
	require.Nil(t, err)
	require.Equal(t, big.NewInt(100), allowance)
	flag, err = c.IncreaseAllowance(test.ToAddress(2), big.NewInt(100))
	require.True(t, flag)
	require.Nil(t, err)
	allowance, err = erc20.Allowance(test.ToAddress(1), test.ToAddress(2))
	require.Nil(t, err)
	require.Equal(t, big.NewInt(200), allowance)
	flag, err = c.DecreaseAllowance(test.ToAddress(2), big.NewInt(100))
	require.True(t, flag)
	require.Nil(t, err)
	allowance, err = erc20.Allowance(test.ToAddress(1), test.ToAddress(2))
	require.Nil(t, err)
	require.Equal(t, big.NewInt(100), allowance)
}

func TestTransfer(t *testing.T) {
	oldBalance, err := erc20.BalanceOf(test.ToAddress(1))
	require.Nil(t, err)
	err = erc20.Mint(test.ToAddress(1), big.NewInt(10))
	require.Nil(t, err)
	contract := vm.NewContract(vm.AccountRef(test.ToAddress(1)), vm.AccountRef(test.To), big.NewInt(0), 100000)
	c, _ := NewERC20(evm, contract, false)
	flag, err := c.Transfer(test.ToAddress(2), big.NewInt(10))
	require.Nil(t, err)
	require.True(t, flag)
	toBalance, err := erc20.BalanceOf(test.ToAddress(2))
	require.Nil(t, err)
	require.Equal(t, big.NewInt(10), toBalance)
	newBalance, err := erc20.BalanceOf(test.ToAddress(1))
	require.Nil(t, err)
	require.Equal(t, oldBalance, newBalance)
}

func TestTransferFrom(t *testing.T) {
	contract := vm.NewContract(vm.AccountRef(test.ToAddress(1)), vm.AccountRef(test.To), big.NewInt(0), 100000)
	c, _ := NewERC20(evm, contract, false)
	flag, err := c.Approve(test.ToAddress(2), big.NewInt(100))
	require.True(t, flag)
	require.Nil(t, err)
	oldBalance, err := erc20.BalanceOf(test.ToAddress(1))
	require.Nil(t, err)
	err = erc20.Mint(test.ToAddress(1), big.NewInt(10))
	require.Nil(t, err)
	contract = vm.NewContract(vm.AccountRef(test.ToAddress(2)), vm.AccountRef(test.To), big.NewInt(0), 100000)
	c, _ = NewERC20(evm, contract, false)
	flag, err = c.TransferFrom(test.ToAddress(1), test.ToAddress(3), big.NewInt(10))
	require.Nil(t, err)
	require.True(t, flag)
	toBalance, err := erc20.BalanceOf(test.ToAddress(3))
	require.Nil(t, err)
	require.Equal(t, big.NewInt(10), toBalance)
	newBalance, err := erc20.BalanceOf(test.ToAddress(1))
	require.Nil(t, err)
	require.Equal(t, oldBalance, newBalance)
}
