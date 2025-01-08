package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
)

// tmplMethod is a wrapper around an abi.Method that contains a few preprocessed
// and cached data fields.
type tmplMethod struct {
	Original   abi.Method // Original method as parsed by the abi package
	Normalized abi.Method // Normalized version of the parsed method (capitalized names, non-anonymous args/returns)
	Structured bool       // Whether the returns should be accumulated into a struct
}

// tmplEvent is a wrapper around an abi.Event that contains a few preprocessed
// and cached data fields.
type tmplEvent struct {
	Original   abi.Event // Original event as parsed by the abi package
	Normalized abi.Event // Normalized version of the parsed fields
}

// tmplField is a wrapper around a struct field with binding language
// struct type definition and relative filed name.
type tmplField struct {
	Type    string   // Field type representation depends on target binding language
	Name    string   // Field name converted from the raw user-defined field name
	SolKind abi.Type // Raw abi type information
}

// tmplStruct is a wrapper around an abi.tuple and contains an auto-generated
// struct name.
type tmplStruct struct {
	Name   string       // Auto-generated struct name(before solidity v0.5.11) or raw name.
	Fields []*tmplField // Struct fields definition depends on the binding language.
}
type tmplContract struct {
	Type        string                 // Type name of the main contract binding
	InputABI    string                 // JSON ABI used as the input to generate the binding from
	InputBin    string                 // Optional EVM bytecode used to generate deploy code from
	FuncSigs    map[string]string      // Optional map: string signature -> 4-byte signature
	Constructor abi.Method             // Contract constructor for deploy parametrization
	Calls       map[string]*tmplMethod // Contract calls that only read state data
	Transacts   map[string]*tmplMethod // Contract calls that write state data
	Events      map[string]*tmplEvent  // Contract events accessors
}

type tmplUpgradeData struct {
	*tmplData
	Version    uint64
	Entries    map[string]string
	NewStructs map[string]*tmplStruct
}

// tmplData is the data structure required to fill the binding template.
type tmplData struct {
	ReceiverName string
	NoStruct     bool
	Package      string                 // Name of the package to place the generated file in
	Contract     *tmplContract          // List of contracts to generate into this file
	Structs      map[string]*tmplStruct // Contract struct type definitions
}

