package contracts

import (
	"encoding/hex"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"regexp"
	"strings"
)

var (
	methodNormalizer = abi.ToCamelCase
	bindStructType   = bindStructTypeGo
	capitalise       = abi.ToCamelCase
)

// bindBasicTypeGo converts basic solidity types(except array, slice and tuple) to Go ones.
func bindBasicTypeGo(kind abi.Type) string {
	switch kind.T {
	case abi.AddressTy:
		return "common.Address"
	case abi.IntTy, abi.UintTy:
		parts := regexp.MustCompile(`(u)?int([0-9]*)`).FindStringSubmatch(kind.String())
		switch parts[2] {
		case "8", "16", "32", "64":
			return fmt.Sprintf("%sint%s", parts[1], parts[2])
		}
		return "*big.Int"
	case abi.FixedBytesTy:
		if kind.Size == 32 {
			return "common.Hash"
		}
		return fmt.Sprintf("[%d]byte", kind.Size)
	case abi.BytesTy:
		return "[]byte"
	case abi.FunctionTy:
		return "[24]byte"
	default:
		// string, bool types
		return kind.String()
	}
}

// decapitalise makes a camel-case string which starts with a lower case character.
func decapitalise(input string) string {
	if len(input) == 0 {
		return input
	}

	goForm := abi.ToCamelCase(input)
	return strings.ToLower(goForm[:1]) + goForm[1:]
}

func hexId(input []byte) string {
	return hex.EncodeToString(input)
}

// alias returns an alias of the given string based on the aliasing rules
// or returns itself if no rule is matched.
func alias(aliases map[string]string, n string) string {
	if alias, exist := aliases[n]; exist {
		return alias
	}
	return n
}

// bindTypeGo converts solidity types to Go ones. Since there is no clear mapping
// from all Solidity types to Go ones (e.g. uint17), those that cannot be exactly
// mapped will use an upscaled type (e.g. BigDecimal).
func bindTypeGo(kind abi.Type, structs map[string]*tmplStruct) string {
	switch kind.T {
	case abi.TupleTy:
		return structs[kind.TupleRawName+kind.String()].Name
	case abi.ArrayTy:
		return fmt.Sprintf("[%d]", kind.Size) + bindTypeGo(*kind.Elem, structs)
	case abi.SliceTy:
		return "[]" + bindTypeGo(*kind.Elem, structs)
	default:
		return bindBasicTypeGo(kind)
	}
}

// bindTopicTypeGo converts a Solidity topic type to a Go one. It is almost the same
// functionality as for simple types, but dynamic types get converted to hashes.
func bindTopicTypeGo(kind abi.Type, structs map[string]*tmplStruct) string {
	bound := bindTypeGo(kind, structs)

	// todo(rjl493456442) according solidity documentation, indexed event
	// parameters that are not value types i.e. arrays and structs are not
	// stored directly but instead a keccak256-hash of an encoding is stored.
	//
	// We only convert stringS and bytes to hash, still need to deal with
	// array(both fixed-size and dynamic-size) and struct.
	if bound == "string" || bound == "[]byte" {
		bound = "common.Hash"
	}
	return bound
}

// bindStructTypeGo converts a Solidity tuple type to a Go one and records the mapping
// in the given map.
// Notably, this function will resolve and record nested struct recursively.
func bindStructTypeGo(kind abi.Type, structs map[string]*tmplStruct) string {
	switch kind.T {
	case abi.TupleTy:
		// We compose a raw struct name and a canonical parameter expression
		// together here. The reason is before solidity v0.5.11, kind.TupleRawName
		// is empty, so we use canonical parameter expression to distinguish
		// different struct definition. From the consideration of backward
		// compatibility, we concat these two together so that if kind.TupleRawName
		// is not empty, it can have unique id.
		id := kind.TupleRawName + kind.String()
		if s, exist := structs[id]; exist {
			return s.Name
		}
		var fields []*tmplField
		for i, elem := range kind.TupleElems {
			field := bindStructTypeGo(*elem, structs)
			fields = append(fields, &tmplField{Type: field, Name: capitalise(kind.TupleRawNames[i]), SolKind: *elem})
		}
		name := kind.TupleRawName
		name = strings.TrimPrefix(name, "Structs")
		if name == "" {
			name = fmt.Sprintf("Struct%d", len(structs))
		}
		structs[id] = &tmplStruct{
			Name:   name,
			Fields: fields,
		}
		return name
	case abi.ArrayTy:
		return fmt.Sprintf("[%d]", kind.Size) + bindStructTypeGo(*kind.Elem, structs)
	case abi.SliceTy:
		return "[]" + bindStructTypeGo(*kind.Elem, structs)
	default:
		return bindBasicTypeGo(kind)
	}
}
