package main

import (
	"fmt"
	deploytools "github.com/PlatONnetwork/AppChain-SDK/tools/deploy"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
)

func createAnsibleNodeCommand() *cli.Command {
	cmd := deploytools.CreateAnsibleNodeCommand
	cmd.Action = createAnsiblenode
	cmd.Flags = append(cmd.Flags, []cli.Flag{
		RootChianUrlFlag,
		CheckpointKeystoreFileFlag,
		L2TxKeystoreFileFlag,
		ChildchainContractFileFlag,
		ExtraArgsFlag,
	}...)
	return &cmd
}
func createAnsiblenode(ctx *cli.Context) error {
	conf, err := readChildChainContractConfig(filepath.Join(ctx.String(deploytools.OutputFlag.Name), ctx.String(ChildchainContractFileFlag.Name)))
	if err != nil {
		return err
	}
	return deploytools.CreateAnsibleNodeExtra(ctx, func(nodeDir string, node *deploytools.Node) (string, error) {
		checkpointKeystore := ctx.String(CheckpointKeystoreFileFlag.Name)
		l2txKeystore := ctx.String(L2TxKeystoreFileFlag.Name)
		checkpointKeystoreFile := fmt.Sprintf("%s.json", checkpointKeystore)

		if err := os.WriteFile(filepath.Join(nodeDir, checkpointKeystoreFile),
			[]byte(node.Extra.(map[string]interface{})["checkpoint_keystore"].(string)), 0755); err != nil {
			return "", err
		}
		checkpointKeystorePassword := fmt.Sprintf("%s_password", checkpointKeystore)
		if err := os.WriteFile(filepath.Join(nodeDir, checkpointKeystorePassword),
			[]byte(node.Extra.(map[string]interface{})["checkpoint_keystore_password"].(string)), 0755); err != nil {
			return "", err
		}
		l2txKeystoreFile := fmt.Sprintf("%s.json", l2txKeystore)
		if err := os.WriteFile(filepath.Join(nodeDir, l2txKeystoreFile),
			[]byte(node.Extra.(map[string]interface{})["l2_keystore"].(string)), 0755); err != nil {
			return "", err
		}
		l2txKeystorePassword := fmt.Sprintf("%s_password", l2txKeystore)

		if err := os.WriteFile(filepath.Join(nodeDir, l2txKeystorePassword),
			[]byte(node.Extra.(map[string]interface{})["l2_keystore_password"].(string)), 0755); err != nil {
			return "", err
		}
		args := fmt.Sprintf("--checkpoint.keystore ./%s --checkpoint.password ./%s --l2.keystore ./%s --l2.password ./%s --statesync.startblock %d --rootchain-node-rpc %s %s",
			checkpointKeystoreFile, checkpointKeystorePassword, l2txKeystoreFile, l2txKeystorePassword, conf.DeployNumber, ctx.String(RootChianUrlFlag.Name), ctx.String(ExtraArgsFlag.Name))
		return args, nil
	})
}
