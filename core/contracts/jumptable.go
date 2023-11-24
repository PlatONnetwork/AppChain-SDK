package contracts

import (
	"github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/common/math"
	"github.com/PlatONnetwork/PlatON-Go/core/types"
	"github.com/PlatONnetwork/PlatON-Go/core/vm"
)

type operation struct {
	constantGas uint64
	dynamicGas  costGas
}

type JumpTable [256]*operation

func newJumpTable() JumpTable {
	return JumpTable{
		CREATEACCOUNT: &operation{
			constantGas: GasCreateAccount,
		},
		SUBBALANCE: &operation{
			constantGas: GasBalance,
		},
		ADDBALANCE: &operation{
			constantGas: GasBalance,
		},
		GETBALANCE: &operation{
			constantGas: GasBalance,
		},
		GETNONCE: &operation{
			constantGas: GasQuickStep,
		},
		SETNONCE: &operation{
			constantGas: GasQuickStep,
		},
		GETCODEHASH: &operation{
			constantGas: GasQuickStep,
		},
		GETCODE: &operation{
			constantGas: GasQuickStep,
		},
		SETCODE: &operation{
			constantGas: GasQuickStep,
			dynamicGas:  copyGas(),
		},
		GETCODESIZE: &operation{
			constantGas: GasQuickStep,
		},
		GETSTATE: &operation{
			constantGas: GasSLoad,
		},
		SETSTATE: &operation{
			dynamicGas: setStateGas(),
		},
		ADDLOG: &operation{
			constantGas: GasLog,
			dynamicGas:  addLogGas(),
		},
	}
}

type costGas func(params ...interface{}) uint64

func copyGas() costGas {
	return func(params ...interface{}) uint64 {
		return toWordSize(params[0].(uint64)) * GasCopy
	}

}

func setStateGas() costGas {
	return func(params ...interface{}) uint64 {
		evm := params[0].(*vm.EVM)
		address := params[1].(common.Address)
		key := params[2].([]byte)
		value := params[3].([]byte)

		currentValue := evm.StateDB.GetState(address, key)
		oldWordSize := toWordSize(uint64(len(key)) + uint64(len(currentValue)))
		newWordSize := toWordSize(uint64(len(key)) + uint64(len(value)))

		switch {
		case 0 == len(currentValue) && 0 != len(value):
			return newWordSize * SstoreSetGas
		case 0 != len(currentValue) && 0 == len(value):
			evm.StateDB.AddRefund(oldWordSize * SstoreRefundGas)
			return oldWordSize * SstoreClearGas
		default:
			var (
				addWordSize    uint64 = 0
				deleteWordSize uint64 = 0
				resetWordSize  uint64 = 0
			)

			if newWordSize >= oldWordSize {
				addWordSize = newWordSize - oldWordSize
				resetWordSize = toWordSize(uint64(len(currentValue)))
			} else {
				deleteWordSize = oldWordSize - newWordSize
				resetWordSize = toWordSize(uint64(len(value)))
			}

			if 0 == resetWordSize {
				resetWordSize = 1
			}

			evm.StateDB.AddRefund(deleteWordSize * SstoreRefundGas)
			return addWordSize*SstoreSetGas + deleteWordSize*SstoreClearGas + resetWordSize*SstoreResetGas

		}
		return 0
	}
}

func addLogGas() costGas {
	return func(params ...interface{}) uint64 {
		log := params[0].(*types.Log)
		return logGas(uint64(len(log.Topics)), uint64(len(log.Data)))
	}
}

func logGas(topicNum, dataSize uint64) uint64 {
	gas := topicNum * GasLogTopicGas
	gas += dataSize * GasLogData
	return gas
}

func toWordSize(size uint64) uint64 {
	if size > math.MaxUint64-31 {
		return math.MaxUint64/32 + 1
	}

	return (size + 31) / 32
}
