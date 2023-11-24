package types

import (
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/crypto/bls"
	"math/big"
	"strings"
)

const (
	/**
	######   ######   ######   ######
	#	  THE VALIDATOR  STATUS     #
	######   ######   ######   ######
	*/
	Invalided    ValidatorStatus = 1 << iota // 0001: The validator is deactivated
	LowBlocks                                // 0010: The validator was low block rate
	LowThreshold                             // 0100: The validator's stake was lower than minimum stake threshold
	Duplicated                               // 1000: The validator was duplicate block or duplicate signature
	Unstaked                                 // 0010,0000: The validator was unstaked
	Valided      = 0                         // 0000: The validator was activated
	NotExist     = 1 << 31                   // 1000,xxxx,... : The validator is not exist
)

type ValidatorStatus uint32

func (status ValidatorStatus) IsValid() bool {
	return !status.IsInvalid()
}

func (status ValidatorStatus) IsInvalid() bool {
	return status&Invalided == Invalided
}

func (status ValidatorStatus) IsOnlyInvalid() bool {
	return status&Invalided == status|Invalided
}

func (status ValidatorStatus) IsLowBlocks() bool {
	return status&LowBlocks == LowBlocks
}

func (status ValidatorStatus) IsOnlyLowBlocks() bool {
	return status&LowBlocks == status|LowBlocks
}

func (status ValidatorStatus) IsInvalidLowBlocks() bool {
	return status&(Invalided|LowBlocks) == (Invalided | LowBlocks)
}

func (status ValidatorStatus) IsLowThreshold() bool {
	return status&LowThreshold == LowThreshold
}

func (status ValidatorStatus) IsOnlyLowThreshold() bool {
	return status&LowThreshold == status|LowThreshold
}

func (status ValidatorStatus) IsInvalidLowThreshold() bool {
	return status&(Invalided|LowThreshold) == (Invalided | LowThreshold)
}

func (status ValidatorStatus) IsDuplicated() bool {
	return status&Duplicated == Duplicated
}

func (status ValidatorStatus) IsInvalidDuplicated() bool {
	return status&(Duplicated|Invalided) == (Duplicated | Invalided)
}

func (status ValidatorStatus) IsUnstaked() bool {
	return status&Unstaked == Unstaked
}

func (status ValidatorStatus) IsOnlyUnstaked() bool {
	return status&Unstaked == status|Unstaked
}

func (status ValidatorStatus) IsInvalidUnstaked() bool {
	return status&(Invalided|Unstaked) == (Invalided | Unstaked)
}

type Validator struct {
	Owner          common.Address
	StakeAmount    *big.Int
	DelegateAmount *big.Int
	PubKey         *ecdsa.PublicKey
	BlsKey         *bls.PublicKey
	Status         ValidatorStatus
	CommissionRate uint64
	Epoch          uint64
	StakeIndex     uint64
}

func NewValidator(owner common.Address, stakeAmount, delegateAmount *big.Int, blsKey *bls.PublicKey, pubKey *ecdsa.PublicKey, commissionRate, epoch, stakeIndex uint64) *Validator {
	return &Validator{
		Owner:          owner,
		StakeAmount:    stakeAmount,
		DelegateAmount: delegateAmount,
		PubKey:         pubKey,
		BlsKey:         blsKey,
		CommissionRate: commissionRate,
		Epoch:          epoch,
		StakeIndex:     stakeIndex,
	}
}

func (v *Validator) String() string {
	return fmt.Sprintf(`{"Owner": "%s","StakeAmount": "%d","DelegateAmount": "%d","PubKey": %s,"BlsKey": %s,"Status": %d,"CommissionRate": "%d", "Epoch": "%d", "StakeIndex": "%d"}`,
		fmt.Sprintf("%x", v.Owner.Bytes()),
		v.StakeAmount,
		v.DelegateAmount,
		hex.EncodeToString(crypto.FromECDSAPub(v.PubKey)),
		hex.EncodeToString(v.BlsKey.Serialize()),
		v.Status,
		v.CommissionRate,
		v.Epoch,
		v.StakeIndex)
}

func (v *Validator) IsNotEmpty() bool {
	return !v.IsEmpty()
}

func (v *Validator) IsEmpty() bool {
	return nil == v
}

func (v *Validator) SetStatus(status ValidatorStatus) {
	v.Status = status
}

func (v *Validator) AppendStatus(status ValidatorStatus) {
	v.Status |= status
}

func (v *Validator) CleanStatus(status ValidatorStatus) {
	v.Status &^= status
}

