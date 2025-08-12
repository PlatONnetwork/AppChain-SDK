// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositmanager

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

// DepositManagerMetaData contains all meta data concerning the DepositManager contract.
var DepositManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEPOSIT_SIG\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"WITHDRAW_SIG\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newRegistryManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newStateSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newChildDepositHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"newExitHelper\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"onL2StateReceive\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registryManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIRegistryManager\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDeposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608080604052346100bf575f549060ff8260081c1661006d575060ff80821603610033575b604051610b1c90816100c48239f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f610024565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fdfe6080604081815260049081361015610015575f80fd5b5f925f3560e01c9081630c63109e146108095750806347e7ef24146105c1578063b176806514610586578063d41f177114610547578063f43cda8b146102105763f8c8765e14610063575f80fd5b3461020c57608036600319011261020c5761007c61082e565b90610085610844565b6001600160a01b0393906044358581169190829003610208576064359186831680930361020857875460ff8160081c1615948580966101fb575b80156101e4575b1561018a575060ff198116600117895584610179575b5087549562010000600160b01b039060101b16968762010000600160b01b031988161789556bffffffffffffffffffffffff60a01b92168260015416176001558160025416176002556003541617600355610135578380f35b610100600160b01b0319909116909117825551600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249890602090a15f80808380f35b61ffff19166101011788555f6100dc565b608490602088519162461bcd60e51b8352820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152fd5b50303b1580156100c65750600160ff8316146100c6565b50600160ff8316106100bf565b5f80fd5b8280fd5b50903461020c57606036600319011261020c5761022b610844565b67ffffffffffffffff90604435828111610543573660238201121561054357808401359283116105435760249083810136838201116104a8576003546001600160a01b03949085163303610501578480600254169116036104ac5760209485116104a857818301357f7a8dc26796a1e50e6e190b70259f58f6a4edd5b22280ceecc82b687b8e9828690361045557816080910312610451576102cf6044820161085a565b918360846102df6064850161085a565b93013592169584895460101c16908851637c8211ff60e11b8152308282015287818581865afa90811561044757889285918d93610413575b508b51948593849263ea78803f60e01b84528301525afa9089821561040857926103ae877f0cc8e26c5fe2f82346d6755e53b81d7301a600a4ada570aa7281c99f1c983d5c999a9b946103b4946103d59897916103db575b50855163a9059cbb60e01b9c81019c909c526001600160a01b038d16938c0193845260208401879052169991829060400103601f19810183528261089e565b8861091d565b51939092166001600160a01b03168352602083019190915281906040820190565b0390a380f35b6103fb91508c8d3d10610401575b6103f3818361089e565b8101906108c0565b5f61036f565b503d6103e9565b8951903d90823e3d90fd5b93849193508092503d8311610440575b61042d818361089e565b810103126102085783889251915f610317565b503d610423565b8a513d8d823e3d90fd5b8680fd5b865162461bcd60e51b81528087018690526028818501527f4465706f7369744d616e616765723a20494e56414c49445f4d4554484f445f5360448201526749474e415455524560c01b6064820152608490fd5b8780fd5b865162461bcd60e51b8152602081880152602a818501527f4465706f7369744d616e616765723a204f4e4c595f4348494c445f4445504f5360448201526924aa2fa420a7222622a960b11b6064820152608490fd5b875162461bcd60e51b81526020818901819052818601527f4465706f7369744d616e616765723a204f4e4c595f455849545f48454c5045526044820152606490fd5b8580fd5b838234610582578160031936011261058257602090517f87a7811f4bfedea3d341ad165680ae306b01aaeacc205d227629cf157dd9f8218152f35b5080fd5b838234610582578160031936011261058257602090517f7a8dc26796a1e50e6e190b70259f58f6a4edd5b22280ceecc82b687b8e9828698152f35b50346102085780600319360112610208576105da61082e565b9060249081359160018060a01b0390815f5460101c1695835196637c8211ff60e11b8852308289015260209788818581855afa9081156107ff579089915f916107d0575b508487518094819363ea78803f60e01b8352878301525afa9081156107c6579084915f916107a9575b50169561067985516323b872dd60e01b8a8201523385820152306044820152876064820152606481526103ae8161086e565b8360015416938060025416917f87a7811f4bfedea3d341ad165680ae306b01aaeacc205d227629cf157dd9f82187519a8b015233878b01521697886060820152866080820152608081526106cc8161086e565b843b15610208576106ff945f928388518098819582946316f1983160e01b8452898401528b8a84015260448301906108df565b03925af1801561079f5761074a575b505090513381526020810192909252507fedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a059080604081016103d5565b909192965067ffffffffffffffff831161078e57505084525f93816103d57fedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a0561070e565b604190634e487b7160e01b5f52525ffd5b84513d5f823e3d90fd5b6107c09150893d8b11610401576103f3818361089e565b5f610647565b85513d5f823e3d90fd5b82819392503d83116107f8575b6107e7818361089e565b81010312610208578890515f61061e565b503d6107dd565b86513d5f823e3d90fd5b34610208575f366003190112610208575f5460101c6001600160a01b03168152602090f35b600435906001600160a01b038216820361020857565b602435906001600160a01b038216820361020857565b35906001600160a01b038216820361020857565b60a0810190811067ffffffffffffffff82111761088a57604052565b634e487b7160e01b5f52604160045260245ffd5b90601f8019910116810190811067ffffffffffffffff82111761088a57604052565b9081602091031261020857516001600160a01b03811681036102085790565b91908251928382525f5b848110610909575050825f602080949584010152601f8019910116010190565b6020818301810151848301820152016108e9565b60018060a01b03166040516040810167ffffffffffffffff908281108282111761088a576040525f806020958685527f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c656487860152868151910182875af13d15610a40573d91821161088a576109b093604051926109a387601f19601f840116018561089e565b83523d5f8785013e610a49565b8051828115918215610a20575b50509050156109c95750565b6084906040519062461bcd60e51b82526004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152fd5b83809293500103126102085781015180151581036102085780825f6109bd565b6109b093606092505b91929015610aab5750815115610a5d575090565b3b15610a665790565b60405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606490fd5b825190915015610abe5750805190602001fd5b60405162461bcd60e51b815260206004820152908190610ae29060248301906108df565b0390fdfea26469706673582212208a7f5242605184d19ee9970902b74c8a8ba69a79a0493026c43d31bb4a6a422f64736f6c63430008160033",
}

// DepositManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositManagerMetaData.ABI instead.
var DepositManagerABI = DepositManagerMetaData.ABI

// DepositManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DepositManagerMetaData.Bin instead.
var DepositManagerBin = DepositManagerMetaData.Bin

// DeployDepositManager deploys a new platon contract, binding an instance of DepositManager to it.
func DeployDepositManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DepositManager, error) {
	parsed, err := DepositManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DepositManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DepositManager{DepositManagerCaller: DepositManagerCaller{contract: contract}, DepositManagerTransactor: DepositManagerTransactor{contract: contract}, DepositManagerFilterer: DepositManagerFilterer{contract: contract}}, nil
}

// DepositManager is an auto generated Go binding around an platon contract.
type DepositManager struct {
	DepositManagerCaller     // Read-only binding to the contract
	DepositManagerTransactor // Write-only binding to the contract
	DepositManagerFilterer   // Log filterer for contract events
}

// DepositManagerCaller is an auto generated read-only Go binding around an platon contract.
type DepositManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerTransactor is an auto generated write-only Go binding around an platon contract.
type DepositManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerFilterer is an auto generated log filtering Go binding around an platon contract events.
type DepositManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositManagerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type DepositManagerSession struct {
	Contract     *DepositManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositManagerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type DepositManagerCallerSession struct {
	Contract *DepositManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DepositManagerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type DepositManagerTransactorSession struct {
	Contract     *DepositManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DepositManagerRaw is an auto generated low-level Go binding around an platon contract.
type DepositManagerRaw struct {
	Contract *DepositManager // Generic contract binding to access the raw methods on
}

// DepositManagerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type DepositManagerCallerRaw struct {
	Contract *DepositManagerCaller // Generic read-only contract binding to access the raw methods on
}

// DepositManagerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type DepositManagerTransactorRaw struct {
	Contract *DepositManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositManager creates a new instance of DepositManager, bound to a specific deployed contract.
func NewDepositManager(address common.Address, backend bind.ContractBackend) (*DepositManager, error) {
	contract, err := bindDepositManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DepositManager{DepositManagerCaller: DepositManagerCaller{contract: contract}, DepositManagerTransactor: DepositManagerTransactor{contract: contract}, DepositManagerFilterer: DepositManagerFilterer{contract: contract}}, nil
}

// NewDepositManagerCaller creates a new read-only instance of DepositManager, bound to a specific deployed contract.
func NewDepositManagerCaller(address common.Address, caller bind.ContractCaller) (*DepositManagerCaller, error) {
	contract, err := bindDepositManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerCaller{contract: contract}, nil
}

// NewDepositManagerTransactor creates a new write-only instance of DepositManager, bound to a specific deployed contract.
func NewDepositManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositManagerTransactor, error) {
	contract, err := bindDepositManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositManagerTransactor{contract: contract}, nil
}

// NewDepositManagerFilterer creates a new log filterer instance of DepositManager, bound to a specific deployed contract.
func NewDepositManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositManagerFilterer, error) {
	contract, err := bindDepositManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositManagerFilterer{contract: contract}, nil
}

