package contracts

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
	BenchTokenGenesisABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]"
	BenchTokenCode       = "60806040523480156200001157600080fd5b5060405162000e0038038062000e0083398101604081905262000034916200022a565b6040518060400160405280600a8152602001692132b731b42a37b5b2b760b11b8152506040518060400160405280600381526020016208486960eb1b81525081600390805190602001906200008b92919062000184565b508051620000a190600490602084019062000184565b505050620000b68282620000be60201b60201c565b5050620002ca565b6001600160a01b038216620001195760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b80600260008282546200012d919062000266565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b82805462000192906200028d565b90600052602060002090601f016020900481019282620001b6576000855562000201565b82601f10620001d157805160ff191683800117855562000201565b8280016001018555821562000201579182015b8281111562000201578251825591602001919060010190620001e4565b506200020f92915062000213565b5090565b5b808211156200020f576000815560010162000214565b600080604083850312156200023e57600080fd5b82516001600160a01b03811681146200025657600080fd5b6020939093015192949293505050565b600082198211156200028857634e487b7160e01b600052601160045260246000fd5b500190565b600181811c90821680620002a257607f821691505b60208210811415620002c457634e487b7160e01b600052602260045260246000fd5b50919050565b610b2680620002da6000396000f3fe608060405234801561001057600080fd5b50600436106100c95760003560e01c80633950935111610081578063a457c2d71161005b578063a457c2d714610194578063a9059cbb146101a7578063dd62ed3e146101ba57600080fd5b8063395093511461014357806370a082311461015657806395d89b411461018c57600080fd5b806318160ddd116100b257806318160ddd1461010f57806323b872dd14610121578063313ce5671461013457600080fd5b806306fdde03146100ce578063095ea7b3146100ec575b600080fd5b6100d6610200565b6040516100e391906109ea565b60405180910390f35b6100ff6100fa3660046109c0565b610292565b60405190151581526020016100e3565b6002545b6040519081526020016100e3565b6100ff61012f366004610984565b6102aa565b604051601281526020016100e3565b6100ff6101513660046109c0565b6102ce565b61011361016436600461092f565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6100d661031a565b6100ff6101a23660046109c0565b610329565b6100ff6101b53660046109c0565b6103ff565b6101136101c8366004610951565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b60606003805461020f90610a9c565b80601f016020809104026020016040519081016040528092919081815260200182805461023b90610a9c565b80156102885780601f1061025d57610100808354040283529160200191610288565b820191906000526020600020905b81548152906001019060200180831161026b57829003601f168201915b5050505050905090565b6000336102a081858561040d565b5060019392505050565b6000336102b88582856105c0565b6102c3858585610697565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091906102a09082908690610315908790610a5d565b61040d565b60606004805461020f90610a9c565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff87168452909152812054909190838110156103f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102c3828686840361040d565b6000336102a0818585610697565b73ffffffffffffffffffffffffffffffffffffffff83166104af576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f726573730000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff8216610552576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f737300000000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146106915781811015610684576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103e9565b610691848484840361040d565b50505050565b73ffffffffffffffffffffffffffffffffffffffff831661073a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f647265737300000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff82166107dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f657373000000000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610893576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e6365000000000000000000000000000000000000000000000000000060648201526084016103e9565b73ffffffffffffffffffffffffffffffffffffffff848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a3610691565b803573ffffffffffffffffffffffffffffffffffffffff8116811461092a57600080fd5b919050565b60006020828403121561094157600080fd5b61094a82610906565b9392505050565b6000806040838503121561096457600080fd5b61096d83610906565b915061097b60208401610906565b90509250929050565b60008060006060848603121561099957600080fd5b6109a284610906565b92506109b060208501610906565b9150604084013590509250925092565b600080604083850312156109d357600080fd5b6109dc83610906565b946020939093013593505050565b600060208083528351808285015260005b81811015610a17578581018301518582016040015282016109fb565b81811115610a29576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008219821115610a97577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b500190565b600181811c90821680610ab057607f821691505b60208210811415610aea577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b5091905056fea26469706673582212200a9e719132b0e157139134318e7df3b4e77dbfb493ffef3f355f41d9a1cd5d0e64736f6c63430008070033"
)

type BenchTokenGenesisCaller struct {
	evmFunc func(address common.Address) *vm.EVM
	abi     *abi.ABI
	to      common.Address
	caller  common.Address
}