const tmplFrameSource = `
package {{.Package}}

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = contracts.Context{}
    _ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = math.ReadBits
	_ = binary.BigEndian
	_ = types.BloomLookup
	_ = event.NewSubscription
	versionKey = []byte("__version")
	createBlockKey = []byte("__createBlock")
)

{{$ReceiverName := .ReceiverName}}
{{$structs := .Structs}}
{{if eq .NoStruct false }}
{{range $structs}}
	// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
	type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
	}
{{end}}
{{end}}
{{$contract := .Contract}}
var (
    ABI = "{{$contract.InputABI}}"
    Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func ({{$ReceiverName}} *{{$contract.Type}})Run(input []byte) (ret []byte, err error) {
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
	if err := {{$ReceiverName}}.loadMethodABI(); err != nil {
		return nil, errors.New("load version failed")
	}

    if len(input) < 4 {
        return nil, errors.New("input too short")
    }
    id := input[0:4]
    entry, ok := {{$ReceiverName}}.methodEntry[hex.EncodeToString(id)]
    if !ok {
		if {{$ReceiverName}}.fallback != nil {
			return {{$ReceiverName}}.fallback(input)
		}
        return nil, errors.New("methods not found")
    }
    return entry(input[4:])
}

func ({{$ReceiverName}} *{{$contract.Type}}) initABI() {
	V0 := uint64(0)
	{{$ReceiverName}}.abis[V0] = &Abi
}

func ({{$ReceiverName}} *{{$contract.Type}}) initMethodEntry() {

    methodEntry := map[string]func([]byte) ([]byte, error){
        {{range .Contract.Calls}}"{{hexid .Original.ID}}" : {{$ReceiverName}}.{{.Normalized.Name}}Entry,
        {{end}}
        {{range .Contract.Transacts}}
        "{{hexid .Original.ID}}" : {{$ReceiverName}}.{{.Normalized.Name}}Entry,{{end}}
    }
	V0 := uint64(0)
	{{$ReceiverName}}.methodEntries[V0] = methodEntry

}
func ({{$ReceiverName}} *{{$contract.Type}}) loadMethodABI() error {
	version := {{$ReceiverName}}.GetVersion()
	entries, ok := {{$ReceiverName}}.methodEntries[version]
	if !ok {
		return errors.New("unknown version")
	}
	{{$ReceiverName}}.methodEntry = entries
	abi, ok := {{$ReceiverName}}.abis[version]
	if !ok {
		return errors.New("unknown version")
	}
	{{$ReceiverName}}.abi = abi
	return nil
}

func ({{$ReceiverName}} *{{$contract.Type}}) InitGenesis(blockNumber uint64) {
	{{$ReceiverName}}.stateDb.SetNonce({{$ReceiverName}}.contract.Address(), 1)
	{{$ReceiverName}}.SetCreateBlock(blockNumber)
	{{$ReceiverName}}.stateDb.SetCode({{$ReceiverName}}.contract.Address(), []byte("code"))
}

func ({{$ReceiverName}} *{{$contract.Type}}) SetCreateBlock(blockNumber uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], blockNumber)
	{{$ReceiverName}}.evm.StateDB.SetState({{$ReceiverName}}.contract.Address(), createBlockKey, data[:])
}

func ({{$ReceiverName}} *{{$contract.Type}}) GetCreateBlock() uint64 {
	blockNumber := {{$ReceiverName}}.evm.StateDB.GetState({{$ReceiverName}}.contract.Address(), createBlockKey)
	if len(blockNumber) == 0 {
		return math.MaxUint64
	}
	return binary.BigEndian.Uint64(blockNumber)
}

func ({{$ReceiverName}} *{{$contract.Type}}) SetVersion(version uint64) {
	var data [8]byte
	binary.BigEndian.PutUint64(data[:], version)
	{{$ReceiverName}}.evm.StateDB.SetState({{$ReceiverName}}.contract.Address(), versionKey, data[:])
}

func ({{$ReceiverName}} *{{$contract.Type}}) GetVersion() uint64 {
	version := {{$ReceiverName}}.evm.StateDB.GetState({{$ReceiverName}}.contract.Address(), versionKey)
	if len(version) == 0 {
		return 0
	}
	return binary.BigEndian.Uint64(version)
}

{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}Entry(input []byte) ([]byte, error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    {{if or (ne $inputLen 0) (ne $outputLen 0) }}
    method := {{$ReceiverName}}.abi.Methods["{{.Original.Name}}"]
    {{end}}
    var err error
    {{ $length := len .Normalized.Inputs }}
    {{if ne $length 0 }}
    args, err := method.Inputs.Unpack(input)
    if err != nil {
        return nil, err
    }
    {{end}}
    {{range $i, $_ := .Normalized.Outputs}}res{{$i}}, {{end}} err {{if ne $outputLen 0 }} := {{else}} ={{end}} {{$ReceiverName}}.{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} *abi.ConvertType(args[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}) {{end}})
    if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
        return nil, err
    }
    var output []byte
    {{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
    output, err = method.Outputs.Pack({{range $i, $_ := .Normalized.Outputs}}{{if ne $i 0}},{{end}}res{{$i}}{{end}})
    if err != nil {
        return nil, err
    }
    {{end}}
    return output, err
}
{{end}}

{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}Entry(input []byte) ([]byte, error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    {{if or (ne $inputLen 0) (ne $outputLen 0) }}
    method := {{$ReceiverName}}.abi.Methods["{{.Original.Name}}"]
    {{end}}
    var err error

    {{ $length := len .Normalized.Inputs }}
    {{if ne $length 0 }}
    args, err := method.Inputs.Unpack(input)
    if err != nil {
        return nil, err
    }
    {{end}}
    {{range $i, $_ := .Normalized.Outputs}}res{{$i}}, {{end}} err {{if ne $outputLen 0 }} := {{else}} ={{end}} {{$ReceiverName}}.{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} *abi.ConvertType(args[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}) {{end}})
    if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
        return nil, err
    }
    var output []byte
    {{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
    output, err = method.Outputs.Pack({{range $i, $_ := .Normalized.Outputs}}{{if ne $i 0}},{{end}}res{{$i}}{{end}})
    if err != nil {
        return nil, err
    }
    {{end}}
    return output, err
}
{{end}}
{{range .Contract.Events}}
    {{ $length := len .Normalized.Inputs }}
        func ({{$ReceiverName}} *{{$contract.Type}}){{.Normalized.Name}}Event({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}} {{bindtype .Type $structs}}{{end}}) (*types.Log, error){
        event := {{$ReceiverName}}.abi.Events["{{.Normalized.Name}}"]
        hashes, err := contracts.PackEventTopics(event.ID, event.Inputs {{range $i, $_ := .Normalized.Inputs}},{{.Name}}{{end}})
        if err != nil {
            return nil, err
        }
        data, err := contracts.PackEventData(event.Inputs {{range $i, $_ := .Normalized.Inputs}},{{.Name}}{{end}})
        if err != nil {
            return nil, err
        }
        return &types.Log{
            Address: {{$ReceiverName}}.contract.Address(),
            Topics: hashes,
            Data:   data,
            BlockNumber: {{$ReceiverName}}.evm.Context.BlockNumber.Uint64(),
        }, nil
        }
		func ({{$ReceiverName}} *{{$contract.Type}})Emit{{.Normalized.Name}}Event({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}} {{bindtype .Type $structs}}{{end}}){
			log, err := {{$ReceiverName}}.{{.Normalized.Name}}Event({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
			contracts.Require(err == nil, "{{$contract.Type}}: emit {{.Normalized.Name}} event failed")
			{{$ReceiverName}}.stateDb.AddLog(log)
		}
{{end}}
`
const implSource = `
package {{.Package}}

import (
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = typesdk.RevertError{}
    _ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)
{{$contract := .Contract}}
{{$structs := .Structs}}
{{$ReceiverName := .ReceiverName}}
type {{$contract.Type}} struct {
    abi *abi.ABI
    abis          map[uint64]*abi.ABI
	methodEntry   map[string]func([]byte) ([]byte, error)
	methodEntries map[uint64]map[string]func([]byte) ([]byte, error)
    readOnly bool
    contract *vm.Contract
    evm *vm.EVM
	burner       contracts.Burn
	stateDb      *contracts.StateDB
	context       *contracts.Context
	fallback func(input []byte) ([]byte, error)
}

func New{{$contract.Type}}(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*{{$contract.Type}}, error) {
    s := &{{$contract.Type}}{
		abi:           nil,
        abis:          make(map[uint64]*abi.ABI),
		methodEntry:   make(map[string]func([]byte) ([]byte, error)),
		methodEntries: make(map[uint64]map[string]func([]byte) ([]byte, error)),
        evm:evm,
        contract: contract,
		burner:   contracts.NewBurner(contract),
		stateDb:  contracts.NewStateDB(evm, contract),
		context:  contracts.NewContext(evm, contract),
        readOnly: readOnly,
    }
	s.initABI()
    s.initMethodEntry()
	s.loadMethodABI()
    return s, nil
}

{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    panic("implement")
}
{{end}}


{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    panic("implement")
}
{{end}}
`

