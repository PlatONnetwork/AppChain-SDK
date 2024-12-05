package proxy

import (
	vm2 "github.com/PlatONnetwork/AppChain-SDK/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math"
	"math/big"
	"strings"
)

var (
	InitializableTransparentUpgradeableProxyABI  = "[{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"admin\",\"inputs\":[],\"outputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeAdmin\",\"inputs\":[{\"name\":\"newAdmin\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"implementation\",\"inputs\":[],\"outputs\":[{\"name\":\"implementation_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_logic\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false}]"
	InitializableTransparentUpgradeableProxyCode = "608060405234801561001057600080fd5b50610b10806100206000396000f3fe6080604052600436106100595760003560e01c80633659cfe6146100705780634f1ef286146100905780635c60da1b146100a35780638f283970146100d4578063cf7a1d77146100f4578063f851a4401461010757610068565b366100685761006661011c565b005b61006661011c565b34801561007c57600080fd5b5061006661008b366004610877565b610136565b61006661009e366004610964565b61017d565b3480156100af57600080fd5b506100b86101ee565b6040516001600160a01b03909116815260200160405180910390f35b3480156100e057600080fd5b506100666100ef366004610877565b610229565b610066610102366004610892565b610253565b34801561011357600080fd5b506100b861035c565b610124610387565b61013461012f610421565b61042b565b565b61013e61044f565b6001600160a01b0316336001600160a01b031614156101755761017281604051806020016040528060008152506000610482565b50565b61017261011c565b61018561044f565b6001600160a01b0316336001600160a01b031614156101e6576101e18383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525060019250610482915050565b505050565b6101e161011c565b60006101f861044f565b6001600160a01b0316336001600160a01b0316141561021e57610219610421565b905090565b61022661011c565b90565b61023161044f565b6001600160a01b0316336001600160a01b0316141561017557610172816104ad565b600061025d610421565b6001600160a01b0316146102f05760405162461bcd60e51b815260206004820152604960248201527f496e697469616c697a61626c655472616e73706172656e74557067726164656160448201527f626c6550726f78793a20636f6e747261637420697320616c726561647920696e6064820152681a5d1a585b1a5e995960ba1b608482015260a4015b60405180910390fd5b6102fa8382610501565b61032560017fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6104610a36565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61031461035357610353610a87565b6101e1826104ad565b600061036661044f565b6001600160a01b0316336001600160a01b0316141561021e5761021961044f565b61038f61044f565b6001600160a01b0316336001600160a01b031614156101345760405162461bcd60e51b815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f78792074617267606482015261195d60f21b608482015260a4016102e7565b600061021961056a565b3660008037600080366000845af43d6000803e80801561044a573d6000f35b3d6000fd5b60007fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b546001600160a01b0316919050565b61048b83610592565b6000825111806104985750805b156101e1576104a783836105d2565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6104d661044f565b604080516001600160a01b03928316815291841660208301520160405180910390a1610172816105fe565b61052c60017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd610a36565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc1461055a5761055a610a87565b61056682826000610482565b5050565b60007f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610473565b61059b816106a7565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60606105f78383604051806060016040528060278152602001610ab46027913961073b565b9392505050565b6001600160a01b0381166106635760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b60648201526084016102e7565b807fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381163b6107145760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016102e7565b807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc610686565b6060600080856001600160a01b03168560405161075891906109e7565b600060405180830381855af49150503d8060008114610793576040519150601f19603f3d011682016040523d82523d6000602084013e610798565b606091505b50915091506107a9868383876107b3565b9695505050505050565b6060831561081f578251610818576001600160a01b0385163b6108185760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016102e7565b5081610829565b6108298383610831565b949350505050565b8151156108415781518083602001fd5b8060405162461bcd60e51b81526004016102e79190610a03565b80356001600160a01b038116811461087257600080fd5b919050565b60006020828403121561088957600080fd5b6105f78261085b565b6000806000606084860312156108a757600080fd5b6108b08461085b565b92506108be6020850161085b565b9150604084013567ffffffffffffffff808211156108db57600080fd5b818601915086601f8301126108ef57600080fd5b81358181111561090157610901610a9d565b604051601f8201601f19908116603f0116810190838211818310171561092957610929610a9d565b8160405282815289602084870101111561094257600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b60008060006040848603121561097957600080fd5b6109828461085b565b9250602084013567ffffffffffffffff8082111561099f57600080fd5b818601915086601f8301126109b357600080fd5b8135818111156109c257600080fd5b8760208285010111156109d457600080fd5b6020830194508093505050509250925092565b600082516109f9818460208701610a5b565b9190910192915050565b6020815260008251806020840152610a22816040850160208701610a5b565b601f01601f19169190910160400192915050565b600082821015610a5657634e487b7160e01b600052601160045260246000fd5b500390565b60005b83811015610a76578181015183820152602001610a5e565b838111156104a75750506000910152565b634e487b7160e01b600052600160045260246000fd5b634e487b7160e01b600052604160045260246000fdfe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212205619a115fd28a99f78d94c45ad168a27e545e65588f705c495c7a33fbe59e45d64736f6c63430008070033"
)

