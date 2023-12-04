package address

import "github.com/PlatONnetwork/PlatON-Go/common"

var (
	StateSenderAddress              = common.HexToAddress("")
	StateReceiverAddress            = common.HexToAddress("")
	RootchainStakeManagerAddress    = common.HexToAddress("") // CustomChildChainManager address on rootchain
	StakeHandlerAddress             = common.HexToAddress("")
	RootchainDepositManagerAddress  = common.HexToAddress("") // DepositManager address on rootchain
	DepositHandlerAddress           = common.HexToAddress("")
	RootchainWithdrawHandlerAddress = common.HexToAddress("") // WithdrawHandler address on rootchain
	WithdrawManagerAddress          = common.HexToAddress("")
	RewardManagerAddress            = common.HexToAddress("")
	VRFHandlerAddress               = common.HexToAddress("")
)
