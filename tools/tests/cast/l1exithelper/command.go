package l1exithelper

import (
	"context"
	"fmt"
	"math/big"

	"github.com/PlatONnetwork/AppChain-SDK/client"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/transaction"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/ethclient"
	"gopkg.in/urfave/cli.v1"
)

var (
	appchainRPCFlag = cli.StringFlag{
		Name:  "appchain-rpc",
		Usage: "AppChain RPC endpoint",
	}

	exitIDFlag = cli.Uint64Flag{
		Name:  "exit-id",
		Usage: "Exit event identifier",
	}

	Command = cli.Command{
		Action:    utils.MigrateFlags(run),
		Name:      "l1.exithelper",
		Usage:     "l1.exithelper",
		ArgsUsage: "",
		Flags: []cli.Flag{
			appchainRPCFlag,
			flags.RPCFlags,
			flags.KeyFlags,
			flags.TypeFlags,
			flags.AddressFlags,
			exitIDFlag,
		},
		Category:           "L1 EXIT HELPER",
		Description:        ``,
		HelpName:           "cast l1.exithelper",
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func initGlobal(ctx *cli.Context) (*client.Client, *ethclient.Client, *Exithelper, *bind.TransactOpts, error) {
	rootClient, exitHelperAddr, opt, err := flags.InitGlobal(ctx)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	exitHelper, err := NewExithelper(exitHelperAddr, rootClient)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	appClient, err := client.Dial(ctx.String(appchainRPCFlag.Name))
	if err != nil {
		return nil, nil, nil, nil, err
	}
	return appClient, rootClient, exitHelper, opt, nil
}

func run(ctx *cli.Context) error {
	appClient, rootClient, exitHelper, opt, err := initGlobal(ctx)
	if err != nil {
		return err
	}

	exitID, ok := new(big.Int).SetString(ctx.String(exitIDFlag.Name), 10)
	if !ok {
		return fmt.Errorf("invalid exit event identifier")
	}

	proof, err := appClient.GenerateExitProof(context.Background(), exitID.Uint64())
	if err != nil {
		return fmt.Errorf("Generate exit proof error: %v\n", err)
	}

	blockNumber := new(big.Int).SetUint64(proof.Metadata.CheckpointBlock)
	unhashedLeaf := common.Hex2Bytes(proof.Metadata.ExitEvent)
	leafIndex := new(big.Int).SetUint64(proof.Metadata.LeafIndex)
	var proofs [][32]byte
	for _, d := range proof.Data {
		proofs = append(proofs, d)
	}
	tx, err := exitHelper.Exit(opt, blockNumber, leafIndex, unhashedLeaf, proofs)
	if err != nil {
		return fmt.Errorf("Send exit transaction error: %v\n", err)
	}
	_, err = transaction.WaitTx(rootClient, tx.Hash())
	if err != nil {
		return err
	}
	fmt.Println("send success", tx.Hash())
	return nil
}
