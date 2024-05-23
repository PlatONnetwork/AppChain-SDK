package contracts

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts/test"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/erc20"
	"github.com/PlatONnetwork/AppChain-SDK/x/gov/contracts/erc20vote"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/stretchr/testify/require"
	abi2 "github.com/umbracle/ethgo/abi"
	"math/big"
	"testing"
)

var (
	evm         *vm.EVM
	contract    *vm.Contract
	governance  *Governance
	vote        *erc20vote.ERC20Vote
	amount      = big.NewInt(100)
	govName     = "OZ-Governor"
	tokenName   = "MockToken"
	tokenSymbol = "MTKN"
	version     = "1"
	voteDelay   = big.NewInt(4)
	votePeriod  = big.NewInt(16)
	voteAddr    = common.BigToAddress(big.NewInt(1001))
	govAddr     = common.BigToAddress(big.NewInt(1002))
	app         vm.ContractsApp
	token       *erc20.ERC20
	tokenAddr   = common.BigToAddress(big.NewInt(1003))
)

func init() {
	initContract()
}

func initContract() {
	app = test.NewContractsApp([]*test.SDKContract{
		&test.SDKContract{
			Addr: voteAddr,
			RunFunc: func(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
				ins, _ := erc20vote.NewERC20Vote(evm, contract, readOnly)
				return ins.Run(input)
			},
		},
		&test.SDKContract{
			Addr: govAddr,
			RunFunc: func(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
				ins, _ := NewGovernance(evm, contract, readOnly)
				return ins.Run(input)
			},
		},
		&test.SDKContract{
			Addr: tokenAddr,
			RunFunc: func(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
				ins, _ := erc20.NewERC20(evm, contract, readOnly)
				return ins.Run(input)
			},
		},
	})

	evm = test.NewEVM(test.NewMemoryStateDB(), vm.Config{}, vm.TxContext{}, test.NewBlockContext(), app)
	contract = vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(voteAddr), big.NewInt(0), 10000000)
	vote, _ = erc20vote.NewERC20Vote(evm, contract, false)
	vote.Init(tokenName, tokenSymbol, version, test.From)
	contract = vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(govAddr), big.NewInt(0), 10000000)
	governance, _ = NewGovernance(evm, contract, false)
	governance.Init(govName, version, voteDelay, votePeriod, big.NewInt(1), big.NewInt(1), test.From, voteAddr)
	contract = vm.NewContract(vm.AccountRef(test.From), vm.AccountRef(tokenAddr), big.NewInt(0), 10000000)
	token, _ = erc20.NewERC20(evm, contract, false)
	token.Init(tokenName, tokenSymbol, govAddr)
}

func TestParams(t *testing.T) {
	name, _ := governance.Name()
	require.Equal(t, govName, name)
	govVersion, _ := governance.Version()
	require.Equal(t, version, govVersion)
}

func TestVoteFlow(t *testing.T) {
	sk, _ := crypto.GenerateKey()
	signer := crypto.PubkeyToAddress(sk.PublicKey)
	sign := func(proposalId *big.Int, support uint8) (common.Hash, common.Hash, uint8) {
		data, _ := abi2.Encode([]interface{}{BALLOT_TYPEHASH, proposalId, support}, CastVoteType)
		hash := governance.eip712.HashTypedData(crypto.Keccak256Hash(data))
		sig, _ := crypto.Sign(hash.Bytes(), sk)
		r, s, v := common.BytesToHash(sig[:32]), common.BytesToHash(sig[32:64]), sig[64]
		return r, s, v
	}

	vote.Mint(test.From, big.NewInt(1000))
	vote.Mint(test.OneAddr, big.NewInt(1000))
	vote.Mint(signer, big.NewInt(1000))
	totalSupply, _ := vote.TotalSupply()
	calldata, _ := erc20.Abi.Pack("mint", test.OneAddr, amount)
	targets := []common.Address{
		tokenAddr,
	}
	values := []*big.Int{
		big.NewInt(0),
	}
	calldatas := [][]byte{
		calldata,
	}
	description := "hello"
	blockContext := test.NewBlockContext()
	blockContext.BlockNumber.Add(blockContext.BlockNumber, big.NewInt(1))
	evm := test.NewEVM(evm.StateDB.(*state.StateDB), evm.GetVMConfig(), evm.TxContext, blockContext, app)
	contract := vm.NewContract(vm.AccountRef(test.OneAddr), vm.AccountRef(govAddr), big.NewInt(0), 10000000)
	voterGov, _ := NewGovernance(evm, contract, false)

	proposalId, err := voterGov.Propose(targets, values, calldatas, description)
	require.Nil(t, err)

	deadline, err := voterGov.ProposalDeadline(proposalId)
	require.Nil(t, err)
	require.Equal(t, new(big.Int).Add(evm.Context.BlockNumber, new(big.Int).Add(voteDelay, votePeriod)), deadline)

	snapshot, err := voterGov.ProposalSnapshot(proposalId)
	require.Nil(t, err)
	require.Equal(t, new(big.Int).Add(evm.Context.BlockNumber, voteDelay), snapshot)

	require.Panics(t, func() {
		voterGov.CastVote(proposalId, For)
	})

	evm.Context.BlockNumber.Add(evm.Context.BlockNumber, voteDelay)
	evm.Context.BlockNumber.Add(evm.Context.BlockNumber, voteDelay)

	weight, err := voterGov.CastVote(proposalId, For)
	require.Nil(t, err)
	require.Equal(t, big.NewInt(1000), weight)
	r, s, v := sign(proposalId, For)
	weight, err = voterGov.CastVoteBySig(proposalId, For, v, r, s)
	require.Nil(t, err)
	require.Equal(t, big.NewInt(1000), weight)
	voterGov, _ = NewGovernance(evm, contract, false)
	state, err := voterGov.State(proposalId)
	require.Nil(t, err)
	require.Equal(t, Active, state)

	evm.Context.BlockNumber.Add(evm.Context.BlockNumber, votePeriod)
	state, err = voterGov.State(proposalId)
	require.Nil(t, err)
	require.Equal(t, Succeeded, state)

	proposalId, err = voterGov.Execute(targets, values, calldatas, crypto.Keccak256Hash([]byte(description)))
	require.Nil(t, err)
	balance, err := token.BalanceOf(test.OneAddr)
	require.Nil(t, err)
	require.Equal(t, amount, balance)
	numberator, err := voterGov.QuorumNumerator()
	require.Equal(t, big.NewInt(1), numberator)
	quorum, err := voterGov.Quorum(big.NewInt(10))
	require.Equal(t, totalSupply.Div(totalSupply, big.NewInt(100)), quorum)
	denominator, err := voterGov.QuorumDenominator()
	require.Equal(t, big.NewInt(100), denominator)
}
