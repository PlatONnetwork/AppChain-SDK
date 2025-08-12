// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package registrymanager

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

// RegistryInit is an auto generated low-level Go binding around an user-defined struct.
type RegistryInit struct {
	ChildChainId      *big.Int
	Token             common.Address
	StateSender       common.Address
	CheckpointManager common.Address
	ExitHelper        common.Address
	ChildChainManager common.Address
	StakeManager      common.Address
	DepositManager    common.Address
	WithdrawHandler   common.Address
	GovernanceHandler common.Address
}

// RegistryManagerMetaData contains all meta data concerning the RegistryManager contract.
var RegistryManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCheckpointManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getChildChainManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDepositManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getExitHelper\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getGovernanceManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStakeManager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateSender\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWithdrawHandler\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"idFor\",\"inputs\":[{\"name\":\"_contract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"factory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isChildChainId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerChildChain\",\"inputs\":[{\"name\":\"init\",\"type\":\"tuple\",\"internalType\":\"structRegistryInit\",\"components\":[{\"name\":\"childChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stateSender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"checkpointManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"exitHelper\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"childChainManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stakeManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"depositManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"withdrawHandler\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"governanceHandler\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRegistryFactory\",\"inputs\":[{\"name\":\"factory\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"tokenOf\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenizeOf\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ContractAdded\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"_contract\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RegistryFactorySet\",\"inputs\":[{\"name\":\"_factory\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608080604052346100bf575f549060ff8260081c1661006d575060ff80821603610033575b60405161106f90816100c48239f35b60ff90811916175f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160ff8152a15f610024565b62461bcd60e51b815260206004820152602760248201527f496e697469616c697a61626c653a20636f6e747261637420697320696e697469604482015266616c697a696e6760c81b6064820152608490fd5b5f80fdfe6080604052600480361015610012575f80fd5b5f3560e01c80631c886fd514610e91578063312a49ab14610e3a578063333af42e14610de35780633a9560c314610d8c578063485cc95514610b5c57806362e0ed3e14610af0578063715018a614610a9557806382192c2614610a5857806384f08a63146104005780638da5cb5b146103d8578063901a44a714610381578063994a73231461032a578063be8d41d5146102f5578063ea78803f146102c1578063ec68f1af1461026a578063f2fde38b146101da578063f90423fe146101425763fec9aaf7146100e0575f80fd5b3461013e57602036600319011261013e57355f9081527f8c3dfe924544c373b1467173210c19c6d75521e5d0fb3b21258e69f5cc2f289160209081526040909120546001600160a01b0316610136811515610f14565b604051908152f35b5f80fd5b503461013e57602036600319011261013e576001600160a01b03610164610ee8565b165f52606760205260405f205490811561018357602082604051908152f35b60849060206040519162461bcd60e51b8352820152602a60248201527f52656769737472794d616e616765724368696c64446174613a20494e56414c496044820152691117d0d3d395149050d560b21b6064820152fd5b503461013e57602036600319011261013e576101f4610ee8565b906101fd610fc1565b6001600160a01b038216156102175761021582610f79565b005b60849060206040519162461bcd60e51b8352820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152fd5b503461013e57602036600319011261013e57355f9081527f60e78cb0dede97912306a60329c90e362b80a616e1dd79e831a1430c612a988360209081526040909120546001600160a01b0316610136811515610f14565b503461013e57602036600319011261013e57355f526069602052602060018060a01b0360405f205416610136811515610f14565b503461013e57602036600319011261013e57610320602091355f52606560205260ff60405f20541690565b6040519015158152f35b503461013e57602036600319011261013e57355f9081527f11be16498e6c58109fa5de15d4c95750bcba290e5610d16526a7050e9cc50b8c60209081526040909120546001600160a01b0316610136811515610f14565b503461013e57602036600319011261013e57355f9081527f3911e5eeb62ff60e2c56e47545e4907690684b4952df7e48ce54e1128c6611e060209081526040909120546001600160a01b0316610136811515610f14565b3461013e575f36600319011261013e576033546040516001600160a01b039091168152602090f35b503461013e57610140908160031936011261013e57609d546001600160a01b031633148015610a44575b156109f65760405191820182811067ffffffffffffffff8211176109e357604052803590818352610459610efe565b60208401908152916044356001600160a01b038116810361013e5760408501526064356001600160a01b038116810361013e5760608501526084356001600160a01b038116810361013e57608085015260a4356001600160a01b038116810361013e5760a085015260c4356001600160a01b038116810361013e5760c085015260e4356001600160a01b038116810361013e5760e0850152610104356001600160a01b038116810361013e57610100850152610124356001600160a01b038116810361013e576105459161053f916101208701525f52606560205260ff60405f20541690565b15610f14565b81516001600160a01b031615610987575081515f52606560205260405f2060ff199060018282541617905560018060a01b038251165f52606a602052600160405f209182541617905560018060a01b0390511681515f52606960205260405f20906001600160601b0360a01b8254161790557fa6604f6f9e958c3372fa784685d6216654aef3be0a2255a92dfbab50f7d0b8545f8051602061101a8339815191526040835160018060a01b038286015116845f526066602052825f20825f52602052825f20816001600160601b0360a01b825416179055805f52606760205281835f2055606860205284835f205582519182526020820152a280515f8051602061101a833981519152604060018060a01b03606085015116927f67ad25e2500d2587e09a708c08f2e60b448c098548a678ba2e5f76c73b79462c93845f526066602052825f20825f52602052825f20816001600160601b0360a01b825416179055805f52606760205281835f2055606860205284835f205582519182526020820152a280515f8051602061101a833981519152604060018060a01b03608085015116927f9a99c9f8222cebf6db4e5de38508c68e5fd909dc83df01f3744575c89926eef793845f526066602052825f20825f52602052825f20816001600160601b0360a01b825416179055805f52606760205281835f2055606860205284835f205582519182526020820152a280515f8051602061101a833981519152604060018060a01b0360a085015116927fc49826ee6f2781a5d4dab675a9b7a6eee240ad636e92bb25dbfbf06e526064f693845f526066602052825f20825f52602052825f20816001600160601b0360a01b825416179055805f52606760205281835f2055606860205284835f205582519182526020820152a280515f8051602061101a833981519152604060018060a01b0360c085015116927f56e86af72b94d3aa725a2e35243d6acbf3dc1ada7212033defd5140c5fcb6a9d93845f526066602052825f20825f52602052825f20816001600160601b0360a01b825416179055805f52606760205281835f2055606860205284835f205582519182526020820152a280515f8051602061101a833981519152604060018060a01b0360e085015116927f396a39c7e290685f408e5373e677285002a403b06145527a7a84a38a30d9ef1093845f526066602052825f20825f52602052825f20816001600160601b0360a01b825416179055805f52606760205281835f2055606860205284835f205582519182526020820152a25f8051602061101a833981519152604082519261012060018060a01b0391015116927fbf6982b8971dec85d8e0112d468bc219587904bbfae37e7509ab2241912d7fce93845f526066602052825f20825f52602052825f20816001600160601b0360a01b825416179055805f52606760205281835f2055606860205284835f205582519182526020820152a2602060405160018152f35b60849060206040519162461bcd60e51b8352820152602f60248201527f52656769737472794d616e616765724368696c64446174613a20494e56414c4960448201526e445f544f4b454e5f4144445245535360881b6064820152fd5b604182634e487b7160e01b5f525260245ffd5b60849060206040519162461bcd60e51b8352820152602160248201527f52656769737472794d616e616765723a20494e56414c49445f4f50455241544f6044820152602960f91b6064820152fd5b506033546001600160a01b0316331461042a565b3461013e57602036600319011261013e576001600160a01b03610a79610ee8565b165f52606a602052602060ff60405f2054166040519015158152f35b3461013e575f36600319011261013e57610aad610fc1565b603380546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461013e57602036600319011261013e577fa52fc0e9feff5c76ba6970bc828096ea2a1d1568431a177b6ccc106c63b309806020610b2c610ee8565b610b34610fc1565b609d80546001600160a01b0319166001600160a01b03929092169182179055604051908152a1005b503461013e57604036600319011261013e57610b76610ee8565b610b7e610efe565b905f549260ff8460081c161593848095610d7f575b8015610d68575b15610d0d5760ff1981166001175f5584610cfb575b506001600160a01b0382811615610cb7578316928315159081610cac575b5015610c5657507fa52fc0e9feff5c76ba6970bc828096ea2a1d1568431a177b6ccc106c63b3098091610c01602092610f79565b609d80546001600160a01b03191682179055604051908152a1610c2057005b61ff00195f54165f557f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498602060405160018152a1005b60849060206040519162461bcd60e51b8352820152602960248201527f52656769737472794d616e616765723a20494e56414c49445f52454749535452604482015268595f464143544f525960b81b6064820152fd5b90503b15155f610bcd565b60405162461bcd60e51b8152602081840152601e60248201527f52656769737472794d616e616765723a20494e56414c49445f4f574e455200006044820152606490fd5b61ffff1916610101175f908155610baf565b60405162461bcd60e51b8152602081840152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b6064820152608490fd5b50303b158015610b9a5750600160ff821614610b9a565b50600160ff821610610b93565b503461013e57602036600319011261013e57355f9081527fd32020d1d00bf75933b0a7b2a33c81a3ea3ae639a2391ba3390c6bf0d081bea160209081526040909120546001600160a01b0316610136811515610f14565b503461013e57602036600319011261013e57355f9081527f19a33844d41777a46798aabacb0cc282149fc714638433e60d0eb1298546c3aa60209081526040909120546001600160a01b0316610136811515610f14565b503461013e57602036600319011261013e57355f9081527f62e63321775e0d155584bc57d50654a7ddb332aac538423d38f77b9e8d13cd4e60209081526040909120546001600160a01b0316610136811515610f14565b503461013e57602036600319011261013e57355f9081527f8e3270a6517de1a6b43eafffa6c3615025cdb8aac96397bc62609241045c89ea60209081526040909120546001600160a01b0316610136811515610f14565b600435906001600160a01b038216820361013e57565b602435906001600160a01b038216820361013e57565b15610f1b57565b60405162461bcd60e51b815260206004820152603060248201527f52656769737472794d616e616765724368696c64446174613a20494e56414c4960448201526f1117d0d212531117d0d210525397d25160821b6064820152608490fd5b603380546001600160a01b039283166001600160a01b0319821681179092559091167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b6033546001600160a01b03163303610fd557565b606460405162461bcd60e51b815260206004820152602060248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152fdfe70f25c84ae6bbb35a4a4a2844b2e022a224d83074aa3ac2a728f1d4c214ee4c6a26469706673582212206dda3e2b08e1383696b833e2e085352c609ee82f5b65aac895e810555762173b64736f6c63430008160033",
}

// RegistryManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistryManagerMetaData.ABI instead.
var RegistryManagerABI = RegistryManagerMetaData.ABI

// RegistryManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RegistryManagerMetaData.Bin instead.
var RegistryManagerBin = RegistryManagerMetaData.Bin

// DeployRegistryManager deploys a new platon contract, binding an instance of RegistryManager to it.
func DeployRegistryManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RegistryManager, error) {
	parsed, err := RegistryManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RegistryManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RegistryManager{RegistryManagerCaller: RegistryManagerCaller{contract: contract}, RegistryManagerTransactor: RegistryManagerTransactor{contract: contract}, RegistryManagerFilterer: RegistryManagerFilterer{contract: contract}}, nil
}

// RegistryManager is an auto generated Go binding around an platon contract.
type RegistryManager struct {
	RegistryManagerCaller     // Read-only binding to the contract
	RegistryManagerTransactor // Write-only binding to the contract
	RegistryManagerFilterer   // Log filterer for contract events
}

// RegistryManagerCaller is an auto generated read-only Go binding around an platon contract.
type RegistryManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryManagerTransactor is an auto generated write-only Go binding around an platon contract.
type RegistryManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryManagerFilterer is an auto generated log filtering Go binding around an platon contract events.
type RegistryManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistryManagerSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type RegistryManagerSession struct {
	Contract     *RegistryManager  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistryManagerCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type RegistryManagerCallerSession struct {
	Contract *RegistryManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// RegistryManagerTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type RegistryManagerTransactorSession struct {
	Contract     *RegistryManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// RegistryManagerRaw is an auto generated low-level Go binding around an platon contract.
type RegistryManagerRaw struct {
	Contract *RegistryManager // Generic contract binding to access the raw methods on
}

// RegistryManagerCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type RegistryManagerCallerRaw struct {
	Contract *RegistryManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RegistryManagerTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type RegistryManagerTransactorRaw struct {
	Contract *RegistryManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistryManager creates a new instance of RegistryManager, bound to a specific deployed contract.
func NewRegistryManager(address common.Address, backend bind.ContractBackend) (*RegistryManager, error) {
	contract, err := bindRegistryManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistryManager{RegistryManagerCaller: RegistryManagerCaller{contract: contract}, RegistryManagerTransactor: RegistryManagerTransactor{contract: contract}, RegistryManagerFilterer: RegistryManagerFilterer{contract: contract}}, nil
}

// NewRegistryManagerCaller creates a new read-only instance of RegistryManager, bound to a specific deployed contract.
func NewRegistryManagerCaller(address common.Address, caller bind.ContractCaller) (*RegistryManagerCaller, error) {
	contract, err := bindRegistryManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryManagerCaller{contract: contract}, nil
}

// NewRegistryManagerTransactor creates a new write-only instance of RegistryManager, bound to a specific deployed contract.
func NewRegistryManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistryManagerTransactor, error) {
	contract, err := bindRegistryManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistryManagerTransactor{contract: contract}, nil
}

// NewRegistryManagerFilterer creates a new log filterer instance of RegistryManager, bound to a specific deployed contract.
func NewRegistryManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistryManagerFilterer, error) {
	contract, err := bindRegistryManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistryManagerFilterer{contract: contract}, nil
}

// bindRegistryManager binds a generic wrapper to an already deployed contract.
func bindRegistryManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistryManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryManager *RegistryManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryManager.Contract.RegistryManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryManager *RegistryManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryManager.Contract.RegistryManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryManager *RegistryManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryManager.Contract.RegistryManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistryManager *RegistryManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistryManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistryManager *RegistryManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistryManager *RegistryManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistryManager.Contract.contract.Transact(opts, method, params...)
}

