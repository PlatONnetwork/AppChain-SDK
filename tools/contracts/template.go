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
	FuncSigs    map[string]string      // Optional map: string signature -> 4-byte signature
	Constructor abi.Method             // Contract constructor for deploy parametrization
	Calls       map[string]*tmplMethod // Contract calls that only read state data
	Transacts   map[string]*tmplMethod // Contract calls that write state data
	Events      map[string]*tmplEvent  // Contract events accessors
}

// tmplData is the data structure required to fill the binding template.
type tmplData struct {
	Package  string                 // Name of the package to place the generated file in
	Contract *tmplContract          // List of contracts to generate into this file
	Structs  map[string]*tmplStruct // Contract struct type definitions
}

const tmplFrameSource = `
package {{.Package}}

import (
	"encoding/hex"
	"errors"
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

{{$structs := .Structs}}
{{range $structs}}
	// {{.Name}} is an auto generated low-level Go binding around an user-defined struct.
	type {{.Name}} struct {
	{{range $field := .Fields}}
	{{$field.Name}} {{$field.Type}}{{end}}
	}
{{end}}

{{$contract := .Contract}}
var (
    ABI = "{{$contract.InputABI}}"
    Abi, _ = abi.JSON(strings.NewReader(ABI))
)

func (c *{{$contract.Type}})Run(input []byte) ([]byte, error) {
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
func (c *{{$contract.Type}}) initMethodEntry() {

    c.methodEntry = map[string]func([]byte) ([]byte, error){
        {{range .Contract.Calls}}"{{hexid .Original.ID}}" : c.{{.Normalized.Name}}Entry,
        {{end}}
        {{range .Contract.Transacts}}
        "{{hexid .Original.ID}}" : c.{{.Normalized.Name}}Entry,{{end}}
    }

}
{{range .Contract.Calls}}
func (c *{{$contract.Type}}) {{.Normalized.Name}}Entry(input []byte) ([]byte, error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    {{if or (ne $inputLen 0) (ne $outputLen 0) }}
    method := c.abi.Methods["{{.Original.Name}}"]
    {{end}}
    var err error
    {{ $length := len .Normalized.Inputs }}
    {{if ne $length 0 }}
    args, err := method.Inputs.Unpack(input)
    if err != nil {
        return nil, err
    }
    {{end}}
    {{range $i, $_ := .Normalized.Outputs}}res{{$i}}, {{end}} err {{if ne $outputLen 0 }} := {{else}} ={{end}} c.{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} *abi.ConvertType(args[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}) {{end}})
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
func (c *{{$contract.Type}}) {{.Normalized.Name}}Entry(input []byte) ([]byte, error) {
    {{ $inputLen := len .Normalized.Inputs }}
    {{ $outputLen := len .Normalized.Outputs }}
    {{if or (ne $inputLen 0) (ne $outputLen 0) }}
    method := c.abi.Methods["{{.Original.Name}}"]
    {{end}}
    var err error

    {{ $length := len .Normalized.Inputs }}
    {{if ne $length 0 }}
    args, err := method.Inputs.Unpack(input)
    if err != nil {
        return nil, err
    }
    {{end}}
    {{range $i, $_ := .Normalized.Outputs}}res{{$i}}, {{end}} err {{if ne $outputLen 0 }} := {{else}} ={{end}} c.{{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} *abi.ConvertType(args[{{$i}}], new({{bindtype .Type $structs}})).(*{{bindtype .Type $structs}}) {{end}})
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
        func (c *{{$contract.Type}})Emit{{.Normalized.Name}}Event({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}} {{bindtype .Type $structs}}{{end}}) (*types.Log, error){
        event := c.abi.Events["{{.Normalized.Name}}"]
        hashes, err := abi.PackTopics(event.Inputs, {{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
        if err != nil {
            return nil, err
        }
        data, err := event.Inputs.Pack({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}}{{.Name}}{{end}})
        if err != nil {
            return nil, err
        }
        return &types.Log{
            Address: c.contract.Address(),
            Topics: hashes,
            Data:   data,
            BlockNumber: c.evm.Context.BlockNumber.Uint64(),
        }, nil
        }
{{end}}
`
const implSource = `
package {{.Package}}

import (
	"errors"
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

type {{$contract.Type}} struct {
    abi *abi.ABI
    methodEntry map[string]func([]byte) ([]byte, error)
    readOnly bool
    contract *vm.Contract
    evm *vm.EVM
	fallback func(input []byte) ([]byte, error)
}

func New{{$contract.Type}}(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*{{$contract.Type}}, error) {
    s := &{{$contract.Type}}{
        abi: &Abi,
        evm:evm,
        contract: contract,
        readOnly: readOnly,
    }
    s.initMethodEntry()
    return s, nil
}

{{range .Contract.Calls}}
func (c *{{$contract.Type}}) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    panic("implement")
}
{{end}}


{{range .Contract.Transacts}}
func (c *{{$contract.Type}}) {{.Normalized.Name}}({{range $i, $_ := .Normalized.Inputs}}{{if ne $i 0}},{{end}} {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
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

type {{$contract.Type}}Caller struct {
    contracts.BoundContract
}
func New{{$contract.Type}}Caller(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*{{$contract.Type}}Caller, error) {
    s := &{{$contract.Type}}Caller{
		contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
    }
    return s, nil
}



{{range .Contract.Calls}}
func (c *{{$contract.Type}}Caller) {{.Normalized.Name}}(to common.Address {{range $i, $_ := .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    var out []interface{}
	err := c.BoundContract.Caller(to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
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
func (c *{{$contract.Type}}Caller) {{.Normalized.Name}}(to common.Address {{range $i, $_ := .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
        var out []interface{}
	err := c.BoundContract.Caller(to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
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
}
func New{{$contract.Type}}DelegateCaller(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*{{$contract.Type}}DelegateCaller, error) {
    s := &{{$contract.Type}}DelegateCaller{
		contracts.BoundContract{
			Abi:      &Abi,
			Evm:      evm,
			Contract: contract,
		},
    }
    return s, nil
}



{{range .Contract.Calls}}
func (c *{{$contract.Type}}DelegateCaller) {{.Normalized.Name}}(to common.Address {{range $i, $_ := .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
    var out []interface{}
	err := c.BoundContract.DelegateCaller(to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
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
func (c *{{$contract.Type}}DelegateCaller) {{.Normalized.Name}}(to common.Address {{range $i, $_ := .Normalized.Inputs}}, {{.Name}} {{bindtype .Type $structs}} {{end}}) ({{if .Structured}}struct{ {{range .Normalized.Outputs}}{{.Name}} {{bindtype .Type $structs}};{{end}} },{{else}}{{range .Normalized.Outputs}}{{bindtype .Type $structs}},{{end}}{{end}} error) {
        var out []interface{}
	err := c.BoundContract.DelegateCaller(to, &out, "{{.Original.Name}}" {{range .Normalized.Inputs}}, {{.Name}}{{end}})
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
