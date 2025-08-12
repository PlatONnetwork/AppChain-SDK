// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawhandler

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

// WithdrawHandlerMetaData contains all meta data concerning the WithdrawHandler contract.
var WithdrawHandlerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEPOSIT_SIG\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAW_SIG\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newRegistryManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newStateSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newChildWithdrawManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newExitHelper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onL2StateReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintableTokenDeposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MintableTokenWithdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608080604052346100bf575f549060ff8260081c1661006d575060ff80821603610033575b604051610a4090816100c48239f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f610024565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fdfe6080604081815260049081361015610015575f80fd5b5f925f3560e01c9081630c63109e1461093857508063b1768065146108fe578063d41f1771146108c4578063f3fef3a3146105c3578063f43cda8b146102105763f8c8765e14610063575f80fd5b3461020c57608036600319011261020c5761007c61095d565b90610085610973565b6001600160a01b0393906044358581169190829003610208576064359186831680930361020857875460ff8160081c1615948580966101fb575b80156101e4575b1561018a575060ff198116600117895584610179575b5087549562010000600160b01b039060101b16968762010000600160b01b031988161789556bffffffffffffffffffffffff60a01b92168260015416176001558160025416176002556003541617600355610135578380f35b610100600160b01b0319909116909117825551600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a15f80808380f35b61ffff19166101011788555f6100dc565b608490602088519162461bcd60e51b8352820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152fd5b50303b1580156100c65750600160ff8316146100c6565b50600160ff8316106100bf565b5f80fd5b8280fd5b508290346105bf5760603660031901126105bf5761022c610973565b604493843567ffffffffffffffff938482116105bb57366023830112156105bb57818301359485116105bb576024928583019736858a01116104c6576003546001600160a01b039690871633036105705786806002541691160361051b5760209687116104c657838501357f87a7811f4bfedea3d341ad165680ae306b01aaeacc205d227629cf157dd9f821036104ca57608084899a03126104c6576102d3818501610989565b948660846102e360648801610989565b960135951697878a5460101c168551637c8211ff60e11b8152308682015282818581855afa9081156104bc578c9161048a575b5082908488518094819363ea78803f60e01b83528a8301525afa90811561048057878c8b61037d948e9f94879591610453575b50169c8d918a518096819582946340c10f1960e01b84528d840160209093929193604081019460018060a01b031681520152565b03925af1908115610449578c9161041c575b50156103dd5750509151949093166001600160a01b03168452506020830152507f198a0b99a541e7fd58217d300068f0b344ea513f245e95ca0a1883e7a3b137b19080604081015b0390a380f35b7f576974686472617748616e646c65723a204d494e545f4641494c4544000000009291601c91606496519562461bcd60e51b8752860152840152820152fd5b61043c9150823d8411610442575b610434818361099d565b8101906109f2565b8c61038f565b503d61042a565b86513d8e823e3d90fd5b6104739150853d8711610479575b61046b818361099d565b8101906109d3565b5f610349565b503d610461565b86513d8d823e3d90fd5b809c50838092503d83116104b5575b6104a3818361099d565b8101031261020857818c9b5190610316565b503d610499565b87513d8e823e3d90fd5b8780fd5b7f576974686472617748616e646c65723a20494e56414c49445f4d4554484f445f9060298689608496519562461bcd60e51b8752860152840152820152685349474e415455524560b81b6064820152fd5b7f576974686472617748616e646c65723a204f4e4c595f4348494c445f5749544890602c866020608496519562461bcd60e51b87528601528401528201526b222920abafa6a0a720a3a2a960a11b6064820152fd5b507f576974686472617748616e646c65723a204f4e4c595f455849545f48454c5045906021866020608496519562461bcd60e51b8752860152840152820152602960f91b6064820152fd5b8580fd5b5080fd5b50346102085780600319360112610208576105dc61095d565b9160249182359160018060a01b0393845f5460101c1691835192637c8211ff60e11b8452308285015260209384818581855afa90811561086e579085915f91610895575b508487518094819363ea78803f60e01b8352878301525afa9081156107dc579087915f91610878575b50855163079cc67960e41b8152338482019081526020810189905292909116979185908290819060400103815f8c5af190811561086e575f91610851575b50156108105780600154168160025416918651997f7a8dc26796a1e50e6e190b70259f58f6a4edd5b22280ceecc82b687b8e982869878c015233888c015216988960608201528760808201526080815260a081019567ffffffffffffffff87119382881085176107fe57878952833b15610208576316f1983160e01b885260a483015260c48201889052815160e483018190528792915f5b8281106107e65750505f60648383836101048896829a98010152601f801991011681010301925af180156107dc57610790575b505091513381526020810193909352507f85f88af14d28fa465d5461b53de077269e7c6833da437c2647e7b0d788730a8091905080604081016103d7565b9091929397506107cb57505084525f93816103d77f85f88af14d28fa465d5461b53de077269e7c6833da437c2647e7b0d788730a8087610752565b604190634e487b7160e01b5f52525ffd5b85513d5f823e3d90fd5b8084018083015161010490910152899450810161071f565b86604187634e487b7160e01b5f52525ffd5b50601c6064928486519362461bcd60e51b85528401528201527f576974686472617748616e646c65723a204255524e5f4641494c4544000000006044820152fd5b6108689150853d871161044257610434818361099d565b5f610687565b86513d5f823e3d90fd5b61088f9150853d87116104795761046b818361099d565b5f610649565b82819392503d83116108bd575b6108ac818361099d565b81010312610208578490515f610620565b503d6108a2565b5034610208575f36600319011261020857602090517f87a7811f4bfedea3d341ad165680ae306b01aaeacc205d227629cf157dd9f8218152f35b5034610208575f36600319011261020857602090517f7a8dc26796a1e50e6e190b70259f58f6a4edd5b22280ceecc82b687b8e9828698152f35b34610208575f366003190112610208575f5460101c6001600160a01b03168152602090f35b600435906001600160a01b038216820361020857565b602435906001600160a01b038216820361020857565b35906001600160a01b038216820361020857565b90601f8019910116810190811067ffffffffffffffff8211176109bf57604052565b634e487b7160e01b5f52604160045260245ffd5b9081602091031261020857516001600160a01b03811681036102085790565b9081602091031261020857518015158103610208579056fea26469706673582212209ca11d629d115e0dc41eaf707063239432fcec1baf8a1a13c9514c0d0365c88264736f6c63430008160033",
}

// WithdrawHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawHandlerMetaData.ABI instead.
var WithdrawHandlerABI = WithdrawHandlerMetaData.ABI

// WithdrawHandlerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use WithdrawHandlerMetaData.Bin instead.
var WithdrawHandlerBin = WithdrawHandlerMetaData.Bin

// DeployWithdrawHandler deploys a new platon contract, binding an instance of WithdrawHandler to it.
func DeployWithdrawHandler(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WithdrawHandler, error) {
	parsed, err := WithdrawHandlerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(WithdrawHandlerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WithdrawHandler{WithdrawHandlerCaller: WithdrawHandlerCaller{contract: contract}, WithdrawHandlerTransactor: WithdrawHandlerTransactor{contract: contract}, WithdrawHandlerFilterer: WithdrawHandlerFilterer{contract: contract}}, nil
}

// WithdrawHandler is an auto generated Go binding around an platon contract.
type WithdrawHandler struct {
	WithdrawHandlerCaller     // Read-only binding to the contract
	WithdrawHandlerTransactor // Write-only binding to the contract
	WithdrawHandlerFilterer   // Log filterer for contract events
}

// WithdrawHandlerCaller is an auto generated read-only Go binding around an platon contract.
type WithdrawHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawHandlerTransactor is an auto generated write-only Go binding around an platon contract.
type WithdrawHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawHandlerFilterer is an auto generated log filtering Go binding around an platon contract events.
type WithdrawHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawHandlerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type WithdrawHandlerSession struct {
	Contract     *WithdrawHandler  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawHandlerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type WithdrawHandlerCallerSession struct {
	Contract *WithdrawHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// WithdrawHandlerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type WithdrawHandlerTransactorSession struct {
	Contract     *WithdrawHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// WithdrawHandlerRaw is an auto generated low-level Go binding around an platon contract.
type WithdrawHandlerRaw struct {
	Contract *WithdrawHandler // Generic contract binding to access the raw methods on
}

// WithdrawHandlerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type WithdrawHandlerCallerRaw struct {
	Contract *WithdrawHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawHandlerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type WithdrawHandlerTransactorRaw struct {
	Contract *WithdrawHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawHandler creates a new instance of WithdrawHandler, bound to a specific deployed contract.
func NewWithdrawHandler(address common.Address, backend bind.ContractBackend) (*WithdrawHandler, error) {
	contract, err := bindWithdrawHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WithdrawHandler{WithdrawHandlerCaller: WithdrawHandlerCaller{contract: contract}, WithdrawHandlerTransactor: WithdrawHandlerTransactor{contract: contract}, WithdrawHandlerFilterer: WithdrawHandlerFilterer{contract: contract}}, nil
}

// NewWithdrawHandlerCaller creates a new read-only instance of WithdrawHandler, bound to a specific deployed contract.
func NewWithdrawHandlerCaller(address common.Address, caller bind.ContractCaller) (*WithdrawHandlerCaller, error) {
	contract, err := bindWithdrawHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawHandlerCaller{contract: contract}, nil
}

// NewWithdrawHandlerTransactor creates a new write-only instance of WithdrawHandler, bound to a specific deployed contract.
func NewWithdrawHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawHandlerTransactor, error) {
	contract, err := bindWithdrawHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawHandlerTransactor{contract: contract}, nil
}

// NewWithdrawHandlerFilterer creates a new log filterer instance of WithdrawHandler, bound to a specific deployed contract.
func NewWithdrawHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawHandlerFilterer, error) {
	contract, err := bindWithdrawHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawHandlerFilterer{contract: contract}, nil
}

// bindWithdrawHandler binds a generic wrapper to an already deployed contract.
func bindWithdrawHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawHandler *WithdrawHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawHandler.Contract.WithdrawHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawHandler *WithdrawHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.WithdrawHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawHandler *WithdrawHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.WithdrawHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WithdrawHandler *WithdrawHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WithdrawHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WithdrawHandler *WithdrawHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WithdrawHandler *WithdrawHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_WithdrawHandler *WithdrawHandlerCaller) DEPOSITSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WithdrawHandler.contract.Call(opts, &out, "DEPOSIT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_WithdrawHandler *WithdrawHandlerSession) DEPOSITSIG() ([32]byte, error) {
	return _WithdrawHandler.Contract.DEPOSITSIG(&_WithdrawHandler.CallOpts)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_WithdrawHandler *WithdrawHandlerCallerSession) DEPOSITSIG() ([32]byte, error) {
	return _WithdrawHandler.Contract.DEPOSITSIG(&_WithdrawHandler.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_WithdrawHandler *WithdrawHandlerCaller) WITHDRAWSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WithdrawHandler.contract.Call(opts, &out, "WITHDRAW_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_WithdrawHandler *WithdrawHandlerSession) WITHDRAWSIG() ([32]byte, error) {
	return _WithdrawHandler.Contract.WITHDRAWSIG(&_WithdrawHandler.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_WithdrawHandler *WithdrawHandlerCallerSession) WITHDRAWSIG() ([32]byte, error) {
	return _WithdrawHandler.Contract.WITHDRAWSIG(&_WithdrawHandler.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_WithdrawHandler *WithdrawHandlerCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WithdrawHandler.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_WithdrawHandler *WithdrawHandlerSession) RegistryManager() (common.Address, error) {
	return _WithdrawHandler.Contract.RegistryManager(&_WithdrawHandler.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_WithdrawHandler *WithdrawHandlerCallerSession) RegistryManager() (common.Address, error) {
	return _WithdrawHandler.Contract.RegistryManager(&_WithdrawHandler.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildWithdrawManager, address newExitHelper) returns()
func (_WithdrawHandler *WithdrawHandlerTransactor) Initialize(opts *bind.TransactOpts, newRegistryManager common.Address, newStateSender common.Address, newChildWithdrawManager common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _WithdrawHandler.contract.Transact(opts, "initialize", newRegistryManager, newStateSender, newChildWithdrawManager, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildWithdrawManager, address newExitHelper) returns()
func (_WithdrawHandler *WithdrawHandlerSession) Initialize(newRegistryManager common.Address, newStateSender common.Address, newChildWithdrawManager common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.Initialize(&_WithdrawHandler.TransactOpts, newRegistryManager, newStateSender, newChildWithdrawManager, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildWithdrawManager, address newExitHelper) returns()
func (_WithdrawHandler *WithdrawHandlerTransactorSession) Initialize(newRegistryManager common.Address, newStateSender common.Address, newChildWithdrawManager common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.Initialize(&_WithdrawHandler.TransactOpts, newRegistryManager, newStateSender, newChildWithdrawManager, newExitHelper)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_WithdrawHandler *WithdrawHandlerTransactor) OnL2StateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawHandler.contract.Transact(opts, "onL2StateReceive", arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_WithdrawHandler *WithdrawHandlerSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.OnL2StateReceive(&_WithdrawHandler.TransactOpts, arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_WithdrawHandler *WithdrawHandlerTransactorSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.OnL2StateReceive(&_WithdrawHandler.TransactOpts, arg0, sender, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_WithdrawHandler *WithdrawHandlerTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WithdrawHandler.contract.Transact(opts, "withdraw", recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_WithdrawHandler *WithdrawHandlerSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.Withdraw(&_WithdrawHandler.TransactOpts, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address recipient, uint256 amount) returns()
func (_WithdrawHandler *WithdrawHandlerTransactorSession) Withdraw(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WithdrawHandler.Contract.Withdraw(&_WithdrawHandler.TransactOpts, recipient, amount)
}

// WithdrawHandlerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the WithdrawHandler contract.
type WithdrawHandlerInitializedIterator struct {
	Event *WithdrawHandlerInitialized // Event containing the contract specifics and raw log

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
func (it *WithdrawHandlerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawHandlerInitialized)
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
		it.Event = new(WithdrawHandlerInitialized)
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
func (it *WithdrawHandlerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawHandlerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawHandlerInitialized represents a Initialized event raised by the WithdrawHandler contract.
type WithdrawHandlerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawHandler *WithdrawHandlerFilterer) FilterInitialized(opts *bind.FilterOpts) (*WithdrawHandlerInitializedIterator, error) {

	logs, sub, err := _WithdrawHandler.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &WithdrawHandlerInitializedIterator{contract: _WithdrawHandler.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawHandler *WithdrawHandlerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *WithdrawHandlerInitialized) (event.Subscription, error) {

	logs, sub, err := _WithdrawHandler.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawHandlerInitialized)
				if err := _WithdrawHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_WithdrawHandler *WithdrawHandlerFilterer) ParseInitialized(log types.Log) (*WithdrawHandlerInitialized, error) {
	event := new(WithdrawHandlerInitialized)
	if err := _WithdrawHandler.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawHandlerMintableTokenDepositIterator is returned from FilterMintableTokenDeposit and is used to iterate over the raw logs and unpacked data for MintableTokenDeposit events raised by the WithdrawHandler contract.
type WithdrawHandlerMintableTokenDepositIterator struct {
	Event *WithdrawHandlerMintableTokenDeposit // Event containing the contract specifics and raw log

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
func (it *WithdrawHandlerMintableTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawHandlerMintableTokenDeposit)
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
		it.Event = new(WithdrawHandlerMintableTokenDeposit)
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
func (it *WithdrawHandlerMintableTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawHandlerMintableTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawHandlerMintableTokenDeposit represents a MintableTokenDeposit event raised by the WithdrawHandler contract.
type WithdrawHandlerMintableTokenDeposit struct {
	Token     common.Address
	Recipient common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMintableTokenDeposit is a free log retrieval operation binding the contract event 0x198a0b99a541e7fd58217d300068f0b344ea513f245e95ca0a1883e7a3b137b1.
//
// Solidity: event MintableTokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_WithdrawHandler *WithdrawHandlerFilterer) FilterMintableTokenDeposit(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*WithdrawHandlerMintableTokenDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WithdrawHandler.contract.FilterLogs(opts, "MintableTokenDeposit", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawHandlerMintableTokenDepositIterator{contract: _WithdrawHandler.contract, event: "MintableTokenDeposit", logs: logs, sub: sub}, nil
}

// WatchMintableTokenDeposit is a free log subscription operation binding the contract event 0x198a0b99a541e7fd58217d300068f0b344ea513f245e95ca0a1883e7a3b137b1.
//
// Solidity: event MintableTokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_WithdrawHandler *WithdrawHandlerFilterer) WatchMintableTokenDeposit(opts *bind.WatchOpts, sink chan<- *WithdrawHandlerMintableTokenDeposit, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WithdrawHandler.contract.WatchLogs(opts, "MintableTokenDeposit", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawHandlerMintableTokenDeposit)
				if err := _WithdrawHandler.contract.UnpackLog(event, "MintableTokenDeposit", log); err != nil {
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

// ParseMintableTokenDeposit is a log parse operation binding the contract event 0x198a0b99a541e7fd58217d300068f0b344ea513f245e95ca0a1883e7a3b137b1.
//
// Solidity: event MintableTokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_WithdrawHandler *WithdrawHandlerFilterer) ParseMintableTokenDeposit(log types.Log) (*WithdrawHandlerMintableTokenDeposit, error) {
	event := new(WithdrawHandlerMintableTokenDeposit)
	if err := _WithdrawHandler.contract.UnpackLog(event, "MintableTokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawHandlerMintableTokenWithdrawIterator is returned from FilterMintableTokenWithdraw and is used to iterate over the raw logs and unpacked data for MintableTokenWithdraw events raised by the WithdrawHandler contract.
type WithdrawHandlerMintableTokenWithdrawIterator struct {
	Event *WithdrawHandlerMintableTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *WithdrawHandlerMintableTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WithdrawHandlerMintableTokenWithdraw)
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
		it.Event = new(WithdrawHandlerMintableTokenWithdraw)
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
func (it *WithdrawHandlerMintableTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WithdrawHandlerMintableTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WithdrawHandlerMintableTokenWithdraw represents a MintableTokenWithdraw event raised by the WithdrawHandler contract.
type WithdrawHandlerMintableTokenWithdraw struct {
	Token      common.Address
	Recipient  common.Address
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMintableTokenWithdraw is a free log retrieval operation binding the contract event 0x85f88af14d28fa465d5461b53de077269e7c6833da437c2647e7b0d788730a80.
//
// Solidity: event MintableTokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_WithdrawHandler *WithdrawHandlerFilterer) FilterMintableTokenWithdraw(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*WithdrawHandlerMintableTokenWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WithdrawHandler.contract.FilterLogs(opts, "MintableTokenWithdraw", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &WithdrawHandlerMintableTokenWithdrawIterator{contract: _WithdrawHandler.contract, event: "MintableTokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchMintableTokenWithdraw is a free log subscription operation binding the contract event 0x85f88af14d28fa465d5461b53de077269e7c6833da437c2647e7b0d788730a80.
//
// Solidity: event MintableTokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_WithdrawHandler *WithdrawHandlerFilterer) WatchMintableTokenWithdraw(opts *bind.WatchOpts, sink chan<- *WithdrawHandlerMintableTokenWithdraw, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _WithdrawHandler.contract.WatchLogs(opts, "MintableTokenWithdraw", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WithdrawHandlerMintableTokenWithdraw)
				if err := _WithdrawHandler.contract.UnpackLog(event, "MintableTokenWithdraw", log); err != nil {
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

// ParseMintableTokenWithdraw is a log parse operation binding the contract event 0x85f88af14d28fa465d5461b53de077269e7c6833da437c2647e7b0d788730a80.
//
// Solidity: event MintableTokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_WithdrawHandler *WithdrawHandlerFilterer) ParseMintableTokenWithdraw(log types.Log) (*WithdrawHandlerMintableTokenWithdraw, error) {
	event := new(WithdrawHandlerMintableTokenWithdraw)
	if err := _WithdrawHandler.contract.UnpackLog(event, "MintableTokenWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
