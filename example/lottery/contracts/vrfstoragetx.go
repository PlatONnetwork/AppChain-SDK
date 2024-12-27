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
	VRFStorageTxBuilderABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addProve\",\"inputs\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"nonceProof\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getNonceProof\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonceProofs\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vrfAddr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AddProve\",\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]"
)

type VRFStorageTxBuilder struct {
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

func NewVRFStorageTxBuilder(to common.Address, sk *ecdsa.PrivateKey, chainId *big.Int) (*VRFStorageTxBuilder, error) {
	abi, err := abi.JSON(strings.NewReader(VRFStorageTxBuilderABI))
	if err != nil {
		return nil, err
	}
	return &VRFStorageTxBuilder{
		abi:      &abi,
		sk:       sk,
		chainId:  chainId,
		signer:   types.NewEIP155Signer(chainId),
		gasLimit: 1000000,
		gasPrice: big.NewInt(0),
		to:       to,
	}, nil
}

func (c *VRFStorageTxBuilder) WithPrivateKey(sk *ecdsa.PrivateKey) *VRFStorageTxBuilder {
	c.sk = sk
	return c
}
func (c *VRFStorageTxBuilder) WithNonce(nonce uint64) *VRFStorageTxBuilder {
	c.nonce = nonce
	return c
}
func (c *VRFStorageTxBuilder) WithChainId(chainId *big.Int) *VRFStorageTxBuilder {
	c.chainId = chainId
	return c
}

func (c *VRFStorageTxBuilder) WithSigner(signer types.Signer) *VRFStorageTxBuilder {
	c.signer = signer
	return c
}

func (c *VRFStorageTxBuilder) WithGasLimit(gasLimit uint64) *VRFStorageTxBuilder {
	c.gasLimit = gasLimit
	return c
}

func (c *VRFStorageTxBuilder) WithGasPrice(gasPrice *big.Int) *VRFStorageTxBuilder {
	c.gasPrice = gasPrice
	return c
}

func (c *VRFStorageTxBuilder) WithTo(to common.Address) *VRFStorageTxBuilder {
	c.to = to
	return c
}

func (c *VRFStorageTxBuilder) PackGetNonceProof(blockNumber *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("getNonceProof", blockNumber)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *VRFStorageTxBuilder) GetNonceProof(blockNumber *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackGetNonceProof(blockNumber)
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

func (c *VRFStorageTxBuilder) PackNonceProofs(arg0 *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("nonceProofs", arg0)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *VRFStorageTxBuilder) NonceProofs(arg0 *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackNonceProofs(arg0)
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

func (c *VRFStorageTxBuilder) PackVrfAddr() ([]byte, error) {
	input, err := c.abi.Pack("vrfAddr")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *VRFStorageTxBuilder) VrfAddr() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackVrfAddr()
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

func (c *VRFStorageTxBuilder) PackAddProve(pubKey []byte, nonceProof []byte) ([]byte, error) {
	input, err := c.abi.Pack("addProve", pubKey, nonceProof)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *VRFStorageTxBuilder) AddProve(pubKey []byte, nonceProof []byte) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackAddProve(pubKey, nonceProof)
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

func (c *VRFStorageTxBuilder) PackGetNonce(blockNumber *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("getNonce", blockNumber)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *VRFStorageTxBuilder) GetNonce(blockNumber *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackGetNonce(blockNumber)
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

func (c *VRFStorageTxBuilder) ABI() *abi.ABI {
	return c.abi
}
