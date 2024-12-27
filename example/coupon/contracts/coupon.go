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
	CouponABI  = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"applyCoupon\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"roundBalance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Apply\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"balance\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]"
	CouponCode = "60806040523480156200001157600080fd5b5060405162000d5838038062000d588339810160408190526200003491620001cb565b8151829082906200004d9060039060208501906200006e565b508051620000639060049060208401906200006e565b505050505062000288565b8280546200007c9062000235565b90600052602060002090601f016020900481019282620000a05760008555620000eb565b82601f10620000bb57805160ff1916838001178555620000eb565b82800160010185558215620000eb579182015b82811115620000eb578251825591602001919060010190620000ce565b50620000f9929150620000fd565b5090565b5b80821115620000f95760008155600101620000fe565b600082601f8301126200012657600080fd5b81516001600160401b038082111562000143576200014362000272565b604051601f8301601f19908116603f011681019082821181831017156200016e576200016e62000272565b816040528381526020925086838588010111156200018b57600080fd5b600091505b83821015620001af578582018301518183018401529082019062000190565b83821115620001c15760008385830101525b9695505050505050565b60008060408385031215620001df57600080fd5b82516001600160401b0380821115620001f757600080fd5b620002058683870162000114565b935060208501519150808211156200021c57600080fd5b506200022b8582860162000114565b9150509250929050565b600181811c908216806200024a57607f821691505b602082108114156200026c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052604160045260246000fd5b610ac080620002986000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c806357e871e71161008c578063a457c2d711610066578063a457c2d7146101ba578063a9059cbb146101cd578063dd62ed3e146101e0578063df6c2767146101f357600080fd5b806357e871e71461018057806370a082311461018957806395d89b41146101b257600080fd5b806323b872dd116100c857806323b872dd14610142578063313ce5671461015557806339509351146101645780633e9317da1461017757600080fd5b806306fdde03146100ef578063095ea7b31461010d57806318160ddd14610130575b600080fd5b6100f76101fd565b604051610104919061099f565b60405180910390f35b61012061011b366004610975565b61028f565b6040519015158152602001610104565b6002545b604051908152602001610104565b610120610150366004610939565b6102a7565b60405160128152602001610104565b610120610172366004610975565b6102cb565b61013460065481565b61013460055481565b6101346101973660046108e4565b6001600160a01b031660009081526020819052604090205490565b6100f76102ed565b6101206101c8366004610975565b6102fc565b6101206101db366004610975565b61037c565b6101346101ee366004610906565b61038a565b6101fb6103b5565b005b60606003805461020c90610a23565b80601f016020809104026020016040519081016040528092919081815260200182805461023890610a23565b80156102855780601f1061025a57610100808354040283529160200191610285565b820191906000526020600020905b81548152906001019060200180831161026857829003601f168201915b5050505050905090565b60003361029d818585610496565b5060019392505050565b6000336102b58582856105ba565b6102c0858585610634565b506001949350505050565b60003361029d8185856102de838361038a565b6102e891906109f4565b610496565b60606004805461020c90610a23565b6000338161030a828661038a565b90508381101561036f5760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102c08286868403610496565b60003361029d818585610634565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b43600554146103cb57633b9aca00600655436005555b6000806103e4600a6006546107d890919063ffffffff16565b91509150816104355760405162461bcd60e51b815260206004820152601f60248201527f436f75706f6e3a2062616c616e636520697320696e73756666696369656e74006044820152606401610366565b61043f3382610809565b8060065461044d9190610a0c565b6006819055604080513381526020810184905280820192909252517fdafdd0f2a3af66f1d9ab59928d5fb2e9eb5428db4b63362502d21e2511624ff79181900360600190a15050565b6001600160a01b0383166104f85760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610366565b6001600160a01b0382166105595760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610366565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006105c6848461038a565b9050600019811461062e57818110156106215760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610366565b61062e8484848403610496565b50505050565b6001600160a01b0383166106985760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610366565b6001600160a01b0382166106fa5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610366565b6001600160a01b038316600090815260208190526040902054818110156107725760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610366565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361062e565b600080826107eb57506000905080610802565b60018385816107fc576107fc610a74565b04915091505b9250929050565b6001600160a01b03821661085f5760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610366565b806002600082825461087191906109f4565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b80356001600160a01b03811681146108df57600080fd5b919050565b6000602082840312156108f657600080fd5b6108ff826108c8565b9392505050565b6000806040838503121561091957600080fd5b610922836108c8565b9150610930602084016108c8565b90509250929050565b60008060006060848603121561094e57600080fd5b610957846108c8565b9250610965602085016108c8565b9150604084013590509250925092565b6000806040838503121561098857600080fd5b610991836108c8565b946020939093013593505050565b600060208083528351808285015260005b818110156109cc578581018301518582016040015282016109b0565b818111156109de576000604083870101525b50601f01601f1916929092016040019392505050565b60008219821115610a0757610a07610a5e565b500190565b600082821015610a1e57610a1e610a5e565b500390565b600181811c90821680610a3757607f821691505b60208210811415610a5857634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fdfea2646970667358221220287933c1dc9d3cc0711ae94be2cc28fb217b8baca129587cb95a586485b7aa8664736f6c63430008070033"
)

