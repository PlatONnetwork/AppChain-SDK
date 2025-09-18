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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"burnFrom\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234620003055762000ebe803803806200001d8162000309565b928339810190604081830312620003055780516001600160401b03908181116200030557836200004f9184016200032f565b916020938482015183811162000305576200006b92016200032f565b825182811162000210576003918254916001958684811c94168015620002fa575b88851014620002e6578190601f9485811162000293575b50889085831160011462000230575f9262000224575b50505f1982861b1c191690861b1783555b8051938411620002105760049586548681811c9116801562000205575b82821014620001f257838111620001aa575b50809285116001146200014057509383949184925f9562000134575b50501b925f19911b1c19161790555b604051610b1e9081620003a08239f35b015193505f8062000115565b92919084601f198116885f52855f20955f905b898383106200018f575050501062000175575b50505050811b01905562000124565b01519060f8845f19921b161c191690555f80808062000166565b85870151895590970196948501948893509081019062000153565b875f52815f208480880160051c820192848910620001e8575b0160051c019087905b828110620001dc575050620000f9565b5f8155018790620001cc565b92508192620001c3565b602288634e487b7160e01b5f525260245ffd5b90607f1690620000e7565b634e487b7160e01b5f52604160045260245ffd5b015190505f80620000b9565b90889350601f19831691875f528a5f20925f5b8c8282106200027c575050841162000264575b505050811b018355620000ca565b01515f1983881b60f8161c191690555f808062000256565b8385015186558c9790950194938401930162000243565b909150855f52885f208580850160051c8201928b8610620002dc575b918a91869594930160051c01915b828110620002cd575050620000a3565b5f81558594508a9101620002bd565b92508192620002af565b634e487b7160e01b5f52602260045260245ffd5b93607f16936200008c565b5f80fd5b6040519190601f01601f191682016001600160401b038111838210176200021057604052565b919080601f84011215620003055782516001600160401b038111620002105760209062000365601f8201601f1916830162000309565b9281845282828701011162000305575f5b8181106200038b5750825f9394955001015290565b85810183015184820184015282016200037656fe6080604090808252600480361015610015575f80fd5b5f3560e01c91826306fdde031461057857508163095ea7b31461054f57816318160ddd1461053157816323b872dd146104f5578163313ce567146104da578163395093511461048e57816340c10f191461038257816342966c681461036357816370a082311461032d57816379cc6790146102fa57816395d89b41146101db578163a457c2d71461013657508063a9059cbb146101065763dd62ed3e146100ba575f80fd5b346101025780600319360112610102576020906100d5610697565b6100dd6106ad565b9060018060a01b038091165f5260018452825f2091165f528252805f20549051908152f35b5f80fd5b503461010257806003193601126101025760209061012f610125610697565b6024359033610875565b5160018152f35b905034610102578160031936011261010257610150610697565b9060243590335f526001602052835f2060018060a01b0384165f52602052835f20549082821061018a5760208561012f85850387336106e4565b608490602086519162461bcd60e51b8352820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b6064820152fd5b8234610102575f366003190112610102578051905f835460018160011c90600183169283156102f0575b60209384841081146102dd578388529081156102c1575060011461026d575b505050829003601f01601f191682019267ffffffffffffffff84118385101761025a5750829182610256925282610650565b0390f35b604190634e487b7160e01b5f525260245ffd5b5f878152929350837f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5b8385106102ad5750505050830101848080610224565b805488860183015293019284908201610297565b60ff1916878501525050151560051b8401019050848080610224565b602289634e487b7160e01b5f525260245ffd5b91607f1691610205565b823461010257806003193601126101025760209061012f610319610697565b602435906103288233836107e2565b6109e2565b8234610102576020366003190112610102576020906001600160a01b03610352610697565b165f525f8252805f20549051908152f35b82346101025760203660031901126101025761012f60209235336109e2565b823461010257806003193601126101025761039b610697565b9160243592831561044b576001600160a01b031690811561040857505f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6020856103e982976002546106c3565b6002558484528382528584208181540190558551908152a35160018152f35b606490602084519162461bcd60e51b8352820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152fd5b506020606492519162461bcd60e51b8352820152601e60248201527f6d696e7420616d6f756e74206e6f742067726561746572207468616e203000006044820152fd5b823461010257806003193601126101025760209061012f6104ad610697565b335f5260018452825f2060018060a01b0382165f5284526104d3602435845f20546106c3565b90336106e4565b8234610102575f366003190112610102576020905160128152f35b82346101025760603660031901126101025760209061012f610515610697565b61051d6106ad565b6044359161052c8333836107e2565b610875565b8234610102575f366003190112610102576020906002549051908152f35b823461010257806003193601126101025760209061012f61056e610697565b60243590336106e4565b8334610102575f366003190112610102575f60035460018160011c9060018316928315610646575b60209384841081146102dd578388529081156102c157506001146105f057505050829003601f01601f191682019267ffffffffffffffff84118385101761025a5750829182610256925282610650565b60035f908152929350837fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5b8385106106325750505050830101848080610224565b80548886018301529301928490820161061c565b91607f16916105a0565b602080825282518183018190529093925f5b82811061068357505060409293505f838284010152601f8019910116010190565b818101860151848201604001528501610662565b600435906001600160a01b038216820361010257565b602435906001600160a01b038216820361010257565b919082018092116106d057565b634e487b7160e01b5f52601160045260245ffd5b6001600160a01b0390811691821561079157169182156107415760207f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591835f526001825260405f20855f5282528060405f2055604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608490fd5b60405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608490fd5b9060018060a01b038083165f52600160205260405f209082165f5260205260405f2054925f198403610815575b50505050565b808410610830576108279303916106e4565b5f80808061080f565b60405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606490fd5b6001600160a01b0390811691821561098f571691821561093e57815f525f60205260405f20548181106108ea57817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92602092855f525f84520360405f2055845f5260405f20818154019055604051908152a3565b60405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608490fd5b60405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608490fd5b6001600160a01b03168015610a9957805f525f60205260405f205491808310610a49576020817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef925f958587528684520360408620558060025403600255604051908152a3565b60405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608490fd5b60405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608490fdfea26469706673582212208f6aa6b6d7787035aa038358627405ab49ae9816ff88b9946cf83e17ad1a1ce364736f6c63430008160033",
}

// TestnetTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TestnetTokenMetaData.ABI instead.
var TestnetTokenABI = TestnetTokenMetaData.ABI

// TestnetTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestnetTokenMetaData.Bin instead.
var TestnetTokenBin = TestnetTokenMetaData.Bin

// DeployTestnetToken deploys a new platon contract, binding an instance of TestnetToken to it.
func DeployTestnetToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *TestnetToken, error) {
	parsed, err := TestnetTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestnetTokenBin), backend, name_, symbol_)
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
// Solidity: function decimals() view returns(uint8)
func (_TestnetToken *TestnetTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TestnetToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestnetToken *TestnetTokenSession) Decimals() (uint8, error) {
	return _TestnetToken.Contract.Decimals(&_TestnetToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TestnetToken *TestnetTokenCallerSession) Decimals() (uint8, error) {
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

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestnetToken *TestnetTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.DecreaseAllowance(&_TestnetToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.DecreaseAllowance(&_TestnetToken.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestnetToken *TestnetTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestnetToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestnetToken *TestnetTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.IncreaseAllowance(&_TestnetToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TestnetToken *TestnetTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TestnetToken.Contract.IncreaseAllowance(&_TestnetToken.TransactOpts, spender, addedValue)
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

// TestnetTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TestnetToken contract.
type TestnetTokenApprovalIterator struct {
	Event *TestnetTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestnetTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestnetTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestnetTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestnetTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestnetTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestnetTokenApproval represents a Approval event raised by the TestnetToken contract.
type TestnetTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestnetToken *TestnetTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TestnetTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestnetToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TestnetTokenApprovalIterator{contract: _TestnetToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestnetToken *TestnetTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TestnetTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TestnetToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestnetTokenApproval)
				if err := _TestnetToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TestnetToken *TestnetTokenFilterer) ParseApproval(log types.Log) (*TestnetTokenApproval, error) {
	event := new(TestnetTokenApproval)
	if err := _TestnetToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TestnetTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TestnetToken contract.
type TestnetTokenTransferIterator struct {
	Event *TestnetTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log      // Log channel receiving the found contract events
	sub  platon.Subscription // Subscription for errors, completion and termination
	done bool                // Whether the subscription completed delivering logs
	fail error               // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TestnetTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TestnetTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TestnetTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TestnetTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TestnetTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TestnetTokenTransfer represents a Transfer event raised by the TestnetToken contract.
type TestnetTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestnetToken *TestnetTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TestnetTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestnetToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TestnetTokenTransferIterator{contract: _TestnetToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestnetToken *TestnetTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TestnetTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TestnetToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TestnetTokenTransfer)
				if err := _TestnetToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TestnetToken *TestnetTokenFilterer) ParseTransfer(log types.Log) (*TestnetTokenTransfer, error) {
	event := new(TestnetTokenTransfer)
	if err := _TestnetToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
