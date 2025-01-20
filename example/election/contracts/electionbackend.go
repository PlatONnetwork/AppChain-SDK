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
	ElectionBackendBackendABI       = "[{\"type\":\"function\",\"name\":\"changeEpoch\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Add\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Apply\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Audit\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"auditStat\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumAuditState\"},{\"name\":\"auditReason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChangeEpoch\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delete\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Update\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]"
	ElectionBackendBackendCode      = "608060405234801561001057600080fd5b50610d1e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638da5cb5b1161005b5780638da5cb5b1461010d578063900cf0cf14610135578063f2fde38b1461015d578063fe6f86f81461017057600080fd5b8063208f2a311461008d5780634851375c146100b6578063572d356e146100ef578063715018a614610103575b600080fd5b6100a061009b3660046109c8565b610178565b6040516100ad9190610ab4565b60405180910390f35b6073546100d69068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100ad565b6073546100d69067ffffffffffffffff1681565b61010b610224565b005b60335460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100ad565b6073546100d690700100000000000000000000000000000000900467ffffffffffffffff1681565b61010b61016b36600461098b565b610238565b61010b6102f4565b6072818154811061018857600080fd5b9060005260206000200160009150905080546101a390610b6b565b80601f01602080910402602001604051908101604052809291908181526020018280546101cf90610b6b565b801561021c5780601f106101f15761010080835404028352916020019161021c565b820191906000526020600020905b8154815290600101906020018083116101ff57829003601f168201915b505050505081565b61022c610546565b61023660006105c7565b565b610240610546565b73ffffffffffffffffffffffffffffffffffffffff81166102e8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b6102f1816105c7565b50565b60735467ffffffffffffffff16600061030d8243610c20565b15905080156104105760738054700100000000000000000000000000000000900467ffffffffffffffff1690601061034483610bf8565b825467ffffffffffffffff9182166101009390930a928302919092021990911617905550607380547fffffffffffffff00ffffffffffffffffffffffffffffffffffffffffffffffff167801000000000000000000000000000000000000000000000000179055606a80546066906103bf9082908490610771565b506001828101549082015560028083015490820155600391820154910155606e8054606a906103f19082908490610771565b5060018281015490820155600280830154908201556003918201549101555b60735460009083906104389068010000000000000000900467ffffffffffffffff1643610b27565b6104429190610c20565b1590508015610541576073546000906104719068010000000000000000900467ffffffffffffffff1643610b27565b61047c906001610b27565b60735490915060009085906104a79068010000000000000000900467ffffffffffffffff1643610b27565b6104b19190610b27565b604080516000808252602082019092529192506104de565b60608152602001906001900390816104c95790505b5080516104f391606e916020909101906107d7565b506104fe606e61063e565b60735461052a90700100000000000000000000000000000000900467ffffffffffffffff166001610b3f565b67ffffffffffffffff16606f556070919091556071555b505050565b60335473ffffffffffffffffffffffffffffffffffffffff163314610236576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102df565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60005b60725481101561076d57600160656072838154811061066257610662610cb9565b9060005260206000200160405161067991906109e1565b9081526040519081900360200190206007015460ff6401000000009091041660048111156106a9576106a9610c8a565b148015610708575060006065607283815481106106c8576106c8610cb9565b906000526020600020016040516106df91906109e1565b9081526040519081900360200190206003015460ff16600181111561070657610706610c8a565b145b1561075b57816000016072828154811061072457610724610cb9565b60009182526020808320845460018101865594845292209101805491909201919061074e90610b6b565b610759929190610824565b505b8061076581610bbf565b915050610641565b5050565b8280548282559060005260206000209081019282156107c75760005260206000209182015b828111156107c75782829080546107ac90610b6b565b6107b7929190610824565b5091600101919060010190610796565b506107d39291506108ab565b5090565b8280548282559060005260206000209081019282156107c7579160200282015b828111156107c757825180516108149184916020909101906108c8565b50916020019190600101906107f7565b82805461083090610b6b565b90600052602060002090601f016020900481019282610852576000855561089f565b82601f10610863578054855561089f565b8280016001018555821561089f57600052602060002091601f016020900482015b8281111561089f578254825591600101919060010190610884565b506107d392915061093c565b808211156107d35760006108bf8282610951565b506001016108ab565b8280546108d490610b6b565b90600052602060002090601f0160209004810192826108f6576000855561089f565b82601f1061090f57805160ff191683800117855561089f565b8280016001018555821561089f579182015b8281111561089f578251825591602001919060010190610921565b5b808211156107d3576000815560010161093d565b50805461095d90610b6b565b6000825580601f1061096d575050565b601f0160209004906000526020600020908101906102f1919061093c565b60006020828403121561099d57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff811681146109c157600080fd5b9392505050565b6000602082840312156109da57600080fd5b5035919050565b600080835481600182811c9150808316806109fd57607f831692505b6020808410821415610a36577f4e487b710000000000000000000000000000000000000000000000000000000086526022600452602486fd5b818015610a4a5760018114610a7957610aa6565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00861689528489019650610aa6565b60008a81526020902060005b86811015610a9e5781548b820152908501908301610a85565b505084890196505b509498975050505050505050565b600060208083528351808285015260005b81811015610ae157858101830151858201604001528201610ac5565b81811115610af3576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008219821115610b3a57610b3a610c5b565b500190565b600067ffffffffffffffff808316818516808303821115610b6257610b62610c5b565b01949350505050565b600181811c90821680610b7f57607f821691505b60208210811415610bb9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610bf157610bf1610c5b565b5060010190565b600067ffffffffffffffff80831681811415610c1657610c16610c5b565b6001019392505050565b600082610c56577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea2646970667358221220436e6df7f4fa6a1ac29ffa30336eba5f93cfc47baed1353ef8e8f9c103fac2c164736f6c63430008070033"
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

func (c *ElectionBackendBackendCaller) ChangeEpoch(ctx vmsdk.BackendCallerContext) error {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, err := ctx.Backend().GetEVMWithState(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header(), ctx.StateDB())
	if err != nil {
		return err
	}

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
