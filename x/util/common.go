package util

import (
	"bytes"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	basetypes "github.com/PlatONnetwork/PlatON-Go/core/types"
)

func isWorker(header *basetypes.Header) bool {
	return len(header.Extra) > 32 && len(header.Extra[32:]) >= basecommon.ExtraSeal && bytes.Equal(header.Extra[32:97], make([]byte, basecommon.ExtraSeal))
}

func IsNotWorker(header *basetypes.Header) bool {
	return !isWorker(header)
}
