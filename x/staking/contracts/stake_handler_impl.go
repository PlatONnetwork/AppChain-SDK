package contracts

import (
	"bytes"
	"errors"

	"github.com/PlatONnetwork/AppChain-SDK/core/contracts"
	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	staketypes "github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/votetoken/contracts/erc20vote"
	platon "github.com/PlatONnetwork/PlatON-Go"
	"github.com/PlatONnetwork/PlatON-Go/common/math"

	"math/big"
	"strings"

	"github.com/PlatONnetwork/PlatON-Go/accounts/abi"
	"github.com/PlatONnetwork/PlatON-Go/accounts/abi/bind"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/PlatONnetwork/PlatON-Go/event"
	"github.com/PlatONnetwork/PlatON-Go/log"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = typesdk.RevertError{}
	_ = vm.EVM{}
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = platon.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

type StakeHandler struct {
	abi           *abi.ABI
	abis          map[uint64]*abi.ABI
	methodEntry   map[string]func([]byte) ([]byte, error)
	methodEntries map[uint64]map[string]func([]byte) ([]byte, error)
	readOnly      bool
	contract      *vm.Contract
	evm           *vm.EVM
	burner        contracts.Burn
	stateDb       *contracts.StateDB
	context       *contracts.Context
	fallback      func(input []byte) ([]byte, error)
	l1Module      staketypes.L1Moduler
	stageModule   staketypes.StageModuler
	stakeModule   staketypes.StakeModuler
	rewardModule  staketypes.RewardModuler
}

func NewStakeHandler(evm *vm.EVM, contract *vm.Contract, readOnly bool) (*StakeHandler, error) {
	s := &StakeHandler{
		abi:           nil,
		abis:          make(map[uint64]*abi.ABI),
		methodEntry:   make(map[string]func([]byte) ([]byte, error)),
		methodEntries: make(map[uint64]map[string]func([]byte) ([]byte, error)),
		evm:           evm,
		contract:      contract,
		burner:        contracts.NewBurner(contract),
		stateDb:       contracts.NewStateDB(evm, contract),
		context:       contracts.NewContext(evm, contract),
		readOnly:      readOnly,
	}
	s.initABI()
	s.initMethodEntry()
	s.loadMethodABI()
	return s, nil
}

// external

// @notice Query the list of information on the number of sealed blocks of validators for a certain period
// @dev For the convenience of expanding the list of validators with multiple period properties
// @param periodType represents a period of a certain type
// @param period represents the number of intervals
// @return BlocksOfValidator array for query
//
// ## NOTE ##
// periodType options:
// 0: unknown
// 1: round
// 2: epoch
// ...
func (c *StakeHandler) GetBlocksOfValidators(periodType uint8, period *big.Int) ([]BlocksOfValidator, error) {
	switch periodType {
	case 1: // round
		return c.getBlocksOfValidatorsForRound(period.Uint64())
	case 2: // epoch
		return c.getBlocksOfValidatorsForEpoch(period.Uint64())
	default:
		return nil, typesdk.NewRevertError("StakeHandler: UNKNOWN PERIOD TYPE")
	}
}

func (c *StakeHandler) GetDelegationsWithValidator(validators []common.Address, delegator common.Address) ([]DelegationInfo, error) {
	delegationQueue := make([]DelegationInfo, 0)
	for _, validatorAddr := range validators {
		epochs, _ := c.getValidatorDelegationRcPendingAndEpoch(validatorAddr, math.MaxUint64)
		for _, stakeEpoch := range epochs {
			delegation := c.getDelegation(delegator, validatorAddr, stakeEpoch)
			if delegation.IsEmpty() {
				continue
			}
			delegationQueue = append(delegationQueue, DelegationInfo{
				ValidatorAddr: validatorAddr,
				DelegatorAddr: delegator,
				Amount:        delegation.Amount,
				StakeEpoch:    new(big.Int).SetUint64(stakeEpoch),
				DelegateEpoch: new(big.Int).SetUint64(delegation.Epoch),
			})
		}
	}
	return delegationQueue, nil
}

// @notice Query the list of validators for a certain period
// @dev For the convenience of expanding the list of validators with multiple period properties
// @param periodType represents a period of a certain type
// @param period represents the number of intervals
// @return validator address array
//
// ## NOTE ##
// periodType options:
// 0: unknown
// 1: round
// 2: epoch
// ...
func (c *StakeHandler) GetValidatorAddrs(periodType uint8, period *big.Int) ([]common.Address, error) {

	switch periodType {
	case 1: // round
		return c.getRoundValidatorIds(period.Uint64()), nil
	case 2: // epoch
		return c.getEpochValidatorIds(period.Uint64()), nil
	default:
		return nil, typesdk.NewRevertError("StakeHandler: UNKNOWN PERIOD TYPE")
	}
}

func (c *StakeHandler) GetValidators(start []byte, size *big.Int) ([]byte, []ValidatorInfo, error) {
	next, validatorAddrQueue, validatorQueue := c.getValidatorsByPriorityKey(start, size.Uint64())
	validatorInfoQueue := make([]ValidatorInfo, len(validatorQueue))
	for i, validator := range validatorQueue {
		validatorInfoQueue[i] = ValidatorInfo{
			ValidatorAddr:  validatorAddrQueue[i],
			Owner:          validator.Owner,
			StakeAmount:    validator.StakeAmount,
			DelegateAmount: validator.DelegateAmount,
			CommissionRate: new(big.Int).SetUint64(validator.CommissionRate),
			Status:         new(big.Int).SetUint64(uint64(validator.Status)),
			Epoch:          new(big.Int).SetUint64(validator.Epoch),
			StakeIndex:     new(big.Int).SetUint64(validator.StakeIndex),
			PubKey:         validator.PubKey.Bytes(),
			BlsKey:         validator.BlsKey,
		}
	}
	return next, validatorInfoQueue, nil
}

func (c *StakeHandler) GetValidatorsWithAddr(validatorAddrs []common.Address) ([]ValidatorInfo, error) {

	validatorInfoQueue := make([]ValidatorInfo, 0)
	for _, validatorAddr := range validatorAddrs {
		validator := c.getValidator(validatorAddr)
		if validator.IsEmpty() {
			continue
		}
		validatorInfoQueue = append(validatorInfoQueue, ValidatorInfo{
			ValidatorAddr:  validatorAddr,
			Owner:          validator.Owner,
			StakeAmount:    validator.StakeAmount,
			DelegateAmount: validator.DelegateAmount,
			CommissionRate: new(big.Int).SetUint64(validator.CommissionRate),
			Status:         new(big.Int).SetUint64(uint64(validator.Status)),
			Epoch:          new(big.Int).SetUint64(validator.Epoch),
			StakeIndex:     new(big.Int).SetUint64(validator.StakeIndex),
			PubKey:         validator.PubKey.Bytes(),
			BlsKey:         validator.BlsKey,
		})
	}
	return validatorInfoQueue, nil
}

func (c *StakeHandler) PendingWithdrawalsOfDelegate(validator common.Address, delegator common.Address) (*big.Int, error) {
	return c.getDelegateWithdrawalPending(delegator, validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) PendingWithdrawalsOfStake(validator common.Address) (*big.Int, error) {
	return c.getStakeWithdrawalPending(validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) VerifyAggregateSignature(blockNumber *big.Int, validatorIndexs []*big.Int, data common.Hash, signatues []byte) (bool, error) {
	return c.verifyBLSAggregateSignature(blockNumber, validatorIndexs, data, signatues)
}

func (c *StakeHandler) VerifyAggregateSignatureByValidators(validators []common.Address, data common.Hash, signatues []byte) (bool, error) {
	return c.verifyBLSAggregateSignatureByValidators(validators, data, signatues)
}

func (c *StakeHandler) WithdrawableOfDelegate(validator common.Address, delegator common.Address) (*big.Int, error) {
	return c.getDelegateWithdrawable(delegator, validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) WithdrawableOfStake(validator common.Address) (*big.Int, error) {
	return c.getStakeWithdrawable(validator, c.getCurrentEpoch()), nil
}

func (c *StakeHandler) OnStateReceive(id *big.Int, sender common.Address, data []byte) error {

	rootchainStakeManagerAddress, err := c.l1Module.GetStakeManagerAddress()
	if nil != err {
		return typesdk.NewRevertError("StakeHandler: NOT FOUND STAKE MANAGER ADDR")
	}

	if c.contract.Caller() != constants.StateSyncAddress || sender != rootchainStakeManagerAddress {
		log.Error("Invalid sender", "caller", c.contract.Caller().Hex(), "address", constants.StateSyncAddress.Hex(),
			"sender", sender.Hex(), "rootchain", rootchainStakeManagerAddress.Hex())
		return typesdk.NewRevertError("StakeHandler: INVALID_SENDER")
	}
	if bytes.Compare(data[:METHODID_SIZE], STAKE_SIG.Bytes()) == 0 {
		return c.onStake(data)
	} else if bytes.Compare(data[:METHODID_SIZE], ADDSTAKE_SIG.Bytes()) == 0 {
		return c.onAddStake(data)
	} else if bytes.Compare(data[:METHODID_SIZE], SLASH_SIG.Bytes()) == 0 {
		return c.onSlash(data)
	} else if bytes.Compare(data[:METHODID_SIZE], DELEGATE_SIG.Bytes()) == 0 {
		return c.onDelegate(data)
	} else {
		return typesdk.NewRevertError("StakeHandler: INVALID_METHOD_SIGN")
	}
}

func (c *StakeHandler) Slash() error {

	currentRound := c.getCurrentRound()
	minBlocksOfRoundValidator := c.stakeModule.GetMinBlocksOfRoundValidator(c.evm.StateDB)
	validatorAddrs := db.CheckLowBlocksValidatorForPreviousRound(c.evm.StateDB, c.contract.Address(), currentRound, minBlocksOfRoundValidator)

	slashingValidatorAddrCache := make(map[common.Address]struct{}, 0)
	// NOTE: update validator status (add log for lowBlocks slashing)
	for _, validatorAddr := range validatorAddrs {
		validator := c.getValidator(validatorAddr)
		if validator.IsEmpty() {
			log.Warn("Not found validator when Slash", "validatorAddr", validatorAddr.Hex(), "currentRound", currentRound, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
			continue
		}
		if validator.IsInvalidSlashing() {
			log.Warn("Was slashed validator when Slash", "validatorAddr", validatorAddr.Hex(), "currentRound", currentRound, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
			continue
		}
		// 1. add validator status (add: invalida|slashing)
		validator.AppendStatus(staketypes.Invalided | staketypes.Slashing)
		// 2. update validator status (add: invalida|slashing) AND remove validator priority
		if err := c.updateValidatorRemovePriority(validatorAddr, validator); nil != err {
			log.Error("Failed to call updateValidatorRemovePriority", "validatorAddr", validatorAddr.Hex(), "currentRound", currentRound, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber, "error", err)
			return typesdk.NewRevertError("StakeHandler: can not update validator priority")
		}
		// 3. add updateValidatorStatus log
		if err := c.addLogUpdateValidatorStatusEvent(validatorAddr, new(big.Int).SetUint64(uint64(validator.Status))); nil != err {
			return err
		}
		slashingValidatorAddrCache[validatorAddr] = struct{}{}
		c.burnVoteToken(validator.Owner, validator.StakeAmount)
	}
	// ###### NOTE: ######
	// remove validator from epoch validators
	if err := c.removeValidatorsFromEpochValidatorQueue(slashingValidatorAddrCache); nil != err {
		return err
	}

	// sync state to L1
	if err := c.syncStateSlash(validatorAddrs); nil != err {
		return err
	}

	log.Info("Begin Slash for", "validator size", len(validatorAddrs), "minBlocksOfRoundValidator", minBlocksOfRoundValidator,
		"currentRound", currentRound, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) Undelegate(validatorAddr common.Address, amount *big.Int) error {

	delegatorAddr := c.contract.Caller()

	currentEpoch := c.getCurrentEpoch()
	epochs, _ := c.getValidatorDelegationRcPendingAndEpoch(validatorAddr, math.MaxUint64)
	origin := new(big.Int).SetBytes(amount.Bytes())
	paid := big.NewInt(0)

	validator := c.getValidator(validatorAddr)
	for _, stakeEpoch := range epochs {

		if amount.Cmp(common.Big0) == 0 {
			break
		}

		delegation := c.getDelegation(delegatorAddr, validatorAddr, stakeEpoch)
		if delegation.IsEmpty() {
			continue
		}

		// NOTE:
		// Priority must be given to settling commission rewards before proceeding with the `withdraw` operation.
		if err := c.rewardModule.UpdateDelegationRewardsByStakeEpoch(c.evm.StateDB, delegatorAddr, validatorAddr, stakeEpoch); nil != err {
			log.Error("Failed to update delegation rewards by stakeEpoch on StakeHandler.Undelegate()", "delegatorAddr", delegatorAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
				"stakeEpoch", stakeEpoch, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "error", err)
			return typesdk.NewRevertError("StakeHandler: UPDATE DELEGATION REWARDS BY STAKE EPOCH FAILED")
		}

		use := big.NewInt(0)
		if delegation.Amount.Cmp(amount) <= 0 {
			// remove the delegation by stakeEpoch when:
			// 1. had not withdrawable undelegate amount
			c.removeDelegation(delegatorAddr, validatorAddr, stakeEpoch)
			log.Debug("remove delegation on StakeHandler.Undelegate()", "delegatorAddr", delegatorAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
				"stakeEpoch", stakeEpoch, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)

			// decrement validator-delegator-rc
			if err := c.releaseValidatorDelegationRcItem(validatorAddr, stakeEpoch, 1); nil != err {
				log.Error("Failed to release validatorDelegation rc on StakeHandler.Undelegate()", "delegatorAddr", delegatorAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
					"stakeEpoch", stakeEpoch, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "error", err)
				return typesdk.NewRevertError("StakeHandler: RELEASE VALIDATOR DELEGATION RC FAILED")
			}
			use = delegation.Amount
		} else {
			// update delegation with new epoch and new amount
			if delegation.Epoch < currentEpoch {
				delegation.UpdateEpoch(currentEpoch)
				delegation.SnapPreEpochAmount()
			}
			delegation.DecrementAmount(amount)
			if err := c.setDelegation(delegatorAddr, validatorAddr, stakeEpoch, delegation); nil != err {
				log.Error("Failed to update delegation on StakeHandler.Undelegate()", "delegatorAddr", delegatorAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
					"stakeEpoch", stakeEpoch, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "error", err)
				return typesdk.NewRevertError("StakeHandler: SET DELEGATION FAILED")
			}
			log.Debug("update delegation on StakeHandler.Undelegate()", "delegatorAddr", delegatorAddr.Hex(), "validatorAddr", validatorAddr.Hex(),
				"stakeEpoch", stakeEpoch, "currentEpoch", currentEpoch, "delegationEpoch", delegation.Epoch, "delegationAmount", delegation.Amount,
				"blockNumber", c.evm.Context.BlockNumber)

			use = amount
		}

		amount = new(big.Int).Sub(amount, use)
		paid = new(big.Int).Add(paid, use)

		// update validator priority
		if validator.IsValid() && validator.Epoch == stakeEpoch {

			validator.SubDelegateAmount(use)

			log.Debug("update validator on StakeHandler.Undelegate()", "validatorAddr", validatorAddr.Hex(), "stakeEpoch", stakeEpoch, "currentEpoch", currentEpoch,
				"validator.DelegateAmount", validator.DelegateAmount, "blockNumber", c.evm.Context.BlockNumber)

			if err := c.updateValidatorByPriority(validatorAddr, validator); nil != err {
				log.Error("Failed to update validator priority on StakeHandler.Undelegate()", "validatorAddr", validatorAddr.Hex(), "error", err)
				return typesdk.NewRevertError("StakeHandler: SUB DELEGATE AMOUNT OF VALIDATOR FAILED")
			}
		}
	}
	c.burnVoteToken(delegatorAddr, origin)
	if err := c.registerDelegateWithdrawal(delegatorAddr, validatorAddr, paid, true); nil != err {
		return err
	}
	if err := c.addLogUnDelegatedEvent(delegatorAddr, validatorAddr, paid); nil != err {
		return err
	}
	log.Info("Undelegate for", "delegator", delegatorAddr.Hex(), "validator", validatorAddr.Hex(), "expect amount", origin, "use amount", paid, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) Unstake(validatorAddr common.Address, amount *big.Int) error {

	if err := c.unStake(validatorAddr, amount); nil != err {
		return err
	}

	if err := c.registerStakeWithdrawal(validatorAddr, amount, true); nil != err {
		return err
	}

	if err := c.addLogUnStakedEvent(validatorAddr, amount); nil != err {
		return err
	}
	log.Info("Unstake for", "validator", validatorAddr.Hex(), "amount", amount, "currentEpoch", c.getCurrentEpoch(), "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) WithdrawUndelegate(validator common.Address) error {

	currentEpoch := c.getCurrentEpoch()
	delegator := c.contract.Caller()
	amount, err := c.applyDelegateWithdrawable(delegator, validator, currentEpoch)
	if nil != err {
		log.Error("Failed to withdraw undelegate", "validatorAddr", validator.Hex(),
			"currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: CAN NOT UPDATE DELEGATE WITHDRAW PENDDING HEAD")
	}
	if amount.Cmp(common.Big0) == 0 {
		log.Error("has no withdrawable delegate amount", "delegator", delegator, "validator", validator.Hex(),
			"amount", amount, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
		return typesdk.NewRevertError("StakeHandler: HAS NO WITHDRAWABLE DELEGATE AMOUNT")
	}

	if err := c.addLogDelegateWithdrawalEvent(delegator, validator, amount); nil != err {
		return err
	}

	if err := c.syncStateUnDelegate(validator, delegator, amount); nil != err {
		return err
	}

	log.Info("Withdraw undelegate for", "delegator", delegator, "validator", validator.Hex(), "amount", amount, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) WithdrawUnstake(validatorAddr common.Address) error {

	validator := c.getValidator(validatorAddr)

	// #### NOTE ####
	// When the validator is in the period of slashing,
	// the validator does not accept any action until the slashing process is completed
	if validator.IsInvalidSlashing() {
		return typesdk.NewRevertError("StakeHandler: SLASHING_VALIDATOR")
	}

	currentEpoch := c.getCurrentEpoch()
	amount, err := c.applyStakeWithdrawable(validatorAddr, currentEpoch)
	if nil != err {
		log.Error("Failed to withdraw unstake", "validatorAddr", validatorAddr.Hex(),
			"currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber, "amount", amount, "error", err)
		return typesdk.NewRevertError("StakeHandler: CAN NOT UPDATE STAKE WITHDRAW PENDING HEAD")
	}

	if amount.Cmp(common.Big0) == 0 {
		log.Error("has no withdrawable stake amount", "validatorAddr", validatorAddr.Hex(),
			"amount", amount, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
		return typesdk.NewRevertError("StakeHandler: HAS NO WITHDRAWABLE STAKE AMOUNT")
	}

	// remove validator when:
	// 1. validator status has 'unstake'
	// 2. had not withdrawable unstake amount
	if validator.IsInvalidUnstaked() && (common.Big0.Cmp(c.getStakeWithdrawalPending(validatorAddr, currentEpoch)) == 0) {
		c.removeValidator(validatorAddr)
	}

	if err := c.addLogStakeWithdrawalEvent(validatorAddr, amount); nil != err {
		return err
	}

	if err := c.syncStateUnStake(validatorAddr, amount); nil != err {
		return err
	}

	log.Info("Withdraw unstake for", "validatorAddr", validatorAddr.Hex(), "amount", amount, "currentEpoch", currentEpoch, "blockNumber", c.evm.Context.BlockNumber)
	return nil
}

func (c *StakeHandler) mintVoteToken(account common.Address, amount *big.Int) {
	log.Debug("Mint vote token", "account", account.Hex(), "amount", amount)
	caller, err := erc20vote.NewERC20VoteCaller(c.evm, c.contract, constants.VoteTokenAddress)
	contracts.Require(err == nil, "StakeHandler: CREATE VOTE CALLER FAILED")
	contracts.Require(caller.Mint(account, amount) == nil, "StakeHandler: MINT VOTE TOKEN FAILED")
}

func (c *StakeHandler) burnVoteToken(account common.Address, amount *big.Int) {
	log.Debug("Burn vote token", "account", account.Hex(), "amount", amount)
	caller, err := erc20vote.NewERC20VoteCaller(c.evm, c.contract, constants.VoteTokenAddress)
	contracts.Require(err == nil, "StakeHandler: CREATE VOTE CALLER FAILED")
	contracts.Require(caller.Burn(account, amount) == nil, "StakeHandler: BURN VOTE TOKEN FAILED")
}
