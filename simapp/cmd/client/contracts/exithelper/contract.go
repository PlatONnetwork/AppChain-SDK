// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exithelper

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

// IExitHelperBatchExitInput is an auto generated low-level Go binding around an user-defined struct.
type IExitHelperBatchExitInput struct {
	BlockNumber  *big.Int
	LeafIndex    *big.Int
	UnhashedLeaf []byte
	Proof        [][32]byte
}

// ExitHelperMetaData contains all meta data concerning the ExitHelper contract.
var ExitHelperMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"batchExit\",\"inputs\":[{\"name\":\"inputs\",\"type\":\"tuple[]\",\"internalType\":\"structIExitHelper.BatchExitInput[]\",\"components\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"leafIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unhashedLeaf\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"caller\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"checkpointManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractICheckpointManager\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"exit\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"leafIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"unhashedLeaf\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newCheckpointManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"processedExits\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"ExitProcessed\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false}]",
	Bin: "0x608080604052346100bf575f549060ff8260081c1661006d575060ff80821603610033575b604051610ac490816100c48239f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f610024565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fdfe604060808152600480361015610013575f80fd5b5f3560e01c90816350607b35146105c7578163aa209cc314610283578163bd88ea7914610257578163c0857ba01461022f578163c4d66de81461008a575063fc9c8d391461005f575f80fd5b34610086575f3660031901126100865760035490516001600160a01b039091168152602090f35b5f80fd5b905034610086576020366003190112610086578035906001600160a01b03821690818303610086575f549260ff8460081c161593848095610222575b801561020b575b156101b15760ff1981166001175f558461019f575b508215159081610194575b501561014257506001600160601b0360a01b600254161760025561010d57005b60207f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989161ff00195f54165f555160018152a1005b608490602085519162461bcd60e51b8352820152602660248201527f4578697448656c7065723a20494e56414c49445f434845434b504f494e545f4d60448201526520a720a3a2a960d11b6064820152fd5b90503b15155f6100ed565b61ffff1916610101175f9081556100e2565b855162461bcd60e51b8152602081850152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b1580156100cd5750600160ff8216146100cd565b50600160ff8216106100c6565b8234610086575f3660031901126100865760025490516001600160a01b039091168152602090f35b823461008657602036600319011261008657602091355f526001825260ff815f20541690519015158152f35b9050346100865760803660031901126100865760448035929067ffffffffffffffff9060248286116100865736602387011215610086578585013595838711610086578181018782019083820197368911610086576064908135888111610086576102f190369083016106bc565b6002546001600160a01b0396908716929060809089906103128615156106ed565b03126100865785359c6103268b8a0161076f565b9b610332878b0161076f565b996084810135918211610086570190806043830112156100865761035c91808c0135908d016107d5565b958d5f5260209e8f600190528d5f205460ff16610578578f916103809136916107d5565b80519101208c51630c34044160e31b81528535868201528a81019190915289358b82015260808682015260848101839052926001600160fb1b03831161008657838f9381809460a49260051b80918484013781010301915afa90811561056e575f91610538575b50156104fe575088918a91835f5260018352895f209760ff1998898154166001179055856001600160601b0360a01b9a338c60035416176003558c5196879687019a63f43cda8b60e01b8c5287015216908401528201606090526084820161044e9161080b565b03601f19810182526104609082610783565b519116905a925f8094938194f17f8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec936104db913d156104f6573d946104a4866107b9565b956104b184519788610783565b86523d5f8a88013e5b600354166003551592836104e0575b5051928392878452159683019061080b565b0390a3005b865f5260018852815f209081541690555f6104c9565b6060946104ba565b885162461bcd60e51b81529081018b90526019818701527822bc34ba2432b63832b91d1024a72b20a624a22fa82927a7a360391b81880152fd5b90508b81813d8311610567575b61054f8183610783565b8101031261008657518015158103610086575f6103e7565b503d610545565b8a513d5f823e3d90fd5b505050505061115160f21b6084927f4578697448656c7065723a20455849545f414c52454144595f50524f434553538960228f8b908f519762461bcd60e51b8952880152860152840152820152fd5b9050346100865760208060031936011261008657813567ffffffffffffffff92838211610086576105fa913691016106bc565b60025490939190610615906001600160a01b031615156106ed565b5f5b84811061062057005b61062b818684610739565b359084610639828886610739565b013591610647828886610739565b88810135601e199182813603018212156100865701908135918783116100865788019082360382136100865761067e858b89610739565b60608101359181360301821215610086570191823592888411610086578901918360051b36038313610086576001966106b695610849565b01610617565b9181601f840112156100865782359167ffffffffffffffff8311610086576020808501948460051b01011161008657565b156106f457565b60405162461bcd60e51b815260206004820152601b60248201527f4578697448656c7065723a204e4f545f494e495449414c495a454400000000006044820152606490fd5b919081101561075b5760051b81013590607e1981360301821215610086570190565b634e487b7160e01b5f52603260045260245ffd5b35906001600160a01b038216820361008657565b90601f8019910116810190811067ffffffffffffffff8211176107a557604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff81116107a557601f01601f191660200190565b9291926107e1826107b9565b916107ef6040519384610783565b829481845281830111610086578281602093845f960137010152565b91908251928382525f5b848110610835575050825f602080949584010152601f8019910116010190565b602081830181015184830182015201610815565b93909291828101916080828403126100865781359560209761086c89850161076f565b9260409661087b88870161076f565b96606087013567ffffffffffffffff81116100865787019080601f8301121561008657818d6108ac933591016107d5565b9860018060a01b03968b5f5260018d5260ff8a5f205416610a7f576108d89088600254169336916107d5565b8051908d01208951630c34044160e31b815260048101969096526024860152604485015260806064850152608484018390526001600160fb1b038311610086578360a48180948e9660051b80918484013781010301915afa908115610a75575f91610a3f575b50156109ff57925f80889488826109fa967f8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec9a9983835260018a526109e48884209a60ff199b60018d8254161790556109d66001600160601b0360a01b9d8e33906003541617600355858c5195869485019963f43cda8b60e01b8b52602486015216604484015260606064840152608483019061080b565b03601f198101835282610783565b5193165af13d156104f6573d946104a4866107b9565b0390a3565b835162461bcd60e51b815260048101889052601960248201527822bc34ba2432b63832b91d1024a72b20a624a22fa82927a7a360391b6044820152606490fd5b90508781813d8311610a6e575b610a568183610783565b8101031261008657518015158103610086575f61093e565b503d610a4c565b85513d5f823e3d90fd5b5050505050505050505050505056fea26469706673582212203112ddb530212c91fd988658b6cccb2dbd4f693094802b99f56c005b62f108bb64736f6c63430008160033",
}

// ExitHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use ExitHelperMetaData.ABI instead.
var ExitHelperABI = ExitHelperMetaData.ABI

// ExitHelperBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ExitHelperMetaData.Bin instead.
var ExitHelperBin = ExitHelperMetaData.Bin

// DeployExitHelper deploys a new platon contract, binding an instance of ExitHelper to it.
func DeployExitHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ExitHelper, error) {
	parsed, err := ExitHelperMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ExitHelperBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExitHelper{ExitHelperCaller: ExitHelperCaller{contract: contract}, ExitHelperTransactor: ExitHelperTransactor{contract: contract}, ExitHelperFilterer: ExitHelperFilterer{contract: contract}}, nil
}

// ExitHelper is an auto generated Go binding around an platon contract.
type ExitHelper struct {
	ExitHelperCaller     // Read-only binding to the contract
	ExitHelperTransactor // Write-only binding to the contract
	ExitHelperFilterer   // Log filterer for contract events
}

// ExitHelperCaller is an auto generated read-only Go binding around an platon contract.
type ExitHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExitHelperTransactor is an auto generated write-only Go binding around an platon contract.
type ExitHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExitHelperFilterer is an auto generated log filtering Go binding around an platon contract events.
type ExitHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExitHelperSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type ExitHelperSession struct {
	Contract     *ExitHelper       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExitHelperCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type ExitHelperCallerSession struct {
	Contract *ExitHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ExitHelperTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type ExitHelperTransactorSession struct {
	Contract     *ExitHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ExitHelperRaw is an auto generated low-level Go binding around an platon contract.
type ExitHelperRaw struct {
	Contract *ExitHelper // Generic contract binding to access the raw methods on
}

// ExitHelperCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type ExitHelperCallerRaw struct {
	Contract *ExitHelperCaller // Generic read-only contract binding to access the raw methods on
}

// ExitHelperTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type ExitHelperTransactorRaw struct {
	Contract *ExitHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExitHelper creates a new instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelper(address common.Address, backend bind.ContractBackend) (*ExitHelper, error) {
	contract, err := bindExitHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExitHelper{ExitHelperCaller: ExitHelperCaller{contract: contract}, ExitHelperTransactor: ExitHelperTransactor{contract: contract}, ExitHelperFilterer: ExitHelperFilterer{contract: contract}}, nil
}

// NewExitHelperCaller creates a new read-only instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelperCaller(address common.Address, caller bind.ContractCaller) (*ExitHelperCaller, error) {
	contract, err := bindExitHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExitHelperCaller{contract: contract}, nil
}

// NewExitHelperTransactor creates a new write-only instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*ExitHelperTransactor, error) {
	contract, err := bindExitHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExitHelperTransactor{contract: contract}, nil
}

