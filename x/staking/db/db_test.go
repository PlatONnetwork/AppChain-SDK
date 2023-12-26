package db

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/constants"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/common/mock"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func Test_RegisterDelegateWithdrawal(t *testing.T) {
	statedb := mock.NewMockStateDB()
	//delegator := common.HexToAddress("0x1000000000000000000000000000000000000001")
	validatorAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	delegatorAddr := common.HexToAddress("0x2222222222222222222222222222222222222222")
	currentEpoch := uint64(1)
	var err error
	err = AppendDelegateWithdrawal(statedb, constants.StakeHandlerAddress, delegatorAddr, validatorAddr, 3, new(big.Int).SetUint64(1000))
	if nil != err {
		t.Error(err)
	}
	err = AppendDelegateWithdrawal(statedb, constants.StakeHandlerAddress, delegatorAddr, validatorAddr, 6, new(big.Int).SetUint64(300))
	if nil != err {
		t.Error(err)
	}

	withdrawableAmount := GetDelegateWithdrawable(statedb, constants.StakeHandlerAddress, delegatorAddr, validatorAddr, currentEpoch)
	//t.Log("withdrawableAmount:", withdrawableAmount)
	assert.Equal(t, withdrawableAmount, common.Big0)

	withdrawalPendingAmount := GetDelegateWithdrawalPending(statedb, constants.StakeHandlerAddress, delegatorAddr, validatorAddr, currentEpoch)
	//t.Log("withdrawalPendingAmount:", withdrawalPendingAmount)
	assert.Equal(t, withdrawalPendingAmount, new(big.Int).SetUint64(1300))

	currentEpoch = 4

	withdrawableAmount = GetDelegateWithdrawable(statedb, constants.StakeHandlerAddress, delegatorAddr, validatorAddr, currentEpoch)
	//t.Log("withdrawableAmount:", withdrawableAmount)
	assert.Equal(t, withdrawableAmount, new(big.Int).SetUint64(1000))

	withdrawalPendingAmount = GetDelegateWithdrawalPending(statedb, constants.StakeHandlerAddress, delegatorAddr, validatorAddr, currentEpoch)
	//t.Log("withdrawalPendingAmount:", withdrawalPendingAmount)
	assert.Equal(t, withdrawalPendingAmount, new(big.Int).SetUint64(300))

}

func Test_AppendValidatorDelegationRc(t *testing.T) {
	statedb := mock.NewMockStateDB()
	validatorAddr := common.HexToAddress("0x1111111111111111111111111111111111111111")
	var err error
	err = AppendValidatorDelegationRc(statedb, constants.StakeHandlerAddress, validatorAddr, 1, 1)
	if nil != err {
		t.Error(err)
	}
	epochs, _ := GetValidatorDelegationRcPendingAndEpoch(statedb, constants.StakeHandlerAddress, validatorAddr, math.MaxUint64)
	//t.Log("epochs", fmt.Sprintf("%v", epochs))
	assert.Equal(t, len(epochs), 1)

	err = AppendValidatorDelegationRc(statedb, constants.StakeHandlerAddress, validatorAddr, 10, 2)
	if nil != err {
		t.Error(err)
	}

	epochs, _ = GetValidatorDelegationRcPendingAndEpoch(statedb, constants.StakeHandlerAddress, validatorAddr, math.MaxUint64)
	//t.Log("epochs", fmt.Sprintf("%v", epochs))
	assert.Equal(t, len(epochs), 2)

}

func Test_MakeSlice(t *testing.T) {
	arr := make([]uint64, 100000)
	t.Log(len(arr))
}
