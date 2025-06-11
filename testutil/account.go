package testutil

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/core"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"math/big"
	"net"
)

type Account struct {
	BlsKey      string
	NodeKey     string
	CoinbaseKey string
	Host        net.IP
	P2PPort     int
	HTTP        int
	Pprof       int
}

func (a *Account) BlsSecretKey() *bls.SecretKey {
	var blsKey bls.SecretKey
	keyBuf, err := hex.DecodeString(a.BlsKey)
	if err != nil {
		panic(err)
	}
	if err := blsKey.Deserialize(keyBuf); err != nil {
		panic(err)
	}

	return &blsKey
}
func (a *Account) NodePrivateKey() *ecdsa.PrivateKey {
	nodePrivate, err := crypto.HexToECDSA(a.NodeKey)
	if err != nil {
		panic(err)
	}
	return nodePrivate
}

func (a *Account) NodeAddress() common.Address {
	return crypto.PubkeyToAddress(a.NodePrivateKey().PublicKey)
}

func (a *Account) CoinbasePrivateKey() *ecdsa.PrivateKey {
	coinbasePrivate, err := crypto.HexToECDSA(a.CoinbaseKey)
	if err != nil {
		panic(err)
	}
	return coinbasePrivate
}

func (a *Account) CoinBaseAddress() common.Address {
	return crypto.PubkeyToAddress(a.CoinbasePrivateKey().PublicKey)
}

func (a *Account) Node() *enode.Node {
	node := enode.NewV4(&a.NodePrivateKey().PublicKey, a.Host, a.P2PPort, a.P2PPort)
	return node
}

var DefaultAccount = []*Account{
	{
		BlsKey:      "045b76979077ceb756073ace213d0a76e11b4c4093dc07439bda23010dd67682",
		NodeKey:     "db8c4595889e472b4cba14f1430c8e45bfadcd5d81c49f9696432311c2d63702",
		CoinbaseKey: "db8c4595889e472b4cba14f1430c8e45bfadcd5d81c49f9696432311c2d63702",
		Host:        net.IP{127, 0, 0, 1},
		P2PPort:     18001,
		HTTP:        8801,
		Pprof:       7001,
	}, {
		BlsKey:      "225e0e20da3b2f9556ccbc3c6031fceede4078b6aed7025ed8a465767d5d79b8",
		NodeKey:     "5cd44b68d91f4b42dc240c658fee14922a0e26245ed6f406c0015bb50cf66431",
		CoinbaseKey: "5cd44b68d91f4b42dc240c658fee14922a0e26245ed6f406c0015bb50cf66431",
		Host:        net.IP{127, 0, 0, 1},
		P2PPort:     18002,
		HTTP:        8802,
		Pprof:       7002,
	}, {
		BlsKey:      "1be36ecdc012e37a6bf682de21e4d9de2071ca9b94362b4b20a6b89f56ba6a99",
		NodeKey:     "d00af106a861e19cda22141b09e402acf869f3b6e366363625f4b4cabe9323d2",
		CoinbaseKey: "d00af106a861e19cda22141b09e402acf869f3b6e366363625f4b4cabe9323d2",
		Host:        net.IP{127, 0, 0, 1},
		P2PPort:     18003,
		HTTP:        8803,
		Pprof:       7003,
	}, {
		BlsKey:      "3538a0bd63d0b9cbf51ed4e107f8337815a6efa69383c24aa50f2a7b3de52a96",
		NodeKey:     "8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd",
		CoinbaseKey: "8d592db252683a647031a77f95d9f0ed4115cdd6aefe1e4cb98d03efe2f524dd",
		Host:        net.IP{127, 0, 0, 1},
		P2PPort:     18004,
		HTTP:        8804,
		Pprof:       7004,
	},
}