func (v *Validator) CleanStakeAmount() {
	v.StakeAmount = new(big.Int).SetInt64(0)
}

func (v *Validator) CleanDelegateAmount() {
	v.DelegateAmount = new(big.Int).SetInt64(0)
}

func (v *Validator) AddStakeAmount(amount *big.Int) {
	v.StakeAmount = new(big.Int).Add(v.StakeAmount, amount)
}

func (v *Validator) SubStakeAmount(amount *big.Int) error {
	if v.StakeAmount.Cmp(amount) == -1 {
		return errors.New("amount exceeds stake amount of validator")
	}
	v.StakeAmount = new(big.Int).Sub(v.StakeAmount, amount)
	return nil
}

func (v *Validator) AddDelegateAmount(amount *big.Int) {
	v.DelegateAmount = new(big.Int).Add(v.DelegateAmount, amount)
}

func (v *Validator) SubDelegateAmount(amount *big.Int) error {
	if v.DelegateAmount.Cmp(amount) == -1 {
		return errors.New("amount exceeds delegate amount of validator")
	}
	v.DelegateAmount = new(big.Int).Sub(v.DelegateAmount, amount)
	return nil
}

func (v *Validator) Shares() *big.Int {
	return new(big.Int).Add(v.StakeAmount, v.DelegateAmount)
}

func (v *Validator) IsValid() bool {
	return v.Status.IsValid()
}

func (v *Validator) IsInvalid() bool {
	return v.Status.IsInvalid()
}

func (v *Validator) IsOnlyInvalid() bool {
	return v.Status.IsOnlyInvalid()
}

func (v *Validator) IsLowBlocks() bool {
	return v.Status.IsLowBlocks()
}

func (v *Validator) IsOnlyLowBlocks() bool {
	return v.Status.IsOnlyLowBlocks()
}

func (v *Validator) IsInvalidLowBlocks() bool {
	return v.Status.IsInvalidLowBlocks()
}

func (v *Validator) IsLowThreshold() bool {
	return v.Status.IsLowThreshold()
}

func (v *Validator) IsOnlyLowThreshold() bool {
	return v.Status.IsOnlyLowThreshold()
}

func (v *Validator) IsInvalidLowThreshold() bool {
	return v.Status.IsInvalidLowThreshold()
}

func (v *Validator) IsDuplicated() bool {
	return v.Status.IsDuplicated()
}

func (v *Validator) IsInvalidDuplicated() bool {
	return v.Status.IsInvalidDuplicated()
}

func (v *Validator) IsUnstaked() bool {
	return v.Status.IsUnstaked()
}

func (v *Validator) IsOnlyUnstaked() bool {
	return v.Status.IsOnlyUnstaked()
}

func (v *Validator) IsInvalidUnstaked() bool {
	return v.Status.IsInvalidUnstaked()
}

type PriorityValidator struct {
	PreKey        []byte // previous priority validator key in statedb
	NextKey       []byte // next priority validator key in statedb
	ValidatorAddr common.Address
}

func NewPriorityValidator(preKey, nextKey []byte, addr common.Address) *PriorityValidator {
	return &PriorityValidator{
		PreKey:        preKey,
		NextKey:       nextKey,
		ValidatorAddr: addr,
	}
}

func (pv *PriorityValidator) UpdatePreKey(key []byte) {
	pv.PreKey = key
}

func (pv *PriorityValidator) UpdateNextKey(key []byte) {
	pv.NextKey = key
}

type ValidatorIds []common.Address

func (eids ValidatorIds) String() string {
	arr := make([]string, len(eids))
	for i, id := range eids {
		arr[i] = id.Hex()
	}
	return "[" + strings.Join(arr, ",") + "]"
}

func (eids ValidatorIds) Has(validatorId common.Address) bool {
	for _, id := range eids {
		if id == validatorId {
			return true
		}
	}
	return false
}

func (eids ValidatorIds) Remove(validatorIds ...common.Address) ValidatorIds {
	cache := make(map[common.Address]struct{}, 0)
	for _, id := range validatorIds {
		cache[id] = struct{}{}
	}
	for i := 0; i < len(eids); i++ {
		id := eids[i]
		if _, ok := cache[id]; ok {
			eids = append(eids[:i], eids[i+1:]...)
			i--
		}
	}
	return eids
}

type EpochItem struct {
	PreEpoch   uint64
	NextEpoch  uint64
	StartBlock uint64
	EndBlock   uint64
}

