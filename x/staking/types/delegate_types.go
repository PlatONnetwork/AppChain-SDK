package types

import "math/big"

type Delegation struct {
	delegateBlock uint64
	delegateShaes *big.Int
}
