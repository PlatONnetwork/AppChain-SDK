package constants

import "github.com/PlatONnetwork/PlatON-Go/common"

var (
	StateSenderAddress     = common.HexToAddress("0x1000000000000000000000000000000000000001") // l2 state sender
	StateSyncAddress       = common.HexToAddress("0x1000000000000000000000000000000000000002") // l2 state sync (received l1 state)
	VRFManagerAddress      = common.HexToAddress("0x1000000000000000000000000000000000000003") // vrf manager (management vrf nonce and proof)
	StageManagerAddress    = common.HexToAddress("0x1000000000000000000000000000000000000004") // stage manager (management epoch and round)
	StakeHandlerAddress    = common.HexToAddress("0x1000000000000000000000000000000000000005") // stake handler (handle staking and delegating and slashing)
	RewardManagerAddress   = common.HexToAddress("0x1000000000000000000000000000000000000006") // reward manager (management reward, stake reward of epoch and blocks reward of round)
	DepositHandlerAddress  = common.HexToAddress("0x1000000000000000000000000000000000000007") // deposit handler (dealing with the circulation of assets issued by L1 between L1 and L2)
	WithdrawManagerAddress = common.HexToAddress("0x1000000000000000000000000000000000000008") // TODO ### Not yet enabled !!! ###  withdrwa manager (dealing with the circulation of assets issued by L2 between L2 and L1)
	UpgradeAddress         = common.HexToAddress("0x1000000000000000000000000000000000000009") // upgrade contract address
)
