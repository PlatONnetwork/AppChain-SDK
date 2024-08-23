package erc20vote

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/test"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

var (
	evm       *vm.EVM
	contract  *vm.Contract
	erc20Vote *ERC20Vote
	amount    = big.NewInt(100)
)

func init() {
	initContract()
}

func initContract() {
	evm = test.NewEVM(test.NewMemoryStateDB(), vm.Config{}, vm.TxContext{}, test.NewBlockContext(), nil)
	contract = vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(test.To), big.NewInt(0), 10000000)
	erc20Vote, _ = NewERC20Vote(evm, contract, false)
	erc20Vote.Init("vote", "vote", "1", test.From)
}

func TestMint(t *testing.T) {
	oldBalance, err := erc20Vote.BalanceOf(test.OneAddr)
	require.Nil(t, err)
	erc20Vote.Mint(test.OneAddr, amount)
	newBalance, err := erc20Vote.BalanceOf(test.OneAddr)
	require.Nil(t, err)
	require.Equal(t, oldBalance.Add(oldBalance, amount), newBalance)
	vote, err := erc20Vote.GetVotes(test.OneAddr)
	require.Nil(t, err)
	require.Equal(t, newBalance, vote)

	require.Panics(t, func() {
		erc20Vote.GetPastTotalSupply(evm.Context.BlockNumber)
	})

}
func TestBurn(t *testing.T) {
	oldBalance, err := erc20Vote.BalanceOf(test.OneAddr)
	require.Nil(t, err)
	erc20Vote.Mint(test.OneAddr, amount)

	err = erc20Vote.Burn(test.OneAddr, amount)
	require.Nil(t, err)
	newBalance, err := erc20Vote.BalanceOf(test.OneAddr)
	require.Nil(t, err)
	vote, err := erc20Vote.GetVotes(test.OneAddr)
	require.Nil(t, err)
	require.Equal(t, newBalance.Sub(newBalance, oldBalance), vote)
}

func TestCheckpoint(t *testing.T) {

	erc20Vote.Mint(test.OneAddr, amount)
	erc20Vote.Mint(test.OneAddr, amount)
	erc20Vote.Mint(test.OneAddr, amount)

	num, err := erc20Vote.NumCheckpoints(test.OneAddr)
	require.Nil(t, err)
	require.Equal(t, uint32(1), num)

}

func TestPastVotes(t *testing.T) {
	erc20Vote.Mint(test.OneAddr, amount)

	blockContext := test.NewBlockContext()
	blockContext.BlockNumber = big.NewInt(10)
	evm := test.NewEVM(evm.StateDB.(*state.StateDB), vm.Config{}, vm.TxContext{}, blockContext, nil)
	contract := vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(test.To), big.NewInt(0), 10000000)
	erc20Vote, _ := NewERC20Vote(evm, contract, false)
	erc20Vote.Mint(test.OneAddr, amount)
	num, err := erc20Vote.NumCheckpoints(test.OneAddr)
	require.Nil(t, err)
	require.Equal(t, uint32(2), num)
	evm.Context.BlockNumber = big.NewInt(11)
	votes, err := erc20Vote.GetPastVotes(test.OneAddr, big.NewInt(10))
	require.Nil(t, err)
	balance, _ := erc20Vote.BalanceOf(test.OneAddr)
	require.Equal(t, balance, votes)
}

func TestTransfer(t *testing.T) {
	erc20Vote.Mint(test.From, amount)
	blockContext := test.NewBlockContext()
	blockContext.BlockNumber = big.NewInt(10)
	evm := test.NewEVM(evm.StateDB.(*state.StateDB), vm.Config{}, vm.TxContext{}, blockContext, nil)
	contract := vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(test.To), big.NewInt(0), 10000000)
	erc20Vote, _ := NewERC20Vote(evm, contract, false)
	erc20Vote.Transfer(test.TwoAddr, amount)
	evm.Context.BlockNumber = big.NewInt(11)
	votes, err := erc20Vote.GetPastVotes(test.TwoAddr, big.NewInt(10))
	require.Nil(t, err)
	require.Equal(t, big.NewInt(100), votes)
}
