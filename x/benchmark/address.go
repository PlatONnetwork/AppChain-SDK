package benchmark

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"math/big"
	"sync"
)

var recipientBase = hexutil.MustDecode("0x1100000000000000000000000000000000000001")
var contractBase = hexutil.MustDecode("0x1200000000000000000000000000000000000001")
var recipientAddrCache sync.Map
var contractAddrCache sync.Map

func GenRecipient(index uint64) common.Address {
	if addr, ok := recipientAddrCache.Load(index); ok {
		return addr.(common.Address)
	}
	recipient := big.NewInt(0).SetBytes(recipientBase)
	recipient.Add(recipient, big.NewInt(int64(index)))

	addr := common.BytesToAddress(crypto.Keccak256(recipient.Bytes()))
	recipientAddrCache.Store(index, addr)
	return addr
}

func GenContract(index uint64) common.Address {
	if addr, ok := contractAddrCache.Load(index); ok {
		return addr.(common.Address)
	}
	recipient := big.NewInt(0).SetBytes(contractBase)
	recipient.Add(recipient, big.NewInt(int64(index)))
	addr := common.BytesToAddress(crypto.Keccak256(recipient.Bytes()))
	contractAddrCache.Store(index, addr)
	return addr
}
