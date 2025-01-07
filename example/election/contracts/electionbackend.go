package contracts

import (
	vmsdk "github.com/PlatONnetwork/AppChain-SDK/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"math"
	"math/big"
	"strings"
	"sync"
)

var (
	ElectionBackendBackendABI       = "[{\"type\":\"function\",\"name\":\"changeEpoch\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Add\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Apply\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Audit\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"auditStat\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumAuditState\"},{\"name\":\"auditReason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChangeEpoch\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delete\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Update\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]"
	ElectionBackendBackendCode      = "608060405234801561001057600080fd5b50610c4e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a35760003560e01c8063715018a611610076578063900cf0cf1161005b578063900cf0cf14610167578063f2fde38b1461018f578063fe6f86f8146101a257600080fd5b8063715018a6146101355780638da5cb5b1461013f57600080fd5b8063208f2a31146100a85780634851375c146100d157806348cd4cb11461010a578063572d356e14610121575b600080fd5b6100bb6100b63660046108f8565b6101aa565b6040516100c891906109e4565b60405180910390f35b6074546100f19068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100c8565b61011360735481565b6040519081526020016100c8565b6074546100f19067ffffffffffffffff1681565b61013d610256565b005b60335460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100c8565b6074546100f190700100000000000000000000000000000000900467ffffffffffffffff1681565b61013d61019d3660046108bb565b61026a565b61013d610326565b607281815481106101ba57600080fd5b9060005260206000200160009150905080546101d590610a9b565b80601f016020809104026020016040519081016040528092919081815260200182805461020190610a9b565b801561024e5780601f106102235761010080835404028352916020019161024e565b820191906000526020600020905b81548152906001019060200180831161023157829003601f168201915b505050505081565b61025e610537565b61026860006105b8565b565b610272610537565b73ffffffffffffffffffffffffffffffffffffffff811661031a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b610323816105b8565b50565b60745467ffffffffffffffff16600061033f8243610b50565b15905080156104425760748054700100000000000000000000000000000000900467ffffffffffffffff1690601061037683610b28565b825467ffffffffffffffff9182166101009390930a928302919092021990911617905550607480547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000179055606a80546066906103f19082908490610762565b506001828101549082015560028083015490820155600391820154910155606e8054606a906104239082908490610762565b5060018281015490820155600280830154908201556003918201549101555b607454600090839061046a9068010000000000000000900467ffffffffffffffff1643610a57565b6104749190610b50565b1590508015610532576074546000906104a39068010000000000000000900467ffffffffffffffff1643610a57565b6104ae906001610a57565b60745490915060009085906104d99068010000000000000000900467ffffffffffffffff1643610a57565b6104e39190610a57565b90506104ef606e61062f565b60745461051b90700100000000000000000000000000000000900467ffffffffffffffff166001610a6f565b67ffffffffffffffff16606f556070919091556071555b505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610268576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610311565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60005b60725481101561075e57600160656072838154811061065357610653610be9565b9060005260206000200160405161066a9190610911565b9081526040519081900360200190206007015460ff64010000000090910416600481111561069a5761069a610bba565b1480156106f9575060006065607283815481106106b9576106b9610be9565b906000526020600020016040516106d09190610911565b9081526040519081900360200190206003015460ff1660018111156106f7576106f7610bba565b145b1561074c57816000016072828154811061071557610715610be9565b60009182526020808320845460018101865594845292209101805491909201919061073f90610a9b565b61074a9291906107c8565b505b8061075681610aef565b915050610632565b5050565b8280548282559060005260206000209081019282156107b85760005260206000209182015b828111156107b857828290805461079d90610a9b565b6107a89291906107c8565b5091600101919060010190610787565b506107c492915061084f565b5090565b8280546107d490610a9b565b90600052602060002090601f0160209004810192826107f65760008555610843565b82601f106108075780548555610843565b8280016001018555821561084357600052602060002091601f016020900482015b82811115610843578254825591600101919060010190610828565b506107c492915061086c565b808211156107c45760006108638282610881565b5060010161084f565b5b808211156107c4576000815560010161086d565b50805461088d90610a9b565b6000825580601f1061089d575050565b601f016020900490600052602060002090810190610323919061086c565b6000602082840312156108cd57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146108f157600080fd5b9392505050565b60006020828403121561090a57600080fd5b5035919050565b600080835481600182811c91508083168061092d57607f831692505b6020808410821415610966577f4e487b710000000000000000000000000000000000000000000000000000000086526022600452602486fd5b81801561097a57600181146109a9576109d6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008616895284890196506109d6565b60008a81526020902060005b868110156109ce5781548b8201529085019083016109b5565b505084890196505b509498975050505050505050565b600060208083528351808285015260005b81811015610a11578581018301518582016040015282016109f5565b81811115610a23576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008219821115610a6a57610a6a610b8b565b500190565b600067ffffffffffffffff808316818516808303821115610a9257610a92610b8b565b01949350505050565b600181811c90821680610aaf57607f821691505b60208210811415610ae9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610b2157610b21610b8b565b5060010190565b600067ffffffffffffffff80831681811415610b4657610b46610b8b565b6001019392505050565b600082610b86577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea264697066735822122079aead9dd5aa87372d5d10719223470bd2fdb18464eedddd4e14deac4e60853b64736f6c63430008070033"
	ElectionBackendDeployedCodeOnce sync.Once
	ElectionBackendDeployedCode     []byte
)