const tmplCaller = `
package {{.Package}}

import (
	"errors"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/tools/contracts"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = abi.ABI{}
	_ = typesdk.RevertError{}
    _ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)
{{$contract := .Contract}}
{{$structs := .Structs}}
{{$ReceiverName := .ReceiverName}}
type {{$contract.Type}}Caller struct {
    contracts.BoundContract
	to common.Address
}
func New{{$contract.Type}}Caller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*{{$contract.Type}}Caller, error) {
    s := &{{$contract.Type}}Caller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to: to,
    }
    return s, nil
}



{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}Caller) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    var out []interface{}
	err := {{$ReceiverName}}.BoundContract.Caller({{$ReceiverName}}.to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
	{{if .Structured}}
	outstruct := new(struct{ {{range .Normalized.Outputs}} {{.Name}} {{bindtype .Type $structs}}; {{end}} })
	if err != nil {
		return *outstruct, err
	}
	{{range $i, $t := .Normalized.Outputs}} 
	outstruct.{{.Name}} = *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}

	return *outstruct, err
	{{else}}
	if err != nil {
		return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
	}
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
	{{end}}
}
{{end}}


{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}Caller) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
        var out []interface{}
	err := {{$ReceiverName}}.BoundContract.Caller({{$ReceiverName}}.to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
	{{if .Structured}}
	outstruct := new(struct{ {{range .Normalized.Outputs}} {{.Name}} {{bindtype .Type $structs}}; {{end}} })
	if err != nil {
		return *outstruct, err
	}
	{{range $i, $t := .Normalized.Outputs}} 
	outstruct.{{.Name}} = *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}

	return *outstruct, err
	{{else}}
	if err != nil {
		return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
	}
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
	{{end}}
}
{{end}}

type {{$contract.Type}}DelegateCaller struct {
    contracts.BoundContract
	to common.Address
}
func New{{$contract.Type}}DelegateCaller(evm *vm.EVM, contract *vm.Contract, to common.Address) (*{{$contract.Type}}DelegateCaller, error) {
    s := &{{$contract.Type}}DelegateCaller{
		BoundContract: contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
		to : to,
    }
    return s, nil
}



{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}DelegateCaller) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    var out []interface{}
	err := {{$ReceiverName}}.BoundContract.DelegateCaller({{$ReceiverName}}.to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
	{{if .Structured}}
	outstruct := new(struct{ {{range .Normalized.Outputs}} {{.Name}} {{bindtype .Type $structs}}; {{end}} })
	if err != nil {
		return *outstruct, err
	}
	{{range $i, $t := .Normalized.Outputs}} 
	outstruct.{{.Name}} = *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}

	return *outstruct, err
	{{else}}
	if err != nil {
		return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
	}
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
	{{end}}
}
{{end}}


{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}DelegateCaller) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
        var out []interface{}
	err := {{$ReceiverName}}.BoundContract.DelegateCaller({{$ReceiverName}}.to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
	{{if .Structured}}
	outstruct := new(struct{ {{range .Normalized.Outputs}} {{.Name}} {{bindtype .Type $structs}}; {{end}} })
	if err != nil {
		return *outstruct, err
	}
	{{range $i, $t := .Normalized.Outputs}} 
	outstruct.{{.Name}} = *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}

	return *outstruct, err
	{{else}}
	if err != nil {
		return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
	}
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
	{{end}}
}
{{end}}
`