// GetCheckpointManager is a free data retrieval call binding the contract method 0xfec9aaf7.
//
// Solidity: function getCheckpointManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetCheckpointManager(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getCheckpointManager", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCheckpointManager is a free data retrieval call binding the contract method 0xfec9aaf7.
//
// Solidity: function getCheckpointManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetCheckpointManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetCheckpointManager(&_RegistryManager.CallOpts, id)
}

// GetCheckpointManager is a free data retrieval call binding the contract method 0xfec9aaf7.
//
// Solidity: function getCheckpointManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetCheckpointManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetCheckpointManager(&_RegistryManager.CallOpts, id)
}

// GetChildChainManager is a free data retrieval call binding the contract method 0x333af42e.
//
// Solidity: function getChildChainManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetChildChainManager(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getChildChainManager", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChildChainManager is a free data retrieval call binding the contract method 0x333af42e.
//
// Solidity: function getChildChainManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetChildChainManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetChildChainManager(&_RegistryManager.CallOpts, id)
}

// GetChildChainManager is a free data retrieval call binding the contract method 0x333af42e.
//
// Solidity: function getChildChainManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetChildChainManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetChildChainManager(&_RegistryManager.CallOpts, id)
}

// GetDepositManager is a free data retrieval call binding the contract method 0xec68f1af.
//
// Solidity: function getDepositManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetDepositManager(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getDepositManager", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDepositManager is a free data retrieval call binding the contract method 0xec68f1af.
//
// Solidity: function getDepositManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetDepositManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetDepositManager(&_RegistryManager.CallOpts, id)
}