type ElectionBackendBackendCaller struct {
	abi    *abi.ABI
	proxy  common.Address
	caller common.Address
}

func NewElectionBackendBackendCaller(proxyAddress common.Address) (*ElectionBackendBackendCaller, error) {
	ElectionBackendDeployedCodeOnce.Do(func() {
		ElectionBackendDeployedCode = vmsdk.MustDeployCode(common.FromHex(ElectionBackendBackendCode), nil)
	})
	abi, err := abi.JSON(strings.NewReader(ElectionBackendBackendABI))
	if err != nil {
		return nil, err
	}
	return &ElectionBackendBackendCaller{
		abi:   &abi,
		proxy: proxyAddress,
	}, nil
}

func (c *ElectionBackendBackendCaller) Distance(ctx vmsdk.BackendCallerContext) (uint64, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(uint64), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("distance")
	if err != nil {
		return *new(uint64), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(uint64), err
	}
	out, err := c.abi.Unpack("distance", output)
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err
}

func (c *ElectionBackendBackendCaller) Epoch(ctx vmsdk.BackendCallerContext) (uint64, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(uint64), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("epoch")
	if err != nil {
		return *new(uint64), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(uint64), err
	}
	out, err := c.abi.Unpack("epoch", output)
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err
}

func (c *ElectionBackendBackendCaller) EpochSize(ctx vmsdk.BackendCallerContext) (uint64, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(uint64), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("epochSize")
	if err != nil {
		return *new(uint64), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(uint64), err
	}
	out, err := c.abi.Unpack("epochSize", output)
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err
}

func (c *ElectionBackendBackendCaller) NodeList(ctx vmsdk.BackendCallerContext, arg0 *big.Int) (string, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(string), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("nodeList", arg0)
	if err != nil {
		return *new(string), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(string), err
	}
	out, err := c.abi.Unpack("nodeList", output)
	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err
}

func (c *ElectionBackendBackendCaller) Owner(ctx vmsdk.BackendCallerContext) (common.Address, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(common.Address), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("owner")
	if err != nil {
		return *new(common.Address), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(common.Address), err
	}
	out, err := c.abi.Unpack("owner", output)
	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err
}

func (c *ElectionBackendBackendCaller) StartBlock(ctx vmsdk.BackendCallerContext) (*big.Int, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(*big.Int), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("startBlock")
	if err != nil {
		return *new(*big.Int), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(*big.Int), err
	}
	out, err := c.abi.Unpack("startBlock", output)
	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err
}

func (c *ElectionBackendBackendCaller) ChangeEpoch(ctx vmsdk.BackendCallerContext) error {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("changeEpoch")
	if err != nil {
		return err
	}

	_, err = evm.Interpreter().Run(con, input, false)

	return err
}

func (c *ElectionBackendBackendCaller) RenounceOwnership(ctx vmsdk.BackendCallerContext) error {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("renounceOwnership")
	if err != nil {
		return err
	}

	_, err = evm.Interpreter().Run(con, input, false)

	return err
}

func (c *ElectionBackendBackendCaller) TransferOwnership(ctx vmsdk.BackendCallerContext, newOwner common.Address) error {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("transferOwnership", newOwner)
	if err != nil {
		return err
	}

	_, err = evm.Interpreter().Run(con, input, false)

	return err
}

func (c *ElectionBackendBackendCaller) ABI() *abi.ABI {
	return c.abi
}
func (c *ElectionBackendBackendCaller) WithCaller(caller common.Address) *ElectionBackendBackendCaller {
	c.caller = caller
	return c
}
