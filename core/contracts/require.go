package contracts

import "github.com/PlatONnetwork/AppChain-SDK/types"

func Require(cond bool, msg string) {
	if !cond {
		panic(types.NewRevertError(msg))
	}
}
