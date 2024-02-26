package contracts

import (
	"bytes"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	abi "github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"go/format"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
	"unicode"
)

var (
	ContractCommand = cli.Command{
		Name: "contract",

		Subcommands: []cli.Command{
			{
				Action:    utils.MigrateFlags(contractCreate),
				Name:      "new",
				Usage:     "Generate golang contract code",
				ArgsUsage: "<genesisPath>",
				Flags: []cli.Flag{
					abiFlag,
					outputFlag,
					typeFlag,
					pkgFlag,
					aliasFlag,
					receiverNameFlag,
				},
				Category:           "BLOCKCHAIN COMMANDS",
				Description:        `Output golang contract`,
				CustomHelpTemplate: flags.CommandHelpTemplate,
			},
			{
				Action:    utils.MigrateFlags(contractUpgrade),
				Name:      "upgrade",
				Usage:     "Generate golang upgrade contract code",
				ArgsUsage: "<genesisPath>",
				Flags: []cli.Flag{
					abiFlag,
					outputFlag,
					typeFlag,
					pkgFlag,
					aliasFlag,
					versionFlag,
					receiverNameFlag,
					upgradeMethodsFlag,
					newMethodsFlag,
					newStructsFlag,
					newEventsFlag,
				},
				Category:           "BLOCKCHAIN COMMANDS",
				Description:        `Output golang contract`,
				CustomHelpTemplate: flags.CommandHelpTemplate,
			},
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}

	abiFlag = cli.StringFlag{
		Name:        "abi",
		Usage:       "abi",
		EnvVar:      "abi",
		Hidden:      false,
		Value:       "",
		Destination: nil,
	}
	outputFlag = cli.StringFlag{
		Name:        "output",
		Usage:       "output golang file",
		EnvVar:      "output",
		Hidden:      false,
		Value:       "",
		Destination: nil,
	}
	typeFlag = cli.StringFlag{
		Name:  "type",
		Usage: "Struct name for the binding (default = package name)",
	}
	pkgFlag = cli.StringFlag{
		Name:  "pkg",
		Usage: "Package name to generate the binding into",
	}
	aliasFlag = cli.StringFlag{
		Name:  "alias",
		Usage: "Comma separated aliases for function and event renaming, e.g. original1=alias1, original2=alias2",
	}
	receiverNameFlag = cli.StringFlag{
		Name:  "receiver-name",
		Usage: "receiver name",
		Value: "c",
	}
	versionFlag = cli.Uint64Flag{
		Name:  "version",
		Usage: "upgrade version",
	}
	upgradeMethodsFlag = cli.StringFlag{
		Name:  "upgrade-method",
		Usage: "Upgrade method",
	}
	newMethodsFlag = cli.StringFlag{
		Name:  "new-method",
		Usage: "New method,comma separated events, e.g. method1,method2",
	}
	newStructsFlag = cli.StringFlag{
		Name:  "new-struct",
		Usage: "New event,comma separated events, e.g. event1,event2",
	}
	newEventsFlag = cli.StringFlag{
		Name:  "new-event",
		Usage: "New event,comma separated events, e.g. event1,event2",
	}
)

func contractUpgrade(ctx *cli.Context) error {
	if ctx.String(pkgFlag.Name) == "" {
		fmt.Println("No destination package specified (--pkg)")
		os.Exit(1)
	}
	abiFile := ctx.String(abiFlag.Name)

	abiJson, err := os.ReadFile(abiFile)
	if err != nil {
		return err
	}
	version := ctx.Uint64(versionFlag.Name)
	aliases := make(map[string]string)
	upgradeMethods := make(map[string]struct{})
	newMethods := make(map[string]struct{})
	newEvents := make(map[string]struct{})
	newStructs := make(map[string]struct{})
	// Extract all aliases from the flags
	if ctx.IsSet(aliasFlag.Name) {
		// We support multi-versions for aliasing
		// e.g.
		//      foo=bar,foo2=bar2
		//      foo:bar,foo2:bar2
		re := regexp.MustCompile(`(?:(\w+)[:=](\w+))`)
		submatches := re.FindAllStringSubmatch(ctx.String(aliasFlag.Name), -1)
		for _, match := range submatches {
			aliases[match[1]] = match[2]

		}
	}

	if ctx.IsSet(upgradeMethodsFlag.Name) {
		methods := strings.Split(ctx.String(upgradeMethodsFlag.Name), ",")
		for _, m := range methods {
			mv := fmt.Sprintf("%sV%d", m, version)
			upgradeMethods[m] = struct{}{}
			aliases[m] = mv
		}
	}

	if ctx.IsSet(newMethodsFlag.Name) {
		methods := strings.Split(ctx.String(newMethodsFlag.Name), ",")
		for _, m := range methods {
			newMethods[m] = struct{}{}
		}
	}

	if ctx.IsSet(newEventsFlag.Name) {
		methods := strings.Split(ctx.String(newEventsFlag.Name), ",")
		for _, m := range methods {
			newEvents[m] = struct{}{}
		}
	}

	if ctx.IsSet(newStructsFlag.Name) {
		methods := strings.Split(ctx.String(newStructsFlag.Name), ",")
		for _, m := range methods {
			newStructs[m] = struct{}{}
		}
	}
	receiverName := ""
	if ctx.IsSet(receiverNameFlag.Name) {
		receiverName = ctx.String(receiverNameFlag.Name)
	}

	var types string
	if ctx.IsSet(typeFlag.Name) {
		types = ctx.String(typeFlag.Name)
	} else {
		types = ctx.String(pkgFlag.Name)
	}

	funcs, bindData, err := BindData(string(abiJson), types, ctx.String(pkgFlag.Name), aliases)
	if err != nil {
		return err
	}
	data := &tmplUpgradeData{
		tmplData:   bindData,
		NewStructs: make(map[string]*tmplStruct),
	}
	data.Version = uint16(version)
	entries := make(map[string]string)
	for _, e := range data.Contract.Calls {
		entries[hexId(e.Original.ID)] = e.Normalized.Name
	}
	for _, e := range data.Contract.Transacts {
		entries[hexId(e.Original.ID)] = e.Normalized.Name
	}
	data.Entries = entries
	for name, _ := range data.Contract.Calls {
		_, upgrade := upgradeMethods[name]
		_, new := newMethods[name]
		if !upgrade && !new {
			delete(data.Contract.Calls, name)
		}
	}

	for name, _ := range data.Contract.Transacts {
		_, upgrade := upgradeMethods[name]
		_, new := newMethods[name]
		if !upgrade && !new {
			delete(data.Contract.Transacts, name)
		}
	}

	for name, _ := range data.Contract.Events {
		if _, ok := newEvents[name]; !ok {
			delete(data.Contract.Events, name)
		}
	}
	for name, _ := range newStructs {
		for _, v := range data.Structs {
			if v.Name == name {
				data.NewStructs[name] = v
				break
			}
		}
	}

	data.ReceiverName = receiverName
	upgradeCode, err := GenerateCode(funcs, data, tmplUpgrade)
	if err != nil {
		return err
	}
	if !ctx.IsSet(outputFlag.Name) {
		fmt.Printf("%s\n", upgradeCode)
		return nil
	}
	if err := os.WriteFile(filepath.Join(ctx.String(outputFlag.Name), strings.ToLower(types)+fmt.Sprintf("v%d.go", version)), []byte(upgradeCode), 0600); err != nil {
		fmt.Printf("Failed to write ABI binding: %v", err)
		os.Exit(1)
	}
	return nil
}

func contractCreate(ctx *cli.Context) error {
	if ctx.String(pkgFlag.Name) == "" {
		fmt.Println("No destination package specified (--pkg)")
		os.Exit(1)
	}
	abiFile := ctx.String(abiFlag.Name)

	abiJson, err := os.ReadFile(abiFile)
	if err != nil {
		return err
	}
	aliases := make(map[string]string)
	// Extract all aliases from the flags
	if ctx.IsSet(aliasFlag.Name) {
		// We support multi-versions for aliasing
		// e.g.
		//      foo=bar,foo2=bar2
		//      foo:bar,foo2:bar2
		re := regexp.MustCompile(`(?:(\w+)[:=](\w+))`)
		submatches := re.FindAllStringSubmatch(ctx.String(aliasFlag.Name), -1)
		for _, match := range submatches {
			aliases[match[1]] = match[2]
		}
	}
	var types string
	if ctx.IsSet(typeFlag.Name) {
		types = ctx.String(typeFlag.Name)
	} else {
		types = ctx.String(pkgFlag.Name)
	}

	frame, impl, caller, err := Bind(string(abiJson), types, ctx.String(pkgFlag.Name), aliases, ctx.String(receiverNameFlag.Name))
	if err != nil {
		return err
	}
	if !ctx.IsSet(outputFlag.Name) {
		fmt.Printf("%s\n", frame)
		fmt.Printf("%s\n", impl)
		fmt.Printf("%s\n", caller)
		return nil
	}

	if err := os.WriteFile(filepath.Join(ctx.String(outputFlag.Name), strings.ToLower(types)+".go"), []byte(frame), 0600); err != nil {
		fmt.Printf("Failed to write ABI binding: %v", err)
		os.Exit(1)
	}

	if err := os.WriteFile(filepath.Join(ctx.String(outputFlag.Name), strings.ToLower(types)+"_impl.go"), []byte(impl), 0600); err != nil {
		fmt.Printf("Failed to write ABI binding: %v", err)
		os.Exit(1)
	}

	if err := os.WriteFile(filepath.Join(ctx.String(outputFlag.Name), strings.ToLower(types)+"_caller.go"), []byte(caller), 0600); err != nil {
		fmt.Printf("Failed to write ABI binding: %v", err)
		os.Exit(1)
	}
	return nil
}

func Bind(abiJson string, types string, pkg string, aliases map[string]string, receiverName string) (string, string, string, error) {
	funcs, data, err := BindData(abiJson, types, pkg, aliases)
	if err != nil {
		return "", "", "", err
	}
	data.ReceiverName = receiverName
	frameCode, err := GenerateCode(funcs, data, tmplFrameSource)
	if err != nil {
		return "", "", "", err
	}
	sourceCode, err := GenerateCode(funcs, data, implSource)
	if err != nil {
		return "", "", "", err
	}
	callerCode, err := GenerateCode(funcs, data, tmplCaller)
	if err != nil {
		return "", "", "", err
	}
	return frameCode, sourceCode, callerCode, nil
}
func BindData(abiJson string, types string, pkg string, aliases map[string]string) (map[string]interface{}, *tmplData, error) {
	evmABI, err := abi.JSON(strings.NewReader(abiJson))
	if err != nil {
		return nil, nil, err
	}

	// Strip any whitespace from the JSON ABI
	strippedABI := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, abiJson)

	var (
		structs = make(map[string]*tmplStruct)
	)
	var (
		calls               = make(map[string]*tmplMethod)
		transacts           = make(map[string]*tmplMethod)
		events              = make(map[string]*tmplEvent)
		callIdentifiers     = make(map[string]bool)
		transactIdentifiers = make(map[string]bool)
		eventIdentifiers    = make(map[string]bool)
	)
	for _, original := range evmABI.Methods {
		// Normalize the method for capital cases and non-anonymous inputs/outputs
		normalized := original
		normalizedName := methodNormalizer(alias(aliases, original.Name))
		// Ensure there is no duplicated identifier
		var identifiers = callIdentifiers
		if !original.IsConstant() {
			identifiers = transactIdentifiers
		}
		if identifiers[normalizedName] {
			return nil, nil, fmt.Errorf("duplicated identifier \"%s\"(normalized \"%s\"), use --alias for renaming", original.Name, normalizedName)
		}
		identifiers[normalizedName] = true
		normalized.Name = normalizedName
		normalized.Inputs = make([]abi.Argument, len(original.Inputs))
		copy(normalized.Inputs, original.Inputs)
		for j, input := range normalized.Inputs {
			if input.Name == "" {
				normalized.Inputs[j].Name = fmt.Sprintf("arg%d", j)
			}
			if hasStruct(input.Type) {
				bindStructType(input.Type, structs)
			}
		}
		normalized.Outputs = make([]abi.Argument, len(original.Outputs))
		copy(normalized.Outputs, original.Outputs)
		for j, output := range normalized.Outputs {
			if output.Name != "" {
				normalized.Outputs[j].Name = capitalise(output.Name)
			}
			if hasStruct(output.Type) {
				bindStructType(output.Type, structs)
			}
		}
		// Append the methods to the call or transact lists
		if original.IsConstant() {
			calls[original.Name] = &tmplMethod{Original: original, Normalized: normalized, Structured: structured(original.Outputs)}
		} else {
			transacts[original.Name] = &tmplMethod{Original: original, Normalized: normalized, Structured: structured(original.Outputs)}
		}
	}

	for _, original := range evmABI.Events {
		// Skip anonymous events as they don't support explicit filtering
		if original.Anonymous {
			continue
		}
		// Normalize the event for capital cases and non-anonymous outputs
		normalized := original

		// Ensure there is no duplicated identifier
		normalizedName := methodNormalizer(alias(aliases, original.Name))
		if eventIdentifiers[normalizedName] {
			return nil, nil, fmt.Errorf("duplicated identifier \"%s\"(normalized \"%s\"), use --alias for renaming", original.Name, normalizedName)
		}
		eventIdentifiers[normalizedName] = true
		normalized.Name = normalizedName

		normalized.Inputs = make([]abi.Argument, len(original.Inputs))
		copy(normalized.Inputs, original.Inputs)
		for j, input := range normalized.Inputs {
			if input.Name == "" {
				normalized.Inputs[j].Name = fmt.Sprintf("arg%d", j)
			}
			if hasStruct(input.Type) {
				bindStructType(input.Type, structs)
			}
		}
		// Append the event to the accumulator list
		events[original.Name] = &tmplEvent{Original: original, Normalized: normalized}
	}

	contract := &tmplContract{
		Type:        capitalise(types),
		InputABI:    strings.Replace(strippedABI, "\"", "\\\"", -1),
		Constructor: evmABI.Constructor,
		Calls:       calls,
		Transacts:   transacts,
		Events:      events,
	}

	data := &tmplData{
		Package:  pkg,
		Contract: contract,
		Structs:  structs,
	}

	funcs := map[string]interface{}{
		"bindtype":      bindTypeGo,
		"bindtopictype": bindTopicTypeGo,
		"namedtype":     func(string, abi.Type) string { panic("this shouldn't be needed") },
		"capitalise":    capitalise,
		"decapitalise":  decapitalise,
		"hexid":         hexId,
	}
	return funcs, data, nil
}

