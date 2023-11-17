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
	Benefit        common.Address
	StakeAmount    *big.Int
	DelegateAmount *big.Int
	PubKey         *ecdsa.PublicKey
	BlsKey         *bls.PublicKey
	Status         ValidatorStatus

	StakeIndex uint64
}

func NewValidator(benefit common.Address, stakeAmount, delegateAmount *big.Int, blsKey *bls.PublicKey, pubKey *ecdsa.PublicKey) *Validator {
	return &Validator{
		Benefit:        benefit,
		StakeAmount:    stakeAmount,
		DelegateAmount: delegateAmount,
		BlsKey:         blsKey,
		PubKey:         pubKey,
	}
}

func (v *Validator) String() string {
	return fmt.Sprintf(`{"Benefit": "%s","StakeAmount": "%d","DelegateAmount": "%d","PubKey": %s,"BlsKey": %s,"Status": %d,"StakeIndex": "%d"}`,
		//fmt.Sprintf("%x", v.Addr.Bytes()),
		fmt.Sprintf("%x", v.Benefit.Bytes()),
		v.StakeAmount,
		v.DelegateAmount,
		hex.EncodeToString(crypto.FromECDSAPub(v.PubKey)),
		hex.EncodeToString(v.BlsKey.Serialize()),
		v.Status,
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
	Previous      []byte // previous priority validator key in statedb
	Next          []byte // next priority validator key in statedb
	ValidatorAddr common.Address
}

func NewPriorityValidator(previous, next []byte, addr common.Address) *PriorityValidator {
	return &PriorityValidator{
		Previous:      previous,
		Next:          next,
		ValidatorAddr: addr,
	}
}

func (pv *PriorityValidator) UpdatePrevious(previous []byte) {
	pv.Previous = previous
}

func (pv *PriorityValidator) UpdateNext(next []byte) {
	pv.Next = next
}

type ValidatorIds []common.Address

func (eids ValidatorIds) String() string {
	arr := make([]string, len(eids))
	for i, id := range eids {
		arr[i] = id.Hex()
	}
	return "[" + strings.Join(arr, ",") + "]"
}

func (eids ValidatorIds) Append(validatorId common.Address) ValidatorIds {
	eids = append(eids, validatorId)
	return eids
}

func (eids ValidatorIds) Has(validatorId common.Address) bool {
	for _, id := range eids {
		if id == validatorId {
			return true
		}
	}
	return false
}

//
//func (eids ValidatorIds) remove(validatorId common.Address) ValidatorIds {
//	for i := 0; i < len(eids); i++ {
//		if eids[i] == validatorId {
//			eids = append(eids[:i], eids[i+1:]...)
//			break
//		}
//	}
//	return eids
//}

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

type Epoch struct {
	Start uint64
	End   uint64
}

type StakeWithdrawalBound struct {
	Head uint64
	Tail uint64
}

func NewStakeWithdrawalBound(head, tail uint64) *StakeWithdrawalBound {
	return &StakeWithdrawalBound{
		Head: head,
		Tail: tail,
	}
}

func (bound *StakeWithdrawalBound) UpdateHead(newHead uint64) {
	bound.Head = newHead
}

func (bound *StakeWithdrawalBound) UpdateTail(newTail uint64) {
	bound.Tail = newTail
}

func (bound *StakeWithdrawalBound) IncrementHead(increment uint64) {
	bound.Head += increment
}

func (bound *StakeWithdrawalBound) IncrementTail(increment uint64) {
	bound.Tail += increment
}

type StakeWithdrawalItem struct {
	Epoch  uint64   // release epoch
	Amount *big.Int // withdraw amount
}

func NewStakeWithdrawalItem(epoch uint64, amount *big.Int) *StakeWithdrawalItem {
	return &StakeWithdrawalItem{
		Epoch:  epoch,
		Amount: amount,
	}
}
func (item *StakeWithdrawalItem) IncrementAmount(increment *big.Int) {
	item.Amount = new(big.Int).Add(item.Amount, increment)
}

// -------------

type DelegateWithdrawalBound struct {
	Head uint64
	Tail uint64
}

func NewDelegateWithdrawalBound(head, tail uint64) *DelegateWithdrawalBound {
	return &DelegateWithdrawalBound{
		Head: head,
		Tail: tail,
	}
}

func (bound *DelegateWithdrawalBound) UpdateHead(newHead uint64) {
	bound.Head = newHead
}

func (bound *DelegateWithdrawalBound) UpdateTail(newTail uint64) {
	bound.Tail = newTail
}

func (bound *DelegateWithdrawalBound) IncrementHead(increment uint64) {
	bound.Head += increment
}

func (bound *DelegateWithdrawalBound) IncrementTail(increment uint64) {
	bound.Tail += increment
}

type DelegateWithdrawalItem struct {
	Epoch  uint64   // release epoch
	Amount *big.Int // withdraw amount
}

func NewDelegateWithdrawalItem(epoch uint64, amount *big.Int) *DelegateWithdrawalItem {
	return &DelegateWithdrawalItem{
		Epoch:  epoch,
		Amount: amount,
	}
}
func (item *DelegateWithdrawalItem) IncrementAmount(increment *big.Int) {
	item.Amount = new(big.Int).Add(item.Amount, increment)
}
