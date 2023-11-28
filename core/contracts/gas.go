package contracts

const (
	GasQuickStep uint64 = 2

	GasCreateAccount uint64 = 100
	GasBalance       uint64 = 20
	GasSLoad         uint64 = 200
	GasCopy          uint64 = 1
	GasLog           uint64 = 375
	GasLogTopicGas   uint64 = 375
	GasLogData       uint64 = 8

	SstoreSetGas    uint64 = 20000 // Once per SLOAD operation.
	SstoreResetGas  uint64 = 5000  // Once per SSTORE operation if the zeroness changes from zero.
	SstoreClearGas  uint64 = 5000  // Once per SSTORE operation if the zeroness doesn't change.
	SstoreRefundGas uint64 = 15000 // Once per SSTORE operation if the zeroness changes to zero.
)