func GenerateCode(funcs map[string]interface{}, data any, source string) (string, error) {
	buffer := new(bytes.Buffer)

	tmpl := template.Must(template.New("").Funcs(funcs).Parse(source))

	if err := tmpl.Execute(buffer, data); err != nil {
		return "", err
	}
	// For Go bindings pass the code through gofmt to clean it up
	code, err := format.Source(buffer.Bytes())
	return string(code), err
}

// structured checks whether a list of ABI data types has enough information to
// operate through a proper Go struct or if flat returns are needed.
func structured(args abi.Arguments) bool {
	return false
	//if len(args) < 2 {
	//	return false
	//}
	//exists := make(map[string]bool)
	//for _, out := range args {
	//	// If the name is anonymous, we can't organize into a struct
	//	if out.Name == "" {
	//		return false
	//	}
	//	// If the field name is empty when normalized or collides (var, Var, _var, _Var),
	//	// we can't organize into a struct
	//	field := capitalise(out.Name)
	//	if field == "" || exists[field] {
	//		return false
	//	}
	//	exists[field] = true
	//}
	//return true
}

// hasStruct returns an indicator whether the given type is struct, struct slice
// or struct array.
func hasStruct(t abi.Type) bool {
	switch t.T {
	case abi.SliceTy:
		return hasStruct(*t.Elem)
	case abi.ArrayTy:
		return hasStruct(*t.Elem)
	case abi.TupleTy:
		return true
	default:
		return false
	}
}
