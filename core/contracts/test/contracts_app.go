package test

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type SDKContract struct {
	Addr    common.Address
	RunFunc func(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error)
}
type ContractsApp struct {
	ContractModules []*SDKContract
}

func (c *ContractsApp) Contracts(statedb sdk.StateDBReader, blockNumber uint64) []vm.SDKContract {
	var contracts []vm.SDKContract
	for _, c := range c.ContractModules {
		contracts = append(contracts, c)
	}
	return contracts
}

func (s *SDKContract) Address() common.Address {
	return s.Addr
}
func (s *SDKContract) Run(evm *vm.EVM, contract *vm.Contract, input []byte, readOnly bool) ([]byte, error) {
	return s.RunFunc(evm, contract, input, readOnly)
}

func (s *SDKContract) ContractCreateBlockNumber(statedb sdk.StateDBReader) uint64 {
	return 0
}

func NewContractsApp(contracts []*SDKContract) vm.ContractsApp {
	return &ContractsApp{
		ContractModules: contracts,
	}
}
