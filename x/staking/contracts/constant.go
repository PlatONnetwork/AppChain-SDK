package contracts

// todo maybe be governanced
var (
	STAKE_WITHDRAWAL_WAIT_PERIOD    = uint64(2)
	DELEGATE_WITHDRAWAL_WAIT_PERIOD = uint64(2)
	SLASHING_PERCENTAGE             = uint64(50) // to be read through NetworkParams later
	SLASH_INCENTIVE_PERCENTAGE      = uint64(30) // exitor reward, to be read through NetworkParams later
)