// GetDepositManager is a free data retrieval call binding the contract method 0xec68f1af.
//
// Solidity: function getDepositManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetDepositManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetDepositManager(&_RegistryManager.CallOpts, id)
}

// GetExitHelper is a free data retrieval call binding the contract method 0x1c886fd5.
//
// Solidity: function getExitHelper(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetExitHelper(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getExitHelper", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetExitHelper is a free data retrieval call binding the contract method 0x1c886fd5.
//
// Solidity: function getExitHelper(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetExitHelper(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetExitHelper(&_RegistryManager.CallOpts, id)
}

// GetExitHelper is a free data retrieval call binding the contract method 0x1c886fd5.
//
// Solidity: function getExitHelper(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetExitHelper(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetExitHelper(&_RegistryManager.CallOpts, id)
}

// GetGovernanceManager is a free data retrieval call binding the contract method 0x3a9560c3.
//
// Solidity: function getGovernanceManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetGovernanceManager(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getGovernanceManager", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGovernanceManager is a free data retrieval call binding the contract method 0x3a9560c3.
//
// Solidity: function getGovernanceManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetGovernanceManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetGovernanceManager(&_RegistryManager.CallOpts, id)
}

// GetGovernanceManager is a free data retrieval call binding the contract method 0x3a9560c3.
//
// Solidity: function getGovernanceManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetGovernanceManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetGovernanceManager(&_RegistryManager.CallOpts, id)
}

