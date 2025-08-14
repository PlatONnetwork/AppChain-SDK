package deploy

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v3"
	"os"

	"path/filepath"
)

type ExtraFunc func(nodeDir string, node *Node) (string, error)

var (
	CLOSE                    = 0
	LOCAL                    = 1
	SERVER                   = 2
	ANY                      = 3
	CreateAnsibleNodeCommand = cli.Command{
		Name:   "createansiblenode",
		Action: createAnsibleNode,
		Flags: []cli.Flag{
			OutputFlag,
			BinFlag,
			AnsibleDirFlag,
			NodeConfigFileFlag,
			GenesisFileFlag,
			PProfFlag,
			HttpFlag,
			StartArgsFlag,
			UsernameFlag,
			PasswordFlag,
			LocalFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func createAnsibleNode(ctx *cli.Context) error {
	return CreateAnsibleNodeExtra(ctx, nil)
}
func CreateAnsibleNodeExtra(ctx *cli.Context, extraFunc ExtraFunc) error {
	username := ctx.String(UsernameFlag.Name)
	password := ctx.String(PasswordFlag.Name)
	outputDir := ctx.String(OutputFlag.Name)
	ansibleDir := filepath.Join(outputDir, ctx.String(AnsibleDirFlag.Name))
	nodeConfig, err := ReadNodeConfig(filepath.Join(outputDir, ctx.String(NodeConfigFileFlag.Name)))
	if err != nil {
		return err
	}
	log.Debug("Read child chain node", "path", ctx.String(NodeConfigFileFlag.Name))

	filesDir := filepath.Join(ansibleDir, "playbooks", "files")
	os.Mkdir(filesDir, 0755)
	varsDir := filepath.Join(ansibleDir, "playbooks", "vars")
	os.Mkdir(varsDir, 0755)
	if err := GenerateAnsible(ctx.Bool(LocalFlag.Name), username, password, ansibleDir, nodeConfig.Nodes); err != nil {
		return err
	}
	genesisJson, err := os.ReadFile(filepath.Join(outputDir, ctx.String(GenesisFileFlag.Name)))
	if err != nil {
		log.Error("Read genesis json failed", "path", filepath.Join(outputDir, ctx.String(GenesisFileFlag.Name)), "err", err)
		return err
	}
	return GenerateNodes(nodeConfig.Nodes, ctx.Int(PProfFlag.Name), ctx.Int(HttpFlag.Name), ctx.String(StartArgsFlag.Name), filesDir,
		string(genesisJson), ctx.String(BinFlag.Name), extraFunc)
}
func GenerateNodes(accs []*Node, pprof, http int, baseArgs, filesDir, genesis, binary string, extraFunc ExtraFunc) error {
	for i, acc := range accs {
		nodeDir := filepath.Join(filesDir, fmt.Sprintf("node%d", i))
		os.RemoveAll(nodeDir)

		err := os.Mkdir(nodeDir, 0755)
		if err != nil {
			log.Error("mkdir failed", "path", nodeDir)
			return err
		}

		log.Debug("Create node dir", "dir", nodeDir)
		if err = CopyFile(binary, filepath.Join(nodeDir, "node")); err != nil {
			return err
		}
		log.Debug("Copy binary success", "host", acc.IP)

		if err = os.WriteFile(filepath.Join(nodeDir, "blskey"), []byte(acc.BlsPrivateKey), 0755); err != nil {
			return err
		}
		log.Debug("Write blskey success", "host", acc.IP)

		if err = os.WriteFile(filepath.Join(nodeDir, "blspub"), []byte(acc.BlsPublicKey), 0755); err != nil {
			return err
		}
		log.Debug("Write blspub success", "host", acc.IP)

		if err = os.WriteFile(filepath.Join(nodeDir, "nodekey"), []byte(acc.NodePrivateKey), 0755); err != nil {
			return err
		}
		log.Debug("Write nodekey success", "host", acc.IP)

		if err = os.WriteFile(filepath.Join(nodeDir, "nodepub"), []byte(acc.NodePublicKey), 0755); err != nil {
			return err
		}
		log.Debug("Write nodepub success", "host", acc.IP)

		extraArgs := baseArgs
		if extraFunc != nil {
			args, err := extraFunc(nodeDir, acc)
			if err != nil {
				log.Error("Execute extra func failed", "node", acc)
				return err
			}
			extraArgs = baseArgs + " " + args
		}
		httpAddr, pprofAddr := convertIP(http, acc.IP), convertIP(pprof, acc.IP)
		if err := os.WriteFile(filepath.Join(nodeDir, "start.sh"), []byte(generateStart(acc.P2pPort, pprofAddr, acc.PprofPort, httpAddr, acc.RpcPort, extraArgs)), 0755); err != nil {
			return err
		}
		log.Debug("Write start.sh success", "host", acc.IP)

		if err := os.WriteFile(filepath.Join(nodeDir, "stop.sh"), []byte(generateStop(acc.P2pPort)), 0755); err != nil {
			return err
		}
		log.Debug("Write stop.sh success", "host", acc.IP)

		if err := os.WriteFile(filepath.Join(nodeDir, "init.sh"), []byte(generateInit()), 0755); err != nil {
			return err
		}
		log.Debug("Write init.sh success", "host", acc.IP)

		if err := os.WriteFile(filepath.Join(nodeDir, "clean.sh"), []byte(generateClean()), 0755); err != nil {
			return err
		}
		log.Debug("Write clean.sh success", "host", acc.IP)

		if err = os.WriteFile(filepath.Join(nodeDir, "genesis.json"), []byte(genesis), 0755); err != nil {
			return err
		}
		log.Debug("Copy genesis.json success", "host", acc.IP)

	}
	return nil
}
func convertIP(ty int, ip string) string {
	switch ty {
	case LOCAL:
		return "127.0.0.1"
	case SERVER:
		return ip
	case ANY:
		return "0.0.0.0"
	}
	return ""
}
func GenerateAnsible(local bool, username, password, ansibleDir string, accs []*Node) error {
	hosts := make(map[string]Host)
	var envConfig EnvConfig
	for i, acc := range accs {
		name := fmt.Sprintf("node%d", i)
		archive := fmt.Sprintf("node%d.tar.gz", i)
		hosts[name] = Host{
			AnsibleHost:     acc.IP,
			AnsibleUser:     username,
			AnsiblePassword: password,
		}
		envConfig.NodeFolders = append(envConfig.NodeFolders, NodeFolder{
			Local:   name,
			Archive: archive,
		})
	}
	hc := HostConfig{
		All: All{
			Hosts: hosts,
			Vars: Vars{
				AnsibleConnection:    "ssh",
				AnsibleSSHCommonArgs: "-o StrictHostKeyChecking=no",
			},
		},
	}
	if !local {
		hostsData, err := yaml.Marshal(hc)
		if err != nil {
			return err
		}
		if err = os.WriteFile(filepath.Join(ansibleDir, "inventories", "hosts.yml"), []byte(hostsData), 0755); err != nil {
			log.Error("Write hosts.yml failed", "err", err)
			return err
		}
	}

	envData, err := yaml.Marshal(envConfig)
	if err = os.WriteFile(filepath.Join(ansibleDir, "playbooks", "vars", "env.yml"), []byte(envData), 0755); err != nil {
		log.Error("Write env.yml failed", "err", err)
		return err
	}

	return nil
}

func generateStart(tcp int, pprofAddr string, pprofPort int, httpAddr string, httpPort int, args string) string {
	httpArgs, pprofArgs := "", ""
	if len(httpAddr) != 0 {
		httpArgs = fmt.Sprintf("--http --http.ethcompatible --http.addr %s --http.port %d", httpAddr, httpPort)
	}
	if len(pprofAddr) != 0 {
		pprofArgs = fmt.Sprintf("--pprof --pprof.addr %s --pprof.port %d", pprofAddr, pprofPort)
	}
	format := "#!/bin/bash\nnohup ./node --datadir ./data --nodekey ./nodekey --cbft.blskey ./blskey --port %d  %s %s %s > ./node.log 2>&1 &"
	return fmt.Sprintf(format, tcp, pprofArgs, httpArgs, args)
}
func generateStop(tcp int) string {
	format := "#!/bin/bash\nps aux | grep node | grep %d | grep -v grep | awk '{print $2}' | xargs kill -9"
	return fmt.Sprintf(format, tcp)
}
func generateInit() string {
	format := "#!/bin/bash\n./node init --datadir ./data genesis.json"
	return fmt.Sprintf(format)
}

func generateClean() string {
	format := "#!/bin/bash\nrm -rf data/*"
	return fmt.Sprintf(format)
}