// NewExitHelperFilterer creates a new log filterer instance of ExitHelper, bound to a specific deployed contract.
func NewExitHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*ExitHelperFilterer, error) {
	contract, err := bindExitHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExitHelperFilterer{contract: contract}, nil
}

// bindExitHelper binds a generic wrapper to an already deployed contract.
func bindExitHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExitHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExitHelper *ExitHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExitHelper.Contract.ExitHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExitHelper *ExitHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExitHelper.Contract.ExitHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExitHelper *ExitHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExitHelper.Contract.ExitHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExitHelper *ExitHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExitHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExitHelper *ExitHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExitHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExitHelper *ExitHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExitHelper.Contract.contract.Transact(opts, method, params...)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_ExitHelper *ExitHelperCaller) Caller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExitHelper.contract.Call(opts, &out, "caller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_ExitHelper *ExitHelperSession) Caller() (common.Address, error) {
	return _ExitHelper.Contract.Caller(&_ExitHelper.CallOpts)
}

// Caller is a free data retrieval call binding the contract method 0xfc9c8d39.
//
// Solidity: function caller() view returns(address)
func (_ExitHelper *ExitHelperCallerSession) Caller() (common.Address, error) {
	return _ExitHelper.Contract.Caller(&_ExitHelper.CallOpts)
}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_ExitHelper *ExitHelperCaller) CheckpointManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExitHelper.contract.Call(opts, &out, "checkpointManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_ExitHelper *ExitHelperSession) CheckpointManager() (common.Address, error) {
	return _ExitHelper.Contract.CheckpointManager(&_ExitHelper.CallOpts)
}

