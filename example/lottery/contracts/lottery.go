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
	LotteryABI  = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"vrfStore\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"drawing\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"guessing\",\"inputs\":[{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Drawing\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"bonus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Guessing\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]"
	LotteryCode = "60806040523480156200001157600080fd5b5060405162000fdb38038062000fdb8339810160408190526200003491620001ed565b8151829082906200004d90600390602085019062000090565b5080516200006390600490602084019062000090565b5050600680546001600160a01b0319166001600160a01b03959095169490941790935550620002ca915050565b8280546200009e9062000277565b90600052602060002090601f016020900481019282620000c257600085556200010d565b82601f10620000dd57805160ff19168380011785556200010d565b828001600101855582156200010d579182015b828111156200010d578251825591602001919060010190620000f0565b506200011b9291506200011f565b5090565b5b808211156200011b576000815560010162000120565b600082601f8301126200014857600080fd5b81516001600160401b0380821115620001655762000165620002b4565b604051601f8301601f19908116603f01168101908282118183101715620001905762000190620002b4565b81604052838152602092508683858801011115620001ad57600080fd5b600091505b83821015620001d15785820183015181830184015290820190620001b2565b83821115620001e35760008385830101525b9695505050505050565b6000806000606084860312156200020357600080fd5b83516001600160a01b03811681146200021b57600080fd5b60208501519093506001600160401b03808211156200023957600080fd5b620002478783880162000136565b935060408601519150808211156200025e57600080fd5b506200026d8682870162000136565b9150509250925092565b600181811c908216806200028c57607f821691505b60208210811415620002ae57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b610d0180620002da6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80633ae7358d1161008c578063a457c2d711610066578063a457c2d7146101a2578063a9059cbb146101b5578063b77200b1146101c8578063dd62ed3e146101d057600080fd5b80633ae7358d1461015c57806370a082311461017157806395d89b411461019a57600080fd5b806306fdde03146100d4578063095ea7b3146100f257806318160ddd1461011557806323b872dd14610127578063313ce5671461013a5780633950935114610149575b600080fd5b6100dc6101e3565b6040516100e99190610bc5565b60405180910390f35b610105610100366004610b69565b610275565b60405190151581526020016100e9565b6002545b6040519081526020016100e9565b610105610135366004610b2d565b61028f565b604051601281526020016100e9565b610105610157366004610b69565b6102b3565b61016f61016a366004610bac565b6102d5565b005b61011961017f366004610adf565b6001600160a01b031660009081526020819052604090205490565b6100dc610371565b6101056101b0366004610b69565b610380565b6101056101c3366004610b69565b610400565b61016f61040e565b6101196101de366004610afa565b61061d565b6060600380546101f290610c49565b80601f016020809104026020016040519081016040528092919081815260200182805461021e90610c49565b801561026b5780601f106102405761010080835404028352916020019161026b565b820191906000526020600020905b81548152906001019060200180831161024e57829003601f168201915b5050505050905090565b600033610283818585610648565b60019150505b92915050565b60003361029d85828561076c565b6102a88585856107e6565b506001949350505050565b6000336102838185856102c6838361061d565b6102d09190610c1a565b610648565b4360009081526005602090815260408083208151808301835233808252818501878152835460018082018655858952978790209351600290910290930180546001600160a01b0319166001600160a01b03909416939093178355519190950155815193845291830184905290917f1506736cb302149d691074ef1881cc011675c4737d3fa169ab3aa0fc74c2030c910160405180910390a15050565b6060600480546101f290610c49565b6000338161038e828661061d565b9050838110156103f35760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102a88286868403610648565b6000336102838185856107e6565b6006546000906001600160a01b0316633d46b81961042d600143610c32565b6040518263ffffffff1660e01b815260040161044b91815260200190565b602060405180830381600087803b15801561046557600080fd5b505af1158015610479573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061049d9190610b93565b905060006005816104af600143610c32565b8152602081019190915260400160002080549091506104cc575050565b6000816000815481106104e1576104e1610cb5565b906000526020600020906002020190506000600190505b825481101561057b576001820154610510908561098a565b61054084838154811061052557610525610cb5565b9060005260206000209060020201600101548660001c61098a565b10156105695782818154811061055857610558610cb5565b906000526020600020906002020191505b8061057381610c84565b9150506104f8565b508054610594906001600160a01b0316629896806109b6565b80546001820154604080516001600160a01b039093168352602083019190915281018490526298968060608201527f696d5367e8a6c822e93ca923e1eac0acf5433369e16d9db3824dc11bfa2e085f9060800160405180910390a1600560006105fe600143610c32565b815260200190815260200160002060006106189190610a75565b505050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166106aa5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103ea565b6001600160a01b03821661070b5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103ea565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6000610778848461061d565b905060001981146107e057818110156107d35760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103ea565b6107e08484848403610648565b50505050565b6001600160a01b03831661084a5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103ea565b6001600160a01b0382166108ac5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103ea565b6001600160a01b038316600090815260208190526040902054818110156109245760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103ea565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36107e0565b6000818311156109a55761099e8284610c32565b9050610289565b6109af8383610c32565b9392505050565b6001600160a01b038216610a0c5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103ea565b8060026000828254610a1e9190610c1a565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b5080546000825560020290600052602060002090810190610a969190610a99565b50565b5b80821115610abf5780546001600160a01b031916815560006001820155600201610a9a565b5090565b80356001600160a01b0381168114610ada57600080fd5b919050565b600060208284031215610af157600080fd5b6109af82610ac3565b60008060408385031215610b0d57600080fd5b610b1683610ac3565b9150610b2460208401610ac3565b90509250929050565b600080600060608486031215610b4257600080fd5b610b4b84610ac3565b9250610b5960208501610ac3565b9150604084013590509250925092565b60008060408385031215610b7c57600080fd5b610b8583610ac3565b946020939093013593505050565b600060208284031215610ba557600080fd5b5051919050565b600060208284031215610bbe57600080fd5b5035919050565b600060208083528351808285015260005b81811015610bf257858101830151858201604001528201610bd6565b81811115610c04576000604083870101525b50601f01601f1916929092016040019392505050565b60008219821115610c2d57610c2d610c9f565b500190565b600082821015610c4457610c44610c9f565b500390565b600181811c90821680610c5d57607f821691505b60208210811415610c7e57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415610c9857610c98610c9f565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052603260045260246000fdfea2646970667358221220a049f4c7a602fd0de9fa9043468eb4b51917e86ceca4fb7950b03b37c3057a1764736f6c63430008070033"
)