func NewEpochItem(preEpoch, nextEpoch, startBlock, endBlock uint64) *EpochItem {
	return &EpochItem{
		PreEpoch:   preEpoch,
		NextEpoch:  nextEpoch,
		StartBlock: startBlock,
		EndBlock:   endBlock,
	}
}

func (item *EpochItem) UpdatePreEpoch(epoch uint64) {
	item.PreEpoch = epoch
}

func (item *EpochItem) UpdateNextEpoch(epoch uint64) {
	item.NextEpoch = epoch
}

func (item *EpochItem) IsEmpty() bool {
	return nil == item
}

func (item *EpochItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type EpochQueue []*EpochItem

func NewEpochQueue(size uint64) EpochQueue {
	queue := make(EpochQueue, size)
	return queue
}

func (queue EpochQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue EpochQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

type StakeWithdrawalItem struct {
	PreEpoch  uint64   // pre release epoch
	NextEpoch uint64   // next release epoch
	Amount    *big.Int // withdraw amount
}

func NewStakeWithdrawalItem(preEpoch, nextEpoch uint64, amount *big.Int) *StakeWithdrawalItem {
	return &StakeWithdrawalItem{
		PreEpoch:  preEpoch,
		NextEpoch: nextEpoch,
		Amount:    amount,
	}
}

func (item *StakeWithdrawalItem) UpdatePreEpoch(epoch uint64) {
	item.PreEpoch = epoch
}

func (item *StakeWithdrawalItem) UpdateNextEpoch(epoch uint64) {
	item.NextEpoch = epoch
}

func (item *StakeWithdrawalItem) IncrementAmount(increment *big.Int) {
	item.Amount = new(big.Int).Add(item.Amount, increment)
}

func (item *StakeWithdrawalItem) IsEmpty() bool {
	return nil == item
}

func (item *StakeWithdrawalItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

// -------------

type DelegateWithdrawalItem struct {
	PreEpoch  uint64   // pre release epoch
	NextEpoch uint64   // next release epoch
	Amount    *big.Int // withdraw amount
}

func NewDelegateWithdrawalItem(preEpoch, nextEpoch uint64, amount *big.Int) *DelegateWithdrawalItem {
	return &DelegateWithdrawalItem{
		PreEpoch:  preEpoch,
		NextEpoch: nextEpoch,
		Amount:    amount,
	}
}

func (item *DelegateWithdrawalItem) UpdatePreEpoch(epoch uint64) {
	item.PreEpoch = epoch
}

func (item *DelegateWithdrawalItem) UpdateNextEpoch(epoch uint64) {
	item.NextEpoch = epoch
}

func (item *DelegateWithdrawalItem) IncrementAmount(increment *big.Int) {
	item.Amount = new(big.Int).Add(item.Amount, increment)
}

type UnStakeDelegationRcItem struct {
	PreStakeEpoch  uint64
	NextStakeEpoch uint64
	// This means how many delegates in the current `stakeEpoch` have not been fully withdrawn
	// eg. delegaterA:validatorA:stakeEpoch(100)、 ...、 delegaterN:validatorA:stakeEpoch(100)
	Rc uint64
}

func NewUnStakeDelegationRcItem(preStakeEpoch, nextStakeEpoch, delegationRc uint64) *UnStakeDelegationRcItem {
	return &UnStakeDelegationRcItem{
		PreStakeEpoch:  preStakeEpoch,
		NextStakeEpoch: nextStakeEpoch,
		Rc:             delegationRc,
	}
}

func (item *UnStakeDelegationRcItem) UpdatePreStakeEpoch(stakeEpoch uint64) {
	item.PreStakeEpoch = stakeEpoch
}

func (item *UnStakeDelegationRcItem) UpdateNextStakeEpoch(stakeEpoch uint64) {
	item.NextStakeEpoch = stakeEpoch
}

func (item *UnStakeDelegationRcItem) IncrementRc(increment uint64) {
	item.Rc += increment
}

func (item *UnStakeDelegationRcItem) DecrementRc(decrement uint64) {
	if item.Rc < decrement {
		item.Rc = 0
	} else {
		item.Rc -= decrement
	}
}

type UnStakeDelegationRcQueue []*UnStakeDelegationRcItem

func NewUnStakeDelegationRcQueue(size uint64) UnStakeDelegationRcQueue {
	queue := make(UnStakeDelegationRcQueue, size)
	return queue
}

func (queue UnStakeDelegationRcQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue UnStakeDelegationRcQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}
