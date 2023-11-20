package contracts

import (
	"crypto/ecdsa"
	"encoding/hex"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"github.com/PlatONnetwork/PlatON-Go/log"
	"github.com/umbracle/ethgo"
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
	_STAKE_PARAMS_TYPE             = abi.MustNewType("tuple(address validatorAddr, address benefitAddr, uint256 amount, uint256[2] bksKey, bytes pubKey)")
	_ADDSTAKE_PARAMS_TYPE          = abi.MustNewType("tuple(address validatorAddr, uint256 amount)")
	_UNSTAKE_PARAMS_TYPE           = abi.MustNewType("tuple(address validatorAddr, uint256 amount)")
	_ROOT_CHAIN_SLASH_PARAMS_TYPE  = abi.MustNewType("tuple(address[] validatorAddrs, uint256 slashingPercentage, uint256 slashIncentivePercentage)")
	_CHILD_CHAIN_SLASH_PARAMS_TYPE = abi.MustNewType("tuple(uint256 handleEventId, []address validatorAddrs)")
	_DELEGATE_PARAMS_TYPE          = abi.MustNewType("tuple(address validatorAddr, address delegterAddr, uint256 amount)")
	_UNDELEGATE_PARAMS_TYPE        = abi.MustNewType("tuple(address validatorAddr, address delegterAddr, uint256 amount)")
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

	validatorAddr, ok := res["validatorAddr"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	benefitAddr, ok := res["benefitAddr"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_BENEFIT")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNT")
	}

	blsKeyArr, ok := res["blsKey"].([2]*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_BLSKEY")
	}
	blsKeyBytes := append(blsKeyArr[0].Bytes(), blsKeyArr[1].Bytes()...)
	blsKey := bls.PublicKey{}
	(&blsKey).Deserialize(blsKeyBytes)

	pubKey, ok := res["pubKey"].([]byte)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_PUBKEY")
	}

	publicKey, err := crypto.UnmarshalPubkey(pubKey)
	if nil != err {
		log.Error("Failed to unmarshal publicKey", "error", err)
		return typesdk.NewRevertError("StakeHandler: INVALID_PUBKEY")
	}
	return c.stake(common.Address(validatorAddr), common.Address(benefitAddr), amount, &blsKey, publicKey)
}

func (c *StakeHandler) onAddStake(input []byte) error {
	decoded, err := abi.Decode(_ADDSTAKE_PARAMS_TYPE, input)
	if nil != err {
		return typesdk.NewRevertError("StakeHandler: DECODE_ADD_STAKE_DATA_FAILED")
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_ADD_STAKE_DATA")
	}

	validatorAddr, ok := res["validatorAddr"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNT")
	}

	return c.addStake(common.Address(validatorAddr), amount)
}

func (c *StakeHandler) onSlash(input []byte) error {
	decoded, err := abi.Decode(_CHILD_CHAIN_SLASH_PARAMS_TYPE, input)
	if nil != err {
		return typesdk.NewRevertError("StakeHandler: DECODE_SLASH_DATA_FAILED")
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_SLASH_DATA")
	}

	handleEventId, ok := res["handleEventId"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_HANDLEEVENTID")
	}

	validatorAddrs, ok := res["validatorAddrs"].([]ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATORADDRS")
	}
	addrs := make([]common.Address, len(validatorAddrs))

	for i, v := range validatorAddrs {
		addrs[i] = common.Address(v)
	}

	return c.slash(handleEventId, addrs)
}

func (c *StakeHandler) onDelegate(input []byte) error {
	decoded, err := abi.Decode(_DELEGATE_PARAMS_TYPE, input)
	if nil != err {
		return typesdk.NewRevertError("StakeHandler: DECODE_DELEGATE_DATA_FAILED")
	}
	res, ok := decoded.(map[string]interface{})
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_DELEGATE_DATA")
	}

	validatorAddr, ok := res["validatorAddr"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	delegterAddr, ok := res["delegterAddr"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_DELEGTERADDR")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNT")
	}

	return c.delegate(common.Address(validatorAddr), common.Address(delegterAddr), amount)
}