type LotteryGenesisCaller struct {
	evmFunc func(address common.Address) *vm.EVM
	abi     *abi.ABI
	to      common.Address
	caller  common.Address
}

func NewLotteryGenesisCaller(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig) (*LotteryGenesisCaller, error) {
	abi, err := abi.JSON(strings.NewReader(LotteryABI))
	if err != nil {
		return nil, err
	}
	return &LotteryGenesisCaller{
		evmFunc: func(address common.Address) *vm.EVM {
			return vm2.NewEVM(vm2.NewGenesisBlockContext(), address, db, chainConfig, nil)
		},
		abi: &abi,
	}, nil
}

func (c *LotteryGenesisCaller) DeployLottery(vrfStore common.Address, name string, symbol string) error {

	var data []byte
	var err error

	data, err = c.abi.Constructor.Inputs.Pack(vrfStore, name, symbol)
	if err != nil {
		return err
	}

	evm := c.evmFunc(c.caller)
	_, _, _, err = evm.CreateAppContract(vm.AccountRef(c.caller), append(common.FromHex(LotteryCode), data...), c.to, math.MaxUint64, big.NewInt(0))
	return err
}

func (c *LotteryGenesisCaller) PackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	input, err := c.abi.Pack("allowance", owner, spender)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackAllowance(owner, spender)
	if err != nil {
		return *new(*big.Int), err
	}
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

func (c *LotteryGenesisCaller) PackBalanceOf(account common.Address) ([]byte, error) {
	input, err := c.abi.Pack("balanceOf", account)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) BalanceOf(account common.Address) (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackBalanceOf(account)
	if err != nil {
		return *new(*big.Int), err
	}
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

func (c *LotteryGenesisCaller) PackDecimals() ([]byte, error) {
	input, err := c.abi.Pack("decimals")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Decimals() (uint8, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackDecimals()
	if err != nil {
		return *new(uint8), err
	}
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

func (c *LotteryGenesisCaller) PackName() ([]byte, error) {
	input, err := c.abi.Pack("name")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Name() (string, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackName()
	if err != nil {
		return *new(string), err
	}
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

func (c *LotteryGenesisCaller) PackSymbol() ([]byte, error) {
	input, err := c.abi.Pack("symbol")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Symbol() (string, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackSymbol()
	if err != nil {
		return *new(string), err
	}
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

func (c *LotteryGenesisCaller) PackTotalSupply() ([]byte, error) {
	input, err := c.abi.Pack("totalSupply")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) TotalSupply() (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackTotalSupply()
	if err != nil {
		return *new(*big.Int), err
	}
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

func (c *LotteryGenesisCaller) PackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("approve", spender, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Approve(spender common.Address, amount *big.Int) (bool, error) {
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

func (c *LotteryGenesisCaller) PackDecreaseAllowance(spender common.Address, subtractedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("decreaseAllowance", spender, subtractedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
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

func (c *LotteryGenesisCaller) PackDrawing() ([]byte, error) {
	input, err := c.abi.Pack("drawing")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Drawing() error {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackDrawing()
	if err != nil {
		return err
	}

	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))

	return err
}

func (c *LotteryGenesisCaller) PackGuessing(number *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("guessing", number)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Guessing(number *big.Int) error {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackGuessing(number)
	if err != nil {
		return err
	}

	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))

	return err
}

func (c *LotteryGenesisCaller) PackIncreaseAllowance(spender common.Address, addedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("increaseAllowance", spender, addedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
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

func (c *LotteryGenesisCaller) PackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transfer", to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) Transfer(to common.Address, amount *big.Int) (bool, error) {
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

func (c *LotteryGenesisCaller) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transferFrom", from, to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryGenesisCaller) TransferFrom(from common.Address, to common.Address, amount *big.Int) (bool, error) {
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

func (c *LotteryGenesisCaller) ABI() *abi.ABI {
	return c.abi
}
func (c *LotteryGenesisCaller) WithCaller(caller common.Address) *LotteryGenesisCaller {
	c.caller = caller
	return c
}

func (c *LotteryGenesisCaller) WithTo(to common.Address) *LotteryGenesisCaller {
	c.to = to
	return c
}
