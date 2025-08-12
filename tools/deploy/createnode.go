package deploy

import (
	"encoding/hex"
	"github.com/PlatONnetwork/AppChain-SDK/tools/tests/cast/flags"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/pelletier/go-toml/v2"
	"gopkg.in/urfave/cli.v1"
	"os"
	"path/filepath"
	"strings"
)

var (
	CreateNodeCommand = cli.Command{
		Name:   "createnode",
		Action: createNode,
		Flags: []cli.Flag{
			GenesisNodeFlag,
			NormalNodeFlag,
			OutputFlag,
			NodeConfigFileFlag,
			NodePProfPortFlag,
			NodeRpcPortFlag,
			NodeP2pPortFlag,
			NodePortIncFlag,
		},
		CustomHelpTemplate: flags.CommandHelpTemplate,
	}
)

func createNode(ctx *cli.Context) error {
	return CreateNodeExtra(ctx, nil)
}
func CreateNodeExtra(ctx *cli.Context, extra func(*Node) (*Node, error)) error {
	genesisHosts := strings.Split(ctx.String(GenesisNodeFlag.Name), ",")
	normalHosts := strings.Split(ctx.String(NormalNodeFlag.Name), ",")
	pprofPort := ctx.Int(NodePProfPortFlag.Name)
	p2pPort := ctx.Int(NodeP2pPortFlag.Name)
	rpcPort := ctx.Int(NodeRpcPortFlag.Name)
	increment := ctx.Bool(NodePortIncFlag.Name)
	var conf NodeConfig
	genNode := func(host string, genesis bool) *Node {
		nodeAddress, nodePrivateKey, nodePubKey := genNodeKey()
		blsPrivateKey, blsPubKey, blsUncompressedPubKey := genBlskey()
		node := &Node{
			Genesis:                  genesis,
			IP:                       host,
			PprofPort:                pprofPort,
			P2pPort:                  p2pPort,
			RpcPort:                  rpcPort,
			NodeAddress:              nodeAddress,
			NodePrivateKey:           nodePrivateKey,
			NodePublicKey:            nodePubKey,
			BlsPrivateKey:            blsPrivateKey,
			BlsPublicKey:             blsPubKey,
			BlsUncompressedPublicKey: blsUncompressedPubKey,
		}
		if extra != nil {
			var err error
			node, err = extra(node)
			if err != nil {
				log.Crit("Create node extra params failed", "err", err)
			}
		}
		return node
	}
	for _, host := range genesisHosts {
		if len(host) > 0 {
			conf.Nodes = append(conf.Nodes, genNode(host, true))
			if increment {
				pprofPort, p2pPort, rpcPort = pprofPort+1, p2pPort+1, rpcPort+1
			}
		}
	}
	for _, host := range normalHosts {
		if len(host) > 0 {
			conf.Nodes = append(conf.Nodes, genNode(host, false))
			if increment {
				pprofPort, p2pPort, rpcPort = pprofPort+1, p2pPort+1, rpcPort+1
			}
		}
	}
	data, err := toml.Marshal(conf)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(ctx.String(OutputFlag.Name), ctx.String(NodeConfigFileFlag.Name)), data, 0666); err != nil {
		log.Error("Write config file failed", "err", err)
	}

	return nil
}

func genNodeKey() (string, string, string) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex(), hex.EncodeToString(crypto.FromECDSA(privateKey)), hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)[1:])
}

func genBlskey() (string, string, string) {
	var privateKey bls.SecretKey
	privateKey.SetByCSPRNG()
	pubKey := privateKey.GetPublicKey()
	return hex.EncodeToString(privateKey.Serialize()), hex.EncodeToString(pubKey.Serialize()), hex.EncodeToString(pubKey.SerializeUncompressed())

}
