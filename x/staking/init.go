package staking

import (
	"fmt"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/config"
	stakingdb "github.com/PlatONnetwork/AppChain-SDK/x/staking/db"
	"github.com/PlatONnetwork/AppChain-SDK/x/staking/types"
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/crypto"
	"github.com/PlatONnetwork/PlatON-Go/params"
	"github.com/PlatONnetwork/PlatON-Go/rlp"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
	"math/big"
)

func initStakeConfigParams(statedb sdk.StateDB, addr common.Address, stakeNetworkParams *config.StakeNetworkParams) {

	statedb.SetState(addr, stakingdb.EncodeStakeWithdrawalWaitPeriodKey(), common.Uint64ToBytes(stakeNetworkParams.StakeWithdrawalWaitPeriod))
	statedb.SetState(addr, stakingdb.EncodeDelegateWithdrawalWaitPeriodKey(), common.Uint64ToBytes(stakeNetworkParams.DelegateWithdrawalWaitPeriod))
	statedb.SetState(addr, stakingdb.EncodeSlashingPercentageKey(), common.Uint64ToBytes(stakeNetworkParams.SlashingPercentage))
	statedb.SetState(addr, stakingdb.EncodeSlashIncentivePercentageKey(), common.Uint64ToBytes(stakeNetworkParams.SlashIncentivePercentage))
	statedb.SetState(addr, stakingdb.EncodeMaxRoundValidatorsSizeKey(), common.Uint64ToBytes(stakeNetworkParams.MaxRoundValidatorsSize))
	statedb.SetState(addr, stakingdb.EncodeMaxEpochValidatorsSizeKey(), common.Uint64ToBytes(stakeNetworkParams.MaxEpochValidatorsSize))
	statedb.SetState(addr, stakingdb.EncodeMinRoundValidatorBlockNumberKey(), common.Uint64ToBytes(stakeNetworkParams.MinRoundValidatorBlockNumber))
}

func initValidatorGenesisPriority(statedb sdk.StateDB, addr common.Address) error {
	head := types.NewPriorityValidator(
		stakingdb.EncodePriorityValidatorTailKey(),
		stakingdb.EncodePriorityValidatorTailKey(),
		common.ZeroAddr,
	)
	tail := types.NewPriorityValidator(
		stakingdb.EncodePriorityValidatorHeadKey(),
		stakingdb.EncodePriorityValidatorHeadKey(),
		common.ZeroAddr,
	)
	hvalue, err := rlp.EncodeToBytes(head)
	if nil != err {
		return fmt.Errorf("rlp encode head validator priority %s", err)
	}
	tvalue, err := rlp.EncodeToBytes(tail)
	if nil != err {
		return fmt.Errorf("rlp encode tail validator priority %s", err)
	}

	statedb.SetState(addr, stakingdb.EncodePriorityValidatorHeadKey(), hvalue)
	statedb.SetState(addr, stakingdb.EncodePriorityValidatorTailKey(), tvalue)
	return nil
}

func initValidators(statedb sdk.StateDB, addr common.Address, chainConfig *params.ChainConfig, stakeNetworkParams *config.StakeNetworkParams) error {

	if err := initValidatorGenesisPriority(statedb, addr); nil != err {
		return err
	}

	genesisValidatorQueueSize := uint64(len(chainConfig.Cbft.InitialNodes))

	if stakeNetworkParams.MaxRoundValidatorsSize <= uint64(len(chainConfig.Cbft.InitialNodes)) {
		genesisValidatorQueueSize = stakeNetworkParams.MaxRoundValidatorsSize
	} else {
		genesisValidatorQueueSize = uint64(len(chainConfig.Cbft.InitialNodes))
	}

	initialNodeQueue := chainConfig.Cbft.InitialNodes

	validatorShareSnapshotQueue := types.NewValidatorSharesSnapshotQueue(0)

	genesisStakeAmount := new(big.Int).SetUint64(stakeNetworkParams.GenesisStakeAmount)
	genesisDelegateAmount := common.Big0

	cache := make(map[common.Address]struct{}, 0)
	for index := uint64(0); index < genesisValidatorQueueSize; index++ {

		initialNode := initialNodeQueue[index]

		pubKey := initialNode.Node.Pubkey()
		validatorAddr := crypto.PubkeyToAddress(*pubKey)

		if _, ok := cache[validatorAddr]; ok {
			continue
		} else {
			cache[validatorAddr] = struct{}{}
		}

		stakeIndex := stakingdb.IncrementValidatorNonce(statedb, addr)

		validator := types.NewValidator(
			stakeNetworkParams.GenesisValidatorOwner, genesisStakeAmount, genesisDelegateAmount,
			initialNode.BlsPubKey.Serialize(), initialNode.Node.IDv0(), 0, 1, stakeIndex)
		if err := stakingdb.SetValidator(statedb, addr, validatorAddr, validator); nil != err {
			return fmt.Errorf("set validator info '%s' %s", validatorAddr, err)
		}

		if err := stakingdb.SetValidatorPriority(statedb, addr, validatorAddr, 1, stakeIndex, genesisStakeAmount); nil != err {
			return fmt.Errorf("set validator priority '%s' %s", validatorAddr, err)
		}

		validatorShareSnapshot := types.NewValidatorSharesSnapshot(validatorAddr, 1, stakeIndex, genesisStakeAmount, genesisDelegateAmount)
		validatorShareSnapshotQueue = append(validatorShareSnapshotQueue, validatorShareSnapshot)
	}

	var (
		roundValidatorSnapshotQueue types.ValidatorSortSnapshotQueue
		epochValidatorSnapshotQueue types.ValidatorSortSnapshotQueue
	)

	if uint64(len(validatorShareSnapshotQueue)) > stakeNetworkParams.MaxRoundValidatorsSize {
		roundValidatorSnapshotQueue = validatorShareSnapshotQueue[:stakeNetworkParams.MaxRoundValidatorsSize]
	} else {
		roundValidatorSnapshotQueue = validatorShareSnapshotQueue
	}

	// set round validator queue
	if err := stakingdb.SetRoundValidatorSharesSnapshotQueue(statedb, addr, 0, roundValidatorSnapshotQueue); nil != err {
		return fmt.Errorf("set genesis validatorQueue for round %d, %s", 0, err)
	}
	if err := stakingdb.SetRoundValidatorSharesSnapshotQueue(statedb, addr, 1, roundValidatorSnapshotQueue); nil != err {
		return fmt.Errorf("set genesis validatorQueue for round %d, %s", 1, err)
	}

	if uint64(len(validatorShareSnapshotQueue)) > stakeNetworkParams.MaxEpochValidatorsSize {
		epochValidatorSnapshotQueue = validatorShareSnapshotQueue[:stakeNetworkParams.MaxEpochValidatorsSize]
	} else {
		epochValidatorSnapshotQueue = validatorShareSnapshotQueue
	}

	// set epoch validator queue
	if err := stakingdb.SetEpochValidatorSharesSnapshotQueue(statedb, addr, 1, epochValidatorSnapshotQueue); nil != err {
		return fmt.Errorf("set genesis validatorQueue for epoch %d, %s", 1, err)
	}

	return nil
}