// GetStakeManager is a free data retrieval call binding the contract method 0x312a49ab.
//
// Solidity: function getStakeManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetStakeManager(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getStakeManager", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakeManager is a free data retrieval call binding the contract method 0x312a49ab.
//
// Solidity: function getStakeManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetStakeManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetStakeManager(&_RegistryManager.CallOpts, id)
}

// GetStakeManager is a free data retrieval call binding the contract method 0x312a49ab.
//
// Solidity: function getStakeManager(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetStakeManager(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetStakeManager(&_RegistryManager.CallOpts, id)
}

// GetStateSender is a free data retrieval call binding the contract method 0x994a7323.
//
// Solidity: function getStateSender(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetStateSender(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getStateSender", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStateSender is a free data retrieval call binding the contract method 0x994a7323.
//
// Solidity: function getStateSender(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetStateSender(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetStateSender(&_RegistryManager.CallOpts, id)
}

// GetStateSender is a free data retrieval call binding the contract method 0x994a7323.
//
// Solidity: function getStateSender(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetStateSender(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetStateSender(&_RegistryManager.CallOpts, id)
}

// GetWithdrawHandler is a free data retrieval call binding the contract method 0x901a44a7.
//
// Solidity: function getWithdrawHandler(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) GetWithdrawHandler(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "getWithdrawHandler", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWithdrawHandler is a free data retrieval call binding the contract method 0x901a44a7.
//
// Solidity: function getWithdrawHandler(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) GetWithdrawHandler(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetWithdrawHandler(&_RegistryManager.CallOpts, id)
}

// GetWithdrawHandler is a free data retrieval call binding the contract method 0x901a44a7.
//
// Solidity: function getWithdrawHandler(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) GetWithdrawHandler(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.GetWithdrawHandler(&_RegistryManager.CallOpts, id)
}

// IdFor is a free data retrieval call binding the contract method 0xf90423fe.
//
// Solidity: function idFor(address _contract) view returns(uint256 id)
func (_RegistryManager *RegistryManagerCaller) IdFor(opts *bind.CallOpts, _contract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "idFor", _contract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdFor is a free data retrieval call binding the contract method 0xf90423fe.
//
// Solidity: function idFor(address _contract) view returns(uint256 id)
func (_RegistryManager *RegistryManagerSession) IdFor(_contract common.Address) (*big.Int, error) {
	return _RegistryManager.Contract.IdFor(&_RegistryManager.CallOpts, _contract)
}

// IdFor is a free data retrieval call binding the contract method 0xf90423fe.
//
// Solidity: function idFor(address _contract) view returns(uint256 id)
func (_RegistryManager *RegistryManagerCallerSession) IdFor(_contract common.Address) (*big.Int, error) {
	return _RegistryManager.Contract.IdFor(&_RegistryManager.CallOpts, _contract)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryManager *RegistryManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryManager *RegistryManagerSession) Owner() (common.Address, error) {
	return _RegistryManager.Contract.Owner(&_RegistryManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) Owner() (common.Address, error) {
	return _RegistryManager.Contract.Owner(&_RegistryManager.CallOpts)
}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCaller) TokenOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "tokenOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerSession) TokenOf(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.TokenOf(&_RegistryManager.CallOpts, id)
}

// TokenOf is a free data retrieval call binding the contract method 0xea78803f.
//
// Solidity: function tokenOf(uint256 id) view returns(address)
func (_RegistryManager *RegistryManagerCallerSession) TokenOf(id *big.Int) (common.Address, error) {
	return _RegistryManager.Contract.TokenOf(&_RegistryManager.CallOpts, id)
}