func NewBenchTokenGenesisCaller(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig) (*BenchTokenGenesisCaller, error) {
	abi, err := abi.JSON(strings.NewReader(BenchTokenGenesisABI))
	if err != nil {
		return nil, err
	}
	return &BenchTokenGenesisCaller{
		evmFunc: func(address common.Address) *vm.EVM {
			return vm2.NewEVM(vm2.NewGenesisBlockContext(), address, db, chainConfig, nil)
		},
		abi: &abi,
	}, nil
}

func (c *BenchTokenGenesisCaller) DeployBenchToken(recipient common.Address, initialSupply *big.Int) error {

	var data []byte
	var err error

	data, err = c.abi.Constructor.Inputs.Pack(recipient, initialSupply)
	if err != nil {
		return err
	}

	evm := c.evmFunc(c.caller)
	_, _, _, err = evm.CreateAppContract(vm.AccountRef(c.caller), append(common.FromHex(BenchTokenCode), data...), c.to, math.MaxUint64, big.NewInt(0))
	return err
}

func (c *BenchTokenGenesisCaller) PackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	input, err := c.abi.Pack("allowance", owner, spender)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	input, err = c.PackAllowance(owner, spender)
	if err != nil {
		return *new(*big.Int), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(*big.Int), err
	}
	out, err := c.abi.Unpack("allowance", output)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackBalanceOf(account common.Address) ([]byte, error) {
	input, err := c.abi.Pack("balanceOf", account)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) BalanceOf(account common.Address) (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	input, err = c.PackBalanceOf(account)
	if err != nil {
		return *new(*big.Int), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(*big.Int), err
	}
	out, err := c.abi.Unpack("balanceOf", output)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackDecimals() ([]byte, error) {
	input, err := c.abi.Pack("decimals")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) Decimals() (uint8, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	input, err = c.PackDecimals()
	if err != nil {
		return *new(uint8), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(uint8), err
	}
	out, err := c.abi.Unpack("decimals", output)
	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackName() ([]byte, error) {
	input, err := c.abi.Pack("name")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) Name() (string, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	input, err = c.PackName()
	if err != nil {
		return *new(string), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(string), err
	}
	out, err := c.abi.Unpack("name", output)
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackSymbol() ([]byte, error) {
	input, err := c.abi.Pack("symbol")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) Symbol() (string, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	input, err = c.PackSymbol()
	if err != nil {
		return *new(string), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(string), err
	}
	out, err := c.abi.Unpack("symbol", output)
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackTotalSupply() ([]byte, error) {
	input, err := c.abi.Pack("totalSupply")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) TotalSupply() (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	input, err = c.PackTotalSupply()
	if err != nil {
		return *new(*big.Int), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(*big.Int), err
	}
	out, err := c.abi.Unpack("totalSupply", output)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("approve", spender, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) Approve(spender common.Address, amount *big.Int) (bool, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackApprove(spender, amount)
	if err != nil {
		return *new(bool), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(bool), err
	}
	out, err := c.abi.Unpack("approve", output)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackDecreaseAllowance(spender common.Address, subtractedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("decreaseAllowance", spender, subtractedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackDecreaseAllowance(spender, subtractedValue)
	if err != nil {
		return *new(bool), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(bool), err
	}
	out, err := c.abi.Unpack("decreaseAllowance", output)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackIncreaseAllowance(spender common.Address, addedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("increaseAllowance", spender, addedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackIncreaseAllowance(spender, addedValue)
	if err != nil {
		return *new(bool), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(bool), err
	}
	out, err := c.abi.Unpack("increaseAllowance", output)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transfer", to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) Transfer(to common.Address, amount *big.Int) (bool, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackTransfer(to, amount)
	if err != nil {
		return *new(bool), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(bool), err
	}
	out, err := c.abi.Unpack("transfer", output)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

func (c *BenchTokenGenesisCaller) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transferFrom", from, to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *BenchTokenGenesisCaller) TransferFrom(from common.Address, to common.Address, amount *big.Int) (bool, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackTransferFrom(from, to, amount)
	if err != nil {
		return *new(bool), err
	}

	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(bool), err
	}
	out, err := c.abi.Unpack("transferFrom", output)
	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err
}

func (c *BenchTokenGenesisCaller) ABI() *abi.ABI {
	return c.abi
}
func (c *BenchTokenGenesisCaller) WithCaller(caller common.Address) *BenchTokenGenesisCaller {
	c.caller = caller
	return c
}

func (c *BenchTokenGenesisCaller) WithTo(to common.Address) *BenchTokenGenesisCaller {
	c.to = to
	return c
}
func (c *BenchTokenGenesisCaller) WithEvmFunc(evmFunc func(address common.Address) *vm.EVM) *BenchTokenGenesisCaller {
	c.evmFunc = evmFunc
	return c
}
