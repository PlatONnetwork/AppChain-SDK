package contracts

import (
	"math/big"

	typesdk "github.com/PlatONnetwork/AppChain-SDK/types"
	"github.com/PlatONnetwork/AppChain-SDK/x/stage/db"
)

func (c *StageManager) getRoundByBlockNumber(blockNumber uint64) (uint64, error) {

	round, _ := db.GetRoundItemAndIndexByBlockNumber(c.evm.StateDB, c.contract.Address(), blockNumber)

	return round, nil
}

func (c *StageManager) getEpochByBlockNumber(blockNumber uint64) (uint64, error) {

	epoch, _ := db.GetEpochItemAndIndexByBlockNumber(c.evm.StateDB, c.contract.Address(), blockNumber)

	return epoch, nil
}

func (c *StageManager) getPeriodEdgeForRound(round uint64) (*PeriodEdge, error) {
	item := db.GetRoundItem(c.evm.StateDB, c.contract.Address(), round)

	if nil == item {
		return nil, typesdk.NewRevertError("StageManager: INVALID ROUND NUMBER")
	}

	return &PeriodEdge{
		StartBlock: new(big.Int).SetUint64(item.StartBlock),
		EndBlock:   new(big.Int).SetUint64(item.EndBlock),
	}, nil
}

func (c *StageManager) getPeriodEdgeForEpoch(epoch uint64) (*PeriodEdge, error) {
	item := db.GetEpochItem(c.evm.StateDB, c.contract.Address(), epoch)

	if nil == item {
		return nil, typesdk.NewRevertError("StageManager: INVALID EPOCH NUMBER")
	}

	return &PeriodEdge{
		StartBlock: new(big.Int).SetUint64(item.StartBlock),
		EndBlock:   new(big.Int).SetUint64(item.EndBlock),
	}, nil
}

func (c *StageManager) getPeriodEdgesForRound(start, size uint64) (*big.Int, []*big.Int, []PeriodEdge, error) {

	indexs, queue := db.GetRoundQueueAndIndexSince(c.evm.StateDB, c.contract.Address(), start, size)

	if len(queue) == 0 {
		return big.NewInt(0), nil, nil, nil
	}
	periodEdges := make([]PeriodEdge, len(queue))
	rounds := make([]*big.Int, len(queue))
	for i := 0; i < len(queue); i++ {

		rounds[i] = new(big.Int).SetUint64(indexs[i])

		periodEdges[i] = PeriodEdge{
			StartBlock: new(big.Int).SetUint64(queue[i].StartBlock),
			EndBlock:   new(big.Int).SetUint64(queue[i].EndBlock),
		}
	}

	return new(big.Int).SetUint64(indexs[len(indexs)-1] + 1), rounds, periodEdges, nil
}

func (c *StageManager) getPeriodEdgesForEpoch(start, size uint64) (*big.Int, []*big.Int, []PeriodEdge, error) {

	indexs, queue := db.GetEpochQueueAndIndexSince(c.evm.StateDB, c.contract.Address(), start, size)

	if len(queue) == 0 {
		return big.NewInt(0), nil, nil, nil
	}
	periodEdges := make([]PeriodEdge, len(queue))
	epochs := make([]*big.Int, len(queue))

	for i := 0; i < len(queue); i++ {

		epochs[i] = new(big.Int).SetUint64(indexs[i])

		periodEdges[i] = PeriodEdge{
			StartBlock: new(big.Int).SetUint64(queue[i].StartBlock),
			EndBlock:   new(big.Int).SetUint64(queue[i].EndBlock),
		}
	}

	return new(big.Int).SetUint64(indexs[len(indexs)-1] + 1), epochs, periodEdges, nil
}
