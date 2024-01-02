# 合约工具

合约工具是生成合约模块工具，用来将 Solidity 合约接口生成 Go 代码，abigen 利用 PlatON 中 ABI 的解析，利用 Go 模板定义合约模块的基础骨架。主要有函数参数及合约参数结构体定义，合约入口函数，事件触发，函数参数编码解析


## 代码框架

使用合约工具将生成3个合约代码文件，contract.go，contract_impl.go，contract_caller.go，下面看看各个文件生成的代码结构

根据以下Solidity 合约生成代码

```solidity
interface Test {
    event SetNum(int);
    function Set(int num) external;
    function Get()external returns (int);
}
```

执行命令

```shell
tool contract --abi abi.json --output . --type test --pkg test
```

* test.go

```go
func (c *Test) Run(input []byte) (ret []byte, err error) {
	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case error:
				if r, ok := e.(*typesdk.RevertError); ok {
					ret, err = r.ReturnData, vm.ErrExecutionReverted
				} else {
					ret, err = nil, e
				}
			default:
				ret, err = typesdk.UndefinedError, vm.ErrExecutionReverted
			}
		}
	}()
	if len(input) < 4 {
		return nil, errors.New("input too short")
	}
	id := input[0:4]
	entry, ok := c.methodEntry[hex.EncodeToString(id)]
	if !ok {
		if c.fallback != nil {
			return c.fallback(input)
		}
		return nil, errors.New("methods not found")
	}
	return entry(input[4:])
}
func (c *Test) initMethodEntry() {

	c.methodEntry = map[string]func([]byte) ([]byte, error){

		"b1976a02": c.GetEntry,
		"03b067c1": c.SetEntry,
	}

}

func (c *Test) GetEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["Get"]

	var err error

	res0, err := c.Get()
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	output, err = method.Outputs.Pack(res0)
	if err != nil {
		return nil, err
	}

	return output, err
}

func (c *Test) SetEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["Set"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Set(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}

func (c *Test) EmitSetNumEvent(arg0 *big.Int) (*types.Log, error) {
	event := c.abi.Events["SetNum"]
	hashes, err := contracts.PackEventTopics(event.ID, event.Inputs, arg0)
	if err != nil {
		return nil, err
	}
	data, err := contracts.PackEventData(event.Inputs, arg0)
	if err != nil {
		return nil, err
	}
	return &types.Log{
		Address:     c.contract.Address(),
		Topics:      hashes,
		Data:        data,
		BlockNumber: c.evm.Context.BlockNumber.Uint64(),
	}, nil
}
```

代码实现了调用框架代码，代码结构不会被改动，其中 Run 函数为合约的入口函数，Entry 函数则解析具体函数后调用具体实现函数，EmitEvent 用于构建合约事件

* test_impl.go

```go
type Test struct {
	abi         *abi.ABI
	methodEntry map[string]func([]byte) ([]byte, error)
	readOnly    bool
	contract    *vm.Contract
	evm         *vm.EVM
	burner      contracts.Burn
	stateDb     *contracts.StateDB
	fallback    func(input []byte) ([]byte, error)
}

func NewTest(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Test, error) {
	s := &Test{
		abi:      &Abi,
		evm:      evm,
		contract: contract,
		burner:   contracts.NewBurner(contract),
		stateDb:  contracts.NewStateDB(evm, contract),
		readOnly: readOnly,
	}
	s.initMethodEntry()
	return s, nil
}

func (c *Test) Get() (*big.Int, error) {
	panic("implement")
}

func (c *Test) Set(num *big.Int) error {
	panic("implement")
}
```

代码定义了合约的结构、构造方法，合约函数的实现

* contract_caller.go

```go
type TestCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewTestCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*TestCaller, error) {
	s := &TestCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *TestCaller) Get() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "Get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *TestCaller) Set(num *big.Int) error {
	var out []interface{}
	err := c.BoundContract.Caller(c.to, &out, "Set", num)

	if err != nil {
		return err
	}

	return err

}

type TestDelegateCaller struct {
	contracts.BoundContract
	to common.Address
}

func NewTestDelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*TestDelegateCaller, error) {
	s := &TestDelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
	}
	return s, nil
}

func (c *TestDelegateCaller) Get() (*big.Int, error) {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "Get")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

func (c *TestDelegateCaller) Set(num *big.Int) error {
	var out []interface{}
	err := c.BoundContract.DelegateCaller(c.to, &out, "Set", num)

	if err != nil {
		return err
	}

	return err

}
```

代码定义了其它模块调用该合约的框架代码，调用分为 Call，DelagateCall 调用