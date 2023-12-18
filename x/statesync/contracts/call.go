package contracts

import (
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"math/big"
	"strings"
)

var (
	abiJson              = `[{"inputs":[{"internalType":"uint256","name":"","type":"uint256"},{"internalType":"address","name":"sender","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"onStateReceive","outputs":[],"stateMutability":"nonpayable","type":"function"}]`
	onStateReceiveAbi, _ = abi.JSON(strings.NewReader(abiJson))
)

func CallOnStateReceive(evm *vm.EVM, contract *vm.Contract, gas uint64, obj *StateSync) ([]byte, error) {
	method := onStateReceiveAbi.Methods["onStateReceive"]
	input, err := method.Inputs.Pack(obj.Id, obj.Sender, obj.Data)
	if err != nil {
		return nil, err
	}
	input = append(method.ID, input...)
	return contracts.Call(evm, contract, obj.Receiver, input, gas, big.NewInt(0))
}
