package deploy

import (
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/pelletier/go-toml/v2"
	"os"
)

type Node struct {
	Genesis                  bool        `toml:"genesis" comment:"Is it a genesis node"`
	IP                       string      `toml:"ip" comment:"IP"`
	PprofPort                int         `toml:"pprof_port" comment:"Pprof port"`
	P2pPort                  int         `toml:"p2p_port" comment:"P2p port"`
	RpcPort                  int         `toml:"rpc_port" comment:"Rpc port"`
	NodeAddress              string      `toml:"node_address" comment:"Address of node"`
	NodePrivateKey           string      `toml:"node_private_key" comment:"Private key of node"`
	NodePublicKey            string      `toml:"node_public_key" comment:"Public key of node"`
	BlsPrivateKey            string      `toml:"bls_private_key" comment:"Bls private key for consensus"`
	BlsPublicKey             string      `toml:"bls_public_key" comment:"Bls public key for consensus"`
	BlsUncompressedPublicKey string      `toml:"bls_uncompressed_key" comment:"Bls public key for root chain contract"`
	Extra                    interface{} `toml:"extra" comment:"Extra Parameters"`
}

type NodeConfig struct {
	Nodes []*Node `toml:"nodes"`
}

func ReadNodeConfig(file string) (*NodeConfig, error) {
	var config NodeConfig
	data, err := os.ReadFile(file)
	if err != nil {
		log.Error("Read childchain config failed", "err", err)
		return nil, err
	}

	if err := toml.Unmarshal(data, &config); err != nil {
		log.Error("Decode childchain config failed", "err", err)
		return nil, err
	}

	return &config, nil
}
