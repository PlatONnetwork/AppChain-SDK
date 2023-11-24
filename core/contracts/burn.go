package contracts

import "github.com/PlatONnetwork/PlatON-Go/core/vm"

type Burn interface {
	UseGas(gas uint64)
	Gas() uint64
}

type Burner struct {
	contract *vm.Contract
}

func NewBurner(contract *vm.Contract) Burn {
	return &Burner{
		contract: contract,
	}
}

func (b *Burner) UseGas(gas uint64) {
	if !b.contract.UseGas(gas) {
		panic(vm.ErrOutOfGas)
	}
}

func (b *Burner) Gas() uint64 {
	return b.contract.Gas
}
