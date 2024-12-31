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
	ElectionBackendBackendABI       = "[{\"type\":\"function\",\"name\":\"amount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeEpoch\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"distance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"period\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Add\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Apply\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Audit\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"auditStat\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumAuditState\"},{\"name\":\"auditReason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChangeEpoch\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delete\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Update\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]"
	ElectionBackendBackendCode      = "608060405234801561001057600080fd5b50610d00806100206000396000f3fe608060405234801561001057600080fd5b50600436106100be5760003560e01c8063900cf0cf11610076578063ef78d4fd1161005b578063ef78d4fd146101c6578063f2fde38b146101da578063fe6f86f8146101ed57600080fd5b8063900cf0cf14610176578063aa8c217c146101a657600080fd5b806348cd4cb1116100a757806348cd4cb11461012d578063715018a6146101445780638da5cb5b1461014e57600080fd5b8063208f2a31146100c35780634851375c146100ec575b600080fd5b6100d66100d136600461097a565b6101f5565b6040516100e39190610a66565b60405180910390f35b60745461011490700100000000000000000000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020016100e3565b61013660735481565b6040519081526020016100e3565b61014c6102a1565b005b60335460405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100e3565b607454610114907801000000000000000000000000000000000000000000000000900467ffffffffffffffff1681565b6074546101149068010000000000000000900467ffffffffffffffff1681565b6074546101149067ffffffffffffffff1681565b61014c6101e836600461093d565b6102b5565b61014c610371565b6072818154811061020557600080fd5b90600052602060002001600091509050805461022090610b4d565b80601f016020809104026020016040519081016040528092919081815260200182805461024c90610b4d565b80156102995780601f1061026e57610100808354040283529160200191610299565b820191906000526020600020905b81548152906001019060200180831161027c57829003601f168201915b505050505081565b6102a96105b9565b6102b3600061063a565b565b6102bd6105b9565b73ffffffffffffffffffffffffffffffffffffffff8116610365576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b61036e8161063a565b50565b6074546000906103999067ffffffffffffffff68010000000000000000820481169116610b1d565b67ffffffffffffffff16905060006103b18243610c02565b15905080156104a457607480547801000000000000000000000000000000000000000000000000900467ffffffffffffffff169060186103f083610bda565b825467ffffffffffffffff9182166101009390930a928302919092021990911617905550607580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055606a805460669061045390829084906107e4565b506001828101549082015560028083015490820155600391820154910155606e8054606a9061048590829084906107e4565b5060018281015490820155600280830154908201556003918201549101555b60745460009083906104d490700100000000000000000000000000000000900467ffffffffffffffff1643610ad9565b6104de9190610c02565b15905080156105b45760745460009061051590700100000000000000000000000000000000900467ffffffffffffffff1643610ad9565b610520906001610ad9565b607454909150600090859061055390700100000000000000000000000000000000900467ffffffffffffffff1643610ad9565b61055d9190610ad9565b9050610569606e6106b1565b60745461059d907801000000000000000000000000000000000000000000000000900467ffffffffffffffff166001610af1565b67ffffffffffffffff16606f556070919091556071555b505050565b60335473ffffffffffffffffffffffffffffffffffffffff1633146102b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161035c565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60005b6072548110156107e05760016065607283815481106106d5576106d5610c9b565b906000526020600020016040516106ec9190610993565b9081526040519081900360200190206007015460ff64010000000090910416600481111561071c5761071c610c6c565b14801561077b5750600060656072838154811061073b5761073b610c9b565b906000526020600020016040516107529190610993565b9081526040519081900360200190206003015460ff16600181111561077957610779610c6c565b145b156107ce57816000016072828154811061079757610797610c9b565b6000918252602080832084546001810186559484529220910180549190920191906107c190610b4d565b6107cc92919061084a565b505b806107d881610ba1565b9150506106b4565b5050565b82805482825590600052602060002090810192821561083a5760005260206000209182015b8281111561083a57828290805461081f90610b4d565b61082a92919061084a565b5091600101919060010190610809565b506108469291506108d1565b5090565b82805461085690610b4d565b90600052602060002090601f01602090048101928261087857600085556108c5565b82601f1061088957805485556108c5565b828001600101855582156108c557600052602060002091601f016020900482015b828111156108c55782548255916001019190600101906108aa565b506108469291506108ee565b808211156108465760006108e58282610903565b506001016108d1565b5b8082111561084657600081556001016108ef565b50805461090f90610b4d565b6000825580601f1061091f575050565b601f01602090049060005260206000209081019061036e91906108ee565b60006020828403121561094f57600080fd5b813573ffffffffffffffffffffffffffffffffffffffff8116811461097357600080fd5b9392505050565b60006020828403121561098c57600080fd5b5035919050565b600080835481600182811c9150808316806109af57607f831692505b60208084108214156109e8577f4e487b710000000000000000000000000000000000000000000000000000000086526022600452602486fd5b8180156109fc5760018114610a2b57610a58565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00861689528489019650610a58565b60008a81526020902060005b86811015610a505781548b820152908501908301610a37565b505084890196505b509498975050505050505050565b600060208083528351808285015260005b81811015610a9357858101830151858201604001528201610a77565b81811115610aa5576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60008219821115610aec57610aec610c3d565b500190565b600067ffffffffffffffff808316818516808303821115610b1457610b14610c3d565b01949350505050565b600067ffffffffffffffff80831681851681830481118215151615610b4457610b44610c3d565b02949350505050565b600181811c90821680610b6157607f821691505b60208210811415610b9b577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610bd357610bd3610c3d565b5060010190565b600067ffffffffffffffff80831681811415610bf857610bf8610c3d565b6001019392505050565b600082610c38577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea264697066735822122066adff91099b455f4bf33a37c4b310ed880727717b98c15b011b671b868534cf64736f6c63430008070033"
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

func (c *ElectionBackendBackendCaller) Amount(ctx vmsdk.BackendCallerContext) (uint64, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(uint64), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("amount")
	if err != nil {
		return *new(uint64), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(uint64), err
	}
	out, err := c.abi.Unpack("amount", output)
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err
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

func (c *ElectionBackendBackendCaller) Period(ctx vmsdk.BackendCallerContext) (uint64, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(uint64), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("period")
	if err != nil {
		return *new(uint64), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(uint64), err
	}
	out, err := c.abi.Unpack("period", output)
	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

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
