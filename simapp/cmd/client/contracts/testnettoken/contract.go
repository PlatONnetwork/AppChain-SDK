// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testnettoken

import (
	"errors"
	"math/big"
	"strings"

	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TestnetTokenMetaData contains all meta data concerning the TestnetToken contract.
var TestnetTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x608060405234620002f15762000cef803803806200001d81620002f5565b928339810190604081830312620002f15780516001600160401b0390818111620002f157836200004f9184016200031b565b9160209384820151838111620002f1576200006b92016200031b565b928251908282116200020f575f54916001948584811c94168015620002e6575b83851014620001f0578190601f9485811162000293575b5083908583116001146200022f575f9262000223575b50505f19600383901b1c191690851b175f555b84519283116200020f5783548481811c9116801562000204575b82821014620001f057828111620001a8575b50809183116001146200014457508192935f9262000138575b50505f19600383901b1c191690821b1790555b600660025560405161096390816200038c8239f35b015190505f8062000110565b90601f19831694845f52825f20925f905b878210620001905750508385961062000177575b505050811b01905562000123565b01515f1960f88460031b161c191690555f808062000169565b80878596829496860151815501950193019062000155565b845f52815f208380860160051c820192848710620001e6575b0160051c019085905b828110620001da575050620000f7565b5f8155018590620001ca565b92508192620001c1565b634e487b7160e01b5f52602260045260245ffd5b90607f1690620000e5565b634e487b7160e01b5f52604160045260245ffd5b015190505f80620000b8565b90879350601f198316915f8052855f20925f5b878282106200027c575050841162000263575b505050811b015f55620000cb565b01515f1960f88460031b161c191690555f808062000255565b8385015186558b9790950194938401930162000242565b9091505f8052835f208580850160051c820192868610620002dc575b918991869594930160051c01915b828110620002cd575050620000a2565b5f8155859450899101620002bd565b92508192620002af565b93607f16936200008b565b5f80fd5b6040519190601f01601f191682016001600160401b038111838210176200020f57604052565b919080601f84011215620002f15782516001600160401b0381116200020f5760209062000351601f8201601f19168301620002f5565b92818452828287010111620002f1575f5b818110620003775750825f9394955001015290565b85810183015184820184015282016200036256fe6080604090808252600480361015610015575f80fd5b5f3560e01c91826306fdde03146105a357508163095ea7b3146104ce57816318160ddd146104b057816323b872dd146103ef578163313ce567146103d157816340c10f191461030257816342966c68146102e357816370a08231146102ae57816379cc67901461024257816395d89b411461012057508063a9059cbb146100f05763dd62ed3e146100a4575f80fd5b346100ec57806003193601126100ec576020906100bf6106bf565b6100c76106d5565b9060018060a01b038091165f5260058452825f2091165f528252805f20549051908152f35b5f80fd5b50346100ec57806003193601126100ec5760209061011961010f6106bf565b60243590336106eb565b5160018152f35b82346100ec575f3660031901126100ec578051905f60018054908160011c9060018316928315610238575b60209384841081146102255783885290811561020957506001146101b3575b505050829003601f01601f191682019267ffffffffffffffff8411838510176101a0575082918261019c925282610678565b0390f35b604190634e487b7160e01b5f525260245ffd5b60015f908152929350837fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf65b8385106101f5575050505083010184808061016a565b8054888601830152930192849082016101df565b60ff1916878501525050151560051b840101905084808061016a565b602289634e487b7160e01b5f525260245ffd5b91607f169161014b565b82346100ec57806003193601126100ec5760209061025e6106bf565b6102946024359161026f8382610856565b6001600160a01b03165f81815260058652848120338252865284902054909290610835565b905f5260058352815f20335f528352815f20555160018152f35b82346100ec5760203660031901126100ec576020916001600160a01b036102d36106bf565b165f528252805f20549051908152f35b82346100ec5760203660031901126100ec576101196020923533610856565b82346100ec57806003193601126100ec5761031b6106bf565b60243590811561038e5760035482810180911161037b576003556001600160a01b03165f81815260208590528390205491820191821061036857926020935f528352815f20555160018152f35b601184634e487b7160e01b5f525260245ffd5b601185634e487b7160e01b5f525260245ffd5b825162461bcd60e51b8152602081860152601e60248201527f6d696e7420616d6f756e74206e6f742067726561746572207468616e203000006044820152606490fd5b82346100ec575f3660031901126100ec576020906002549051908152f35b82346100ec5760603660031901126100ec576104096106bf565b6104116106d5565b6044359060018060a01b03831692835f5260209560058752855f20335f528752855f20548411610463575082610294939261044b926106eb565b825f5260058552835f20335f528552835f2054610835565b855162461bcd60e51b8152908101879052602160248201527f7472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636044820152606560f81b6064820152608490fd5b82346100ec575f3660031901126100ec576020906003549051908152f35b82346100ec57806003193601126100ec576104e76106bf565b3315610560576001600160a01b0316801561051d5760209250335f5260058352815f20905f528252602435815f20555160018152f35b815162461bcd60e51b8152602081850152601b60248201527f617070726f766520746f20746865207a65726f206164647265737300000000006044820152606490fd5b815162461bcd60e51b8152602081850152601d60248201527f617070726f76652066726f6d20746865207a65726f20616464726573730000006044820152606490fd5b83346100ec575f3660031901126100ec575f805460018160011c906001831692831561066e575b602093848410811461022557838852908115610209575060011461061a57505050829003601f01601f191682019267ffffffffffffffff8411838510176101a0575082918261019c925282610678565b5f808052929350837f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5635b83851061065a575050505083010184808061016a565b805488860183015293019284908201610644565b91607f16916105ca565b602080825282518183018190529093925f5b8281106106ab57505060409293505f838284010152601f8019910116010190565b81810186015184820160400152850161068a565b600435906001600160a01b03821682036100ec57565b602435906001600160a01b03821682036100ec57565b6001600160a01b039081169182156107f0571680156107ab57815f5260049081602052604092835f2054851161076857805f528260205261072f85855f2054610835565b905f5282602052835f2055805f52825f2054938401809411610755575f526020525f2055565b601182634e487b7160e01b5f525260245ffd5b835162461bcd60e51b8152602081850152601f60248201527f7472616e7366657220616d6f756e7420657863656564732062616c616e6365006044820152606490fd5b60405162461bcd60e51b815260206004820152601c60248201527f7472616e7366657220746f20746865207a65726f2061646472657373000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601e60248201527f7472616e736665722066726f6d20746865207a65726f206164647265737300006044820152606490fd5b9190820391821161084257565b634e487b7160e01b5f52601160045260245ffd5b6001600160a01b03165f81815260046020526040902054909180156108e8578181116108a3576108949161088c82600354610835565b600355610835565b905f52600460205260405f2055565b60405162461bcd60e51b815260206004820152601b60248201527f6275726e20616d6f756e7420657863656564732062616c616e636500000000006044820152606490fd5b60405162461bcd60e51b815260206004820152601e60248201527f6275726e20616d6f756e74206e6f742067726561746572207468616e203000006044820152606490fdfea2646970667358221220980af8c27004fa8ad3ccc9941ed1a706bdf7fec48ed6aa92c357e934abfeb2ec64736f6c63430008160033",
}

// TestnetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TestnetTokenMetaData.ABI instead.
var TestnetTokenABI = TestnetTokenMetaData.ABI

// TestnetTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestnetTokenMetaData.Bin instead.
var TestnetTokenBin = TestnetTokenMetaData.Bin

// DeployTestnetToken deploys a new platon contract, binding an instance of TestnetToken to it.
func DeployTestnetToken(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string) (common.Address, *types.Transaction, *TestnetToken, error) {
	parsed, err := TestnetTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestnetTokenBin), backend, _name, _symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TestnetToken{TestnetTokenCaller: TestnetTokenCaller{contract: contract}, TestnetTokenTransactor: TestnetTokenTransactor{contract: contract}, TestnetTokenFilterer: TestnetTokenFilterer{contract: contract}}, nil
}

// TestnetToken is an auto generated Go binding around an platon contract.
type TestnetToken struct {
	TestnetTokenCaller     // Read-only binding to the contract
	TestnetTokenTransactor // Write-only binding to the contract
	TestnetTokenFilterer   // Log filterer for contract events
}

// TestnetTokenCaller is an auto generated read-only Go binding around an platon contract.
type TestnetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnetTokenTransactor is an auto generated write-only Go binding around an platon contract.
type TestnetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnetTokenFilterer is an auto generated log filtering Go binding around an platon contract events.
type TestnetTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestnetTokenSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type TestnetTokenSession struct {
	Contract     *TestnetToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestnetTokenCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type TestnetTokenCallerSession struct {
	Contract *TestnetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TestnetTokenTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type TestnetTokenTransactorSession struct {
	Contract     *TestnetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TestnetTokenRaw is an auto generated low-level Go binding around an platon contract.
type TestnetTokenRaw struct {
	Contract *TestnetToken // Generic contract binding to access the raw methods on
}

// TestnetTokenCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type TestnetTokenCallerRaw struct {
	Contract *TestnetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TestnetTokenTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type TestnetTokenTransactorRaw struct {
	Contract *TestnetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestnetToken creates a new instance of TestnetToken, bound to a specific deployed contract.
func NewTestnetToken(address common.Address, backend bind.ContractBackend) (*TestnetToken, error) {
	contract, err := bindTestnetToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestnetToken{TestnetTokenCaller: TestnetTokenCaller{contract: contract}, TestnetTokenTransactor: TestnetTokenTransactor{contract: contract}, TestnetTokenFilterer: TestnetTokenFilterer{contract: contract}}, nil
}

// NewTestnetTokenCaller creates a new read-only instance of TestnetToken, bound to a specific deployed contract.
func NewTestnetTokenCaller(address common.Address, caller bind.ContractCaller) (*TestnetTokenCaller, error) {
	contract, err := bindTestnetToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestnetTokenCaller{contract: contract}, nil
}

// NewTestnetTokenTransactor creates a new write-only instance of TestnetToken, bound to a specific deployed contract.
func NewTestnetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TestnetTokenTransactor, error) {
	contract, err := bindTestnetToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestnetTokenTransactor{contract: contract}, nil
}

// NewTestnetTokenFilterer creates a new log filterer instance of TestnetToken, bound to a specific deployed contract.
func NewTestnetTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TestnetTokenFilterer, error) {
	contract, err := bindTestnetToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestnetTokenFilterer{contract: contract}, nil
}

// bindTestnetToken binds a generic wrapper to an already deployed contract.
func bindTestnetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestnetTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestnetToken *TestnetTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestnetToken.Contract.TestnetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestnetToken *TestnetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestnetToken.Contract.TestnetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestnetToken *TestnetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestnetToken.Contract.TestnetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestnetToken *TestnetTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestnetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestnetToken *TestnetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TestnetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestnetToken *TestnetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TestnetToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestnetToken *TestnetTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestnetToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestnetToken *TestnetTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestnetToken.Contract.Allowance(&_TestnetToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TestnetToken *TestnetTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TestnetToken.Contract.Allowance(&_TestnetToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestnetToken *TestnetTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TestnetToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestnetToken *TestnetTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestnetToken.Contract.BalanceOf(&_TestnetToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TestnetToken *TestnetTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TestnetToken.Contract.BalanceOf(&_TestnetToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_TestnetToken *TestnetTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestnetToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_TestnetToken *TestnetTokenSession) Decimals() (*big.Int, error) {
	return _TestnetToken.Contract.Decimals(&_TestnetToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_TestnetToken *TestnetTokenCallerSession) Decimals() (*big.Int, error) {
	return _TestnetToken.Contract.Decimals(&_TestnetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestnetToken *TestnetTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestnetToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestnetToken *TestnetTokenSession) Name() (string, error) {
	return _TestnetToken.Contract.Name(&_TestnetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TestnetToken *TestnetTokenCallerSession) Name() (string, error) {
	return _TestnetToken.Contract.Name(&_TestnetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestnetToken *TestnetTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TestnetToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestnetToken *TestnetTokenSession) Symbol() (string, error) {
	return _TestnetToken.Contract.Symbol(&_TestnetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TestnetToken *TestnetTokenCallerSession) Symbol() (string, error) {
	return _TestnetToken.Contract.Symbol(&_TestnetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestnetToken *TestnetTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TestnetToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestnetToken *TestnetTokenSession) TotalSupply() (*big.Int, error) {
	return _TestnetToken.Contract.TotalSupply(&_TestnetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TestnetToken *TestnetTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TestnetToken.Contract.TotalSupply(&_TestnetToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Approve(&_TestnetToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Approve(&_TestnetToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Burn(&_TestnetToken.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Burn(&_TestnetToken.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.BurnFrom(&_TestnetToken.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.BurnFrom(&_TestnetToken.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Mint(&_TestnetToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Mint(&_TestnetToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Transfer(&_TestnetToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.Transfer(&_TestnetToken.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.TransferFrom(&_TestnetToken.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.TransferFrom(&_TestnetToken.TransactOpts, from, to, amount)
}
