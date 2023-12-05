package constants

import "github.com/PlatONnetwork/PlatON-Go/common"

var (
	StateSenderAddress              = common.HexToAddress("")
	StateReceiverAddress            = common.HexToAddress("")
	RootchainStakeManagerAddress    = common.HexToAddress("") // CustomChildChainManager constants on rootchain
	StakeHandlerAddress             = common.HexToAddress("")
	RootchainDepositManagerAddress  = common.HexToAddress("") // DepositManager constants on rootchain
	DepositHandlerAddress           = common.HexToAddress("")
	RootchainWithdrawHandlerAddress = common.HexToAddress("") // WithdrawHandler constants on rootchain
	WithdrawManagerAddress          = common.HexToAddress("")
	RewardManagerAddress            = common.HexToAddress("")
	StageManagerAddress             = common.HexToAddress("") // management epoch/round ...
	VRFHandlerAddress               = common.HexToAddress("")
)
