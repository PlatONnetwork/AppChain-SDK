package staking

import (
	"github.com/PlatONnetwork/AppChain-SDK/x/address"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/contracts"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

func initStakeHandler(db sdk.StateDB) error {

	if err := initValidatorPriority(db); nil != err {
		return err
	}
	//
	//if err := initUnStakeWithdrawQueueItem(db); nil != err {
	//	return err
	//}

	return nil
}

func initValidatorPriority(db sdk.StateDB) error {
	head := types.NewPriorityValidator(
		contracts.PriorityValidatorTailKey,
		contracts.PriorityValidatorTailKey,
		common.ZeroAddr,
	)
	tail := types.NewPriorityValidator(
		contracts.PriorityValidatorHeadKey,
		contracts.PriorityValidatorHeadKey,
		common.ZeroAddr,
	)
	hvalue, err := rlp.EncodeToBytes(head)
	if nil != err {
		return contracts.ErrRlpEncode
	}
	tvalue, err := rlp.EncodeToBytes(tail)
	if nil != err {
		return contracts.ErrRlpEncode
	}

	db.SetState(address.StakeHandlerAddres, contracts.PriorityValidatorHeadKey, hvalue)
	db.SetState(address.StakeHandlerAddres, contracts.PriorityValidatorTailKey, tvalue)
	return nil
}

//func initUnStakeWithdrawQueueItem(db sdk.StateDB) error {
//
//	return nil
//}
