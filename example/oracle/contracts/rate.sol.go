// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// QuorumCert is an auto generated low-level Go binding around an user-defined struct.
type QuorumCert struct {
	Epoch       uint64
	ViewNumber  uint64
	BlockHash   [32]byte
	BlockNumber uint64
	BlockIndex  uint32
	ExtendHash  [32]byte
}

// RoundNodeList is an auto generated low-level Go binding around an user-defined struct.
type RoundNodeList struct {
	Nodes []string
	Epoch *big.Int
	Start *big.Int
	End   *big.Int
}

// RateMetaData contains all meta data concerning the Rate contract.
var RateMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"findNodes\",\"inputs\":[{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getCurrentRoundValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRoundNodeList\",\"components\":[{\"name\":\"nodes\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getLastRoundValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structRoundNodeList\",\"components\":[{\"name\":\"nodes\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"newRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"qc\",\"type\":\"tuple\",\"internalType\":\"structQuorumCert\",\"components\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"viewNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"bitmap\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"leafIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"UpdateRate\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611fb2380380611fb283398101604081905261002f91610048565b6002805460ff191660ff92909216919091179055610072565b60006020828403121561005a57600080fd5b815160ff8116811461006b57600080fd5b9392505050565b611f31806100816000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063d1e2cbb91161005b578063d1e2cbb9146100c6578063f866c30b146100db578063fa811847146100e3578063fc8045c31461010357600080fd5b80632c4e722e14610082578063313ce5671461009e57806357e871e7146100bd575b600080fd5b61008b60015481565b6040519081526020015b60405180910390f35b6002546100ab9060ff1681565b60405160ff9091168152602001610095565b61008b60005481565b6100ce610118565b6040516100959190611ab0565b6100ce6101e0565b6100f66100f136600461178a565b610289565b60405161009591906119b6565b6101166101113660046117a3565b6102f2565b005b6101436040518060800160405280606081526020016000815260200160008152602001600081525090565b6040805160048152602481018252602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd1e2cbb900000000000000000000000000000000000000000000000000000000179052825160608101909352602880845291926101c19260659285929190611ed490830139610489565b90506000818060200190518101906101d9919061167b565b9392505050565b61020b6040518060800160405280606081526020016000815260200160008152602001600081525090565b6040805160048152602481018252602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167ff866c30b00000000000000000000000000000000000000000000000000000000179052825160608101909352602580845291926101c19260659285929190611eaf90830139610489565b606060006102956101e0565b9050828160400151111580156102af575082816060015110155b156102bb575192915050565b6102c3610118565b9050828160400151111580156102dd575082816060015110155b156102e9575192915050565b50606092915050565b876060015167ffffffffffffffff1660005410610370576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f526174653a206578706972656420626c6f636b2071630000000000000000000060448201526064015b60405180910390fd5b600061037b8a6104a0565b8051602082012060a08b01519192509061039a908290879087876104b9565b610400576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601a60248201527f526174653a20696e76616c6964206d65726b6c652070726f6f660000000000006044820152606401610367565b600061040b8b61050d565b90506104298b6060015167ffffffffffffffff16828a8a8e8e6106a9565b60018c905560608b015167ffffffffffffffff16600081905560408051918252602082018e90527ffe3b4c3f6e2efdc3dd612c3eb2817308b7568ad72d3f9c9200a08a069c27dff8910160405180910390a1505050505050505050505050565b6060610498848460008561084b565b949350505050565b60606104b36104ae83610964565b610ad2565b92915050565b60008115610502578160051b8301835b8660011660051b8881528135602082185250604060002097508660011c96506020810190508181106104fa576104ff565b6104c9565b50505b505050901591141690565b60408051600680825260e08201909252600091829190816020015b606081526020019060019003908161052857505083519091506105549067ffffffffffffffff166104a0565b8160008151811061056757610567611e37565b6020026020010181905250610589836020015167ffffffffffffffff166104a0565b8160018151811061059c5761059c611e37565b60200260200101819052506105d583604001516040516020016105c191815260200190565b604051602081830303815290604052610ad2565b816002815181106105e8576105e8611e37565b602002602001018190525061060a836060015167ffffffffffffffff166104a0565b8160038151811061061d5761061d611e37565b602002602001018190525061063b836080015163ffffffff166104a0565b8160048151811061064e5761064e611e37565b60200260200101819052506106738360a001516040516020016105c191815260200190565b8160058151811061068657610686611e37565b602002602001018190525061069a81610b21565b80519060200120915050919050565b60006106eb8784848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250610b4592505050565b905060008686868460405160240161070694939291906119c9565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f5fbe99b3000000000000000000000000000000000000000000000000000000001790528151808301909252601982527f526174653a207665726966792070726f6f66206661696c656400000000000000908201529091506107bf9060cc908390610489565b90506000818060200190518101906107d791906114c9565b905080610840576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f526174653a207665726966792070726f6f66206661696c6564000000000000006044820152606401610367565b505050505050505050565b6060824710156108dd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c00000000000000000000000000000000000000000000000000006064820152608401610367565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051610906919061199a565b60006040518083038185875af1925050503d8060008114610943576040519150601f19603f3d011682016040523d82523d6000602084013e610948565b606091505b509150915061095987838387610c24565b979650505050505050565b6040805160208082528183019092526060916000919060208201818036833701905050905082602082015260005b60208110156109f0578181815181106109ad576109ad611e37565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016156109de576109f0565b806109e881611d8c565b915050610992565b60006109fd826020611d45565b67ffffffffffffffff811115610a1557610a15611e66565b6040519080825280601f01601f191660200182016040528015610a3f576020820181803683370190505b50905060005b8151811015610ac9578383610a5981611d8c565b945081518110610a6b57610a6b611e37565b602001015160f81c60f81b828281518110610a8857610a88611e37565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080610ac181611d8c565b915050610a45565b50949350505050565b60608082516001148015610b005750608083600081518110610af657610af6611e37565b016020015160f81c105b15610b0c5750816104b3565b6101d9610b1b84516080610cc1565b84610ec0565b60606000610b2e83610f5b565b90506101d9610b3f825160c0610cc1565b82610ec0565b60606000610b5284610289565b805190915060008167ffffffffffffffff811115610b7257610b72611e66565b604051908082528060200260200182016040528015610ba557816020015b6060815260200190600190039081610b905790505b50905060005b82811015610c1a57610bbd8682611135565b15610c08576000610be6858381518110610bd957610bd9611e37565b6020026020010151611197565b905080838381518110610bfb57610bfb611e37565b6020026020010181905250505b80610c1281611d8c565b915050610bab565b5095945050505050565b60608315610cb7578251610cb05773ffffffffffffffffffffffffffffffffffffffff85163b610cb0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610367565b5081610498565b6104988383611289565b6060806038841015610d435760408051600180825281830190925290602082018180368337019050509050610cf68385611bb5565b601f1a60f81b81600081518110610d0f57610d0f611e37565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506101d9565b600060015b610d528187611bcd565b15610d785781610d6181611d8c565b9250610d71905061010082611d08565b9050610d48565b610d83826001611bb5565b67ffffffffffffffff811115610d9b57610d9b611e66565b6040519080825280601f01601f191660200182016040528015610dc5576020820181803683370190505b509250610dd28583611bb5565b610ddd906037611bb5565b601f1a60f81b83600081518110610df657610df6611e37565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600190505b818111610eb757610100610e3e8284611d45565b610e4a90610100611c42565b610e549088611bcd565b610e5e9190611dc5565b601f1a60f81b838281518110610e7657610e76611e37565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535080610eaf81611d8c565b915050610e2a565b50509392505050565b6060806040519050835180825260208201818101602087015b81831015610ef1578051835260209283019201610ed9565b50855184518101855292509050808201602086015b81831015610f1e578051835260209283019201610f06565b508651929092011591909101601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660405250905092915050565b6060815160001415610f7b57505060408051600081526020810190915290565b6000805b835181101561106c576000848281518110610f9c57610f9c611e37565b60200260200101515111611032576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f416e206974656d20696e20746865206c69737420746f20626520524c5020656e60448201527f636f646564206973206e756c6c2e0000000000000000000000000000000000006064820152608401610367565b83818151811061104457611044611e37565b602002602001015151826110589190611bb5565b91508061106481611d8c565b915050610f7f565b60008267ffffffffffffffff81111561108757611087611e66565b6040519080825280601f01601f1916602001820160405280156110b1576020820181803683370190505b50600092509050602081015b8551831015610ac95760008684815181106110da576110da611e37565b6020026020010151905060006020820190506110f8838284516112cd565b87858151811061110a5761110a611e37565b6020026020010151518361111e9190611bb5565b92505050828061112d90611d8c565b9350506110bd565b600080611143600884611bcd565b90506000611152600885611dc5565b905084518210611167576000925050506104b3565b60008160ff166001901b86848151811061118357611183611e37565b016020015160f81c16119250505092915050565b60606000826040516024016111ac9190611a9d565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152918152602080830180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fb336ad83000000000000000000000000000000000000000000000000000000001790528151808301909252601582527f526174653a20676574206e6f6465206661696c6564000000000000000000000090820152909150611265906065908390610489565b905060008180602001905181019061127d91906114e4565b60a00151949350505050565b8151156112995781518083602001fd5b806040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103679190611a9d565b8282825b6020811061130957815183526112e8602084611bb5565b92506112f5602083611bb5565b9150611302602082611d45565b90506112d1565b60006001611318836020611d45565b61132490610100611c42565b61132e9190611d45565b925184518416931916929092179092525050505050565b805173ffffffffffffffffffffffffffffffffffffffff8116811461136957600080fd5b919050565b60008083601f84011261138057600080fd5b50813567ffffffffffffffff81111561139857600080fd5b6020830191508360208260051b85010111156113b357600080fd5b9250929050565b8051801515811461136957600080fd5b60008083601f8401126113dc57600080fd5b50813567ffffffffffffffff8111156113f457600080fd5b6020830191508360208285010111156113b357600080fd5b600082601f83011261141d57600080fd5b815167ffffffffffffffff81111561143757611437611e66565b61146860207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f84011601611b66565b81815284602083860101111561147d57600080fd5b610498826020830160208701611d5c565b80516002811061136957600080fd5b80516005811061136957600080fd5b805161ffff8116811461136957600080fd5b805161136981611e95565b6000602082840312156114db57600080fd5b6101d9826113ba565b6000602082840312156114f657600080fd5b815167ffffffffffffffff8082111561150e57600080fd5b908301906101a0828603121561152357600080fd5b61152b611af6565b82518281111561153a57600080fd5b6115468782860161140c565b82525061155560208401611345565b602082015260408301518281111561156c57600080fd5b6115788782860161140c565b60408301525061158a6060840161148e565b60608201526080830151828111156115a157600080fd5b6115ad8782860161140c565b60808301525060a0830151828111156115c557600080fd5b6115d18782860161140c565b60a08301525060c0830151828111156115e957600080fd5b6115f58782860161140c565b60c08301525061160760e084016114ac565b60e0820152610100915061161c8284016114ac565b82820152610120915061163082840161149d565b8282015261014091506116448284016113ba565b8282015261016091506116588284016114be565b82820152610180915061166c8284016114be565b91810191909152949350505050565b6000602080838503121561168e57600080fd5b825167ffffffffffffffff808211156116a657600080fd5b90840190608082870312156116ba57600080fd5b6116c2611b20565b8251828111156116d157600080fd5b8301601f810188136116e257600080fd5b8051838111156116f4576116f4611e66565b8060051b611703878201611b66565b828152878101908489018386018a018d101561171e57600080fd5b60009350835b8581101561175b5781518981111561173a578586fd5b6117488f8d838b010161140c565b855250928a0192908a0190600101611724565b505085525050505082840151938101939093525060408082015190830152606090810151908201529392505050565b60006020828403121561179c57600080fd5b5035919050565b6000806000806000806000806000898b036101608112156117c357600080fd5b8a35995060c07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0820112156117f757600080fd5b50611800611b43565b60208b013561180e81611e95565b815260408b013561181e81611e95565b602082015260608b0135604082015260808b013561183b81611e95565b606082015260a08b013563ffffffff8116811461185757600080fd5b608082015260c08b013560a0820152975060e08a013567ffffffffffffffff8082111561188357600080fd5b61188f8d838e016113ca565b90995097506101008c01359150808211156118a957600080fd5b6118b58d838e016113ca565b90975095506101208c013594506101408c01359150808211156118d757600080fd5b506118e48c828d0161136e565b915080935050809150509295985092959850929598565b600081518084526020808501808196508360051b8101915082860160005b85811015611943578284038952611931848351611950565b98850198935090840190600101611919565b5091979650505050505050565b60008151808452611968816020860160208601611d5c565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b600082516119ac818460208701611d5c565b9190910192915050565b6020815260006101d960208301846118fb565b84815260006020606081840152846060840152848660808501376000608086850101527fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f860116830160808101608085830301604086015280865180835260a08401915060a08160051b850101925084880160005b82811015611a8c577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60868603018452611a7a858351611950565b94509286019290860190600101611a40565b50929b9a5050505050505050505050565b6020815260006101d96020830184611950565b602081526000825160806020840152611acc60a08401826118fb565b90506020840151604084015260408401516060840152606084015160808401528091505092915050565b6040516101a0810167ffffffffffffffff81118282101715611b1a57611b1a611e66565b60405290565b6040516080810167ffffffffffffffff81118282101715611b1a57611b1a611e66565b60405160c0810167ffffffffffffffff81118282101715611b1a57611b1a611e66565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff81118282101715611bad57611bad611e66565b604052919050565b60008219821115611bc857611bc8611dd9565b500190565b600082611bdc57611bdc611e08565b500490565b600181815b80851115611c3a57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115611c2057611c20611dd9565b80851615611c2d57918102915b93841c9390800290611be6565b509250929050565b60006101d98383600082611c58575060016104b3565b81611c65575060006104b3565b8160018114611c7b5760028114611c8557611ca1565b60019150506104b3565b60ff841115611c9657611c96611dd9565b50506001821b6104b3565b5060208310610133831016604e8410600b8410161715611cc4575081810a6104b3565b611cce8383611be1565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04821115611d0057611d00611dd9565b029392505050565b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611d4057611d40611dd9565b500290565b600082821015611d5757611d57611dd9565b500390565b60005b83811015611d77578181015183820152602001611d5f565b83811115611d86576000848401525b50505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611dbe57611dbe611dd9565b5060010190565b600082611dd457611dd4611e08565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b67ffffffffffffffff81168114611eab57600080fd5b5056fe526174653a20676574206c61737420726f756e642076616c696461746f72206661696c6564526174653a206765742063757272656e7420726f756e642076616c696461746f72206661696c6564a2646970667358221220caf2402c1e6aee18267ec22ddbcf8960aa378bf20d21b8f9d2d096031356995a64736f6c63430008070033",
}

