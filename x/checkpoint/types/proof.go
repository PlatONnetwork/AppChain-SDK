package types

import "github.com/PlatONnetwork/PlatON-Go/common"

type Proof struct {
	Data     []common.Hash
	Metadata map[string]interface{}
}
