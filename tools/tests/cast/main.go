package cast

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"gopkg.in/urfave/cli.v1"
	"strings"
)

var (
	Command = cli.Command{
		Action:    utils.MigrateFlags(run),
		Name:      "cast",
		Usage:     "cast",
		ArgsUsage: "",
		Subcommands: []cli.Command{
			ExitHelperCommand,
			RawSendCommand,
		},
		Flags: []cli.Flag{
			flags.AppchainRPCFlag,
			flags.KeyFlags,
			flags.ModuleFlags,
			flags.AbiFileFlags,
			flags.AddressFlags,
			flags.MethodFlags,
			flags.TypeFlags,
			flags.FromFlags,
			flags.StartFlags,
			flags.EndFlags,
			flags.GasLimitFlags,
			flags.GasPriceFlags,
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
		callFunc := func() {
			chainid, _ := client.ChainID(context.Background())
			_, err := client.CallContract(context.Background(), platon.CallMsg{
				From:       tx.FromAddr(types.NewLondonSigner(chainid)),
				To:         tx.To(),
				Gas:        tx.Gas(),
				GasPrice:   tx.GasPrice(),
				GasFeeCap:  tx.GasFeeCap(),
				GasTipCap:  tx.GasTipCap(),
				Value:      tx.Value(),
				Data:       tx.Data(),
				AccessList: tx.AccessList(),
			}, nil)

			fmt.Println("send failed", tx.Hash(), "err", err)
		}
		if err != nil {
			callFunc()
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
		if receipt.Status == 0 {
			callFunc()
			return fmt.Errorf("receipt status is failed")
		}
		bytes, _ = json.MarshalIndent(receipt, " ", " ")

		fmt.Println("receipt:", string(bytes))

	case "call":
		client, addr, _, err := InitGlobal(ctx)
		if err != nil {
			return err
		}
		result, err := Call(moduleAbi, method, inputs, addr, client, &bind.CallOpts{
			From: common.HexToAddress(ctx.String(flags.FromFlags.Name)),
		})
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
		fmt.Println("methods:")
		for _, method := range moduleAbi.Methods {
			rawAux := []string{}
			for _, i := range method.Inputs.TupleElems() {
				name := i.Elem.Format(false) + " " + i.Name

				rawAux = append(rawAux, name)
			}
			outputs := []string{}
			for _, i := range method.Outputs.TupleElems() {
				name := i.Elem.Format(false)
				rawAux = append(outputs, name)
			}
			fmt.Println("  ", fmt.Sprintf("%s(%s)(%s)", method.Name, strings.Join(rawAux, ","), strings.Join(outputs, ",")))
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
