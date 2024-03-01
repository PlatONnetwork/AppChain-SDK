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
tool contract new --abi abi.json --output . --type test --pkg test
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
	if err := c.loadMethodABI(); err != nil {
		return nil, errors.New("load version failed")
	}

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

func (c *Test) initABI() {
	V0 := uint16(0)
	c.abis[V0] = &Abi
}

func (c *Test) initMethodEntry() {

	methodEntry := map[string]func([]byte) ([]byte, error){

		"b1976a02": c.GetEntry,
		"03b067c1": c.SetEntry,
	}
	V0 := uint16(0)
	c.methodEntries[V0] = methodEntry

}
func (c *Test) loadMethodABI() error {
	version := c.GetVersion()
	entries, ok := c.methodEntries[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.methodEntry = entries
	abi, ok := c.abis[version]
	if !ok {
		return errors.New("unknown version")
	}
	c.abi = abi
	return nil
}
func (c *Test) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	c.evm.StateDB.SetState(c.contract.Address(), createBlockKey, data[:])
}

func (c *Test) GetCreateBlock() uint64 {
	blockNumber := c.evm.StateDB.GetState(c.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func (c *Test) SetVersion(version uint16) {
	var data [2]byte
	binary.BigEndian.PutUint16(data[:], version)
	c.evm.StateDB.SetState(c.contract.Address(), versionKey, data[:])
}

func (c *Test) GetVersion() uint16 {
	version := c.evm.StateDB.GetState(c.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint16(version)
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
	abi           *abi.ABI
	abis          map[uint16]*abi.ABI
	methodEntry   map[string]func([]byte) ([]byte, error)
	methodEntries map[uint16]map[string]func([]byte) ([]byte, error)
	readOnly      bool
	contract      *vm.Contract
	evm           *vm.EVM
	burner        contracts.Burn
	stateDb       *contracts.StateDB
	fallback      func(input []byte) ([]byte, error)
}

func NewTest(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*Test, error) {
	s := &Test{
		abi:           nil,
		abis:          make(map[uint16]*abi.ABI),
		methodEntry:   make(map[string]func([]byte) ([]byte, error)),
		methodEntries: make(map[uint16]map[string]func([]byte) ([]byte, error)),
		evm:           evm,
		contract:      contract,
		burner:        contracts.NewBurner(contract),
		stateDb:       contracts.NewStateDB(evm, contract),
		readOnly:      readOnly,
	}
	s.initABI()
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

代码定义了其它模块调用该合约的框架代码，调用分为 Call，DelegateCall 调用

## 升级

合约升级增加 结构体`struct A` 、事件 `DelNum`、函数 `Del`，升级 `Set` 函数

```solidity
struct A{
    int a;
}
interface Test {

    event SetNum(int);
    event DelNum(int);
    function Set(int num) external;
    function Get()external returns (int);
    function Del(int num) external;
}
```

执行生成新代码命令

```shell
tool contract upgrade --abi  abi2.json --output . --pkg test --version 1 --upgrade-method "Set" --new-method "Del" --new-struct "A"  --new-event DelNum
```
* version 指定升级的版本
* upgrade-method 升级的函数
* new-method 新增的函数
* new-struct 新增的结构体
* new-event 新增的事件


生成新的 ABI 初始化函数，新增的结构体，新的Entry函数，新的事件触发函数

```go
// A is an auto generated low-level Go binding around an user-defined struct.
type A struct {
	A *big.Int
}

func (c *Test) initABIV1() {
	V1 := uint16(1)
	c.abis[V1] = &AbiV1
}

func (c *Test) initMethodV1Entry() {
	methodEntry := map[string]func([]byte) ([]byte, error){

		"03b067c1": c.SetV1Entry,
		"2125afe1": c.DelEntry,
		"b1976a02": c.GetEntry,
	}
	V1 := uint16(1)
	c.methodEntries[V1] = methodEntry
}

func (c *Test) DelEntry(input []byte) ([]byte, error) {

	method := c.abi.Methods["Del"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.Del(*abi.ConvertType(args[0], new(A)).(*A))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
func (c *Test) Del(num A) error {
	panic("implement")
}

func (c *Test) SetV1Entry(input []byte) ([]byte, error) {

	method := c.abi.Methods["Set"]

	var err error

	args, err := method.Inputs.Unpack(input)
	if err != nil {
		return nil, err
	}

	err = c.SetV1(*abi.ConvertType(args[0], new(*big.Int)).(**big.Int))
	if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
		return nil, err
	}
	var output []byte

	return output, err
}
func (c *Test) SetV1(num *big.Int) error {
	panic("implement")
}

func (c *Test) EmitDelNumEvent(arg0 *big.Int) (*types.Log, error) {
	event := c.abi.Events["DelNum"]
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