// RateABI is the input ABI used to generate the binding from.
// Deprecated: Use RateMetaData.ABI instead.
var RateABI = RateMetaData.ABI

// RateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RateMetaData.Bin instead.
var RateBin = RateMetaData.Bin

// DeployRate deploys a new platon contract, binding an instance of Rate to it.
func DeployRate(auth *bind.TransactOpts, backend bind.ContractBackend, _decimals uint8) (common.Address, *types.Transaction, *Rate, error) {
	parsed, err := RateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RateBin), backend, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rate{RateCaller: RateCaller{contract: contract}, RateTransactor: RateTransactor{contract: contract}, RateFilterer: RateFilterer{contract: contract}}, nil
}

// Rate is an auto generated Go binding around an platon contract.
type Rate struct {
	RateCaller     // Read-only binding to the contract
	RateTransactor // Write-only binding to the contract
	RateFilterer   // Log filterer for contract events
}

// RateCaller is an auto generated read-only Go binding around an platon contract.
type RateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateTransactor is an auto generated write-only Go binding around an platon contract.
type RateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateFilterer is an auto generated log filtering Go binding around an platon contract events.
type RateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RateSession is an auto generated Go binding around an platon contract,
// with pre-set call and transact options.
type RateSession struct {
	Contract     *Rate             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RateCallerSession is an auto generated read-only Go binding around an platon contract,
// with pre-set call options.
type RateCallerSession struct {
	Contract *RateCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RateTransactorSession is an auto generated write-only Go binding around an platon contract,
// with pre-set transact options.
type RateTransactorSession struct {
	Contract     *RateTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RateRaw is an auto generated low-level Go binding around an platon contract.
type RateRaw struct {
	Contract *Rate // Generic contract binding to access the raw methods on
}

// RateCallerRaw is an auto generated low-level read-only Go binding around an platon contract.
type RateCallerRaw struct {
	Contract *RateCaller // Generic read-only contract binding to access the raw methods on
}

// RateTransactorRaw is an auto generated low-level write-only Go binding around an platon contract.
type RateTransactorRaw struct {
	Contract *RateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRate creates a new instance of Rate, bound to a specific deployed contract.
func NewRate(address common.Address, backend bind.ContractBackend) (*Rate, error) {
	contract, err := bindRate(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rate{RateCaller: RateCaller{contract: contract}, RateTransactor: RateTransactor{contract: contract}, RateFilterer: RateFilterer{contract: contract}}, nil
}

// NewRateCaller creates a new read-only instance of Rate, bound to a specific deployed contract.
func NewRateCaller(address common.Address, caller bind.ContractCaller) (*RateCaller, error) {
	contract, err := bindRate(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RateCaller{contract: contract}, nil
}

// NewRateTransactor creates a new write-only instance of Rate, bound to a specific deployed contract.
func NewRateTransactor(address common.Address, transactor bind.ContractTransactor) (*RateTransactor, error) {
	contract, err := bindRate(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RateTransactor{contract: contract}, nil
}

// NewRateFilterer creates a new log filterer instance of Rate, bound to a specific deployed contract.
func NewRateFilterer(address common.Address, filterer bind.ContractFilterer) (*RateFilterer, error) {
	contract, err := bindRate(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RateFilterer{contract: contract}, nil
}

// bindRate binds a generic wrapper to an already deployed contract.
func bindRate(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rate *RateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rate.Contract.RateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rate *RateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rate.Contract.RateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rate *RateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rate.Contract.RateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rate *RateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rate.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rate *RateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rate.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rate *RateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rate.Contract.contract.Transact(opts, method, params...)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Rate *RateCaller) BlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rate.contract.Call(opts, &out, "blockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Rate *RateSession) BlockNumber() (*big.Int, error) {
	return _Rate.Contract.BlockNumber(&_Rate.CallOpts)
}

// BlockNumber is a free data retrieval call binding the contract method 0x57e871e7.
//
// Solidity: function blockNumber() view returns(uint256)
func (_Rate *RateCallerSession) BlockNumber() (*big.Int, error) {
	return _Rate.Contract.BlockNumber(&_Rate.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Rate *RateCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Rate.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Rate *RateSession) Decimals() (uint8, error) {
	return _Rate.Contract.Decimals(&_Rate.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Rate *RateCallerSession) Decimals() (uint8, error) {
	return _Rate.Contract.Decimals(&_Rate.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Rate *RateCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rate.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Rate *RateSession) Rate() (*big.Int, error) {
	return _Rate.Contract.Rate(&_Rate.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Rate *RateCallerSession) Rate() (*big.Int, error) {
	return _Rate.Contract.Rate(&_Rate.CallOpts)
}

// FindNodes is a paid mutator transaction binding the contract method 0xfa811847.
//
// Solidity: function findNodes(uint256 number) returns(string[])
func (_Rate *RateTransactor) FindNodes(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Rate.contract.Transact(opts, "findNodes", number)
}

// FindNodes is a paid mutator transaction binding the contract method 0xfa811847.
//
// Solidity: function findNodes(uint256 number) returns(string[])
func (_Rate *RateSession) FindNodes(number *big.Int) (*types.Transaction, error) {
	return _Rate.Contract.FindNodes(&_Rate.TransactOpts, number)
}

// FindNodes is a paid mutator transaction binding the contract method 0xfa811847.
//
// Solidity: function findNodes(uint256 number) returns(string[])
func (_Rate *RateTransactorSession) FindNodes(number *big.Int) (*types.Transaction, error) {
	return _Rate.Contract.FindNodes(&_Rate.TransactOpts, number)
}

// GetCurrentRoundValidator is a paid mutator transaction binding the contract method 0xd1e2cbb9.
//
// Solidity: function getCurrentRoundValidator() returns((string[],uint256,uint256,uint256))
func (_Rate *RateTransactor) GetCurrentRoundValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rate.contract.Transact(opts, "getCurrentRoundValidator")
}

// GetCurrentRoundValidator is a paid mutator transaction binding the contract method 0xd1e2cbb9.
//
// Solidity: function getCurrentRoundValidator() returns((string[],uint256,uint256,uint256))
func (_Rate *RateSession) GetCurrentRoundValidator() (*types.Transaction, error) {
	return _Rate.Contract.GetCurrentRoundValidator(&_Rate.TransactOpts)
}

// GetCurrentRoundValidator is a paid mutator transaction binding the contract method 0xd1e2cbb9.
//
// Solidity: function getCurrentRoundValidator() returns((string[],uint256,uint256,uint256))
func (_Rate *RateTransactorSession) GetCurrentRoundValidator() (*types.Transaction, error) {
	return _Rate.Contract.GetCurrentRoundValidator(&_Rate.TransactOpts)
}

// GetLastRoundValidator is a paid mutator transaction binding the contract method 0xf866c30b.
//
// Solidity: function getLastRoundValidator() returns((string[],uint256,uint256,uint256))
func (_Rate *RateTransactor) GetLastRoundValidator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rate.contract.Transact(opts, "getLastRoundValidator")
}

// GetLastRoundValidator is a paid mutator transaction binding the contract method 0xf866c30b.
//
// Solidity: function getLastRoundValidator() returns((string[],uint256,uint256,uint256))
func (_Rate *RateSession) GetLastRoundValidator() (*types.Transaction, error) {
	return _Rate.Contract.GetLastRoundValidator(&_Rate.TransactOpts)
}

// GetLastRoundValidator is a paid mutator transaction binding the contract method 0xf866c30b.
//
// Solidity: function getLastRoundValidator() returns((string[],uint256,uint256,uint256))
func (_Rate *RateTransactorSession) GetLastRoundValidator() (*types.Transaction, error) {
	return _Rate.Contract.GetLastRoundValidator(&_Rate.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xfc8045c3.
//
// Solidity: function update(uint256 newRate, (uint64,uint64,bytes32,uint64,uint32,bytes32) qc, bytes bitmap, bytes signature, uint256 leafIndex, bytes32[] proof) returns()
func (_Rate *RateTransactor) Update(opts *bind.TransactOpts, newRate *big.Int, qc QuorumCert, bitmap []byte, signature []byte, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Rate.contract.Transact(opts, "update", newRate, qc, bitmap, signature, leafIndex, proof)
}

// Update is a paid mutator transaction binding the contract method 0xfc8045c3.
//
// Solidity: function update(uint256 newRate, (uint64,uint64,bytes32,uint64,uint32,bytes32) qc, bytes bitmap, bytes signature, uint256 leafIndex, bytes32[] proof) returns()
func (_Rate *RateSession) Update(newRate *big.Int, qc QuorumCert, bitmap []byte, signature []byte, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Rate.Contract.Update(&_Rate.TransactOpts, newRate, qc, bitmap, signature, leafIndex, proof)
}

// Update is a paid mutator transaction binding the contract method 0xfc8045c3.
//
// Solidity: function update(uint256 newRate, (uint64,uint64,bytes32,uint64,uint32,bytes32) qc, bytes bitmap, bytes signature, uint256 leafIndex, bytes32[] proof) returns()
func (_Rate *RateTransactorSession) Update(newRate *big.Int, qc QuorumCert, bitmap []byte, signature []byte, leafIndex *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Rate.Contract.Update(&_Rate.TransactOpts, newRate, qc, bitmap, signature, leafIndex, proof)
}

// RateUpdateRateIterator is returned from FilterUpdateRate and is used to iterate over the raw logs and unpacked data for UpdateRate events raised by the Rate contract.
type RateUpdateRateIterator struct {
	Event *RateUpdateRate // Event containing the contract specifics and raw log

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
func (it *RateUpdateRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RateUpdateRate)
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
		it.Event = new(RateUpdateRate)
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
func (it *RateUpdateRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RateUpdateRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RateUpdateRate represents a UpdateRate event raised by the Rate contract.
type RateUpdateRate struct {
	BlockNumber *big.Int
	Rate        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterUpdateRate is a free log retrieval operation binding the contract event 0xfe3b4c3f6e2efdc3dd612c3eb2817308b7568ad72d3f9c9200a08a069c27dff8.
//
// Solidity: event UpdateRate(uint256 blockNumber, uint256 rate)
func (_Rate *RateFilterer) FilterUpdateRate(opts *bind.FilterOpts) (*RateUpdateRateIterator, error) {

	logs, sub, err := _Rate.contract.FilterLogs(opts, "UpdateRate")
	if err != nil {
		return nil, err
	}
	return &RateUpdateRateIterator{contract: _Rate.contract, event: "UpdateRate", logs: logs, sub: sub}, nil
}

// WatchUpdateRate is a free log subscription operation binding the contract event 0xfe3b4c3f6e2efdc3dd612c3eb2817308b7568ad72d3f9c9200a08a069c27dff8.
//
// Solidity: event UpdateRate(uint256 blockNumber, uint256 rate)
func (_Rate *RateFilterer) WatchUpdateRate(opts *bind.WatchOpts, sink chan<- *RateUpdateRate) (event.Subscription, error) {

	logs, sub, err := _Rate.contract.WatchLogs(opts, "UpdateRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RateUpdateRate)
				if err := _Rate.contract.UnpackLog(event, "UpdateRate", log); err != nil {
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

// ParseUpdateRate is a log parse operation binding the contract event 0xfe3b4c3f6e2efdc3dd612c3eb2817308b7568ad72d3f9c9200a08a069c27dff8.
//
// Solidity: event UpdateRate(uint256 blockNumber, uint256 rate)
func (_Rate *RateFilterer) ParseUpdateRate(log types.Log) (*RateUpdateRate, error) {
	event := new(RateUpdateRate)
	if err := _Rate.contract.UnpackLog(event, "UpdateRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
