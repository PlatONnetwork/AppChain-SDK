package contracts

import (
	"crypto/ecdsa"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/umbracle/ethgo/abi"
	"math/big"
)

const (
	METHODID_SIZE = 32
)

var (
	_STAKE_SIG      = crypto.Keccak256Hash([]byte("STAKE"))
	_ADDSTAKE_SIG   = crypto.Keccak256Hash([]byte("ADDSTAKE"))
	_UNSTAKE_SIG    = crypto.Keccak256Hash([]byte("UNSTAKE"))
	_SLASH_SIG      = crypto.Keccak256Hash([]byte("SLASH"))
	_DELEGATE_SIG   = crypto.Keccak256Hash([]byte("DELEGATE"))
	_UNDELEGATE_SIG = crypto.Keccak256Hash([]byte("UNDELEGATE"))
)

var (
	_STAKE_PARAMS_TYPE = abi.MustNewType("tuple(address validatorAddr, address benefitAddr, uint256 amount, uint256[2] bksKey, bytes pubKey)")
	_ADDSTAKE_TYPE     = abi.MustNewType("tuple(address validatorAddr, uint256 amount)")
	_UNSTAKE_TYPE      = abi.MustNewType("tuple(address validatorAddr, uint256 amount)")
	_SLASH_TYPE        = abi.MustNewType("tuple(address[] validatorAddrs, uint256 slashingPercentage, uint256 slashIncentivePercentage)")
	_DELEGATE_TYPE     = abi.MustNewType("tuple(address validatorAddr, address delegterAddr, uint256 amount)")
	_UNDELEGATE_TYPE   = abi.MustNewType("tuple(address validatorAddr, address delegterAddr, uint256 amount)")
)

func (c *StakeHandler) onStake(input []byte) error {
	decoded, err := abi.Decode(_STAKE_PARAMS_TYPE, input)
	if nil != err {
		return typesdk.NewRevertError("StakeHandler: DECODE_STAKE_DATA_FAILED")
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_STAKE_DATA")
	}

	validatorAddr, ok := res["validatorAddr"].(common.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	benefitAddr, ok := res["benefitAddr"].(common.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_BENEFIT")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNT")
	}

	// TODO parse blsKey
	//blsKey, ok := res["blsKey"].([2]*big.Int)
	//if !ok {
	//	return typesdk.NewRevertError("StakeHandler: INVALID_BLSKEY")
	//}
	//
	pubKey, ok := res["pubKey"].([]byte)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_PUBKEY")
	}

	publicKey, err := crypto.UnmarshalPubkey(pubKey)
	if nil != err {
		log.Error("Failed to unmarshal publicKey", "error", err)
		return typesdk.NewRevertError("StakeHandler: INVALID_PUBKEY")
	}
	return c.stake(validatorAddr, benefitAddr, amount, nil, publicKey) // todo need blsKey
}

func (c *StakeHandler) stake(validatorAddr, benefit common.Address, amount *big.Int, blsKey *bls.PublicKey, pubKey *ecdsa.PublicKey) error {
	if c.hasValidator(validatorAddr) {
		return typesdk.NewRevertError("StakeHandler: VALIDATOR ALREADY STAKE")
	}
	// TODO need check min stake threshold

	if err := c.setValidator(validatorAddr, types.NewValidator(benefit, amount, common.Big0, blsKey, pubKey)); nil != err {
		log.Error("Failed to set validator stake", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: STAKE FAILED")
	}
	return nil
}
