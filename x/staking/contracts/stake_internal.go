package contracts

import (
	"crypto/ecdsa"
	"encoding/hex"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	statesenderC "github.com/PlatONnetwork/AppChain-SDK/x/statesender/contracts"
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
	STAKE_SIG      = crypto.Keccak256Hash([]byte("STAKE"))
	ADDSTAKE_SIG   = crypto.Keccak256Hash([]byte("ADDSTAKE"))
	UNSTAKE_SIG    = crypto.Keccak256Hash([]byte("UNSTAKE"))
	SLASH_SIG      = crypto.Keccak256Hash([]byte("SLASH"))
	DELEGATE_SIG   = crypto.Keccak256Hash([]byte("DELEGATE"))
	UNDELEGATE_SIG = crypto.Keccak256Hash([]byte("UNDELEGATE"))
)

var (
	STAKE_PARAMS_TYPE             = abi.MustNewType("tuple(address validatorAddr, address ownerAddr, uint256 amount, uint256 commissionRate, uint256[2] bksKey, bytes pubKey)")
	ADDSTAKE_PARAMS_TYPE          = abi.MustNewType("tuple(address validatorAddr, uint256 amount)")
	UNSTAKE_PARAMS_TYPE           = abi.MustNewType("tuple(bytes32 sig, address validatorAddr, uint256 amount)")
	ROOT_CHAIN_SLASH_PARAMS_TYPE  = abi.MustNewType("tuple(address[] validatorAddrs, uint256 slashingPercentage, uint256 slashIncentivePercentage)")
	CHILD_CHAIN_SLASH_PARAMS_TYPE = abi.MustNewType("tuple(uint256 handleEventId, address[] validatorAddrs)")
	DELEGATE_PARAMS_TYPE          = abi.MustNewType("tuple(address validatorAddr, address delegterAddr, uint256 amount)")
	UNDELEGATE_PARAMS_TYPE        = abi.MustNewType("tuple(bytes32 sig, address validatorAddr, address delegterAddr, uint256 amount)")
)

func (c *StakeHandler) Initialize() error {

	return nil
}

func (c *StakeHandler) onStake(input []byte) error {
	decoded, err := abi.Decode(STAKE_PARAMS_TYPE, input)
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

	ownerAddr, ok := res["ownerAddr"].(ethgo.Address)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_owner")
	}

	amount, ok := res["amount"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNT")
	}

	commissionRate, ok := res["commissionRate"].(*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_COMMISSION_RATE")
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
	return c.stake(common.Address(validatorAddr), common.Address(ownerAddr), amount, commissionRate.Uint64(), &blsKey, publicKey)
}

