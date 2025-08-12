package main

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/PlatONnetwork/AppChain-SDK/testutil"
	"github.com/PlatONnetwork/AppChain-SDK/tools/deploy"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"gopkg.in/urfave/cli.v1"
	"net"
	"os"
	"path/filepath"
	"strings"
)

var (
	hostsFlag     = cli.StringFlag{Name: "hosts", EnvVar: "BENCHMARK_HOSTS", Usage: "hosts"}
	userFlag      = cli.StringFlag{Name: "user", EnvVar: "BENCHMARK_USERNAME"}
	passwordFlag  = cli.StringFlag{Name: "password", EnvVar: "BENCHMARK_PASSWORD"}
	binFlag       = cli.StringFlag{Name: "bin", EnvVar: "BENCHMARK_BIN"}
	extraArgsFlag = cli.StringFlag{
		Name:   "extra_args",
		EnvVar: "BENCHMARK_EXTRA_ARGS",
		Value:  "--asyncblock.computersenderthread 4 --asyncblock.entrysize 1000 --asyncblock.splitthreshold 3000 --asyncblock.concurrency_level 7 --asyncblock.txs_batch 512 --cbft.wal.disabled"}
)

func Generate(ctx *cli.Context) error {
	username := ctx.String(userFlag.Name)
	password := ctx.String(passwordFlag.Name)
	hosts, err := splitHost(ctx)
	if err != nil {
		return err
	}
	var nodes []*deploy.Node
	var accounts []*testutil.Account
	if len(hosts) <= 4 {
		accounts = testutil.DefaultAccount[0:len(hosts)]
	} else {
		return errors.New("host address too much")
		//TODO create account
	}
	for i, h := range hosts {
		accounts[i].Host = h
		nodes = append(nodes, &deploy.Node{
			Genesis:                  true,
			IP:                       h.String(),
			PprofPort:                accounts[i].Pprof,
			P2pPort:                  accounts[i].P2PPort,
			RpcPort:                  accounts[i].HTTP,
			NodeAddress:              accounts[i].NodeAddress().Hex(),
			NodePrivateKey:           accounts[i].NodeKey,
			NodePublicKey:            hex.EncodeToString(crypto.FromECDSAPub(&accounts[i].NodePrivateKey().PublicKey)),
			BlsPrivateKey:            accounts[i].BlsKey,
			BlsPublicKey:             hex.EncodeToString(accounts[i].BlsSecretKey().GetPublicKey().Serialize()),
			BlsUncompressedPublicKey: hex.EncodeToString(accounts[i].BlsSecretKey().GetPublicKey().SerializeUncompressed()),
			Extra:                    nil,
		})
	}

	genesisJson, err := testutil.GenerateGenesis(accounts, map[string]json.RawMessage{
		"benchmark": json.RawMessage("{}"),
		"election":  json.RawMessage("{}"),
	}, nil)
	if err != nil {
		return err
	}
	binary := ctx.String(binFlag.Name)
	dir := filepath.Join(ctx.String(deploy.OutputFlag.Name), ctx.String(deploy.AnsibleDirFlag.Name))
	filesDir := filepath.Join(dir, "playbooks", "files")
	os.Mkdir(filesDir, 0755)
	varsDir := filepath.Join(dir, "playbooks", "vars")
	os.Mkdir(varsDir, 0755)
	if err = deploy.GenerateAnsible(username, password, dir, nodes); err != nil {
		return err
	}
	if err = deploy.GenerateNodes(nodes, deploy.ANY, deploy.ANY, ctx.String(deploy.StartArgsFlag.Name)+" "+ctx.String(extraArgsFlag.Name), filesDir, string(genesisJson), binary, nil); err != nil {
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