type InitializableTransparentUpgradeableProxyGenesisCaller struct {
	evmFunc func(address common.Address) *vm.EVM
	abi     *abi.ABI
	to      common.Address
	caller  common.Address
}

func NewInitializableTransparentUpgradeableProxyGenesisCaller(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig) (*InitializableTransparentUpgradeableProxyGenesisCaller, error) {
	abi, err := abi.JSON(strings.NewReader(InitializableTransparentUpgradeableProxyABI))
	if err != nil {
		return nil, err
	}
	return &InitializableTransparentUpgradeableProxyGenesisCaller{
		evmFunc: func(address common.Address) *vm.EVM {
			return vm2.NewEVM(vm2.NewGenesisBlockContext(), address, db, chainConfig, nil)
		},
		abi: &abi,
	}, nil
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) DeployInitializableTransparentUpgradeableProxy() error {

	var data []byte

	evm := c.evmFunc(c.caller)
	_, _, _, err := evm.CreateAppContract(vm.AccountRef(c.caller), append(common.FromHex(InitializableTransparentUpgradeableProxyCode), data...), c.to, math.MaxUint64, big.NewInt(0))
	return err
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) Admin() (common.Address, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.abi.Pack("admin")
	if err != nil {
		return *new(common.Address), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(common.Address), err
	}
	out, err := c.abi.Unpack("admin", output)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) ChangeAdmin(newAdmin common.Address) error {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.abi.Pack("changeAdmin", newAdmin)
	if err != nil {
		return err
	}

	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))

	return err
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) Implementation() (common.Address, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.abi.Pack("implementation")
	if err != nil {
		return *new(common.Address), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(common.Address), err
	}
	out, err := c.abi.Unpack("implementation", output)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) Initialize(_logic common.Address, admin_ common.Address, _data []byte) error {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.abi.Pack("initialize", _logic, admin_, _data)
	if err != nil {
		return err
	}

	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))

	return err
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) UpgradeTo(newImplementation common.Address) error {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.abi.Pack("upgradeTo", newImplementation)
	if err != nil {
		return err
	}

	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))

	return err
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) UpgradeToAndCall(newImplementation common.Address, data []byte) error {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.abi.Pack("upgradeToAndCall", newImplementation, data)
	if err != nil {
		return err
	}

	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))

	return err
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) ABI() *abi.ABI {
	return c.abi
}
func (c *InitializableTransparentUpgradeableProxyGenesisCaller) WithCaller(caller common.Address) *InitializableTransparentUpgradeableProxyGenesisCaller {
	c.caller = caller
	return c
}

func (c *InitializableTransparentUpgradeableProxyGenesisCaller) WithTo(to common.Address) *InitializableTransparentUpgradeableProxyGenesisCaller {
	c.to = to
	return c
}