const tmplUpgrade = `
package {{.Package}}

import (
	"encoding/hex"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"math/big"
	"strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = hex.ErrLength
	_ = contracts.Context{}
    _ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)
{{$ReceiverName := .ReceiverName}}
{{$contract := .Contract}}
var (
    ABIV{{.Version}} = "{{$contract.InputABI}}"
    AbiV{{.Version}}, _ = abi.JSON(strings.NewReader(ABIV{{.Version}}))
)

{{$structs := .Structs}}
{{range .NewStructs}}
	// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
	type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
	}
{{end}}

{{$contract := .Contract}}
func ({{$ReceiverName}} *{{$contract.Type}}) initABIV{{.Version}}() {
	V{{.Version}} := uint64({{.Version}})
	{{$ReceiverName}}.abis[V{{.Version}}] = &AbiV{{.Version}}
}

func ({{$ReceiverName}} *{{$contract.Type}}) initMethodV{{.Version}}Entry() {
    methodEntry := map[string]func([]byte) ([]byte, error){
        {{range $id, $name := .Entries}}
        "{{$id}}" : {{$ReceiverName}}.{{$name}}Entry,{{end}}
    }
	V{{.Version}} := uint64({{.Version}})
	{{$ReceiverName}}.methodEntries[V{{.Version}}] = methodEntry
}
{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}Entry(input []byte) ([]byte, error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    {{if or (ne $inputLen 0) (ne $outputLen 0) }}
    method := {{$ReceiverName}}.abi.Methods["{{.Original.Name}}"]
    {{end}}
    var err error
    {{ $length := len .Normalized.Inputs }}
    {{if ne $length 0 }}
    args, err := method.Inputs.Unpack(input)
    if err != nil {
        return nil, err
    }
    {{end}}
    {{range $i, $_ := .Normalized.Outputs}}res{{$i}}, {{end}} err {{if ne $outputLen 0 }} := {{else}} ={{end}} {{$ReceiverName}}.{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} *abi.ConvertType(args[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}) {{end}})
    if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
        return nil, err
    }
    var output []byte
    {{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
    output, err = method.Outputs.Pack({{range $i, $_ := .Normalized.Outputs}}{{if ne $i 0}},{{end}}res{{$i}}{{end}})
    if err != nil {
        return nil, err
    }
    {{end}}
    return output, err
}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    panic("implement")
}
{{end}}

{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}Entry(input []byte) ([]byte, error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    {{if or (ne $inputLen 0) (ne $outputLen 0) }}
    method := {{$ReceiverName}}.abi.Methods["{{.Original.Name}}"]
    {{end}}
    var err error

    {{ $length := len .Normalized.Inputs }}
    {{if ne $length 0 }}
    args, err := method.Inputs.Unpack(input)
    if err != nil {
        return nil, err
    }
    {{end}}
    {{range $i, $_ := .Normalized.Outputs}}res{{$i}}, {{end}} err {{if ne $outputLen 0 }} := {{else}} ={{end}} {{$ReceiverName}}.{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} *abi.ConvertType(args[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}) {{end}})
    if err != nil {
		if r, ok := err.(*typesdk.RevertError); ok {
			return r.ReturnData, vm.ErrExecutionReverted
		}
        return nil, err
    }
    var output []byte
    {{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
    output, err = method.Outputs.Pack({{range $i, $_ := .Normalized.Outputs}}{{if ne $i 0}},{{end}}res{{$i}}{{end}})
    if err != nil {
        return nil, err
    }
    {{end}}
    return output, err
}
func ({{$ReceiverName}} *{{$contract.Type}}) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    panic("implement")
}
{{end}}
{{range .Contract.Events}}
    {{ $length := len .Normalized.Inputs }}
        func ({{$ReceiverName}} *{{$contract.Type}}){{.Normalized.Name}}Event({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}} {{bindtype .Type $structs}}{{end}}) (*types.Log, error){
        event := {{$ReceiverName}}.abi.Events["{{.Normalized.Name}}"]
        hashes, err := contracts.PackEventTopics(event.ID, event.Inputs {{range $i, $_ := .Normalized.Inputs}},{{.Name}}{{end}})
        if err != nil {
            return nil, err
        }
        data, err := contracts.PackEventData(event.Inputs {{range $i, $_ := .Normalized.Inputs}},{{.Name}}{{end}})
        if err != nil {
            return nil, err
        }
        return &types.Log{
            Address: {{$ReceiverName}}.contract.Address(),
            Topics: hashes,
            Data:   data,
            BlockNumber: {{$ReceiverName}}.evm.Context.BlockNumber.Uint64(),
        }, nil
        }
		func ({{$ReceiverName}} *{{$contract.Type}})Emit{{.Normalized.Name}}Event({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}} {{bindtype .Type $structs}}{{end}}){
			log, err := {{$ReceiverName}}.{{.Normalized.Name}}Event({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
			contracts.Require(err == nil, "{{$contract.Type}}: emit {{.Normalized.Name}} event failed")
			{{$ReceiverName}}.stateDb.AddLog(log)
		}
{{end}}
`

