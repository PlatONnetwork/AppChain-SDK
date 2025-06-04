package types

import "github.com/PlatONnetwork/PlatON-Go/sdk"

type ContractsApp interface {
	Contracts(sdk.StateDBReader, uint64) []sdk.SDKContract
}