func (c *StakeHandler) stake(validatorAddr, benefit common.Address, amount *big.Int, blsKey *bls.PublicKey, pubKey *ecdsa.PublicKey) error {
	if c.hasValidator(validatorAddr) {
		return typesdk.NewRevertError("StakeHandler: VALIDATOR ALREADY STAKE")
	}

	if amount.Cmp(MIN_STAKE) < 0 {
		return typesdk.NewRevertError("StakeHandler: NOT ENOUGH STAKE")
	}

	blockNumber := c.evm.Context.BlockNumber.Uint64()
	stakeIndex := c.incrementValidatorNonce()

	if err := c.setValidatorByPriority(validatorAddr, types.NewValidator(benefit, amount, common.Big0, blsKey, pubKey, blockNumber, stakeIndex)); nil != err {
		log.Error("Failed to set validator stake", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: STAKE FAILED")
	}
	log.Info("Stake for", "validator", validatorAddr.Hex(), "benefit", benefit.Hex(), "amount", amount, "blsKey", string(blsKey.Bytes()),
		"pubKey", hex.EncodeToString(crypto.FromECDSAPub(pubKey)), "blockNumber", blockNumber, "stakeIndex", stakeIndex)
	return nil
}

func (c *StakeHandler) addStake(validatorAddr common.Address, amount *big.Int) error {
	if c.hasNotValidator(validatorAddr) {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	validator := c.GetValidator(validatorAddr)
	if validator.IsInvalid() {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	// update validator priority
	validator.AddStakeAmount(amount)

	if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
		log.Error("Failed to add validator stake amount", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: ADD STAKE FAILED")
	}
	log.Info("AddStake for", "validator", validatorAddr.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) slash(handleEventId *big.Int, validatorAddrs []common.Address) error {
	if c.hasSlashProcessed(handleEventId) {
		return typesdk.NewRevertError("StakeHandler: SLASH_ALREADY_PROCESSED")
	}

	// TODO

	return nil
}

func (c *StakeHandler) delegate(validatorAddr, delegaterAddr common.Address, amount *big.Int) error {
	if c.hasNotValidator(validatorAddr) {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	validator := c.GetValidator(validatorAddr)
	if validator.IsInvalid() {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	if amount.Cmp(MIN_DELEGATE) < 0 {
		return typesdk.NewRevertError("StakeHandler: NOT ENOUGH DELEGATE")
	}

	// update validator priority
	validator.AddDelegateAmount(amount)

	if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
		log.Error("Failed to add validator delegate amount", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: ADD DELEGATE AMOUNT OF VALIDATOR FAILED")
	}
	// update delegation
	if err := c.updateDelegation(delegaterAddr, validatorAddr, types.NewDelegation(c.evm.Context.BlockNumber.Uint64(), amount)); nil != err {
		log.Error("Failed to set delegation", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: SET DELEGATION FAILED")
	}
	log.Info("Delegate for", "delegaterAddr", delegaterAddr.Hex(), "validator", validatorAddr.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) registerStakeWithdrawal(validatorAddr common.Address, amount *big.Int) error {
	currentEpoch := c.GetCurrentEpoch()
	releaseEpoch := currentEpoch + STAKE_WITHDRAWAL_WAIT_PERIOD
	if err := c.appendStakeWithdrawal(validatorAddr, releaseEpoch, amount); nil != err {
		log.Error("Failed to register stake withdraw", "validatorAddr", validatorAddr.Hex(),
			"currentEpoch", currentEpoch, "releaseEpoch", releaseEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: SET REGISTER STAKE WITHDRAW FAILED")
	}

	if err := c.addLogStakeWithdrawalRegisteredEvent(validatorAddr, amount); nil != err {
		return err
	}
	log.Info("Register stake withdrawal for", "validator", validatorAddr.Hex(), "releaseEpoch", releaseEpoch, "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) registerDelegateWithdrawal(delegater, validatorAddr common.Address, amount *big.Int) error {
	currentEpoch := c.GetCurrentEpoch()
	releaseEpoch := currentEpoch + DELEGATE_WITHDRAWAL_WAIT_PERIOD
	if err := c.appendDelegateWithdrawal(delegater, validatorAddr, releaseEpoch, amount); nil != err {
		log.Error("Failed to register delegate withdraw", "delegater", delegater.Hex(), "validatorAddr", validatorAddr.Hex(),
			"currentEpoch", currentEpoch, "releaseEpoch", releaseEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: SET REGISTER DELEGATE WITHDRAW FAILED")
	}

	if err := c.addLogDelegateWithdrawalRegisteredEvent(delegater, validatorAddr, amount); nil != err {
		return err
	}
	log.Info("Register delegate withdrawal for", "delegater", delegater.Hex(), "validator", validatorAddr.Hex(), "releaseEpoch", releaseEpoch, "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}