// CheckpointManager is a free data retrieval call binding the contract method 0xc0857ba0.
//
// Solidity: function checkpointManager() view returns(address)
func (_ExitHelper *ExitHelperCallerSession) CheckpointManager() (common.Address, error) {
	return _ExitHelper.Contract.CheckpointManager(&_ExitHelper.CallOpts)
}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_ExitHelper *ExitHelperCaller) ProcessedExits(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _ExitHelper.contract.Call(opts, &out, "processedExits", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_ExitHelper *ExitHelperSession) ProcessedExits(arg0 *big.Int) (bool, error) {
	return _ExitHelper.Contract.ProcessedExits(&_ExitHelper.CallOpts, arg0)
}

// ProcessedExits is a free data retrieval call binding the contract method 0xbd88ea79.
//
// Solidity: function processedExits(uint256 ) view returns(bool)
func (_ExitHelper *ExitHelperCallerSession) ProcessedExits(arg0 *big.Int) (bool, error) {
	return _ExitHelper.Contract.ProcessedExits(&_ExitHelper.CallOpts, arg0)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_ExitHelper *ExitHelperTransactor) BatchExit(opts *bind.TransactOpts, inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _ExitHelper.contract.Transact(opts, "batchExit", inputs)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_ExitHelper *ExitHelperSession) BatchExit(inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _ExitHelper.Contract.BatchExit(&_ExitHelper.TransactOpts, inputs)
}

// BatchExit is a paid mutator transaction binding the contract method 0x50607b35.
//
// Solidity: function batchExit((uint256,uint256,bytes,bytes32[])[] inputs) returns()
func (_ExitHelper *ExitHelperTransactorSession) BatchExit(inputs []IExitHelperBatchExitInput) (*types.Transaction, error) {
	return _ExitHelper.Contract.BatchExit(&_ExitHelper.TransactOpts, inputs)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_ExitHelper *ExitHelperTransactor) Exit(opts *bind.TransactOpts, blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _ExitHelper.contract.Transact(opts, "exit", blockNumber, leafIndex, unhashedLeaf, proof)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_ExitHelper *ExitHelperSession) Exit(blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _ExitHelper.Contract.Exit(&_ExitHelper.TransactOpts, blockNumber, leafIndex, unhashedLeaf, proof)
}

// Exit is a paid mutator transaction binding the contract method 0xaa209cc3.
//
// Solidity: function exit(uint256 blockNumber, uint256 leafIndex, bytes unhashedLeaf, bytes32[] proof) returns()
func (_ExitHelper *ExitHelperTransactorSession) Exit(blockNumber *big.Int, leafIndex *big.Int, unhashedLeaf []byte, proof [][32]byte) (*types.Transaction, error) {
	return _ExitHelper.Contract.Exit(&_ExitHelper.TransactOpts, blockNumber, leafIndex, unhashedLeaf, proof)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_ExitHelper *ExitHelperTransactor) Initialize(opts *bind.TransactOpts, newCheckpointManager common.Address) (*types.Transaction, error) {
	return _ExitHelper.contract.Transact(opts, "initialize", newCheckpointManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_ExitHelper *ExitHelperSession) Initialize(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _ExitHelper.Contract.Initialize(&_ExitHelper.TransactOpts, newCheckpointManager)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address newCheckpointManager) returns()
func (_ExitHelper *ExitHelperTransactorSession) Initialize(newCheckpointManager common.Address) (*types.Transaction, error) {
	return _ExitHelper.Contract.Initialize(&_ExitHelper.TransactOpts, newCheckpointManager)
}

// ExitHelperExitProcessedIterator is returned from FilterExitProcessed and is used to iterate over the raw logs and unpacked data for ExitProcessed events raised by the ExitHelper contract.
type ExitHelperExitProcessedIterator struct {
	Event *ExitHelperExitProcessed // Event containing the contract specifics and raw log

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
func (it *ExitHelperExitProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExitHelperExitProcessed)
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
		it.Event = new(ExitHelperExitProcessed)
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
func (it *ExitHelperExitProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExitHelperExitProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExitHelperExitProcessed represents a ExitProcessed event raised by the ExitHelper contract.
type ExitHelperExitProcessed struct {
	Id         *big.Int
	Success    bool
	ReturnData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExitProcessed is a free log retrieval operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_ExitHelper *ExitHelperFilterer) FilterExitProcessed(opts *bind.FilterOpts, id []*big.Int, success []bool) (*ExitHelperExitProcessedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ExitHelper.contract.FilterLogs(opts, "ExitProcessed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return &ExitHelperExitProcessedIterator{contract: _ExitHelper.contract, event: "ExitProcessed", logs: logs, sub: sub}, nil
}

// WatchExitProcessed is a free log subscription operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_ExitHelper *ExitHelperFilterer) WatchExitProcessed(opts *bind.WatchOpts, sink chan<- *ExitHelperExitProcessed, id []*big.Int, success []bool) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _ExitHelper.contract.WatchLogs(opts, "ExitProcessed", idRule, successRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExitHelperExitProcessed)
				if err := _ExitHelper.contract.UnpackLog(event, "ExitProcessed", log); err != nil {
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

// ParseExitProcessed is a log parse operation binding the contract event 0x8bbfa0c9bee3785c03700d2a909592286efb83fc7e7002be5764424b9842f7ec.
//
// Solidity: event ExitProcessed(uint256 indexed id, bool indexed success, bytes returnData)
func (_ExitHelper *ExitHelperFilterer) ParseExitProcessed(log types.Log) (*ExitHelperExitProcessed, error) {
	event := new(ExitHelperExitProcessed)
	if err := _ExitHelper.contract.UnpackLog(event, "ExitProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExitHelperInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ExitHelper contract.
type ExitHelperInitializedIterator struct {
	Event *ExitHelperInitialized // Event containing the contract specifics and raw log

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
func (it *ExitHelperInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExitHelperInitialized)
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
		it.Event = new(ExitHelperInitialized)
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
func (it *ExitHelperInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExitHelperInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExitHelperInitialized represents a Initialized event raised by the ExitHelper contract.
type ExitHelperInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExitHelper *ExitHelperFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExitHelperInitializedIterator, error) {

	logs, sub, err := _ExitHelper.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExitHelperInitializedIterator{contract: _ExitHelper.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExitHelper *ExitHelperFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExitHelperInitialized) (event.Subscription, error) {

	logs, sub, err := _ExitHelper.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExitHelperInitialized)
				if err := _ExitHelper.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ExitHelper *ExitHelperFilterer) ParseInitialized(log types.Log) (*ExitHelperInitialized, error) {
	event := new(ExitHelperInitialized)
	if err := _ExitHelper.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
