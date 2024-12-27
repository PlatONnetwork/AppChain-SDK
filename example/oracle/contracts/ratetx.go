package contracts

import (
	"crypto/ecdsa"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"

	"math/big"
	"strings"
)

var (
	RateTxBuilderABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_electionAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blockNumber\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"rate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"update\",\"inputs\":[{\"name\":\"newRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"qc\",\"type\":\"tuple\",\"internalType\":\"structQuorumCert\",\"components\":[{\"name\":\"epoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"viewNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockIndex\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"extendHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"bitmap\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"leafIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"UpdateRate\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"rate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]"
)

type RateTxBuilder struct {
	abi      *abi.ABI
	sk       *ecdsa.PrivateKey
	nonce    uint64
	chainId  *big.Int
	signer   types.Signer
	gasLimit uint64
	gasPrice *big.Int
	to       common.Address
	value    *big.Int
}

func NewRateTxBuilder(to common.Address, sk *ecdsa.PrivateKey, chainId *big.Int) (*RateTxBuilder, error) {
	abi, err := abi.JSON(strings.NewReader(RateTxBuilderABI))
	if err != nil {
		return nil, err
	}
	return &RateTxBuilder{
		abi:      &abi,
		sk:       sk,
		chainId:  chainId,
		signer:   types.NewEIP155Signer(chainId),
		gasLimit: 1000000,
		gasPrice: big.NewInt(0),
		to:       to,
	}, nil
}

func (c *RateTxBuilder) WithPrivateKey(sk *ecdsa.PrivateKey) *RateTxBuilder {
	c.sk = sk
	return c
}
func (c *RateTxBuilder) WithNonce(nonce uint64) *RateTxBuilder {
	c.nonce = nonce
	return c
}
func (c *RateTxBuilder) WithChainId(chainId *big.Int) *RateTxBuilder {
	c.chainId = chainId
	return c
}

func (c *RateTxBuilder) WithSigner(signer types.Signer) *RateTxBuilder {
	c.signer = signer
	return c
}

func (c *RateTxBuilder) WithGasLimit(gasLimit uint64) *RateTxBuilder {
	c.gasLimit = gasLimit
	return c
}

func (c *RateTxBuilder) WithGasPrice(gasPrice *big.Int) *RateTxBuilder {
	c.gasPrice = gasPrice
	return c
}

func (c *RateTxBuilder) WithTo(to common.Address) *RateTxBuilder {
	c.to = to
	return c
}

func (c *RateTxBuilder) PackBlockNumber() ([]byte, error) {
	input, err := c.abi.Pack("blockNumber")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *RateTxBuilder) BlockNumber() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackBlockNumber()
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(c.nonce, c.to, c.value, c.gasLimit, c.gasPrice, input)
	tx, err = types.SignTx(tx, c.signer, c.sk)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *RateTxBuilder) PackDecimals() ([]byte, error) {
	input, err := c.abi.Pack("decimals")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *RateTxBuilder) Decimals() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackDecimals()
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(c.nonce, c.to, c.value, c.gasLimit, c.gasPrice, input)
	tx, err = types.SignTx(tx, c.signer, c.sk)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *RateTxBuilder) PackRate() ([]byte, error) {
	input, err := c.abi.Pack("rate")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *RateTxBuilder) Rate() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackRate()
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(c.nonce, c.to, c.value, c.gasLimit, c.gasPrice, input)
	tx, err = types.SignTx(tx, c.signer, c.sk)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *RateTxBuilder) PackUpdate(newRate *big.Int, qc QuorumCert, bitmap []byte, signature []byte, leafIndex *big.Int, proof []common.Hash) ([]byte, error) {
	input, err := c.abi.Pack("update", newRate, qc, bitmap, signature, leafIndex, proof)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *RateTxBuilder) Update(newRate *big.Int, qc QuorumCert, bitmap []byte, signature []byte, leafIndex *big.Int, proof []common.Hash) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackUpdate(newRate, qc, bitmap, signature, leafIndex, proof)
	if err != nil {
		return nil, err
	}
	tx := types.NewTransaction(c.nonce, c.to, c.value, c.gasLimit, c.gasPrice, input)
	tx, err = types.SignTx(tx, c.signer, c.sk)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (c *RateTxBuilder) ABI() *abi.ABI {
	return c.abi
}
