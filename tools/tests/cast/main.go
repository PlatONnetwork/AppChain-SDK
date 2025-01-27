package cast

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"gopkg.in/urfave/cli.v1"
	"strings"
)

var (
	Command = cli.Command{
		Action:      utils.MigrateFlags(run),
		Name:        "cast",
		Usage:       "cast",
		ArgsUsage:   "",
		Subcommands: []cli.Command{},
		Flags: []cli.Flag{
			flags.RPCFlags,
			flags.KeyFlags,
			flags.ModuleFlags,
			flags.AbiFileFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
			flags.StartFlags,
			flags.EndFlags,
		},
		Category:           "CAST COMMANDS",
		Description:        ``,
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func run(ctx *cli.Context) error {
	_, inputs, err := FindMethodArgs(ctx)
	if err != nil {
		return err
	}
	method := ctx.String(flags.MethodFlags.Name)
	typeName := ctx.String(flags.TypeFlags.Name)
	moduleAbi, err := GetAbi(ctx)
	if err != nil {
		return err
	}
	switch typeName {
	case "send":
		client, addr, opt, err := InitGlobal(ctx)
		if err != nil {
			return err
		}
		tx, err := Send(moduleAbi, method, inputs, addr, client, opt)
		if err != nil {
			return err
		}
		_, err = WaitTx(client, tx.Hash())
		if err != nil {
			fmt.Println("send failed", tx.Hash())
			return err
		}
		fmt.Println("send success", tx.Hash())
		tx, _, err = client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			return err
		}
		bytes, _ := json.MarshalIndent(tx, " ", " ")
		fmt.Println("transaction:", string(bytes))
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			return err
		}
		bytes, _ = json.MarshalIndent(receipt, " ", " ")

		fmt.Println("receipt:", string(bytes))

	case "call":
		client, addr, _, err := InitGlobal(ctx)
		if err != nil {
			return err
		}
		result, err := Call(moduleAbi, method, inputs, addr, client, &bind.CallOpts{})
		if err != nil {
			return err
		}

		fmt.Println(string(result))

	case "logs":
		startBlock := ctx.Uint64(flags.StartFlags.Name)
		var endBlock *uint64
		end := ctx.Uint64(flags.EndFlags.Name)
		if end != 0 {
			endBlock = &end
		}
		client, addr, _, err := InitGlobal(ctx)
		if err != nil {
			return err
		}
		logs, err := FilterLog(moduleAbi, method, inputs, addr, client, &bind.FilterOpts{
			Start:   startBlock,
			End:     endBlock,
			Context: context.Background(),
		})
		if err != nil {
			return err
		}
		result, err := json.MarshalIndent(logs, "", "")
		if err != nil {
			return err
		}
		fmt.Println(string(result))
	case "abi":
		for _, method := range moduleAbi.Methods {
			fmt.Println("  ", method.Sig)
		}

		fmt.Println("events:")
		for _, event := range moduleAbi.Events {
			rawAux := []string{}
			for _, i := range event.Inputs.TupleElems() {
				name := i.Elem.Format(false)
				if i.Indexed {
					name += " indexed"
				}
				rawAux = append(rawAux, name)
			}
			fmt.Println("  ", fmt.Sprintf("%s(%s)", event.Name, strings.Join(rawAux, ",")))
		}
	}
	return nil
}
