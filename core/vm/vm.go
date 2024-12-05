package vm

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/core/rawdb"
	"github.com/PlatONnetwork/PlatON-Go/core/state"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

func NewEVM(blockCtx vm.BlockContext, origin common.Address, statedb vm.StateDB, chainConfig *params.ChainConfig, app vm.ContractsApp) *vm.EVM {
	evm := vm.NewEVM(blockCtx, vm.TxContext{
		Origin:   origin,
		GasPrice: big.NewInt(0),
	}, statedb, chainConfig, vm.Config{}, app)
	return evm
}

func NewGenesisBlockContext() vm.BlockContext {
	blockContext := vm.BlockContext{
		CanTransfer: func(db vm.StateDB, address common.Address, b *big.Int) bool {
			return core.CanTransfer(db, address, b)

		},
		Transfer: func(db vm.StateDB, address common.Address, address2 common.Address, b *big.Int) {
			core.Transfer(db, address, address2, b)
		},
		GasLimit:    math.MaxUint64,
		BlockNumber: big.NewInt(0),
		Time:        new(big.Int).SetUint64(0),
		BlockHash:   common.ZeroHash,
		Difficulty:  new(big.Int).SetUint64(0), // This one must not be deleted, otherwise the solidity contract will be failed
		Random:      &common.ZeroHash,
	}
	return blockContext
}
func NewBlockContext(header *types.Header, chain core.ChainContext) vm.BlockContext {
	var baseFee *big.Int
	if header.BaseFee != nil {
		baseFee = new(big.Int).Set(header.BaseFee)
	}
	blockContext := vm.BlockContext{
		CanTransfer: func(db vm.StateDB, address common.Address, b *big.Int) bool {
			return core.CanTransfer(db, address, b)

		},
		Transfer: func(db vm.StateDB, address common.Address, address2 common.Address, b *big.Int) {
			core.Transfer(db, address, address2, b)
		},
		GasLimit:    math.MaxUint64,
		Coinbase:    header.Coinbase,
		BlockNumber: new(big.Int).Set(header.Number),
		Time:        new(big.Int).SetUint64(header.Time),
		BaseFee:     baseFee,
		BlockHash:   common.ZeroHash,
		Difficulty:  new(big.Int).SetUint64(0), // This one must not be deleted, otherwise the solidity contract will be failed
		Random:      &common.ZeroHash,
		Nonce:       header.Nonce,
		ParentHash:  header.ParentHash,
	}
	if chain != nil {
		blockContext.GetHash = core.GetHashFn(header, chain)
		blockContext.GetNonce = core.GetNonceFn(header, chain)
	}
	return blockContext
}
func MustDeployCode(rawCode []byte, input []byte) []byte {
	deployCode, err := DeployCode(rawCode, input)
	if err != nil {
		panic(err)
	}
	return deployCode
}

func DeployCode(rawCode []byte, input []byte) ([]byte, error) {
	blockContext := vm.BlockContext{
		CanTransfer: func(db vm.StateDB, address common.Address, b *big.Int) bool {
			return core.CanTransfer(db, address, b)

		},
		Transfer: func(db vm.StateDB, address common.Address, address2 common.Address, b *big.Int) {
			core.Transfer(db, address, address2, b)
		},
		GetHash: func(u uint64) common.Hash {
			return common.BytesToHash(crypto.Keccak256([]byte(big.NewInt(int64(u)).String())))
		},
		BlockNumber: new(big.Int).SetUint64(10),
		Time:        new(big.Int).SetUint64(1),
		GasLimit:    100000000,
	}
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
	evm := vm.NewEVM(blockContext, vm.TxContext{}, statedb, params.MainnetChainConfig, vm.Config{}, nil)
	data := rawCode
	if input != nil {
		data = append(data, input...)
	}
	deployCode, _, _, err := evm.Create(vm.AccountRef(common.ZeroAddr), data, 1000000000000, big.NewInt(0))
	return deployCode, err
}

type BackendCallerContext interface {
	Backend() sdk.Backend
	StateDB() sdk.StateDB
	Header() *types.Header
}
