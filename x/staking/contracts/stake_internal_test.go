package contracts

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/test-go/testify/assert"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func Test_StakeFor(t *testing.T) {

	epoch := uint64(1)
	stakeHandler := stakePrepare(t, epoch, common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10"), true)
	// get mock data
	_, datas := extractStakeForDataList()
	stakeFor(t, stakeHandler, datas, true)
	//
	next, queue, err := stakeHandler.GetValidators([]byte{}, common.Big100)
	//t.Log("next", hexutils.BytesToHex(next))
	//t.Log("queue", fmt.Sprintf("%v", queue))
	//t.Log("err", err)
	assert.Nil(t, err, "Failed to call GetValidators")
	assert.Equal(t, 4, len(queue), "mismatching validator arr size")
	assert.Equal(t, []byte("priorityValidatorTail"), next, "mismatching next priority key")

	cache := getStakeForDataCache()
	for i, validator := range queue {
		item, ok := cache[validator.ValidatorAddr]
		assert.True(t, ok, "Not found validator")
		assert.Equal(t, item.Owner, validator.Owner, fmt.Sprintf("mismatching owner, expect: %s, actual: %s", item.Owner.Hex(), validator.Owner.Hex()))
		assert.Equal(t, item.StakeAmount, validator.StakeAmount, fmt.Sprintf("mismatching stakeAmount, expect: %d, actual: %d", item.StakeAmount, validator.StakeAmount))
		assert.Equal(t, item.CommissionRate, validator.CommissionRate, fmt.Sprintf("mismatching commissionRate, expect: %d, actual: %d", item.CommissionRate, validator.CommissionRate))
		assert.Equal(t, common.Big0, validator.Status, fmt.Sprintf("mismatching status, expect: %d, actual: %d", common.Big0, validator.Status))
		assert.Equal(t, epoch, validator.Epoch.Uint64(), fmt.Sprintf("mismatching epoch, expect: %d, actual: %d", epoch, validator.Epoch))
		assert.Equal(t, uint64(i), validator.StakeIndex.Uint64(), fmt.Sprintf("mismatching stakeIndex, expect: %d, actual: %d", i, validator.StakeIndex))
		assert.Equal(t, item.PubKey, validator.PubKey, fmt.Sprintf("mismatching pubKey, expect: %s, actual: %s", hexutils.BytesToHex(item.PubKey), hexutils.BytesToHex(validator.PubKey)))
		assert.Equal(t, item.BlsKey, validator.BlsKey, fmt.Sprintf("mismatching blsKey, expect: %s, actual: %s", hexutils.BytesToHex(item.BlsKey), hexutils.BytesToHex(validator.BlsKey)))
	}
}

func Test_AddStake(t *testing.T) {

	epoch := uint64(1)
	stakeHandler := stakePrepare(t, epoch, common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10"), true)

	// get mock data
	_, stakeDatas := extractStakeForDataList()
	stakeFor(t, stakeHandler, stakeDatas, false)

	validatorAddrs, addStakeDatas := extractAddStakeDataList()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(validatorAddrs)-1) + 1

	validatorStakeAmountCache := make(map[common.Address]*big.Int, num)

	for _, validatorAddr := range validatorAddrs[:num] {
		validator := stakeHandler.getValidator(validatorAddr)
		assert.False(t, validator.IsEmpty(), fmt.Sprintf("validator is empty, %s", validatorAddr.Hex()))

		validatorStakeAmountCache[validatorAddr] = validator.StakeAmount
	}

	addStake(t, stakeHandler, addStakeDatas[:num], true)

	addStakeDataCache := getAddStakeDataCache()
	for validatorAddr, amount := range validatorStakeAmountCache {
		validator := stakeHandler.getValidator(validatorAddr)
		assert.Equal(t, new(big.Int).Add(addStakeDataCache[validatorAddr].Amount, amount), validator.StakeAmount, "STAKE AMOUNT MISMATCHING")
	}

}

func Test_Delegate(t *testing.T) {

	epoch := uint64(1)
	stakeHandler := stakePrepare(t, epoch, common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10"), true)

	// get mock data
	validatorAddrs, stakeDatas := extractStakeForDataList()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(validatorAddrs)-1) + 1

	stakeFor(t, stakeHandler, stakeDatas[:num], true)

	subValidatorAddrs := validatorAddrs[:num]

	validatorDelegateAmountCache := make(map[common.Address]*big.Int, num)

	for _, validatorAddr := range subValidatorAddrs {
		validator := stakeHandler.getValidator(validatorAddr)
		assert.False(t, validator.IsEmpty(), fmt.Sprintf("validator is empty, %s", validatorAddr.Hex()))

		validatorDelegateAmountCache[validatorAddr] = validator.DelegateAmount

	}

	// delegate for
	delegatorAddrs := extractDelegateDatList()

	num = r.Intn(len(delegatorAddrs)-1) + 1

	delegateRelationShipCache := make(map[common.Address][]common.Address, 0)

	delegateDataCache := getDelegateDataCache()

	for _, delegator := range delegatorAddrs[:num] {

		// delegate for
		var validator common.Address
		if len(subValidatorAddrs) == 1 {
			validator = subValidatorAddrs[0]
		} else {
			validator = subValidatorAddrs[r.Intn(len(subValidatorAddrs)-1)]
		}

		delegate(t, stakeHandler, encodeDelegateData(validator, delegateDataCache[delegator]), true)

		delegatorArr, ok := delegateRelationShipCache[validator]
		if !ok {
			delegatorArr = make([]common.Address, 0)
		}
		delegatorArr = append(delegatorArr, delegator)
		delegateRelationShipCache[validator] = delegatorArr
	}

	for validatorAddr, delegatorArr := range delegateRelationShipCache {

		validator := stakeHandler.getValidator(validatorAddr)

		oldDelegateAmount := validatorDelegateAmountCache[validatorAddr]

		allDelegatorAmount := common.Big0

		for _, delegator := range delegatorArr {
			allDelegatorAmount = new(big.Int).Add(allDelegatorAmount, delegateDataCache[delegator].Amount)
		}

		assert.Equal(t, new(big.Int).Add(oldDelegateAmount, allDelegatorAmount), validator.DelegateAmount, "DELEGATE AMOUNT MISMATCHING")
	}
}
