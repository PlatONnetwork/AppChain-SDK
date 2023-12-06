package types

import (
	"crypto/ecdsa"
	basecommon "github.com/PlatONnetwork/PlatON-Go/common"
	"github.com/PlatONnetwork/PlatON-Go/sdk"
)

type Stage interface {
	GetCurrentRound(stateDB sdk.StateDBReader) uint64
	GetCurrentEpoch(stateDB sdk.StateDBReader) uint64
}

type Stake interface {
	IsValidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool
	IsInvalidValidator(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) bool
	GetValidatorECDSAPubKey(stateDB sdk.StateDBReader, validatorAddr basecommon.Address) *ecdsa.PublicKey
}