// bindDepositManager binds a generic wrapper to an already deployed contract.
func bindDepositManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DepositManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManager *DepositManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositManager.Contract.DepositManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManager *DepositManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManager.Contract.DepositManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManager *DepositManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManager.Contract.DepositManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DepositManager *DepositManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DepositManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DepositManager *DepositManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DepositManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DepositManager *DepositManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DepositManager.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_DepositManager *DepositManagerCaller) DEPOSITSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DepositManager.contract.Call(opts, &out, "DEPOSIT_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_DepositManager *DepositManagerSession) DEPOSITSIG() ([32]byte, error) {
	return _DepositManager.Contract.DEPOSITSIG(&_DepositManager.CallOpts)
}

// DEPOSITSIG is a free data retrieval call binding the contract method 0xd41f1771.
//
// Solidity: function DEPOSIT_SIG() view returns(bytes32)
func (_DepositManager *DepositManagerCallerSession) DEPOSITSIG() ([32]byte, error) {
	return _DepositManager.Contract.DEPOSITSIG(&_DepositManager.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_DepositManager *DepositManagerCaller) WITHDRAWSIG(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DepositManager.contract.Call(opts, &out, "WITHDRAW_SIG")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_DepositManager *DepositManagerSession) WITHDRAWSIG() ([32]byte, error) {
	return _DepositManager.Contract.WITHDRAWSIG(&_DepositManager.CallOpts)
}

// WITHDRAWSIG is a free data retrieval call binding the contract method 0xb1768065.
//
// Solidity: function WITHDRAW_SIG() view returns(bytes32)
func (_DepositManager *DepositManagerCallerSession) WITHDRAWSIG() ([32]byte, error) {
	return _DepositManager.Contract.WITHDRAWSIG(&_DepositManager.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_DepositManager *DepositManagerCaller) RegistryManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DepositManager.contract.Call(opts, &out, "registryManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_DepositManager *DepositManagerSession) RegistryManager() (common.Address, error) {
	return _DepositManager.Contract.RegistryManager(&_DepositManager.CallOpts)
}

// RegistryManager is a free data retrieval call binding the contract method 0x0c63109e.
//
// Solidity: function registryManager() view returns(address)
func (_DepositManager *DepositManagerCallerSession) RegistryManager() (common.Address, error) {
	return _DepositManager.Contract.RegistryManager(&_DepositManager.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_DepositManager *DepositManagerTransactor) Deposit(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "deposit", recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_DepositManager *DepositManagerSession) Deposit(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.Deposit(&_DepositManager.TransactOpts, recipient, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address recipient, uint256 amount) returns()
func (_DepositManager *DepositManagerTransactorSession) Deposit(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _DepositManager.Contract.Deposit(&_DepositManager.TransactOpts, recipient, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildDepositHandler, address newExitHelper) returns()
func (_DepositManager *DepositManagerTransactor) Initialize(opts *bind.TransactOpts, newRegistryManager common.Address, newStateSender common.Address, newChildDepositHandler common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "initialize", newRegistryManager, newStateSender, newChildDepositHandler, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildDepositHandler, address newExitHelper) returns()
func (_DepositManager *DepositManagerSession) Initialize(newRegistryManager common.Address, newStateSender common.Address, newChildDepositHandler common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.Initialize(&_DepositManager.TransactOpts, newRegistryManager, newStateSender, newChildDepositHandler, newExitHelper)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address newRegistryManager, address newStateSender, address newChildDepositHandler, address newExitHelper) returns()
func (_DepositManager *DepositManagerTransactorSession) Initialize(newRegistryManager common.Address, newStateSender common.Address, newChildDepositHandler common.Address, newExitHelper common.Address) (*types.Transaction, error) {
	return _DepositManager.Contract.Initialize(&_DepositManager.TransactOpts, newRegistryManager, newStateSender, newChildDepositHandler, newExitHelper)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_DepositManager *DepositManagerTransactor) OnL2StateReceive(opts *bind.TransactOpts, arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _DepositManager.contract.Transact(opts, "onL2StateReceive", arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_DepositManager *DepositManagerSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _DepositManager.Contract.OnL2StateReceive(&_DepositManager.TransactOpts, arg0, sender, data)
}

// OnL2StateReceive is a paid mutator transaction binding the contract method 0xf43cda8b.
//
// Solidity: function onL2StateReceive(uint256 , address sender, bytes data) returns()
func (_DepositManager *DepositManagerTransactorSession) OnL2StateReceive(arg0 *big.Int, sender common.Address, data []byte) (*types.Transaction, error) {
	return _DepositManager.Contract.OnL2StateReceive(&_DepositManager.TransactOpts, arg0, sender, data)
}

// DepositManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the DepositManager contract.
type DepositManagerInitializedIterator struct {
	Event *DepositManagerInitialized // Event containing the contract specifics and raw log

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
func (it *DepositManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerInitialized)
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
		it.Event = new(DepositManagerInitialized)
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
func (it *DepositManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerInitialized represents a Initialized event raised by the DepositManager contract.
type DepositManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DepositManager *DepositManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*DepositManagerInitializedIterator, error) {

	logs, sub, err := _DepositManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &DepositManagerInitializedIterator{contract: _DepositManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_DepositManager *DepositManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *DepositManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _DepositManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerInitialized)
				if err := _DepositManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_DepositManager *DepositManagerFilterer) ParseInitialized(log types.Log) (*DepositManagerInitialized, error) {
	event := new(DepositManagerInitialized)
	if err := _DepositManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerTokenDepositIterator is returned from FilterTokenDeposit and is used to iterate over the raw logs and unpacked data for TokenDeposit events raised by the DepositManager contract.
type DepositManagerTokenDepositIterator struct {
	Event *DepositManagerTokenDeposit // Event containing the contract specifics and raw log

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
func (it *DepositManagerTokenDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerTokenDeposit)
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
		it.Event = new(DepositManagerTokenDeposit)
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
func (it *DepositManagerTokenDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerTokenDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerTokenDeposit represents a TokenDeposit event raised by the DepositManager contract.
type DepositManagerTokenDeposit struct {
	Token     common.Address
	Recipient common.Address
	Depositor common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposit is a free log retrieval operation binding the contract event 0xedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a05.
//
// Solidity: event TokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) FilterTokenDeposit(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DepositManagerTokenDepositIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DepositManager.contract.FilterLogs(opts, "TokenDeposit", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerTokenDepositIterator{contract: _DepositManager.contract, event: "TokenDeposit", logs: logs, sub: sub}, nil
}

// WatchTokenDeposit is a free log subscription operation binding the contract event 0xedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a05.
//
// Solidity: event TokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) WatchTokenDeposit(opts *bind.WatchOpts, sink chan<- *DepositManagerTokenDeposit, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DepositManager.contract.WatchLogs(opts, "TokenDeposit", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerTokenDeposit)
				if err := _DepositManager.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
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

// ParseTokenDeposit is a log parse operation binding the contract event 0xedde939981675ab6f1bccc178784c22dde9c3da9c409aaa5c222157f656f9a05.
//
// Solidity: event TokenDeposit(address indexed token, address indexed recipient, address depositor, uint256 amount)
func (_DepositManager *DepositManagerFilterer) ParseTokenDeposit(log types.Log) (*DepositManagerTokenDeposit, error) {
	event := new(DepositManagerTokenDeposit)
	if err := _DepositManager.contract.UnpackLog(event, "TokenDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositManagerTokenWithdrawIterator is returned from FilterTokenWithdraw and is used to iterate over the raw logs and unpacked data for TokenWithdraw events raised by the DepositManager contract.
type DepositManagerTokenWithdrawIterator struct {
	Event *DepositManagerTokenWithdraw // Event containing the contract specifics and raw log

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
func (it *DepositManagerTokenWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositManagerTokenWithdraw)
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
		it.Event = new(DepositManagerTokenWithdraw)
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
func (it *DepositManagerTokenWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositManagerTokenWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositManagerTokenWithdraw represents a TokenWithdraw event raised by the DepositManager contract.
type DepositManagerTokenWithdraw struct {
	Token      common.Address
	Recipient  common.Address
	Withdrawer common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdraw is a free log retrieval operation binding the contract event 0x0cc8e26c5fe2f82346d6755e53b81d7301a600a4ada570aa7281c99f1c983d5c.
//
// Solidity: event TokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_DepositManager *DepositManagerFilterer) FilterTokenWithdraw(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*DepositManagerTokenWithdrawIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DepositManager.contract.FilterLogs(opts, "TokenWithdraw", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &DepositManagerTokenWithdrawIterator{contract: _DepositManager.contract, event: "TokenWithdraw", logs: logs, sub: sub}, nil
}

// WatchTokenWithdraw is a free log subscription operation binding the contract event 0x0cc8e26c5fe2f82346d6755e53b81d7301a600a4ada570aa7281c99f1c983d5c.
//
// Solidity: event TokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_DepositManager *DepositManagerFilterer) WatchTokenWithdraw(opts *bind.WatchOpts, sink chan<- *DepositManagerTokenWithdraw, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _DepositManager.contract.WatchLogs(opts, "TokenWithdraw", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositManagerTokenWithdraw)
				if err := _DepositManager.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
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

// ParseTokenWithdraw is a log parse operation binding the contract event 0x0cc8e26c5fe2f82346d6755e53b81d7301a600a4ada570aa7281c99f1c983d5c.
//
// Solidity: event TokenWithdraw(address indexed token, address indexed recipient, address withdrawer, uint256 amount)
func (_DepositManager *DepositManagerFilterer) ParseTokenWithdraw(log types.Log) (*DepositManagerTokenWithdraw, error) {
	event := new(DepositManagerTokenWithdraw)
	if err := _DepositManager.contract.UnpackLog(event, "TokenWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