const tmplBackendCaller = `
package {{.Package}}
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

{{$ReceiverName := .ReceiverName}}
{{$contract := .Contract}}
var (
    {{$contract.Type}}BackendABI = "{{$contract.InputABI}}"
	{{$contract.Type}}BackendCode = "{{$contract.InputBin}}"
	{{$contract.Type}}DeployedCodeOnce sync.Once
	{{$contract.Type}}DeployedCode     []byte
)

{{$contract := .Contract}}
{{$structs := .Structs}}
{{$ReceiverName := .ReceiverName}}
type {{$contract.Type}}BackendCaller struct {
	abi    *abi.ABI
	proxy  common.Address
	caller common.Address
}

func New{{$contract.Type}}BackendCaller(proxyAddress common.Address) (*{{$contract.Type}}BackendCaller, error) {
	{{$contract.Type}}DeployedCodeOnce.Do(func() {
		{{$contract.Type}}DeployedCode = vmsdk.MustDeployCode(common.FromHex({{$contract.Type}}BackendCode), nil)
	})
	abi, err := abi.JSON(strings.NewReader({{$contract.Type}}BackendABI))
	if err != nil {
		return nil, err
	}
	return &{{$contract.Type}}BackendCaller{
		abi: &abi,
		proxy: proxyAddress,
	}, nil
}

{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}BackendCaller) {{.Normalized.Name}}(ctx vmsdk.BackendCallerContext {{range $i, $_ := .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    con := vm.NewContract(vm.AccountRef({{$ReceiverName}}.caller), vm.AccountRef({{$ReceiverName}}.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&{{$ReceiverName}}.proxy, common.Hash{}, {{$contract.Type}}DeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage({{$ReceiverName}}.caller), ctx.Header())
	if err != nil {
		return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
	}
	evm.StateDB = ctx.StateDB()

    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
	var input []byte
	
    input, err = {{$ReceiverName}}.abi.Pack("{{.Original.Name}}",{{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }

    {{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	out, err := c.abi.Unpack("{{.Original.Name}}", output)
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	{{else}}
	_, err = evm.Interpreter().Run(con, input, false)
	{{end}}
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
}
{{end}}

{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}BackendCaller) {{.Normalized.Name}}(ctx vmsdk.BackendCallerContext {{range $i, $_ := .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    con := vm.NewContract(vm.AccountRef({{$ReceiverName}}.caller), vm.AccountRef({{$ReceiverName}}.proxy), big.NewInt(0), math.MaxUint64)
	con.SetCallCode(&{{$ReceiverName}}.proxy, common.Hash{}, {{$contract.Type}}DeployedCode)
	evm, _, err := ctx.Backend().GetEVM(vmsdk.NewOnlyCallMessage({{$ReceiverName}}.caller), ctx.Header())
	if err != nil {
		return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
	}
	evm.StateDB = ctx.StateDB()

    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
	var input []byte
	
    input, err = {{$ReceiverName}}.abi.Pack("{{.Original.Name}}",{{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }

    {{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
	var output []byte
	output, err = evm.Interpreter().Run(con, input, false)
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	out, err := c.abi.Unpack("{{.Original.Name}}", output)
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	{{else}}
	_, err = evm.Interpreter().Run(con, input, false)
	{{end}}
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
}
{{end}}
func ({{$ReceiverName}} *{{$contract.Type}}BackendCaller) ABI() *abi.ABI {
	return {{$ReceiverName}}.abi
}
func ({{$ReceiverName}} *{{$contract.Type}}BackendCaller) WithCaller(caller common.Address) *{{$contract.Type}}BackendCaller {
	{{$ReceiverName}}.caller = caller
	return {{$ReceiverName}}
}
`

