package contracts

import (
	"fmt"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
	"github.com/status-im/keycard-go/hexutils"
	"github.com/test-go/testify/assert"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func Test_Delegate(t *testing.T) {

	//validatorAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	//delegatorAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")
	//ownerAddr := common.HexToAddress("0x3333333333333333333333333333333333333333")
	//stakeHandler := newStakeHandler(common.ZeroAddr)
	//var err error
	//err = stakeHandler.stake(validatorAddr, ownerAddr, common.Big100, 10, []byte{}, enode.IDv0{})
	//if nil != err {
	//	t.Error(err)
	//}
	//err = stakeHandler.delegate(validatorAddr, delegatorAddr, common.Big32)
	//if nil != err {
	//	t.Error(err)
	//}
	//
	//delegations, err := stakeHandler.GetDelegationsWithValidator([]common.Address{validatorAddr}, delegatorAddr)
	//if nil != err {
	//	t.Error(err)
	//}
	//
	//t.Log("delegation size", len(delegations))
}

func Test_StakeFor(t *testing.T) {
	stakeHandlerTestConfig := newStakeHandlerTestConfig()
	epoch := uint64(1)
	from := common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10")

	// init stakeHandler
	stakeHandler := newStakeHandler((stakeHandlerTestConfig.StateDB).(vm.StateDB), from, new(big.Int).SetUint64(2))
	initStakeHandler(stakeHandler, stakeHandlerTestConfig.L1Module, stakeHandlerTestConfig.StageModule, stakeHandlerTestConfig.StakeModule, stakeHandlerTestConfig.RewardModule)

	// mock init genesis
	stakeHandlerTestConfig.StageModule.MockCurrentEpoch(epoch)
	err := stakeHandlerTestConfig.StakeModule.MockInitValidatorGenesisPriority()
	assert.Nil(t, err, "Failed to call MockInitValidatorGenesisPriority")
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

	stakeHandlerTestConfig := newStakeHandlerTestConfig()
	epoch := uint64(1)
	from := common.HexToAddress("0xB0568bF61e3E7AF10623054b9169Eb8030271D10")

	// init stakeHandler
	stakeHandler := newStakeHandler((stakeHandlerTestConfig.StateDB).(vm.StateDB), from, new(big.Int).SetUint64(2))
	initStakeHandler(stakeHandler, stakeHandlerTestConfig.L1Module, stakeHandlerTestConfig.StageModule, stakeHandlerTestConfig.StakeModule, stakeHandlerTestConfig.RewardModule)

	// mock init genesis
	stakeHandlerTestConfig.StageModule.MockCurrentEpoch(epoch)
	err := stakeHandlerTestConfig.StakeModule.MockInitValidatorGenesisPriority()
	assert.Nil(t, err, "Failed to call MockInitValidatorGenesisPriority")
	// get mock data
	_, stakeDatas := extractStakeForDataList()
	stakeFor(t, stakeHandler, stakeDatas, false)

	validatorAddrs, addStakeDatas := extractAddStakeDataList()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := r.Intn(len(validatorAddrs)-1) + 1

	cache := make(map[common.Address]*big.Int, num)

	for _, validatorAddr := range validatorAddrs[:num] {
		validator := stakeHandler.getValidator(validatorAddr)
		assert.False(t, validator.IsEmpty(), fmt.Sprintf("validator is empty, %s", validatorAddr.Hex()))

		cache[validatorAddr] = validator.StakeAmount
	}

	addStake(t, stakeHandler, addStakeDatas[:num], true)

	addStakeDataCache := getAddStakeDataCache()
	for validatorAddr, amount := range cache {
		validator := stakeHandler.getValidator(validatorAddr)
		assert.Equal(t, new(big.Int).Add(addStakeDataCache[validatorAddr].Amount, amount), validator.StakeAmount, "STAKE AMOUNT MISMATCHING")
	}

}
