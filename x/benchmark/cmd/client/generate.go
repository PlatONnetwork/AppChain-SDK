package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"gopkg.in/urfave/cli.v1"
	"gopkg.in/yaml.v3"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var (
	hostsFlag      = cli.StringFlag{Name: "hosts", EnvVar: "BENCHMARK_HOSTS", Usage: "hosts fal"}
	ansibleDirFlag = cli.StringFlag{Name: "ansible_dir", EnvVar: "BENCHMARK_ANSIBLE_DIR"}
	userFlag       = cli.StringFlag{Name: "user", EnvVar: "BENCHMARK_USERNAME"}
	passwordFlag   = cli.StringFlag{Name: "password", EnvVar: "BENCHMARK_PASSWORD"}
	binFlag        = cli.StringFlag{Name: "bin", EnvVar: "BENCHMARK_BIN"}
)

type HostConfig struct {
	All All `yaml:"all"`
}

type All struct {
	Hosts map[string]Host `yaml:"hosts"`
	Vars  Vars            `yaml:"vars"`
}
type Vars struct {
	AnsibleConnection    string `yaml:"ansible_connection"`
	AnsibleSSHCommonArgs string `yaml:"ansible_ssh_common_args"`
}
type Host struct {
	AnsibleHost     string `yaml:"ansible_host"`
	AnsibleUser     string `yaml:"ansible_user"`
	AnsiblePassword string `yaml:"ansible_password"`
}

type NodeFolder struct {
	Local   string `yaml:"local"`
	Archive string `yaml:"archive"`
}

type EnvConfig struct {
	NodeFolders []NodeFolder `yaml:"node_folders"`
}

func Generate(ctx *cli.Context) error {
	hosts, err := splitHost(ctx)
	if err != nil {
		return err
	}
	var accounts []*testutil.Account
	if len(hosts) <= 4 {
		accounts = testutil.DefaultAccount[0:len(hosts)]
	} else {
		return errors.New("host address too much")
		//TODO create account
	}
	for i, h := range hosts {
		accounts[i].Host = h
	}
	genesisJson, err := testutil.GenerateGenesis(accounts, map[string]json.RawMessage{
		"benchmark": json.RawMessage("{}"),
		"election":  json.RawMessage("{}"),
	}, nil)
	if err != nil {
		return err
	}
	binary := ctx.String(binFlag.Name)
	dir := ctx.String(ansibleDirFlag.Name)
	filesDir := filepath.Join(dir, "playbooks", "files")
	for i, _ := range hosts {
		nodeDir := filepath.Join(filesDir, fmt.Sprintf("node%d", i))
		os.RemoveAll(nodeDir)
		err := os.Mkdir(nodeDir, 0755)
		if err != nil {
			return err
		}
		if err = CopyFile(binary, filepath.Join(nodeDir, "benchmark")); err != nil {
			return err
		}
		privateKey := accounts[i].BlsSecretKey()
		blsKey, blsPub := hex.EncodeToString(privateKey.Serialize()), hex.EncodeToString(privateKey.GetPublicKey().Serialize())
		if err = os.WriteFile(filepath.Join(nodeDir, "blskey"), []byte(blsKey), 0755); err != nil {
			return err
		}
		if err = os.WriteFile(filepath.Join(nodeDir, "blspub"), []byte(blsPub), 0755); err != nil {
			return err
		}
		key := accounts[i].NodePrivateKey()
		nodeKey, nodePub := hex.EncodeToString(crypto.FromECDSA(key)), hex.EncodeToString(crypto.FromECDSAPub(&key.PublicKey)[1:])
		if err = os.WriteFile(filepath.Join(nodeDir, "nodekey"), []byte(nodeKey), 0755); err != nil {
			return err
		}
		if err = os.WriteFile(filepath.Join(nodeDir, "nodepub"), []byte(nodePub), 0755); err != nil {
			return err
		}
		if err = os.WriteFile(filepath.Join(nodeDir, "genesis.json"), genesisJson, 0755); err != nil {
			return err
		}
		if err = outputScript(nodeDir, accounts[i]); err != nil {
			return err
		}
	}
	if err := GenerateAnsible(ctx, dir, accounts); err != nil {
		return err
	}
	return nil
}