// TokenizeOf is a free data retrieval call binding the contract method 0x82192c26.
//
// Solidity: function tokenizeOf(address token) view returns(bool)
func (_RegistryManager *RegistryManagerCaller) TokenizeOf(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _RegistryManager.contract.Call(opts, &out, "tokenizeOf", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenizeOf is a free data retrieval call binding the contract method 0x82192c26.
//
// Solidity: function tokenizeOf(address token) view returns(bool)
func (_RegistryManager *RegistryManagerSession) TokenizeOf(token common.Address) (bool, error) {
	return _RegistryManager.Contract.TokenizeOf(&_RegistryManager.CallOpts, token)
}

// TokenizeOf is a free data retrieval call binding the contract method 0x82192c26.
//
// Solidity: function tokenizeOf(address token) view returns(bool)
func (_RegistryManager *RegistryManagerCallerSession) TokenizeOf(token common.Address) (bool, error) {
	return _RegistryManager.Contract.TokenizeOf(&_RegistryManager.CallOpts, token)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newOwner, address factory) returns()
func (_RegistryManager *RegistryManagerTransactor) Initialize(opts *bind.TransactOpts, newOwner common.Address, factory common.Address) (*types.Transaction, error) {
	return _RegistryManager.contract.Transact(opts, "initialize", newOwner, factory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newOwner, address factory) returns()
func (_RegistryManager *RegistryManagerSession) Initialize(newOwner common.Address, factory common.Address) (*types.Transaction, error) {
	return _RegistryManager.Contract.Initialize(&_RegistryManager.TransactOpts, newOwner, factory)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address newOwner, address factory) returns()
func (_RegistryManager *RegistryManagerTransactorSession) Initialize(newOwner common.Address, factory common.Address) (*types.Transaction, error) {
	return _RegistryManager.Contract.Initialize(&_RegistryManager.TransactOpts, newOwner, factory)
}

// IsChildChainId is a paid mutator transaction binding the contract method 0xbe8d41d5.
//
// Solidity: function isChildChainId(uint256 chainId) returns(bool)
func (_RegistryManager *RegistryManagerTransactor) IsChildChainId(opts *bind.TransactOpts, chainId *big.Int) (*types.Transaction, error) {
	return _RegistryManager.contract.Transact(opts, "isChildChainId", chainId)
}

// IsChildChainId is a paid mutator transaction binding the contract method 0xbe8d41d5.
//
// Solidity: function isChildChainId(uint256 chainId) returns(bool)
func (_RegistryManager *RegistryManagerSession) IsChildChainId(chainId *big.Int) (*types.Transaction, error) {
	return _RegistryManager.Contract.IsChildChainId(&_RegistryManager.TransactOpts, chainId)
}

// IsChildChainId is a paid mutator transaction binding the contract method 0xbe8d41d5.
//
// Solidity: function isChildChainId(uint256 chainId) returns(bool)
func (_RegistryManager *RegistryManagerTransactorSession) IsChildChainId(chainId *big.Int) (*types.Transaction, error) {
	return _RegistryManager.Contract.IsChildChainId(&_RegistryManager.TransactOpts, chainId)
}

// RegisterChildChain is a paid mutator transaction binding the contract method 0x84f08a63.
//
// Solidity: function registerChildChain((uint256,address,address,address,address,address,address,address,address,address) init) returns(bool)
func (_RegistryManager *RegistryManagerTransactor) RegisterChildChain(opts *bind.TransactOpts, init RegistryInit) (*types.Transaction, error) {
	return _RegistryManager.contract.Transact(opts, "registerChildChain", init)
}

// RegisterChildChain is a paid mutator transaction binding the contract method 0x84f08a63.
//
// Solidity: function registerChildChain((uint256,address,address,address,address,address,address,address,address,address) init) returns(bool)
func (_RegistryManager *RegistryManagerSession) RegisterChildChain(init RegistryInit) (*types.Transaction, error) {
	return _RegistryManager.Contract.RegisterChildChain(&_RegistryManager.TransactOpts, init)
}

// RegisterChildChain is a paid mutator transaction binding the contract method 0x84f08a63.
//
// Solidity: function registerChildChain((uint256,address,address,address,address,address,address,address,address,address) init) returns(bool)
func (_RegistryManager *RegistryManagerTransactorSession) RegisterChildChain(init RegistryInit) (*types.Transaction, error) {
	return _RegistryManager.Contract.RegisterChildChain(&_RegistryManager.TransactOpts, init)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistryManager *RegistryManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistryManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistryManager *RegistryManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegistryManager.Contract.RenounceOwnership(&_RegistryManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistryManager *RegistryManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegistryManager.Contract.RenounceOwnership(&_RegistryManager.TransactOpts)
}

// SetRegistryFactory is a paid mutator transaction binding the contract method 0x62e0ed3e.
//
// Solidity: function setRegistryFactory(address factory) returns()
func (_RegistryManager *RegistryManagerTransactor) SetRegistryFactory(opts *bind.TransactOpts, factory common.Address) (*types.Transaction, error) {
	return _RegistryManager.contract.Transact(opts, "setRegistryFactory", factory)
}

// SetRegistryFactory is a paid mutator transaction binding the contract method 0x62e0ed3e.
//
// Solidity: function setRegistryFactory(address factory) returns()
func (_RegistryManager *RegistryManagerSession) SetRegistryFactory(factory common.Address) (*types.Transaction, error) {
	return _RegistryManager.Contract.SetRegistryFactory(&_RegistryManager.TransactOpts, factory)
}

// SetRegistryFactory is a paid mutator transaction binding the contract method 0x62e0ed3e.
//
// Solidity: function setRegistryFactory(address factory) returns()
func (_RegistryManager *RegistryManagerTransactorSession) SetRegistryFactory(factory common.Address) (*types.Transaction, error) {
	return _RegistryManager.Contract.SetRegistryFactory(&_RegistryManager.TransactOpts, factory)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistryManager *RegistryManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RegistryManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistryManager *RegistryManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistryManager.Contract.TransferOwnership(&_RegistryManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistryManager *RegistryManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistryManager.Contract.TransferOwnership(&_RegistryManager.TransactOpts, newOwner)
}

// RegistryManagerContractAddedIterator is returned from FilterContractAdded and is used to iterate over the raw logs and unpacked data for ContractAdded events raised by the RegistryManager contract.
type RegistryManagerContractAddedIterator struct {
	Event *RegistryManagerContractAdded // Event containing the contract specifics and raw log

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
func (it *RegistryManagerContractAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryManagerContractAdded)
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
		it.Event = new(RegistryManagerContractAdded)
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
func (it *RegistryManagerContractAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryManagerContractAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryManagerContractAdded represents a ContractAdded event raised by the RegistryManager contract.
type RegistryManagerContractAdded struct {
	Key      [32]byte
	Id       *big.Int
	Contract common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterContractAdded is a free log retrieval operation binding the contract event 0x70f25c84ae6bbb35a4a4a2844b2e022a224d83074aa3ac2a728f1d4c214ee4c6.
//
// Solidity: event ContractAdded(bytes32 indexed key, uint256 id, address _contract)
func (_RegistryManager *RegistryManagerFilterer) FilterContractAdded(opts *bind.FilterOpts, key [][32]byte) (*RegistryManagerContractAddedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _RegistryManager.contract.FilterLogs(opts, "ContractAdded", keyRule)
	if err != nil {
		return nil, err
	}
	return &RegistryManagerContractAddedIterator{contract: _RegistryManager.contract, event: "ContractAdded", logs: logs, sub: sub}, nil
}

// WatchContractAdded is a free log subscription operation binding the contract event 0x70f25c84ae6bbb35a4a4a2844b2e022a224d83074aa3ac2a728f1d4c214ee4c6.
//
// Solidity: event ContractAdded(bytes32 indexed key, uint256 id, address _contract)
func (_RegistryManager *RegistryManagerFilterer) WatchContractAdded(opts *bind.WatchOpts, sink chan<- *RegistryManagerContractAdded, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _RegistryManager.contract.WatchLogs(opts, "ContractAdded", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryManagerContractAdded)
				if err := _RegistryManager.contract.UnpackLog(event, "ContractAdded", log); err != nil {
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

// ParseContractAdded is a log parse operation binding the contract event 0x70f25c84ae6bbb35a4a4a2844b2e022a224d83074aa3ac2a728f1d4c214ee4c6.
//
// Solidity: event ContractAdded(bytes32 indexed key, uint256 id, address _contract)
func (_RegistryManager *RegistryManagerFilterer) ParseContractAdded(log types.Log) (*RegistryManagerContractAdded, error) {
	event := new(RegistryManagerContractAdded)
	if err := _RegistryManager.contract.UnpackLog(event, "ContractAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RegistryManager contract.
type RegistryManagerInitializedIterator struct {
	Event *RegistryManagerInitialized // Event containing the contract specifics and raw log

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
func (it *RegistryManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryManagerInitialized)
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
		it.Event = new(RegistryManagerInitialized)
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
func (it *RegistryManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryManagerInitialized represents a Initialized event raised by the RegistryManager contract.
type RegistryManagerInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RegistryManager *RegistryManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistryManagerInitializedIterator, error) {

	logs, sub, err := _RegistryManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistryManagerInitializedIterator{contract: _RegistryManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RegistryManager *RegistryManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistryManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _RegistryManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryManagerInitialized)
				if err := _RegistryManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_RegistryManager *RegistryManagerFilterer) ParseInitialized(log types.Log) (*RegistryManagerInitialized, error) {
	event := new(RegistryManagerInitialized)
	if err := _RegistryManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegistryManager contract.
type RegistryManagerOwnershipTransferredIterator struct {
	Event *RegistryManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistryManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryManagerOwnershipTransferred)
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
		it.Event = new(RegistryManagerOwnershipTransferred)
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
func (it *RegistryManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryManagerOwnershipTransferred represents a OwnershipTransferred event raised by the RegistryManager contract.
type RegistryManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistryManager *RegistryManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistryManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistryManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistryManagerOwnershipTransferredIterator{contract: _RegistryManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistryManager *RegistryManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistryManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistryManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryManagerOwnershipTransferred)
				if err := _RegistryManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistryManager *RegistryManagerFilterer) ParseOwnershipTransferred(log types.Log) (*RegistryManagerOwnershipTransferred, error) {
	event := new(RegistryManagerOwnershipTransferred)
	if err := _RegistryManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistryManagerRegistryFactorySetIterator is returned from FilterRegistryFactorySet and is used to iterate over the raw logs and unpacked data for RegistryFactorySet events raised by the RegistryManager contract.
type RegistryManagerRegistryFactorySetIterator struct {
	Event *RegistryManagerRegistryFactorySet // Event containing the contract specifics and raw log

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
func (it *RegistryManagerRegistryFactorySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistryManagerRegistryFactorySet)
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
		it.Event = new(RegistryManagerRegistryFactorySet)
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
func (it *RegistryManagerRegistryFactorySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistryManagerRegistryFactorySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistryManagerRegistryFactorySet represents a RegistryFactorySet event raised by the RegistryManager contract.
type RegistryManagerRegistryFactorySet struct {
	Factory common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistryFactorySet is a free log retrieval operation binding the contract event 0xa52fc0e9feff5c76ba6970bc828096ea2a1d1568431a177b6ccc106c63b30980.
//
// Solidity: event RegistryFactorySet(address _factory)
func (_RegistryManager *RegistryManagerFilterer) FilterRegistryFactorySet(opts *bind.FilterOpts) (*RegistryManagerRegistryFactorySetIterator, error) {

	logs, sub, err := _RegistryManager.contract.FilterLogs(opts, "RegistryFactorySet")
	if err != nil {
		return nil, err
	}
	return &RegistryManagerRegistryFactorySetIterator{contract: _RegistryManager.contract, event: "RegistryFactorySet", logs: logs, sub: sub}, nil
}

// WatchRegistryFactorySet is a free log subscription operation binding the contract event 0xa52fc0e9feff5c76ba6970bc828096ea2a1d1568431a177b6ccc106c63b30980.
//
// Solidity: event RegistryFactorySet(address _factory)
func (_RegistryManager *RegistryManagerFilterer) WatchRegistryFactorySet(opts *bind.WatchOpts, sink chan<- *RegistryManagerRegistryFactorySet) (event.Subscription, error) {

	logs, sub, err := _RegistryManager.contract.WatchLogs(opts, "RegistryFactorySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistryManagerRegistryFactorySet)
				if err := _RegistryManager.contract.UnpackLog(event, "RegistryFactorySet", log); err != nil {
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

// ParseRegistryFactorySet is a log parse operation binding the contract event 0xa52fc0e9feff5c76ba6970bc828096ea2a1d1568431a177b6ccc106c63b30980.
//
// Solidity: event RegistryFactorySet(address _factory)
func (_RegistryManager *RegistryManagerFilterer) ParseRegistryFactorySet(log types.Log) (*RegistryManagerRegistryFactorySet, error) {
	event := new(RegistryManagerRegistryFactorySet)
	if err := _RegistryManager.contract.UnpackLog(event, "RegistryFactorySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
