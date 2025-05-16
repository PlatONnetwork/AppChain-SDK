package benchmark

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/hexutil"
	"math/big"
)

var recipientBase = hexutil.MustDecode("0x1100000000000000000000000000000000000001")
var contractBase = hexutil.MustDecode("0x1200000000000000000000000000000000000001")

func GenRecipient(index uint64) common.Address {
	recipient := big.NewInt(0).SetBytes(recipientBase)
	recipient.Add(recipient, big.NewInt(int64(index)))
	return common.BigToAddress(recipient)
}

func GenContract(index uint64) common.Address {
	recipient := big.NewInt(0).SetBytes(contractBase)
	recipient.Add(recipient, big.NewInt(int64(index)))
	return common.BigToAddress(recipient)
}