var DefaultGenesisTemplate = core.Genesis{
	Config: &params.ChainConfig{
		ChainID:         big.NewInt(123083),
		AddressHRP:      "lat",
		EmptyBlock:      "",
		EIP155Block:     big.NewInt(0),
		EWASMBlock:      big.NewInt(0),
		CopernicusBlock: big.NewInt(0),
		NewtonBlock:     big.NewInt(0),
		EinsteinBlock:   big.NewInt(0),
		HubbleBlock:     big.NewInt(0),
		PauliBlock:      big.NewInt(0),
		Cbft: &params.CbftConfig{
			Period:        10000,
			Amount:        10,
			InitialNodes:  nil,
			ValidatorMode: "",
		},
		GenesisVersion: 65792,
		Modules:        nil,
	},
	Nonce:      hexutil.MustDecode("0x0376e56dffd12ab53bb149bda4e0cbce2b6aabe4cccc0df0b5a39e12977a2fcd23"),
	Timestamp:  1602973620000,
	ExtraData:  hexutil.MustDecode("0x22cf84e1bf86cf8220e1bc84cebdcf8920e1bd81ceb4cebfe1bfa620e1bc80ceb5e1bdb620e1bc91cebecf8ccebcceb5ceb8ceb120cebaceb1e1bdb620ceb4ceb9cebaceb1ceb9cebfcf83cf8dcebdceb7cebd20cebcceb5cf84e1bdb020cf86cf81cebfcebdceaecf83ceb5cf89cf8220cf80ceb1cebdcf84e1bdb620cf84cf81cf8ccf80e1bfb320e1bc90cf80ceb9cf84ceb7ceb4ceb5cf8dcf83cebfcebcceb5cebd2e220a225765207368616c6c20616c77617973206b65657020746f2074686520757070657220726f616420616e64207072616374696365206a75737469636520776974682070727564656e636520696e206576657279207761792e220a22e8aea9e68891e4bbace6b0b8e8bf9ce59d9ae68c81e8b5b0e59091e4b88ae79a84e8b7afefbc8ce585a8e58a9be4bba5e5aea1e6858ee8b7b5e8a18ce6ada3e4b989e3808222"),
	GasLimit:   1201600000,
	Coinbase:   common.Address{},
	Alloc:      nil,
	Number:     0,
	GasUsed:    0,
	ParentHash: common.Hash{},
}

func GenerateGenesis(accounts []*Account, modules map[string]json.RawMessage, users []common.Address) ([]byte, error) {
	conf := DefaultGenesisTemplate
	//init cbft node
	var cbftNode []params.CbftNode
	alloc := make(map[common.Address]core.GenesisAccount)
	for _, acc := range accounts {
		var blsKey bls.SecretKey
		keyBuf, _ := hex.DecodeString(acc.BlsKey)
		blsKey.Deserialize(keyBuf)
		nodePrivate, err := crypto.HexToECDSA(acc.NodeKey)
		if err != nil {
			return nil, err
		}
		node := enode.NewV4(&nodePrivate.PublicKey, acc.Host, acc.P2PPort, acc.P2PPort)
		nodeAddr := crypto.PubkeyToAddress(nodePrivate.PublicKey)
		balance, _ := new(big.Int).SetString("2000000000000000000000000000000000000", 16)
		alloc[nodeAddr] = core.GenesisAccount{
			Balance: balance,
			Nonce:   0,
		}
		cbftNode = append(cbftNode, params.CbftNode{
			Node:      node,
			BlsPubKey: *blsKey.GetPublicKey(),
		})
	}
	for _, acc := range DefaultAccount {
		nodePrivate, err := crypto.HexToECDSA(acc.NodeKey)
		if err != nil {
			return nil, err
		}
		nodeAddr := crypto.PubkeyToAddress(nodePrivate.PublicKey)
		balance, _ := new(big.Int).SetString("2000000000000000000000000000000000000", 16)
		alloc[nodeAddr] = core.GenesisAccount{
			Balance: balance,
			Nonce:   0,
		}
	}
	for _, ac := range users {
		balance, _ := new(big.Int).SetString("2000000000000000000000000000000000000", 16)
		alloc[ac] = core.GenesisAccount{
			Balance: balance,
			Nonce:   0,
		}
	}
	conf.Config.Cbft.InitialNodes = cbftNode
	conf.Alloc = alloc
	conf.Config.Modules = modules
	return json.MarshalIndent(conf, "", "")
}
