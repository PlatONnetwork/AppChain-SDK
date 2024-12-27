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
	LotteryTxBuilderABI = "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"vrfStore\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"drawing\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"guessing\",\"inputs\":[{\"name\":\"number\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Drawing\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"bonus\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Guessing\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"number\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]"
)

type LotteryTxBuilder struct {
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

func NewLotteryTxBuilder(to common.Address, sk *ecdsa.PrivateKey, chainId *big.Int) (*LotteryTxBuilder, error) {
	abi, err := abi.JSON(strings.NewReader(LotteryTxBuilderABI))
	if err != nil {
		return nil, err
	}
	return &LotteryTxBuilder{
		abi:      &abi,
		sk:       sk,
		chainId:  chainId,
		signer:   types.NewEIP155Signer(chainId),
		gasLimit: 1000000,
		gasPrice: big.NewInt(0),
		to:       to,
	}, nil
}

func (c *LotteryTxBuilder) WithPrivateKey(sk *ecdsa.PrivateKey) *LotteryTxBuilder {
	c.sk = sk
	return c
}
func (c *LotteryTxBuilder) WithNonce(nonce uint64) *LotteryTxBuilder {
	c.nonce = nonce
	return c
}
func (c *LotteryTxBuilder) WithChainId(chainId *big.Int) *LotteryTxBuilder {
	c.chainId = chainId
	return c
}

func (c *LotteryTxBuilder) WithSigner(signer types.Signer) *LotteryTxBuilder {
	c.signer = signer
	return c
}

func (c *LotteryTxBuilder) WithGasLimit(gasLimit uint64) *LotteryTxBuilder {
	c.gasLimit = gasLimit
	return c
}

func (c *LotteryTxBuilder) WithGasPrice(gasPrice *big.Int) *LotteryTxBuilder {
	c.gasPrice = gasPrice
	return c
}

func (c *LotteryTxBuilder) WithTo(to common.Address) *LotteryTxBuilder {
	c.to = to
	return c
}

func (c *LotteryTxBuilder) PackAllowance(owner common.Address, spender common.Address) ([]byte, error) {
	input, err := c.abi.Pack("allowance", owner, spender)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackAllowance(owner, spender)
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

func (c *LotteryTxBuilder) PackBalanceOf(account common.Address) ([]byte, error) {
	input, err := c.abi.Pack("balanceOf", account)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) BalanceOf(account common.Address) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackBalanceOf(account)
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

func (c *LotteryTxBuilder) PackDecimals() ([]byte, error) {
	input, err := c.abi.Pack("decimals")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Decimals() (*types.Transaction, error) {

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

func (c *LotteryTxBuilder) PackName() ([]byte, error) {
	input, err := c.abi.Pack("name")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Name() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackName()
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

func (c *LotteryTxBuilder) PackSymbol() ([]byte, error) {
	input, err := c.abi.Pack("symbol")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Symbol() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackSymbol()
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

func (c *LotteryTxBuilder) PackTotalSupply() ([]byte, error) {
	input, err := c.abi.Pack("totalSupply")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) TotalSupply() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackTotalSupply()
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

func (c *LotteryTxBuilder) PackApprove(spender common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("approve", spender, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackApprove(spender, amount)
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

func (c *LotteryTxBuilder) PackDecreaseAllowance(spender common.Address, subtractedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("decreaseAllowance", spender, subtractedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackDecreaseAllowance(spender, subtractedValue)
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

func (c *LotteryTxBuilder) PackDrawing() ([]byte, error) {
	input, err := c.abi.Pack("drawing")
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Drawing() (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackDrawing()
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

func (c *LotteryTxBuilder) PackGuessing(number *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("guessing", number)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Guessing(number *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackGuessing(number)
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

func (c *LotteryTxBuilder) PackIncreaseAllowance(spender common.Address, addedValue *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("increaseAllowance", spender, addedValue)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackIncreaseAllowance(spender, addedValue)
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

func (c *LotteryTxBuilder) PackTransfer(to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transfer", to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackTransfer(to, amount)
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

func (c *LotteryTxBuilder) PackTransferFrom(from common.Address, to common.Address, amount *big.Int) ([]byte, error) {
	input, err := c.abi.Pack("transferFrom", from, to, amount)
	if err != nil {
		return nil, err
	}
	return input, nil
}
func (c *LotteryTxBuilder) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {

	var err error
	var input []byte
	input, err = c.PackTransferFrom(from, to, amount)
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

func (c *LotteryTxBuilder) ABI() *abi.ABI {
	return c.abi
}