type CouponGenesisCaller struct {
	evmFunc func(address common.Address) *vm.EVM
	abi     *abi.ABI
	to      common.Address
	caller  common.Address
}

func NewCouponGenesisCaller(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig) (*CouponGenesisCaller, error) {
	abi, err := abi.JSON(strings.NewReader(CouponABI))
	if err != nil {
		return nil, err
	}
	return &CouponGenesisCaller{
		evmFunc: func(address common.Address) *vm.EVM {
			return vm2.NewEVM(vm2.NewGenesisBlockContext(), address, db, chainConfig, nil)
		},
		abi: &abi,
	}, nil
}

func (c *CouponGenesisCaller) DeployCoupon(name string, symbol string) error {

	var data []byte
	var err error

	data, err = c.abi.Constructor.Inputs.Pack(name, symbol)
	if err != nil {
		return err
	}

	evm := c.evmFunc(c.caller)
	_, _, _, err = evm.CreateAppContract(vm.AccountRef(c.caller), append(common.FromHex(CouponCode), data...), c.to, math.MaxUint64, big.NewInt(0))
	return err
}

func (c *CouponGenesisCaller) PackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	input, err := c.abi.Pack("allowance", owner, spender)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
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

func (c *CouponGenesisCaller) PackBalanceOf(account common.Address) ([]byte, error) {
	input, err := c.abi.Pack("balanceOf", account)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) BalanceOf(account common.Address) (*big.Int, error) {
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

func (c *CouponGenesisCaller) PackBlockNumber() ([]byte, error) {
	input, err := c.abi.Pack("blockNumber")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) BlockNumber() (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackBlockNumber()
	if err != nil {
		return *new(*big.Int), err
	}
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(*big.Int), err
	}

	out, err := c.abi.Unpack("blockNumber", output)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

func (c *CouponGenesisCaller) PackDecimals() ([]byte, error) {
	input, err := c.abi.Pack("decimals")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) Decimals() (uint8, error) {
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

func (c *CouponGenesisCaller) PackName() ([]byte, error) {
	input, err := c.abi.Pack("name")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) Name() (string, error) {
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

func (c *CouponGenesisCaller) PackRoundBalance() ([]byte, error) {
	input, err := c.abi.Pack("roundBalance")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) RoundBalance() (*big.Int, error) {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte
	var output []byte
	input, err = c.PackRoundBalance()
	if err != nil {
		return *new(*big.Int), err
	}
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	if err != nil {
		return *new(*big.Int), err
	}

	out, err := c.abi.Unpack("roundBalance", output)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

func (c *CouponGenesisCaller) PackSymbol() ([]byte, error) {
	input, err := c.abi.Pack("symbol")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) Symbol() (string, error) {
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

func (c *CouponGenesisCaller) PackTotalSupply() ([]byte, error) {
	input, err := c.abi.Pack("totalSupply")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) TotalSupply() (*big.Int, error) {
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

func (c *CouponGenesisCaller) PackApplyCoupon() ([]byte, error) {
	input, err := c.abi.Pack("applyCoupon")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) ApplyCoupon() error {
	evm := c.evmFunc(c.caller)

	var err error
	var input []byte

	input, err = c.PackApplyCoupon()
	if err != nil {
		return err
	}

	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))

	return err
}

func (c *CouponGenesisCaller) PackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("approve", spender, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) Approve(spender common.Address, amount *big.Int) (bool, error) {
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

func (c *CouponGenesisCaller) PackDecreaseAllowance(spender common.Address, subtractedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("decreaseAllowance", spender, subtractedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (bool, error) {
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

func (c *CouponGenesisCaller) PackIncreaseAllowance(spender common.Address, addedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("increaseAllowance", spender, addedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) IncreaseAllowance(spender common.Address, addedValue *big.Int) (bool, error) {
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

func (c *CouponGenesisCaller) PackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transfer", to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) Transfer(to common.Address, amount *big.Int) (bool, error) {
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

func (c *CouponGenesisCaller) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transferFrom", from, to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *CouponGenesisCaller) TransferFrom(from common.Address, to common.Address, amount *big.Int) (bool, error) {
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

func (c *CouponGenesisCaller) ABI() *abi.ABI {
	return c.abi
}
func (c *CouponGenesisCaller) WithCaller(caller common.Address) *CouponGenesisCaller {
	c.caller = caller
	return c
}

func (c *CouponGenesisCaller) WithTo(to common.Address) *CouponGenesisCaller {
	c.to = to
	return c
}
