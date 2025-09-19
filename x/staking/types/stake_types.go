package types

import (
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/p2p/enode"
	"github.com/status-im/keycard-go/hexutils"
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
	Slashing                                 // 0100,0000: The validator is being slashed
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
func (status ValidatorStatus) IsOnlyInvalidLowBlocks() bool {
	return status&(Invalided|LowBlocks) == status|(Invalided|LowBlocks)
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

func (status ValidatorStatus) IsUnstaked() bool { return status&Unstaked == Unstaked }
func (status ValidatorStatus) IsOnlyUnstaked() bool {
	return status&Unstaked == status|Unstaked
}
func (status ValidatorStatus) IsInvalidUnstaked() bool {
	return status&(Invalided|Unstaked) == (Invalided | Unstaked)
}
func (status ValidatorStatus) IsOnlyInvalidUnstaked() bool {
	return status&(Invalided|Unstaked) == status|(Invalided|Unstaked)
}

func (status ValidatorStatus) IsSlashing() bool     { return status&Slashing == Slashing }
func (status ValidatorStatus) IsOnlySlashing() bool { return status&Slashing == status|Slashing }
func (status ValidatorStatus) IsInvalidSlashing() bool {
	return status&(Invalided|Slashing) == (Invalided | Slashing)
}
func (status ValidatorStatus) IsOnlyInvalidSlashing() bool {
	return status&(Invalided|Slashing) == status|(Invalided|Slashing)
}

func (status ValidatorStatus) IsNotExist() bool {
	return status&NotExist == NotExist
}

type Validator struct {
	Status         ValidatorStatus
	CommissionRate uint64
	Epoch          uint64
	StakeIndex     uint64
	Owner          common.Address
	StakeAmount    *big.Int
	DelegateAmount *big.Int
	//PubKey         *ecdsa.PublicKey
	//BlsKey         *bls.PublicKey
	PubKey enode.IDv0
	BlsKey []byte
}

func NewValidator(owner common.Address, stakeAmount, delegateAmount *big.Int, blsKey []byte, pubKey enode.IDv0, commissionRate, epoch, stakeIndex uint64) *Validator {
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
	//blsKey := bls.PublicKey{}
	//(&blsKey).Deserialize(v.BlsKey)
	return fmt.Sprintf(`{"Owner": "%s","StakeAmount": "%d","DelegateAmount": "%d","PubKey": %s,"BlsKey": %s,"Status": %d,"CommissionRate": "%d", "Epoch": "%d", "StakeIndex": "%d"}`,
		fmt.Sprintf("%x", v.Owner.Bytes()),
		v.StakeAmount,
		v.DelegateAmount,
		fmt.Sprintf("%x", v.PubKey.Bytes()),
		hexutils.BytesToHex(v.BlsKey),
		//hex.EncodeToString(crypto.FromECDSAPub(v.PubKey)),
		//hex.EncodeToString((&blsKey).Serialize()),

		v.Status,
		v.CommissionRate,
		v.Epoch,
		v.StakeIndex)
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
	return v.IsNotEmpty() && v.Status.IsValid()
}
func (v *Validator) IsEmptyOrInvalid() bool {
	return v.IsEmpty() || (v.IsNotEmpty() && v.Status.IsInvalid())
}
func (v *Validator) IsInvalid() bool {
	return v.IsNotEmpty() && v.Status.IsInvalid()
}
func (v *Validator) IsOnlyInvalid() bool {
	return v.IsNotEmpty() && v.Status.IsOnlyInvalid()
}

func (v *Validator) IsLowBlocks() bool {
	return v.IsNotEmpty() && v.Status.IsLowBlocks()
}
func (v *Validator) IsOnlyLowBlocks() bool {
	return v.IsNotEmpty() && v.Status.IsOnlyLowBlocks()
}
func (v *Validator) IsInvalidLowBlocks() bool {
	return v.IsNotEmpty() && v.Status.IsInvalidLowBlocks()
}
func (v *Validator) IsOnlyInvalidLowBlocks() bool {
	return v.IsNotEmpty() && v.Status.IsOnlyInvalidLowBlocks()
}

func (v *Validator) IsLowThreshold() bool {
	return v.IsNotEmpty() && v.Status.IsLowThreshold()
}
func (v *Validator) IsOnlyLowThreshold() bool {
	return v.IsNotEmpty() && v.Status.IsOnlyLowThreshold()
}
func (v *Validator) IsInvalidLowThreshold() bool {
	return v.IsNotEmpty() && v.Status.IsInvalidLowThreshold()
}

func (v *Validator) IsDuplicated() bool {
	return v.IsNotEmpty() && v.Status.IsDuplicated()
}
func (v *Validator) IsInvalidDuplicated() bool {
	return v.IsNotEmpty() && v.Status.IsInvalidDuplicated()
}

func (v *Validator) IsUnstaked() bool {
	return v.IsNotEmpty() && v.Status.IsUnstaked()
}
func (v *Validator) IsOnlyUnstaked() bool {
	return v.IsNotEmpty() && v.Status.IsOnlyUnstaked()
}
func (v *Validator) IsInvalidUnstaked() bool {
	return v.IsNotEmpty() && v.Status.IsInvalidUnstaked()
}
func (v *Validator) IsOnlyInvalidUnstaked() bool {
	return v.IsNotEmpty() && v.Status.IsOnlyInvalidUnstaked()
}

func (v *Validator) IsSlashing() bool {
	return v.IsNotEmpty() && v.Status.IsSlashing()
}
func (v *Validator) IsOnlySlashing() bool {
	return v.IsNotEmpty() && v.Status.IsOnlySlashing()
}
func (v *Validator) IsInvalidSlashing() bool {
	return v.IsNotEmpty() && v.Status.IsInvalidSlashing()
}
func (v *Validator) IsOnlyInvalidSlashing() bool {
	return v.IsNotEmpty() && v.Status.IsOnlyInvalidSlashing()
}

func (v *Validator) IsEmpty() bool {
	return nil == v
}

func (v *Validator) IsNotEmpty() bool {
	return !v.IsEmpty()
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

func (pv *PriorityValidator) IsEmpty() bool {
	return nil == pv
}

func (pv *PriorityValidator) IsNotEmpty() bool {
	return !pv.IsEmpty()
}

type ValidatorAddrQueue []common.Address

func NewValidatorAddrQueue(size uint64) ValidatorAddrQueue {
	queue := make(ValidatorAddrQueue, size)
	return queue
}

func (ids ValidatorAddrQueue) String() string {
	arr := make([]string, len(ids))
	for i, id := range ids {
		arr[i] = id.Hex()
	}
	return "[" + strings.Join(arr, ",") + "]"
}

func (ids ValidatorAddrQueue) Has(validatorId common.Address) bool {
	for _, id := range ids {
		if id == validatorId {
			return true
		}
	}
	return false
}

func (ids ValidatorAddrQueue) Remove(validatorIds ...common.Address) ValidatorAddrQueue {
	cache := make(map[common.Address]struct{}, 0)
	for _, id := range validatorIds {
		cache[id] = struct{}{}
	}
	for i := 0; i < len(ids); i++ {
		id := ids[i]
		if _, ok := cache[id]; ok {
			ids = append(ids[:i], ids[i+1:]...)
			i--
		}
	}
	return ids
}

func (ids ValidatorAddrQueue) IsEmpty() bool {
	return len(ids) == 0
}

func (ids ValidatorAddrQueue) IsNotEmpty() bool {
	return !ids.IsEmpty()
}

// used for sorting round validators in elections
type ValidatorSortSnapshot struct {
	ValidatorAddr  common.Address
	Epoch          uint64
	StakeIndex     uint64
	ValidatorTerm  uint64
	CommissionRate uint64
	StakeAmount    *big.Int
	DelegateAmount *big.Int
}

func NewValidatorSharesSnapshot(validatorAddr common.Address, epoch, stakeIndex, commissionRate uint64, stakeAmount, delegateAmount *big.Int) *ValidatorSortSnapshot {
	return &ValidatorSortSnapshot{
		ValidatorAddr:  validatorAddr,
		Epoch:          epoch,
		StakeIndex:     stakeIndex,
		CommissionRate: commissionRate,
		StakeAmount:    stakeAmount,
		DelegateAmount: delegateAmount,
	}
}

func (item *ValidatorSortSnapshot) Shares() *big.Int {
	return new(big.Int).Add(item.StakeAmount, item.DelegateAmount)
}

func (item *ValidatorSortSnapshot) IncrementValidatorTerm(increment uint64) {
	item.ValidatorTerm += increment
}

func (item *ValidatorSortSnapshot) IsEmpty() bool {
	return nil == item
}

func (item *ValidatorSortSnapshot) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type ValidatorSortSnapshotQueue []*ValidatorSortSnapshot

func NewValidatorSharesSnapshotQueue(size uint64) ValidatorSortSnapshotQueue {
	queue := make(ValidatorSortSnapshotQueue, size)
	return queue
}

func (queue ValidatorSortSnapshotQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue ValidatorSortSnapshotQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}

// prioty order: large -> small
// eg. 8 -> 5- > 3
func (queue ValidatorSortSnapshotQueue) ValidatorSort(cache MaybeRemoveValidatorStatusCache,
	compare func(cache MaybeRemoveValidatorStatusCache, c, can *ValidatorSortSnapshot) int) {
	if len(queue) <= 1 {
		return
	}

	if nil == compare {
		queue.quickSort(cache, 0, len(queue)-1, CompareDefault)
	} else {
		queue.quickSort(cache, 0, len(queue)-1, compare)
	}
}
func (queue ValidatorSortSnapshotQueue) quickSort(cache MaybeRemoveValidatorStatusCache, left, right int,
	compare func(cache MaybeRemoveValidatorStatusCache, c, can *ValidatorSortSnapshot) int) {
	if left < right {
		pivot := queue.partition(cache, left, right, compare)
		queue.quickSort(cache, left, pivot-1, compare)
		queue.quickSort(cache, pivot+1, right, compare)
	}
}
func (queue ValidatorSortSnapshotQueue) partition(cache MaybeRemoveValidatorStatusCache, left, right int,
	compare func(cache MaybeRemoveValidatorStatusCache, c, can *ValidatorSortSnapshot) int) int {
	for left < right {
		for left < right && compare(cache, queue[left], queue[right]) >= 0 {
			right--
		}
		if left < right {
			queue[left], queue[right] = queue[right], queue[left]
			left++
		}
		for left < right && compare(cache, queue[left], queue[right]) >= 0 {
			left++
		}
		if left < right {
			queue[left], queue[right] = queue[right], queue[left]
			right--
		}
	}
	return left
}

// #### NOTE: ####
// Sort By Default
// ###############
//
// order rules:
//
//	ValidaotorTerm -> Shares > Epoch > StakeIndex > Invalid
//
// (no remove -> maybe remove)
//
// eg:  term(2 -> 5) -> shares (200w -> 100w ) -> epoch (2<early> -> 10<later>) -> stakeIndex(3<early> -> 8<later>) -> status(invalid)
//
// Compare Left And Right
// 1: Left > Right	=> Left -> Right
// 0: Left == Right	=> Left(Right)
// -1:Left < Right	=> Right -> Left
func CompareDefault(cache MaybeRemoveValidatorStatusCache, left, right *ValidatorSortSnapshot) int {

	//
	compareStakeIndexFunc := func(l, r *ValidatorSortSnapshot) int {
		switch {
		case l.StakeIndex > r.StakeIndex:
			return -1
		case l.StakeIndex < r.StakeIndex:
			return 1
		default:
			return 0
		}
	}

	compareEpochFunc := func(l, r *ValidatorSortSnapshot) int {

		switch {
		case l.Epoch > r.Epoch:
			return -1
		case l.Epoch < r.Epoch:
			return 1
		default:
			return compareStakeIndexFunc(l, r)
		}
	}

	compareSharesFunc := func(l, r *ValidatorSortSnapshot) int {

		switch {
		case l.Shares().Cmp(r.Shares()) < 0:
			return -1
		case l.Shares().Cmp(r.Shares()) > 0:
			return 1
		default:
			return compareEpochFunc(l, r)
		}
	}

	// Compare Term
	compareTermFunc := func(l, r *ValidatorSortSnapshot) int {
		switch {
		case l.ValidatorTerm < r.ValidatorTerm:
			return 1
		case l.ValidatorTerm > r.ValidatorTerm:
			return -1
		default:
			return compareSharesFunc(l, r)
		}
	}

	_, leftOk := cache[left.ValidatorAddr]
	_, rightOk := cache[right.ValidatorAddr]

	if leftOk && !rightOk {
		return -1
	} else if !leftOk && rightOk {
		return 1
	} else {

		return compareTermFunc(left, right)
	}

}

// #### NOTE: ####
// These are sorted by priority that will be removed
// ###############
//
// order rules:
//
// Invalid > ValidaotorTerm  > Shares > Epoch > StakeIndex
// (maybe remove -> no remove)
//
// eg: status(duplicated -> lowBlocks -> unstake -> valid) -> term(5 -> 2) -> shares (100w -> 200 w) -> epoch (10<later> -> 2<early>) -> stakeIndex(8<later> -> 3<early>)
//
// What is the invalid ?  That are slashed and withdrew&NotInEpochValidators
//
// Compare Left And Right
// 1: Left > Right   => Left -> Right
// 0: Left == Right  => Left(Right)
// -1:Left < Right   => Right -> Left
func CompareForRemoveFromHead(cache MaybeRemoveValidatorStatusCache, left, right *ValidatorSortSnapshot) int {

	// Compare StakeIndex
	compareStakeIndexFunc := func(l, r *ValidatorSortSnapshot) int {

		switch {
		case l.StakeIndex > r.StakeIndex:
			return 1
		case l.StakeIndex < r.StakeIndex:
			return -1
		default:
			return 0
		}
	}

	// Compare epoch
	compareEpochFunc := func(l, r *ValidatorSortSnapshot) int {
		switch {
		case l.Epoch > r.Epoch:
			return 1
		case l.Epoch < r.Epoch:
			return -1
		default:
			return compareStakeIndexFunc(l, r)
		}
	}

	// Compare Shares
	compareSharesFunc := func(l, r *ValidatorSortSnapshot) int {

		switch {
		case l.Shares().Cmp(r.Shares()) < 0:
			return 1
		case l.Shares().Cmp(r.Shares()) > 0:
			return -1
		default:
			return compareEpochFunc(l, r)
		}
	}

	// Compare Term
	compareTermFunc := func(l, r *ValidatorSortSnapshot) int {
		switch {
		case l.ValidatorTerm < r.ValidatorTerm:
			return -1
		case l.ValidatorTerm > r.ValidatorTerm:
			return 1
		default:
			return compareSharesFunc(l, r)
		}
	}

	lstatus, lok := cache[left.ValidatorAddr]
	rstatus, rok := cache[right.ValidatorAddr]

	/**
	Start Compare
	*/

	switch {
	case !lok && rok: // left need not removed AND right need removed   right -> left
		return -1
	case !lok && !rok: // both need not removed

		return compareTermFunc(left, right)

	case lok && !rok: // left need removed AND right need not removed
		return 1
	default: // both need removed

		// notexist -> exist
		if lstatus.IsNotExist() && !rstatus.IsNotExist() {
			return 1
		} else if !lstatus.IsNotExist() && rstatus.IsNotExist() {
			return -1
		} else {
			// duplicated -> unduplicated
			switch {
			case lstatus.IsInvalidDuplicated() && !rstatus.IsInvalidDuplicated():
				return 1
			case !lstatus.IsInvalidDuplicated() && rstatus.IsInvalidDuplicated():
				return -1
			case lstatus.IsInvalidDuplicated() && rstatus.IsInvalidDuplicated():
				// compare shares
				return compareSharesFunc(left, right)
			default:

				// lowBlocks -> unlowBlocks
				switch {
				case lstatus.IsInvalidLowBlocks() && !rstatus.IsInvalidLowBlocks():
					return 1
				case !lstatus.IsInvalidLowBlocks() && rstatus.IsInvalidLowBlocks():
					return -1
				case lstatus.IsInvalidLowBlocks() && rstatus.IsInvalidLowBlocks():
					// compare shares
					return compareSharesFunc(left, right)
				default:

					// unstake -> valid
					if lstatus.IsInvalidUnstaked() && !rstatus.IsInvalidUnstaked() {
						return 1
					} else if !lstatus.IsInvalidUnstaked() && rstatus.IsInvalidUnstaked() {
						return -1
					} else {
						// compare term
						return compareTermFunc(left, right)
					}
				}
			}
		}
	}
}

type MaybeRemoveValidatorStatusCache map[common.Address]ValidatorStatus

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

func (item *DelegateWithdrawalItem) IsEmpty() bool {
	return nil == item
}

func (item *DelegateWithdrawalItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type ValidatorDelegationRcItem struct {
	PreStakeEpoch  uint64
	NextStakeEpoch uint64
	// This means how many delegates in the current `stakeEpoch` have not been fully withdrawn
	// eg. delegatorA:validatorA:stakeEpoch(100)、 ...、 delegatorN:validatorA:stakeEpoch(100)
	Rc uint64
}

func NewValidatorDelegationRcItem(preStakeEpoch, nextStakeEpoch, delegationRc uint64) *ValidatorDelegationRcItem {
	return &ValidatorDelegationRcItem{
		PreStakeEpoch:  preStakeEpoch,
		NextStakeEpoch: nextStakeEpoch,
		Rc:             delegationRc,
	}
}

func (item *ValidatorDelegationRcItem) UpdatePreStakeEpoch(stakeEpoch uint64) {
	item.PreStakeEpoch = stakeEpoch
}

func (item *ValidatorDelegationRcItem) UpdateNextStakeEpoch(stakeEpoch uint64) {
	item.NextStakeEpoch = stakeEpoch
}

func (item *ValidatorDelegationRcItem) IncrementRc(increment uint64) {
	item.Rc += increment
}

func (item *ValidatorDelegationRcItem) DecrementRc(decrement uint64) {
	if item.Rc < decrement {
		item.Rc = 0
	} else {
		item.Rc -= decrement
	}
}

func (item *ValidatorDelegationRcItem) IsEmpty() bool {
	return nil == item
}

func (item *ValidatorDelegationRcItem) IsNotEmpty() bool {
	return !item.IsEmpty()
}

type ValidatorDelegationRcQueue []*ValidatorDelegationRcItem

func NewValidatorDelegationRcQueue(size uint64) ValidatorDelegationRcQueue {
	queue := make(ValidatorDelegationRcQueue, size)
	return queue
}

func (queue ValidatorDelegationRcQueue) IsEmpty() bool {
	return len(queue) == 0
}

func (queue ValidatorDelegationRcQueue) IsNotEmpty() bool {
	return !queue.IsEmpty()
}
