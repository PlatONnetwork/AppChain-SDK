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
	ElectionBackendBackendABI       = "[{\"type\":\"function\",\"name\":\"amount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeEpoch\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentRoundValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"distance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nextRoundValidator\",\"inputs\":[],\"outputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nodeList\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"period\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Add\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Apply\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"publicKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"blsPubKey\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Audit\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"auditStat\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumAuditState\"},{\"name\":\"auditReason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChangeEpoch\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"start\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"end\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delete\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Update\",\"inputs\":[{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"hostAddress\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rpcPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"p2pPort\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"desc\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false}]"
	ElectionBackendBackendCode      = "608060405234801561001057600080fd5b50610ac3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80638da5cb5b116100715780638da5cb5b14610174578063900cf0cf1461018f578063aa8c217c146101a9578063ef78d4fd146101c3578063f2fde38b146101d6578063fe6f86f8146101e957600080fd5b8063208f2a31146100b95780632b033ea7146100e25780633bdb31441461010f5780634851375c1461012157806348cd4cb114610153578063715018a61461016a575b600080fd5b6100cc6100c7366004610830565b6101f1565b6040516100d991906108e5565b60405180910390f35b606b54606c54606d546100f492919083565b604080519384526020840192909252908201526060016100d9565b6067546068546069546100f492919083565b60705461013b90600160801b90046001600160401b031681565b6040516001600160401b0390911681526020016100d9565b61015c606f5481565b6040519081526020016100d9565b61017261029d565b005b6033546040516001600160a01b0390911681526020016100d9565b60705461013b90600160c01b90046001600160401b031681565b60705461013b90600160401b90046001600160401b031681565b60705461013b906001600160401b031681565b6101726101e4366004610800565b6102b1565b61017261032f565b606e818154811061020157600080fd5b90600052602060002001600091509050805461021c906109ac565b80601f0160208091040260200160405190810160405280929190818152602001828054610248906109ac565b80156102955780601f1061026a57610100808354040283529160200191610295565b820191906000526020600020905b81548152906001019060200180831161027857829003601f168201915b505050505081565b6102a56104c8565b6102af6000610522565b565b6102b96104c8565b6001600160a01b0381166103235760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b61032c81610522565b50565b607054600090610351906001600160401b03600160401b82048116911661097d565b6001600160401b0316905060006103688243610a29565b15905080156103f45760708054600160c01b90046001600160401b031690601861039183610a02565b82546001600160401b039182166101009390930a9283029190920219909116179055506071805460ff19166001179055606a80546066906103d590829084906106a7565b5060018281015490820155600280830154908201556003918201549101555b607054600090839061041690600160801b90046001600160401b03164361093a565b6104209190610a29565b15905080156104c35760705460009061044990600160801b90046001600160401b03164361093a565b61045490600161093a565b607054909150600090859061047990600160801b90046001600160401b03164361093a565b610483919061093a565b905061048f606a610574565b6070546104ad90600160c01b90046001600160401b03166001610952565b6001600160401b0316606b55606c91909155606d555b505050565b6033546001600160a01b031633146102af5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161031a565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60005b606e548110156106a35760016065606e838154811061059857610598610a77565b906000526020600020016040516105af9190610849565b9081526040519081900360200190206007015460ff6401000000009091041660048111156105df576105df610a61565b14801561063e575060006065606e83815481106105fe576105fe610a77565b906000526020600020016040516106159190610849565b9081526040519081900360200190206003015460ff16600181111561063c5761063c610a61565b145b156106915781600001606e828154811061065a5761065a610a77565b600091825260208083208454600181018655948452922091018054919092019190610684906109ac565b61068f92919061070d565b505b8061069b816109e7565b915050610577565b5050565b8280548282559060005260206000209081019282156106fd5760005260206000209182015b828111156106fd5782829080546106e2906109ac565b6106ed92919061070d565b50916001019190600101906106cc565b50610709929150610794565b5090565b828054610719906109ac565b90600052602060002090601f01602090048101928261073b5760008555610788565b82601f1061074c5780548555610788565b8280016001018555821561078857600052602060002091601f016020900482015b8281111561078857825482559160010191906001019061076d565b506107099291506107b1565b808211156107095760006107a882826107c6565b50600101610794565b5b8082111561070957600081556001016107b2565b5080546107d2906109ac565b6000825580601f106107e2575050565b601f01602090049060005260206000209081019061032c91906107b1565b60006020828403121561081257600080fd5b81356001600160a01b038116811461082957600080fd5b9392505050565b60006020828403121561084257600080fd5b5035919050565b600080835481600182811c91508083168061086557607f831692505b602080841082141561088557634e487b7160e01b86526022600452602486fd5b81801561089957600181146108aa576108d7565b60ff198616895284890196506108d7565b60008a81526020902060005b868110156108cf5781548b8201529085019083016108b6565b505084890196505b509498975050505050505050565b600060208083528351808285015260005b81811015610912578581018301518582016040015282016108f6565b81811115610924576000604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561094d5761094d610a4b565b500190565b60006001600160401b0380831681851680830382111561097457610974610a4b565b01949350505050565b60006001600160401b03808316818516818304811182151516156109a3576109a3610a4b565b02949350505050565b600181811c908216806109c057607f821691505b602082108114156109e157634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156109fb576109fb610a4b565b5060010190565b60006001600160401b0380831681811415610a1f57610a1f610a4b565b6001019392505050565b600082610a4657634e487b7160e01b600052601260045260246000fd5b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052602160045260246000fd5b634e487b7160e01b600052603260045260246000fdfea2646970667358221220e009c6347c2b5329140ce30684e0b4d5450e265f135da38d3f4bd6927940b68464736f6c63430008070033"
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

func (c *ElectionBackendBackendCaller) CurrentRoundValidator(ctx vmsdk.BackendCallerContext) (*big.Int, *big.Int, *big.Int, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("currentRoundValidator")
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}
	out, err := c.abi.Unpack("currentRoundValidator", output)
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err
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

func (c *ElectionBackendBackendCaller) NextRoundValidator(ctx vmsdk.BackendCallerContext) (*big.Int, *big.Int, *big.Int, error) {
	con := vm.NewContract(vm.AccountRef(c.caller), vm.AccountRef(c.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&c.proxy, common.Hash{}, ElectionBackendDeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage(c.caller), ctx.Header())
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}
	evm.StateDB = ctx.StateDB()

	var input []byte

	input, err = c.abi.Pack("nextRoundValidator")
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}
	out, err := c.abi.Unpack("nextRoundValidator", output)
	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err
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