func GenerateAnsible(ctx *cli.Context, ansibleDir string, accs []*testutil.Account) error {
	username := ctx.String(userFlag.Name)
	password := ctx.String(passwordFlag.Name)
	hosts := make(map[string]Host)
	var envConfig EnvConfig
	for i, acc := range accs {
		name := fmt.Sprintf("node%d", i)
		archive := fmt.Sprintf("node%d.tar.gz", i)
		hosts[name] = Host{
			AnsibleHost:     acc.Host.String(),
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
	hostsData, err := yaml.Marshal(hc)
	if err != nil {
		return err
	}
	if err = os.WriteFile(filepath.Join(ansibleDir, "inventories", "hosts.yml"), []byte(hostsData), 0755); err != nil {
		return err
	}

	envData, err := yaml.Marshal(envConfig)
	if err = os.WriteFile(filepath.Join(ansibleDir, "playbooks", "vars", "env.yml"), []byte(envData), 0755); err != nil {
		return err
	}
	return nil
}

func splitHost(ctx *cli.Context) ([]net.IP, error) {
	hosts := strings.Split(ctx.String(hostsFlag.Name), ",")
	if len(hosts) == 0 {
		return nil, errors.New("empty host")
	}
	var ips []net.IP
	for _, h := range hosts {
		ips = append(ips, net.ParseIP(h))
	}
	return ips, nil
}

func CopyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
func generateBls() (string, string) {
	var privateKey bls.SecretKey
	privateKey.SetByCSPRNG()
	pubKey := privateKey.GetPublicKey()
	return hex.EncodeToString(privateKey.Serialize()), hex.EncodeToString(pubKey.Serialize())
}

func generateEcdsa() (string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		utils.Fatalf("Failed to generate random private key: %v", err)
	}
	return hex.EncodeToString(crypto.FromECDSA(privateKey)), hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])
}

func outputScript(path string, acc *testutil.Account) error {
	if err := os.WriteFile(filepath.Join(path, "start.sh"), []byte(generateStart(acc.P2PPort, acc.Pprof, acc.HTTP)), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(path, "stop.sh"), []byte(generateStop(acc.P2PPort)), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(path, "init.sh"), []byte(generateInit()), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(path, "clean.sh"), []byte(generateClean()), 0755); err != nil {
		return err
	}

	return nil
}

func generateStart(tcp int, pprof, http int) string {
	format := "#!/bin/bash\nnohup ./benchmark --identity platon --datadir ./data --nodekey ./nodekey --cbft.blskey ./blskey --port %d --verbosity 1 --pprof --pprof.addr 0.0.0.0 --pprof.port %d  --http --http.ethcompatible --http.addr 0.0.0.0 --http.port %d --http.api platon,debug,personal,admin,net,web3,txpool,benchmark --http.vhosts \"*\" --cache 256 --metrics --ipcdisable  --maxpeers 100 --maxconsensuspeers 75 --txpool.globaltxcount 1000 --nodiscover  --networkid 102 --allow-insecure-unlock > ./platon.log 2>&1 &"
	return fmt.Sprintf(format, tcp, pprof, http)
}
func generateStop(tcp int) string {
	format := "#!/bin/bash\nps aux | grep benchmark | grep %d | grep -v grep | awk '{print $2}' | xargs kill -9"
	return fmt.Sprintf(format, tcp)
}
func generateInit() string {
	format := "#!/bin/bash\n./benchmark init --datadir ./data genesis.json"
	return fmt.Sprintf(format)
}

func generateClean() string {
	format := "#!/bin/bash\nrm -rf data/*"
	return fmt.Sprintf(format)
}