const tmplGenesis = `
package {{.Package}}
import (
	vm2 "github.com/PlatONnetwork/AppChain-SDK/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math"
	"math/big"
	"strings"
)

{{if eq .NoStruct false }}
	{{$ReceiverName := .ReceiverName}}
	{{$structs := .Structs}}
	{{range $structs}}
		
		// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
		type {{.Name}} struct {
		{{range $field := .Fields}}
		{{$field.Name}} {{$field.Type}}{{end}}
		}
	{{end}}
{{end}}

{{$ReceiverName := .ReceiverName}}
{{$contract := .Contract}}
var (
    {{$contract.Type}}GenesisABI = "{{$contract.InputABI}}"
	{{$contract.Type}}Code = "{{$contract.InputBin}}"
)

{{$contract := .Contract}}
{{$structs := .Structs}}
{{$ReceiverName := .ReceiverName}}
type {{$contract.Type}}GenesisCaller struct {
   	evmFunc func(address common.Address) *vm.EVM
	abi     *abi.ABI
	to      common.Address
	caller  common.Address
}
func New{{$contract.Type}}GenesisCaller(ctx sdk.Context, db sdk.StateDB, chainConfig *params.ChainConfig) (*{{$contract.Type}}GenesisCaller, error) {
	abi, err := abi.JSON(strings.NewReader({{$contract.Type}}GenesisABI))
	if err != nil {
		return nil, err
	}
   	return &{{$contract.Type}}GenesisCaller{
		evmFunc: func(address common.Address) *vm.EVM {
			return vm2.NewEVM(vm2.NewGenesisBlockContext(), address, db, chainConfig, nil)
		},
		abi: &abi,
	}, nil
}

func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) Deploy{{$contract.Type}}({{range $i, $_ := $contract.Constructor.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) error {
    {{ $length := len $contract.Constructor.Inputs }}
	var data []byte
	var err error
    {{if ne $length 0 }}
    data, err = {{$ReceiverName}}.abi.Constructor.Inputs.Pack({{range $i, $_ := $contract.Constructor.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
        return err
    }
    {{end}}

	evm := {{$ReceiverName}}.evmFunc({{$ReceiverName}}.caller)
	_, _, _, err = evm.CreateAppContract(vm.AccountRef({{$ReceiverName}}.caller), append(common.FromHex({{$contract.Type}}Code), data...), {{$ReceiverName}}.to, math.MaxUint64, big.NewInt(0))
	return err
}

{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ([]byte,error) {
    input, err := {{$ReceiverName}}.abi.Pack("{{.Original.Name}}",{{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
		return nil, err
	}
	return input, nil
}
func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
   	evm := c.evmFunc(c.caller)
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    var err error
	var input []byte
    input, err = {{$ReceiverName}}.Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	{{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	out, err := c.abi.Unpack("{{.Original.Name}}", output)
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	{{else}}
	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	{{end}}
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
}
{{end}}

{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ([]byte,error) {
    input, err := {{$ReceiverName}}.abi.Pack("{{.Original.Name}}",{{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
		return nil, err
	}
	return input, nil
}
func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
   	evm := c.evmFunc(c.caller)
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    var err error
	var input []byte
	
    input, err = {{$ReceiverName}}.Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
    {{ $length := len .Normalized.Outputs }}
    {{if ne $length 0 }}
	var output []byte
	output, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	out, err := c.abi.Unpack("{{.Original.Name}}", output)
    if err != nil {
        return {{range $i, $_ := .Normalized.Outputs}}*new({{bindtype .Type $structs}}), {{end}} err
    }
	
	{{range $i, $t := .Normalized.Outputs}}
	out{{$i}} := *abi.ConvertType(out[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}){{end}}
	{{else}}
	_, _, err = evm.Call(vm.AccountRef(c.caller), c.to, input, math.MaxUint64, big.NewInt(0))
	{{end}}
	return {{range $i, $t := .Normalized.Outputs}}out{{$i}}, {{end}} err
}
{{end}}
func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) ABI() *abi.ABI {
	return {{$ReceiverName}}.abi
}
func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) WithCaller(caller common.Address) *{{$contract.Type}}GenesisCaller {
	{{$ReceiverName}}.caller = caller
	return {{$ReceiverName}}
}

func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) WithTo(to common.Address) *{{$contract.Type}}GenesisCaller {
	{{$ReceiverName}}.to = to
	return {{$ReceiverName}}
}
func ({{$ReceiverName}} *{{$contract.Type}}GenesisCaller) WithEvmFunc(evmFunc func(address common.Address) *vm.EVM) *{{$contract.Type}}GenesisCaller {
	{{$ReceiverName}}.evmFunc = evmFunc
	return {{$ReceiverName}}
}
`
const tmplTxBuilder = `
package {{.Package}}
import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	
	"math/big"
	"strings"
)

{{if eq .NoStruct false }}
{{$ReceiverName := .ReceiverName}}
{{$structs := .Structs}}
{{range $structs}}
	// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
	type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
	}
{{end}}
{{end}}
{{$ReceiverName := .ReceiverName}}
{{$contract := .Contract}}
var (
    {{$contract.Type}}TxBuilderABI = "{{$contract.InputABI}}"
)

{{$contract := .Contract}}
{{$structs := .Structs}}
{{$ReceiverName := .ReceiverName}}
type {{$contract.Type}}TxBuilder struct {
	abi     *abi.ABI
	sk *ecdsa.PrivateKey
	nonce uint64
	chainId *big.Int
	signer types.Signer
	gasLimit uint64
	gasPrice *big.Int
	to      common.Address
	value *big.Int
}
func New{{$contract.Type}}TxBuilder(to common.Address, sk *ecdsa.PrivateKey, chainId *big.Int) (*{{$contract.Type}}TxBuilder, error) {
	abi, err := abi.JSON(strings.NewReader({{$contract.Type}}TxBuilderABI))
	if err != nil {
		return nil, err
	}
   	return &{{$contract.Type}}TxBuilder{
		abi: &abi,
		sk: sk,
		chainId: chainId,
		signer: types.NewEIP155Signer(chainId),
		gasLimit: 1000000,
		gasPrice: big.NewInt(0),
		to: to,
	}, nil
}


func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) WithPrivateKey(sk *ecdsa.PrivateKey) *{{$contract.Type}}TxBuilder {
	{{$ReceiverName}}.sk = sk
	return {{$ReceiverName}}
}
func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) WithNonce(nonce uint64) *{{$contract.Type}}TxBuilder {
	{{$ReceiverName}}.nonce = nonce
	return {{$ReceiverName}}
}
func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) WithChainId(chainId *big.Int) *{{$contract.Type}}TxBuilder {
	{{$ReceiverName}}.chainId = chainId
	return {{$ReceiverName}}
}

func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) WithSigner(signer types.Signer) *{{$contract.Type}}TxBuilder {
	{{$ReceiverName}}.signer = signer
	return {{$ReceiverName}}
}

func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) WithGasLimit(gasLimit uint64) *{{$contract.Type}}TxBuilder {
	{{$ReceiverName}}.gasLimit = gasLimit
	return {{$ReceiverName}}
}

func({{$ReceiverName}} *{{$contract.Type}}TxBuilder) WithGasPrice(gasPrice *big.Int) *{{$contract.Type}}TxBuilder {
	{{$ReceiverName}}.gasPrice = gasPrice
	return {{$ReceiverName}}
}

func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) WithTo(to common.Address) *{{$contract.Type}}TxBuilder {
	{{$ReceiverName}}.to = to
	return {{$ReceiverName}}
}

{{range .Contract.Calls}}
func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ([]byte,error) {
    input, err := {{$ReceiverName}}.abi.Pack("{{.Original.Name}}",{{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
		return nil, err
	}
	return input, nil
}
func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) (*types.Transaction,error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    var err error
	var input []byte
    input, err = {{$ReceiverName}}.Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
        return nil, err
    }
	tx := types.NewTransaction({{$ReceiverName}}.nonce, {{$ReceiverName}}.to, {{$ReceiverName}}.value, {{$ReceiverName}}.gasLimit, {{$ReceiverName}}.gasPrice, input)
	tx, err = types.SignTx(tx, {{$ReceiverName}}.signer, {{$ReceiverName}}.sk)
	if err != nil {
		return nil, err
	}
	
	return tx, nil
}
{{end}}

{{range .Contract.Transacts}}
func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ([]byte,error) {
    input, err := {{$ReceiverName}}.abi.Pack("{{.Original.Name}}",{{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
		return nil, err
	}
	return input, nil
}
func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) (*types.Transaction,error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    var err error
	var input []byte
    input, err = {{$ReceiverName}}.Pack{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
    if err != nil {
        return nil, err
    }
	tx := types.NewTransaction({{$ReceiverName}}.nonce, {{$ReceiverName}}.to, {{$ReceiverName}}.value, {{$ReceiverName}}.gasLimit, {{$ReceiverName}}.gasPrice, input)
	tx, err = types.SignTx(tx, {{$ReceiverName}}.signer, {{$ReceiverName}}.sk)
	if err != nil {
		return nil, err
	}
	
	return tx, nil
}
{{end}}
func ({{$ReceiverName}} *{{$contract.Type}}TxBuilder) ABI() *abi.ABI {
	return {{$ReceiverName}}.abi
}

`
