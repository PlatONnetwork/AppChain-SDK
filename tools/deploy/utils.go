package deploy

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"math/big"
	"os/exec"
)

func MustStringToAddress(str string) common.Address {
	addr, err := common.StringToAddress(str)
	if err != nil {
		panic(fmt.Sprintf("decode address failed, [%s]", str))
	}
	return addr
}

func MustDecodeString(str string) []byte {
	data, err := hex.DecodeString(str)
	if err != nil {
		panic(fmt.Sprintf("decode hex failed, [%s]", str))
	}
	return data
}

func MustDecodePubKey(str string) *ecdsa.PublicKey {
	pubKey := MustDecodeString(str)
	var pubKey65 []byte
	switch len(pubKey) {
	case 64:
		// add 'uncompressed key' flag
		pubKey65 = append([]byte{0x04}, pubKey...)
	case 65:
		pubKey65 = pubKey
	default:
		panic(fmt.Sprintf("invalid public key length %v (expect 64/65)", len(pubKey)))
	}
	pubkey, err := crypto.UnmarshalPubkey(pubKey65)
	if err != nil {
		panic(fmt.Sprintf("decode public failed, [%s], err;%s", str, err.Error()))
	}
	return pubkey
}
func CopyFile(src, dst string) error {
	log.Debug("Copy file", "src", src, "dst", dst)
	cpCmd := exec.Command("cp", "-f", src, dst)
	return cpCmd.Run()
}

func MustStringToDecimal(s string) *big.Int {
	n, flag := new(big.Int).SetString(s, 10)
	if !flag {
		log.Crit("Decode decimal failed", "value", s)
	}
	return n
}
