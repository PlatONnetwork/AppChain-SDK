package contracts

import (
	"crypto/ecdsa"
	"encoding/hex"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	statesenderC "github.com/PlatONnetwork/AppChain-SDK/x/statesender/contracts"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
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
	ROOT_CHAIN_SLASH_PARAMS_TYPE  = abi.MustNewType("tuple(bytes32 sig, address[] validatorAddrs, uint256 slashingPercentage, uint256 slashIncentivePercentage)")
	CHILD_CHAIN_SLASH_PARAMS_TYPE = abi.MustNewType("tuple(bytes32 sig, uint256 handleEventId, address[] validatorAddrs, uint256[] amounts)")
	DELEGATE_PARAMS_TYPE          = abi.MustNewType("tuple(address validatorAddr, address delegterAddr, uint256 amount)")
	UNDELEGATE_PARAMS_TYPE        = abi.MustNewType("tuple(bytes32 sig, address validatorAddr, address delegterAddr, uint256 amount)")
)

func (c *StakeHandler) Initialize() error {

	return nil
}

func (c *StakeHandler) onStake(input []byte) error {
	decoded, err := abi.Decode(STAKE_PARAMS_TYPE, input)
	if nil != err {
		log.Error("Failed to decode stake data", "error", err)
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
	return c.stake(basecommon.Address(validatorAddr), basecommon.Address(ownerAddr), amount, commissionRate.Uint64(), &blsKey, publicKey)
}

func (c *StakeHandler) onAddStake(input []byte) error {
	decoded, err := abi.Decode(ADDSTAKE_PARAMS_TYPE, input)
	if nil != err {
		log.Error("Failed to decode addStake data", "error", err)
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

	return c.addStake(basecommon.Address(validatorAddr), amount)
}

func (c *StakeHandler) onSlash(input []byte) error {
	decoded, err := abi.Decode(CHILD_CHAIN_SLASH_PARAMS_TYPE, input)
	if nil != err {
		log.Error("Failed to decode slash data", "error", err)
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

	addrs := make([]basecommon.Address, len(validatorAddrs))
	for i, v := range validatorAddrs {
		addrs[i] = basecommon.Address(v)
	}

	amounts, ok := res["amounts"].([]*big.Int)
	if !ok {
		return typesdk.NewRevertError("StakeHandler: INVALID_AMOUNTS")
	}

	return c.slash(handleEventId, addrs, amounts)
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

	return c.delegate(basecommon.Address(validatorAddr), basecommon.Address(delegterAddr), amount)
}

func (c *StakeHandler) stake(validatorAddr, owner basecommon.Address, amount *big.Int, commissionRate uint64, blsKey *bls.PublicKey, pubKey *ecdsa.PublicKey) error {
	// ## NOTE ##
	//
	// Because there is a validator's stake information in the rootchain,
	// it is impossible to initiate duplicate `stake` operations from the rootchain
	// as long as the `unstake` has not been fully performed.
	if c.hasValidator(validatorAddr) {
		return typesdk.NewRevertError("StakeHandler: VALIDATOR ALREADY STAKE")
	}

	stakeIndex := c.incrementValidatorNonce()

	if err := c.setValidatorByPriority(validatorAddr, types.NewValidator(owner, amount, basecommon.Big0, blsKey, pubKey, commissionRate, c.getCurrentEpoch(), stakeIndex)); nil != err {
		log.Error("Failed to set validator stake", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: STAKE FAILED")
	}

	if err := c.addLogStakeAddedEvent(validatorAddr, amount); nil != err {
		return err
	}

	log.Info("Stake for", "validator", validatorAddr.Hex(), "owner", owner.Hex(), "amount", amount, "blsKey", string(blsKey.Bytes()),
		"pubKey", hex.EncodeToString(crypto.FromECDSAPub(pubKey)), "epoch", c.getCurrentEpoch(), "stakeIndex", stakeIndex, "blockNumber", c.evm.Context.BlockNumber.Uint64())
	return nil
}

func (c *StakeHandler) addStake(validatorAddr basecommon.Address, amount *big.Int) error {
	validator := c.getValidator(validatorAddr)

	// ## NOTE ##
	// Reason:
	//		The invalid status of the childchain is synchronized to the rootchain, resulting in the ability to initiate an `addstake` on the rootchain.
	//		or due to the time difference between the `unstake` on the childchain and the `addstake` transaction on the rootchain.
	//
	// If the validator has already invalid (nonexistent) on the childchain,
	// but the `addstake` sent by the rootchain should be  appended as a stackewithdrawl item.
	if validator.IsInvalid() {

		var err error
		lastEpoch := c.getStakeWithdrawalLastEpoch(validatorAddr)
		if lastEpoch != 0 {
			err = c.registerStakeWithdrawalByEpoch(validatorAddr, amount, lastEpoch) // append amount to last one
		} else {
			err = c.registerStakeWithdrawal(validatorAddr, amount, false) // append amount with currentEpoch
		}

		if nil != err {
			log.Error("Failed to call registerStakeWithdrawal", "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
			return typesdk.NewRevertError("StakeHandler: ADD STAKE FAILED")

		}
	}

	// update validator priority
	validator.AddStakeAmount(amount)

	if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
		log.Error("Failed to add validator stake amount", "validatorAddr", validatorAddr.Hex(), "error", err)
		return typesdk.NewRevertError("StakeHandler: ADD STAKE FAILED")
	}

	if err := c.addLogStakeAddedEvent(validatorAddr, amount); nil != err {
		return err
	}

	log.Info("AddStake for", "validator", validatorAddr.Hex(), "amount", amount, "epoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) unStake(validatorAddr basecommon.Address, amount *big.Int) error {
	validator := c.getValidator(validatorAddr)

	if validator.IsInvalid() {
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

	if validator.StakeAmount.Cmp(basecommon.Big0) == 0 {
		validator.AppendStatus(types.Invalided | types.Unstaked)
		if err := c.updateValidatorRemovePriority(validatorAddr, validator); nil != err {
			log.Error("Failed to call updateValidatorRemovePriority", "validatorAddr", validatorAddr.Hex(), "error", err)
			return typesdk.NewRevertError("StakeHandler: can not update validator priority")
		}
	} else {
		if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
			log.Error("Failed to call updateValidatorByPriority", "validatorAddr", validatorAddr.Hex(), "error", err)
			return typesdk.NewRevertError("StakeHandler: can not update validator priority")
		}
	}

	return nil
}

func (c *StakeHandler) slash(handleEventId *big.Int, validatorAddrs []basecommon.Address, amounts []*big.Int) error {
	if c.hasSlashProcessed(handleEventId) {
		return typesdk.NewRevertError("StakeHandler: SLASH_ALREADY_PROCESSED")
	}
	if len(validatorAddrs) != len(amounts) {
		return typesdk.NewRevertError("StakeHandler: INVALID_PARAMS")
	}

	queue := types.NewSlashValidatorWithdrawItemQueue(uint64(len(validatorAddrs)))
	cache := make(map[basecommon.Address]struct{}, 0)
	for i, validatorAddr := range validatorAddrs {

		queue[i] = types.NewSlashValidatorWithdrawItem(validatorAddr, amounts[i])
		cache[validatorAddr] = struct{}{}
		// NOTE: unstake short circuit
		c.removeValidator(validatorAddr)
		c.cleanStakeWithdrawable(validatorAddr)

	}

	if err := c.setSlashProcessed(handleEventId, queue); nil != err {
		log.Error("Failed to call setSlashProcessed", "handleEventId", handleEventId, "error", err)
		return typesdk.NewRevertError("StakeHandler: INVALID_PARAMS")
	}

	// ###### NOTE: ######
	// To prevent transaction time differences between L1 and L2,
	// after processing L1's crash, return to L2 to remove the validator
	// and try again to remove the validator from Epoch validators.
	//(as validators may be selected again during time differences)
	currentEpoch := c.getCurrentEpoch()
	epochValidatorAddrQueue := db.GetEpochValidatorSharesSnapshotQueue(c.evm.StateDB, c.contract.Address(), currentEpoch)
	oldSize := len(epochValidatorAddrQueue)
	for i := 0; i < len(epochValidatorAddrQueue); i++ {
		validator := epochValidatorAddrQueue[i]
		if _, ok := cache[validator.ValidatorAddr]; !ok {
			// remove the validatorAddr from epoch validatorAddrQueue
			epochValidatorAddrQueue = append(epochValidatorAddrQueue[:i], epochValidatorAddrQueue[i+1:]...)
			i--
		}
	}
	// NTOE: update epoch validator snapshot queue (after remove low blocks validators)
	if len(epochValidatorAddrQueue) != oldSize {
		if err := db.SetEpochValidatorSharesSnapshotQueue(c.evm.StateDB, c.contract.Address(), currentEpoch, epochValidatorAddrQueue); nil != err {
			log.Error("Failed to update epochValidators", "epoch", currentEpoch, "error", err)
			return typesdk.NewRevertError("StakeHandler: UPDATE EPOCH VALIDATORS FAILED")
		}
	}

	if err := c.addLogSlashedEvent(handleEventId, validatorAddrs, amounts); nil != err {
		return err
	}

	log.Info("Slash for", "handleEventId", handleEventId, "validator size", len(validatorAddrs), "epoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) delegate(validatorAddr, delegaterAddr basecommon.Address, amount *big.Int) error {

	validator := c.getValidator(validatorAddr)

	var err error
	if validator.IsInvalid() {
		if err = c.registerDelegateWithdrawal(delegaterAddr, validatorAddr, amount, false); nil != err {
			return err
		}
	} else {

		currentEpoch := c.getCurrentEpoch()
		delegation := c.getDelegation(delegaterAddr, validatorAddr, validator.Epoch)

		if delegation.IsNotEmpty() {
			// NOTE:
			// Priority must be given to settling commission rewards before proceeding with the `withdraw` operation.
			if err := c.reward.UpdateDelegationRewardsByStakeEpoch(c.evm.StateDB, delegaterAddr, validatorAddr, validator.Epoch); nil != err {
				log.Error("Failed to update delegation rewards by stakeEpoch", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
					"stakeEpoch", validator.Epoch, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "error", err)
				return typesdk.NewRevertError("StakeHandler: UPDATE DELEGATION REWARDS BY STAKE EPOCH FAILED")
			}

			// update delegation
			delegation.UpdateEpoch(currentEpoch)
			delegation.IncrementAmount(amount)
		} else {
			delegation = types.NewDelegation(currentEpoch, amount)

			// increment validator-delegater-rc
			if err = c.appendValidatorDelegationRc(validatorAddr, validator.Epoch, 1); nil != err {
				log.Error("Failed to append validatorDelegation rc", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
					"stakeEpoch", validator.Epoch, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "error", err)
				return typesdk.NewRevertError("StakeHandler: APPEND VALIDATOR DELEGATION RC FAILED")
			}
		}

		if err := c.setDelegation(delegaterAddr, validatorAddr, validator.Epoch, delegation); nil != err {
			log.Error("Failed to set delegation", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
				"stakeEpoch", validator.Epoch, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "error", err)
			return typesdk.NewRevertError("StakeHandler: SET DELEGATION FAILED")
		}

		// update validator priority
		validator.AddDelegateAmount(amount)
		if err = c.updateValidatorByPriority(validatorAddr, validator); nil != err {
			log.Error("Failed to add validator delegate amount", "validatorAddr", validatorAddr.Hex(), "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "error", err)
			return typesdk.NewRevertError("StakeHandler: ADD DELEGATE AMOUNT OF VALIDATOR FAILED")
		}

		if err := c.addLogDelegationAddedEvent(delegaterAddr, validatorAddr, amount); nil != err {
			return err
		}

		log.Info("Delegate for", "delegaterAddr", delegaterAddr.Hex(), "validator", validatorAddr.Hex(), "amount", amount, "epoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	}
	return nil
}

func (c *StakeHandler) registerStakeWithdrawalByEpoch(validatorAddr basecommon.Address, amount *big.Int, epoch uint64) error {

	item := db.GetStakeWithdrawalQueueItem(c.evm.StateDB, c.contract.Address(), validatorAddr, epoch)
	if item.IsNotEmpty() {
		item.IncrementAmount(amount)
		if err := c.setStakeWithdrawalQueueItem(validatorAddr, epoch, item); nil != err {
			log.Error("Failed to register stake withdraw", "validatorAddr", validatorAddr.Hex(),
				"currentEpoch", c.getCurrentEpoch(), "releaseEpoch", epoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
			return typesdk.NewRevertError("StakeHandler: SET REGISTER STAKE WITHDRAW FAILED")
		}

	} else {
		if err := c.appendStakeWithdrawal(validatorAddr, epoch, amount); nil != err {
			log.Error("Failed to register stake withdraw", "validatorAddr", validatorAddr.Hex(),
				"currentEpoch", c.getCurrentEpoch(), "releaseEpoch", epoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
			return typesdk.NewRevertError("StakeHandler: SET REGISTER STAKE WITHDRAW FAILED")
		}
	}

	if err := c.addLogStakeWithdrawalRegisteredEvent(validatorAddr, amount); nil != err {
		return err
	}
	log.Info("Register stake withdrawal for", "validator", validatorAddr.Hex(), "currentEpoch", c.getCurrentEpoch(),
		"releaseEpoch", epoch, "amount", amount, "blockNumber", c.evm.Context.BlockNumber)

	return nil
}

func (c *StakeHandler) registerStakeWithdrawal(validatorAddr basecommon.Address, amount *big.Int, wait bool) error {
	currentEpoch := c.getCurrentEpoch()
	var releaseEpoch uint64
	if wait {
		releaseEpoch = currentEpoch + constants.STAKE_WITHDRAWAL_WAIT_PERIOD
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
	log.Info("Register stake withdrawal for", "validator", validatorAddr.Hex(), "currentEpoch", c.getCurrentEpoch(), "releaseEpoch", releaseEpoch, "amount", amount, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) registerDelegateWithdrawal(delegater, validatorAddr basecommon.Address, amount *big.Int, wait bool) error {
	currentEpoch := c.getCurrentEpoch()
	var releaseEpoch uint64
	if wait {
		releaseEpoch = currentEpoch + constants.DELEGATE_WITHDRAWAL_WAIT_PERIOD
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

func (c *StakeHandler) syncStateUnStake(validatorAddr basecommon.Address, amount *big.Int) error {

	data, err := abi.Encode([]interface{}{UNSTAKE_SIG, validatorAddr, amount}, UNSTAKE_PARAMS_TYPE)
	if nil != err {
		log.Error("Failed to encode unstake syncState data", "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("encode L2StateSender unstake data failed")
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, constants.StateSenderAddress)
	if nil != err {
		log.Error("Failed to call NewL2StateSenderCaller", "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("call unstake by L2StateSender failed")
	}

	if err := l2statesender.SyncState(constants.RootchainStakeManagerAddress, data); nil != err {
		log.Error("Failed to call SyncState", "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("call unstake by L2StateSender failed")
	}
	return nil
}

func (c *StakeHandler) syncStateUnDelegate(validatorAddr, delegaterAddr basecommon.Address, amount *big.Int) error {
	data, err := abi.Encode([]interface{}{UNDELEGATE_SIG, validatorAddr, delegaterAddr, amount}, UNDELEGATE_PARAMS_TYPE)
	if nil != err {
		log.Error("Failed to encode undelegate syncState data", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("encode L2StateSender undelegate data failed")
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, constants.StateSenderAddress)
	if nil != err {
		log.Error("Failed to call NewL2StateSenderCaller", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("call undelegate by L2StateSender failed")
	}

	if err := l2statesender.SyncState(constants.RootchainStakeManagerAddress, data); nil != err {
		log.Error("Failed to call SyncState", "delegaterAddr", delegaterAddr.Hex(), "validatorAddr", validatorAddr.Hex(), "amount", amount, "error", err)
		return typesdk.NewRevertError("call undelegate by L2StateSender failed")
	}
	return nil
}

func (c *StakeHandler) syncStateSlash(validators []basecommon.Address) error {

	data, err := abi.Encode([]interface{}{SLASH_SIG, validators, constants.SLASHING_PERCENTAGE, constants.SLASH_INCENTIVE_PERCENTAGE}, ROOT_CHAIN_SLASH_PARAMS_TYPE)
	if nil != err {
		log.Error("Failed to encode slash syncState data", "validators size", len(validators), "error", err)
		return typesdk.NewRevertError("encode L2StateSender slash data failed")
	}

	l2statesender, err := statesenderC.NewL2StateSenderCaller(c.evm, c.contract, constants.StateSenderAddress)
	if nil != err {
		log.Error("Failed to call NewL2StateSenderCaller", "validators size", len(validators), "error", err)
		return typesdk.NewRevertError("call slash by L2StateSender failed")
	}

	if err := l2statesender.SyncState(constants.RootchainStakeManagerAddress, data); nil != err {
		log.Error("Failed to call SyncState", "validators size", len(validators), "error", err)
		return typesdk.NewRevertError("call slash by L2StateSender failed")
	}
	return nil
}
