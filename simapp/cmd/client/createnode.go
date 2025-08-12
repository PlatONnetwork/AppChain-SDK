package main

import (
	"crypto/ecdsa"
	deploytools "github.com/PlatONnetwork/AppChain-SDK/tools/deploy"
	"github.com/PlatONnetwork/PlatON-Go/accounts/keystore"
	"github.com/PlatONnetwork/PlatON-Go/cmd/utils"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/google/uuid"
	"gopkg.in/urfave/cli.v1"
	"os"
)

var (
	DefaultPassword = "123456"
)

type ExtraParams struct {
	CheckpointKeystoreAddress string `toml:"checkpoint_keystore_address"                     comment:"Checkpoint keystore address, send a checkpoint tx to rootchain"`

	CheckpointKeystore string `toml:"checkpoint_keystore"                     comment:"Checkpoint keystore, send a checkpoint tx to rootchain"`
	CheckpointPassword string `toml:"checkpoint_keystore_password"                     comment:"Checkpoint keystore password"`
	L2KeystoreAddress  string `toml:"l2_keystore_address"                     comment:"Layer2 keystore address, send a checkpoint tx to rootchain"`

	L2Keystore string `toml:"l2_keystore"                     comment:"Layer2 keystore, send a checkpoint tx to rootchain"`
	L2Password string `toml:"l2_keystore_password"                     comment:"Layer2 keystore, send a internal tx to layer2"`
}

func createNodeCommand() *cli.Command {
	cmd := deploytools.CreateNodeCommand
	cmd.Action = createnode
	cmd.Flags = append(cmd.Flags, []cli.Flag{
		CheckpointKeystoreFlag,
		CheckpointKeystorePasswordFlag,
		L2TxKeystoreFlag,
		L2TxKeystorePasswordFlag,
	}...)
	return &cmd
}

func createnode(ctx *cli.Context) error {
	var checkpointKey *ecdsa.PrivateKey
	var checkpointKeyStore []byte
	var checkpointPassword []byte
	var l2key *ecdsa.PrivateKey
	var l2keystore []byte
	var l2keystorePassword []byte
	var err error
	if len(ctx.String(CheckpointKeystoreFlag.Name)) != 0 {
		if checkpointKeyStore, err = os.ReadFile(ctx.String(CheckpointKeystoreFlag.Name)); err != nil {
			log.Error("Read file failed", "path", ctx.String(CheckpointKeystoreFlag.Name))
			return err
		}

		if checkpointPassword, err = os.ReadFile(ctx.String(CheckpointKeystorePasswordFlag.Name)); err != nil {
			log.Error("Read file failed", "path", ctx.String(CheckpointKeystorePasswordFlag.Name))
			return err
		}
		if checkpointKey, err = decodeKey(checkpointKeyStore, checkpointPassword); err != nil {
			log.Error("Decode checkpoint key failed", "err", err)
			return err
		}
	}
	if len(ctx.String(L2TxKeystoreFlag.Name)) != 0 {

		if l2keystore, err = os.ReadFile(ctx.String(L2TxKeystoreFlag.Name)); err != nil {
			log.Error("Read file failed", "path", ctx.String(L2TxKeystoreFlag.Name))
			return err
		}

		if l2keystorePassword, err = os.ReadFile(ctx.String(L2TxKeystorePasswordFlag.Name)); err != nil {
			log.Error("Read file failed", "path", ctx.String(L2TxKeystorePasswordFlag.Name))
			return err
		}
		if l2key, err = decodeKey(l2keystore, l2keystorePassword); err != nil {
			log.Error("Decode l2 key failed", "err", err)
			return err
		}
	}
	return deploytools.CreateNodeExtra(ctx, func(node *deploytools.Node) (*deploytools.Node, error) {

		extra := &ExtraParams{}
		if len(checkpointKeyStore) != 0 {

			extra.CheckpointKeystoreAddress = crypto.PubkeyToAddress(checkpointKey.PublicKey).Hex()
			extra.CheckpointKeystore = string(checkpointKeyStore)
			extra.CheckpointPassword = string(checkpointPassword)
		} else {
			extra.CheckpointKeystoreAddress, extra.CheckpointKeystore = createkey(DefaultPassword)
			extra.CheckpointPassword = DefaultPassword
		}
		if len(l2keystore) != 0 {
			extra.L2KeystoreAddress = crypto.PubkeyToAddress(l2key.PublicKey).Hex()

			extra.L2Keystore = string(l2keystore)
			extra.L2Password = string(l2keystorePassword)
		} else {
			extra.L2KeystoreAddress, extra.L2Keystore = createkey(DefaultPassword)
			extra.L2Password = DefaultPassword
		}
		node.Extra = extra
		return node, nil
	})
}

func createkey(passphrase string) (string, string) {
	privateKey, _ := crypto.GenerateKey()
	UUID, err := uuid.NewRandom()
	if err != nil {
		utils.Fatalf("Failed to generate random uuid: %v", err)
	}
	key := &keystore.Key{
		Id:         UUID,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}

	scryptN, scryptP := keystore.StandardScryptN, keystore.StandardScryptP
	keyjson, err := keystore.EncryptKey(key, passphrase, scryptN, scryptP)
	if err != nil {
		utils.Fatalf("Error encrypting key: %v", err)
	}
	return key.Address.Hex(), string(keyjson)
}

func decodeKey(keyjson, passphrase []byte) (*ecdsa.PrivateKey, error) {
	key, err := keystore.DecryptKey(keyjson, string(passphrase))
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, err
}