func (c *StakeHandler) onAddStake(input []byte) error {
	decoded, err := abi.Decode(ADDSTAKE_PARAMS_TYPE, input)
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
	decoded, err := abi.Decode(CHILD_CHAIN_SLASH_PARAMS_TYPE, input)
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
	decoded, err := abi.Decode(DELEGATE_PARAMS_TYPE, input)
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

func (c *StakeHandler) stake(validatorAddr, owner common.Address, amount *big.Int, commissionRate uint64, blsKey *bls.PublicKey, pubKey *ecdsa.PublicKey) error {
	// ## NOTE ##
	//
	// Because there is a validator's stake information in the rootchain,
	// it is impossible to initiate duplicate `stake` operations from the rootchain
	// as long as the `unstake` has not been fully performed.
	if c.hasValidator(validatorAddr) {
		return typesdk.NewRevertError("StakeHandler: VALIDATOR ALREADY STAKE")
	}

	stakeIndex := c.incrementValidatorNonce()

	if err := c.setValidatorByPriority(validatorAddr, types.NewValidator(owner, amount, common.Big0, blsKey, pubKey, commissionRate, c.GetCurrentEpoch(), stakeIndex)); nil != err {
		log.Error("Failed to set validator stake", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: STAKE FAILED")
	}
	log.Info("Stake for", "validator", validatorAddr.Hex(), "owner", owner.Hex(), "amount", amount, "blsKey", string(blsKey.Bytes()),
		"pubKey", hex.EncodeToString(crypto.FromECDSAPub(pubKey)), "epoch", c.GetCurrentEpoch(), "stakeIndex", stakeIndex, "blockNumber", c.evm.Context.BlockNumber.Uint64())
	return nil
}

func (c *StakeHandler) addStake(validatorAddr common.Address, amount *big.Int) error {
	validator := c.GetValidator(validatorAddr)

	// ## NOTE ##
	// Reason:
	//		The invalid status of the childchain is synchronized to the rootchain, resulting in the ability to initiate an `addstake` on the rootchain.
	//		or due to the time difference between the `unstake` on the childchain and the `addstake` transaction on the rootchain.
	//
	// If the validator has already invalid (nonexistent) on the childchain,
	// but the `addstake` sent by the rootchain should be  appended as a stackewithdrawl item.
	if validator.IsEmpty() || validator.IsInvalid() {

		var err error
		lastEpoch := c.getStakeWithdrawalLastEpoch(validatorAddr)
		if lastEpoch != 0 {
			err = c.registerStakeWithdrawalByEpoch(validatorAddr, amount, lastEpoch) // append amount to last one
		} else {
			err = c.registerStakeWithdrawal(validatorAddr, amount, false) // append amount with currentEpoch
		}

		if nil != err {
			return typesdk.NewRevertError("StakeHandler: ADDSTAKE FAILED")
		}
	}

	// update validator priority
	validator.AddStakeAmount(amount)

	if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
		log.Error("Failed to add validator stake amount", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: ADD STAKE FAILED")
	}
	log.Info("AddStake for", "validator", validatorAddr.Hex(), "amount", amount, "epoch", c.GetCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) unStake(validatorAddr common.Address, amount *big.Int) error {
	validator := c.GetValidator(validatorAddr)

	if validator.IsEmpty() || validator.IsInvalid() {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	if validator.Owner != c.contract.Caller() {
		return typesdk.NewRevertError("StakeHandler: INVALID_SENDER")
	}

	if amount.Cmp(validator.StakeAmount) > 0 {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNT")
	}

	// update validator priority
	validator.SubStakeAmount(amount)

	var err error
	if validator.StakeAmount.Cmp(common.Big0) == 0 {
		validator.AppendStatus(types.Invalided | types.Unstaked)
		err = c.updateValidatorRemovePriority(validatorAddr, validator)
	} else {
		err = c.updateValidatorByPriority(validatorAddr, validator)
	}

	log.Info("UnStake for", "validator", validatorAddr.Hex(), "amount", amount, "blockNumber", c.evm.Context.BlockNumber)

	return err
}

func (c *StakeHandler) slash(handleEventId *big.Int, validatorAddrs []common.Address) error {
	if c.hasSlashProcessed(handleEventId) {
		return typesdk.NewRevertError("StakeHandler: SLASH_ALREADY_PROCESSED")
	}

	// TODO

	return nil
}

// TODO  如果存在 失效的 validator 或者 不存在的 validator 应该将 delegate Amount 追加到 delegateWithdrawal 中
func (c *StakeHandler) delegate(validatorAddr, delegaterAddr common.Address, amount *big.Int) error {

	validator := c.GetValidator(validatorAddr)

	if validator.IsEmpty() || validator.IsInvalid() {
		return typesdk.NewRevertError("StakeHandler: INVALID_VALIDATOR")
	}

	// update validator priority
	validator.AddDelegateAmount(amount)

	if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
		log.Error("Failed to add validator delegate amount", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: ADD DELEGATE AMOUNT OF VALIDATOR FAILED")
	}
	// update delegation
	if err := c.updateDelegation(delegaterAddr, validatorAddr, validator.Epoch, types.NewDelegation(c.GetCurrentEpoch(), amount)); nil != err {
		log.Error("Failed to set delegation", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: SET DELEGATION FAILED")
	}
	log.Info("Delegate for", "delegaterAddr", delegaterAddr.Hex(), "validator", validatorAddr.Hex(), "amount", amount, "epoch", c.GetCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) registerStakeWithdrawalByEpoch(validatorAddr common.Address, amount *big.Int, epoch uint64) error {

	item := c.getStakeWithdrawalQueueItem(validatorAddr, epoch)
	if item.IsNotEmpty() {
		item.IncrementAmount(amount)
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, epoch, item); nil != err {
			log.Error("Failed to register stake withdraw", "validatorAddr", validatorAddr.Hex(),
				"currentEpoch", c.GetCurrentEpoch(), "releaseEpoch", epoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
			return typesdk.NewRevertError("StakeHandler: SET REGISTER STAKE WITHDRAW FAILED")
		}

		if err := c.addLogStakeWithdrawalRegisteredEvent(validatorAddr, amount); nil != err {
			return err
		}
		log.Info("Register stake withdrawal for", "validator", validatorAddr.Hex(), "currentEpoch", c.GetCurrentEpoch(), "releaseEpoch", epoch, "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	}

	return nil
}

func (c *StakeHandler) registerStakeWithdrawal(validatorAddr common.Address, amount *big.Int, wait bool) error {
	currentEpoch := c.GetCurrentEpoch()
	var releaseEpoch uint64
	if wait {
		releaseEpoch = currentEpoch + DELEGATE_WITHDRAWAL_WAIT_PERIOD
	} else {
		releaseEpoch = currentEpoch
	}
	if err := c.appendStakeWithdrawal(validatorAddr, releaseEpoch, amount); nil != err {
		log.Error("Failed to register stake withdraw", "validatorAddr", validatorAddr.Hex(),
			"currentEpoch", currentEpoch, "releaseEpoch", releaseEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: SET REGISTER STAKE WITHDRAW FAILED")
	}

	if err := c.addLogStakeWithdrawalRegisteredEvent(validatorAddr, amount); nil != err {
		return err
	}
	log.Info("Register stake withdrawal for", "validator", validatorAddr.Hex(), "currentEpoch", c.GetCurrentEpoch(), "releaseEpoch", releaseEpoch, "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) registerDelegateWithdrawal(delegater, validatorAddr common.Address, amount *big.Int, wait bool) error {
	currentEpoch := c.GetCurrentEpoch()
	var releaseEpoch uint64
	if wait {
		releaseEpoch = currentEpoch + DELEGATE_WITHDRAWAL_WAIT_PERIOD
	} else {
		releaseEpoch = currentEpoch
	}
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

func (c *StakeHandler) syncStateUnStake(validatorAddr common.Address, amount *big.Int) error {

	data, err := abi.Encode([]interface{}{UNSTAKE_SIG, validatorAddr, amount}, UNSTAKE_PARAMS_TYPE)
	if nil != err {
		return typesdk.NewRevertError("encode L2StateSender unstake data failed")
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, address.StakeSenderAddress)
	if nil != err {
		return typesdk.NewRevertError("call unstake by L2StateSender failed")
	}

	if err := l2statesender.SyncState(address.RootchainStakeManagerAddress, data); nil != err {
		return typesdk.NewRevertError("call unstake by L2StateSender failed")
	}
	return nil
}

func (c *StakeHandler) syncStateUnDelegate(validatorAddr, delegaterAddr common.Address, amount *big.Int) error {
	data, err := abi.Encode([]interface{}{UNDELEGATE_SIG, validatorAddr, delegaterAddr, amount}, UNDELEGATE_PARAMS_TYPE)
	if nil != err {
		return typesdk.NewRevertError("encode L2StateSender undelegate data failed")
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, address.StakeSenderAddress)
	if nil != err {
		return typesdk.NewRevertError("call undelegate by L2StateSender failed")
	}

	if err := l2statesender.SyncState(address.RootchainStakeManagerAddress, data); nil != err {
		return typesdk.NewRevertError("call undelegate by L2StateSender failed")
	}
	return nil
}
